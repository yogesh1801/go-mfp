// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// rotation value

package wsscan

import (
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// RotationValue defines the supported rotation value for a scan device.
type RotationValue int

// RotationValue represents possible rotation angles in degrees.
// Rotation is applied in the clockwise direction.
const (
	UnknownRotationValue RotationValue = iota // rotation unknown or not specified
	Rotation0                                 // no rotation
	Rotation90                                // 90 degrees
	Rotation180                               // 180 degrees
	Rotation270                               // 270 degrees
)

// decodeRotationValue decodes [RotationValue] from the XML tree.
func decodeRotationValue(root xmldoc.Element) (RotationValue, error) {
	return decodeEnum(root, DecodeRotationValue)
}

// toXML generates XML tree for the [RotationValue].
func (rv RotationValue) toXML(name string) xmldoc.Element {
	return xmldoc.Element{
		Name: name,
		Text: rv.String(),
	}
}

// String returns a string representation of the [RotationValue].
func (rv RotationValue) String() string {
	switch rv {
	case Rotation0:
		return "0"
	case Rotation90:
		return "90"
	case Rotation180:
		return "180"
	case Rotation270:
		return "270"
	}
	return "Unknown"
}

// DecodeRotationValue decodes [RotationValue] out of its XML string representation.
func DecodeRotationValue(s string) RotationValue {
	switch s {
	case "0":
		return Rotation0
	case "90":
		return Rotation90
	case "180":
		return Rotation180
	case "270":
		return Rotation270
	}
	return UnknownRotationValue
}
