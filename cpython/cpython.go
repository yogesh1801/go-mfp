// MFP - Miulti-Function Printers and scanners toolkit
// CPython binding.
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// CPython glue -- the Go side

package cpython

import (
	"errors"
	"math/big"
	"runtime"
	"sync"
	"unsafe"

	"github.com/OpenPrinting/go-mfp/internal/assert"
)

// #cgo pkg-config: python3
// #cgo LDFLAGS: -l python3
//
// #include "cpython.h"
import "C"

type (
	// pyInterp is the Go name for the *C.PyInterpreterState
	pyInterp = *C.PyInterpreterState

	// pyObject is the Go name for the *C.PyObject
	pyObject = *C.PyObject

	// pyObject is the Go name for the *C.PyObject
	pyTypeObject = *C.PyTypeObject
)

var (
	// pyInitError holds Python initialization error, if any.
	pyInitError error

	// pyNoNativeValue is returned by gate.decodeObject
	// when Python object is None
	pyNone = struct{}{}
)

// pyInterpNewRequestChan is the channel where requests to create
// new pyInterp are sent to.
//
// These requests are handled by the dedicated thread. The request
// itself is the channel of pyInterp, where response is sent.
var pyInterpNewRequestChan = make(chan chan *C.PyInterpreterState)

// pyNewInterp creates a new pyInterp.
func pyNewInterp() (pyInterp, error) {
	if pyInitError != nil {
		return nil, pyInitError
	}

	rsp := make(chan *C.PyInterpreterState)
	pyInterpNewRequestChan <- rsp
	interp := <-rsp

	return interp, nil
}

// pyInterpDelete releases the pyInterp.
func pyInterpDelete(interp pyInterp) {
	C.py_interp_close(interp)
}

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
			return Error{msg}
		}
	}

	if etype != nil {
		msg, ok := gate.str(evalue)
		if ok {
			return Error{msg}
		}
	}

	return Error{"Unknown Python exception"}
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

// delattr returns Object attribute with the specified name.
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

// setattr sets Object has attribute with the specified name.
func (gate pyGate) setattr(pyobj pyObject, name string, val pyObject) (ok bool) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	ok = bool(C.py_obj_setattr(pyobj, cname, val))
	return
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
	var val C.long

	ok := bool(C.py_long_get(pyobj, &val, &overflow))
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

// pyInterpThread runs Python dedicated thread.
//
// We need this thread, because CPython pollutes the thread local
// storage with its own stuff while doing operations like Python
// initialization and new interpreter creation, and we don't want
// this litter to be distributed across all Go threads.
//
// Once initialization is completed, pyInterpThread signals it
// over the supplied sync.WaitGroup parameter.
func pyInterpThread(initilized *sync.WaitGroup) {
	// We need the dedicated thread...
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	// Initialize Python library
	msg := C.py_init()
	pyInitErrorCheck(msg)
	initilized.Done()

	// If no error, serve incoming requests
	if pyInitError == nil {
		for rq := range pyInterpNewRequestChan {
			interp := C.py_new_interp()
			rq <- interp
		}
	}
}

// pyInitErrorCheck sets the pyInitError, if error message is not nil.
func pyInitErrorCheck(msg *C.char) {
	if msg != nil {
		pyInitError = errors.New(C.GoString(msg))
	}
}

// pyInitErrorCheckTest is the test interface to the pyInitErrorCheck
// function. We cannot call pyInitErrorCheck directly from tests, because
// CGo is not available in the _test.go files, so we cannot construct
// *C.char string out of there.
func pyInitErrorCheckTest(s string) {
	msg := C.CString(s)
	defer C.free(unsafe.Pointer(msg))
	pyInitErrorCheck(msg)
}

// init starts a dedicated Python thread.
func init() {
	var initilized sync.WaitGroup

	initilized.Add(1)
	go pyInterpThread(&initilized)
	initilized.Wait()
}
