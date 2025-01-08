// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scan CCD channel test

package escl

import "testing"

// TestCcdChannelString tests CcdChannel.String
func TestCcdChannelString(t *testing.T) {
	type testData struct {
		ccd CcdChannel
		s   string
	}

	tests := []testData{
		{Red, "Red"},
		{Green, "Green"},
		{Blue, "Blue"},
		{NTSC, "NTSC"},
		{GrayCcd, "GrayCcd"},
		{GrayCcdEmulated, "GrayCcdEmulated"},
		{-1, "Unknown"},
	}

	for _, test := range tests {
		s := test.ccd.String()
		if s != test.s {
			t.Errorf("%d: extected %q, present %q",
				int(test.ccd), test.s, s)
		}
	}
}

// TestDecodeCcdChannel tests DecodeCcdChannel
func TestDecodeCcdChannel(t *testing.T) {
	type testData struct {
		ccd CcdChannel
		s   string
	}

	tests := []testData{
		{Red, "Red"},
		{Green, "Green"},
		{Blue, "Blue"},
		{NTSC, "NTSC"},
		{GrayCcd, "GrayCcd"},
		{GrayCcdEmulated, "GrayCcdEmulated"},
		{UnknownCcdChannel, "XXX"},
	}

	for _, test := range tests {
		ccd := DecodeCcdChannel(test.s)
		if ccd != test.ccd {
			t.Errorf("%q: extected %q, present %q",
				test.s, test.ccd, ccd)
		}
	}
}
