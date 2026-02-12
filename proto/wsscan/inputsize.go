// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// InputSize: specifies the size of the original scan media

package wsscan

import (
	// Added for error formatting
	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// InputSize specifies the size of the original scan media.
type InputSize struct {
	MustHonor              optional.Val[BooleanElement]
	DocumentSizeAutoDetect optional.Val[BooleanElement]
	InputMediaSize         InputMediaSize
}

// decodeInputSize decodes an InputSize from an XML element.
func decodeInputSize(root xmldoc.Element) (InputSize, error) {
	var is InputSize

	// Decode MustHonor attribute if present
	if attr, found := root.AttrByName(NsWSCN + ":MustHonor"); found {
		val := BooleanElement(attr.Value)
		if err := val.Validate(); err != nil {
			return is, err
		}
		is.MustHonor = optional.New(val)
	}

	// Setup lookups for child elements
	documentSizeAutoDetectLookup := xmldoc.Lookup{
		Name: NsWSCN + ":DocumentSizeAutoDetect",
	}
	inputMediaSizeLookup := xmldoc.Lookup{
		Name:     NsWSCN + ":InputMediaSize",
		Required: true,
	}

	missed := root.Lookup(&documentSizeAutoDetectLookup, &inputMediaSizeLookup)
	if missed != nil {
		return is, xmldoc.XMLErrMissed(missed.Name)
	}

	// Decode DocumentSizeAutoDetect if present
	if documentSizeAutoDetectLookup.Elem.Name != "" {
		val := BooleanElement(documentSizeAutoDetectLookup.Elem.Text)
		if err := val.Validate(); err != nil {
			return is, err
		}
		is.DocumentSizeAutoDetect = optional.New(val)
	}

	// Decode InputMediaSize (required)
	ims, err := decodeInputMediaSize(inputMediaSizeLookup.Elem)
	if err != nil {
		return is, err
	}
	is.InputMediaSize = ims

	return is, nil
}

// toXML converts an InputSize to an XML element.
func (is InputSize) toXML(name string) xmldoc.Element {
	elm := xmldoc.Element{Name: name}

	// Add MustHonor attribute if present
	if is.MustHonor != nil {
		elm.Attrs = []xmldoc.Attr{
			{
				Name:  NsWSCN + ":MustHonor",
				Value: string(optional.Get(is.MustHonor)),
			},
		}
	}

	var children []xmldoc.Element

	// Add DocumentSizeAutoDetect if present
	if is.DocumentSizeAutoDetect != nil {
		children = append(children, xmldoc.Element{
			Name: NsWSCN + ":DocumentSizeAutoDetect",
			Text: string(optional.Get(is.DocumentSizeAutoDetect)),
		})
	}

	// Add InputMediaSize (required)
	children = append(children,
		is.InputMediaSize.toXML(NsWSCN+":InputMediaSize"))

	elm.Children = children
	return elm
}
