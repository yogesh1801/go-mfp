// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scanner feed directions test

package escl

import "testing"

// TestDecodeFeedDirectionString tests FeedDirection.String
func TestDecodeFeedDirectionString(t *testing.T) {
	type testData struct {
		feed FeedDirection
		s    string
	}

	tests := []testData{
		{LongEdgeFeed, "LongEdgeFeed"},
		{ShortEdgeFeed, "ShortEdgeFeed"},
		{-1, "Unknown"},
	}

	for _, test := range tests {
		s := test.feed.String()
		if s != test.s {
			t.Errorf("%d: extected %q, present %q",
				int(test.feed), test.s, s)
		}
	}
}

// TestDecodeFeedDirection tests DecodeFeedDirection
func TestDecodeFeedDirection(t *testing.T) {
	type testData struct {
		feed FeedDirection
		s    string
	}

	tests := []testData{
		{LongEdgeFeed, "LongEdgeFeed"},
		{ShortEdgeFeed, "ShortEdgeFeed"},
		{UnknownFeedDirection, "XXX"},
	}

	for _, test := range tests {
		feed := DecodeFeedDirection(test.s)
		if feed != test.feed {
			t.Errorf("%q: extected %q, present %q",
				test.s, test.feed, feed)
		}
	}
}
