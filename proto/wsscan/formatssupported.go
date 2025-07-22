// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// formats supported (FormatsSupported)

package wsscan

import (
	"fmt"

	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// FormatsSupported represents the <wscn:FormatsSupported> element,
// containing a list of FormatValue elements.
type FormatsSupported struct {
	Values []FormatValue
}

// toXML generates XML tree for the [FormatsSupported].
func (fs FormatsSupported) toXML(name string) xmldoc.Element {
	children := make([]xmldoc.Element, len(fs.Values))
	for i, v := range fs.Values {
		children[i] = v.toXML(NsWSCN + ":FormatValue")
	}
	return xmldoc.Element{
		Name:     name,
		Children: children,
	}
}

// decodeFormatsSupported decodes [FormatsSupported] from the XML tree.
func decodeFormatsSupported(root xmldoc.Element) (
	fs FormatsSupported, err error) {
	defer func() { err = xmldoc.XMLErrWrap(root, err) }()

	var found bool
	for _, child := range root.Children {
		if child.Name == NsWSCN+":FormatValue" {
			val, err := decodeFormatValue(child)
			if err != nil {
				return fs, fmt.Errorf("invalid FormatValue: %w", err)
			}
			fs.Values = append(fs.Values, val)
			found = true
		}
	}
	if !found {
		return fs, fmt.Errorf("at least one FormatValue is required")
	}

	return fs, nil
}
