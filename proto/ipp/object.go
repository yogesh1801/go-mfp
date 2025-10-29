// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Object gives access to raw IPP attributes

package ipp

import (
	"github.com/OpenPrinting/goipp"
)

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

	// Get returns [goipp.Attibute] by name
	Get(name string) (goipp.Attribute, bool)

	// Set sets [goipp.Attibute] by name. It updates both
	// object raw attributes and the corresponding field
	// in the object structure (if any).
	Set(attr goipp.Attribute) error
}

// ObjectRawAttrs MUST be embedded into every IPP-encodable structure.
// It gives access to raw IPP attributes and implements [Object]
// interface.
type ObjectRawAttrs struct {
	attrs  goipp.Attributes // Raw attributes
	byName map[string]int   // Attribute indices by name
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

	i, found := rawattrs.byName[name]
	if found {
		attr = rawattrs.attrs[i]
	}

	return
}

// set sets attribute by name and updates the outer structure that
// contains the ObjectRawAttrs (hence the need for ippCodec to
// obtain information about the attribute).
func (rawattrs *ObjectRawAttrs) set(attr goipp.Attribute, outer Object) error {
	// Update the outer structure
	dec := ippDecoder{}
	err := dec.DecodeSingle(outer, attr)
	if err != nil {
		return err
	}

	// Update raw attributes
	i, found := rawattrs.byName[attr.Name]
	if !found {
		i = len(rawattrs.attrs)
		rawattrs.byName[attr.Name] = i
		rawattrs.attrs = append(rawattrs.attrs, goipp.Attribute{})
	}

	rawattrs.attrs[i] = attr

	return nil
}

// setattrs saves all raw IPP attributes
func (rawattrs *ObjectRawAttrs) setattrs(attrs goipp.Attributes) {
	rawattrs.attrs = make(goipp.Attributes, 0, len(attrs))
	rawattrs.byName = make(map[string]int, len(attrs))

	for _, attr := range attrs {
		// If we see some attribute, the second occurrence is
		// Silently ignored. Note, CUPS does the same
		//
		// For details, see discussion here:
		//   https://lore.kernel.org/printing-architecture/84EEF38C-152E-4779-B1E8-578D6BB896E6@msweet.org/
		if _, found := rawattrs.byName[attr.Name]; !found {
			rawattrs.byName[attr.Name] = len(attrs)
			rawattrs.attrs = append(rawattrs.attrs, attr)
		}
	}
}
