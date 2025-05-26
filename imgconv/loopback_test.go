// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Decoder/Encoder loopback test

package imgconv

import (
	"image/color"
	"io"
	"testing"

	"github.com/OpenPrinting/go-mfp/internal/testutils"
)

// TestLoopback tests Decoder/Encoder loopback
func TestLoopback(t *testing.T) {
	wid := 100
	hei := 75
	model := color.RGBAModel
	c := color.RGBA{0x10, 0x20, 0x30, 0xff}

	decoder, encoder := NewLoopback(wid, hei, model)
	row1 := decoder.NewRow()
	row2 := decoder.NewRow()
	row1.Fill(c)

	// Check decoder/encoder parameters
	if decoder.ColorModel() != model {
		t.Errorf("Loopback Decoder.ColorModel:\n"+
			"expected: %v\n"+
			"present:  %v\n",
			model, decoder.ColorModel())
	}

	if encoder.ColorModel() != model {
		t.Errorf("Loopback Encoder.ColorModel:\n"+
			"expected: %v\n"+
			"present:  %v\n",
			model, encoder.ColorModel())
	}

	wid2, hei2 := decoder.Size()
	if wid != wid2 || hei != hei2 {
		t.Errorf("Loopback Decoder.Size:\n"+
			"expected: %d x %d\n"+
			"present:  %d x %d\n",
			wid, hei, wid2, hei2)
	}

	wid2, hei2 = encoder.Size()
	if wid != wid2 || hei != hei2 {
		t.Errorf("Loopback Encoder.Size:\n"+
			"expected: %d x %d\n"+
			"present:  %d x %d\n",
			wid, hei, wid2, hei2)
	}

	// Up to hei rows must be passed transparently
	for y := 0; y < hei; y++ {
		row2.Fill(color.Opaque)
		err := encoder.Write(row1)
		if err != nil {
			t.Errorf("Loopback: unexpected Decoder.Write error: %s",
				err)
		}

		_, err = decoder.Read(row2)
		if err != nil {
			t.Errorf("Loopback: unexpected Encoder.Read error: %s",
				err)
		}

		diff := testutils.Diff(row1, row2)
		if diff != "" {
			t.Errorf("Loopback: read/write mismatch:\n%s", diff)
		}
	}

	// Subsequent writes must be ignored and reads must return io.EOF
	for y := 0; y < hei; y++ {
		err := encoder.Write(row1)
		if err != nil {
			t.Errorf("Loopback: unexpected Decoder.Write error: %s",
				err)
		}

		_, err = decoder.Read(row2)
		if err != io.EOF {
			t.Errorf("Loopback: Encoder.Read mismatch:\n"+
				"expected: %v\n"+
				"present:  %v\n",
				io.EOF, err)
		}
	}

	decoder.Close()
	encoder.Close()

	// Now test loopback Decoder.Read, when Encoder is closed
	decoder, encoder = NewLoopback(wid, hei, model)
	lim := hei / 2

	for y := 0; y < hei; y++ {
		encoder.Write(row1)
		if y == lim {
			encoder.Close()
		}

		_, err := decoder.Read(row2)
		var errExpected error
		if y > lim {
			errExpected = io.ErrUnexpectedEOF
		}

		if err != errExpected {
			t.Errorf("Loopback: Encoder.Read mismatch:\n"+
				"expected: %v\n"+
				"present:  %v\n",
				errExpected, err)
		}
	}

	decoder.Close()
	encoder.Close()

	// Encoder.Write must not be blocked, if Decoder is
	// closed and inactive

	decoder, encoder = NewLoopback(wid, hei, model)
	decoder.Close()

	for y := 0; y < hei; y++ {
		encoder.Write(row1)
	}

	encoder.Close()
}
