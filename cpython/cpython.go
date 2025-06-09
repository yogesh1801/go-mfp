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

// pyInterpEval evaluates string as a Python statement.
func pyInterpEval(interp pyInterp, s string) (*Object, error) {
	// Lock the interpreter
	ref := pyRefAcquire(interp)
	defer ref.release()

	// Convert expression to the C string
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))

	// Execute the expression
	pyobj := C.py_interp_eval(cs)
	if pyobj == nil {
		return nil, ref.lastError()
	}

	// Decode result
	native, ok := ref.decodeObject(pyobj)
	if !ok {
		err := ref.lastError()
		if err != nil {
			return nil, err
		}
	}

	return newObjectFromPython(interp, pyobj, native, ok), nil
}

// pyRef represents the locked (attached to the current thread
// and with the GIL acquired) state of the Python interpreter.
//
// It works as a reference to the interpreter and implements
// all interpreter operations that require locking.
type pyRef struct {
	prev *C.PyThreadState // Previous current thread state
}

// pyRefAcquire temporary attaches the calling thread to the
// Python interpreter.
//
// It returns the pyRef object, that must be released after
// use with the [pyRef.release] call.
func pyRefAcquire(interp pyInterp) pyRef {
	runtime.LockOSThread()
	prev := C.py_enter(interp)
	return pyRef{prev}
}

// release detaches the calling thread from the Python interpreter.
func (ref pyRef) release() {
	C.py_leave(ref.prev)
	runtime.UnlockOSThread()
}

// lastError returns a last error, nil if none.
func (ref pyRef) lastError() error {
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
		msg, ok := ref.str(evalue)
		if ok {
			return Error{msg}
		}
	}

	if etype != nil {
		msg, ok := ref.str(evalue)
		if ok {
			return Error{msg}
		}
	}

	return Error{"Unknown Python exception"}
}

// str returns str(pyobj), decoded as Go string.
func (ref pyRef) str(pyobj pyObject) (s string, ok bool) {
	str := C.py_obj_str(pyobj)
	if str != nil {
		defer C.py_obj_unref(str)
		s, ok = ref.decodeString(str)
	}

	return
}

// repr returns repr(pyobj), decoded as Go string.
func (ref pyRef) repr(pyobj pyObject) (s string, ok bool) {
	repr := C.py_obj_repr(pyobj)
	if repr != nil {
		defer C.py_obj_unref(repr)
		s, ok = ref.decodeString(repr)
	}

	return
}

// decodeObject decodes PyObject value as Go value.
func (ref pyRef) decodeObject(pyobj pyObject) (any, bool) {
	switch pyObjectType(pyobj) {
	case C.PyBool_Type_p:
		return C.py_obj_is_true(pyobj) != 0, true
	case C.PyByteArray_Type_p:
		return ref.decodeByteArray(pyobj)
	case C.PyBytes_Type_p:
		return ref.decodeBytes(pyobj)
	case C.PyCFunction_Type_p:
	case C.PyComplex_Type_p:
	case C.PyDict_Type_p:
	case C.PyDictKeys_Type_p:
	case C.PyFloat_Type_p:
	case C.PyFrozenSet_Type_p:
	case C.PyList_Type_p:
	case C.PyLong_Type_p:
		return ref.decodeInteger(pyobj)
	case C.PyMemoryView_Type_p:
	case C.PyModule_Type_p:
	case C.PySet_Type_p:
	case C.PySlice_Type_p:
	case C.PyTuple_Type_p:
	case C.PyType_Type_p:
	case C.PyUnicode_Type_p:
		return ref.decodeString(pyobj)
	default:
		if C.py_obj_is_none(pyobj) != 0 {
			return nil, true
		}
	}

	return nil, false
}

// decodeByteArray decodes Python byte array object as []byte slice.
//
// Python byte array is mutable object, so we return a slice,
// backed by the Python memory.
func (ref pyRef) decodeByteArray(pyobj pyObject) ([]byte, bool) {
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
func (ref pyRef) decodeBytes(pyobj pyObject) ([]byte, bool) {
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
func (ref pyRef) decodeInteger(pyobj pyObject) (any, bool) {
	var overflow C.bool
	var val C.long

	ok := bool(C.py_long_get(pyobj, &val, &overflow))
	if !ok {
		return nil, true
	}

	if !bool(overflow) && C.long(int(val)) == val {
		return int(val), true
	}

	s, ok := ref.repr(pyobj)
	if !ok {
		return nil, false
	}

	v := big.NewInt(0)
	_, ok = v.SetString(s, 10)
	assert.Must(ok) // FIXME

	return v, true
}

// decodeString decodes Python Unicode object as a string.
func (ref pyRef) decodeString(pyobj pyObject) (string, bool) {
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
