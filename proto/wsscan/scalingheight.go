// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// scaling range supported (ScalingHeight)

package wsscan

import (
	"fmt"

	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// ScalingHeight represents the <wscn:ScalingHeight> element,
// containing MinValue and MaxValue.
type ScalingHeight struct {
	MinValue MinValue
	MaxValue MaxValue
}

// toXML generates XML tree for the [ScalingHeight].
func (sh ScalingHeight) toXML(name string) xmldoc.Element {
	return xmldoc.Element{
		Name: name,
		Children: []xmldoc.Element{
			sh.MinValue.toXML(NsWSCN + ":MinValue"),
			sh.MaxValue.toXML(NsWSCN + ":MaxValue"),
		},
	}
}

// decodeScalingHeight decodes [ScalingHeight] from the XML tree.
func decodeScalingHeight(root xmldoc.Element) (sh ScalingHeight, err error) {
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
		return sh, xmldoc.XMLErrMissed(missed.Name)
	}

	min, err := decodeMinValue(minLookup.Elem)
	if err != nil {
		return sh, fmt.Errorf("invalid MinValue: %w", err)
	}
	sh.MinValue = min

	max, err := decodeMaxValue(maxLookup.Elem)
	if err != nil {
		return sh, fmt.Errorf("invalid MaxValue: %w", err)
	}
	sh.MaxValue = max

	if err := sh.Validate(); err != nil {
		return sh, err
	}
	return sh, nil
}

// Validate checks that MinValue and MaxValue are within [1, 1000] and MinValue <= MaxValue.
func (sh ScalingHeight) Validate() error {
	if sh.MinValue < 1 || sh.MinValue > 1000 {
		return fmt.Errorf("MinValue must be between 1 and 1000")
	}
	if sh.MaxValue < 1 || sh.MaxValue > 1000 {
		return fmt.Errorf("MaxValue must be between 1 and 1000")
	}
	if int(sh.MinValue) > int(sh.MaxValue) {
		return fmt.Errorf("MinValue must be less than or equal to MaxValue")
	}
	return nil
}
