// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scan color mode

package escl

import "testing"

// TestADFOptionString tests ADFOption.String
func TestADFOptionString(t *testing.T) {
	type testData struct {
		opt ADFOption
		s   string
	}

	tests := []testData{
		{DetectPaperLoaded, "DetectPaperLoaded"},
		{SelectSinglePage, "SelectSinglePage"},
		{Duplex, "Duplex"},
		{-1, "Unknown"},
	}

	for _, test := range tests {
		s := test.opt.String()
		if s != test.s {
			t.Errorf("%d: extected %q, present %q",
				int(test.opt), test.s, s)
		}
	}
}

// TestDecodeADFOption tests DecodeADFOption
func TestDecodeADFOption(t *testing.T) {
	type testData struct {
		opt ADFOption
		s   string
	}

	tests := []testData{
		{DetectPaperLoaded, "DetectPaperLoaded"},
		{SelectSinglePage, "SelectSinglePage"},
		{Duplex, "Duplex"},
		{UnknownADFOption, "XXX"},
	}

	for _, test := range tests {
		opt := DecodeADFOption(test.s)
		if opt != test.opt {
			t.Errorf("%q: extected %q, present %q",
				test.s, test.opt, opt)
		}
	}
}
