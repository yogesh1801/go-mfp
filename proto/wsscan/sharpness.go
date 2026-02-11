// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// Sharpness: specifies the sharpness adjustment for scanned images

package wsscan

import "github.com/OpenPrinting/go-mfp/util/xmldoc"

// Sharpness specifies the relative amount to reduce or enhance the sharpness
// of the scanned document.
type Sharpness ValWithOptions[int]

// decodeSharpness decodes a Sharpness from an XML element.
func decodeSharpness(root xmldoc.Element) (Sharpness, error) {
	var base ValWithOptions[int]
	decoded, err := base.decodeValWithOptions(root, intValueDecoder)
	if err != nil {
		return Sharpness{}, err
	}
	return Sharpness(decoded), nil
}

// toXML converts a Sharpness to an XML element.
func (s Sharpness) toXML(name string) xmldoc.Element {
	return ValWithOptions[int](s).toXML(name, intValueEncoder)
}
