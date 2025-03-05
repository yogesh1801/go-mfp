// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scan color mode

package escl

import (
	"github.com/alexpevzner/mfp/util/xmldoc"
)

// ColorMode specifies combination of the color mode (color/grayscale/1-bit
// black and white) with the bit depth.
type ColorMode int

// Known color modes:
const (
	UnknownColorMode ColorMode = iota // Unknown color mode
	BlackAndWhite1                    // 1-bit black and white
	Grayscale8                        // 8-bit grayscale
	Grayscale16                       // 16-bit grayscale
	RGB24                             // 8-bit per channel RGB
	RGB48                             // 16-bit per channel RGB
)

// decodeColorMode decodes [ColorMode] from the XML tree.
func decodeColorMode(root xmldoc.Element) (cm ColorMode, err error) {
	return decodeEnum(root, DecodeColorMode)
}

// toXML generates XML tree for the [ColorMode].
func (cm ColorMode) toXML(name string) xmldoc.Element {
	return xmldoc.Element{
		Name: name,
		Text: cm.String(),
	}
}

// String returns a string representation of the [ColorMode]
func (cm ColorMode) String() string {
	switch cm {
	case BlackAndWhite1:
		return "BlackAndWhite1"
	case Grayscale8:
		return "Grayscale8"
	case Grayscale16:
		return "Grayscale16"
	case RGB24:
		return "RGB24"
	case RGB48:
		return "RGB48"
	}

	return "Unknown"
}

// DecodeColorMode decodes [ColorMode] out of its XML string representation.
func DecodeColorMode(s string) ColorMode {
	switch s {
	case "BlackAndWhite1":
		return BlackAndWhite1
	case "Grayscale8":
		return Grayscale8
	case "Grayscale16":
		return Grayscale16
	case "RGB24":
		return RGB24
	case "RGB48":
		return RGB48
	}

	return UnknownColorMode
}
