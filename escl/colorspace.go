// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scan color space

package escl

import "github.com/alexpevzner/mfp/xmldoc"

// ColorSpace defines the color space used for scanning.
type ColorSpace int

// Known color modes:
const (
	UnknownColorSpace ColorSpace = iota // Unknown color mode
	SRGB                                // sRGG
)

// decodeColorSpace decodes [ColorSpace] from the XML tree.
func decodeColorSpace(root xmldoc.Element) (sps ColorSpace, err error) {
	return decodeEnum(root, DecodeColorSpace)
}

// toXML generates XML tree for the [ColorSpace].
func (sps ColorSpace) toXML(name string) xmldoc.Element {
	return xmldoc.Element{
		Name: name,
		Text: sps.String(),
	}
}

// String returns a string representation of the [ColorSpace]
func (sps ColorSpace) String() string {
	switch sps {
	case SRGB:
		return "sRGB"
	}

	return "Unknown"
}

// DecodeColorSpace decodes [ColorSpace] out of its XML string representation.
func DecodeColorSpace(s string) ColorSpace {
	switch s {
	case "sRGB":
		return SRGB
	}

	return UnknownColorSpace
}
