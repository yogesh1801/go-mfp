// MFP       - Miulti-Function Printers and scanners toolkit
// TRANSPORT - Transport protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Loopback transport test

package transport

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"sync"
	"testing"
)

// TestLoopbackBasic performs a basic testing of loopback
func TestLoopbackBasic(t *testing.T) {
	message := []byte("Hello, world!\n")

	// Create loopbacked Client and Server
	tr, l := NewLoopback()

	handler := func(w http.ResponseWriter, rq *http.Request) {
		w.WriteHeader(200)
		w.Write(message)
	}

	srv := &http.Server{Handler: http.HandlerFunc(handler)}
	clnt := &http.Client{Transport: tr}

	// Run srv.Serve() on its own goroutine
	var done sync.WaitGroup
	done.Add(1)

	go func() {
		srv.Serve(l)
		done.Done()
	}()

	defer func() {
		srv.Close()
		done.Wait()
	}()

	// Perform HTTP transaction and test results
	resp, err := clnt.Get("http://loopback/")
	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		t.Errorf("HTTP error: %s", err)
		return
	}

	if resp.StatusCode != 200 {
		t.Errorf("HTTP response status mismatch:\n"+
			"expected: %d\n"+
			"present:  %d",
			200, resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("HTTP body read error: %s", err)
		return
	}

	if !bytes.Equal(body, message) {
		t.Errorf("HTTP response body mismatch:\n"+
			"expected: %q\n"+
			"present:  %q",
			message, body)
	}
}

// TestLoopbackDoubleClose closes twice net.Listener, returned by
// NewLoopback, and checks for results.
func TestLoopbackDoubleClose(t *testing.T) {
	_, l := NewLoopback()

	err := l.Close()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	err = l.Close()
	if err != ErrLoopbackClosed {
		t.Errorf("Error expected %s, present %s",
			ErrLoopbackClosed, err)
	}
}

// TestLoopbackCloseUnblocksAll tests that closing Loobpack
// net.Listener unblocks all waiters.
func TestLoopbackCloseUnblocksAll(t *testing.T) {
	tr, l := NewLoopback()

	var done sync.WaitGroup

	addr := l.Addr()

	for i := 0; i < 10; i++ {
		done.Add(1)
		go func() {
			tr.DialContext(context.Background(),
				addr.Network(), addr.String())
			done.Done()
		}()
	}

	for i := 0; i < 10; i++ {
		done.Add(1)
		go func() {
			l.Accept()
			done.Done()
		}()
	}

	l.Close()
	done.Wait()
}

// TestLoopbackBysyError tests that Loopback properly generates
// ErrLoopbackBusy error when there are too many connections attempts.
func TestLoopbackBysyError(t *testing.T) {
	tr, l := NewLoopback()

	addr := l.Addr()
	connCount := 0

	var err error

	for i := 0; i < LoopbackMaxPendingConnections*2; i++ {
		_, err = tr.DialContext(context.Background(),
			addr.Network(), addr.String())

		if err != nil {
			break
		}

		connCount++
	}

	if err == nil {
		err = errors.New("")
	}

	if err != ErrLoopbackBusy {
		t.Errorf("Error expected: %q, present: %q",
			ErrLoopbackBusy, err)
		return
	}

	if connCount != LoopbackMaxPendingConnections {
		t.Errorf("Connection count expected: %d, present: %d",
			LoopbackMaxPendingConnections, connCount)
	}

	l.Close()
}
