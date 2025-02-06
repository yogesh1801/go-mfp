// MFP - Miulti-Function Printers and scanners toolkit
// The "proxy" command
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Package documentation

package proxy

import (
	"bytes"
	"context"
	"errors"
	"net"
	"net/http"
	"strings"
	"sync"

	"github.com/OpenPrinting/goipp"
	"github.com/alexpevzner/mfp/log"
)

// proxy implements an IPP/eSCL/WSD proxy
type proxy struct {
	ctx       context.Context // Logging/shutdown context
	cancel    func()          // ctx cancel function
	m         mapping         // Local/remote mapping
	l         net.Listener    // TCP listener for incoming connections
	srv       *http.Server    // HTTP server for incoming connections
	closeWait sync.WaitGroup  // Wait for proxy.Close completion
}

// newProxy creates a new proxy for the specified mapping.
func newProxy(ctx context.Context, m mapping) (*proxy, error) {
	log.Debug(ctx, "proxy started: %d->%s", m.localPort, m.targetURL)

	// Create TCP listener
	l, err := newListener(ctx, m.localPort)
	if err != nil {
		return nil, err
	}

	// Create cancelable context
	ctx, cancel := context.WithCancel(ctx)

	// Create proxy structure
	p := &proxy{
		ctx:    ctx,
		cancel: cancel,
		m:      m,
		l:      l,
	}

	// Ensure cancellation propagation
	p.closeWait.Add(1)
	go p.kill()

	// Create HTTP server
	p.srv = &http.Server{
		Handler: p,
	}

	p.closeWait.Add(1)
	go func() {
		p.srv.Serve(l)
		p.closeWait.Done()
	}()

	return p, nil
}

// kill closes the proxy and terminates all active session when proxy.ctx
// is canceled.
func (p *proxy) kill() {
	<-p.ctx.Done()

	p.srv.Close()

	p.closeWait.Done()
}

// Shutdown performs proxy shutdown.
func (p *proxy) Shutdown() {
	p.cancel()
	p.closeWait.Wait()

	log.Debug(p.ctx, "proxy finished: %d->%s",
		p.m.localPort, p.m.targetURL)
}

// ServeHTTP handles incoming HTTP requests.
// It implements [http.Handler] interface.
func (p *proxy) ServeHTTP(w http.ResponseWriter, rq *http.Request) {
	// Catch panics to log
	defer func() {
		v := recover()
		if v != nil {
			log.Panic(p.ctx, v)
		}
	}()

	// Handle request
	log.Debug(p.ctx, "%s %s", rq.Method, rq.URL)

	ct := strings.ToLower(rq.Header.Get("Content-Type"))

	switch {
	case p.m.proto == protoIPP && rq.Method == "POST" &&
		ct == "application/ipp":
		p.doIPP(w, rq)

	default:
		p.httpReject(w, rq,
			http.StatusBadRequest, errors.New("Bad Request"))
	}
}

// doIPP handles incoming IPP requests
func (p *proxy) doIPP(w http.ResponseWriter, rq *http.Request) {
	ops := goipp.DecoderOptions{EnableWorkarounds: true}

	// Fetch IPP Request message
	var msg goipp.Message
	err := msg.DecodeEx(rq.Body, ops)
	if err != nil {
		p.httpReject(w, rq, 503, errors.New("oops"))
		return
	}

	var buf bytes.Buffer
	msg.Print(&buf, true)
	log.Debug(p.ctx, buf.String())
}

// httpRemoveHopByHopHeaders removes HTTP hop-by-hop headers,
// RFC 7230, section 6.1
func (p *proxy) httpRemoveHopByHopHeaders(hdr http.Header) {
	// Per RFC 7230, section 6.1:
	//
	// Hence, the Connection header field provides a declarative way of
	// distinguishing header fields that are only intended for the immediate
	// recipient ("hop-by-hop") from those fields that are intended for all
	// recipients on the chain ("end-to-end"), enabling the message to be
	// self-descriptive and allowing future connection-specific extensions
	// to be deployed without fear that they will be blindly forwarded by
	// older intermediaries.
	if c := hdr.Get("Connection"); c != "" {
		for _, f := range strings.Split(c, ",") {
			if f = strings.TrimSpace(f); f != "" {
				hdr.Del(f)
			}
		}
	}

	// These headers are always considered hop-by-hop.
	for _, c := range []string{"Connection", "Keep-Alive",
		"Proxy-Authenticate", "Proxy-Connection",
		"Proxy-Authorization", "Te", "Trailer", "Transfer-Encoding"} {
		hdr.Del(c)
	}
}

// httpReject completes request with a error
func (p *proxy) httpReject(w http.ResponseWriter, rq *http.Request,
	status int, err error) {

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	p.httpNoCache(w)
	w.WriteHeader(status)

	w.Write([]byte(err.Error()))
	w.Write([]byte("\n"))
}

// httpNoCache set response headers to disable client-side
// response cacheing.
func (p *proxy) httpNoCache(w http.ResponseWriter) {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
}
