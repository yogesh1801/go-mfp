// MFP - Miulti-Function Printers and scanners toolkit
// XML mini library
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// XML decoder

package xml

import (
	"encoding/xml"
	"io"
	"strings"
)

// Decode parses XML document, and represents it as a linear
// sequence of XML elements
//
// Each element has a Name, Path, which is a full path to the element,
// starting from root and Text, which is XML element body, stripped
// from leading and trailing space, and Children, which includes
// its direct children, children of children and so on.
//
// Path uses '/' character as a path separator and starts with '/'
//
// Decoded elements are ordered according to their layout in the
// input document (i.e., root, then first child of the root, then
// its first child and so on), and [Element.Children] are ordered
// the same way.
//
// Namespace prefixes are rewritten according to the 'ns' map.
// Full namespace URL used as map index, and value that corresponds
// to the index replaced with map value. If URL is not found in the
// map, prefix replaced with "-" string
func Decode(ns Namespace, in io.Reader) (Element, error) {
	var elem Element
	stack := []Element{}
	decoder := xml.NewDecoder(in)

	for {
		token, err := decoder.Token()
		if err != nil {
			return Element{}, err
		}

		switch t := token.(type) {
		case xml.StartElement:
			// Decode name and path.
			// Namespace translation is handled here.
			var name string
			if t.Name.Space != "" {
				var ok bool
				name, ok = ns.ByURL(t.Name.Space)
				if !ok {
					name = "-"
				}
			}

			if name != "" {
				name += ":"
			}
			name += t.Name.Local

			// Create an element
			stack = append(stack, elem)
			elem = Element{Name: name}

		case xml.EndElement:
			elem.Text = strings.TrimSpace(elem.Text)

			if len(stack) == 1 {
				return elem, nil
			}

			parent := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			parent.Children = append(parent.Children, elem)
			elem = parent

		case xml.CharData:
			elem.Text += string(t)
		}
	}
}
