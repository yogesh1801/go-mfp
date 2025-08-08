// MFP - Multi-Function Printers and scanners toolkit
// Virtual USB/IP device emulator for testing and fuzzing
//
// Copyright (C) 2025 and up by GO-MFP authors.
// See LICENSE for license terms and conditions
//
// USBIP protocol connection

package usbip

import (
	"context"
	"net"
	"sync/atomic"
	"time"
)

// EndpointConn implements the [net.Conn] interface on a top of [Endpoint].
//
// This is suitable for usage with the [http.Server] to implement the
// IPP over USB protocol or for the similar tasks.
//
// Data, that was written into the endpoint from the USB side appears
// as data received from the connection and visa versa.
type EndpointConn struct {
	ep            *Endpoint          // Underlying endpoint
	local, remote net.Addr           // Local and Remote addresses
	closectx      context.Context    // Canceled by Close()
	closecancel   context.CancelFunc // Cancels closectx
	listener      atomic.Pointer[    // Parent listener, nil if none
	EndpointListener]
}

// NewEndpointConn creates the [EndpointConn] on a top of the existent
// [Endpoint].
//
// The local and remote addresses are only used to implement the
// [net.Conn.LocalAddr] and [net.Conn.RemoteAddr] interfaces.
func NewEndpointConn(ep *Endpoint, local, remote net.Addr) *EndpointConn {
	conn := &EndpointConn{
		ep:     ep,
		local:  local,
		remote: remote,
	}

	ctx := context.Background()
	conn.closectx, conn.closecancel = context.WithCancel(ctx)

	return conn
}

// Close closes the connection and unblocks all pending Reads and Writes.
func (conn *EndpointConn) Close() error {
	conn.closecancel()
	return nil
}

// Read received data from the underlying [Endpoint].
func (conn *EndpointConn) Read(buf []byte) (int, error) {
	n, err := conn.ep.ReadContext(conn.closectx, buf)
	if err == context.Canceled {
		err = net.ErrClosed
	}
	return n, err
}

// Write sends data into the underlying [Endpoint].
func (conn *EndpointConn) Write(buf []byte) (int, error) {
	n, err := conn.ep.WriteContext(conn.closectx, buf)
	if err == context.Canceled {
		err = net.ErrClosed
	}
	return n, err
}

// LocalAddr returns the local network address, if known.
func (conn *EndpointConn) LocalAddr() net.Addr {
	return conn.local
}

// RemoteAddr returns the remote network address, if known.
func (conn *EndpointConn) RemoteAddr() net.Addr {
	return conn.remote
}

// SetDeadline sets the read and write deadlines associated
// with the connection.
func (conn *EndpointConn) SetDeadline(t time.Time) error {
	return nil
}

// SetReadDeadline sets the deadline for future Read calls
// and any currently-blocked Read call.
func (conn *EndpointConn) SetReadDeadline(t time.Time) error {
	return nil
}

// SetWriteDeadline sets the deadline for future Write calls
// and any currently-blocked Write call.
func (conn *EndpointConn) SetWriteDeadline(t time.Time) error {
	return nil
}

// EndpointListener implements the [net.Listener] interface on
// a top of the group of [Endpoint]s.
//
// This is suitable for usage with the [http.Server] to implement the
// IPP over USB protocol or for the similar tasks.
//
// Its Accept methods waits until and returns some idle [Endpoint],
// wrapped into the [EndpointConn] type. The [EndpointConn.Close]
// returns connection into the pool.
type EndpointListener struct {
	endpoints     chan *Endpoint // Queue of available endpoints
	closechan     chan struct{}  // Closed by EndpointListener.Close
	local, remote net.Addr       // Listener's addresses
}

// NewEndpointListener creates the new [EndpointListener] on a top
// of the group of existing [Endpoint]s.
//
// The provided local and remote addresses become local and remote
// addresses of the [EndpointConn] connections, "accepted" from the
// listener.
//
// The local address is also returned by the [net.Listener.Addr]
// method of the listener.
func NewEndpointListener(local, remote net.Addr,
	endpoints []*Endpoint) *EndpointListener {

	listener := &EndpointListener{
		endpoints: make(chan *Endpoint, len(endpoints)),
		closechan: make(chan struct{}),
		local:     local,
		remote:    remote,
	}

	for _, ep := range endpoints {
		listener.endpoints <- ep
	}

	return listener
}

// Accept waits for and returns the next connection to the listener.
func (listener *EndpointListener) Accept() (net.Conn, error) {
	select {
	case ep := <-listener.endpoints:
		conn := NewEndpointConn(ep, listener.local, listener.remote)
		conn.listener.Store(listener)
		return conn, nil
	case <-listener.closechan:
		return nil, net.ErrClosed
	}
}

// Close closes the listener.
// Any blocked Accept operations will be unblocked and return errors.
func (listener *EndpointListener) Close() error {
	close(listener.closechan)
	return nil
}

// Addr returns the listener's network address.
func (listener *EndpointListener) Addr() net.Addr {
	return listener.local
}
