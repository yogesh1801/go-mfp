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

func TestPNGDecode(t *testing.T) {
	type testData struct {
		name string // Image name, for logging
		data []byte // Image data
	}

	tests := []testData{
		{"PNG100x75rgb8", testutils.Images.PNG100x75rgb8},
		{"PNG100x75gray8", testutils.Images.PNG100x75gray8},
		{"PNG100x75gray16", testutils.Images.PNG100x75gray16},
		{"PNG100x75rgb16", testutils.Images.PNG100x75rgb16},
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
}
