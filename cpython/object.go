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
import "runtime"

// Object represents a Python value.
//
// Objects lifetime is managed by the Go garbage collector.
// There is no need to explicitly release the Objects.
type Object struct {
	py     *Python // Interpreter that owns the Object
	oid    objid   // Object ID of the underlying *C.PyObject
	native any     // Native Go value (may be *Object itself)
}

// newObjectFromPython constructs new Object, decoded from PyObject.
// If native is nil, it means, native Go value not available for
// the object. Python None passed as pyNone.
func newObjectFromPython(py *Python, oid objid, native any) *Object {
	obj := &Object{py: py, oid: oid, native: native}

	switch native {
	case nil:
		obj.native = obj
	case pyNone:
		obj.native = nil
	}

	runtime.SetFinalizer(obj, func(obj *Object) {
		obj.finalizer()
	})

	return obj
}

// finalizer is called when Object is garbage-collected.
// It released *C.PyObject, associated with the Object.
func (obj *Object) finalizer() {
	if !obj.py.closed() {
		gate := obj.py.gate()
		defer gate.release()

		obj.py.delObjID(gate, obj.oid)
	}
}

// DelAttr deletes Object attribute with the specified name.
// It returns:
//   - (true, nil) if attribute was found and deleted
//   - (false, nil) if attribute was not found
//   - (false, error) in a case of error
func (obj *Object) DelAttr(name string) (bool, error) {
	gate := obj.py.gate()
	defer gate.release()

	pyobj := obj.py.lookupObjID(gate, obj.oid)

	found, ok := gate.hasattr(pyobj, name)
	if found && ok {
		ok = gate.delattr(pyobj, name)
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
	pyobj := obj.py.lookupObjID(gate, obj.oid)

	found, ok := gate.hasattr(pyobj, name)
	if !found && ok {
		return nil, nil
	}

	// Try to get
	var pyattr pyObject
	if ok {
		pyattr, ok = gate.getattr(pyobj, name)
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
	attrid := obj.py.newObjID(gate, pyattr)
	attr := newObjectFromPython(obj.py, attrid, native)

	return attr, nil
}

// HasAttr reports if Object has the attribute with the specified name.
// This is equivalent to the Python expression hasattr(obj, name).
func (obj *Object) HasAttr(name string) (bool, error) {
	gate := obj.py.gate()
	defer gate.release()

	pyobj := obj.py.lookupObjID(gate, obj.oid)

	has, ok := gate.hasattr(pyobj, name)
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

	pyobj := obj.py.lookupObjID(gate, obj.oid)
	valobj := obj.py.lookupObjID(gate, val.oid)

	ok := gate.setattr(pyobj, name, valobj)
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
