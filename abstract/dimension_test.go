// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Dimension test

package abstract

import "testing"

// TestDimension tests Dimension.Dots
func TestDimensionDots(t *testing.T) {
	type testData struct {
		dim  Dimension // Input dimension
		dpi  int       // Dimension.Dots() parameter
		dots int       // Dimension.Dots output
	}

	tests := []testData{
		{
			dim:  Inch,
			dpi:  300,
			dots: 300,
		},
		{
			dim:  Inch + Inch/300,
			dpi:  300,
			dots: 301,
		},
		{
			dim:  Inch - Inch/300,
			dpi:  300,
			dots: 299,
		},
		{
			dim:  Inch + Inch/600,
			dpi:  300,
			dots: 300,
		},
		{
			dim:  Inch + Inch/600 + 1,
			dpi:  300,
			dots: 301,
		},
	}

	for _, test := range tests {
		dots := test.dim.Dots(test.dpi)
		if dots != test.dots {
			t.Errorf("Dimension(%d).Dots(%d):\n"+
				"expected: %d\n"+
				"present:  %d",
				test.dim, test.dpi, test.dots, dots)
		}
	}
}
