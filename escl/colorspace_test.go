// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scan color space test

package escl

import "testing"

var testColorSpace = testEnum[ColorSpace]{
	decodeStr: DecodeColorSpace,
	decodeXML: decodeColorSpace,
	ns:        NsScan,
	dataset: []testEnumData[ColorSpace]{
		{SRGB, "sRGB"},
	},
}

// TestColorSpace tests [ColorSpace] common methods and functions.
func TestColorSpace(t *testing.T) {
	testColorSpace.run(t)
}
