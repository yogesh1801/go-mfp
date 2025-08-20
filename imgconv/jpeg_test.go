// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// JPEG tests

package imgconv

import (
	"bytes"
	"errors"
	"image"
	"image/color"
	"image/jpeg"
	"testing"

	"github.com/OpenPrinting/go-mfp/internal/testutils"
	"github.com/OpenPrinting/go-mfp/util/generic"
)

// TestJPEGDecode tests JPEG reader
func TestJPEGDecode(t *testing.T) {
	type testData struct {
		name  string      // Image name, for logging
		data  []byte      // Image data
		model color.Model // Expected color model
		img   image.Image // Decoded reference image
	}

	tests := []testData{
		{
			name:  "JPEG100x75rgb8",
			data:  testutils.Images.JPEG100x75rgb8,
			model: color.RGBAModel,
		},

		{
			name:  "JPEG100x75gray8",
			data:  testutils.Images.JPEG100x75gray8,
			model: color.GrayModel,
		},
	}

	// Decode reference images. We need to do it only once.
	for i := range tests {
		test := &tests[i]
		reference, err := jpeg.Decode(bytes.NewReader(test.data))
		if err != nil {
			panic(err)
		}
		test.img = reference
	}

	// Test image decoding
	for _, test := range tests {
		// Create a reader
		in := bytes.NewReader(test.data)
		reader, err := NewJPEGReader(in)
		if err != nil {
			t.Errorf("%s: NewJPEGReader: %s",
				test.name, err)
			continue
		}

		// Decode test image
		img, err := decodeImage(reader)
		if err != nil {
			t.Errorf("%s: decodeImage: %s",
				test.name, err)
			reader.Close()
			continue
		}

		// Check ColorModel()
		model := reader.ColorModel()
		if model != test.model {
			t.Errorf("%s: Reader.ColorModel mismatch",
				test.name)
			reader.Close()
			continue
		}

		// Compare images
		//
		// Note, as JPEG internally uses the YCbCr color model,
		// while we operate in the RBG terms, some small conversion
		// error is OK. Se we must use imageEuclideanDistance
		// metric instead of the exact match.
		dist := imageEuclideanDistance(test.img, img)
		if dist > 1.0/100 {
			t.Errorf("%s: images too different", test.name)
			reader.Close()
			continue
		}

		reader.Close()
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

			// Create a reader
			in := bytes.NewReader(data)
			reader, err := NewJPEGReader(in)
			if err != nil {
				damagedHeader = true
				continue // Error is expected
			}

			// Decode test image
			img, err := decodeImage(reader)
			if err != nil {
				damagedData = true
				continue // Error is expected
			}

			// JPEG doesn't have any integrity check,
			// so comparing images is meaningless
			_ = img

			reader.Close()
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
			rd := newIoReaderWithError(test.data[:off], expectedErr)
			data := generic.CopySlice(test.data)
			data[off] = ^data[off]

			// Create a reader
			reader, err := NewJPEGReader(rd)
			if err != nil {
				damagedHeader = true
				if err != expectedErr {
					fail = true
					t.Errorf("%s:in NewJPEGReader\n"+
						"error expected: %s\n"+
						"error present:  %s\n",
						test.name, expectedErr, err)
				}
				continue // Error is expected
			}

			img, err := decodeImage(reader)
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

			// JPEG doesn't have any integrity check,
			// so comparing images is meaningless
			_ = img

			reader.Close()
		}
	}
}
