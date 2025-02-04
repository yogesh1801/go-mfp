// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// ADF image justification test.

package escl

import "testing"

var testJustification = testEnum[Justification]{
	decodeStr: DecodeJustification,
	decodeXML: decodeJustification,
	ns:        NsScan,
	dataset: []testEnumData[Justification]{
		{Left, "Left"},
		{Right, "Right"},
		{Top, "Top"},
		{Bottom, "Bottom"},
		{Center, "Center"},
	},
}

// TestJustification tests [Justification] common methods and functions.
func TestJustification(t *testing.T) {
	testJustification.run(t)
}
