// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// scaling range supported (ScalingWidth)

package wsscan

import (
	"fmt"

	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// ScalingWidth represents the <wscn:ScalingWidth> element,
// containing MinValue and MaxValue.
type ScalingWidth struct {
	MinValue MinValue
	MaxValue MaxValue
}

// toXML generates XML tree for the [ScalingWidth].
func (sw ScalingWidth) toXML(name string) xmldoc.Element {
	return xmldoc.Element{
		Name: name,
		Children: []xmldoc.Element{
			sw.MinValue.toXML(NsWSCN + ":MinValue"),
			sw.MaxValue.toXML(NsWSCN + ":MaxValue"),
		},
	}
}

// decodeScalingWidth decodes [ScalingWidth] from the XML tree.
func decodeScalingWidth(root xmldoc.Element) (sw ScalingWidth, err error) {
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
		return sw, xmldoc.XMLErrMissed(missed.Name)
	}

	min, err := decodeMinValue(minLookup.Elem)
	if err != nil {
		return sw, fmt.Errorf("invalid MinValue: %w", err)
	}
	sw.MinValue = min

	max, err := decodeMaxValue(maxLookup.Elem)
	if err != nil {
		return sw, fmt.Errorf("invalid MaxValue: %w", err)
	}
	sw.MaxValue = max

	if err := sw.Validate(); err != nil {
		return sw, err
	}
	return sw, nil
}

// Validate checks that MinValue and MaxValue are within [1, 1000] and MinValue <= MaxValue.
func (sw ScalingWidth) Validate() error {
	if sw.MinValue < 1 || sw.MinValue > 1000 {
		return fmt.Errorf("MinValue must be between 1 and 1000")
	}
	if sw.MaxValue < 1 || sw.MaxValue > 1000 {
		return fmt.Errorf("MaxValue must be between 1 and 1000")
	}
	if int(sw.MinValue) > int(sw.MaxValue) {
		return fmt.Errorf("MinValue must be less than or equal to MaxValue")
	}
	return nil
}
