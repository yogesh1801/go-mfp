// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// scaling range supported (ScalingRangeSupported)

package wsscan

import (
	"fmt"

	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// ScalingRangeSupported represents the <wscn:ScalingRangeSupported> element,
// containing ScalingWidth and ScalingHeight.
type ScalingRangeSupported struct {
	ScalingWidth  ScalingWidth
	ScalingHeight ScalingHeight
}

// toXML generates XML tree for the [ScalingRangeSupported].
func (srs ScalingRangeSupported) toXML(name string) xmldoc.Element {
	return xmldoc.Element{
		Name: name,
		Children: []xmldoc.Element{
			srs.ScalingWidth.toXML(NsWSCN + ":ScalingWidth"),
			srs.ScalingHeight.toXML(NsWSCN + ":ScalingHeight"),
		},
	}
}

// decodeScalingRangeSupported decodes [ScalingRangeSupported] from the XML tree.
func decodeScalingRangeSupported(root xmldoc.Element) (srs ScalingRangeSupported, err error) {
	defer func() { err = xmldoc.XMLErrWrap(root, err) }()

	widthLookup := xmldoc.Lookup{
		Name:     NsWSCN + ":ScalingWidth",
		Required: true,
	}
	heightLookup := xmldoc.Lookup{
		Name:     NsWSCN + ":ScalingHeight",
		Required: true,
	}

	missed := root.Lookup(&widthLookup, &heightLookup)
	if missed != nil {
		return srs, xmldoc.XMLErrMissed(missed.Name)
	}

	width, err := decodeScalingWidth(widthLookup.Elem)
	if err != nil {
		return srs, fmt.Errorf("invalid ScalingWidth: %w", err)
	}
	srs.ScalingWidth = width

	height, err := decodeScalingHeight(heightLookup.Elem)
	if err != nil {
		return srs, fmt.Errorf("invalid ScalingHeight: %w", err)
	}
	srs.ScalingHeight = height

	if err := srs.Validate(); err != nil {
		return srs, err
	}
	return srs, nil
}

// Validate checks that ScalingWidth and ScalingHeight are valid.
func (srs ScalingRangeSupported) Validate() error {
	if err := srs.ScalingWidth.Validate(); err != nil {
		return fmt.Errorf("ScalingWidth: %w", err)
	}
	if err := srs.ScalingHeight.Validate(); err != nil {
		return fmt.Errorf("ScalingHeight: %w", err)
	}
	return nil
}
