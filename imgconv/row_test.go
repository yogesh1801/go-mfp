// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Row tests

package imgconv

import (
	"image/color"
	"reflect"
	"testing"

	"github.com/OpenPrinting/go-mfp/internal/testutils"
)

// TestRowBasic tests Row basic functionality:
//   - NewRow
//   - Row.Width
//   - Row.Fill
func TestRowBasic(t *testing.T) {
	type testData struct {
		model    color.Model
		expected Row
	}

	tests := []testData{
		{
			model: color.GrayModel,
			expected: RowGray8{
				color.Gray{0xff},
				color.Gray{0xff},
				color.Gray{0xff},
				color.Gray{0xff},
				color.Gray{0xff},
			},
		},

		{
			model: color.Gray16Model,
			expected: RowGray16{
				color.Gray16{0xffff},
				color.Gray16{0xffff},
				color.Gray16{0xffff},
				color.Gray16{0xffff},
				color.Gray16{0xffff},
			},
		},

		{
			model: color.RGBAModel,
			expected: RowRGBA32{
				color.RGBA{0xff, 0xff, 0xff, 0xff},
				color.RGBA{0xff, 0xff, 0xff, 0xff},
				color.RGBA{0xff, 0xff, 0xff, 0xff},
				color.RGBA{0xff, 0xff, 0xff, 0xff},
				color.RGBA{0xff, 0xff, 0xff, 0xff},
			},
		},

		{
			model: color.RGBA64Model,
			expected: RowRGBA64{
				color.RGBA64{0xffff, 0xffff, 0xffff, 0xffff},
				color.RGBA64{0xffff, 0xffff, 0xffff, 0xffff},
				color.RGBA64{0xffff, 0xffff, 0xffff, 0xffff},
				color.RGBA64{0xffff, 0xffff, 0xffff, 0xffff},
				color.RGBA64{0xffff, 0xffff, 0xffff, 0xffff},
			},
		},
	}

	for _, test := range tests {
		name := reflect.TypeOf(test.expected).String()
		wid := reflect.ValueOf(test.expected).Len()

		row := NewRow(test.model, wid)
		if row.Width() != wid {
			t.Errorf("%s: Row.Width test failed", name)
		}

		filler := color.Opaque
		row.Fill(filler)

		diff := testutils.Diff(test.expected, row)
		if diff != "" {
			t.Errorf("%s: Row.Fill test failed:\n%s", name, diff)
		}

		for x := 0; x < wid; x++ {
			c := row.At(x)
			expected := test.model.Convert(filler)
			if !colorEqual(c, expected) {
				t.Errorf("%s: Row.At(%d) failed: %#v != %#v",
					name, x, c, expected)
			}
		}
	}
}
