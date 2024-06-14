// MFP       - Miulti-Function Printers and scanners toolkit
// TRANSPORT - Transport protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Loopback transport

package transport

import (
	"context"
	"errors"
	"net"
	"net/http"
	"sync"
)

// Loopback limits:
const (
	// LoopbackMaxPendingConnections defines a limit of pending
	// connections (attempted by Client but not yet accepted by Server).
	LoopbackMaxPendingConnections = 1024
)

// Loopback errors:
var (
	ErrLoopbackClosed = errors.New("Loopback: closed")
	ErrLoopbackBusy   = errors.New("Loopback: too many dial attempts")
)

// NewLoopback creates a new [Transport] and corresponding [net.Listener].
//
// Every connection created by the Transport appears as incoming
// connection on the Listener. It is suitable for use with the
// [http.Server.Serve] method of the [http.Server].
//
// When Listener is being closed, all subsequent Accept() calls and
// all subsequent attempts to establish a new connection fails with
// [ErrLoopbackClosed] error.
//
// If there are too many pending connections, an attempt to establish
// a new connection may fail with [ErrLoopbackBusy] error.
//
// The primary purpose of this functionality is client/server testing.
func NewLoopback() (*Transport, net.Listener) {
	// Create loopback
	l := &loopback{
		conns: make(chan net.Conn, LoopbackMaxPendingConnections),
	}

	// Create a Transport
	template := (http.DefaultTransport.(*http.Transport)).Clone()
	template.DialContext = l.dial
	tr := NewTransport(template)

	return tr, l
}

// loopback implements underlying low-level transport for the
// Loopback() function.
type loopback struct {
	lock  sync.Mutex
	conns chan net.Conn
}

// dial establishes a new connection.
func (l *loopback) dial(ctx context.Context, network, addr string) (net.Conn, error) {
	// Acquire loopback lock
	l.lock.Lock()
	defer l.lock.Unlock()

	// Crete piped connection and try to push one into the listener's queue
	c1, c2 := net.Pipe()
	err := ErrLoopbackClosed
	if l.conns != nil {
		select {
		case l.conns <- c2:
			return c1, nil
		default:
			err = ErrLoopbackBusy
		}
	}

	c1.Close()
	c2.Close()

	return nil, err

}

// Accept waits for and returns the next connection to the listener.
func (l *loopback) Accept() (net.Conn, error) {
	l.lock.Lock()
	conns := l.conns
	l.lock.Unlock()

	if conns == nil {
		// If we are here, loopback was closed before us
		return nil, ErrLoopbackClosed
	}

	conn := <-conns
	if conn == nil {
		// If we are here, loopback was closed while we were waiting
		return nil, ErrLoopbackClosed
	}

	return conn, nil
}

// Close closes the listener.
func (l *loopback) Close() error {
	// Steal the connections queue
	l.lock.Lock()
	conns := l.conns
	l.conns = nil
	l.lock.Unlock()

	// Close and purge connections queue
	err := ErrLoopbackClosed
	if conns != nil {
		close(conns)
		for conn := range conns {
			conn.Close()
		}

		err = nil
	}

	return err
}

// Addr returns the listener's network address.
func (*loopback) Addr() net.Addr {
	return &net.UnixAddr{
		Name: "loopback",
		Net:  "unix",
	}
}
