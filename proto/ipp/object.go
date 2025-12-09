// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Object gives access to raw IPP attributes

package ipp

import (
	"reflect"

	"github.com/OpenPrinting/go-mfp/util/generic"
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

	// Errors returns a slice of errors that has occurred during
	// the [Object] decoding.
	//
	// If [DecodeOptions.KeepTrying] set to true, non-fatal errors
	// doesn't interrupt decoding but instead saved here (and may
	// be reported as decode warnings).
	Errors() []error

	// This is to force embedding of some attributesGroup into
	// the object
	//attributesGroup
}

// ObjectRegisteredAttrNames returns names of attributes specific
// for the particular [Object] type (but not necessarily present
// in the particular Object instance).
func ObjectRegisteredAttrNames(obj Object) []string {
	return ippRegisteredAttrNames(reflect.TypeOf(obj))
}

// ObjectGetAttr returns [goipp.Attibute] by name
func ObjectGetAttr(obj Object, name string) (attr goipp.Attribute, found bool) {
	rawattrs := obj.RawAttrs()

	i, found := rawattrs.byName[name]
	if found {
		attr = rawattrs.attrs[i]
	}

	return
}

// ObjectSetAttr sets [goipp.Attibute] by name. It updates both
// object raw attributes and the corresponding field in the
// object structure (if any).
func ObjectSetAttr(obj Object, attr goipp.Attribute) error {
	// Update the Object's outer structure
	dec := ippDecoder{}
	err := dec.DecodeSingle(obj, attr)
	if err != nil {
		return err
	}

	// Update raw attributes
	rawattrs := obj.RawAttrs()
	i, found := rawattrs.byName[attr.Name]
	if !found {
		i = len(rawattrs.attrs)
		rawattrs.byName[attr.Name] = i
		rawattrs.attrs = append(rawattrs.attrs, goipp.Attribute{})
	}

	rawattrs.attrs[i] = attr

	return nil
}

// ObjectRawAttrs MUST be embedded into every IPP-encodable structure.
// It gives access to raw IPP attributes and implements [Object]
// interface.
type ObjectRawAttrs struct {
	attrs  goipp.Attributes // Raw attributes
	byName map[string]int   // Attribute indices by name
	errors []error          // Possible decode errors
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

// Errors returns a slice of errors that has occurred during
// the [Object] decoding.
//
// If [DecodeOptions.KeepTrying] set to true, non-fatal errors
// doesn't interrupt decoding but instead saved here (and may
// be reported as decode warnings).
func (rawattrs *ObjectRawAttrs) Errors() []error {
	return rawattrs.errors
}

// save saves all raw IPP attributes and decode errors.
func (rawattrs *ObjectRawAttrs) save(attrs goipp.Attributes, errors []error) {
	rawattrs.attrs = make(goipp.Attributes, 0, len(attrs))
	rawattrs.byName = make(map[string]int, len(attrs))
	rawattrs.errors = generic.CopySlice(errors)

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
