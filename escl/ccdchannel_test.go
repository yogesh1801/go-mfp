// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scan CCD channel test

package escl

import "testing"

var testCCDChannel = testEnum[CCDChannel]{
	decodeStr: DecodeCCDChannel,
	decodeXML: decodeCCDChannel,
	dataset: []testEnumData[CCDChannel]{
		{Red, "Red"},
		{Green, "Green"},
		{Blue, "Blue"},
		{NTSC, "NTSC"},
		{GrayCcd, "GrayCcd"},
		{GrayCcdEmulated, "GrayCcdEmulated"},
	},
}

// TestCCDChannel tests [CCDChannel] common methods and functions.
func TestCCDChannel(t *testing.T) {
	testCCDChannel.run(t)
}
