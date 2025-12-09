// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Helper functions for reflection

package ipp

import (
	"reflect"

	"github.com/OpenPrinting/go-mfp/proto/ipp/iana"
)

// reflectIsObject reports if type implements the [Object] interface.
func reflectIsObject(t reflect.Type) bool {
	var obj Object
	objtype := reflect.TypeOf(&obj).Elem()

	return reflect.PointerTo(t).AssignableTo(objtype)
}

// reflecRegistrations extracts attribute registrations from
// the Object of type t.
//
// t.Kind must be reflect.Struct.
//
// If t doesn't implement the Object interface, this function
// returns nil.
func reflecRegistrations(t reflect.Type) []map[string]*iana.DefAttr {
	if t.Kind() != reflect.Struct || !reflectIsObject(t) {
		return nil
	}

	regs := []map[string]*iana.DefAttr{}

	n := t.NumField()
	for i := 0; i < n; i++ {
		fld := t.Field(i)
		if !fld.IsExported() {
			continue
		}

		if grp, ok := reflect.New(fld.Type).Interface().(attributesGroup); ok {
			regs = append(regs, grp.registrations())
		}
	}

	return regs
}
