// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scanner feed directions test

package escl

import "testing"

var testFeedDirection = testEnum[FeedDirection]{
	decodeStr: DecodeFeedDirection,
	decodeXML: decodeFeedDirection,
	ns:        NsScan,
	dataset: []testEnumData[FeedDirection]{
		{LongEdgeFeed, "LongEdgeFeed"},
		{ShortEdgeFeed, "ShortEdgeFeed"},
	},
}

// TestFeedDirection tests [FeedDirection] common methods and functions.
func TestFeedDirection(t *testing.T) {
	testFeedDirection.run(t)
}
