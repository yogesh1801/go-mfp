// MFP - Miulti-Function Printers and scanners toolkit
// CPython binding.
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Binary operations on objects

package cpython

// Add adds value to Object:
//
//	ret = obj + val
func (obj *Object) Add(val any) *Object {
	return obj.binop("__add__", val)
}

// Sub subtracts value from Object:
//
//	ret = obj - val
func (obj *Object) Sub(val any) *Object {
	return obj.binop("__sub__", val)
}

// Mul multiplies Object by value:
//
//	ret = obj * val
func (obj *Object) Mul(val any) *Object {
	return obj.binop("__mul__", val)
}

// TrueDiv divides Object to value, returning a float:
//
//	ret = obj / val
func (obj *Object) TrueDiv(val any) *Object {
	return obj.binop("__truediv__", val)
}

// FloorDiv divides Object to value and returns integer part of the quotient:
//
//	ret = obj // val
func (obj *Object) FloorDiv(val any) *Object {
	return obj.binop("__floordiv__", val)
}

// Mod computes the remainder of Object divided by val:
//
//	ret = obj % val
func (obj *Object) Mod(val any) *Object {
	return obj.binop("__mod__", val)
}

// Pow returns Object raised to the power of val:
//
//	ret = obj ** val
func (obj *Object) Pow(val any) *Object {
	return obj.binop("__pow__", val)
}

// Lt tells if Object is less that val:
//
//	ret = obj < val
func (obj *Object) Lt(val any) *Object {
	return obj.binop("__lt__", val)
}

// Gt tells if Object is greater that val:
//
//	ret = obj > val
func (obj *Object) Gt(val any) *Object {
	return obj.binop("__gt__", val)
}

// Le tells if Object is less or equal that val:
//
//	ret = obj <= val
func (obj *Object) Le(val any) *Object {
	return obj.binop("__le__", val)
}

// Ge tells if Object is less or equal that val:
//
//	ret = obj >= val
func (obj *Object) Ge(val any) *Object {
	return obj.binop("__ge__", val)
}

// Eq tells if Object is equal to val:
//
//	ret = obj == val
func (obj *Object) Eq(val any) *Object {
	return obj.binop("__eq__", val)
}

// Ne tells if Object is not equal to val:
//
//	ret = obj != val
func (obj *Object) Ne(val any) *Object {
	return obj.binop("__ne__", val)
}

// binop performs the named binary operation on two objects.
func (obj *Object) binop(name string, val any) *Object {
	// Acquire the gate and obtain the first *C.PyObject
	gate, pyobj, err := obj.begin()
	if err != nil {
		return newErrorObject(obj.py, err)
	}
	defer gate.release()

	// Obtain the second *C.PyObject
	pyobj2, err := obj.py.newPyObject(gate, val)
	if err != nil {
		return newErrorObject(obj.py, err)
	}
	defer gate.unref(pyobj2)

	// Obtain callable *C.PyObject for the operation
	pyop, err := gate.getattr(pyobj, name)
	if err != nil {
		return newErrorObject(obj.py, err)
	}
	defer gate.unref(pyop)

	// Construct arguments
	pyargs, err := gate.makeTuple(1)
	if err != nil {
		return newErrorObject(obj.py, err)
	}
	defer gate.unref(pyargs)

	err = gate.setTupleItem(pyargs, pyobj2, 0)
	if err != nil {
		return newErrorObject(obj.py, err)
	}

	// Call the operation
	pyres, err := gate.call(pyop, pyargs, nil)
	if err != nil {
		return newErrorObject(obj.py, err)
	}

	// Create the result object
	return newObjectFromPython(obj.py, gate, pyres)
}
