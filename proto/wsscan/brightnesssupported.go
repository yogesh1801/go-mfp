// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// brightness supported

package wsscan

import (
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// BrightnessSupported is a BooleanElement representing whether the scan
// device supports user control of the scan brightness setting.
type BrightnessSupported BooleanElement

// IsValid returns true if the value is a valid BrightnessSupported value.
func (bs BrightnessSupported) IsValid() bool {
	return BooleanElement(bs).IsValid()
}

// Bool returns the boolean value of BrightnessSupported.
func (bs BrightnessSupported) Bool() (bool, error) {
	return BooleanElement(bs).Bool()
}

// toXML converts a BrightnessSupported to an XML element.
func (bs BrightnessSupported) toXML(name string) xmldoc.Element {
	return BooleanElement(bs).toXML(name)
}

// decodeBrightnessSupported decodes a BrightnessSupported
// from an XML element.
func decodeBrightnessSupported(root xmldoc.Element) (
	BrightnessSupported, error) {
	val, err := decodeBooleanElement(root)
	return BrightnessSupported(val), err
}
