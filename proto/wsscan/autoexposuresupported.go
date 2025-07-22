// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// auto exposure supported

package wsscan

import (
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// AutoExposureSupported is a BooleanElement representing whether the scan
// device supports automatic exposure adjustment.
type AutoExposureSupported BooleanElement

// IsValid returns true if the value is a valid AutoExposureSupported value.
func (aes AutoExposureSupported) IsValid() bool {
	return BooleanElement(aes).IsValid()
}

// Bool returns the boolean value of AutoExposureSupported.
func (aes AutoExposureSupported) Bool() (bool, error) {
	return BooleanElement(aes).Bool()
}

// toXML converts an AutoExposureSupported to an XML element.
func (aes AutoExposureSupported) toXML(name string) xmldoc.Element {
	return BooleanElement(aes).toXML(name)
}

// decodeAutoExposureSupported decodes an AutoExposureSupported
// from an XML element.
func decodeAutoExposureSupported(
	root xmldoc.Element,
) (AutoExposureSupported, error) {
	val, err := decodeBooleanElement(root)
	return AutoExposureSupported(val), err
}
