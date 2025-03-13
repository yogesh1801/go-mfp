// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Resolution test

package abstract

import "testing"

// TestResolutionRangeIsZero tests ResolutionRange.IsZero method
func TestResolutionRangeIsZero(t *testing.T) {
	type testData struct {
		res  ResolutionRange
		zero bool
	}

	tests := []testData{
		{
			res:  ResolutionRange{},
			zero: true,
		},

		{
			res:  ResolutionRange{XMin: 100},
			zero: false,
		},
		{
			res:  ResolutionRange{XMax: 100},
			zero: false,
		},
		{
			res:  ResolutionRange{XStep: 100},
			zero: false,
		},
		{
			res:  ResolutionRange{XNormal: 100},
			zero: false,
		},

		{
			res:  ResolutionRange{YMin: 100},
			zero: false,
		},
		{
			res:  ResolutionRange{YMax: 100},
			zero: false,
		},
		{
			res:  ResolutionRange{YStep: 100},
			zero: false,
		},
		{
			res:  ResolutionRange{YNormal: 100},
			zero: false,
		},
	}

	for _, test := range tests {
		zero := test.res.IsZero()

		if zero != test.zero {
			t.Errorf("ResolutionRange%v.IsZero:\n"+
				"expected: %v\n"+
				"present:  %v",
				test.res, test.zero, zero)
		}
	}
}
