// MFP - Miulti-Function Printers and scanners toolkit
// CPython binding.
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Call gate into the Python interpreter

package cpython

import (
	"math/big"
	"runtime"
	"unsafe"

	"github.com/OpenPrinting/go-mfp/internal/assert"
)

// #include "cpython.h"
import "C"

// pyGate represents the locked (attached to the current thread
// and with the GIL acquired) state of the Python interpreter.
//
// It works as a call gate into the interpreter and implements
// all interpreter operations that require locking.
type pyGate struct {
	prev *C.PyThreadState // Previous current thread state
}

// pyGateAcquire temporary attaches the calling thread to the
// Python interpreter.
//
// It returns the pyGate object, that must be released after
// use with the [pyGate.release] call.
func pyGateAcquire(interp pyInterp) pyGate {
	runtime.LockOSThread()
	prev := C.py_enter(interp)
	return pyGate{prev}
}

// release detaches the calling thread from the Python interpreter.
func (gate pyGate) release() {
	C.py_leave(gate.prev)
	runtime.UnlockOSThread()
}

// lastError returns a last error, nil if none.
func (gate pyGate) lastError() error {
	var etype, evalue, trace pyObject

	// Fetch Python error information
	C.py_err_fetch(&etype, &evalue, &trace)
	if etype == nil && evalue == nil && trace == nil {
		return nil
	}

	defer C.py_obj_unref(etype)
	defer C.py_obj_unref(evalue)
	defer C.py_obj_unref(trace)

	// Decode the error
	if evalue != nil {
		msg, ok := gate.str(evalue)
		if ok {
			return ErrPython{msg}
		}
	}

	if etype != nil {
		msg, ok := gate.str(etype)
		if ok {
			return ErrPython{msg}
		}
	}

	return ErrPython{"Unknown Python exception"}
}

// ref increments PyObject's reference count.
func (gate pyGate) ref(pyobj pyObject) {
	C.py_obj_ref(pyobj)
}

// unref decrements PyObject's reference count.
func (gate pyGate) unref(pyobj pyObject) {
	C.py_obj_unref(pyobj)
}

// str returns str(pyobj), decoded as Go string.
func (gate pyGate) str(pyobj pyObject) (s string, ok bool) {
	str := C.py_obj_str(pyobj)
	if str != nil {
		defer C.py_obj_unref(str)
		s, ok = gate.decodeString(str)
	}

	return
}

// repr returns repr(pyobj), decoded as Go string.
func (gate pyGate) repr(pyobj pyObject) (s string, ok bool) {
	repr := C.py_obj_repr(pyobj)
	if repr != nil {
		defer C.py_obj_unref(repr)
		s, ok = gate.decodeString(repr)
	}

	return
}

// delattr deletes Object attribute with the specified name.
func (gate pyGate) delattr(pyobj pyObject, name string) (ok bool) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	ok = bool(C.py_obj_delattr(pyobj, cname))
	return
}

// getattr returns Object attribute with the specified name.
func (gate pyGate) getattr(pyobj pyObject, name string) (attr pyObject, ok bool) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	ok = bool(C.py_obj_getattr(pyobj, cname, &attr))
	return
}

// hasattr reports if Object has attribute with the specified name.
func (gate pyGate) hasattr(pyobj pyObject, name string) (answer, ok bool) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	var canswer C.bool
	ok = bool(C.py_obj_hasattr(pyobj, cname, &canswer))
	answer = bool(canswer)

	return
}

// setattr sets Object attribute with the specified name.
func (gate pyGate) setattr(pyobj pyObject, name string, val pyObject) (ok bool) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	ok = bool(C.py_obj_setattr(pyobj, cname, val))
	return
}

// delitem deletes Object item with the specified key:
//
//	del(pyobj[key])
func (gate pyGate) delitem(pyobj, key pyObject) (ok bool) {
	ok = bool(C.py_obj_delitem(pyobj, key))
	return
}

// getitem returns Object item with the specified key.
//
//	pyobj[key]
func (gate pyGate) getitem(pyobj, key pyObject) (item pyObject, ok bool) {
	ok = bool(C.py_obj_getitem(pyobj, key, &item))
	return
}

// hasitem reports if Object has item with the specified key.
//
//	key in pyobj
func (gate pyGate) hasitem(pyobj, key pyObject) (answer, ok bool) {
	var canswer C.bool
	ok = bool(C.py_obj_hasitem(pyobj, key, &canswer))
	answer = bool(canswer)

	return
}

// setitem sets Object item with the specified key.
//
//	pyobj[key] = val
func (gate pyGate) setitem(pyobj, key, val pyObject) (ok bool) {
	ok = bool(C.py_obj_setitem(pyobj, key, val))
	return
}

