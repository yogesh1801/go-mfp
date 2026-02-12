// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// CompressionQualityFactor:
// specifies an idealized integer amount of image quality

package wsscan

import (
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// CompressionQualityFactor specifies an idealized integer amount of
// image quality, on a scale from 0 through 100.
type CompressionQualityFactor ValWithOptions[int]

// decodeCompressionQualityFactor decodes a CompressionQualityFactor
// from an XML element.
func decodeCompressionQualityFactor(root xmldoc.Element) (
	CompressionQualityFactor, error) {
	var base ValWithOptions[int]
	decoded, err := base.decodeValWithOptions(root, intValueDecoder)
	if err != nil {
		return CompressionQualityFactor{}, err
	}
	return CompressionQualityFactor(decoded), nil
}

// toXML converts a CompressionQualityFactor to an XML element.
func (cqf CompressionQualityFactor) toXML(name string) xmldoc.Element {
	return ValWithOptions[int](cqf).toXML(name, intValueEncoder)
}
