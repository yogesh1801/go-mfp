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
func newObjectFromPython(py *Python, gate pyGate,
	pyobj pyObject, native any) *Object {

	// Obtain Object ID
	oid := py.newObjID(gate, pyobj)
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

// Del deletes Object item with the specified key:
//
// In Python:
//
//	del(obj, key)
//
// The Object must be container (array, dict, etc).
// The key may be any value that [Python.NewObject] accepts.
//
// It returns:
//   - (true, nil) if item was found and deleted
//   - (false, nil) if item was not found
//   - (false, error) in a case of error
func (obj *Object) Del(key any) (bool, error) {
	gate := obj.py.gate()
	defer gate.release()

	// Obtain *C.PyIbject references for all relevant objects.
	pyobj := obj.py.lookupObjID(gate, obj.oid)

	pykey, err := obj.py.newPyObject(gate, key)
	if err != nil {
		return false, err
	}
	defer gate.unref(pykey)

	// Check for item existence, then delete, if found.
	found, ok := gate.hasitem(pyobj, pykey)
	if found && ok {
		ok = gate.delitem(pyobj, pykey)
	}

	if !ok {
		return false, gate.lastError()
	}

	return found, nil
}

// Get returns Object item with the specified key:
//
// In Python:
//
//	obj[key]
//
// The Object must be container (array, dict, etc).
// The key may be any value that [Python.NewObject] accepts.
//
// It returns:
//   - (*Object, nil) if item was found
//   - (nil, nil) if item was not found
//   - (nil, error) in a case of error
func (obj *Object) Get(key any) (*Object, error) {
	gate := obj.py.gate()
	defer gate.release()

	// Obtain *C.PyIbject references for all relevant objects.
	pyobj := obj.py.lookupObjID(gate, obj.oid)

	pykey, err := obj.py.newPyObject(gate, key)
	if err != nil {
		return nil, err
	}
	defer gate.unref(pykey)

	// Check for item existence, then retrieve, if found.
	found, ok := gate.hasitem(pyobj, pykey)
	if !found && ok {
		return nil, nil
	}

	var pyitem pyObject
	if ok {
		pyitem, ok = gate.getitem(pyobj, pykey)
	}

	// Try to decode
	var native any
	if ok {
		native, ok = gate.decodeObject(pyitem)
	}

	// Handle possible error
	if !ok {
		return nil, gate.lastError()
	}

	// Create the item object
	item := newObjectFromPython(obj.py, gate, pyitem, native)

	return item, nil
}

// Contains reports if Object has the item with the specified key.
//
// In Python:
//
//	key in obj
//
// The Object must be container (array, dict, etc).
// The key may be any value that [Python.NewObject] accepts.
func (obj *Object) Contains(key any) (bool, error) {
	gate := obj.py.gate()
	defer gate.release()

	// Obtain *C.PyIbject references for all relevant objects.
	pyobj := obj.py.lookupObjID(gate, obj.oid)

	pykey, err := obj.py.newPyObject(gate, key)
	if err != nil {
		return false, err
	}
	defer gate.unref(pykey)

	// Check for item existence
	found, ok := gate.hasitem(pyobj, pykey)
	if !ok {
		return false, gate.lastError()
	}
	return found, nil
}

// Set sets Object item with the specified name.
//
// In Python:
//
//	obj.name = val
//
// The Object must be container (array, dict, etc).
// The key and val may be any value that [Python.NewObject] accepts.
func (obj *Object) Set(key, val any) error {
	gate := obj.py.gate()
	defer gate.release()

	// Obtain *C.PyIbject references for all relevant objects.
	pyobj := obj.py.lookupObjID(gate, obj.oid)

	pykey, err := obj.py.newPyObject(gate, key)
	if err != nil {
		return err
	}
	defer gate.unref(pykey)

	pyval, err := obj.py.newPyObject(gate, val)
	if err != nil {
		return err
	}
	defer gate.unref(pyval)

	// Set the item
	ok := gate.setitem(pyobj, pykey, pyval)
	if !ok {
		return gate.lastError()
	}

	return nil
}

// DelAttr deletes Object attribute with the specified name:
//
// In Python:
//
//	delattr(obj, name)
//
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

// GetAttr returns Object attribute with the specified name:
//
// In Python:
//
//	obj.name
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
	attr := newObjectFromPython(obj.py, gate, pyattr, native)

	return attr, nil
}

