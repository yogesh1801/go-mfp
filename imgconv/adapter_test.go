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
	"io"
	"testing"

	"github.com/OpenPrinting/go-mfp/internal/testutils"
	"golang.org/x/image/draw"
)

// TestSourceImageAdapter tests SourceImageAdapter
func TestSourceImageAdapter(t *testing.T) {
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
		// Create SourceImageAdapter on a top of PNG decoder
		decoder, err := NewPNGDecoder(bytes.NewReader(test.data))
		if err != nil {
			panic(err)
		}
		source := NewSourceImageAdapter(decoder)

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
		if decoder.ColorModel() != source.ColorModel() {
			t.Errorf("%s: color model mismatch:\n"+
				"expected: %v\n"+
				"present:  %v\n",
				test.name,
				decoder.ColorModel(), source.ColorModel())
		}

		// Compare two images
		diff := imageDiff(reference, source)
		if diff != "" {
			t.Errorf("%s: %s", test.name, diff)
		}

		// Cleanup
		source.Close()
	}
}

// TestSourceImageAdapterErrors tests I/O errors handling
// by SourceImageAdapter
func TestSourceImageAdapterErrors(t *testing.T) {
	wid := 100
	hei := 75
	lim := hei / 2
	ioerr := errors.New("I/O error")

	input := newDecoderWithError(color.RGBA64Model, wid, hei, lim, ioerr)
	source := NewSourceImageAdapter(input)
	reference := image.NewRGBA(image.Rect(0, 0, wid, hei))

	draw.Draw(reference, image.Rect(0, 0, wid, lim),
		&image.Uniform{color.White}, image.ZP, draw.Over)

	defer source.Close()

	diff := imageDiff(reference, source)

	if source.Error() != ioerr {
		t.Errorf("error mismatch:\n"+
			"expected: %v\n"+
			"present:  %v\n",
			ioerr, source.Error())
	}

	if diff != "" {
		t.Errorf("decoded image mismatch:\n%s", diff)
	}
}

// TestTargetImageAdapter tests TargetImageAdapter
func TestTargetImageAdapter(t *testing.T) {
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
		// Create SourceImageAdapter on a top of PNG decoder
		decoder, err := NewPNGDecoder(bytes.NewReader(test.data))
		if err != nil {
			panic(err)
		}

		source := NewSourceImageAdapter(decoder)

		// Create TargetImageAdapter
		wid, hei := decoder.Size()
		model := decoder.ColorModel()
		target := NewTargetImageAdapter(wid, hei, model)

		// Test TargetImageAdapter parameters
		if target.ColorModel() != model {
			t.Errorf("%s: TargetImageAdapter.ColorModel:\n"+
				"expected: %v\n"+
				"present:  %v\n",
				test.name, model, target.ColorModel())
		}

		bounds := image.Rect(0, 0, wid, hei)
		if target.Bounds() != bounds {
			t.Errorf("%s: TargetImageAdapter.ColorModel:\n"+
				"expected: %s\n"+
				"present:  %s\n",
				test.name, bounds, target.Bounds())
		}

		// Copy image from decoder to target
		go func() {
			draw.Draw(target, image.Rect(0, 0, wid, hei),
				source, image.ZP, draw.Over)

			target.Flush()
			source.Close()
		}()

		// Read image from Decoder side of target
		recoded, err := decodeImage(target)
		if err != nil {
			t.Errorf("%s: TargetImageAdapter.Read: %s",
				test.name, err)
		}

		// Decode image with stdlib decoder, for reference
		reference, err := png.Decode(bytes.NewReader(test.data))
		if err != nil {
			panic(err)
		}

		// Compare images
		diff := imageDiff(reference, recoded)
		if diff != "" {
			t.Errorf("%s: Encoded image mismatch:\n%s",
				test.name, diff)
		}

		// Subsequent reads from the target must return io.EOF
		row := target.NewRow()
		_, err = target.Read(row)
		if err != io.EOF {
			t.Errorf("%s: TargetImageAdapter.Read:\n"+
				"error expected: %v\n"+
				"error present:  %v\n",
				test.name, io.EOF, err)
		}

		// Cleanup after test
		target.Close()
	}
}

// TestTargetImageAdapterClose tests TargetImageAdapter.Close
func TestTargetImageAdapterClose(t *testing.T) {
	wid := 100
	hei := 75
	model := color.RGBAModel

	target := NewTargetImageAdapter(wid, hei, model)
	target.Close()

	// Draw into the closed target must not block
	draw.Draw(target, image.Rect(0, 0, wid, hei),
		image.White, image.ZP, draw.Over)

	target.Flush()
}
