// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Set of scan color modes

package escl

import (
	"fmt"
	"slices"
	"strings"
	"testing"
)

// TestColorModesAddDel tests ColorModes.Add and ColorModes.Del operations
func TestColorModesAddDel(t *testing.T) {
	type testOp struct {
		action string
		val    ColorMode
	}

	type testData struct {
		seq []testOp
		res ColorModes
	}

	tests := []testData{
		{
			seq: nil,
			res: 0,
		},

		{
			seq: []testOp{
				{"add", BlackAndWhite1},
			},
			res: 1 << BlackAndWhite1,
		},

		{
			seq: []testOp{
				{"add", BlackAndWhite1},
				{"add", Grayscale8},
				{"add", Grayscale16},
			},
			res: 1<<BlackAndWhite1 |
				1<<Grayscale8 |
				1<<Grayscale16,
		},

		{
			seq: []testOp{
				{"add", BlackAndWhite1},
				{"add", Grayscale8},
				{"add", Grayscale16},
				{"del", BlackAndWhite1},
				{"add", RGB24},
			},
			res: 1<<Grayscale8 |
				1<<Grayscale16 |
				1<<RGB24,
		},
	}

	for _, test := range tests {
		var cmodes ColorModes
		seq := ""

		for _, op := range test.seq {
			seq += fmt.Sprintf("  %s %s\n", op.action, op.val)

			switch op.action {
			case "add":
				cmodes.Add(op.val)
			case "del":
				cmodes.Del(op.val)

			default:
				panic(fmt.Errorf("unknown action %q", op.action))
			}
		}

		if cmodes != test.res {
			t.Errorf("\nfor the sequence:\n%s"+
				"expected: %s\npresent:  %s",
				seq, test.res, cmodes)
		}
	}
}

// TestMakeColorModes tests MakeColorModes
func TestMakeColorModes(t *testing.T) {
	type testData struct {
		in  []ColorMode
		res ColorModes
	}

	tests := []testData{
		{[]ColorMode{}, 0},
		{[]ColorMode{BlackAndWhite1},
			1 << BlackAndWhite1,
		},
		{[]ColorMode{BlackAndWhite1, Grayscale8, RGB24},
			1<<BlackAndWhite1 | 1<<Grayscale8 | 1<<RGB24,
		},
	}

	for _, test := range tests {
		cmodes := MakeColorModes(test.in...)
		if cmodes != test.res {
			t.Errorf("\n"+
				"input:    %s\n"+
				"expected: %s\n"+
				"present:  %s",
				test.in, test.res, cmodes)
		}
	}
}

// TestColorModesString tests ColorModes.String
func TestColorModesString(t *testing.T) {
	type testData struct {
		cmodes ColorModes
		s      string
	}

	tests := []testData{
		{0, ""},
		{MakeColorModes(BlackAndWhite1), "BlackAndWhite1"},
		{MakeColorModes(BlackAndWhite1, RGB24),
			"BlackAndWhite1,RGB24"},
		{MakeColorModes(BlackAndWhite1, RGB24),
			"RGB24,BlackAndWhite1"},
		{MakeColorModes(BlackAndWhite1, RGB24, RGB48),
			"RGB24,BlackAndWhite1,RGB48"},
	}

	for _, test := range tests {
		s := test.cmodes.String()

		// Compare resulting strings, ignoring the order
		// of color modes in the output.
		out := strings.Split(s, ",")
		exp := strings.Split(test.s, ",")

		slices.Sort(out)
		slices.Sort(exp)

		if !slices.Equal(out, exp) {
			t.Errorf("%s:\n"+
				"expected: %s\n"+
				"present:  %s",
				test.cmodes, exp, out)
		}
	}
}
