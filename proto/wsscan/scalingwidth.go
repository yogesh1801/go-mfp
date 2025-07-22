// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// scaling range supported (ScalingWidth)

package wsscan

import (
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// ScalingWidth represents the <wscn:ScalingWidth> element,
// containing a range of supported scaling width values.
type ScalingWidth struct {
	RangeElement
}

// toXML generates XML tree for the [ScalingWidth].
func (sw ScalingWidth) toXML(name string) xmldoc.Element {
	return xmldoc.Element{
		Name:     name,
		Children: sw.RangeElement.toXML(),
	}
}

// decodeScalingWidth decodes [ScalingWidth] from the XML tree.
func decodeScalingWidth(root xmldoc.Element) (sw ScalingWidth, err error) {
	defer func() { err = xmldoc.XMLErrWrap(root, err) }()

	sw.RangeElement, err = decodeRangeElement(root)
	if err != nil {
		return sw, err
	}

	if err = sw.Validate(); err != nil {
		return sw, err
	}

	return sw, nil
}

// Validate checks that MinValue and MaxValue are within [1, 1000] and MinValue <= MaxValue.
func (sw ScalingWidth) Validate() error {
	return sw.RangeElement.Validate(1, 1000)
}
