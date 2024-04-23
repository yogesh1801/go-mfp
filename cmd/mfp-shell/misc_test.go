// MFP   - Miulti-Function Printers and scanners toolkit
// mains - Main functions for all commands
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Miscellaneous tests

package main

import (
	"errors"
	"reflect"
	"testing"
)

// TestTokenize tests tokenize() function in the main.go
func TestTokenize(t *testing.T) {
	// testData contains test date for a single test
	type testData struct {
		input  string
		output []string
		err    error
	}

	// tests contains collection of tokenize() tests
	tests := []testData{
		{
			input:  ``,
			output: []string{},
		},

		{
			input:  `   `,
			output: []string{},
		},

		{
			input:  `  ABC  `,
			output: []string{"ABC"},
		},

		{
			input:  `a b c`,
			output: []string{`a`, `b`, `c`},
		},

		{
			input:  `a "b" c`,
			output: []string{`a`, `b`, `c`},
		},

		{
			input:  `a "b c" d`,
			output: []string{`a`, `b c`, `d`},
		},

		{
			input:  `"\x12"`,
			output: []string{"\x12"},
		},

		{
			input:  `"\x1f"`,
			output: []string{"\x1f"},
		},

		{
			input:  `"\x1F"`,
			output: []string{"\x1f"},
		},

		{
			input:  `"\x1"`,
			output: []string{"\x01"},
		},

		{
			input:  `"\x123"`,
			output: []string{"\x12" + "3"},
		},

		{
			input:  `"\123"`,
			output: []string{"\123"},
		},

		{
			input:  `"\1234"`,
			output: []string{"\123" + "4"},
		},

		{
			input:  `"\12"`,
			output: []string{"\012"},
		},

		{
			input:  `"\1"`,
			output: []string{"\001"},
		},

		{
			input:  `"\12X"`,
			output: []string{"\012" + "X"},
		},

		{
			input:  `"\1X"`,
			output: []string{"\001" + "X"},
		},

		{
			//input:  `"\123\321"`,
			input:  `"\123\103"`,
			output: []string{"\123\103"},
		},

		{
			input:  `"\x12\x34"`,
			output: []string{"\x12\x34"},
		},

		{
			input:  `"\x2Z"`,
			output: []string{"\x02" + "Z"},
		},

		{
			input:  `"A\aB"`,
			output: []string{"A\aB"},
		},

		{
			input:  `"A\bB"`,
			output: []string{"A\bB"},
		},

		{
			input:  `"A\fB"`,
			output: []string{"A\fB"},
		},

		{
			input:  `"A\nB"`,
			output: []string{"A\nB"},
		},

		{
			input:  `"A\rB"`,
			output: []string{"A\rB"},
		},

		{
			input:  `"A\tB"`,
			output: []string{"A\tB"},
		},

		{
			input:  `"A\vB"`,
			output: []string{"A\vB"},
		},

		{
			input:  `"A\"B"`,
			output: []string{"A\"B"},
		},

		{
			input: `"`,
			err:   errors.New("unterminated string"),
		},
	}

	for _, test := range tests {
		output, err := tokenize(test.input)

		switch {
		case err == nil && test.err != nil:
			t.Errorf("%q->%q, expected %s",
				test.input, output, test.err)
		case err != nil && test.err == nil:
			t.Errorf("%q->%s, expected %q",
				test.input, err, test.output)
		case !reflect.DeepEqual(output, test.output):
			t.Errorf("%q->%q, expected %q",
				test.input, output, test.output)
		}
	}
}
