// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Conversions between Rows and byte slices tests

package imgconv

import (
	"bytes"
	"image/color"
	"reflect"
	"strings"
	"testing"

	"github.com/OpenPrinting/go-mfp/internal/testutils"
)

// TestBytesToRow tests bytesXXXtoRow and bytesXXXfromRow family of functions
func TestBytesToFromRow(t *testing.T) {
	type testData struct {
		name  string                // Test name
		bytes []byte                // Inputs byte slice
		row   Row                   // Expected row
		to    func(Row, []byte) int // []byte->Row
		from  func([]byte, Row) int // Row->[]byte
	}

	tests := []testData{
		{
			name:  "Gray8",
			bytes: []byte{0x00, 0x10, 0x20, 0x30},
			row: RowGray8{
				color.Gray{0x00},
				color.Gray{0x10},
				color.Gray{0x20},
				color.Gray{0x30},
			},
			to:   bytesGray8toRow,
			from: bytesGray8fromRow,
		},

		{
			name:  "Gray8",
			bytes: []byte{0x00, 0x10, 0x20, 0x30},
			row: RowGrayFP32{
				0x00 / 255.,
				0x10 / 255.,
				0x20 / 255.,
				0x30 / 255.,
			},
			to:   bytesGray8toRow,
			from: bytesGray8fromRow,
		},

		{
			name: "Gray16BE",
			bytes: []byte{
				0x10, 0x01,
				0x20, 0x02,
				0x30, 0x03,
				0x40, 0x04,
			},
			row: RowGray16{
				color.Gray16{0x1001},
				color.Gray16{0x2002},
				color.Gray16{0x3003},
				color.Gray16{0x4004},
			},
			to:   bytesGray16BEtoRow,
			from: bytesGray16BEfromRow,
		},

		{
			name: "Gray16BE",
			bytes: []byte{
				0x10, 0x01,
				0x20, 0x02,
				0x30, 0x03,
				0x40, 0x04,
			},
			row: RowGrayFP32{
				0x1001 / 65535.,
				0x2002 / 65535.,
				0x3003 / 65535.,
				0x4004 / 65535.,
			},
			to:   bytesGray16BEtoRow,
			from: bytesGray16BEfromRow,
		},

		{
			name:  "RGB8",
			bytes: []byte{0x00, 0x10, 0x20, 0x30, 0x40, 0x50},
			row: RowRGBA32{
				color.RGBA{R: 0x00, G: 0x10, B: 0x20, A: 0xff},
				color.RGBA{R: 0x30, G: 0x40, B: 0x50, A: 0xff},
			},
			to:   bytesRGB8toRow,
			from: bytesRGB8fromRow,
		},

		{
			name:  "RGB8",
			bytes: []byte{0x00, 0x10, 0x20, 0x30, 0x40, 0x50},
			row: RowRGBAFP32{
				0x00 / 255.,
				0x10 / 255.,
				0x20 / 255.,
				0xff / 255.,
				0x30 / 255.,
				0x40 / 255.,
				0x50 / 255.,
				0xff / 255.,
			},
			to:   bytesRGB8toRow,
			from: bytesRGB8fromRow,
		},

		{
			name: "RGB16",
			bytes: []byte{
				0x00, 0x00,
				0x10, 0x01,
				0x20, 0x02,
				0x30, 0x03,
				0x40, 0x04,
				0x50, 0x05,
			},
			row: RowRGBA64{
				color.RGBA64{
					R: 0x0000,
					G: 0x1001,
					B: 0x2002,
					A: 0xffff,
				},
				color.RGBA64{
					R: 0x3003,
					G: 0x4004,
					B: 0x5005,
					A: 0xffff,
				},
			},
			to:   bytesRGB16toRow,
			from: bytesRGB16fromRow,
		},

		{
			name: "RGB16",
			bytes: []byte{
				0x00, 0x00,
				0x10, 0x01,
				0x20, 0x02,
				0x30, 0x03,
				0x40, 0x04,
				0x50, 0x05,
			},
			row: RowRGBAFP32{
				0x0000 / 65535.,
				0x1001 / 65535.,
				0x2002 / 65535.,
				0xffff / 65535.,
				0x3003 / 65535.,
				0x4004 / 65535.,
				0x5005 / 65535.,
				0xffff / 65535.,
			},
			to:   bytesRGB16toRow,
			from: bytesRGB16fromRow,
		},

		{
			name: "RGBA8",
			bytes: []byte{
				0x00, 0x10, 0x20, 0x30,
				0x40, 0x50, 0x60, 0x70,
			},
			row: RowRGBA32{
				color.RGBA{R: 0x00, G: 0x10, B: 0x20, A: 0x30},
				color.RGBA{R: 0x40, G: 0x50, B: 0x60, A: 0x70},
			},
			to:   bytesRGBA8toRow,
			from: bytesRGBA8fromRow,
		},

		{
			name: "RGBA8",
			bytes: []byte{
				0x00, 0x10, 0x20, 0x30,
				0x40, 0x50, 0x60, 0x70,
			},
			row: RowRGBAFP32{
				0x00 / 255.,
				0x10 / 255.,
				0x20 / 255.,
				0x30 / 255.,
				0x40 / 255.,
				0x50 / 255.,
				0x60 / 255.,
				0x70 / 255.,
			},
			to:   bytesRGBA8toRow,
			from: bytesRGBA8fromRow,
		},

		{
			name: "RGBA16",
			bytes: []byte{
				0x00, 0x00,
				0x10, 0x01,
				0x20, 0x02,
				0x30, 0x03,
				0x40, 0x04,
				0x50, 0x05,
				0x60, 0x06,
				0x70, 0x07,
			},
			row: RowRGBA64{
				color.RGBA64{
					R: 0x0000,
					G: 0x1001,
					B: 0x2002,
					A: 0x3003,
				},
				color.RGBA64{
					R: 0x4004,
					G: 0x5005,
					B: 0x6006,
					A: 0x7007,
				},
			},
			to:   bytesRGBA16toRow,
			from: bytesRGBA16fromRow,
		},

		{
			name: "RGBA16",
			bytes: []byte{
				0x00, 0x00,
				0x10, 0x01,
				0x20, 0x02,
				0x30, 0x03,
				0x40, 0x04,
				0x50, 0x05,
				0x60, 0x06,
				0x70, 0x07,
			},
			row: RowRGBAFP32{
				0x0000 / 65535.,
				0x1001 / 65535.,
				0x2002 / 65535.,
				0x3003 / 65535.,
				0x4004 / 65535.,
				0x5005 / 65535.,
				0x6006 / 65535.,
				0x7007 / 65535.,
			},
			to:   bytesRGBA16toRow,
			from: bytesRGBA16fromRow,
		},
	}

	newrow := func(row Row) Row {
		l := reflect.ValueOf(row).Len()
		return reflect.MakeSlice(reflect.TypeOf(row), l, l).
			Interface().(Row)
	}

	for _, test := range tests {
		// Prepare test name
		name := reflect.TypeOf(test.row).String()
		if i := strings.IndexByte(name, '.'); i >= 0 {
			name = name[i+1:]
		}
		name = test.name + "/" + name

		// Perform bytes->Row conversion the "normal" way
		row := newrow(test.row)
		wid := test.to(row, test.bytes)
		if wid != test.row.Width() {
			t.Errorf("%s: length mismatch:\n"+
				"expected: %d\n"+
				"present:  %d\n",
				name, test.row.Width(), wid)
		}

		diff := testutils.Diff(test.row, row)
		if diff != "" {
			t.Errorf("%s: data mismatch:\n%s", name, diff)
		}

		// Wrap the row, which effectively prohibits the fast path
		type wrap struct{ Row }
		row = newrow(test.row)
		test.to(wrap{row}, test.bytes)

		diff = testutils.Diff(test.row, row)
		if diff != "" {
			t.Errorf("%s: slow path failed:\n%s", name, diff)
		}

		if test.from == nil {
			continue
		}

		// Perform Row->bytes conversion the "normal" way
		buf := make([]byte, len(test.bytes))
		wid = test.from(buf, test.row)
		if wid != len(test.bytes) {
			t.Errorf("%s: length mismatch:\n"+
				"expected: %d\n"+
				"present:  %d\n",
				name, len(test.bytes), wid)
		}

		if !bytes.Equal(buf, test.bytes) {
			t.Errorf("%s: data mismatch:\n"+
				"expected: %x\n"+
				"present:  %x\n",
				name, test.bytes, buf)
		}

		// Wrap the row, which effectively prohibits the fast path
		buf = make([]byte, len(test.bytes))
		test.from(buf, wrap{row})
		buf = buf[:wid]

		if !bytes.Equal(buf, test.bytes) {
			t.Errorf("%s: slow path failed:\n"+
				"expected: %x\n"+
				"present:  %x\n",
				name, test.bytes, buf)
		}
	}
}
