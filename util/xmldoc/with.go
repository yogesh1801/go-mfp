// MFP - Miulti-Function Printers and scanners toolkit
// XML mini library
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// XML tree constructors

package xmldoc

// WithText is the convenience function for constricting XML tree.
//
// It construct Element with Text.
func WithText(name, text string) Element {
	return Element{Name: name, Text: text}
}

// WithChildren is the convenience function for constricting XML tree.
//
// It construct Element with Children.
func WithChildren(name string, children ...Element) Element {
	return Element{Name: name, Children: children}
}

// WithAttrs is the convenience function for constricting XML tree.
//
// It construct Element with attributes.
func WithAttrs(name string, attrs ...Attr) Element {
	return Element{Name: name, Attrs: attrs}
}
