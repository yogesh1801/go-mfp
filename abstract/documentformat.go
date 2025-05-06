// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Document format detection

package abstract

import "bytes"

// MIME types for known document formats:
const (
	DocumentFormatBMP  = "image/bmp"
	DocumentFormatJPEG = "image/jpeg"
	DocumentFormatPDF  = "application/pdf"
	DocumentFormatPNG  = "image/png"
	DocumentFormatTIFF = "image/tiff"
)

// DocumentFormatDetect detects document type by its few starting bytes
// and returns its MIME type.
//
// If format cannot be guessed, "" is returned.
func DocumentFormatDetect(image []byte) string {
	for _, ent := range formatTable {
		if bytes.HasPrefix(image, ent.prefix) {
			return ent.mime
		}
	}

	return ""
}

// formatTable contains a table of known file magic prefixes
// for the document format autodetection.
var formatTable = []struct {
	prefix []byte
	mime   string
}{
	{[]byte{'B', 'M'}, DocumentFormatBMP},
	{[]byte{0xff, 0xd8}, DocumentFormatJPEG},
	{[]byte{'%', 'P', 'D', 'F', '-'}, DocumentFormatPDF},
	{[]byte{0x89, 'P', 'N', 'G', 0x0d, 0x0a, 0x1a, 0x0a},
		DocumentFormatPNG},
	{[]byte{'I', 'I', '*', 0}, DocumentFormatTIFF},
	{[]byte{'M', 'M', 0, '*'}, DocumentFormatTIFF},
}
