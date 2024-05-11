// MFP  - Miulti-Function Printers and scanners toolkit
// argv - Argv parsing mini-library
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// String functions test

package argv

import "testing"

// TestStrCommonPrefix tests strCommonPrefix
func TestStrCommonPrefix(t *testing.T) {
	type testData struct {
		s1, s2 string
		pfx    string
	}

	tests := []testData{
		{s1: "foo", s2: "bar", pfx: ""},
		{s1: "foo", s2: "foobar", pfx: "foo"},
		{s1: "foobar", s2: "foo", pfx: "foo"},
		{s1: "foo", s2: "", pfx: ""},
		{s1: "", s2: "bar", pfx: ""},
		{s1: "", s2: "", pfx: ""},
	}

	for _, test := range tests {
		pfx := strCommonPrefix(test.s1, test.s2)
		if pfx != test.pfx {
			t.Errorf("{%q, %q}:\nexpected: %q\nreceived: %q",
				test.s1, test.s2, test.pfx, pfx)
		}
	}
}

// TestStrCommonPrefixSlice tests strCommonPrefixSlice
func TestStrCommonPrefixSlice(t *testing.T) {
	type testData struct {
		in  []string
		pfx string
	}

	tests := []testData{
		{in: []string{"1", "2", "3"}, pfx: ""},
		{in: []string{"foo-1", "foo-2", "foo-3"}, pfx: "foo-"},
		{in: []string{"foo-3", "foo-2", "foo-1"}, pfx: "foo-"},
		{in: []string{"foo-1", "", "foo-3"}, pfx: ""},
		{in: []string{}, pfx: ""},
		{in: nil, pfx: ""},
	}

	for _, test := range tests {
		pfx := strCommonPrefixSlice(test.in)
		if pfx != test.pfx {
			t.Errorf("%q:\nexpected: %q\nreceived: %q",
				test.in, test.pfx, pfx)
		}
	}
}
