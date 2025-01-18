// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// ADF image justification test.

package escl

import "testing"

// TestDecodeJustificationString tests Justification.String
func TestDecodeJustificationString(t *testing.T) {
	type testData struct {
		jst Justification
		s   string
	}

	tests := []testData{
		{Left, "Left"},
		{Right, "Right"},
		{Top, "Top"},
		{Bottom, "Bottom"},
		{Center, "Center"},
		{-1, "Unknown"},
	}

	for _, test := range tests {
		s := test.jst.String()
		if s != test.s {
			t.Errorf("%d: extected %q, present %q",
				int(test.jst), test.s, s)
		}
	}
}

// TestDecodeJustification tests DecodeJustification
func TestDecodeJustification(t *testing.T) {
	type testData struct {
		jst Justification
		s   string
	}

	tests := []testData{
		{Left, "Left"},
		{Right, "Right"},
		{Top, "Top"},
		{Bottom, "Bottom"},
		{Center, "Center"},
		{UnknownJustification, "XXX"},
	}

	for _, test := range tests {
		jst := DecodeJustification(test.s)
		if jst != test.jst {
			t.Errorf("%q: extected %q, present %q",
				test.s, test.jst, jst)
		}
	}
}
