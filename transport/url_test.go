// MFP       - Miulti-Function Printers and scanners toolkit
// TRANSPORT - Transport protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Tests for IPP-specific URL parsing

package transport

import (
	"errors"
	"testing"
)

// TestParseUrl tests ParseURL function
func TestParseURL(t *testing.T) {
	type testData struct {
		in  string
		out string
		err string
	}

	tests := []testData{
		// HTTP schemes
		{
			in:  "http://127.0.0.1/ipp/print",
			out: "http://127.0.0.1/ipp/print",
		},

		{
			in:  "http://127.0.0.1:80/ipp/print",
			out: "http://127.0.0.1/ipp/print",
		},

		{
			in:  "http://127.0.0.1:81/ipp/print",
			out: "http://127.0.0.1:81/ipp/print",
		},

		{
			in:  "https://127.0.0.1:443/ipp/print",
			out: "https://127.0.0.1/ipp/print",
		},

		{
			in:  "https://127.0.0.1:444/ipp/print",
			out: "https://127.0.0.1:444/ipp/print",
		},

		{
			in:  "http://[fe80::aec5:1bff:fe1c:6fa7%252]/ipp/print",
			out: "http://[fe80::aec5:1bff:fe1c:6fa7%252]/ipp/print",
		},

		// IPP schemes
		{
			in:  "ipp://127.0.0.1/ipp/print",
			out: "ipp://127.0.0.1/ipp/print",
		},

		{
			in:  "ipp://127.0.0.1:631/ipp/print",
			out: "ipp://127.0.0.1/ipp/print",
		},

		{
			in:  "ipps://127.0.0.1:631/ipp/print",
			out: "ipps://127.0.0.1/ipp/print",
		},

		{
			in:  "ipps://127.0.0.1:632/ipp/print",
			out: "ipps://127.0.0.1:632/ipp/print",
		},

		// UNIX schema
		{
			in:  "unix:///var/run/cups/cups.sock",
			out: "unix:/var/run/cups/cups.sock",
		},

		{
			in:  "unix:/var/run/cups/cups.sock",
			out: "unix:/var/run/cups/cups.sock",
		},

		{
			in:  "unix://localhost/var/run/cups/cups.sock",
			out: "unix:/var/run/cups/cups.sock",
		},

		{
			in:  "unix://LoCaLhOsT/var/run/cups/cups.sock",
			out: "unix:/var/run/cups/cups.sock",
		},

		{
			in:  "unix://localhost:80/var/run/cups/cups.sock",
			err: ErrURLUNIXHost.Error(),
		},

		{
			in:  "unix://example.com/var/run/cups/cups.sock",
			err: ErrURLUNIXHost.Error(),
		},

		// Path handling
		{
			in:  "http://127.0.0.1/",
			out: "http://127.0.0.1/",
		},

		{
			in:  "http://127.0.0.1",
			out: "http://127.0.0.1/",
		},

		{
			in:  "http://127.0.0.1/foo/",
			out: "http://127.0.0.1/foo/",
		},

		{
			in:  "http://127.0.0.1/foo//////bar",
			out: "http://127.0.0.1/foo/bar",
		},

		{
			in:  "http://127.0.0.1/foo/./bar/../foobar",
			out: "http://127.0.0.1/foo/foobar",
		},

		// Scheme errors
		{
			in:  "foo",
			err: ErrURLSchemeMissed.Error(),
		},

		{
			in:  "foo:",
			err: ErrURLSchemeInvalid.Error(),
		},

		// Other errors:
		{
			in:  "http://Invalid URL",
			err: ErrURLInvalid.Error(),
		},

		{
			in:  "",
			err: ErrURLSchemeMissed.Error(),
		},

		{
			in:  "unix:///var/run/cups/cups.sock",
			out: "unix:/var/run/cups/cups.sock",
		},

		{
			in:  "unix:/var/run/cups/cups.sock",
			out: "unix:/var/run/cups/cups.sock",
		},
	}

	for _, test := range tests {
		u, err := ParseURL(test.in)
		if err == nil {
			err = errors.New("")
		}

		switch {
		case err.Error() != test.err:
			t.Errorf("%q: error mismatch:\nexpected: %s\npresent:  %s",
				test.in, test.err, err)

		case test.err != "":
			// Error as expected; nothing to do

		case u.String() != test.out:
			t.Errorf("%q: output mismatch:\nexpected: %s\npresent:  %s",
				test.in, test.out, u)
		}
	}
}

