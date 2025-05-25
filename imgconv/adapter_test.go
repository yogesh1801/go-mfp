// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Decoder/Encoder image adapter test

package imgconv

import (
	"bytes"
	"errors"
	"image"
	"image/color"
	"image/png"
	"testing"

	"github.com/OpenPrinting/go-mfp/internal/testutils"
	"golang.org/x/image/draw"
)

// TestDecoderImageAdapter tests DecoderImageAdapter
func TestDecoderImageAdapter(t *testing.T) {
	type testData struct {
		name string // test name
		data []byte // Image data (PNG)
	}

	tests := []testData{
		{
			name: "PNG100x75rgb8",
			data: testutils.Images.PNG100x75rgb8,
		},
		{
			name: "PNG100x75rgb8paletted",
			data: testutils.Images.PNG100x75rgb8paletted,
		},
		{
			name: "PNG100x75gray1",
			data: testutils.Images.PNG100x75gray1,
		},
		{
			name: "PNG100x75gray8",
			data: testutils.Images.PNG100x75gray8,
		},
		{
			name: "PNG100x75gray16",
			data: testutils.Images.PNG100x75gray16,
		},
		{
			name: "PNG100x75rgb16",
			data: testutils.Images.PNG100x75rgb16,
		},
	}

	for _, test := range tests {
		// Create DecoderImageAdapter on a top of PNG decoder
		decoder, err := NewPNGDecoder(bytes.NewReader(test.data))
		if err != nil {
			panic(err)
		}
		adapter := NewDecoderImageAdapter(decoder)

		// Decode image with stdlib decoder, for reference
		reference, err := png.Decode(bytes.NewReader(test.data))
		if err != nil {
			panic(err)
		}

		// Compare color models.
		//
		// Note, we check color mode against the decoder,
		// not the reference image. Stdlib and PNG decoders
		// may have slightly different ideas on the image
		// color modes (stdlib may use NRGBA/NRGBA64, while
		// we alsways prefer RGBA/GRBA64).
		if decoder.ColorModel() != adapter.ColorModel() {
			t.Errorf("%s: color model mismatch:\n"+
				"expected: %v\n"+
				"present:  %v\n",
				test.name,
				decoder.ColorModel(), adapter.ColorModel())
		}

		// Compare two images
		diff := imageDiff(reference, adapter)
		if diff != "" {
			t.Errorf("%s: %s", test.name, diff)
		}

		// Cleanup
		adapter.Close()
	}
}

// TestDecoderImageAdapterErrors tests I/O errors handling
// by DecoderImageAdapter
func TestDecoderImageAdapterErrors(t *testing.T) {
	wid := 100
	hei := 75
	lim := hei / 2
	ioerr := errors.New("I/O error")

	source := newDecoderWithError(color.RGBA64Model, wid, hei, lim, ioerr)
	adapter := NewDecoderImageAdapter(source)
	reference := image.NewRGBA(image.Rect(0, 0, wid, hei))

	draw.Draw(reference, image.Rect(0, 0, wid, lim),
		&image.Uniform{color.White}, image.ZP, draw.Over)

	defer adapter.Close()

	diff := imageDiff(reference, adapter)

	if adapter.Error() != ioerr {
		t.Errorf("error mismatch:\n"+
			"expected: %v\n"+
			"present:  %v\n",
			ioerr, adapter.Error())
	}

	if diff != "" {
		t.Errorf("decoded image mismatch:\n%s", diff)
	}
}

// TestEncoderImageAdapter tests EncoderImageAdapter
func TestEncoderImageAdapter(t *testing.T) {
	type testData struct {
		name string // Test name
		data []byte // Image data (PNG)
	}

	tests := []testData{
		{
			name: "PNG100x75rgb8",
			data: testutils.Images.PNG100x75rgb8,
		},
		{
			name: "PNG100x75rgb8paletted",
			data: testutils.Images.PNG100x75rgb8paletted,
		},
		{
			name: "PNG100x75gray1",
			data: testutils.Images.PNG100x75gray1,
		},
		{
			name: "PNG100x75gray8",
			data: testutils.Images.PNG100x75gray8,
		},
		{
			name: "PNG100x75gray16",
			data: testutils.Images.PNG100x75gray16,
		},
		{
			name: "PNG100x75rgb16",
			data: testutils.Images.PNG100x75rgb16,
		},
	}

	for _, test := range tests {
		// Create DecoderImageAdapter on a top of PNG decoder
		decoder, err := NewPNGDecoder(bytes.NewReader(test.data))
		if err != nil {
			panic(err)
		}

		decoderAdapter := NewDecoderImageAdapter(decoder)

		// Create EncoderImageAdapter on a top of PNG encoder
		wid, hei := decoder.Size()
		model := decoder.ColorModel()
		buf := &bytes.Buffer{}
		encoder, err := NewPNGEncoder(buf, wid, hei, model)
		if err != nil {
			panic(err)
		}

		encoderAdapter := NewEncoderImageAdapter(encoder)

		// Test encoderAdapter parameters
		if encoderAdapter.ColorModel() != model {
			t.Errorf("%s: EncoderImageAdapter.ColorModel:\n"+
				"expected: %v\n"+
				"present:  %v\n",
				test.name, model, encoderAdapter.ColorModel())
		}

		bounds := image.Rect(0, 0, wid, hei)
		if encoderAdapter.Bounds() != bounds {
			t.Errorf("%s: EncoderImageAdapter.ColorModel:\n"+
				"expected: %s\n"+
				"present:  %s\n",
				test.name, bounds, encoderAdapter.Bounds())
		}

		// Copy image from decoder to encoder
		draw.Draw(encoderAdapter, image.Rect(0, 0, wid, hei),
			decoderAdapter, image.ZP, draw.Over)

		encoderAdapter.Close() // Flushes last line
		decoderAdapter.Close()

		// Decode image with stdlib decoder, for reference
		reference, err := png.Decode(bytes.NewReader(test.data))
		if err != nil {
			panic(err)
		}

		// Decode resulting image, for comparison
		recoded, err := png.Decode(buf)
		if err != nil {
			panic(err)
		}

		diff := imageDiff(reference, recoded)
		if diff != "" {
			t.Errorf("%s: Encoded image mismatch:\n%s",
				test.name, diff)
		}
	}
}

// TestEncoderImageAdapterErrors tests I/O errors handling by
// the EncoderImageAdapter
func TestEncoderImageAdapterErrors(t *testing.T) {
	wid := 100
	hei := 75
	lim := hei / 2
	model := color.RGBAModel
	ioerr := errors.New("I/O error")

	encoder := newEncoderWithError(model, wid, hei, lim, ioerr)
	encoderAdapter := NewEncoderImageAdapter(encoder)

	draw.Draw(encoderAdapter, image.Rect(0, 0, wid, hei),
		&image.Uniform{color.White}, image.ZP, draw.Over)

	if encoderAdapter.Error() != ioerr {
		t.Errorf("EncoderImageAdapter.Error:\n"+
			"expected: %s\n"+
			"present:  %s\n",
			ioerr, encoderAdapter.Error())
	}
}
