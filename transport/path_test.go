// MFP       - Miulti-Function Printers and scanners toolkit
// TRANSPORT - Transport protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Functions for paths tests

package transport

import "testing"

// TestCleanURLPath tests CleanURLPath
func TestCleanURLPath(t *testing.T) {
	type testData struct {
		in  string // Input path
		out string // Expected output
	}

	tests := []testData{
		{in: "", out: "/"},
		{in: "/", out: "/"},
		{in: "/./", out: "/"},
		{in: "//./", out: "/"},
		{in: "foo", out: "/foo"},
		{in: "foo/", out: "/foo/"},
		{in: "foo//", out: "/foo/"},
		{in: "////foo//", out: "/foo/"},
	}

	for _, test := range tests {
		out := CleanURLPath(test.in)
		if out != test.out {
			t.Errorf("CleanURLPath(%q):\n"+
				"expected: %q\n"+
				"present:  %q\n",
				test.in, test.out, out)
		}
	}
}
