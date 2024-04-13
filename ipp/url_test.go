// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Tests for IPP-specific URL parsing

package ipp

import (
	"errors"
	"testing"
)

type urlParseTest struct {
	in, ipp, http string // Input and expected IPP/HTTP output
	err           error  // Expected error
}

var urlParseTestData = []urlParseTest{
	{
		in:   "ipp://127.0.0.1/ipp/print",
		ipp:  "ipp://127.0.0.1/ipp/print",
		http: "http://127.0.0.1:631/ipp/print",
	},
	{
		in:   "ipp://127.0.0.1/ipp/xxx/../print",
		ipp:  "ipp://127.0.0.1/ipp/print",
		http: "http://127.0.0.1:631/ipp/print",
	},
	{
		in:   "ipp://127.0.0.1/../ipp/print",
		ipp:  "ipp://127.0.0.1/ipp/print",
		http: "http://127.0.0.1:631/ipp/print",
	},
	{
		in:   "ipp://127.0.0.1:80/ipp/print",
		ipp:  "ipp://127.0.0.1:80/ipp/print",
		http: "http://127.0.0.1/ipp/print",
	},
	{
		in:   "ipps://127.0.0.1:443/ipp/print",
		ipp:  "ipps://127.0.0.1:443/ipp/print",
		http: "https://127.0.0.1/ipp/print",
	},
	{
		in:   "ipp://127.0.0.1:631/ipp/print",
		ipp:  "ipp://127.0.0.1/ipp/print",
		http: "http://127.0.0.1:631/ipp/print",
	},
	{
		in:   "ipp://[fe80::aec5:1bff:fe1c:6fa7%252]/ipp/print",
		ipp:  "ipp://[fe80::aec5:1bff:fe1c:6fa7%252]/ipp/print",
		http: "http://[fe80::aec5:1bff:fe1c:6fa7%252]:631/ipp/print",
	},
	{
		in:  "http://127.0.0.1/ipp/print",
		err: errors.New("Printer URL: scheme must be ipp or ipps"),
	},
	{
		in:  "https://127.0.0.1/ipp/print",
		err: errors.New("Printer URL: scheme must be ipp or ipps"),
	},
	{
		in:  "",
		err: errors.New("Printer URL: invalid URL"),
	},
	{
		in:  "http://Invalid URL",
		err: errors.New("Printer URL: invalid URL"),
	},
}

// TestUrlParse tests urlParse function
func TestUrlParse(t *testing.T) {
	for _, test := range urlParseTestData {
		http, norm, err := urlParse(test.in)

		switch {
		case test.err == nil && err != nil:
			t.Errorf("%s: unexpected error: %q", test.in, err)
		case test.err != nil && err == nil:
			t.Errorf("%s: no error, expected: %q", test.in,
				test.err)
		case test.err != nil && err != nil:
			if test.err.Error() != err.Error() {
				t.Errorf("%s: error expected %q, got %q",
					test.in, test.err, err)
			}
		case norm != test.ipp:
			t.Errorf("%s: IPP expected %q, got %q", test.in,
				test.ipp, norm)
		case http.String() != test.http:
			t.Errorf("%s: HTTP expected %q, got %q", test.in,
				test.http, http.String())
		}
	}
}
