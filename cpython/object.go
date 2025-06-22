// MFP - Miulti-Function Printers and scanners toolkit
// CPython binding.
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Python objects

package cpython

import (
	"math/big"
	"runtime"
)

// Object represents a Python value.
//
// Objects lifetime is managed by the Go garbage collector.
// There is no need to explicitly release the Objects.
type Object struct {
	py  *Python // Interpreter that owns the Object
	oid objid   // Object ID of the underlying *C.PyObject
}

// newObjectFromPython constructs new Object, decoded from PyObject.
//
// If pyObject is nil, returned *Object is nil as well. This is
// convenient for functions like dictionary retrieval.
func newObjectFromPython(py *Python, gate pyGate, pyobj pyObject) *Object {
	if pyobj == nil {
		return nil
	}

	obj := &Object{
		py:  py,
		oid: py.newObjID(gate, pyobj),
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

// Py returns the [Python] interpreter Object belongs to.
func (obj *Object) Py() *Python {
	return obj.py
}

// Len returns Object length, in items. It works with container
// objects (lists, tuples, dict, ...).
//
// In Python:
//
//	len(obj)
func (obj *Object) Len() (int, error) {
	return objDo(obj, pyGate.length)
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
	found, err := gate.hasitem(pyobj, pykey)
	if found {
		err = gate.delitem(pyobj, pykey)
	}

	return found, err
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
	found, err := gate.hasitem(pyobj, pykey)
	var pyitem pyObject
	if found {
		pyitem, err = gate.getitem(pyobj, pykey)
	}

	if err != nil {
		return nil, err
	}

	// Create the item object
	return newObjectFromPython(obj.py, gate, pyitem), nil
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
	return gate.hasitem(pyobj, pykey)
}

// Set sets Object item with the specified name.
//
// In Python:
//
//	obj[name] = val
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
	return gate.setitem(pyobj, pykey, pyval)
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

	// Check for attribute existence, then delete
	pyobj := obj.py.lookupObjID(gate, obj.oid)
	found, err := gate.hasattr(pyobj, name)
	if found {
		err = gate.delattr(pyobj, name)
	}

	return found, err
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
	found, err := gate.hasattr(pyobj, name)
	if !found {
		return nil, err
	}

	// Try to get
	pyattr, err := gate.getattr(pyobj, name)
	if err != nil {
		return nil, err
	}

	// Create the attribute object
	return newObjectFromPython(obj.py, gate, pyattr), nil
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
	return gate.hasattr(pyobj, name)
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

	// Obtain *C.PyIbject references for all relevant objects.
	pyobj := obj.py.lookupObjID(gate, obj.oid)

	pyval, err := obj.py.newPyObject(gate, val)
	if err != nil {
		return err
	}
	defer gate.unref(pyval)

	// Set the attribute
	return gate.setattr(pyobj, name, pyval)
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
	pyargs, err := gate.makeTuple(len(args))
	if err != nil {
		return nil, err
	}

	defer gate.unref(pyargs)

	for i, arg := range args {
		pyarg, err := obj.py.newPyObject(gate, arg)
		if err != nil {
			return nil, err
		}

		err = gate.setTupleItem(pyargs, pyarg, i)
		if err != nil {
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

	pyret, err := gate.call(pyobj, pyargs, pykwargs)
	if err != nil {
		return nil, err
	}

	// Create response Object
	return newObjectFromPython(obj.py, gate, pyret), nil
}

// Str returns string representation of the Object.
// This is the equivalent of the Python expression str(x).
func (obj *Object) Str() (string, error) {
	return objDo(obj, pyGate.str)
}

// Repr returns string representation of the Object.
// This is the equivalent of the Python expression repr(x).
func (obj *Object) Repr() (string, error) {
	return objDo(obj, pyGate.repr)
}

// Bigint returns Object value as *[big.Int] or an error.
func (obj *Object) Bigint() (*big.Int, error) {
	return objDo(obj, pyGate.decodeBigint)
}

// Bool returns Object value as bool or an error.
func (obj *Object) Bool() (bool, error) {
	return objDo(obj, pyGate.decodeBool)
}

// Bytes returns Object value as []byte slice or an error.
func (obj *Object) Bytes() ([]byte, error) {
	return objDo(obj, pyGate.decodeBytes)
}

// Complex returns Object value as complex128 number or an error.
func (obj *Object) Complex() (complex128, error) {
	return objDo(obj, pyGate.decodeComplex)
}

// Float returns Object value as float64 number or an error.
func (obj *Object) Float() (float64, error) {
	return objDo(obj, pyGate.decodeFloat)
}

// Int returns Object value as int64 number or an error.
func (obj *Object) Int() (int64, error) {
	return objDo(obj, pyGate.decodeInt64)
}

// Keys returns Object mapping keys as the []*Object slice.
// The Object must support mapping.
func (obj *Object) Keys() ([]*Object, error) {
	gate := obj.py.gate()
	defer gate.release()

	// Obtain keys as a list (or tuple)
	pyobj := obj.py.lookupObjID(gate, obj.oid)
	pykeys, err := gate.keys(pyobj)
	if err != nil {
		return nil, err
	}

	defer gate.unref(pykeys)

	// Convert into []*Object slice
	return objSlice(obj.py, gate, pykeys)
}

// Slice returns Object value as []*Object slice or an error.
// It works with sequence objects (lists, tuples, ...).
func (obj *Object) Slice() ([]*Object, error) {
	gate := obj.py.gate()
	defer gate.release()

	pyobj := obj.py.lookupObjID(gate, obj.oid)
	return objSlice(obj.py, gate, pyobj)
}

// Uint returns Object value as uint64 number or an error.
func (obj *Object) Uint() (uint64, error) {
	return objDo(obj, pyGate.decodeUint64)
}

// Unicode returns Object value as UNICODE string or an error.
func (obj *Object) Unicode() (string, error) {
	return objDo(obj, pyGate.decodeUnicode)
}

// IsCallable reports if Object is callable.
func (obj *Object) IsCallable() bool {
	gate := obj.py.gate()
	defer gate.release()

	pyobj := obj.py.lookupObjID(gate, obj.oid)
	return gate.callable(pyobj)
}

// IsMap reports if Object is map (i.e., dict, ...),
func (obj *Object) IsMap() bool {
	return objDoNoError(obj, pyGate.isMap)
}

// IsNone reports if Object is Python None.
func (obj *Object) IsNone() bool {
	return objDoNoError(obj, pyGate.isNone)
}

// IsSeq reports if Object is sequnce (i.e., list, tuple, ...).
func (obj *Object) IsSeq() bool {
	return objDoNoError(obj, pyGate.isSeq)
}

// objDo is the convenience wrapper for the pyGate methods
// with the following signature:
//
//	pyGate.method(pyObject) (T, error)
//
// It acquires the pyGate for the Python interpreter that owns
// the Object, calls the method, releases the gate and return
// result.
func objDo[T any](obj *Object, f func(pyGate, pyObject) (T, error)) (T, error) {
	gate := obj.py.gate()
	defer gate.release()

	pyobj := obj.py.lookupObjID(gate, obj.oid)
	return f(gate, pyobj)
}

// objDo is the convenience wrapper for the pyGate methods
// with the following signature:
//
//	pyGate.method(pyObject) T
//
// It acquires the pyGate for the Python interpreter that owns
// the Object, calls the method, releases the gate and return
// result.
func objDoNoError[T any](obj *Object, f func(pyGate, pyObject) T) T {
	gate := obj.py.gate()
	defer gate.release()

	pyobj := obj.py.lookupObjID(gate, obj.oid)
	return f(gate, pyobj)
}

// objSlice is the helper function that extracts contained objects from the
// sequence object and converts extracted objects into the []*Object slice.
func objSlice(py *Python, gate pyGate, pyobj pyObject) ([]*Object, error) {
	// Check that object is sequence
	if !gate.isSeq(pyobj) {
		return nil, gate.decodeError(pyobj, "[]*Object")
	}

	// Obtain length
	length, err := gate.length(pyobj)
	if err != nil {
		return nil, err
	}

	// Obtain items
	pyobjects := make([]pyObject, length)
	for i := range pyobjects {
		pyobjects[i], err = gate.getSeqItem(pyobj, i)
		if err != nil {
			for j := 0; j < i; j++ {
				gate.unref(pyobjects[j])
			}
			return nil, err
		}
	}

	// Convert into []*Object
	objects := make([]*Object, length)
	for i := range pyobjects {
		objects[i] = newObjectFromPython(py, gate, pyobjects[i])
	}

	return objects, nil
}
