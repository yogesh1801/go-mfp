// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Image format detection

package imgconv

import "bytes"

// MIME types for known document formats:
const (
	MIMETypeBMP  = "image/bmp"
	MIMETypeJPEG = "image/jpeg"
	MIMETypePDF  = "application/pdf"
	MIMETypePNG  = "image/png"
	MIMETypeTIFF = "image/tiff"
	MIMETypeData = "application/octet-stream"
)

// MIMETypeDetect detects image type by its few starting bytes
// and returns its MIME type.
//
// If format cannot be guessed, "" is returned.
func MIMETypeDetect(image []byte) string {
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
	{[]byte{'B', 'M'}, MIMETypeBMP},
	{[]byte{0xff, 0xd8}, MIMETypeJPEG},
	{[]byte{'%', 'P', 'D', 'F', '-'}, MIMETypePDF},
	{[]byte{0x89, 'P', 'N', 'G', 0x0d, 0x0a, 0x1a, 0x0a},
		MIMETypePNG},
	{[]byte{'I', 'I', '*', 0}, MIMETypeTIFF},
	{[]byte{'M', 'M', 0, '*'}, MIMETypeTIFF},
}
