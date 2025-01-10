// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Set of CCD intents test

package escl

import (
	"fmt"
	"slices"
	"strings"
	"testing"
)

// TestIntentsAddDel tests Intents.Add and Intents.Del operations
func TestIntentsAddDel(t *testing.T) {
	type testOp struct {
		action string
		val    Intent
	}

	type testData struct {
		seq []testOp
		res Intents
	}

	tests := []testData{
		{
			seq: nil,
			res: Intents{},
		},

		{
			seq: []testOp{
				{"add", Document},
			},
			res: Intents{
				1 << Document,
				UnknownIntent,
			},
		},

		{
			seq: []testOp{
				{"add", Document},
				{"add", TextAndGraphic},
			},
			res: Intents{
				1<<Document | 1<<TextAndGraphic,
				UnknownIntent,
			},
		},

		{
			seq: []testOp{
				{"add", Document},
				{"add", TextAndGraphic},
				{"del", Document},
				{"add", Photo},
			},
			res: Intents{
				1<<TextAndGraphic | 1<<Photo,
				UnknownIntent,
			},
		},
	}

	for _, test := range tests {
		var opts Intents
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

// TestMakeIntents tests MakeIntents
func TestMakeIntents(t *testing.T) {
	type testData struct {
		in  []Intent
		res Intents
	}

	tests := []testData{
		{[]Intent{}, Intents{}},
		{[]Intent{Document},
			Intents{
				1 << Document,
				UnknownIntent,
			},
		},
		{[]Intent{Document, Photo, TextAndGraphic},
			Intents{
				1<<Document |
					1<<Photo |
					1<<TextAndGraphic,
				UnknownIntent,
			},
		},
	}

	for _, test := range tests {
		opts := MakeIntents(test.in...)
		if opts != test.res {
			t.Errorf("\n"+
				"input:    %s\n"+
				"expected: %s\n"+
				"present:  %s",
				test.in, test.res, opts)
		}
	}
}

// TestIntentsString tests Intents.String
func TestIntentsString(t *testing.T) {
	type testData struct {
		opts Intents
		s    string
	}

	tests := []testData{
		{Intents{}, ""},
		{MakeIntents(Document), "Document"},
		{MakeIntents(Document, TextAndGraphic),
			"Document,TextAndGraphic"},
		{MakeIntents(Document, TextAndGraphic),
			"TextAndGraphic,Document"},
		{MakeIntents(Document, Photo, TextAndGraphic),
			"Document,Photo,TextAndGraphic"},
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
