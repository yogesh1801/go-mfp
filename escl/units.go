// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Units for coordinates and resolutions.

package escl

import "github.com/alexpevzner/mfp/xmldoc"

// Units specifies the feed direction of the input media
// (affects the resulting image orientation).
type Units int

// Known Units.
//
// The only supported value for eSCL is 300 DPI.
const (
	UnknownUnits            Units = iota // Unknown CCD
	ThreeHundredthsOfInches              // 300 DPI
)

// decodeUnits decodes [Units] from the XML tree.
func decodeUnits(root xmldoc.Element) (units Units, err error) {
	return decodeEnum(root, DecodeUnits, NsScan)
}

// toXML generates XML tree for the [Units].
func (units Units) toXML(name string) xmldoc.Element {
	return xmldoc.Element{
		Name: name,
		Text: NsScan + ":" + units.String(),
	}
}

// String returns a string representation of the [Units]
func (units Units) String() string {
	switch units {
	case ThreeHundredthsOfInches:
		return "ThreeHundredthsOfInches"
	}

	return "Unknown"
}

// DecodeUnits decodes [Units] out of its XML
// string representation.
func DecodeUnits(s string) Units {
	switch s {
	case "ThreeHundredthsOfInches":
		return ThreeHundredthsOfInches
	}

	return UnknownUnits
}
