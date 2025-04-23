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
		{
			rng:  Range{Step: 1},
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

// TestRangeWithin tests Range.Within method
func TestRangeWithin(t *testing.T) {
	type testData struct {
		rng    Range
		v      int
		within bool
	}

	tests := []testData{
		{
			rng:    Range{},
			v:      0,
			within: false,
		},

		{
			rng:    Range{},
			v:      1,
			within: false,
		},

		{
			rng:    Range{Min: 1, Max: 10},
			v:      0,
			within: false,
		},

		{
			rng:    Range{Min: 1, Max: 10},
			v:      1,
			within: true,
		},

		{
			rng:    Range{Min: 1, Max: 10},
			v:      5,
			within: true,
		},

		{
			rng:    Range{Min: 1, Max: 10},
			v:      10,
			within: true,
		},

		{
			rng:    Range{Min: 1, Max: 10},
			v:      11,
			within: false,
		},

		{
			rng:    Range{Min: 1, Max: 11, Step: 2},
			v:      1,
			within: true,
		},

		{
			rng:    Range{Min: 1, Max: 11, Step: 2},
			v:      2,
			within: false,
		},

		{
			rng:    Range{Min: 1, Max: 11, Step: 2},
			v:      3,
			within: true,
		},

		{
			rng:    Range{Min: 1, Max: 11, Step: 2},
			v:      11,
			within: true,
		},
	}

	for _, test := range tests {
		within := test.rng.Within(test.v)

		if within != test.within {
			t.Errorf("Range%v.Within(%v):\n"+
				"expected: %v\n"+
				"present:  %v",
				test.rng, test.v, test.within, within)
		}
	}
}
