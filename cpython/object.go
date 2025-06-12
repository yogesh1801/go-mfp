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
	py     *Python  // Interpreter that owns the Object
	pyobj  pyObject // Underlying *C.PyObject
	native any      // Native Go value (may be *Object)
}

// newObjectFromPython constructs new Object, decoded from PyObject.
// If native is nil, it means, native Go value not available for
// the object. Python None passed as pyNone.
func newObjectFromPython(py *Python, pyobj pyObject, native any) *Object {
	obj := &Object{py: py, pyobj: pyobj, native: native}

	switch native {
	case nil:
		obj.native = obj
	case pyNone:
		obj.native = nil
	}

	return obj
}

// Unref decrements Object's reference count.
// Object should not be accessed after that.
func (obj *Object) Unref() {
	gate := obj.py.gate()
	defer gate.release()

	gate.unref(obj.pyobj)
	obj.pyobj = nil
}

// DelAttr deletes Object attribute with the specified name.
// It returns:
//   - (true, nil) if attribute was found and deleted
//   - (false, nil) if attribute was not found
//   - (false, error) in a case of error
func (obj *Object) DelAttr(name string) (bool, error) {
	gate := obj.py.gate()
	defer gate.release()

	found, ok := gate.hasattr(obj.pyobj, name)
	if found && ok {
		ok = gate.delattr(obj.pyobj, name)
	}

	if !ok {
		return false, gate.lastError()
	}

	return found, nil
}

// GetAttr returns Object attribute with the specified name.
// This is the equivalent of the Python expression obj.name
//
// It returns:
//   - (*Object, nil) if attribute was found
//   - (nil, nil) if attribute was not found
//   - (nil, error) in a case of error
func (obj *Object) GetAttr(name string) (*Object, error) {
	gate := obj.py.gate()
	defer gate.release()

	// Check if attribute exists
	found, ok := gate.hasattr(obj.pyobj, name)
	if !found && ok {
		return nil, nil
	}

	// Try to get
	var pyattr pyObject
	if ok {
		pyattr, ok = gate.getattr(obj.pyobj, name)
	}

	// Try to decode
	var native any
	if ok {
		native, ok = gate.decodeObject(pyattr)
	}

	// Handle possible error
	if !ok {
		return nil, gate.lastError()
	}

	// Create the attribute object
	attr := newObjectFromPython(obj.py, pyattr, native)
	return attr, nil
}

// HasAttr reports if Object has the attribute with the specified name.
// This is equivalent to the Python expression hasattr(obj, name).
func (obj *Object) HasAttr(name string) (bool, error) {
	gate := obj.py.gate()
	defer gate.release()

	has, ok := gate.hasattr(obj.pyobj, name)
	if !ok {
		return false, gate.lastError()
	}
	return has, nil
}

// SetAttr sets Object attribute with the specified name.
// This is the equivalent of the Python statement obj.name = val
func (obj *Object) SetAttr(name string, val *Object) error {
	gate := obj.py.gate()
	defer gate.release()

	ok := gate.setattr(obj.pyobj, name, val.pyobj)
	if !ok {
		return gate.lastError()
	}

	return nil
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
