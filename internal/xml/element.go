// MFP - Miulti-Function Printers and scanners toolkit
// XML mini library
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// XML element

package xml

import (
	"os"
	"slices"
)

// Element represents a single decoded XML Element
type Element struct {
	Name     string    // Name of this element (ns:name)
	Text     string    // Element body
	Attrs    []Attr    // Element attributes
	Children []Element // All children
}

// Attr represents an XML element attribute
type Attr struct {
	Name  string // Attribute name
	Value string // Attribute value
}

// Expand replaces ${var} or $var in the XML [Element.Text] and
// [Attr.Value] of the entire XML tree based on the mapping
// function and returns rewritten XML tree.
//
// It is uses [os.Expand] function for strings substitution.
func (root Element) Expand(mapping func(string) string) Element {
	root.Text = os.Expand(root.Text, mapping)
	root.Attrs = slices.Clone(root.Attrs)
	root.Children = slices.Clone(root.Children)

	for i := range root.Attrs {
		attr := &root.Attrs[i]
		attr.Value = os.Expand(attr.Value, mapping)
	}

	for i := range root.Children {
		root.Children[i] = root.Children[i].Expand(mapping)
	}

	return root
}
