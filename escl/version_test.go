// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Protocol version test

package escl

import (
	"fmt"
	"testing"
)

// TestMakeVersion tests MakeVersion, Version.Major and Version.Minor.
func TestMakeVersion(t *testing.T) {
	type testData struct {
		major, minor int
		ver          string
	}

	tests := []testData{
		{2, 0, "2.0"},
		{1, 5, "1.5"},
		{2, 1, "2.1"},
		{2, 12, "2.12"},
		{2, 123, "2.123"},
		{2, 1234, "2.1234"},
	}

	for _, test := range tests {
		name := fmt.Sprintf("%d.%d", test.major, test.minor)

		ver := MakeVersion(test.major, test.minor)
		major := ver.Major()
		minor := ver.Minor()

		if ver.String() != test.ver {
			t.Errorf("%q: version expected %s, present %s",
				name, test.ver, ver)
		}

		if major != test.major {
			t.Errorf("%q: major expected %d, present %d",
				name, test.major, major)
		}

		if minor != test.minor {
			t.Errorf("%q: major expected %d, present %d",
				name, test.minor, minor)
		}
	}
}

// TestVersionString tests Version.String.
func TestVersionString(t *testing.T) {
	type testData struct {
		ver Version
		s   string
	}

	tests := []testData{
		{MakeVersion(1, 2), "1.2"},
		{MakeVersion(2, 0), "2.0"},
		{MakeVersion(2, 1), "2.1"},
		{MakeVersion(2, 12), "2.12"},
		{MakeVersion(2, 123), "2.123"},
		{MakeVersion(2, 1234), "2.1234"},
		{MakeVersion(2, 12345), "2.1234"},
	}

	for _, test := range tests {
		name := fmt.Sprintf("%d.%d",
			test.ver.Major(), test.ver.Minor())

		s := test.ver.String()
		if s != test.s {
			t.Errorf("%q: expected %q, present %q",
				name, test.s, s)
		}
	}
}

// TestDecodeVersion tests DecodeVersion.
func TestDecodeVersion(t *testing.T) {
	type testData struct {
		s   string
		ver Version
		err string
	}

	tests := []testData{
		{"2.0", MakeVersion(2, 0), ""},
		{"xxx", 0, `"xxx": invalid eSCL version`},
		{"2.", 0, `"2.": invalid eSCL version`},
		{".0", 0, `".0": invalid eSCL version`},
		{"2.0a", 0, `"2.0a": invalid eSCL version`},
		{"2.12345", 0, `"2.12345": invalid eSCL version`},
	}

	for _, test := range tests {
		ver, err := DecodeVersion(test.s)
		errstr := ""
		if err != nil {
			errstr = err.Error()
		}

		if errstr != test.err {
			t.Errorf("%q: error expected %q, present %q",
				test.s, test.err, errstr)
			continue
		}

		if ver != test.ver {
			t.Errorf("%q: version expected %q, present %q",
				test.s, test.ver, ver)
		}
	}
}
