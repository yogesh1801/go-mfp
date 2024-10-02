// MFP - Miulti-Function Printers and scanners toolkit
// XML mini library
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// XML element

package xml

// Element represents a single decoded XML Element
type Element struct {
	Name     string     // Name of this element
	Text     string     // Element body
	Attrs    []Attr     // Element attributes
	Children []*Element // All children
	Parent   *Element   // Parent element, nil for root
	Path     string     // Full path to the Element within XML document
}

// Attr represents an XML element attribute
type Attr struct {
	Name  string // Attribute name
	Value string // Attribute value
}
