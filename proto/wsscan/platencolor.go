// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// platen color

package wsscan

import (
	"fmt"

	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// PlatenColor represents the <wscn:PlatenColor> element,
// containing a list of ColorEntry elements.
type PlatenColor struct {
	Values []ColorEntry
}

// toXML generates XML tree for the [PlatenColor].
func (pc PlatenColor) toXML(name string) xmldoc.Element {
	children := make([]xmldoc.Element, len(pc.Values))
	for i, v := range pc.Values {
		children[i] = v.toXML(NsWSCN + ":ColorEntry")
	}
	return xmldoc.Element{
		Name:     name,
		Children: children,
	}
}

// decodePlatenColor decodes [PlatenColor] from the XML tree.
func decodePlatenColor(root xmldoc.Element) (pc PlatenColor, err error) {
	defer func() { err = xmldoc.XMLErrWrap(root, err) }()

	var found bool
	for _, child := range root.Children {
		if child.Name == NsWSCN+":ColorEntry" {
			val, err := decodeColorEntry(child)
			if err != nil {
				return pc, fmt.Errorf("invalid ColorEntry: %w", err)
			}
			pc.Values = append(pc.Values, val)
			found = true
		}
	}
	if !found {
		return pc, fmt.Errorf("at least one ColorEntry is required")
	}

	return pc, nil
}
