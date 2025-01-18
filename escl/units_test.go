// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Units for coordinates and resolutions test.

package escl

import "testing"

// TestDecodeUnitsString tests Units.String
func TestDecodeUnitsString(t *testing.T) {
	type testData struct {
		units Units
		s     string
	}

	tests := []testData{
		{ThreeHundredthsOfInches, "ThreeHundredthsOfInches"},
		{-1, "Unknown"},
	}

	for _, test := range tests {
		s := test.units.String()
		if s != test.s {
			t.Errorf("%d: extected %q, present %q",
				int(test.units), test.s, s)
		}
	}
}

// TestDecodeUnits tests DecodeUnits
func TestDecodeUnits(t *testing.T) {
	type testData struct {
		units Units
		s     string
	}

	tests := []testData{
		{ThreeHundredthsOfInches, "ThreeHundredthsOfInches"},
		{UnknownUnits, "XXX"},
	}

	for _, test := range tests {
		units := DecodeUnits(test.s)
		if units != test.units {
			t.Errorf("%q: extected %q, present %q",
				test.s, test.units, units)
		}
	}
}
