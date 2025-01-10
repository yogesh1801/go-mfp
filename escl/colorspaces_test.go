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

// TestColorSpacesAddDel tests ColorSpaces.Add and ColorSpaces.Del operations
func TestColorSpacesAddDel(t *testing.T) {
	type testOp struct {
		action string
		val    ColorSpace
	}

	type testData struct {
		seq []testOp
		res ColorSpaces
	}

	tests := []testData{
		{
			seq: nil,
			res: ColorSpaces{},
		},

		{
			seq: []testOp{
				{"add", SRGB},
			},
			res: ColorSpaces{1 << SRGB, UnknownColorSpace},
		},

		{
			seq: []testOp{
				{"add", SRGB},
				{"del", SRGB},
			},
			res: ColorSpaces{},
		},
	}

	for _, test := range tests {
		var spaces ColorSpaces
		seq := ""

		for _, op := range test.seq {
			seq += fmt.Sprintf("  %s %s\n", op.action, op.val)

			switch op.action {
			case "add":
				spaces.Add(op.val)
			case "del":
				spaces.Del(op.val)

			default:
				panic(fmt.Errorf("unknown action %q", op.action))
			}
		}

		if spaces != test.res {
			t.Errorf("\nfor the sequence:\n%s"+
				"expected: %s\npresent:  %s",
				seq, test.res, spaces)
		}
	}
}

// TestMakeColorSpaces tests MakeColorSpaces
func TestMakeColorSpaces(t *testing.T) {
	type testData struct {
		in  []ColorSpace
		res ColorSpaces
	}

	tests := []testData{
		{[]ColorSpace{}, ColorSpaces{}},
		{[]ColorSpace{SRGB},
			ColorSpaces{
				1 << SRGB,
				UnknownColorSpace,
			},
		},
	}

	for _, test := range tests {
		spaces := MakeColorSpaces(test.in...)
		if spaces != test.res {
			t.Errorf("\n"+
				"input:    %s\n"+
				"expected: %s\n"+
				"present:  %s",
				test.in, test.res, spaces)
		}
	}
}

// TestColorSpacesString tests ColorSpaces.String
func TestColorSpacesString(t *testing.T) {
	type testData struct {
		spaces ColorSpaces
		s      string
	}

	tests := []testData{
		{ColorSpaces{}, ""},
		{MakeColorSpaces(SRGB), "sRGB"},
	}

	for _, test := range tests {
		s := test.spaces.String()

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
				test.spaces, exp, out)
		}
	}
}
