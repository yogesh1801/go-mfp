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
