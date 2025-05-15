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
	"image/color"
	"image/png"
	"testing"

	"github.com/OpenPrinting/go-mfp/internal/testutils"
	"github.com/OpenPrinting/go-mfp/util/generic"
)

func TestPNGDecode(t *testing.T) {
	type testData struct {
		name  string      // Image name, for logging
		data  []byte      // Image data
		model color.Model // Expected color model
	}

	tests := []testData{
		{"PNG100x75rgb8", testutils.Images.PNG100x75rgb8,
			color.RGBAModel},
		{"PNG100x75gray1", testutils.Images.PNG100x75gray1,
			color.GrayModel},
		{"PNG100x75gray8", testutils.Images.PNG100x75gray8,
			color.GrayModel},
		{"PNG100x75gray16", testutils.Images.PNG100x75gray16,
			color.Gray16Model},
		{"PNG100x75rgb16", testutils.Images.PNG100x75rgb16,
			color.RGBA64Model},
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

		// Check ColorModel()
		model := decoder.ColorModel()
		if model != test.model {
			t.Errorf("%s: Decoder.ColorModel mismatch",
				test.name)
		}

		// Decode reference image
		reference, err := png.Decode(bytes.NewReader(test.data))
		if err != nil {
			panic(err)
		}

		// Compare images
		diff := imageDiff(reference, decoder)
		if diff != "" {
			t.Errorf("%s: %s", test.name, diff)
			decoder.Close()
			continue
		}

		decoder.Close()
	}

	// Test handling of truncated data
	for _, test := range tests {
		// Decode reference image
		reference, err := png.Decode(bytes.NewReader(test.data))
		if err != nil {
			panic(err)
		}

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

			// Compare images.
			//
			// If everything is OK, either images must
			// match or decoder must be in error.
			diff := imageDiff(reference, decoder)
			if diff != "" && decoder.Error() == nil {
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

			height := decoder.Bounds().Dy()
			for y := 0; y < height && decoder.Error() == nil; y++ {
				decoder.At(0, y)
			}

			if decoder.Error() != nil {
				damagedData = true
			}

			decoder.Close()
		}

		if !(damagedHeader && damagedData) {
			t.Errorf("%s: damaged data not handled properly",
				test.name)
		}
	}

	// Test handling of I/O errors
	expectedErr := errors.New("I/O error, for testing")
	damagedHeader := false
	damagedData := false
	for _, test := range tests {
		for off := 0; off < len(test.data); off++ {
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
					t.Errorf("%s:\n"+
						"error expected: %s\n"+
						"error present:  %s\n",
						test.name, expectedErr, err)
				}
				continue // Error is expected
			}

			height := decoder.Bounds().Dy()
			for y := 0; y < height && decoder.Error() == nil; y++ {
				decoder.At(0, y)
			}

			if err := decoder.Error(); err != nil {
				damagedData = true
				if err != expectedErr {
					t.Errorf("%s:\n"+
						"error expected: %s\n"+
						"error present:  %s\n",
						test.name, expectedErr, err)
				}
			}

			decoder.Close()
		}
	}
}
