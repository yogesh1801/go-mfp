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

// TestDimensionFromDots tests DimensionFromDots function
func TestDimensionFromDots(t *testing.T) {
	type testData struct {
		dpi  int       // Input DPI
		dots int       // Input dots output
		dim  Dimension // Expected output
	}

	tests := []testData{
		{
			dpi:  300,
			dots: 300,
			dim:  Inch,
		},

		{
			dpi:  600,
			dots: 300,
			dim:  Inch / 2,
		},

		{
			dpi:  150,
			dots: 300,
			dim:  Inch * 2,
		},
	}

	for _, test := range tests {
		dim := DimensionFromDots(test.dpi, test.dots)
		if dim != test.dim {
			println("    ->", dim, test.dim)
			t.Errorf("DimensionFromDots(%d,%d):\n"+
				"expected: %d\n"+
				"present:  %d",
				test.dpi, test.dots, test.dim, dim)
		}
	}
}
