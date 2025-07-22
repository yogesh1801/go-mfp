// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// compression quality factor supported

package wsscan

import (
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// CompressionQualityFactorSupported represents the
// <wscn:CompressionQualityFactorSupported> element,
// containing a range of supported compression quality factors.
type CompressionQualityFactorSupported struct {
	RangeElement
}

// toXML generates XML tree for the [CompressionQualityFactorSupported].
func (cqfs CompressionQualityFactorSupported) toXML(name string) xmldoc.Element {
	return xmldoc.Element{
		Name:     name,
		Children: cqfs.RangeElement.toXML(),
	}
}

// decodeCompressionQualityFactorSupported decodes
// [CompressionQualityFactorSupported] from the XML tree.
func decodeCompressionQualityFactorSupported(root xmldoc.Element) (cqfs CompressionQualityFactorSupported, err error) {
	defer func() { err = xmldoc.XMLErrWrap(root, err) }()

	cqfs.RangeElement, err = decodeRangeElement(root)
	if err != nil {
		return cqfs, err
	}

	if err = cqfs.Validate(); err != nil {
		return cqfs, err
	}

	return cqfs, nil
}

// Validate checks that MinValue and MaxValue are
// within [1, 100] and MinValue <= MaxValue.
func (cqfs CompressionQualityFactorSupported) Validate() error {
	return cqfs.RangeElement.Validate(1, 100)
}
