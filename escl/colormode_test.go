// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scan color mode

package escl

import "testing"

// TestColorModeString tests ColorMode.String
func TestColorModeString(t *testing.T) {
	type testData struct {
		cm ColorMode
		s  string
	}

	tests := []testData{
		{BlackAndWhite1, "BlackAndWhite1"},
		{Grayscale8, "Grayscale8"},
		{Grayscale16, "Grayscale16"},
		{RGB24, "RGB24"},
		{RGB48, "RGB48"},
		{-1, "Unknown"},
	}

	for _, test := range tests {
		s := test.cm.String()
		if s != test.s {
			t.Errorf("%d: extected %q, present %q",
				int(test.cm), test.s, s)
		}
	}
}

// TestDecodeColorMode tests DecodeColorMode
func TestDecodeColorMode(t *testing.T) {
	type testData struct {
		cm ColorMode
		s  string
	}

	tests := []testData{
		{BlackAndWhite1, "BlackAndWhite1"},
		{Grayscale8, "Grayscale8"},
		{Grayscale16, "Grayscale16"},
		{RGB24, "RGB24"},
		{RGB48, "RGB48"},
		{UnknownColorMode, "XXX"},
	}

	for _, test := range tests {
		cm := DecodeColorMode(test.s)
		if cm != test.cm {
			t.Errorf("%q: extected %q, present %q",
				test.s, test.cm, cm)
		}
	}
}
