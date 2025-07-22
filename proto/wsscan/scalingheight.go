// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// scaling range supported (ScalingHeight)

package wsscan

import (
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// ScalingHeight represents the <wscn:ScalingHeight> element,
// containing a range of supported scaling height values.
type ScalingHeight struct {
	RangeElement
}

// toXML generates XML tree for the [ScalingHeight].
func (sh ScalingHeight) toXML(name string) xmldoc.Element {
	return xmldoc.Element{
		Name:     name,
		Children: sh.RangeElement.toXML(),
	}
}

// decodeScalingHeight decodes [ScalingHeight] from the XML tree.
func decodeScalingHeight(root xmldoc.Element) (sh ScalingHeight, err error) {
	defer func() { err = xmldoc.XMLErrWrap(root, err) }()

	sh.RangeElement, err = decodeRangeElement(root)
	if err != nil {
		return sh, err
	}

	if err = sh.Validate(); err != nil {
		return sh, err
	}

	return sh, nil
}

// Validate checks that MinValue and MaxValue are within
// [1, 1000] and MinValue <= MaxValue.
func (sh ScalingHeight) Validate() error {
	return sh.RangeElement.Validate(1, 1000)
}
