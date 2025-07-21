// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// format value

package wsscan

import (
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// FormatValue defines the file format and compression type supported by the scan device.
type FormatValue int

// Known format values (constants for well-known types)
const (
	UnknownFormatValue FormatValue = iota
	DIB
	EXIF
	JBIG
	JFIF
	JPEG2K
	PDFA
	PNG
	TIFFSingleUncompressed
	TIFFSingleG4
	TIFFSingleG3MH
	TIFFSingleJPEGTN2
	TIFFMultiUncompressed
	TIFFMultiG4
	TIFFMultiG3MH
	TIFFMultiJPEGTN2
	XPS
)

// decodeFormatValue decodes [FormatValue] from the XML tree.
func decodeFormatValue(root xmldoc.Element) (fv FormatValue, err error) {
	return decodeEnum(root, DecodeFormatValue)
}

// toXML generates XML tree for the [FormatValue].
func (fv FormatValue) toXML(name string) xmldoc.Element {
	return xmldoc.Element{
		Name: name,
		Text: fv.String(),
	}
}

// String returns a string representation of the [FormatValue].
func (fv FormatValue) String() string {
	switch fv {
	case DIB:
		return "dib"
	case EXIF:
		return "exif"
	case JBIG:
		return "jbig"
	case JFIF:
		return "jfif"
	case JPEG2K:
		return "jpeg2k"
	case PDFA:
		return "pdf-a"
	case PNG:
		return "png"
	case TIFFSingleUncompressed:
		return "tiff-single-uncompressed"
	case TIFFSingleG4:
		return "tiff-single-g4"
	case TIFFSingleG3MH:
		return "tiff-single-g3mh"
	case TIFFSingleJPEGTN2:
		return "tiff-single-jpeg-tn2"
	case TIFFMultiUncompressed:
		return "tiff-multi-uncompressed"
	case TIFFMultiG4:
		return "tiff-multi-g4"
	case TIFFMultiG3MH:
		return "tiff-multi-g3mh"
	case TIFFMultiJPEGTN2:
		return "tiff-multi-jpeg-tn2"
	case XPS:
		return "xps"
	}
	return "Unknown"
}

// DecodeFormatValue decodes [FormatValue] out of its XML string representation.
func DecodeFormatValue(s string) FormatValue {
	switch s {
	case "dib":
		return DIB
	case "exif":
		return EXIF
	case "jbig":
		return JBIG
	case "jfif":
		return JFIF
	case "jpeg2k":
		return JPEG2K
	case "pdf-a":
		return PDFA
	case "png":
		return PNG
	case "tiff-single-uncompressed":
		return TIFFSingleUncompressed
	case "tiff-single-g4":
		return TIFFSingleG4
	case "tiff-single-g3mh":
		return TIFFSingleG3MH
	case "tiff-single-jpeg-tn2":
		return TIFFSingleJPEGTN2
	case "tiff-multi-uncompressed":
		return TIFFMultiUncompressed
	case "tiff-multi-g4":
		return TIFFMultiG4
	case "tiff-multi-g3mh":
		return TIFFMultiG3MH
	case "tiff-multi-jpeg-tn2":
		return TIFFMultiJPEGTN2
	case "xps":
		return XPS
	}
	return UnknownFormatValue
}
