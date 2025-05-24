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
	"errors"
	"image"
	"image/color"
	"image/png"
	"io"
	"testing"

	"github.com/OpenPrinting/go-mfp/internal/testutils"
	"github.com/OpenPrinting/go-mfp/util/generic"
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
			//     #############   # - target
			//     #############
			//     **********
			data: testutils.Images.PNG100x75rgb8,
			rect: image.Rect(0, 25, 125, 50),
		},

		{
			//     **********      * - source
			//     **###########   # - target
			//     **###########
			//     **********
			data: testutils.Images.PNG100x75rgb8,
			rect: image.Rect(25, 25, 125, 50),
		},

		{
			//     **********      * - source
			//  #############      # - target
			//  #############
			//     **********
			data: testutils.Images.PNG100x75rgb8,
			rect: image.Rect(-25, 25, 100, 50),
		},

		{
			//     **********      * - source
			//  ###########**      # - target
			//  ###########**
			//     **********
			data: testutils.Images.PNG100x75rgb8,
			rect: image.Rect(-25, 25, 75, 50),
		},

		{
			//     **********      * - source
			//  ################   # - target
			//  ################
			//     **********
			data: testutils.Images.PNG100x75rgb8,
			rect: image.Rect(-25, 25, 125, 50),
		},

		{
			//       ######
			//     **######**      * - source
			//     **######**      # - target
			//     **######**
			//     **######**
			data: testutils.Images.PNG100x75rgb8,
			rect: image.Rect(25, -25, 75, 75),
		},

		{
			//       ######
			//     **######**      * - source
			//     **######**      # - target
			//     **######**
			//     **********
			data: testutils.Images.PNG100x75rgb8,
			rect: image.Rect(25, -25, 75, 50),
		},

		{
			//     **######**      * - source
			//     **######**      # - target
			//     **######**
			//     **######**
			//       ######
			data: testutils.Images.PNG100x75rgb8,
			rect: image.Rect(25, 0, 75, 100),
		},

		{
			//       ######
			//     **######**      * - source
			//     **######**      # - target
			//     **######**
			//     **######**
			//       ######
			data: testutils.Images.PNG100x75rgb8,
			rect: image.Rect(-25, 0, 75, 100),
		},

		{
			//  #######
			//  #######******      * - source
			//  #######******      # - target
			//     **********
			//     **********
			data: testutils.Images.PNG100x75rgb8,
			rect: image.Rect(-25, -25, 75, 50),
		},

		{
			//           #######
			//     ******#######   * - source
			//     ******#######   # - target
			//     **********
			//     **********
			data: testutils.Images.PNG100x75rgb8,
			rect: image.Rect(25, -25, 125, 50),
		},

		{
			//     **********      * - source
			//     **********      # - target
			//  #######******
			//  #######******
			//  #######
			data: testutils.Images.PNG100x75rgb8,
			rect: image.Rect(-25, 25, 75, 100),
		},

		{
			//     **********      * - source
			//     **********      # - target
			//     ******#######
			//     ******#######
			//           #######
			data: testutils.Images.PNG100x75rgb8,
			rect: image.Rect(25, 25, 125, 100),
		},

		{
			//  ################
			//  ###**********###   * - source
			//  ###**********###   # - target
			//  ###**********###
			//  ###**********###
			//  ################
			data: testutils.Images.PNG100x75rgb8,
			rect: image.Rect(-25, -25, 125, 100),
		},

		{
			//     **********      * - source
			//     ********** ##   # - target
			//     ********** ##
			//     ********** ##
			//                ##
			data: testutils.Images.PNG100x75rgb8,
			rect: image.Rect(125, 25, 150, 100),
		},

		{
			//     **********      * - source
			//     **********      # - target
			//     **********
			//     **********
			//
			//         #########
			//         #########
			data: testutils.Images.PNG100x75rgb8,
			rect: image.Rect(25, 100, 150, 125),
		},

		{
			//     **********      * - source
			//  ## **********      # - target
			//  ## **********
			//  ## **********
			//  ##
			data: testutils.Images.PNG100x75rgb8,
			rect: image.Rect(-50, 25, -25, 100),
		},

		{
			//  #########
			//  #########
			//
			//     **********      * - source
			//     **********      # - target
			//     **********
			//     **********
			data: testutils.Images.PNG100x75rgb8,
			rect: image.Rect(-50, -25, -25, 75),
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

		// Resizer must return io.EOF after end
		_, err = resizer.Read(row)
		if err != io.EOF {
			t.Errorf("%s: read after end:\n"+
				"error expected: %s\n"+
				"error present:  %v\n",
				test.rect, io.EOF, err)
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

// TestResizerErrors tests how resizer handles I/O errors
func TestResizerErrors(t *testing.T) {
	type testData struct {
		model    color.Model     // Color model
		wid, hei int             // Input (before resizing) image size
		lim      int             // Limit of non-error reads
		rect     image.Rectangle // Target size
	}

	tests := []testData{
		{
			//     **********      *  - source
			//     **######**      #  - target
			//  >> **######**      >> - lim
			//     **********
			model: color.RGBAModel,
			wid:   100,
			hei:   75,
			lim:   37,
			rect:  image.Rect(25, 25, 75, 50),
		},

		{
			//     ################
			//     ###**********###   *  - source
			//     ###**********###   #  - target
			//  >> ###**********###   >> - lim
			//     ###**********###
			//     ################
			model: color.RGBAModel,
			wid:   100,
			hei:   75,
			lim:   50,
			rect:  image.Rect(-25, -25, 125, 100),
		},

		{
			//  >> **********      *  - source
			//     **######**      #  - target
			//     **######**      >> - lim
			//     **********
			model: color.RGBAModel,
			wid:   100,
			hei:   75,
			lim:   10,
			rect:  image.Rect(25, 25, 75, 50),
		},
	}

	for _, test := range tests {
		// Create image conversion pipeline
		ioerr := errors.New("I/O error")

		decoder := newDecoderWithError(test.model,
			test.wid, test.hei, test.lim, ioerr)
		resizer := NewResizer(decoder, test.rect)

		// Rows until limResized expected to be OK
		limResized := test.lim - test.rect.Min.Y
		limResized = generic.Max(0, limResized)
		limResized = generic.Min(test.rect.Dy(), limResized)

		row := resizer.NewRow()
		ok := true
		for y := 0; ok && y < limResized; y++ {
			_, err := resizer.Read(row)
			if err != nil {
				t.Errorf("%s: unexpected error %s",
					test.rect, err)
				ok = false
			}
		}

		// All subsequent reads must return ioerr
		for y := limResized; ok && y < test.rect.Dy(); y++ {
			_, err := resizer.Read(row)
			if err != ioerr {
				t.Errorf("%s: line %d:\n"+
					"error expected: %s\n"+
					"error present:  %v\n",
					test.rect, y, ioerr, err)
				ok = false
			}
		}

		resizer.Close()
		decoder.Close()
	}
}
