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

// TestADFOptionsAddDel tests ADFOptions.Add and ADFOptions.Del operations
func TestADFOptionsAddDel(t *testing.T) {
	type testOp struct {
		action string
		val    ADFOption
	}

	type testData struct {
		seq []testOp
		res ADFOptions
	}

	tests := []testData{
		{
			seq: nil,
			res: 0,
		},

		{
			seq: []testOp{
				{"add", DetectPaperLoaded},
			},
			res: 1 << DetectPaperLoaded,
		},

		{
			seq: []testOp{
				{"add", DetectPaperLoaded},
				{"add", Duplex},
			},
			res: 1<<DetectPaperLoaded | 1<<Duplex,
		},

		{
			seq: []testOp{
				{"add", DetectPaperLoaded},
				{"add", Duplex},
				{"del", DetectPaperLoaded},
				{"add", SelectSinglePage},
			},
			res: 1<<Duplex | 1<<SelectSinglePage,
		},
	}

	for _, test := range tests {
		var opts ADFOptions
		seq := ""

		for _, op := range test.seq {
			seq += fmt.Sprintf("  %s %s\n", op.action, op.val)

			switch op.action {
			case "add":
				opts.Add(op.val)
			case "del":
				opts.Del(op.val)

			default:
				panic(fmt.Errorf("unknown action %q", op.action))
			}
		}

		if opts != test.res {
			t.Errorf("\nfor the sequence:\n%s"+
				"expected: %s\npresent:  %s",
				seq, test.res, opts)
		}
	}
}

// TestMakeADFOptions tests MakeADFOptions
func TestMakeADFOptions(t *testing.T) {
	type testData struct {
		in  []ADFOption
		res ADFOptions
	}

	tests := []testData{
		{[]ADFOption{}, 0},
		{[]ADFOption{DetectPaperLoaded},
			1 << DetectPaperLoaded,
		},
		{[]ADFOption{DetectPaperLoaded, SelectSinglePage, Duplex},
			1<<DetectPaperLoaded | 1<<SelectSinglePage | 1<<Duplex,
		},
	}

	for _, test := range tests {
		opts := MakeADFOptions(test.in...)
		if opts != test.res {
			t.Errorf("\n"+
				"input:    %s\n"+
				"expected: %s\n"+
				"present:  %s",
				test.in, test.res, opts)
		}
	}
}

// TestADFOptionsString tests ADFOptions.String
func TestADFOptionsString(t *testing.T) {
	type testData struct {
		opts ADFOptions
		s    string
	}

	tests := []testData{
		{0, ""},
		{MakeADFOptions(DetectPaperLoaded), "DetectPaperLoaded"},
		{MakeADFOptions(DetectPaperLoaded, Duplex),
			"DetectPaperLoaded,Duplex"},
		{MakeADFOptions(DetectPaperLoaded, Duplex),
			"Duplex,DetectPaperLoaded"},
		{MakeADFOptions(DetectPaperLoaded, SelectSinglePage, Duplex),
			"DetectPaperLoaded,SelectSinglePage,Duplex"},
	}

	for _, test := range tests {
		s := test.opts.String()

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
				test.opts, exp, out)
		}
	}
}
