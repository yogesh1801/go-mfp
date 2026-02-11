// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// Brightness: specifies the brightness adjustment for scanned images

package wsscan

import "github.com/OpenPrinting/go-mfp/util/xmldoc"

// Brightness specifies the relative amount to reduce or enhance the brightness
// of the scanned document.
type Brightness ValWithOptions[int]

// decodeBrightness decodes a Brightness from an XML element.
func decodeBrightness(root xmldoc.Element) (Brightness, error) {
	var base ValWithOptions[int]
	decoded, err := base.decodeValWithOptions(root, intValueDecoder)
	if err != nil {
		return Brightness{}, err
	}
	return Brightness(decoded), nil
}

// toXML converts a Brightness to an XML element.
func (b Brightness) toXML(name string) xmldoc.Element {
	return ValWithOptions[int](b).toXML(name, intValueEncoder)
}
