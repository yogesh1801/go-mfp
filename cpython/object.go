// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Python objects

package cpython

// #include "cpython.h"
import "C"

// pyObject is the Go name for the *C.PyObject
type pyObject = *C.PyObject

// pyObject is the Go name for the *C.PyObject
type pyTypeObject = *C.PyTypeObject

// Object represents a Python value
type Object struct {
	interp pyInterp // Interpreter that owns the Object
	pyobj  pyObject // Underlying *C.PyObject
	native any      // Native Go value (may be *Object)
}

// Unbox returns Object's value as Go value.
//
// If Object cannot be represented as a native Go value,
// the Object itself is returned.
func (obj *Object) Unbox() any {
	return obj.native
}

// pyObjectType returns pyTypeObject for the value object
func pyObjectType(ob pyObject) pyTypeObject {
	return C.Py_TYPE(ob)
}
