// MFP       - Miulti-Function Printers and scanners toolkit
// TRANSPORT - Transport protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Transport test

package transport

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"testing"
)

// TestEscapePath tests that unescapePath(escapePath(s)) == s
func TestEscapePath(t *testing.T) {
	tests := []string{
		"",
		"/usr/bin",
		"usr:bin",
	}

	for _, test := range tests {
		escaped := escapePath(test)
		unescaped := unescapePath(escaped)

		if unescaped != test {
			t.Errorf("%q: escaped: %q, unescaped: %q",
				test, escaped, unescaped)
		}
	}
}

// TestUnescapePath contains additional tests for code paths
// not covered by the TestEscapePath() (mostly handling of
// invalid sequences).
func TestUnescapePath(t *testing.T) {
	type testData struct {
		escaped   string // Escaped string
		unescaped string // Expected unescaped string
	}

	tests := []testData{
		{
			escaped:   "",
			unescaped: "",
		},

		{
			escaped:   "0-1",
			unescaped: "0",
		},

		{
			escaped:   "0-12",
			unescaped: "0\x12",
		},

		{
			escaped:   "0-1a",
			unescaped: "0\x1a",
		},

		{
			escaped:   "0-1A",
			unescaped: "0\x1A",
		},

		{
			escaped:   "0-1x",
			unescaped: "0x",
		},
	}

	for _, test := range tests {
		unescaped := unescapePath(test.escaped)
		if unescaped != test.unescaped {
			t.Errorf("%q: expected: %q, present: %q",
				test.escaped, test.unescaped, unescaped)
		}
	}
}

// TestTransportDial tests how Transport handles connection
// to server, depending on the destination URL.
func TestTransportDial(t *testing.T) {
	type testData struct {
		dest          string // Destination URL
		network, addr string // dial network and address
	}

	tests := []testData{
		{
			dest:    "unix:/var/run/cups/cups.sock",
			network: "unix",
			addr:    "/var/run/cups/cups.sock",
		},

		{
			dest:    "http://[::1]/",
			network: "tcp",
			addr:    "[::1]:80",
		},

		{
			dest:    "http://[::1%25eth0]/",
			network: "tcp",
			addr:    "[::1%eth0]:80",
		},

		{
			dest:    "ipp://localhost:631/",
			network: "tcp",
			addr:    "localhost:631",
		},

		{
			dest:    "ipp://localhost/",
			network: "tcp",
			addr:    "localhost:631",
		},

		{
			dest:    "ipps://localhost/",
			network: "tcp",
			addr:    "localhost:631",
		},

		{
			dest:    "http://localhost/",
			network: "tcp",
			addr:    "localhost:80",
		},

		{
			dest:    "https://localhost/",
			network: "tcp",
			addr:    "localhost:443",
		},

		{
			dest:    "http://127.0.0.1:39205/",
			network: "tcp",
			addr:    "127.0.0.1:39205",
		},

		{
			dest:    "http://[::1]:39205/",
			network: "tcp",
			addr:    "[::1]:39205",
		},

		{
			dest:    "http://localhost:39205/",
			network: "tcp",
			addr:    "localhost:39205",
		},
	}

	var network, addr string

	dial := func(ctx context.Context, n, a string) (net.Conn, error) {
		network, addr = n, a
		return nil, errors.New("not implemented")
	}

	template := (http.DefaultTransport.(*http.Transport)).Clone()
	template.DialContext = dial
	tr := NewTransport(template)

	for _, test := range tests {
		u := MustParseURL(test.dest)
		rq, err := NewRequest("GET", u, nil)
		if err != nil {
			panic(err)
		}

		tr.RoundTrip(rq)

		if network != test.network {
			t.Errorf("%q: network expected %q, present %q",
				test.dest, test.network, network)
		}

		if addr != test.addr {
			t.Errorf("%q: addr expected %q, present %q",
				test.dest, test.addr, addr)
		}
	}
}

func TestTransport(t *testing.T) {

	//return

	rq, err := NewRequest("GET",
		MustParseURL("unix:/var/run/cups/cups.sock"), nil)
	//rq, err := NewRequest("GET", "http://localhost/", nil)

	if err != nil {
		panic(err)
	}

	tr := NewTransport(nil)
	rsp, err := tr.RoundTrip(rq)

	if err != nil {
		t.Errorf("%s", err)
		return
	}

	fmt.Printf("================\n")
	rsp.Header.WriteSubset(os.Stdout, nil)
}
