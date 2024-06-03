// MFP  - Miulti-Function Printers and scanners toolkit
// DEST - Destination URLs hanling
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Tests for IPP-specific URL parsing

package dest

import (
	"errors"
	"testing"
)

// TestUrlParse tests urlParse function
func TestUrlParse(t *testing.T) {
	type testData struct {
		in, ipp, http string // Input and expected IPP/HTTP output
		err           string // Expected error
	}

	tests := []testData{
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
			err: "Printer URL: scheme must be ipp or ipps",
		},
		{
			in:  "https://127.0.0.1/ipp/print",
			err: "Printer URL: scheme must be ipp or ipps",
		},
		{
			in:  "",
			err: "Printer URL: invalid URL",
		},
		{
			in:  "http://Invalid URL",
			err: "Printer URL: invalid URL",
		},
	}

	for _, test := range tests {
		http, norm, err := urlParse(test.in)
		if err == nil {
			err = errors.New("")
		}

		if test.err != "" {
			if err.Error() != test.err {
				t.Errorf("%s: error expected: %q, present: %q",
					test.in, test.err, err)
			}
		} else {
			if norm != test.ipp {
				t.Errorf("%s: IPP URL expected %q, present %q", test.in,
					test.ipp, norm)
			}
			if http.String() != test.http {
				t.Errorf("%s: HTTP URL expected %q, got %q", test.in,
					test.http, http.String())
			}
		}
	}
}
