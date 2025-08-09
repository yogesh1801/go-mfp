// MFP - Miulti-Function Printers and scanners toolkit
// The "proxy" command
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Package documentation

package proxy

import (
	"archive/tar"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/OpenPrinting/go-mfp/log"
	"github.com/OpenPrinting/goipp"
)

// traceWriter writes a protocol trace
type traceWriter struct {
	ctx      context.Context // Logging context
	name     string          // file name
	fp       *os.File        // Underlying file
	tar      *tar.Writer     // TAR writer
	lock     sync.Mutex      // Access lock
	err      error           // First error
	donewait sync.WaitGroup  // Wait for async activities
}

// newTrace creates a new trace writer
func newTraceWriter(ctx context.Context, name string) (*traceWriter, error) {
	nameLog := name + ".log"
	nameTar := name + ".tar"

	// Create name.log
	os.Remove(nameLog)
	backend := log.NewFileBackend(nameLog, 0, 0)
	log.CtxLogger(ctx).Attach(log.LevelTrace, backend)

	// Create name.tar
	const flags = os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	fp, err := os.OpenFile(nameTar, flags, 0644)
	if err != nil {
		return nil, err
	}

	trace := &traceWriter{
		ctx:  ctx,
		name: name,
		fp:   fp,
		tar:  tar.NewWriter(fp),
	}

	return trace, nil
}

// Close closes the traceWriter
func (trace *traceWriter) Close() {
	trace.donewait.Wait()

	trace.lock.Lock()
	defer trace.lock.Unlock()

	err := trace.tar.Close()
	if err != nil {
		trace.setError(err)
	}
	err = trace.fp.Close()
	if err != nil {
		trace.setError(err)
	}
}

// IPPRequest is the [ipp.Sniffer.Request] callback.
func (trace *traceWriter) IPPRequest(seqnum uint64,
	rq *http.Request, msg *goipp.Message, body io.Reader) {

	name := fmt.Sprintf("%8.8d/00-%s.ipp",
		seqnum, goipp.Op(msg.Code))

	data, _ := msg.EncodeBytes()
	trace.Send(name, data)

	trace.donewait.Add(1)
	go func() {
		data, _ := io.ReadAll(body)

		if len(data) != 0 {
			name := fmt.Sprintf("%8.8d/01-odata.%s",
				seqnum, magic(data))

			trace.Send(name, data)
		}

		trace.donewait.Done()
	}()
}

// IPPResponse is the [ipp.Sniffer.Response] callback.
func (trace *traceWriter) IPPResponse(seqnum uint64,
	rsp *http.Response, msg *goipp.Message, body io.Reader) {

	name := fmt.Sprintf("%8.8d/02-%s.ipp",
		seqnum, goipp.Status(msg.Code))

	data, _ := msg.EncodeBytes()
	trace.Send(name, data)

	trace.donewait.Add(1)
	go func() {
		data, _ := io.ReadAll(body)

		if len(data) != 0 {
			name := fmt.Sprintf("%8.8d/03-rdata.%s",
				seqnum, magic(data))

			trace.Send(name, data)
		}

		trace.donewait.Done()
	}()
}

// Send writes a new record (a file) into the trace archive.
func (trace *traceWriter) Send(name string, data []byte) {
	trace.lock.Lock()
	defer trace.lock.Unlock()

	log.Debug(trace.ctx, "%s: %d bytes saved", name, len(data))

	hdr := tar.Header{
		Typeflag: tar.TypeReg,
		Name:     name,
		Size:     int64(len(data)),
		Mode:     0644,
		ModTime:  time.Now(),
		Devmajor: 1,
		Devminor: 1,
	}

	if trace.err == nil {
		err := trace.tar.WriteHeader(&hdr)
		if err != nil {
			trace.setError(err)
		}
	}

	if trace.err == nil {
		_, err := trace.tar.Write(data)
		if err != nil {
			trace.setError(err)
		}
	}

	if trace.err == nil {
		err := trace.tar.Flush()
		if err != nil {
			trace.setError(err)
		}
	}
}

// setError sets trace.err, when error occurs for the first time.
// When it happens, the event is logged.
//
// This function must be called under trace.lock
func (trace *traceWriter) setError(err error) {
	if trace.err == nil {
		trace.err = err
		log.Error(trace.ctx, "%s: %s", trace.name, err)
	}
}
