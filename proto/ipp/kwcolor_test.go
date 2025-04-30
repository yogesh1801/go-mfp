// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// KwColor test

package ipp

import "testing"

// TestKwColorLocalizedName tests (KwColor) LocalizedName() method
func TestKwColorLocalizedName(t *testing.T) {
	type testData struct {
		input  KwColor
		output string
	}

	tests := []testData{
		{"no-color", "Transparent"},
		{"black", "Black"},
		{"clear-black", "Clear Black"},
		{"unknown", ""},
	}

	for _, test := range tests {
		output := test.input.LocalizedName()
		if output != test.output {
			t.Errorf("(KwColor) LocalizedName(): %q->%q; expected %q",
				test.input, output, test.output)
		}
	}
}
