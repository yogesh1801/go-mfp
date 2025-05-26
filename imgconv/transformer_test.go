// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Image transformation filter test.

package imgconv

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"testing"

	"github.com/OpenPrinting/go-mfp/internal/testutils"
	"golang.org/x/image/draw"
)

// TestTransformer tests image filter created by the NewTransformer function
func TestTransformer(t *testing.T) {
	// Create source image
	img, err := png.Decode(
		bytes.NewReader(testutils.Images.PNG100x75rgb8))
	if err != nil {
		panic(err)
	}

	// Define transformation function
	transform := func(dst draw.Image, src image.Image) {
		bounds := src.Bounds()
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				c := src.At(x, y)
				rgb := color.RGBAModel.Convert(c).(color.RGBA)
				rgb.R ^= 0xff
				rgb.G ^= 0xff
				rgb.B ^= 0xff
				dst.Set(x, y, rgb)
			}
		}
	}

	// Transform image for reference
	bounds := img.Bounds()
	expected := image.NewRGBA(bounds)
	transform(expected, img)

	// Make sure they actually differ
	diff := imageDiff(img, expected)
	if diff == "" {
		panic("internal error")
	}

	// Run transformer
	decoder, err := NewPNGDecoder(
		bytes.NewReader(testutils.Images.PNG100x75rgb8))
	if err != nil {
		panic(err)
	}

	transformer := NewTransformer(
		decoder,
		bounds.Dx(), bounds.Dy(), decoder.ColorModel(),
		transform)

	defer transformer.Close()

	transformed, err := decodeImage(transformer)
	if err != nil {
		t.Errorf("Transformer: %s", err)
		decoder.Close()
		return
	}

	diff = imageDiff(expected, transformed)
	if diff != "" {
		t.Errorf("Transformer:\n%s", diff)
	}
}
