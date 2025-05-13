// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// PNG Decoder and Encoder

package imgconv

import (
	"bytes"
	"image/png"
	"testing"

	"github.com/OpenPrinting/go-mfp/internal/testutils"
)

func TestPNG(t *testing.T) {
	in := bytes.NewReader(testutils.Images.PNG100x75rgb8)
	decoder, err := NewPNGDecoder(in)
	if err != nil {
		panic(err)
	}

	image, err := png.Decode(bytes.NewReader(testutils.Images.PNG100x75rgb8))

	if decoder.Bounds() != image.Bounds() {
		t.Errorf("Bounds mismatch:\n"+
			"expected: %v\n"+
			"present:  %v\n",
			image.Bounds(),
			decoder.Bounds(),
		)
	}

	bounds := decoder.Bounds()

	for y := 0; y < bounds.Dy(); y++ {
		for x := 0; x < bounds.Dx(); x++ {
			expected := image.At(x, y)
			present := decoder.At(x, y)
			if !colorEqual(expected, present) {
				t.Errorf("At(%d,%d) mismatch:\n"+
					"expected: %v\n"+
					"present:  %v\n",
					x, y,
					expected, present,
				)
				return
			}
		}
	}
}
