// MFP - Miulti-Function Printers and scanners toolkit
// XML mini library
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// XML elements lookup

package xml

// Lookup contains element name for lookup by name and
// received lookup result.
//
// It is optimized for looking up multiple elements at once.
//
// See also: [Element.Lookup]
type Lookup struct {
	Name  string  // Requested element name
	Elem  Element // Returned element data
	Found bool    // Becomes true, if element was found
}