// HasAttr reports if Object has the attribute with the specified name.
//
// In Python:
//
//	hasattr(obj, name)
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
//
// In Python:
//
//	obj.name = val
//
// The val may be any value that [Python.NewObject] accepts.
func (obj *Object) SetAttr(name string, val any) error {
	gate := obj.py.gate()
	defer gate.release()

	pyobj := obj.py.lookupObjID(gate, obj.oid)

	pyval, err := obj.py.newPyObject(gate, val)
	if err != nil {
		return err
	}
	defer gate.unref(pyval)

	ok := gate.setattr(pyobj, name, pyval)
	if !ok {
		return gate.lastError()
	}

	return nil
}

// Call calls Object as function.
//
// Arguments are automatically converted from Go to Python.
// See [Python.NewObject] for details.
//
// Use [Object.CallKW] for call with keyword arguments.
func (obj *Object) Call(args ...any) (*Object, error) {
	return obj.CallKW(nil, args...)
}

// CallKW calls Object as function with keyword arguments defined
// by the kw parameter and positional arguments defined by the
// args parameter (variadic).
//
// Arguments are automatically converted from Go to Python.
// See [Python.NewObject] for details.
//
// If keyword arguments are not used, kw may be nil.
//
// It returns the function's return value.
func (obj *Object) CallKW(kw map[string]any, args ...any) (*Object, error) {
	gate := obj.py.gate()
	defer gate.release()

	// Convert positional arguments
	pyargs := gate.makeTuple(len(args))
	if pyargs == nil {
		return nil, gate.lastError()
	}

	defer gate.unref(pyargs)

	for i, arg := range args {
		pyarg, err := obj.py.newPyObject(gate, arg)
		if err != nil {
			return nil, err
		}

		ok := gate.setTupleItem(pyargs, pyarg, i)
		if !ok {
			err := gate.lastError()
			gate.unref(pyarg)
			return nil, err
		}
	}

	// Convert keyword arguments
	var pykwargs pyObject
	if len(kw) > 0 {
		var err error
		pykwargs, err = obj.py.newPyObject(gate, kw)
		if err != nil {
			return nil, err
		}

		defer gate.unref(pykwargs)
	}

	// Perform a call
	pyobj := obj.py.lookupObjID(gate, obj.oid)
	pyret := gate.call(pyobj, pyargs, pykwargs)

	if pyret == nil {
		return nil, gate.lastError()
	}

	// Decode response
	native, ok := gate.decodeObject(pyret)
	if !ok {
		return nil, gate.lastError()
	}

	ret := newObjectFromPython(obj.py, gate, pyret, native)

	return ret, nil
}

// Callable reports if Object is callable.
func (obj *Object) Callable() (bool, error) {
	gate := obj.py.gate()
	defer gate.release()

	pyobj := obj.py.lookupObjID(gate, obj.oid)
	return gate.callable(pyobj), nil
}

// Str returns string representation of the Object.
// This is the equivalent of the Python expression str(x).
func (obj *Object) Str() (string, error) {
	gate := obj.py.gate()
	defer gate.release()

	pyobj := obj.py.lookupObjID(gate, obj.oid)
	s, ok := gate.str(pyobj)
	var err error
	if !ok {
		err = gate.lastError()
	}

	return s, err
}

// Repr returns string representation of the Object.
// This is the equivalent of the Python expression repr(x).
func (obj *Object) Repr() (string, error) {
	gate := obj.py.gate()
	defer gate.release()

	pyobj := obj.py.lookupObjID(gate, obj.oid)
	s, ok := gate.repr(pyobj)
	var err error
	if !ok {
		err = gate.lastError()
	}

	return s, err
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