// TestMustParseURL tests how MustParseURL panics in a case of errors
func TestMustParseURL(t *testing.T) {
	defer func() {
		err, ok := recover().(error)
		if !ok || !errors.Is(err, ErrURLSchemeMissed) {
			t.Errorf("Error expected: %q, present: %v",
				ErrURLSchemeMissed, err)
		}
	}()

	MustParseURL("foo")
}

// TestParseAddr tests ParseAddr function
func TestParseAddr(t *testing.T) {
	type testData struct {
		in       string // Input address
		template string // Template URL
		out      string // Expected output
		err      string // Expected error
	}

	tests := []testData{
		// IP4 and IP4 addresses
		{
			in:  "127.0.0.1",
			out: "http://127.0.0.1/",
		},

		{
			in:  "::1",
			out: "http://[::1]/",
		},

		{
			in:  "[::1]",
			out: "http://[::1]/",
		},

		// IP4 and IP4 addresses with port
		{
			in:  "127.0.0.1:80",
			out: "http://127.0.0.1/",
		},

		{
			in:  "127.0.0.1:81",
			out: "http://127.0.0.1:81/",
		},

		{
			in:  "127.0.0.1:443",
			out: "https://127.0.0.1/",
		},

		{
			in:  "127.0.0.1:631",
			out: "ipp://127.0.0.1/",
		},

		{
			in:  "[::1]:80",
			out: "http://[::1]/",
		},

		{
			in:  "[::1]:81",
			out: "http://[::1]:81/",
		},

		// UNIX paths
		{
			in:  "/var/run/cups/cups.sock",
			out: "unix:/var/run/cups/cups.sock",
		},

		// IP address with template
		{
			in:       "127.0.0.1",
			template: "https://localhost/",
			out:      "https://127.0.0.1/",
		},

		{
			in:       "127.0.0.1",
			template: "http://localhost:222/",
			out:      "http://127.0.0.1:222/",
		},

		// IP address and port with template
		{
			in:       "127.0.0.1:1234",
			template: "https://localhost/path",
			out:      "https://127.0.0.1:1234/path",
		},

		// Full URLs
		{
			in:  "http://127.0.0.1/ipp/print",
			out: "http://127.0.0.1/ipp/print",
		},

		{
			in:  "http://127.0.0.1:80/ipp/print",
			out: "http://127.0.0.1/ipp/print",
		},
	}

	for _, test := range tests {
		u, err := ParseAddr(test.in, test.template)
		if err == nil {
			err = errors.New("")
		}

		switch {
		case err.Error() != test.err:
			t.Errorf("%q: error mismatch:\nexpected: %s\npresent:  %s",
				test.in, test.err, err)

		case test.err != "":
			// Error as expected; nothing to do

		case u.String() != test.out:
			t.Errorf("%q: output mismatch:\nexpected: %s\npresent:  %s",
				test.in, test.out, u)
		}

	}
}

// TestURLPort tests URLPort
func TestURLPort(t *testing.T) {
	type testData struct {
		in   string
		port int
	}

	tests := []testData{
		// Port is not set in URL
		{
			in:   "http://127.0.0.1",
			port: 80,
		},
		{
			in:   "https://127.0.0.1",
			port: 443,
		},
		{
			in:   "ipp://127.0.0.1",
			port: 631,
		},
		{
			in:   "ipps://127.0.0.1",
			port: 631,
		},

		// Port explicitly set to default
		{
			in:   "http://127.0.0.1:80",
			port: 80,
		},
		{
			in:   "https://127.0.0.1:443",
			port: 443,
		},
		{
			in:   "ipp://127.0.0.1:631",
			port: 631,
		},
		{
			in:   "ipps://127.0.0.1:631",
			port: 631,
		},

		// Port explicitly set to non-default
		{
			in:   "http://127.0.0.1:1234",
			port: 1234,
		},
		{
			in:   "https://127.0.0.1:1234",
			port: 1234,
		},
		{
			in:   "ipp://127.0.0.1:1234",
			port: 1234,
		},
		{
			in:   "ipps://127.0.0.1:1234",
			port: 1234,
		},

		// Invalid port
		{
			in:   "http://127.0.0.1:66666",
			port: -1,
		},
		{
			in:   "https://127.0.0.1:66666",
			port: -1,
		},
		{
			in:   "ipp://127.0.0.1:66666",
			port: -1,
		},
		{
			in:   "ipps://127.0.0.1:66666",
			port: -1,
		},

		// URL with unix: schema
		{
			in:   "unix:/var/run/cups/cups.sock",
			port: -1,
		},
	}

	for _, test := range tests {
		u := MustParseURL(test.in)
		if port := URLPort(u); port != test.port {
			t.Errorf("%s: URLPort expected: %d, present: %d",
				test.in, test.port, port)
		}
	}
}

