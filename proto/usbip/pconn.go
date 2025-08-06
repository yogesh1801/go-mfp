// MFP - Multi-Function Printers and scanners toolkit
// Virtual USB/IP device emulator for testing and fuzzing
//
// Copyright (C) 2025 and up by GO-MFP authors.
// See LICENSE for license terms and conditions
//
// USBIP protocol connection

package main

import (
	"fmt"
	"io"
	"net"
	"sync"

	"github.com/OpenPrinting/go-mfp/util/generic"
)

// protoConn represents the protocol connection.
type protoConn struct {
	conn    net.Conn          // Underlying network connection
	txqueue []protoIOResponse // Transmit queue
	txlock  sync.Mutex        // txqueue lock
	txready chan struct{}     // Signaled on push to the txqueue
	txdone  sync.WaitGroup    // To wait for txroutine completion
}

// newProtoConn creates a new protoConn.
func newProtoConn(conn net.Conn) *protoConn {
	pconn := &protoConn{
		conn:    conn,
		txready: make(chan struct{}, 1),
	}

	pconn.txdone.Add(1)
	go pconn.txroutine()

	return pconn
}

// Close closes the protocol connection.
func (pconn *protoConn) Close() {
	pconn.txPush(nil)
	pconn.conn.Close()
	pconn.txdone.Wait()
}

// RecvHandshake receives protoHandshakeRequest from the connection.
func (pconn *protoConn) RecvHandshake() (protoHandshakeRequest, error) {
	// Receive protocol version, operation code and status
	var buf [8]byte
	err := pconn.read(buf[:])
	if err != nil {
		return nil, err
	}

	dec := newDecoder(buf[:])
	ver := dec.GetBE16()
	op := dec.GetBE16()

	if ver != protoVersion {
		err := fmt.Errorf("Invalid protocol version 0x%4.4x", ver)
		return nil, err
	}

	var rq protoHandshakeRequest

	switch protoHsCode(op) {
	case protoHsDevlistReq:
		rq = &protoDevlistRequest{}
	case protoHsImportReq:
		rq = &protoImportRequest{}

	default:
		err := fmt.Errorf("Invalid operation code 0x%4.4x", op)
		return nil, err
	}

	return rq, rq.Recv(pconn)
}

// SendHandshake sends protoHandshakeResponse into the connection.
func (pconn *protoConn) SendHandshake(rsp protoHandshakeResponse) error {
	enc := newEncoder(1024)

	enc.PutBE16(protoVersion)
	enc.PutBE16(uint16(rsp.OpCode()))
	st := rsp.OpStatus()
	enc.PutBE32(uint32(st))

	err := pconn.write(enc.Bytes())
	if err == nil && st == protoHsOK {
		err = rsp.Send(pconn)
	}

	return err
}

// RecvIO receives protoIORequest from the connection.
func (pconn *protoConn) RecvIO() (protoIORequest, error) {
	// Receive protoIOHeader
	var hdr protoIOHeader
	err := hdr.Recv(pconn)
	if err != nil {
		return nil, err
	}

	// Receive request body
	switch hdr.Command {
	case protoIOSubmitCmd:
		rq := &protoIOSubmitRequest{protoIOHeader: hdr}
		return rq, rq.Recv(pconn)

	case protoIOUnlinkCmd:
		rq := &protoIOUnlinkRequest{protoIOHeader: hdr}
		return rq, rq.Recv(pconn)
	}

	err = fmt.Errorf("Invalid command code 0x%4.4x", hdr.Command)
	return nil, err
}

// SendIO send protoIOResponse into the connection.
func (pconn *protoConn) SendIO(rsp protoIOResponse) error {
	pconn.txPush(rsp)
	return nil
}

// txroutine runs in the separate goroutine and transmits
// queued responses into the connection
func (pconn *protoConn) txroutine() {
	for {
		// Pull the response
		rsp := pconn.txPull()
		if rsp == nil {
			break
		}

		// Send the response
		hdr := rsp.Header()
		err := hdr.Send(pconn)
		if err == nil {
			err = rsp.Send(pconn)
		}

		if err != nil {
			break
		}
	}

	pconn.conn.Close()
	pconn.txdone.Done()
}

// txPush pushes protoIOResponse into the txqueue
func (pconn *protoConn) txPush(rsp protoIOResponse) {
	pconn.txlock.Lock()
	pconn.txqueue = append(pconn.txqueue, rsp)
	pconn.txlock.Unlock()

	select {
	case pconn.txready <- struct{}{}:
	default:
	}
}

// txPull pulls protoIOResponse from the txqueue
func (pconn *protoConn) txPull() protoIOResponse {
	pconn.txlock.Lock()
	defer pconn.txlock.Unlock()

	for {
		if len(pconn.txqueue) > 0 {
			rsp := pconn.txqueue[0]
			n := copy(pconn.txqueue, pconn.txqueue[1:])
			pconn.txqueue = pconn.txqueue[:n]
			return rsp
		}

		pconn.txlock.Unlock()
		<-pconn.txready
		pconn.txlock.Lock()
	}
}

// read reads the full buffer from the network connection.
// It returns the io.ErrUnexpectedEOF error, if connection was closed
// by peer before all data is received.
func (pconn *protoConn) read(buf []byte) error {
	for len(buf) > 0 {
		n, err := pconn.conn.Read(buf)
		switch {
		case n > 0:
			// Some data was received. Ignore the error, if any
			buf = buf[n:]

		case err == io.EOF:
			// Connection has closed prematurely
			fallthrough

		case err != nil:
			// Read error
			return err
		}
	}

	return nil
}

// drain reads and discards the specified number of bytes from the connection.
func (pconn *protoConn) drain(n uint32) error {
	var buf [65536]byte

	for n > 0 {
		sz := generic.Min(n, uint32(len(buf)))
		err := pconn.read(buf[:sz])
		if err != nil {
			return err
		}
		n -= sz
	}

	return nil
}

// write writes the whole buffer into the network connection.
func (pconn *protoConn) write(buf []byte) error {
	for len(buf) > 0 {
		n, err := pconn.conn.Write(buf)
		switch {
		case n > 0:
			// Some data was sent. Ignore the error, if any
			buf = buf[n:]

		case err != nil:
			// Write error
			return err
		}
	}

	return nil
}
