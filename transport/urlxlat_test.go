// MFP       - Miulti-Function Printers and scanners toolkit
// TRANSPORT - Transport protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// URL proxying test

package transport

import (
	"testing"
)

// TestURLXlat tests URLXlat
func TestURLXlat(t *testing.T) {
	type testData struct {
		local, remote string // Local/remote URLs
		in, out       string // Input/output URLs in forward direction
	}

	tests := []testData{
		{
			local:  "ipp://127.0.0.1:1234",
			remote: "ipp://127.0.0.1",
			in:     "ipp://127.0.0.1:1234/printers/1",
			out:    "ipp://127.0.0.1/printers/1",
		},

		{
			local:  "ipp://127.0.0.1:1234",
			remote: "ipp://127.0.0.1/xxx",
			in:     "ipp://127.0.0.1:1234/printers/1",
			out:    "ipp://127.0.0.1/xxx/printers/1",
		},

		{
			local:  "ipp://127.0.0.1:1234",
			remote: "unix:/var/run/cups/cups.sock",
			in:     "ipp://127.0.0.1:1234/printers/1",
			out:    "unix:/var/run/cups/cups.sock/printers/1",
		},

		{
			local:  "ipp://127.0.0.1/xxx",
			remote: "ipp://127.0.0.2/yyy",
			in:     "ipp://127.0.0.1/xxx/1",
			out:    "ipp://127.0.0.2/yyy/1",
		},

		{
			local:  "ipp://127.0.0.1/xxx",
			remote: "ipp://127.0.0.2/yyy",
			in:     "ipp://127.0.0.1/1",
			out:    "ipp://127.0.0.1/1",
		},

		{
			local:  "ipp://127.0.0.1/xxx",
			remote: "ipp://127.0.0.2/yyy",
			in:     "ipp://127.0.0.1/xxx-xxx/1",
			out:    "ipp://127.0.0.1/xxx-xxx/1",
		},

		{
			local:  "http://127.0.0.1/",
			remote: "ipp://192.168.0.1/",
			in:     "http://127.0.0.1/1",
			out:    "http://192.168.0.1:631/1",
		},

		{
			local:  "http://127.0.0.1:631/",
			remote: "ipp://192.168.0.1/",
			in:     "ipp://127.0.0.1/1",
			out:    "ipp://192.168.0.1/1",
		},

		{
			local:  "http://127.0.0.1:631/",
			remote: "ipp://192.168.0.1/",
			in:     "http://127.0.0.1:631/1",
			out:    "http://192.168.0.1:631/1",
		},

		{
			local:  "http://localhost:60002/ipp/print",
			remote: "ipp://localhost:631/printers/Boomaga",
			in:     "ipp://localhost:60002/ipp/print",
			out:    "ipp://localhost/printers/Boomaga",
		},

		{
			local:  "http://localhost:80/eSCLLL/",
			remote: "http://localhost:50000/eSCL",
			in:     "http://localhost/eSCLLL/ScanJobs/urn:uuid:652a6023-97af-4936-80d7-8962ed514842",
			out:    "http://localhost:50000/eSCL/ScanJobs/urn:uuid:652a6023-97af-4936-80d7-8962ed514842",
		},
	}

	for _, test := range tests {
		ux := NewURLXlat(MustParseURL(test.local),
			MustParseURL(test.remote))

		out := ux.Forward(MustParseURL(test.in)).String()
		if out != test.out {
			t.Errorf("forward %s->%s\n"+
				"input:    %q\n"+
				"expected: %q\n"+
				"present:  %q\n",
				test.local, test.remote,
				test.in, test.out, out)
		}

		in := ux.Reverse(MustParseURL(test.out)).String()
		if in != test.in {
			t.Errorf("reverse %s<-%s\n"+
				"input:    %q\n"+
				"expected: %q\n"+
				"present:  %q\n",
				test.local, test.remote,
				test.out, test.in, in)
		}
	}
}
