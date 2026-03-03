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
// It understands attribute registrations defined by including
// of attributesGroup as well as attribute registrations, defined
// inline, using the appropriate syntax of the ipp: struct tags.
//
// t.Kind must be reflect.Struct.
func reflecRegistrations(t reflect.Type) []map[string]*iana.DefAttr {
	switch t.Kind() {
	case reflect.Slice, reflect.Array, reflect.Pointer:
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return nil
	}

	regs := []map[string]*iana.DefAttr{}
	defs := make(map[string]*iana.DefAttr)

	for _, fld := range reflect.VisibleFields(t) {
		if !fld.IsExported() {
			continue
		}

		if grp, ok := reflect.New(fld.Type).Interface().(attributesGroup); ok {
			// Handle attributesGroup
			regs = append(regs, grp.registrations())
		} else if name, def, _ := attrFieldAnalyze(fld); def != nil {
			// Handle ipp: tag, if any
			if defs[name] == nil {
				defs[name] = def
			}
		}
	}

	if len(defs) != 0 {
		// Prepend inline registrations, if any
		regs = append([]map[string]*iana.DefAttr{defs}, regs...)
	}

	return regs
}
