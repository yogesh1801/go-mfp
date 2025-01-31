// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scan color mode

package escl

import "testing"

var testColorMode = testEnum[ColorMode]{
	decodeStr: DecodeColorMode,
	decodeXML: decodeColorMode,
	ns:        NsScan,
	dataset: []testEnumData[ColorMode]{
		{BlackAndWhite1, "BlackAndWhite1"},
		{Grayscale8, "Grayscale8"},
		{Grayscale16, "Grayscale16"},
		{RGB24, "RGB24"},
		{RGB48, "RGB48"},
	},
}

// TestColorMode tests [ColorMode] common methods and functions.
func TestColorMode(t *testing.T) {
	testColorMode.run(t)
}
