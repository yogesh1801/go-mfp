// MFP - Multi-Function Printers and scanners toolkit
// Virtual USB/IP device emulator for testing and fuzzing
//
// Copyright (C) 2025 and up by GO-MFP authors.
// See LICENSE for license terms and conditions
//
// IPP-USB device implementation

package main

import (
	"net"
	"net/http"
	"sync/atomic"
	"time"
)

// ippusbAddr returned as local and remote address for the IPP over USB connections.
var ippusbAddr = &net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 80, Zone: ""}

// NewIPPUSB creates a new IPP over USB server (represented by the
// [http.Server] and its [Endpoint]s
func NewIPPUSB(numendpoints int) (*http.Server, []*Endpoint) {
	endpoints := make([]*Endpoint, numendpoints)
	for i := range endpoints {
		endpoints[i] = NewEndpoint(EndpointInOut, USBXferBulk, 512)
	}

	srv := &http.Server{}
	listener := newIppusbListener(endpoints)

	srv.Handler = http.HandlerFunc(handler)
	go srv.Serve(listener)

	return srv, endpoints
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
}

// ippusbListener implements the [net.Listener] interface on a top of
// the group of [Endpoint]s
type ippusbListener struct {
	endpoints chan *Endpoint // Queue of available endpoints
	closechan chan struct{}  // Closed by ippusbListener.Close
}

// newIppusbListener returns the new ippusbListener
func newIppusbListener(endpoints []*Endpoint) *ippusbListener {
	listener := &ippusbListener{
		endpoints: make(chan *Endpoint, len(endpoints)),
		closechan: make(chan struct{}),
	}

	for _, ep := range endpoints {
		listener.endpoints <- ep
	}

	return listener
}

// Accept waits for and returns the next connection to the listener.
func (listener *ippusbListener) Accept() (net.Conn, error) {
	select {
	case ep := <-listener.endpoints:
		conn := &ippusbConn{Endpoint: ep}
		conn.listener.Store(listener)
		return conn, nil
	case <-listener.closechan:
		return nil, net.ErrClosed
	}
}

// Close closes the listener.
// Any blocked Accept operations will be unblocked and return errors.
func (listener *ippusbListener) Close() error {
	close(listener.closechan)
	return nil
}

// Addr returns the listener's network address.
func (listener *ippusbListener) Addr() net.Addr {
	return ippusbAddr
}

// ippusbEndpoint wraps the Endpoint and adds stub implementation of the
// missed methods to implement the net.Conn interface
type ippusbConn struct {
	*Endpoint                                // Underlying Endpoint
	listener  atomic.Pointer[ippusbListener] // Parent ippusbListener
}

// Close closes the connection.
// Any blocked Read or Write operations will be unblocked and return errors.
func (conn *ippusbConn) Close() error {
	// Double-close protection
	listener := conn.listener.Swap(nil)
	if listener != nil {
		listener.endpoints <- conn.Endpoint
	}

	return nil
}

// LocalAddr returns the local network address, if known.
func (conn *ippusbConn) LocalAddr() net.Addr {
	return ippusbAddr
}

// RemoteAddr returns the remote network address, if known.
func (conn *ippusbConn) RemoteAddr() net.Addr {
	return ippusbAddr
}

// SetDeadline sets the read and write deadlines associated
// with the connection.
func (conn *ippusbConn) SetDeadline(t time.Time) error {
	return nil
}

// SetReadDeadline sets the deadline for future Read calls
// and any currently-blocked Read call.
func (conn *ippusbConn) SetReadDeadline(t time.Time) error {
	return nil
}

// SetWriteDeadline sets the deadline for future Write calls
// and any currently-blocked Write call.
func (conn *ippusbConn) SetWriteDeadline(t time.Time) error {
	return nil
}
