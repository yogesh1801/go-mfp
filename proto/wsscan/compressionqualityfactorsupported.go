// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// compression quality factor supported

package wsscan

import (
	"fmt"

	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// CompressionQualityFactorSupported represents the
// <wscn:CompressionQualityFactorSupported> element,
// containing MinValue and MaxValue.
type CompressionQualityFactorSupported struct {
	MinValue MinValue
	MaxValue MaxValue
}

// toXML generates XML tree for the [CompressionQualityFactorSupported].
func (cqfs CompressionQualityFactorSupported) toXML(name string) xmldoc.Element {
	return xmldoc.Element{
		Name: name,
		Children: []xmldoc.Element{
			cqfs.MinValue.toXML(NsWSCN + ":MinValue"),
			cqfs.MaxValue.toXML(NsWSCN + ":MaxValue"),
		},
	}
}

// decodeCompressionQualityFactorSupported decodes
// [CompressionQualityFactorSupported] from the XML tree.
func decodeCompressionQualityFactorSupported(root xmldoc.Element) (cqfs CompressionQualityFactorSupported, err error) {
	defer func() { err = xmldoc.XMLErrWrap(root, err) }()

	minLookup := xmldoc.Lookup{
		Name:     NsWSCN + ":MinValue",
		Required: true,
	}
	maxLookup := xmldoc.Lookup{
		Name:     NsWSCN + ":MaxValue",
		Required: true,
	}

	missed := root.Lookup(&minLookup, &maxLookup)
	if missed != nil {
		return cqfs, xmldoc.XMLErrMissed(missed.Name)
	}

	min, err := decodeMinValue(minLookup.Elem)
	if err != nil {
		return cqfs, fmt.Errorf("invalid MinValue: %w", err)
	}
	cqfs.MinValue = min

	max, err := decodeMaxValue(maxLookup.Elem)
	if err != nil {
		return cqfs, fmt.Errorf("invalid MaxValue: %w", err)
	}
	cqfs.MaxValue = max

	if err := cqfs.Validate(); err != nil {
		return cqfs, err
	}
	return cqfs, nil
}

// Validate checks that MinValue and MaxValue are
// within [1, 100] and MinValue <= MaxValue.
func (cqfs CompressionQualityFactorSupported) Validate() error {
	if cqfs.MinValue < 1 || cqfs.MinValue > 100 {
		return fmt.Errorf("MinValue must be between 1 and 100")
	}
	if cqfs.MaxValue < 1 || cqfs.MaxValue > 100 {
		return fmt.Errorf("MaxValue must be between 1 and 100")
	}
	if int(cqfs.MinValue) > int(cqfs.MaxValue) {
		return fmt.Errorf("MinValue must be less than or equal to MaxValue")
	}
	return nil
}
