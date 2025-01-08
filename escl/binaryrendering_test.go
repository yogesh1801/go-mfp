// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scan CCD channel test

package escl

import "testing"

// TestBinaryRenderingString tests BinaryRendering.String
func TestBinaryRenderingString(t *testing.T) {
	type testData struct {
		rnd BinaryRendering
		s   string
	}

	tests := []testData{
		{Halftone, "Halftone"},
		{Threshold, "Threshold"},
		{-1, "Unknown"},
	}

	for _, test := range tests {
		s := test.rnd.String()
		if s != test.s {
			t.Errorf("%d: extected %q, present %q",
				int(test.rnd), test.s, s)
		}
	}
}

// TestDecodeBinaryRendering tests DecodeBinaryRendering
func TestDecodeBinaryRendering(t *testing.T) {
	type testData struct {
		rnd BinaryRendering
		s   string
	}

	tests := []testData{
		{Halftone, "Halftone"},
		{Threshold, "Threshold"},
		{UnknownBinaryRendering, "XXX"},
	}

	for _, test := range tests {
		rnd := DecodeBinaryRendering(test.s)
		if rnd != test.rnd {
			t.Errorf("%q: extected %q, present %q",
				test.s, test.rnd, rnd)
		}
	}
}
