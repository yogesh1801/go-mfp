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
	Path     string     // Full path to the Element within XML document
	Parent   *Element   // Parent element, nil for root
	Children []*Element // All children
}
