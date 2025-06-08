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

// pyInitError holds Python initialization error, if any.
var pyInitError error

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

// pyEnter temporary attaches the calling thread to the
// Python interpreter.
//
// It must be called before any operations with the interpreter
// are performed and must be paired with the py_leave.
//
// The value it returns must be passed to the corresponding
// pyEnter call.
func pyEnter(interp pyInterp) *C.PyThreadState {
	runtime.LockOSThread()
	return C.py_enter(interp)
}

// pyLeave detaches the calling thread from the Python interpreter.
//
// Its parameter must be the value, previously returned by the
// corresponding pyLeave call.
func pyLeave(prev *C.PyThreadState) {
	C.py_leave(prev)
	runtime.UnlockOSThread()
}

// pyInterpEval evaluates string as a Python statement.
func pyInterpEval(interp pyInterp, s string) *Object {
	// Lock the interpreter
	prev := pyEnter(interp)
	defer pyLeave(prev)

	// Convert expression to the C string
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))

	// Execute the expression
	pyobj := C.py_interp_eval(cs)

	// Decode result
	native, ok := pyObjectDecode(pyobj)
	return newObjectFromPython(interp, pyobj, native, ok)
}

// pyObjectDecode decodes PyObject value as Go value.
// It MUST be called between pyEnter/pyLeave calls.
func pyObjectDecode(pyobj pyObject) (any, bool) {
	switch pyObjectType(pyobj) {
	case C.PyBool_Type_p:
		return C.py_obj_is_true(pyobj) != 0, true
	case C.PyByteArray_Type_p:
	case C.PyBytes_Type_p:
	case C.PyCFunction_Type_p:
	case C.PyComplex_Type_p:
	case C.PyDict_Type_p:
	case C.PyDictKeys_Type_p:
	case C.PyFloat_Type_p:
	case C.PyFrozenSet_Type_p:
	case C.PyList_Type_p:
	case C.PyLong_Type_p:
		return pyObjectDecodeInteger(pyobj), true
	case C.PyMemoryView_Type_p:
	case C.PyModule_Type_p:
	case C.PySet_Type_p:
	case C.PySlice_Type_p:
	case C.PyTuple_Type_p:
	case C.PyType_Type_p:
	case C.PyUnicode_Type_p:
		return pyObjectDecodeString(pyobj), true
	default:
		if C.py_obj_is_none(pyobj) != 0 {
			return nil, true
		}
	}

	return nil, false
}

// pyObjectDecodeInteger decodes Python object as int or big.Int
// It MUST be called between pyEnter/pyLeave calls.
func pyObjectDecodeInteger(pyobj pyObject) any {
	var overflow C.bool
	var val C.long

	ok := bool(C.py_long_get(pyobj, &val, &overflow))
	assert.Must(ok) // FIXME

	if !bool(overflow) && C.long(int(val)) == val {
		return int(val)
	}

	repr := C.py_obj_repr(pyobj)
	assert.Must(repr != nil) // FIXME

	s := pyObjectDecodeString(repr)
	C.py_obj_unref(repr)

	v := big.NewInt(0)
	_, ok = v.SetString(s, 10)
	assert.Must(ok) // FIXME

	return v
}

// pyObjectDecodeString decodes Python Unicode object as a string.
// It MUST be called between pyEnter/pyLeave calls.
func pyObjectDecodeString(pyobj pyObject) string {
	sz := C.py_str_len(pyobj)
	assert.Must(sz >= 0)

	if sz > 0 {
		buf := make([]rune, sz)
		p := (*C.Py_UCS4)(unsafe.Pointer(&buf[0]))
		C.py_str_get(pyobj, p, C.size_t(sz))
		return string(buf)
	}

	return ""
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

// init starts a dedicated Python thread.
func init() {
	var initilized sync.WaitGroup

	initilized.Add(1)
	go pyInterpThread(&initilized)
	initilized.Wait()
}
