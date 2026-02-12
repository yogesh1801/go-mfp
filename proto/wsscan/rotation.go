// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// Rotation: specifies the amount to rotate each image of the scanned document

package wsscan

import (
	"fmt"

	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// Rotation specifies the amount to rotate each image of the scanned document.
// The text value is a RotationValue (0, 90, 180, or 270 degrees).
type Rotation struct {
	ValWithOptions[RotationValue]
}

// rotationDecoder is the decoder function for use with ValWithOptions
func rotationDecoder(s string) (RotationValue, error) {
	val := DecodeRotationValue(s)
	if val == UnknownRotationValue {
		return val, fmt.Errorf("unknown Rotation value: %s", s)
	}
	return val, nil
}

// rotationEncoder is the encoder function for use with ValWithOptions
func rotationEncoder(rv RotationValue) string {
	return rv.String()
}

// decodeRotation decodes Rotation from the XML tree
func decodeRotation(root xmldoc.Element) (Rotation, error) {
	var r Rotation
	decoded, err := r.ValWithOptions.decodeValWithOptions(root, rotationDecoder)
	if err != nil {
		return r, err
	}
	r.ValWithOptions = decoded
	return r, nil
}

// toXML generates XML tree for the Rotation
func (r Rotation) toXML(name string) xmldoc.Element {
	return r.ValWithOptions.toXML(name, rotationEncoder)
}
