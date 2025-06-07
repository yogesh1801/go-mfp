// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Python objects

package cpython

import (
	"unsafe"

	"github.com/OpenPrinting/go-mfp/internal/assert"
)

// #include "cpython.h"
import "C"

// pyObject is the Go name for the *C.PyObject
type pyObject = *C.PyObject

// pyObject is the Go name for the *C.PyObject
type pyTypeObject = *C.PyTypeObject

// Object represents a Python value
type Object struct {
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

// objectFromPython decodes Object from *C.PyObject
func objectFromPython(pyobj pyObject) *Object {
	// Translate nil to nil
	if pyobj == nil {
		return nil
	}

	// Construct the Object
	obj := &Object{pyobj: pyobj}
	obj.native = obj

	// Decode native value, if possible.
	switch pyObjectType(pyobj) {
	case C.PyBool_Type_p:
		obj.native = C.py_obj_is_true(pyobj) != 0
	case C.PyByteArray_Type_p:
	case C.PyBytes_Type_p:
	case C.PyCFunction_Type_p:
	case C.PyComplex_Type_p:
	case C.PyDict_Type_p:
	case C.PyDictKeys_Type_p:
	case C.PyFloat_Type_p:
	case C.PyFrozenSet_Type_p:
	case C.PyList_Type_p:
	case C.PyLong_Type_p:
	case C.PyMemoryView_Type_p:
	case C.PyModule_Type_p:
	case C.PySet_Type_p:
	case C.PySlice_Type_p:
	case C.PyTuple_Type_p:
	case C.PyType_Type_p:
	case C.PyUnicode_Type_p:
		sz := C.py_str_len(pyobj)
		assert.Must(sz >= 0)
		buf := make([]rune, sz)
		p := (*C.Py_UCS4)(unsafe.Pointer(&buf[0]))
		C.py_str_get(pyobj, p, C.size_t(sz))
		obj.native = string(buf)
	default:
		if C.py_obj_is_none(pyobj) != 0 {
			obj.native = nil
		}
	}

	return obj
}

// pyObjectType returns pyTypeObject for the value object
func pyObjectType(ob pyObject) pyTypeObject {
	return C.Py_TYPE(ob)
}
