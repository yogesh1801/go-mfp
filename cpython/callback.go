// MFP - Miulti-Function Printers and scanners toolkit
// CPython binding.
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Python->Go callbacks

package cpython

import (
	"reflect"
	"runtime/cgo"
	"unsafe"

	"github.com/OpenPrinting/go-mfp/internal/assert"
)

// #include "cpython.h"
//
// void callbackDestroy(PyObject *self);
// PyObject *callbackCall(PyObject *self, PyObject *args);
import "C"

// callbackFunc is the Go function called from callback
type callbackFunc func(args pyObject) pyObject

var callbackCapsuleName = C.CString("go.function")

// callback implements Python->Go callbacks
type callback struct {
	py       *Python        // interpreter that owns the callback
	def      *C.PyMethodDef // Python method definition
	function reflect.Value  // Go function to be called
}

// newCallback creates a new callback.
//
// function must be the Go function that accepts arguments,
// convertible from the Python values and returns none, a single
// value or single value plus error.
func newCallback(py *Python, name string, function any) *callback {
	rv := reflect.ValueOf(function)
	rt := rv.Type()
	assert.Must(rt.Kind() == reflect.Func)

	// Check types of return values.
	//
	// We only accept functions that:
	//   - returns no value
	//   - returns a single value
	//   - returns a single value plus error
	switch {
	case rt.NumOut() <= 1:
	case rt.NumOut() == 2 && rt.Out(1).Implements(reflectErrorType):
	default:
		return nil
	}

	// Create callback structure
	sz := C.size_t(unsafe.Sizeof(C.PyMethodDef{}))
	def := (*C.PyMethodDef)(C.malloc(sz))

	def.ml_name = C.CString(name)
	def.ml_meth = C.PyCFunction(C.callbackCall)
	def.ml_flags = C.METH_VARARGS

	cb := &callback{
		py:       py,
		def:      def,
		function: reflect.ValueOf(function),
	}

	return cb
}

// object creates Python object for the callback.
// This function requires pyGate to be acquired by the caller.
func (cb *callback) object(gate pyGate) (pyObject, error) {
	// Wrap callback into the PyCapsule_Type
	handle := cgo.NewHandle(cb)
	capsule, err := gate.makeCapsule(
		unsafe.Pointer(&handle), callbackCapsuleName,
		C.PyCapsule_Destructor(C.callbackDestroy))
	if err != nil {
		handle.Delete()
		return nil, err
	}

	// Create PyCFunction_Type
	cfunction, err := gate.makeCfunction(cb.def, capsule)
	if err != nil {
		C.py_obj_unref(capsule)
	}

	return cfunction, nil
}

// call performs Python->Go call
func (cb *callback) call(args pyObject) (pyObject, error) {
	// Handle special case: void function w/o arguments
	rt := cb.function.Type()
	if rt.NumIn() == 0 && rt.NumOut() == 0 {
		cb.function.Call(nil)
		return cb.py.pyNone, nil
	}

	// Acquire pyGate
	gate, err := cb.py.gate()
	if err != nil {
		return nil, err
	}
	defer gate.release()

	// Translate input arguments
	//
	// TODO

	// Call the function. Translate returned values
	ret := cb.function.Call(nil)
	switch len(ret) {
	case 0:
		// No return value
		return cb.py.pyNone, nil

	case 2:
		// Returns value and error
		err = ret[1].Interface().(error)
		if err != nil {
			return nil, err
		}
		fallthrough

	case 1:
		// Single return value
		return cb.py.newPyObject(gate, ret[0].Interface())
	}

	return nil, nil
}

// Delete is the callback destructor.
func (cb *callback) Delete() {
	C.free(unsafe.Pointer(cb.def.ml_name))
	C.free(unsafe.Pointer(cb.def))
}

// callbackDestroy wraps callback.Delete into the C-callable function.
//
//export callbackDestroy
func callbackDestroy(capsule pyObject) {
	p := C.py_capsule_get_ptr(capsule, callbackCapsuleName)
	if p != nil {
		handle := *(*cgo.Handle)(p)
		cb := handle.Value().(*callback)
		cb.Delete()
		handle.Delete()
	}
}

// callbackCall wraps callback.call into the C-callable function.
//
//export callbackCall
func callbackCall(capsule, args pyObject) pyObject {
	p := C.py_capsule_get_ptr(capsule, callbackCapsuleName)
	if p != nil {
		handle := *(*cgo.Handle)(p)
		cb := handle.Value().(*callback)
		pyobj, _ := cb.call(args)
		return pyobj
	}
	return nil
}
