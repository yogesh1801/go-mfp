// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Object gives access to raw IPP attributes

package ipp

import "github.com/OpenPrinting/goipp"

// Object is the interface that MUST be implemented by any
// concrete type that can be encoded to and decoded from the
// set of IPP attributes.
//
// It gives access to the underlying raw IPP attributes.
//
// Every concrete IPP-encodable structure MUST embed ObjectRawAttrs
// to implement this interface.
type Object interface {
	// RawAttrs returns ObjectRawAttrs embedded into the structure
	RawAttrs() *ObjectRawAttrs

	// KnownAttributes return information about known
	// attributes of the Object
	KnownAttrs() []AttrInfo
}

// ObjectRawAttrs MUST be embedded into every IPP-encodable structure.
// It gives access to raw IPP attributes and implements [Object]
// interface.
type ObjectRawAttrs struct {
	attrs  goipp.Attributes
	byName map[string]goipp.Attribute
}

// RawAttrs returns [ObjecRawtAttrs], which gives uniform
// access to the header of any [Object]
func (rawattrs *ObjectRawAttrs) RawAttrs() *ObjectRawAttrs {
	return rawattrs
}

// All returns all [goipp.Attributes].
func (rawattrs *ObjectRawAttrs) All() goipp.Attributes {
	return rawattrs.attrs
}

// Get returns [goipp.Attribute] by name.
func (rawattrs *ObjectRawAttrs) Get(name string) (
	attr goipp.Attribute, found bool) {

	attr, found = rawattrs.byName[name]
	return
}

// set saves raw IPP attributes
func (rawattrs *ObjectRawAttrs) set(attrs goipp.Attributes) {
	rawattrs.attrs = make(goipp.Attributes, 0, len(attrs))
	rawattrs.byName = make(map[string]goipp.Attribute, len(attrs))

	for _, attr := range attrs {
		// If we see some attribute, the second occurrence is
		// Silently ignored. Note, CUPS does the same
		//
		// For details, see discussion here:
		//   https://lore.kernel.org/printing-architecture/84EEF38C-152E-4779-B1E8-578D6BB896E6@msweet.org/
		if _, found := rawattrs.byName[attr.Name]; !found {
			rawattrs.byName[attr.Name] = attr
			rawattrs.attrs = append(rawattrs.attrs, attr)
		}
	}
}
