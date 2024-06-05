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

		// UNIX schama
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
			t.Errorf("%#v", u)
		}
	}
}

// TestMustParseURL tests how MustParseURL panics in a case of errors
func TestMustParseURL(t *testing.T) {
	defer func() {
		err := recover()
		if err != ErrURLSchemeMissed {
			t.Errorf("Error expected: %q, present: %v",
				ErrURLSchemeMissed, err)
		}
	}()

	MustParseURL("foo")
}
