// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Document format detection

package abstract

import (
	"testing"

	"github.com/OpenPrinting/go-mfp/internal/testutils"
)

// TestDocumentFormatDetect tests DocumentFormatDetect function
func TestDocumentFormatDetect(t *testing.T) {
	type testData struct {
		data   []byte // Data sample
		format string // Expected DocumentFormatDetect output
	}

	tests := []testData{
		{
			data:   testutils.Images.BMP100x75,
			format: DocumentFormatBMP,
		},

		{
			data:   testutils.Images.JPEG100x75,
			format: DocumentFormatJPEG,
		},

		{
			data:   testutils.Images.PDF100x75,
			format: DocumentFormatPDF,
		},

		{
			data:   testutils.Images.PNG100x75,
			format: DocumentFormatPNG,
		},

		{
			data:   testutils.Images.TIFF100x75,
			format: DocumentFormatTIFF,
		},

		{
			data:   []byte{},
			format: "",
		},

		{
			data:   make([]byte, 256),
			format: "",
		},
	}

	for _, test := range tests {
		format := DocumentFormatDetect(test.data)
		if format != test.format {
			t.Errorf("Format expected %s, present %s",
				test.format, format)
		}
	}
}
