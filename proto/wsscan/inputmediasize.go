// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// InputMediaSize: specifies the size of the media to be scanned

package wsscan

import (
	"fmt"

	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// InputMediaSize specifies the size of the media to be
// scanned for the current job.
// It contains required Height and Width child elements.
type InputMediaSize struct {
	Height ValWithOptions[int]
	Width  ValWithOptions[int]
}

// decodeInputMediaSize decodes an InputMediaSize from an XML element.
func decodeInputMediaSize(root xmldoc.Element) (InputMediaSize, error) {
	var ims InputMediaSize

	// Setup lookups for required child elements
	heightLookup := xmldoc.Lookup{
		Name:     NsWSCN + ":Height",
		Required: true,
	}
	widthLookup := xmldoc.Lookup{
		Name:     NsWSCN + ":Width",
		Required: true,
	}

	missed := root.Lookup(&heightLookup, &widthLookup)
	if missed != nil {
		return ims, xmldoc.XMLErrMissed(missed.Name)
	}

	// Decode Height
	height, err := ims.Height.decodeValWithOptions(
		heightLookup.Elem, intValueDecoder)
	if err != nil {
		return ims, fmt.Errorf("Height: %w", err)
	}
	ims.Height = height

	// Decode Width
	width, err := ims.Width.decodeValWithOptions(
		widthLookup.Elem, intValueDecoder)
	if err != nil {
		return ims, fmt.Errorf("Width: %w", err)
	}
	ims.Width = width

	return ims, nil
}

// toXML converts an InputMediaSize to an XML element.
func (ims InputMediaSize) toXML(name string) xmldoc.Element {
	return xmldoc.Element{
		Name: name,
		Children: []xmldoc.Element{
			ims.Height.toXML(NsWSCN+":Height", intValueEncoder),
			ims.Width.toXML(NsWSCN+":Width", intValueEncoder),
		},
	}
}
