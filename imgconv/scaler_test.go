// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Image scaler test

package imgconv

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"testing"

	"github.com/OpenPrinting/go-mfp/internal/testutils"
	"golang.org/x/image/draw"
)

// TestScaler tests image scaler
func TestScaler(t *testing.T) {
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
		//{
		//	name: "PNG5100x7016",
		//	data: testutils.Images.PNG5100x7016,
		//},
	}

	for _, test := range tests {
		// Decode reference image
		reference, err := png.Decode(bytes.NewReader(test.data))
		if err != nil {
			panic(err)
		}

		bounds := reference.Bounds()
		wid, hei := bounds.Dx(), bounds.Dy()

		// Run over some scales
		for _, scale := range []float32{0.5, 1.0, 2.0} {
			scaledWid := int(float32(wid) * scale)
			scaledHei := int(float32(hei) * scale)
			name := fmt.Sprintf("%s-%dx%d", test.name,
				scaledWid, scaledHei)

			// Scale image with scaler
			decoder, err := NewPNGDecoder(
				bytes.NewReader(test.data))
			if err != nil {
				panic(err)
			}
			scaler := NewScaler(decoder, scaledWid, scaledHei)
			scaled, err := decodeImage(scaler)
			scaler.Close()

			if err != nil {
				t.Errorf("%s: decode error: %s", name, err)
				continue
			}

			//saveImage(name+".png", scaled)

			// Scale image with library function
			scaledBounds := image.Rect(0, 0, scaledWid, scaledHei)
			var expected draw.Image
			switch decoder.ColorModel() {
			case color.GrayModel:
				expected = image.NewGray(scaledBounds)
			case color.Gray16Model:
				expected = image.NewGray16(scaledBounds)
			case color.RGBAModel:
				expected = image.NewRGBA(scaledBounds)
			case color.RGBA64Model:
				expected = image.NewRGBA64(scaledBounds)
			default:
				panic("internal error")
			}

			draw.BiLinear.Scale(
				expected, expected.Bounds(),
				reference, reference.Bounds(),
				draw.Over, nil)

			//saveImage(name+"-ref.png", expected)

			//diff := imageDiff(expected, scaled)
			//println(diff)

			dist := imageEuclideanDistance(expected, scaled)
			if dist > 1.0/100 {
				t.Errorf("%s: images too different: %g%%",
					name, dist*100)
			}
		}
	}
}
