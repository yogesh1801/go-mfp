// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Image format detection test
package imgconv

import (
	"testing"

	"github.com/OpenPrinting/go-mfp/internal/testutils"
)

// TestMIMETypeDetect tests MIMETypeDetect function
func TestMIMETypeDetect(t *testing.T) {
	type testData struct {
		data   []byte // Data sample
		format string // Expected MIMETypeDetect output
	}

	tests := []testData{
		{
			data:   testutils.Images.BMP100x75,
			format: MIMETypeBMP,
		},

		{
			data:   testutils.Images.JPEG100x75rgb8,
			format: MIMETypeJPEG,
		},

		{
			data:   testutils.Images.PDF100x75,
			format: MIMETypePDF,
		},

		{
			data:   testutils.Images.PNG100x75rgb8,
			format: MIMETypePNG,
		},

		{
			data:   testutils.Images.PNG100x75gray8,
			format: MIMETypePNG,
		},

		{
			data:   testutils.Images.TIFF100x75,
			format: MIMETypeTIFF,
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
		format := MIMETypeDetect(test.data)
		if format != test.format {
			t.Errorf("Format expected %s, present %s",
				test.format, format)
		}
	}
}
