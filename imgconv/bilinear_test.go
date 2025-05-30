// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Bi-linear interpolation tests

package imgconv

import (
	"testing"

	"github.com/OpenPrinting/go-mfp/internal/testutils"
)

// TestMakeScaleCoefficients tests makeScaleCoefficients function
func TestMakeBLCoefficients(t *testing.T) {
	type testData struct {
		slen, dlen int          // Source/destination length
		coeffs     []scaleCoeff // Expected coefficients
	}

	tests := []testData{
		{
			slen:   0,
			dlen:   0,
			coeffs: nil,
		},

		{
			slen: 1,
			dlen: 1,
			coeffs: []scaleCoeff{
				{S: 0, D: 0, W: 1.0},
			},
		},

		{
			slen: 5,
			dlen: 5,
			coeffs: []scaleCoeff{
				{S: 0, D: 0, W: 1.0},
				{S: 1, D: 1, W: 1.0},
				{S: 2, D: 2, W: 1.0},
				{S: 3, D: 3, W: 1.0},
				{S: 4, D: 4, W: 1.0},
			},
		},

		{
			slen: 5,
			dlen: 1,
			coeffs: []scaleCoeff{
				{D: 0, S: 0, W: 0.2},
				{D: 0, S: 1, W: 0.2},
				{D: 0, S: 2, W: 0.2},
				{D: 0, S: 3, W: 0.2},
				{D: 0, S: 4, W: 0.2},
			},
		},

		{
			slen: 1,
			dlen: 5,
			coeffs: []scaleCoeff{
				{S: 0, D: 0, W: 1},
				{S: 0, D: 1, W: 1},
				{S: 0, D: 2, W: 1},
				{S: 0, D: 3, W: 1},
				{S: 0, D: 4, W: 1},
			},
		},

		{
			//    0 --------------> 0
			// 1/3 \ \ 2.3
			//      \ ------------> 1
			//       X
			// 1/3  / ------------> 2
			//     / / 2/3
			//    1 --------------> 3
			slen: 2,
			dlen: 4,
			coeffs: []scaleCoeff{
				{S: 0, D: 0, W: 1},
				{S: 0, D: 1, W: 2.0 / 3.0},
				{S: 1, D: 1, W: 1.0 / 3.0},
				{S: 0, D: 2, W: 1.0 / 3.0},
				{S: 1, D: 2, W: 2.0 / 3.0},
				{S: 1, D: 3, W: 1.0},
			},
		},

		{
			//   0 --------------> 0
			//     \ 1/2
			//      -------------> 1
			//     / 1/2
			//   1 --------------> 2
			//    \  1/2
			//     --------------> 3
			//    /  1/2
			//   2 --------------> 4
			slen: 3,
			dlen: 5,
			coeffs: []scaleCoeff{
				{S: 0, D: 0, W: 1.0},
				{S: 0, D: 1, W: 0.5},
				{S: 1, D: 1, W: 0.5},
				{S: 1, D: 2, W: 1.0},
				{S: 1, D: 3, W: 0.5},
				{S: 2, D: 3, W: 0.5},
				{S: 2, D: 4, W: 1.0},
			},
		},

		{
			//   S:     | 0 | 1 | 2 | 3 |
			//   D: |     0     |     1     |
			slen: 4,
			dlen: 2,
			coeffs: []scaleCoeff{
				{S: 0, D: 0, W: 1. / 3.},
				{S: 1, D: 0, W: 2. / 3.},
				{S: 2, D: 1, W: 2. / 3.},
				{S: 3, D: 1, W: 1. / 3.},
			},
		},

		{
			//   S:    | 0 | 1 | 2 | 3 | 4 |
			//   D:  |   0   |   1   |   2   |
			slen: 5,
			dlen: 3,
			coeffs: []scaleCoeff{
				{S: 0, D: 0, W: 2. / 4.},
				{S: 1, D: 0, W: 2. / 4.},
				{S: 1, D: 1, W: 1. / 4.},
				{S: 2, D: 1, W: 2. / 4.},
				{S: 3, D: 1, W: 1. / 4.},
				{S: 3, D: 2, W: 2. / 4.},
				{S: 4, D: 2, W: 2. / 4.},
			},
		},
	}

	//tests = tests[:len(tests)-2]
	//tests = tests[len(tests)-2:]
	//tests = tests[:1]

	for _, test := range tests {
		coeffs := makeScaleCoefficients(test.slen, test.dlen)

		var expected, present []string

		for _, sc := range test.coeffs {
			expected = append(expected, sc.String())
		}

		for _, sc := range coeffs {
			present = append(present, sc.String())
		}

		diff := testutils.Diff(expected, present)
		if diff != "" {
			t.Errorf("%d->%d:\n%s", test.slen, test.dlen, diff)
		}
	}
}
