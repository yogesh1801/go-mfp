// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Units for coordinates and resolutions test.

package escl

import "testing"

var testUnits = testEnum[Units]{
	decodeStr: DecodeUnits,
	decodeXML: decodeUnits,
	ns:        NsScan,
	dataset: []testEnumData[Units]{
		{ThreeHundredthsOfInches, "ThreeHundredthsOfInches"},
	},
}

// TestUnits tests [Units] common methods and functions.
func TestUnits(t *testing.T) {
	testUnits.run(t)
}
