// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// rotations supported (RotationsSupported)

package wsscan

import (
	"fmt"

	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// RotationsSupported represents the <wscn:RotationsSupported> element,
// containing a list of RotationValue elements.
type RotationsSupported struct {
	Values []RotationValue
}

// toXML generates XML tree for the [RotationsSupported].
func (rs RotationsSupported) toXML(name string) xmldoc.Element {
	children := make([]xmldoc.Element, len(rs.Values))
	for i, v := range rs.Values {
		children[i] = v.toXML(NsWSCN + ":RotationValue")
	}
	return xmldoc.Element{
		Name:     name,
		Children: children,
	}
}

// decodeRotationsSupported decodes [RotationsSupported] from the XML tree.
func decodeRotationsSupported(root xmldoc.Element) (rs RotationsSupported, err error) {
	defer func() { err = xmldoc.XMLErrWrap(root, err) }()

	var found bool
	for _, child := range root.Children {
		if child.Name == NsWSCN+":RotationValue" {
			val, err := decodeRotationValue(child)
			if err != nil {
				return rs, fmt.Errorf("invalid RotationValue: %w", err)
			}
			rs.Values = append(rs.Values, val)
			found = true
		}
	}
	if !found {
		return rs, fmt.Errorf("at least one RotationValue is required")
	}

	return rs, nil
}
