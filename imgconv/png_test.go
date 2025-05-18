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
	"errors"
	"image"
	"image/color"
	"image/png"
	"io"
	"strings"
	"testing"

	"github.com/OpenPrinting/go-mfp/internal/testutils"
	"github.com/OpenPrinting/go-mfp/util/generic"
)

// TestPNGDecode tests PNG decoder
func TestPNGDecode(t *testing.T) {
	type testData struct {
		name  string      // Image name, for logging
		data  []byte      // Image data
		model color.Model // Expected color model
		img   image.Image // Decoded reference images
	}

	tests := []testData{
		{
			name:  "PNG100x75rgb8",
			data:  testutils.Images.PNG100x75rgb8,
			model: color.RGBAModel,
		},
		{
			name:  "PNG100x75rgb8paletted",
			data:  testutils.Images.PNG100x75rgb8paletted,
			model: color.RGBAModel,
		},
		{
			name:  "PNG100x75gray1",
			data:  testutils.Images.PNG100x75gray1,
			model: color.GrayModel,
		},
		{
			name:  "PNG100x75gray8",
			data:  testutils.Images.PNG100x75gray8,
			model: color.GrayModel,
		},
		{
			name:  "PNG100x75gray16",
			data:  testutils.Images.PNG100x75gray16,
			model: color.Gray16Model,
		},
		{
			name:  "PNG100x75rgb16",
			data:  testutils.Images.PNG100x75rgb16,
			model: color.RGBA64Model,
		},
	}

	// Decode reference images. We need to do it only once.
	for i := range tests {
		test := &tests[i]
		reference, err := png.Decode(bytes.NewReader(test.data))
		if err != nil {
			panic(err)
		}
		test.img = reference
	}

	// Test image decoding
	for _, test := range tests {
		// Create a decoder
		in := bytes.NewReader(test.data)
		decoder, err := NewPNGDecoder(in)
		if err != nil {
			t.Errorf("%s: NewPNGDecoder: %s",
				test.name, err)
			continue
		}

		// Decode test image
		img, err := decodeImage(decoder)
		if err != nil {
			t.Errorf("%s: decodeImage: %s",
				test.name, err)
			continue
		}

		// Check ColorModel()
		model := decoder.ColorModel()
		if model != test.model {
			t.Errorf("%s: Decoder.ColorModel mismatch",
				test.name)
		}

		// Compare images
		diff := imageDiff(test.img, img)
		if diff != "" {
			t.Errorf("%s: %s", test.name, diff)
			decoder.Close()
			continue
		}

		decoder.Close()
	}

	// Test handling of truncated data
	for _, test := range tests {
		step := 1
		for trunc := 0; trunc < len(test.data)-1; trunc += step {
			// Create a decoder
			in := bytes.NewReader(test.data[:trunc])
			decoder, err := NewPNGDecoder(in)
			if err != nil {
				continue // Error is expected
			}

			// Once header is decoded, increment step.
			// Otherwise test longs for too long.
			step = 16

			// Decode test image
			img, err := decodeImage(decoder)
			if err != nil {
				decoder.Close()
				continue // Error is expected
			}

			// Compare images.
			//
			// If everything is OK so far, images must match
			diff := imageDiff(test.img, img)
			if diff != "" {
				t.Errorf("%s: %s", test.name, diff)
				decoder.Close()
				continue
			}

			decoder.Close()
		}
	}

	// Test handling of damaged data
	for _, test := range tests {
		damagedHeader := false
		damagedData := false
		for off := 0; off < len(test.data); off++ {
			// Scan until we have seen both damaged header
			// and damaged image data
			if damagedHeader && damagedData {
				break
			}

			// Damage data
			data := generic.CopySlice(test.data)
			data[off] = ^data[off]

			// Create a decoder
			in := bytes.NewReader(data)
			decoder, err := NewPNGDecoder(in)
			if err != nil {
				damagedHeader = true
				continue // Error is expected
			}

			// Decode test image
			img, err := decodeImage(decoder)
			if err != nil {
				damagedData = true
				continue // Error is expected
			}

			// Compare images.
			//
			// If everything is OK so far, images must match
			diff := imageDiff(test.img, img)
			if diff != "" {
				t.Errorf("%s: %s", test.name, diff)
				decoder.Close()
				continue
			}

			decoder.Close()
		}

		if !(damagedHeader && damagedData) {
			t.Errorf("%s: damaged data not handled properly",
				test.name)
		}
	}

	// Test handling of I/O errors
	for _, test := range tests {
		expectedErr := errors.New("I/O error, for testing")
		damagedHeader := false
		damagedData := false
		fail := false
		for off := 0; off < len(test.data) && fail; off++ {
			// Scan until we have seen both damaged header
			// and damaged image data
			if damagedHeader && damagedData {
				break
			}

			// Simulate I/O error
			rd := newReaderWithError(test.data[:off], expectedErr)
			data := generic.CopySlice(test.data)
			data[off] = ^data[off]

			// Create a decoder
			decoder, err := NewPNGDecoder(rd)
			if err != nil {
				damagedHeader = true
				if err != expectedErr {
					fail = true
					t.Errorf("%s:in NewPNGDecoder\n"+
						"error expected: %s\n"+
						"error present:  %s\n",
						test.name, expectedErr, err)
				}
				continue // Error is expected
			}

			img, err := decodeImage(decoder)
			if err != nil {
				damagedData = true
				if err != expectedErr {
					fail = true
					t.Errorf("%s:in decodeImage\n"+
						"error expected: %s\n"+
						"error present:  %s\n",
						test.name, expectedErr, err)
				}
				continue // Error is expected
			}

			// Compare images.
			//
			// If everything is OK so far, images must match
			diff := imageDiff(test.img, img)
			if diff != "" {
				t.Errorf("%s: %s", test.name, diff)
				decoder.Close()
				continue
			}

			decoder.Close()
		}
	}
}

