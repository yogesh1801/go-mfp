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
	gate.lastError() // Reset pending error condition, if any
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
		msg, err := gate.str(evalue)
		if err == nil {
			return ErrPython{msg}
		}
	}

	if etype != nil {
		msg, err := gate.str(etype)
		if err == nil {
			return ErrPython{msg}
		}
	}

	return ErrPython{"Unknown Python exception"}
}

// objOrLastError returns pyobj, if it is not nil, or gate.lastError()
func (gate pyGate) objOrLastError(pyobj pyObject) (pyObject, error) {
	if pyobj != nil {
		return pyobj, nil
	}

	return nil, gate.lastError()
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
func (gate pyGate) str(pyobj pyObject) (s string, err error) {
	str := C.py_obj_str(pyobj)
	if str != nil {
		defer C.py_obj_unref(str)
		s, err = gate.decodeUnicode(str)
	}

	return
}

// repr returns repr(pyobj), decoded as Go string.
func (gate pyGate) repr(pyobj pyObject) (s string, err error) {
	repr := C.py_obj_repr(pyobj)
	if repr != nil {
		defer C.py_obj_unref(repr)
		s, err = gate.decodeUnicode(repr)
	}

	return
}

// typename returns name of the PyObject's type
func (gate pyGate) typename(pyobj pyObject) string {
	name := "unknown"

	t := pyObject(unsafe.Pointer(C.Py_TYPE(pyobj)))
	n, err := gate.getattr(t, "__name__")
	if err == nil {
		s, err := gate.str(n)
		if err == nil {
			name = s
		}
	}

	return name
}

// delattr deletes Object attribute with the specified name.
func (gate pyGate) delattr(pyobj pyObject, name string) error {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	ok := bool(C.py_obj_delattr(pyobj, cname))
	if !ok {
		return gate.lastError()
	}

	return nil
}

// getattr returns Object attribute with the specified name.
func (gate pyGate) getattr(pyobj pyObject, name string) (pyObject, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	var attr pyObject
	if !bool(C.py_obj_getattr(pyobj, cname, &attr)) {
		return nil, gate.lastError()
	}

	return attr, nil
}

// hasattr reports if Object has attribute with the specified name.
func (gate pyGate) hasattr(pyobj pyObject, name string) (bool, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	var found C.bool
	if !bool(C.py_obj_hasattr(pyobj, cname, &found)) {
		return false, gate.lastError()
	}

	return bool(found), nil
}

// setattr sets Object attribute with the specified name.
func (gate pyGate) setattr(pyobj pyObject, name string, val pyObject) error {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	if !bool(C.py_obj_setattr(pyobj, cname, val)) {
		return gate.lastError()
	}

	return nil
}

// delitem deletes Object item with the specified key:
//
//	del(pyobj[key])
func (gate pyGate) delitem(pyobj, key pyObject) error {
	if !bool(C.py_obj_delitem(pyobj, key)) {
		return gate.lastError()
	}

	return nil
}

// getitem returns Object item with the specified key.
//
//	pyobj[key]
func (gate pyGate) getitem(pyobj, key pyObject) (pyObject, error) {
	var item pyObject
	if !bool(C.py_obj_getitem(pyobj, key, &item)) {
		return nil, gate.lastError()
	}

	return item, nil
}

// hasitem reports if Object has item with the specified key.
//
//	key in pyobj
func (gate pyGate) hasitem(pyobj, key pyObject) (bool, error) {
	var found C.bool
	if !bool(C.py_obj_hasitem(pyobj, key, &found)) {
		return false, gate.lastError()
	}

	return bool(found), nil
}

// setitem sets Object item with the specified key.
//
//	pyobj[key] = val
func (gate pyGate) setitem(pyobj, key, val pyObject) error {
	if !bool(C.py_obj_setitem(pyobj, key, val)) {
		return gate.lastError()
	}

	return nil
}

// call calls callable object as a function, with positional
// arguments, defined by args (must be PyTuple_Type) and keyword
// arguments, defined by kwargs (must be PyDict_Type or nil).
//
// It returns strong reference to result on success, nil on an error.
func (gate pyGate) call(callable, args, kwargs pyObject) (pyObject, error) {
	return gate.objOrLastError(C.py_obj_call(callable, args, kwargs))
}

// callable reports if object is callable.
// This function always succeeds.
func (gate pyGate) callable(pyobj pyObject) bool {
	return bool(C.py_obj_callable(pyobj))
}

// isNone reports if PyObject is None
func (gate pyGate) isNone(pyobj pyObject) bool {
	return bool(C.py_obj_is_none(pyobj))
}

// decodeError returns [ErrTypeConversion] for Python->Go conversion.
func (gate pyGate) decodeError(pyobj pyObject, to string) error {
	return ErrTypeConversion{
		from: gate.typename(pyobj),
		to:   to,
	}
}

// decodeBigint decodes PyLong_Type object as *big.Int
func (gate pyGate) decodeBigint(pyobj pyObject) (*big.Int, error) {
	if !bool(C.py_obj_is_long(pyobj)) {
		return nil, gate.decodeError(pyobj, "*big.Int,")
	}

	s, err := gate.repr(pyobj)
	if err != nil {
		return nil, err
	}

	v := big.NewInt(0)
	_, ok := v.SetString(s, 10)
	assert.Must(ok) // FIXME

	return v, nil
}

// decodeBool decodes PyObject into bool
func (gate pyGate) decodeBool(pyobj pyObject) (bool, error) {
	switch {
	case bool(C.py_obj_is_true(pyobj)):
		return true, nil
	case bool(C.py_obj_is_false(pyobj)):
		return false, nil
	}

	return false, gate.decodeError(pyobj, "bool")
}

// decodeBytes decodes PyBytes_Type or PyByteArray_Type object as []byte slice.
func (gate pyGate) decodeBytes(pyobj pyObject) ([]byte, error) {
	if bool(C.py_obj_is_bytes(pyobj)) {
		// PyBytes_Type are immutable, so we return a copy.
		var data unsafe.Pointer
		var size C.size_t

		ok := bool(C.py_bytes_get(pyobj, &data, &size))
		if !ok {
			return nil, gate.lastError()
		}

		bytes := make([]byte, size, size)
		src := unsafe.Slice((*byte)(data), size)
		copy(bytes, src)

		return bytes, nil
	}

	if bool(C.py_obj_is_byte_array(pyobj)) {
		// Python byte array is mutable object, so we return a slice,
		// backed by the Python memory.
		var data unsafe.Pointer
		var size C.size_t

		ok := bool(C.py_bytearray_get(pyobj, &data, &size))
		if !ok {
			return nil, gate.lastError()
		}

		bytes := unsafe.Slice((*byte)(data), size)
		return bytes, nil
	}

	return nil, gate.decodeError(pyobj, "[]byte")
}

// decodeComplex decodes PyComplex_Type object as complex128.
func (gate pyGate) decodeComplex(pyobj pyObject) (complex128, error) {
	if !bool(C.py_obj_is_complex(pyobj)) {
		return 0, gate.decodeError(pyobj, "complex128")
	}

	var real, imag C.double

	ok := bool(C.py_complex_get(pyobj, &real, &imag))
	if !ok {
		return 0, gate.lastError()
	}

	return complex(float64(real), float64(imag)), nil
}

// decodeFloat decodes PyFloat_Type object as float64.
func (gate pyGate) decodeFloat(pyobj pyObject) (float64, error) {
	if !bool(C.py_obj_is_float(pyobj)) {
		return 0, gate.decodeError(pyobj, "float64")
	}

	var val C.double

	ok := bool(C.py_float_get(pyobj, &val))
	if !ok {
		return 0, gate.lastError()
	}

	return float64(val), nil
}

// decodeInt64 decodes PyLong_Type object as int64
func (gate pyGate) decodeInt64(pyobj pyObject) (int64, error) {
	if !bool(C.py_obj_is_long(pyobj)) {
		return 0, gate.decodeError(pyobj, "int64")
	}

	var val C.int64_t
	var ovf C.bool

	ok := bool(C.py_long_get_int64(pyobj, &val, &ovf))
	switch {
	case !ok:
		return 0, gate.lastError()
	case bool(ovf):
		s, _ := gate.repr(pyobj)
		return 0, ErrOverflow{s}
	}

	return int64(val), nil
}

// decodeInt64 decodes PyLong_Type object as uint64
func (gate pyGate) decodeUint64(pyobj pyObject) (uint64, error) {
	if !bool(C.py_obj_is_long(pyobj)) {
		return 0, gate.decodeError(pyobj, "uint64")
	}

	var val C.uint64_t
	var ovf C.bool

	ok := bool(C.py_long_get_uint64(pyobj, &val, &ovf))
	switch {
	case !ok:
		return 0, gate.lastError()
	case bool(ovf):
		s, _ := gate.repr(pyobj)
		return 0, ErrOverflow{s}
	}

	return uint64(val), nil
}

// decodeUnicode decodes PyUnicode_Type object as a string.
func (gate pyGate) decodeUnicode(pyobj pyObject) (string, error) {
	if !bool(C.py_obj_is_unicode(pyobj)) {
		return "", gate.decodeError(pyobj, "string")
	}

	sz := C.py_str_len(pyobj)
	assert.Must(sz >= 0) // It only happens if PyObject not Unicode

	s := ""
	if sz > 0 {
		buf := make([]rune, sz)
		p := (*C.Py_UCS4)(unsafe.Pointer(&buf[0]))
		C.py_str_get(pyobj, p, C.size_t(sz))
		s = string(buf)
	}

	return s, nil
}

// makeBigint makes a new PyLong_type object from *big.Int.
// It returns strong object reference on success, nil on an error.
func (gate pyGate) makeBigint(v *big.Int) (pyObject, error) {
	cs := C.CString(v.String())
	defer C.free(unsafe.Pointer(cs))

	return gate.objOrLastError(C.py_long_from_string(cs))
}

// makeBool makes a new PyBool_Type object.
// It returns strong object reference on success, nil on an error.
func (gate pyGate) makeBool(v bool) (pyObject, error) {
	return gate.objOrLastError(C.py_bool_make(C.bool(v)))
}

// makeBytes makes a new PyList_Bytes object.
// It returns strong object reference on success, nil on an error.
func (gate pyGate) makeBytes(data []byte) (pyObject, error) {
	var p unsafe.Pointer
	if len(data) != 0 {
		p = unsafe.Pointer(&data[0])
	}

	return gate.objOrLastError(C.py_bytes_make(p, C.size_t(len(data))))
}

// makeComplex makes a new PyComlex_Type object.
// It returns strong object reference on success, nil on an error.
func (gate pyGate) makeComplex(v complex128) (pyObject, error) {
	pyobj := C.py_complex_make(C.double(real(v)), C.double(imag(v)))
	return gate.objOrLastError(pyobj)
}

// makeDict makes a new PyDict_Type object.
// It returns strong object reference on success, nil on an error.
func (gate pyGate) makeDict() (pyObject, error) {
	return gate.objOrLastError(C.py_dict_make())
}

// makeFloat makes a new PyFloatType object.
// It returns strong object reference on success, nil on an error.
func (gate pyGate) makeFloat(v float64) (pyObject, error) {
	return gate.objOrLastError(C.py_float_make(C.double(v)))
}

// makeInt makes a new PyLong_type object from int64.
// It returns strong object reference on success, nil on an error.
func (gate pyGate) makeInt(v int64) (pyObject, error) {
	return gate.objOrLastError(C.py_long_from_int64(C.int64_t(v)))
}

// makeList makes a new PyList_Type object of the given size.
// It returns strong object reference on success, nil on an error.
func (gate pyGate) makeList(sz int) (pyObject, error) {
	return gate.objOrLastError(C.py_list_make(C.size_t(sz)))
}

// makeString makes a new PyUnicoder_type object from string.
// It returns strong object reference on success, nil on an error.
func (gate pyGate) makeString(v string) (pyObject, error) {
	cs := C.CString(v)
	defer C.free(unsafe.Pointer(cs))
	return gate.objOrLastError(C.py_str_make(cs, C.size_t(len(v))))
}

// makeTuple makes a new PyTuple_Type object of the given size.
// It returns strong object reference on success, nil on an error.
func (gate pyGate) makeTuple(sz int) (pyObject, error) {
	return gate.objOrLastError(C.py_tuple_make(C.size_t(sz)))
}

// makeUint makes a new PyLong_type object from uint64.
// It returns strong object reference on success, nil on an error.
func (gate pyGate) makeUint(v uint64) (pyObject, error) {
	return gate.objOrLastError(C.py_long_from_uint64(C.uint64_t(v)))
}

// getListItem retrieves the item of the PyList_Type the specified position.
func (gate pyGate) getListItem(list pyObject, idx int) (pyObject, error) {
	var answer pyObject
	if !bool(C.py_list_get(list, C.int(idx), &answer)) {
		return nil, gate.lastError()
	}
	return answer, nil
}

// setListItem sets the item of the PyList_Type the specified position.
// Internally, it creates a new strong reference to the item object.
func (gate pyGate) setListItem(list, item pyObject, idx int) error {
	if !bool(C.py_list_set(list, C.int(idx), item)) {
		return gate.lastError()
	}
	return nil
}

// setTupleItem sets the item of the PyTuple_Type the specified position.
// Internally, it creates a new strong reference to the item object.
func (gate pyGate) setTupleItem(tuple, item pyObject, idx int) error {
	if !bool(C.py_tuple_set(tuple, C.int(idx), item)) {
		return gate.lastError()
	}
	return nil
}

// getTupleItem retrieves the item of the PyTuple_Type the specified position.
func (gate pyGate) getTupleItem(tuple pyObject, idx int) (pyObject, error) {
	var answer pyObject
	if !bool(C.py_tuple_get(tuple, C.int(idx), &answer)) {
		return nil, gate.lastError()
	}
	return answer, nil
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
