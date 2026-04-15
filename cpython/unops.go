// MFP - Miulti-Function Printers and scanners toolkit
// CPython binding.
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Unary operations on objects

package cpython

// Neg negates the Object:
//
//	ret = -obj
func (obj *Object) Neg() *Object {
	return obj.unop("__neg__")
}

// Pos applies the unary '+' operation to the Object:
//
//	ret = +obj
func (obj *Object) Pos() *Object {
	return obj.unop("__pos__")
}

// Invert inverts the Object:
//
//	ret = ~obj
func (obj *Object) Invert() *Object {
	return obj.unop("__invert__")
}

// binop performs the named unary operation on the object.
func (obj *Object) unop(name string) *Object {
	// Acquire the gate and obtain the first object's *C.PyObject
	gate, pyobj, err := obj.begin()
	if err != nil {
		return newErrorObject(obj.py, err)
	}
	defer gate.release()

	// Obtain callable *C.PyObject for the operation
	pyop, err := gate.getattr(pyobj, name)
	if err != nil {
		return newErrorObject(obj.py, err)
	}
	defer gate.unref(pyop)

	// Construct arguments
	pyargs, err := gate.makeTuple(0)
	if err != nil {
		return newErrorObject(obj.py, err)
	}
	defer gate.unref(pyargs)

	// Call the operation
	pyres, err := gate.call(pyop, pyargs, nil)
	if err != nil {
		return newErrorObject(obj.py, err)
	}

	// Create the result object
	return newObjectFromPython(obj.py, gate, pyres)
}