// TestPNGDecodeErrors tests PNG decoder errors, not handled by other tests
func TestPNGDecodeErrors(t *testing.T) {
	// Interlaced images not supported
	in := bytes.NewReader(testutils.Images.PNG100x75rgb8i)
	_, err := NewPNGDecoder(in)

	if err == nil || !strings.Contains(err.Error(), "interlaced") {
		s := "nil"
		if err != nil {
			s = err.Error()
		}
		t.Errorf("PNG: test for interlaced images failed (err=%s)", s)
	}
}

// TestPNGEncode tests PNG encoder
func TestPNGEncode(t *testing.T) {
	type testData struct {
		name string      // Image name, for logging
		data []byte      // Image data
		img  image.Image // Decoded reference images
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

	// Decode reference images. We need to do it only once.
	for i := range tests {
		test := &tests[i]
		reference, err := png.Decode(bytes.NewReader(test.data))
		if err != nil {
			panic(err)
		}
		test.img = reference
	}

	buf := &bytes.Buffer{}
	for _, test := range tests {
		// Create image decoder
		in := bytes.NewReader(test.data)
		decoder, err := NewPNGDecoder(in)
		if err != nil {
			panic(err)
		}

		// Create image encoder
		buf.Reset()
		wid, hei := decoder.Size()
		model := decoder.ColorModel()

		encoder, err := NewPNGEncoder(buf, wid, hei, model)
		if err != nil {
			t.Errorf("%s: NewPNGEncoder: %s", test.name, err)
			decoder.Close()
			continue
		}

		// Recode the image, row by row
		for err == nil {
			var row Row
			row, err = decoder.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				panic(err)
			}

			err = encoder.Write(row)
			if err != nil {
				t.Errorf("%s: Encoder.Write: %s", test.name, err)
				decoder.Close()
				encoder.Close()
				continue
			}
		}

		decoder.Close()
		encoder.Close()

		// Decode just encoded image by reference decoder
		img, err := png.Decode(buf)
		if err != nil {
			t.Errorf("%s: error in encoded image: %s",
				test.name, err)
			continue
		}

		diff := imageDiff(test.img, img)
		if diff != "" {
			t.Errorf("%s: %s", test.name, diff)
		}
	}
}