// call calls callable object as a function, with positional
// arguments, defined by args (must be PyTuple_Type) and keyword
// arguments, defined by kwargs (must be PyDict_Type or nil).
//
// It returns strong reference to result on success, nil on an error.
func (gate pyGate) call(callable, args, kwargs pyObject) (res pyObject) {
	return C.py_obj_call(callable, args, kwargs)
}

// callable reports if object is callable.
// This function always succeeds.
func (gate pyGate) callable(pyobj pyObject) bool {
	return bool(C.py_obj_callable(pyobj))
}

// decodeObject decodes PyObject value as Go value.
// It returns pyNone for Python None and nil if native Go value not available.
func (gate pyGate) decodeObject(pyobj pyObject) (any, bool) {
	switch pyObjectType(pyobj) {
	case C.PyBool_Type_p:
		return C.py_obj_is_true(pyobj) != 0, true
	case C.PyByteArray_Type_p:
		return gate.decodeByteArray(pyobj)
	case C.PyBytes_Type_p:
		return gate.decodeBytes(pyobj)
	case C.PyCFunction_Type_p:
	case C.PyComplex_Type_p:
		return gate.decodeComplex(pyobj)
	case C.PyDict_Type_p:
	case C.PyDictKeys_Type_p:
	case C.PyFloat_Type_p:
		return gate.decodeFloat(pyobj)
	case C.PyFrozenSet_Type_p:
	case C.PyList_Type_p:
	case C.PyLong_Type_p:
		return gate.decodeInteger(pyobj)
	case C.PyMemoryView_Type_p:
	case C.PyModule_Type_p:
	case C.PySet_Type_p:
	case C.PySlice_Type_p:
	case C.PyTuple_Type_p:
	case C.PyType_Type_p:
	case C.PyUnicode_Type_p:
		return gate.decodeString(pyobj)
	default:
		if C.py_obj_is_none(pyobj) != 0 {
			return pyNone, true
		}
	}

	return nil, true
}

// decodeComplex decodes Python complex number object.
func (gate pyGate) decodeComplex(pyobj pyObject) (complex128, bool) {
	var real, imag C.double

	ok := bool(C.py_complex_get(pyobj, &real, &imag))
	var c complex128
	if ok {
		c = complex(float64(real), float64(imag))
	}

	return c, ok
}

// decodeFloat decodes Python float number object.
func (gate pyGate) decodeFloat(pyobj pyObject) (float64, bool) {
	var val C.double

	ok := bool(C.py_float_get(pyobj, &val))

	return float64(val), ok
}

// decodeByteArray decodes Python byte array object as []byte slice.
//
// Python byte array is mutable object, so we return a slice,
// backed by the Python memory.
func (gate pyGate) decodeByteArray(pyobj pyObject) ([]byte, bool) {
	var data unsafe.Pointer
	var size C.size_t

	ok := bool(C.py_bytearray_get(pyobj, &data, &size))
	var bytes []byte
	if ok {
		bytes = unsafe.Slice((*byte)(data), size)
	}

	return bytes, ok
}

// decodeBytes decodes Python bytes object as []byte slice.
//
// Python bytes are immutable, so we return a copy.
func (gate pyGate) decodeBytes(pyobj pyObject) ([]byte, bool) {
	var data unsafe.Pointer
	var size C.size_t

	ok := bool(C.py_bytes_get(pyobj, &data, &size))
	var bytes []byte
	if ok {
		src := unsafe.Slice((*byte)(data), size)
		bytes = make([]byte, size, size)
		copy(bytes, src)
	}

	return bytes, ok
}

// decodeInteger decodes Python integer object as int or big.Int
func (gate pyGate) decodeInteger(pyobj pyObject) (any, bool) {
	var overflow C.bool
	var val C.int64_t

	ok := bool(C.py_long_get_int64(pyobj, &val, &overflow))
	if !ok {
		return nil, true
	}

	if !bool(overflow) && C.long(int(val)) == val {
		return int(val), true
	}

	s, ok := gate.repr(pyobj)
	if !ok {
		return nil, false
	}

	v := big.NewInt(0)
	_, ok = v.SetString(s, 10)
	assert.Must(ok) // FIXME

	return v, true
}

// decodeString decodes Python Unicode object as a string.
func (gate pyGate) decodeString(pyobj pyObject) (string, bool) {
	sz := C.py_str_len(pyobj)
	assert.Must(sz >= 0) // It only happens if PyObject not Unicode

	s := ""
	if sz > 0 {
		buf := make([]rune, sz)
		p := (*C.Py_UCS4)(unsafe.Pointer(&buf[0]))
		C.py_str_get(pyobj, p, C.size_t(sz))
		s = string(buf)
	}

	return s, true
}

