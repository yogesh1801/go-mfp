// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// Test for format value

package wsscan

import "testing"

var testFormatValue = testEnum[FormatValue]{
	decodeStr: DecodeFormatValue,
	decodeXML: decodeFormatValue,
	dataset: []testEnumData[FormatValue]{
		{DIB, "dib"},
		{EXIF, "exif"},
		{JBIG, "jbig"},
		{JFIF, "jfif"},
		{JPEG2K, "jpeg2k"},
		{PDFA, "pdf-a"},
		{PNG, "png"},
		{TIFFSingleUncompressed, "tiff-single-uncompressed"},
		{TIFFSingleG4, "tiff-single-g4"},
		{TIFFSingleG3MH, "tiff-single-g3mh"},
		{TIFFSingleJPEGTN2, "tiff-single-jpeg-tn2"},
		{TIFFMultiUncompressed, "tiff-multi-uncompressed"},
		{TIFFMultiG4, "tiff-multi-g4"},
		{TIFFMultiG3MH, "tiff-multi-g3mh"},
		{TIFFMultiJPEGTN2, "tiff-multi-jpeg-tn2"},
		{XPS, "xps"},
	},
}

// TestFormatValue tests [FormatValue] common methods and functions.
func TestFormatValue(t *testing.T) {
	testFormatValue.run(t)
}
