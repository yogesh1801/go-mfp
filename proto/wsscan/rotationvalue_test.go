// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// Test for rotation value

package wsscan

import "testing"

var testRotationValue = testEnum[RotationValue]{
	decodeStr: DecodeRotationValue,
	decodeXML: decodeRotationValue,
	dataset: []testEnumData[RotationValue]{
		{Rotation0, "0"},
		{Rotation90, "90"},
		{Rotation180, "180"},
		{Rotation270, "270"},
	},
}

// TestRotationValue tests [RotationValue] common methods and functions.
func TestRotationValue(t *testing.T) {
	testRotationValue.run(t)
}
