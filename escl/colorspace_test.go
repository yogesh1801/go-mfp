// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scan color space test

package escl

import "testing"

// TestColorSpaceString tests ColorSpace.String
func TestColorSpaceString(t *testing.T) {
	type testData struct {
		sps ColorSpace
		s   string
	}

	tests := []testData{
		{SRGB, "sRGB"},
		{-1, "Unknown"},
	}

	for _, test := range tests {
		s := test.sps.String()
		if s != test.s {
			t.Errorf("%d: extected %q, present %q",
				int(test.sps), test.s, s)
		}
	}
}

// TestDecodeColorSpace tests DecodeColorSpace
func TestDecodeColorSpace(t *testing.T) {
	type testData struct {
		sps ColorSpace
		s   string
	}

	tests := []testData{
		{SRGB, "sRGB"},
		{UnknownColorSpace, "XXX"},
	}

	for _, test := range tests {
		sps := DecodeColorSpace(test.s)
		if sps != test.sps {
			t.Errorf("%q: extected %q, present %q",
				test.s, test.sps, sps)
		}
	}
}
