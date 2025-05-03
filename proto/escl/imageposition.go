// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Image position after ADF image justification.

package escl

import "github.com/OpenPrinting/go-mfp/util/xmldoc"

// ImagePosition specifies image position after document justification
// made by the ADF.
type ImagePosition int

// Known CCD Channels.
const (
	UnknownImagePosition ImagePosition = iota // Unknown CCD
	Left                                      // Left (horizontal)
	Right                                     // Right (horizontal)
	Top                                       // Top (vertical)
	Bottom                                    // Bottom (vertical)
	Center                                    // Center (both)
)

// decodeImagePosition decodes [ImagePosition] from the XML tree.
func decodeImagePosition(root xmldoc.Element) (pos ImagePosition, err error) {
	return decodeEnum(root, DecodeImagePosition)
}

// toXML generates XML tree for the [ImagePosition].
func (pos ImagePosition) toXML(name string) xmldoc.Element {
	return xmldoc.Element{
		Name: name,
		Text: pos.String(),
	}
}

// String returns a string representation of the [ImagePosition]
func (pos ImagePosition) String() string {
	switch pos {
	case Left:
		return "Left"
	case Right:
		return "Right"
	case Top:
		return "Top"
	case Bottom:
		return "Bottom"
	case Center:
		return "Center"
	}

	return "Unknown"
}

// DecodeImagePosition decodes [ImagePosition] out of its XML
// string representation.
func DecodeImagePosition(s string) ImagePosition {
	switch s {
	case "Left":
		return Left
	case "Right":
		return Right
	case "Top":
		return Top
	case "Bottom":
		return Bottom
	case "Center":
		return Center
	}

	return UnknownImagePosition
}
