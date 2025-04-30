// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Image edges, for edge auto detection test

package escl

import "testing"

var testSupportedEdge = testEnum[SupportedEdge]{
	decodeStr: DecodeSupportedEdge,
	decodeXML: decodeSupportedEdge,
	dataset: []testEnumData[SupportedEdge]{
		{TopEdge, "TopEdge"},
		{LeftEdge, "LeftEdge"},
		{BottomEdge, "BottomEdge"},
		{RightEdge, "RightEdge"},
	},
}

// TestSupportedEdge tests [SupportedEdge] common methods and functions.
func TestSupportedEdge(t *testing.T) {
	testSupportedEdge.run(t)
}
