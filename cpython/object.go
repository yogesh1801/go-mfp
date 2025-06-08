// MFP - Miulti-Function Printers and scanners toolkit
// CPython binding.
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Python objects

package cpython

// #include "cpython.h"
import "C"

// Object represents a Python value
type Object struct {
	interp pyInterp // Interpreter that owns the Object
	pyobj  pyObject // Underlying *C.PyObject
	native any      // Native Go value (may be *Object)
}

// newObjectFromPython constructs new Object, decoded from PyObject.
// If nativeOk is false, the Object't native value becomes reference
// to the object itself.
func newObjectFromPython(interp pyInterp, pyobj pyObject,
	native any, nativeOk bool) *Object {

	obj := &Object{interp: interp, pyobj: pyobj}
	obj.native = obj

	if nativeOk {
		obj.native = native
	}

	return obj
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
