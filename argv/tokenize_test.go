// MFP  - Miulti-Function Printers and scanners toolkit
// argv - Argv parsing mini-library
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Argv tokenizer test

package argv

import (
	"errors"
	"reflect"
	"testing"
)

// TestTokenize performs testing of Tokenize() function
func TestTokenize(t *testing.T) {
	type testData struct {
		in  string   // Input string
		out []string // Expected output
		err string   // Expected error, "" if none
	}

	tests := []testData{
		// Normal cases
		{
			in:  `param1 param2 param3`,
			out: []string{"param1", "param2", "param3"},
		},

		{
			in:  `param1 "param 2" "param3"`,
			out: []string{"param1", "param 2", "param3"},
		},

		{
			in:  `param1 hel"lo wo"rld "param3"`,
			out: []string{"param1", "hello world", "param3"},
		},

		{
			in:  `"\a\b\f\n\r\t\v"`,
			out: []string{"\x07\x08\x0c\x0a\x0d\x09\x0b"},
		},

		{
			in:  `"-\0-"`,
			out: []string{string([]byte{'-', 0, '-'})},
		},

		{
			in:  `"-\1-"`,
			out: []string{string([]byte{'-', 01, '-'})},
		},

		{
			in:  `"-\12-"`,
			out: []string{string([]byte{'-', 012, '-'})},
		},

		{
			in:  `"-\123-"`,
			out: []string{string([]byte{'-', 0123, '-'})},
		},

		{
			in:  `"-\12"3-`,
			out: []string{string([]byte{'-', 012, '3', '-'})},
		},

		{
			in:  `"-\x0-"`,
			out: []string{string([]byte{'-', 0x00, '-'})},
		},

		{
			in:  `"-\x12-"`,
			out: []string{string([]byte{'-', 0x12, '-'})},
		},

		{
			in:  `"-\xaB-"`,
			out: []string{string([]byte{'-', 0xab, '-'})},
		},

		{
			in:  `"-\x1"-`,
			out: []string{string([]byte{'-', 0x01, '-'})},
		},

		{
			in:  `"-\x123-"`,
			out: []string{string([]byte{'-', 0x12, '3', '-'})},
		},

		{
			in:  `"-\"-"`,
			out: []string{string([]byte{'-', '"', '-'})},
		},

		// Errors handling
		{
			in:  `"param1" "param2`,
			out: []string{"param1", "param2"},
			err: `unterminated string`,
		},

		{
			in:  `"param1" "param2\`,
			out: []string{"param1", "param2"},
			err: `unterminated string`,
		},
	}

	for i, test := range tests {
		out, err := Tokenize(test.in)
		if err == nil {
			err = errors.New("")
		}

		if err.Error() != test.err {
			t.Errorf("[%d]: error mismatch: expected `%s`, present `%s`",
				i, test.err, err)
		} else {
			if !reflect.DeepEqual(out, test.out) {
				t.Errorf("[%d]: output mismatch:\nexpected %q\npresent  %q (%#x)",
					i, test.out, out, out)
			}
		}
	}
}
