// MFP       - Miulti-Function Printers and scanners toolkit
// TRANSPORT - Transport protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// HTTP requests test

package transport

import (
	"fmt"
	"testing"
)

// TestRequestHost tests how Request.Host field is set,
// depending on a URL
func TestRequestHost(t *testing.T) {
	type testData struct {
		url  string // URL string
		host string // Expected host
	}

	tests := []testData{
		// localhost targets
		{
			url:  "http://localhost/",
			host: "localhost",
		},

		// http targets
		{
			url:  "http://example.com/",
			host: "example.com",
		},

		{
			url:  "http://example.com:80/",
			host: "example.com",
		},

		{
			url:  "http://example.com:123/",
			host: "example.com:123",
		},

		// https targets
		{
			url:  "https://example.com/",
			host: "example.com",
		},

		{
			url:  "https://example.com:443/",
			host: "example.com",
		},

		{
			url:  "https://example.com:123/",
			host: "example.com:123",
		},

		// ipp targets
		{
			url:  "ipp://example.com/",
			host: "example.com:631",
		},

		{
			url:  "ipp://example.com:80/",
			host: "example.com",
		},

		{
			url:  "ipps://example.com/",
			host: "example.com:631",
		},

		{
			url:  "ipps://example.com:443/",
			host: "example.com",
		},

		// IPv6 targets
		{
			url:  "http://[::1]/",
			host: "[::1]",
		},

		{
			url:  "http://[::1%25lo0]/",
			host: "[::1]",
		},

		{
			url:  "http://[::1%25lo0]:123/",
			host: "[::1]:123",
		},

		// unix targets
		{
			url:  "unix:/var/run/cups/cups.sock",
			host: "localhost",
		},
	}

	for _, test := range tests {
		rq, err := NewRequest("GET", test.url, nil)
		if err != nil {
			panic(fmt.Errorf("%q: %w", test.url, err))
		}

		if rq.Host != test.host {
			t.Errorf("%q: Host expected %q, present %q",
				test.url, test.host, rq.Host)
		}
	}
}
