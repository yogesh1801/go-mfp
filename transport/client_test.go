// MFP       - Miulti-Function Printers and scanners toolkit
// TRANSPORT - Transport protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// HTTP client wrapper test.

package transport

import "testing"

// TestNewClient tests NewClient function
func TestNewClient(t *testing.T) {
	// NewClient(nil) must create a new Transport
	clnt := NewClient(nil)
	if clnt.Transport == nil {
		t.Errorf("NewClient(nil): clnt.Transport == nil")
	}

	// NewClient(tr) must use provided Transport
	tr := NewTransport(nil)
	clnt = NewClient(tr)
	if clnt.Transport != tr {
		t.Errorf("NewClient(tr): clnt.Transport != tr")
	}
}
