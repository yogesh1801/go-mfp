// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//

package wsscan

import (
	"fmt"

	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// MediaSides contains the parameters that are unique
// to each physical side of the scanned media.
type MediaSides struct {
	MediaFront MediaSide
	MediaBack  optional.Val[MediaSide]
	MustHonor  optional.Val[BooleanElement]
}

// decodeMediaSides decodes a MediaSides from an XML element.
func decodeMediaSides(root xmldoc.Element) (MediaSides, error) {
	var ms MediaSides

	// Decode MustHonor attribute if present
	if attr, found := root.AttrByName(NsWSCN + ":MustHonor"); found {
		boolVal := BooleanElement(attr.Value)
		if err := boolVal.Validate(); err != nil {
			return ms, err
		}
		ms.MustHonor = optional.New(boolVal)
	}

	// MediaFront is required, MediaBack is optional
	mediaFront := xmldoc.Lookup{
		Name:     NsWSCN + ":MediaFront",
		Required: true,
	}
	mediaBack := xmldoc.Lookup{
		Name:     NsWSCN + ":MediaBack",
		Required: false,
	}

	missed := root.Lookup(&mediaFront, &mediaBack)
	if missed != nil {
		return ms, xmldoc.XMLErrMissed(missed.Name)
	}

	// Decode MediaFront (required)
	mf, err := decodeMediaSide(mediaFront.Elem)
	if err != nil {
		return ms, fmt.Errorf("MediaFront: %w", err)
	}
	ms.MediaFront = mf

	// Decode MediaBack if present
	if mediaBack.Found {
		mb, err := decodeMediaSide(mediaBack.Elem)
		if err != nil {
			return ms, fmt.Errorf("MediaBack: %w", err)
		}
		ms.MediaBack = optional.New(mb)
	}

	return ms, nil
}

// toXML creates an XML element for MediaSides.
func (ms MediaSides) toXML(name string) xmldoc.Element {
	elm := xmldoc.Element{Name: name}

	// Add MustHonor attribute if present
	if ms.MustHonor != nil {
		elm.Attrs = append(elm.Attrs, xmldoc.Attr{
			Name:  NsWSCN + ":MustHonor",
			Value: string(optional.Get(ms.MustHonor)),
		})
	}

	// Add MediaFront (required)
	elm.Children = append(elm.Children,
		ms.MediaFront.toXML(NsWSCN+":MediaFront"))

	// Add MediaBack if present
	if ms.MediaBack != nil {
		elm.Children = append(elm.Children,
			optional.Get(ms.MediaBack).toXML(NsWSCN+":MediaBack"))
	}

	return elm
}
