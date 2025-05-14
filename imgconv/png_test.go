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

	for _, test := range tests {
		in := bytes.NewReader(test.data)
		decoder, err := NewPNGDecoder(in)
		if err != nil {
			t.Errorf("%s: NewPNGDecoder: %s",
				test.name, err)
			continue
		}

		// Decode reference reference
		reference, err := png.Decode(bytes.NewReader(test.data))
		if err != nil {
			panic(err)
		}

		// Compare image size
		if decoder.Bounds() != reference.Bounds() {
			t.Errorf("%s: Bounds mismatch:\n"+
				"expected: %v\n"+
				"present:  %v\n",
				test.name,
				reference.Bounds(),
				decoder.Bounds(),
			)
			decoder.Close()
			continue
		}

		diff := imageDiff(reference, decoder)
		if diff != "" {
			t.Errorf("%s: %s", test.name, diff)
			decoder.Close()
			continue
		}

		decoder.Close()
	}
}
