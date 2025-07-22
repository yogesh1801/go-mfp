// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// contrast supported

package wsscan

import (
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// ContrastSupported is a BooleanElement representing whether the scan
// device supports user control of the scan contrast setting.
type ContrastSupported BooleanElement

// IsValid returns true if the value is a valid ContrastSupported value.
func (cs ContrastSupported) IsValid() bool {
	return BooleanElement(cs).IsValid()
}

// Bool returns the boolean value of ContrastSupported.
func (cs ContrastSupported) Bool() (bool, error) {
	return BooleanElement(cs).Bool()
}

// toXML converts a ContrastSupported to an XML element.
func (cs ContrastSupported) toXML(name string) xmldoc.Element {
	return BooleanElement(cs).toXML(name)
}

// decodeContrastSupported decodes a ContrastSupported
// from an XML element.
func decodeContrastSupported(root xmldoc.Element) (
	ContrastSupported, error) {
	val, err := decodeBooleanElement(root)
	return ContrastSupported(val), err
}