// makeBigint makes a new PyLong_type object from *big.Int.
// It returns strong object reference on success, nil on an error.
func (gate pyGate) makeBigint(v *big.Int) pyObject {
	cs := C.CString(v.String())
	defer C.free(unsafe.Pointer(cs))
	return C.py_long_from_string(cs)
}

// makeBool makes a new PyBool_Type object.
// It returns strong object reference on success, nil on an error.
func (gate pyGate) makeBool(v bool) pyObject {
	return C.py_bool_make(C.bool(v))
}

// makeBytes makes a new PyList_Bytes object.
// It returns strong object reference on success, nil on an error.
func (gate pyGate) makeBytes(data []byte) pyObject {
	var p unsafe.Pointer
	if len(data) != 0 {
		p = unsafe.Pointer(&data[0])
	}

	return C.py_bytes_make(p, C.size_t(len(data)))
}

// makeComplex makes a new PyComlex_Type object.
// It returns strong object reference on success, nil on an error.
func (gate pyGate) makeComplex(v complex128) pyObject {
	return C.py_complex_make(C.double(real(v)), C.double(imag(v)))
}

// makeDict makes a new PyDict_Type object.
// It returns strong object reference on success, nil on an error.
func (gate pyGate) makeDict() pyObject {
	return C.py_dict_make()
}

// makeFloat makes a new PyFloatType object.
// It returns strong object reference on success, nil on an error.
func (gate pyGate) makeFloat(v float64) pyObject {
	return C.py_float_make(C.double(v))
}

// makeInt makes a new PyLong_type object from int64.
// It returns strong object reference on success, nil on an error.
func (gate pyGate) makeInt(v int64) pyObject {
	return C.py_long_from_int64(C.int64_t(v))
}

// makeList makes a new PyList_Type object of the given size.
// It returns strong object reference on success, nil on an error.
func (gate pyGate) makeList(sz int) pyObject {
	return C.py_list_make(C.size_t(sz))
}

// makeString makes a new PyUnicoder_type object from string.
// It returns strong object reference on success, nil on an error.
func (gate pyGate) makeString(v string) pyObject {
	cs := C.CString(v)
	defer C.free(unsafe.Pointer(cs))
	return C.py_str_make(cs, C.size_t(len(v)))
}

// makeTuple makes a new PyTuple_Type object of the given size.
// It returns strong object reference on success, nil on an error.
func (gate pyGate) makeTuple(sz int) pyObject {
	return C.py_tuple_make(C.size_t(sz))
}

// makeUint makes a new PyLong_type object from uint64.
// It returns strong object reference on success, nil on an error.
func (gate pyGate) makeUint(v uint64) pyObject {
	return C.py_long_from_uint64(C.uint64_t(v))
}

// getListItem retrieves the item of the PyList_Type the specified position.
func (gate pyGate) getListItem(list pyObject, idx int) (pyObject, bool) {
	var answer pyObject
	ok := bool(C.py_list_get(list, C.int(idx), &answer))
	return answer, ok
}

// setListItem sets the item of the PyList_Type the specified position.
// Internally, it creates a new strong reference to the item object.
func (gate pyGate) setListItem(list, item pyObject, idx int) bool {
	return bool(C.py_list_set(list, C.int(idx), item))
}

// setTupleItem sets the item of the PyTuple_Type the specified position.
// Internally, it creates a new strong reference to the item object.
func (gate pyGate) setTupleItem(tuple, item pyObject, idx int) bool {
	return bool(C.py_tuple_set(tuple, C.int(idx), item))
}

// getTupleItem retrieves the item of the PyTuple_Type the specified position.
func (gate pyGate) getTupleItem(tuple pyObject, idx int) (pyObject, bool) {
	var answer pyObject
	ok := bool(C.py_tuple_get(tuple, C.int(idx), &answer))
	return answer, ok
}

// pyInterpEval evaluates string as a Python statement.
//
// The name parameter indicates the Python source file name and
// used only for diagnostic messages.
//
// If expr is true, input interpreted as a Python expression and
// on success, its result returned as a *Object. Otherwise, input
// interpreted as a multi-line Python script, and returned *Object
// will be nil.
func (gate pyGate) eval(s, name string, expr bool) (pyObject, error) {
	// Convert expression and filename to the C strings
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))

	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	// Execute the expression
	var pyobj pyObject
	ok := bool(C.py_interp_eval(cs, cname, C.bool(expr), &pyobj))
	if !ok {
		return nil, gate.lastError()
	}

	return pyobj, nil
}
