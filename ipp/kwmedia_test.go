// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// KwMedia test

package ipp

import "testing"

// TestKwColorLocalizedName tests (KwColor) LocalizedName() method
func TestKwMediaSize(t *testing.T) {
	type testData struct {
		media    KwMedia
		wid, hei int
	}

	tests := []testData{
		{"iso_a4_210x297mm", 21000, 29700},
		{"na_govt-letter_8x10in", 20320, 25400},
		{"unknown", -1, -1},
	}

	for _, test := range tests {
		wid, hei := test.media.Size()
		if wid != test.wid || hei != test.hei {
			t.Errorf("(KwMedia) Size(): %q: expected %dx%d, present %dx%d",
				test.media, test.wid, test.hei, wid, hei)
		}
	}
}
