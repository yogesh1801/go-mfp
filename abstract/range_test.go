// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Range test

package abstract

import "testing"

// TestRangeIsZero tests Range.IsZero method
func TestRangeIsZero(t *testing.T) {
	type testData struct {
		rng  Range
		zero bool
	}

	tests := []testData{
		{
			rng:  Range{},
			zero: true,
		},
		{
			rng:  Range{Min: -1},
			zero: false,
		},
		{
			rng:  Range{Max: 1},
			zero: false,
		},
		{
			rng:  Range{Normal: 1},
			zero: false,
		},
	}

	for _, test := range tests {
		zero := test.rng.IsZero()

		if zero != test.zero {
			t.Errorf("Range%v.IsZero:\n"+
				"expected: %v\n"+
				"present:  %v",
				test.rng, test.zero, zero)
		}
	}
}
