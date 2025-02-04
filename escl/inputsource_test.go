// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Input source test

package escl

import "testing"

var testInputSource = testEnum[InputSource]{
	decodeStr: DecodeInputSource,
	decodeXML: decodeInputSource,
	ns:        NsScan,
	dataset: []testEnumData[InputSource]{
		{InputPlaten, "Platen"},
		{InputFeeder, "Feeder"},
		{InputCamera, "Camera"},
	},
}

// TestInputSource tests [InputSource] common methods and functions.
func TestInputSource(t *testing.T) {
	testInputSource.run(t)
}
