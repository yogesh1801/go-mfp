// MFP - Miulti-Function Printers and scanners toolkit
// The "proxy" command
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Package documentation

package proxy

import (
	"context"
	"net"
	"net/http"
	"sync"

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

	// Create HTTP server
	p.srv = &http.Server{}
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
