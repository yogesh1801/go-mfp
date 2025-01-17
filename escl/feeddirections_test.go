// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Set of CCD feed directions test

package escl

import (
	"fmt"
	"slices"
	"strings"
	"testing"
)

// TestFeedDirectionsAddDel tests FeedDirections.Add and
// FeedDirections.Del operations.
func TestFeedDirectionsAddDel(t *testing.T) {
	type testOp struct {
		action string
		val    FeedDirection
	}

	type testData struct {
		seq []testOp
		res FeedDirections
	}

	tests := []testData{
		{
			seq: nil,
			res: FeedDirections{},
		},

		{
			seq: []testOp{
				{"add", LongEdgeFeed},
			},
			res: FeedDirections{
				1 << LongEdgeFeed,
			},
		},

		{
			seq: []testOp{
				{"add", LongEdgeFeed},
				{"add", ShortEdgeFeed},
			},
			res: FeedDirections{
				1<<LongEdgeFeed | 1<<ShortEdgeFeed,
			},
		},

		{
			seq: []testOp{
				{"add", LongEdgeFeed},
				{"add", ShortEdgeFeed},
				{"del", LongEdgeFeed},
			},
			res: FeedDirections{
				1 << ShortEdgeFeed,
			},
		},
	}

	for _, test := range tests {
		var feeds FeedDirections
		seq := ""

		for _, op := range test.seq {
			seq += fmt.Sprintf("  %s %s\n", op.action, op.val)

			switch op.action {
			case "add":
				feeds.Add(op.val)
			case "del":
				feeds.Del(op.val)

			default:
				panic(fmt.Errorf("unknown action %q", op.action))
			}
		}

		if feeds != test.res {
			t.Errorf("\nfor the sequence:\n%s"+
				"expected: %s\npresent:  %s",
				seq, test.res, feeds)
		}
	}
}

// TestMakeFeedDirections tests MakeFeedDirections
func TestMakeFeedDirections(t *testing.T) {
	type testData struct {
		in  []FeedDirection
		res FeedDirections
	}

	tests := []testData{
		{[]FeedDirection{}, FeedDirections{}},
		{[]FeedDirection{LongEdgeFeed},
			FeedDirections{
				1 << LongEdgeFeed,
			},
		},
		{[]FeedDirection{LongEdgeFeed, ShortEdgeFeed},
			FeedDirections{
				1<<LongEdgeFeed | 1<<ShortEdgeFeed,
			},
		},
	}

	for _, test := range tests {
		feeds := MakeFeedDirections(test.in...)
		if feeds != test.res {
			t.Errorf("\n"+
				"input:    %s\n"+
				"expected: %s\n"+
				"present:  %s",
				test.in, test.res, feeds)
		}
	}
}

// TestFeedDirectionsString tests FeedDirections.String
func TestFeedDirectionsString(t *testing.T) {
	type testData struct {
		feeds FeedDirections
		s     string
	}

	tests := []testData{
		{FeedDirections{}, ""},
		{MakeFeedDirections(LongEdgeFeed), "LongEdgeFeed"},
		{MakeFeedDirections(ShortEdgeFeed), "ShortEdgeFeed"},
		{MakeFeedDirections(LongEdgeFeed, ShortEdgeFeed),
			"LongEdgeFeed,ShortEdgeFeed"},
	}

	for _, test := range tests {
		s := test.feeds.String()

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
				test.feeds, exp, out)
		}
	}
}