// TestURLForcePort the tests URLForcePort function
func TestURLForcePort(t *testing.T) {
	type testData struct {
		in  string // Input URL
		out string // URL after URLForcePort
	}

	tests := []testData{
		{
			in:  "http://127.0.0.1/ipp/print",
			out: "http://127.0.0.1:80/ipp/print",
		},

		{
			in:  "ipp://127.0.0.1/ipp/print",
			out: "ipp://127.0.0.1:631/ipp/print",
		},

		{
			in:  "http://[::1]/",
			out: "http://[::1]:80/",
		},

		{
			in:  "http://[::1]:1234/",
			out: "http://[::1]:1234/",
		},

		{
			in:  "http://localhost/",
			out: "http://localhost:80/",
		},

		{
			in:  "unix:/var/run/cups/cups.sock",
			out: "unix:/var/run/cups/cups.sock",
		},
	}

	for _, test := range tests {
		u := MustParseURL(test.in)
		URLForcePort(u)
		out := u.String()

		if out != test.out {
			t.Errorf("%s:\n"+
				"expected: %s\n"+
				"present:  %s",
				test.in, test.out, out)
		}
	}
}

// TestURLStripPort the tests URLStripPort function
func TestURLStripPort(t *testing.T) {
	type testData struct {
		in  string // Input URL
		out string // URL after URLForcePort
	}

	tests := []testData{
		{
			in:  "http://127.0.0.1/ipp/print",
			out: "http://127.0.0.1/ipp/print",
		},

		{
			in:  "http://127.0.0.1:80/ipp/print",
			out: "http://127.0.0.1/ipp/print",
		},

		{
			in:  "ipp://127.0.0.1/ipp/print",
			out: "ipp://127.0.0.1/ipp/print",
		},

		{
			in:  "ipp://127.0.0.1:631/ipp/print",
			out: "ipp://127.0.0.1/ipp/print",
		},

		{
			in:  "https://127.0.0.1/ipp/print",
			out: "https://127.0.0.1/ipp/print",
		},

		{
			in:  "https://127.0.0.1:443/ipp/print",
			out: "https://127.0.0.1/ipp/print",
		},

		{
			in:  "ipps://127.0.0.1/ipp/print",
			out: "ipps://127.0.0.1/ipp/print",
		},

		{
			in:  "ipps://127.0.0.1:631/ipp/print",
			out: "ipps://127.0.0.1/ipp/print",
		},

		{
			in:  "http://[::1]/",
			out: "http://[::1]/",
		},

		{
			in:  "http://[::1]:80/",
			out: "http://[::1]/",
		},

		{
			in:  "http://[::1]:1234/",
			out: "http://[::1]:1234/",
		},

		{
			in:  "http://localhost/",
			out: "http://localhost/",
		},

		{
			in:  "http://localhost:80/",
			out: "http://localhost/",
		},

		{
			in:  "unix:/var/run/cups/cups.sock",
			out: "unix:/var/run/cups/cups.sock",
		},
	}

	for _, test := range tests {
		u := MustParseURL(test.in)
		URLStripPort(u)
		out := u.String()

		if out != test.out {
			t.Errorf("%s:\n"+
				"expected: %s\n"+
				"present:  %s",
				test.in, test.out, out)
		}
	}
}
