// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scan CCD channel test

package escl

import "testing"

var testCcdChannel = testEnum[CcdChannel]{
	decodeStr: DecodeCcdChannel,
	decodeXML: decodeCcdChannel,
	ns:        NsScan,
	dataset: []testEnumData[CcdChannel]{
		{Red, "Red"},
		{Green, "Green"},
		{Blue, "Blue"},
		{NTSC, "NTSC"},
		{GrayCcd, "GrayCcd"},
		{GrayCcdEmulated, "GrayCcdEmulated"},
	},
}

// TestCcdChannel tests [CcdChannel] common methods and functions.
func TestCcdChannel(t *testing.T) {
	testCcdChannel.run(t)
}
