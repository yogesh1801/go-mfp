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
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"sync"

	"github.com/OpenPrinting/goipp"
	"github.com/alexpevzner/mfp/log"
	"github.com/alexpevzner/mfp/transport"
)

// proxy implements an IPP/eSCL/WSD proxy
type proxy struct {
	ctx       context.Context   // Logging/shutdown context
	cancel    func()            // ctx cancel function
	m         mapping           // Local/remote mapping
	l         net.Listener      // TCP listener for incoming connections
	srv       *transport.Server // HTTP server for incoming connections
	clnt      *transport.Client // HTTP client part of proxy
	closeWait sync.WaitGroup    // Wait for proxy.Close completion
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
		clnt:   transport.NewClient(nil),
	}

	// Ensure cancellation propagation
	p.closeWait.Add(1)
	go p.kill()

	// Start HTTP server
	p.srv = transport.NewServer(nil, p)

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
func (p *proxy) ServeHTTP(w http.ResponseWriter, in *http.Request) {
	// Catch panics to log
	defer func() {
		v := recover()
		if v != nil {
			log.Panic(p.ctx, v)
		}
	}()

	// Handle request
	log.Debug(p.ctx, "%s %s", in.Method, in.URL)

	ct := strings.ToLower(in.Header.Get("Content-Type"))

	switch {
	case p.m.proto == protoIPP && in.Method == "POST" &&
		ct == "application/ipp":
		p.doIPP(w, in)

	default:
		p.httpReject(w, in,
			http.StatusBadRequest, errors.New("Bad Request"))
	}
}

// outreq creates an outgoing HTTP request based on request
// received by the server side of proxy.
func (p *proxy) outreq(in *http.Request, body io.ReadCloser) *http.Request {
	// Create request
	out, _ := transport.NewRequest(p.ctx, in.Method, in.URL, body)
	out.Header = in.Header.Clone()
	p.httpRemoveHopByHopHeaders(out.Header)

	// Adjust target URL
	prq := httputil.ProxyRequest{
		Out: out,
	}
	prq.SetURL(p.m.targetURL)
	out.Host = out.URL.Host

	return out
}

// msgxlat returns goipp.Message translator that rewrites message
// attributes when message is being forwarded via proxy.
//
// Currently, only URLs embedded into the message are translated.
func (p *proxy) msgxlat(in *http.Request) (*msgXlat, error) {
	s := "http://" + in.Host
	u, err := transport.ParseURL(s)
	if err != nil {
		err = fmt.Errorf("%q: can't parse local URL")
		return nil, err
	}

	urlxlat := transport.NewURLXlat(u, p.m.targetURL)
	msgxlat := newMsgXlat(urlxlat)

	return msgxlat, nil
}

// doIPP handles incoming IPP requests
func (p *proxy) doIPP(w http.ResponseWriter, in *http.Request) {
	ops := goipp.DecoderOptions{EnableWorkarounds: true}

	// Create goipp.Message translator
	msgxlat, err := p.msgxlat(in)
	if err != nil {
		p.httpReject(w, in, 503, err)
	}

	// Dump input request
	dump, _ := httputil.DumpRequest(in, false)
	log.Debug(p.ctx, "IPP: request received:")
	log.Debug(p.ctx, "%s", dump)

	// Fetch IPP Request message
	ibody := transport.NewPeeker(in.Body)
	var msg goipp.Message
	err = msg.DecodeEx(ibody, ops)
	if err != nil {
		p.httpReject(w, in, 503, fmt.Errorf("IPP error: %w", err))
		return
	}

	var buf bytes.Buffer
	msg.Print(&buf, true)
	log.Debug(p.ctx, buf.String())

	msg2 := msgxlat.Forward(&msg)
	buf.Reset()
	msg2.Print(&buf, true)
	log.Debug(p.ctx, "IPP: request translated:")
	log.Debug(p.ctx, buf.String())

	// Setup and execute outgoing request
	ibody.Rewind()
	out := p.outreq(in, ibody)
	out.ContentLength = in.ContentLength

	log.Debug(p.ctx, "IPP: forward request to: %s", out.URL)

	rsp, err := p.clnt.Do(out)
	if err != nil {
		log.Debug(p.ctx, "IPP: %s", err)
		p.httpReject(w, in, http.StatusBadGateway, err)
		return
	}

	dump, _ = httputil.DumpResponse(rsp, false)
	log.Debug(p.ctx, "IPP: response received:")
	log.Debug(p.ctx, "%s", dump)

	// Fetch IPP response message
	obody := transport.NewPeeker(rsp.Body)
	msg.Reset()
	err = msg.DecodeEx(obody, ops)
	if err != nil {
		obody.Rewind()
		data, _ := io.ReadAll(obody)
		println(err.Error())
		log.Debug(p.ctx, "%2.2x", data)
		p.httpReject(w, in, 503, errors.New("oops"))
		return
	}

	msg.Print(&buf, false)
	log.Debug(p.ctx, buf.String())

	msg2 = msgxlat.Reverse(&msg)
	buf.Reset()
	msg2.Print(&buf, true)
	log.Debug(p.ctx, "IPP: response translated:")
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
func (p *proxy) httpReject(w http.ResponseWriter, in *http.Request,
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
