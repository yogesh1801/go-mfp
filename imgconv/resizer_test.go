// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Image resizer test

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

func TestResizer(t *testing.T) {
	type testData struct {
		data []byte
		rect image.Rectangle
	}

	tests := []testData{
		{
			data: testutils.Images.PNG100x75rgb8,
			rect: image.Rect(0, 0, 100, 75),
		},

		{
			//     **********      * - source
			//     ##########      # - target
			//     ##########
			//     **********
			data: testutils.Images.PNG100x75rgb8,
			rect: image.Rect(0, 25, 100, 50),
		},

		{
			//     **######**      * - source
			//     **######**      # - target
			//     **######**
			//     **######**
			data: testutils.Images.PNG100x75rgb8,
			rect: image.Rect(25, 0, 75, 75),
		},

		{
			//     **********      * - source
			//     **######**      # - target
			//     **######**
			//     **********
			data: testutils.Images.PNG100x75rgb8,
			rect: image.Rect(25, 25, 75, 50),
		},

		{
			//     **********      * - source
			//     ##############  # - target
			//     ##############
			//     **********
			data: testutils.Images.PNG100x75rgb8,
			rect: image.Rect(0, 25, 125, 50),
		},

		{
			//     **********      * - source
			//     **############  # - target
			//     **############
			//     **********
			data: testutils.Images.PNG100x75rgb8,
			rect: image.Rect(25, 25, 125, 50),
		},
	}

	for _, test := range tests {
		// Create image conversion pipeline
		decoder, err := NewPNGDecoder(bytes.NewReader(test.data))
		if err != nil {
			panic(err)
		}

		resizer := NewResizer(decoder, test.rect)
		buf := &bytes.Buffer{}

		encoder, err := NewPNGEncoder(buf,
			test.rect.Dx(), test.rect.Dy(), decoder.ColorModel())
		if err != nil {
			panic(err)
		}

		// Convert, row by row
		row := resizer.NewRow()
		row.Fill(color.RGBA{R: 0xff})

		for y := 0; y < test.rect.Dy(); y++ {
			_, err = resizer.Read(row)
			if err != nil {
				panic(err)
			}

			err = encoder.Write(row)
			if err != nil {
				panic(err)
			}
		}

		resizer.Close()
		encoder.Close()

		// Now compare converted image with expected
		source, err := png.Decode(bytes.NewReader(test.data))
		if err != nil {
			panic(err)
		}

		resized, err := png.Decode(buf)
		if err != nil {
			panic(err)
		}

		var expected draw.Image
		bounds := image.Rect(0, 0, test.rect.Dx(), test.rect.Dy())
		switch resizer.ColorModel() {
		case color.GrayModel:
			expected = image.NewGray(bounds)
		case color.Gray16Model:
			expected = image.NewGray16(bounds)
		case color.RGBAModel:
			expected = image.NewRGBA(bounds)
		case color.RGBA64Model:
			expected = image.NewRGBA64(bounds)
		}

		draw.Draw(expected, bounds, &image.Uniform{color.White},
			image.ZP, draw.Over)
		draw.Draw(expected, bounds, source, test.rect.Min, draw.Over)

		diff := imageDiff(expected, resized)
		if diff != "" {
			t.Errorf("%s: image mismatch:\n%s", test.rect, diff)
		}
	}
}
