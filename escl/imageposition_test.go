// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Image position tests

package escl

import "testing"

var testImagePosition = testEnum[ImagePosition]{
	decodeStr: DecodeImagePosition,
	decodeXML: decodeImagePosition,
	dataset: []testEnumData[ImagePosition]{
		{Left, "Left"},
		{Right, "Right"},
		{Top, "Top"},
		{Bottom, "Bottom"},
		{Center, "Center"},
	},
}

// TestImagePosition tests [ImagePosition] common methods and functions.
func TestImagePosition(t *testing.T) {
	testImagePosition.run(t)
}
