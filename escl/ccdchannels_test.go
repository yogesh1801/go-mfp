// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Set of CCD channels test

package escl

import (
	"fmt"
	"slices"
	"strings"
	"testing"
)

// TestCcdChannelsAddDel tests CcdChannels.Add and CcdChannels.Del operations
func TestCcdChannelsAddDel(t *testing.T) {
	type testOp struct {
		action string
		val    CcdChannel
	}

	type testData struct {
		seq []testOp
		res CcdChannels
	}

	tests := []testData{
		{
			seq: nil,
			res: CcdChannels{},
		},

		{
			seq: []testOp{
				{"add", Red},
			},
			res: CcdChannels{
				1 << Red,
				UnknownCcdChannel,
			},
		},

		{
			seq: []testOp{
				{"add", Red},
				{"add", Blue},
			},
			res: CcdChannels{
				1<<Red | 1<<Blue,
				UnknownCcdChannel,
			},
		},

		{
			seq: []testOp{
				{"add", Red},
				{"add", Blue},
				{"del", Red},
				{"add", Green},
			},
			res: CcdChannels{
				1<<Blue | 1<<Green,
				UnknownCcdChannel,
			},
		},
	}

	for _, test := range tests {
		var opts CcdChannels
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

// TestMakeCcdChannels tests MakeCcdChannels
func TestMakeCcdChannels(t *testing.T) {
	type testData struct {
		in  []CcdChannel
		res CcdChannels
	}

	tests := []testData{
		{[]CcdChannel{}, CcdChannels{}},
		{[]CcdChannel{Red},
			CcdChannels{
				1 << Red,
				UnknownCcdChannel,
			},
		},
		{[]CcdChannel{Red, Green, Blue},
			CcdChannels{
				1<<Red | 1<<Green | 1<<Blue,
				UnknownCcdChannel,
			},
		},
	}

	for _, test := range tests {
		opts := MakeCcdChannels(test.in...)
		if opts != test.res {
			t.Errorf("\n"+
				"input:    %s\n"+
				"expected: %s\n"+
				"present:  %s",
				test.in, test.res, opts)
		}
	}
}

// TestCcdChannelsString tests CcdChannels.String
func TestCcdChannelsString(t *testing.T) {
	type testData struct {
		opts CcdChannels
		s    string
	}

	tests := []testData{
		{CcdChannels{}, ""},
		{MakeCcdChannels(Red), "Red"},
		{MakeCcdChannels(Red, Blue),
			"Red,Blue"},
		{MakeCcdChannels(Red, Blue),
			"Blue,Red"},
		{MakeCcdChannels(Red, Green, Blue),
			"Red,Green,Blue"},
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
