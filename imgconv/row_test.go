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

// TestRowSliceCopy tests Row.Slice and Row.Copy
func TestRowSliceCopy(t *testing.T) {
	type testData struct {
		model  color.Model
		sample Row
	}

	tests := []testData{
		{
			model: color.GrayModel,
			sample: RowGray8{
				color.Gray{0x00},
				color.Gray{0x20},
				color.Gray{0x40},
				color.Gray{0x60},
				color.Gray{0x80},
				color.Gray{0xa0},
				color.Gray{0xc0},
				color.Gray{0xe0},
			},
		},

		{
			model: color.Gray16Model,
			sample: RowGray16{
				color.Gray16{0x0000},
				color.Gray16{0x2000},
				color.Gray16{0x4000},
				color.Gray16{0x6000},
				color.Gray16{0x8000},
				color.Gray16{0xa000},
				color.Gray16{0xc000},
				color.Gray16{0xe000},
			},
		},

		{
			model: color.RGBAModel,
			sample: RowRGBA32{
				color.RGBA{0x00, 0x00, 0x00, 0x00},
				color.RGBA{0x20, 0x20, 0x20, 0x20},
				color.RGBA{0x40, 0x40, 0x40, 0x40},
				color.RGBA{0x60, 0x60, 0x60, 0x60},
				color.RGBA{0x80, 0x80, 0x80, 0x80},
				color.RGBA{0xa0, 0xa0, 0xa0, 0xa0},
				color.RGBA{0xc0, 0xc0, 0xc0, 0xc0},
				color.RGBA{0xe0, 0xe0, 0xe0, 0xe0},
			},
		},

		{
			model: color.RGBA64Model,
			sample: RowRGBA64{
				color.RGBA64{0x0000, 0x0000, 0x0000, 0x0000},
				color.RGBA64{0x2000, 0x2000, 0x2000, 0x2000},
				color.RGBA64{0x4000, 0x4000, 0x4000, 0x4000},
				color.RGBA64{0x6000, 0x6000, 0x6000, 0x6000},
				color.RGBA64{0x8000, 0x8000, 0x8000, 0x8000},
				color.RGBA64{0xa000, 0xa000, 0xa000, 0xa000},
				color.RGBA64{0xc000, 0xc000, 0xc000, 0xc000},
				color.RGBA64{0xe000, 0xe000, 0xe000, 0xe000},
			},
		},
	}

	for _, test := range tests {
		name := reflect.TypeOf(test.sample).String()
		fullWid := reflect.ValueOf(test.sample).Len()
		wid := fullWid - 3
		mid := wid / 2

		source := test.sample
		expected := reflect.ValueOf(test.sample).Slice(0, wid).Interface()

		// Test copying from the slice of the matched type
		row := NewRow(test.model, wid)
		row.Copy(source.Slice(0, mid))
		row.Slice(mid, row.Width()).Copy(source.Slice(mid, fullWid))

		diff := testutils.Diff(expected, row)
		if diff != "" {
			t.Errorf("%s: direct Copy test failed:\n%s",
				name, diff)
		}

		// Test copying from the wrapped slice - it prohibits
		// the direct copying
		type wrap struct{ Row }

		row = NewRow(test.model, wid)
		row.Copy(wrap{source.Slice(0, mid)})
		row.Slice(mid, row.Width()).Copy(wrap{source.Slice(mid, fullWid)})

		diff = testutils.Diff(expected, row)
		if diff != "" {
			t.Errorf("%s: converted Copy test failed:\n%s",
				name, diff)
		}
	}
}
