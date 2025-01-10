// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Set of CCD binary rendering modes test

package escl

import (
	"fmt"
	"slices"
	"strings"
	"testing"
)

// TestBinaryRenderingsAddDel tests BinaryRenderings.Add and BinaryRenderings.Del operations
func TestBinaryRenderingsAddDel(t *testing.T) {
	type testOp struct {
		action string
		val    BinaryRendering
	}

	type testData struct {
		seq []testOp
		res BinaryRenderings
	}

	tests := []testData{
		{
			seq: nil,
			res: BinaryRenderings{},
		},

		{
			seq: []testOp{
				{"add", Halftone},
			},
			res: BinaryRenderings{
				1 << Halftone,
				UnknownBinaryRendering,
			},
		},

		{
			seq: []testOp{
				{"add", Halftone},
				{"add", Threshold},
			},
			res: BinaryRenderings{
				1<<Halftone | 1<<Threshold,
				UnknownBinaryRendering,
			},
		},

		{
			seq: []testOp{
				{"add", Halftone},
				{"add", Threshold},
				{"del", Halftone},
			},
			res: BinaryRenderings{
				1 << Threshold,
				UnknownBinaryRendering,
			},
		},
	}

	for _, test := range tests {
		var bits BinaryRenderings
		seq := ""

		for _, op := range test.seq {
			seq += fmt.Sprintf("  %s %s\n", op.action, op.val)

			switch op.action {
			case "add":
				bits.Add(op.val)
			case "del":
				bits.Del(op.val)

			default:
				panic(fmt.Errorf("unknown action %q", op.action))
			}
		}

		if bits != test.res {
			t.Errorf("\nfor the sequence:\n%s"+
				"expected: %s\npresent:  %s",
				seq, test.res, bits)
		}
	}
}

// TestMakeBinaryRenderings tests MakeBinaryRenderings
func TestMakeBinaryRenderings(t *testing.T) {
	type testData struct {
		in  []BinaryRendering
		res BinaryRenderings
	}

	tests := []testData{
		{[]BinaryRendering{}, BinaryRenderings{}},
		{
			[]BinaryRendering{Halftone},
			BinaryRenderings{
				1 << Halftone,
				UnknownBinaryRendering,
			},
		},
		{
			[]BinaryRendering{Halftone, Threshold},
			BinaryRenderings{
				1<<Halftone | 1<<Threshold,
				UnknownBinaryRendering,
			},
		},
	}

	for _, test := range tests {
		res := MakeBinaryRenderings(test.in...)
		if res != test.res {
			t.Errorf("\n"+
				"input:    %s\n"+
				"expected: %s\n"+
				"present:  %s",
				test.in, test.res, res)
		}
	}
}

// TestBinaryRenderingsString tests BinaryRenderings.String
func TestBinaryRenderingsString(t *testing.T) {
	type testData struct {
		bits BinaryRenderings
		s    string
	}

	tests := []testData{
		{BinaryRenderings{}, ""},
		{MakeBinaryRenderings(Halftone), "Halftone"},
		{MakeBinaryRenderings(Threshold), "Threshold"},
		{MakeBinaryRenderings(Halftone, Threshold),
			"Halftone,Threshold"},
		{MakeBinaryRenderings(Halftone, Threshold),
			"Threshold,Halftone"},
	}

	for _, test := range tests {
		s := test.bits.String()

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
				test.bits, exp, out)
		}
	}
}
