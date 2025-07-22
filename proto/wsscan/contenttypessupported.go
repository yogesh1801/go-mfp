// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// content types supported (ContentTypesSupported)

package wsscan

import (
	"fmt"

	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// ContentTypesSupported represents the <wscn:ContentTypesSupported> element,
// containing a list of ContentTypeValue elements.
type ContentTypesSupported struct {
	Values []ContentTypeValue
}

// toXML generates XML tree for the [ContentTypesSupported].
func (cts ContentTypesSupported) toXML(name string) xmldoc.Element {
	children := make([]xmldoc.Element, len(cts.Values))
	for i, v := range cts.Values {
		children[i] = v.toXML(NsWSCN + ":ContentTypeValue")
	}
	return xmldoc.Element{
		Name:     name,
		Children: children,
	}
}

// decodeContentTypesSupported decodes [ContentTypesSupported] from the XML tree.
func decodeContentTypesSupported(root xmldoc.Element) (
	cts ContentTypesSupported, err error) {
	defer func() { err = xmldoc.XMLErrWrap(root, err) }()

	var found bool
	for _, child := range root.Children {
		if child.Name == NsWSCN+":ContentTypeValue" {
			val, err := decodeContentTypeValue(child)
			if err != nil {
				return cts, fmt.Errorf("invalid ContentTypeValue: %w", err)
			}
			cts.Values = append(cts.Values, val)
			found = true
		}
	}
	if !found {
		return cts, fmt.Errorf("at least one ContentTypeValue is required")
	}

	return cts, nil
}
