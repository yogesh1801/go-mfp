// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// Contrast: specifies the contrast adjustment for scanned images

package wsscan

import "github.com/OpenPrinting/go-mfp/util/xmldoc"

// Contrast specifies the relative amount to reduce or enhance the contrast
// of the scanned document.
type Contrast ValWithOptions[int]

// decodeContrast decodes a Contrast from an XML element.
func decodeContrast(root xmldoc.Element) (Contrast, error) {
	var base ValWithOptions[int]
	decoded, err := base.decodeValWithOptions(root, intValueDecoder)
	if err != nil {
		return Contrast{}, err
	}
	return Contrast(decoded), nil
}

// toXML converts a Contrast to an XML element.
func (c Contrast) toXML(name string) xmldoc.Element {
	return ValWithOptions[int](c).toXML(name, intValueEncoder)
}
