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
		in   string   // Input string
		argv []string // Expected output
		err  string   // Expected error, "" if none
	}

	tests := []testData{
		// Normal cases
		{
			in:   `param1 param2 param3`,
			argv: []string{"param1", "param2", "param3"},
		},

		{
			in:   `param1 "param 2" "param3"`,
			argv: []string{"param1", "param 2", "param3"},
		},

		{
			in:   `param1 hel"lo wo"rld "param3"`,
			argv: []string{"param1", "hello world", "param3"},
		},

		{
			in:   `"\a\b\f\n\r\t\v"`,
			argv: []string{"\x07\x08\x0c\x0a\x0d\x09\x0b"},
		},

		{
			in:   `"-\0-"`,
			argv: []string{string([]byte{'-', 0, '-'})},
		},

		{
			in:   `"-\1-"`,
			argv: []string{string([]byte{'-', 01, '-'})},
		},

		{
			in:   `"-\12-"`,
			argv: []string{string([]byte{'-', 012, '-'})},
		},

		{
			in:   `"-\123-"`,
			argv: []string{string([]byte{'-', 0123, '-'})},
		},

		{
			in:   `"-\12"3-`,
			argv: []string{string([]byte{'-', 012, '3', '-'})},
		},

		{
			in:   `"-\x0-"`,
			argv: []string{string([]byte{'-', 0x00, '-'})},
		},

		{
			in:   `"-\x12-"`,
			argv: []string{string([]byte{'-', 0x12, '-'})},
		},

		{
			in:   `"-\xaB-"`,
			argv: []string{string([]byte{'-', 0xab, '-'})},
		},

		{
			in:   `"-\x1"-`,
			argv: []string{string([]byte{'-', 0x01, '-'})},
		},

		{
			in:   `"-\x123-"`,
			argv: []string{string([]byte{'-', 0x12, '3', '-'})},
		},

		{
			in:   `"-\"-"`,
			argv: []string{string([]byte{'-', '"', '-'})},
		},

		// Errors handling
		{
			in:   `"param1" "param2`,
			argv: []string{"param1", "param2"},
			err:  `unterminated string`,
		},

		{
			in:   `"param1" "param2\`,
			argv: []string{"param1", "param2"},
			err:  `unterminated string`,
		},
	}

	for i, test := range tests {
		argv, err := Tokenize(test.in)
		if err == nil {
			err = errors.New("")
		}

		if err.Error() != test.err {
			t.Errorf("[%d]: error mismatch: expected `%s`, present `%s`",
				i, test.err, err)
		} else {
			if !reflect.DeepEqual(argv, test.argv) {
				t.Errorf("[%d]: output mismatch:\nexpected %q\npresent  %q",
					i, test.argv, argv)
			}
		}
	}
}

// TestTokenizeEx performs testing of TokenizExe() function
func TestTokenizeEx(t *testing.T) {
	type testData struct {
		in   string   // Input string
		argv []string // Expected output
		tail string   // Expected tail
		err  string   // Expected error, "" if none
	}

	tests := []testData{
		{
			in:   `"param1" "param2`,
			argv: []string{"param1", "param2"},
			tail: ``,
			err:  `unterminated string`,
		},

		{
			in:   `"param1" "param2\`,
			argv: []string{"param1", "param2"},
			tail: `\`,
			err:  `unterminated string`,
		},

		{
			in:   `"param1" "param2\0`,
			argv: []string{"param1", "param2"},
			tail: `\0`,
			err:  `unterminated string`,
		},

		{
			in:   `"param1" "param2\01`,
			argv: []string{"param1", "param2"},
			tail: `\01`,
			err:  `unterminated string`,
		},

		{
			in:   `"param1" "param2\012`,
			argv: []string{"param1", "param2\012"},
			tail: ``,
			err:  `unterminated string`,
		},

		{
			in:   `"param1" "param2\x`,
			argv: []string{"param1", "param2"},
			tail: `\x`,
			err:  `unterminated string`,
		},

		{
			in:   `"param1" "param2\x3`,
			argv: []string{"param1", "param2"},
			tail: `\x3`,
			err:  `unterminated string`,
		},

		{
			in:   `"param1" "param2\x34`,
			argv: []string{"param1", "param2\x34"},
			tail: ``,
			err:  `unterminated string`,
		},
	}

	for i, test := range tests {
		argv, tail, err := TokenizeEx(test.in)
		if err == nil {
			err = errors.New("")
		}

		if tail != test.tail {
			t.Errorf("[%d]: tail mismatch: expected `%s`, present `%s`",
				i, test.tail, tail)
		}

		if !reflect.DeepEqual(argv, test.argv) {
			t.Errorf("[%d]: output mismatch:\nexpected %q\npresent  %q",
				i, test.argv, argv)
		}
	}
}
