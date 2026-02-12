// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// Scaling: specifies the scaling of both width and height of the scanned document

package wsscan

import (
	"fmt"
	"strconv"

	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// Scaling specifies the scaling of both width and height of
// the scanned document.
// It contains required ScalingHeight and ScalingWidth child elements.
type Scaling struct {
	ScalingHeight ValWithOptions[int]
	ScalingWidth  ValWithOptions[int]
	MustHonor     optional.Val[BooleanElement]
}

// decodeScaling decodes a Scaling from an XML element.
func decodeScaling(root xmldoc.Element) (Scaling, error) {
	var s Scaling

	// Decode MustHonor attribute if present
	if attr, found := root.AttrByName(NsWSCN + ":MustHonor"); found {
		boolVal := BooleanElement(attr.Value)
		if err := boolVal.Validate(); err != nil {
			return s, err
		}
		s.MustHonor = optional.New(boolVal)
	}

	// Setup lookups for child elements
	height := xmldoc.Lookup{
		Name:     NsWSCN + ":ScalingHeight",
		Required: true,
	}
	width := xmldoc.Lookup{
		Name:     NsWSCN + ":ScalingWidth",
		Required: true,
	}

	missed := root.Lookup(&height, &width)
	if missed != nil {
		return s, xmldoc.XMLErrMissed(missed.Name)
	}

	// Decode ScalingHeight
	h, err := s.ScalingHeight.decodeValWithOptions(height.Elem,
		func(s string) (int, error) {
			return decodeInt(xmldoc.Element{Text: s})
		})
	if err != nil {
		return s, fmt.Errorf("ScalingHeight: %w", err)
	}
	s.ScalingHeight = h

	// Decode ScalingWidth
	w, err := s.ScalingWidth.decodeValWithOptions(width.Elem,
		func(s string) (int, error) {
			return decodeInt(xmldoc.Element{Text: s})
		})
	if err != nil {
		return s, fmt.Errorf("ScalingWidth: %w", err)
	}
	s.ScalingWidth = w

	return s, nil
}

// toXML creates an XML element for Scaling.
func (s Scaling) toXML(name string) xmldoc.Element {
	elm := xmldoc.Element{Name: name}

	// Add MustHonor attribute if present
	if s.MustHonor != nil {
		elm.Attrs = append(elm.Attrs, xmldoc.Attr{
			Name:  NsWSCN + ":MustHonor",
			Value: string(optional.Get(s.MustHonor)),
		})
	}

	// Add ScalingHeight and ScalingWidth child elements
	intToString := func(i int) string {
		return strconv.Itoa(i)
	}
	elm.Children = append(elm.Children,
		s.ScalingHeight.toXML(NsWSCN+":ScalingHeight", intToString),
		s.ScalingWidth.toXML(NsWSCN+":ScalingWidth", intToString))

	return elm
}
