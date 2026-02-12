// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// Resolution: specifies the resolution of the scanned image

package wsscan

import (
	"fmt"
	"strconv"

	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// Resolution specifies the resolution of the scanned image.
// It contains Height and Width child elements.
type Resolution struct {
	Height    ValWithOptions[int]
	Width     ValWithOptions[int]
	MustHonor optional.Val[BooleanElement]
}

// decodeResolution decodes a Resolution from an XML element.
func decodeResolution(root xmldoc.Element) (Resolution, error) {
	var r Resolution

	// Decode MustHonor attribute if present
	if attr, found := root.AttrByName(NsWSCN + ":MustHonor"); found {
		boolVal := BooleanElement(attr.Value)
		if err := boolVal.Validate(); err != nil {
			return r, err
		}
		r.MustHonor = optional.New(boolVal)
	}

	// Setup lookups for child elements
	height := xmldoc.Lookup{
		Name:     NsWSCN + ":Height",
		Required: true,
	}
	width := xmldoc.Lookup{
		Name:     NsWSCN + ":Width",
		Required: true,
	}

	missed := root.Lookup(&height, &width)
	if missed != nil {
		return r, xmldoc.XMLErrMissed(missed.Name)
	}

	// Decode Height
	h, err := r.Height.decodeValWithOptions(height.Elem,
		func(s string) (int, error) {
			return decodeInt(xmldoc.Element{Text: s})
		})
	if err != nil {
		return r, fmt.Errorf("Height: %w", err)
	}
	r.Height = h

	// Decode Width
	w, err := r.Width.decodeValWithOptions(width.Elem,
		func(s string) (int, error) {
			return decodeInt(xmldoc.Element{Text: s})
		})
	if err != nil {
		return r, fmt.Errorf("Width: %w", err)
	}
	r.Width = w

	return r, nil
}

// toXML creates an XML element for Resolution.
func (r Resolution) toXML(name string) xmldoc.Element {
	elm := xmldoc.Element{Name: name}

	// Add MustHonor attribute if present
	if r.MustHonor != nil {
		elm.Attrs = append(elm.Attrs, xmldoc.Attr{
			Name:  NsWSCN + ":MustHonor",
			Value: string(optional.Get(r.MustHonor)),
		})
	}

	// Add Height and Width child elements
	intToString := func(i int) string {
		return strconv.Itoa(i)
	}
	elm.Children = append(elm.Children,
		r.Height.toXML(NsWSCN+":Height", intToString),
		r.Width.toXML(NsWSCN+":Width", intToString))

	return elm
}
