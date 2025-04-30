// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scan CCD channel test

package escl

import "testing"

var testBinaryRendering = testEnum[BinaryRendering]{
	decodeStr: DecodeBinaryRendering,
	decodeXML: decodeBinaryRendering,
	dataset: []testEnumData[BinaryRendering]{
		{Halftone, "Halftone"},
		{Threshold, "Threshold"},
	},
}

// TestBinaryRendering tests [BinaryRendering] common methods and functions.
func TestBinaryRendering(t *testing.T) {
	testBinaryRendering.run(t)
}
