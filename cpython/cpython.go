// MFP - Miulti-Function Printers and scanners toolkit
// CPython binding.
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// CPython glue -- the Go side

package cpython

// #cgo pkg-config: python3
// #cgo LDFLAGS: -l python3
//
// #include "cpython.h"
import "C"

import (
	"errors"
	"runtime"
	"sync"
	"unsafe"
)

// pyInterp is the Go name for the *C.PyInterpreterState
type pyInterp = *C.PyInterpreterState

// pyInterpError holds Python initialization error, if any.
var pyInterpError error

// pyInterpNewRequestChan is the channel where requests to create
// new pyInterp are sent to.
//
// These requests are handled by the dedicated thread. The request
// itself is the channel of pyInterp, where response is sent.
var pyInterpNewRequestChan = make(chan chan *C.PyInterpreterState)

// pyNewInterp creates a new pyInterp.
func pyNewInterp() (pyInterp, error) {
	if pyInterpError != nil {
		return nil, pyInterpError
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
func pyInterpEval(interp pyInterp, s string) {
	cs := C.CString(s)
	C.py_interp_eval(interp, cs)
	C.free(unsafe.Pointer(cs))
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
	pyInterpErrorCheck(msg)
	initilized.Done()

	// If no error, serve incoming requests
	if pyInterpError == nil {
		for rq := range pyInterpNewRequestChan {
			interp := C.py_new_interp()
			rq <- interp
		}
	}
}

// pyInterpErrorCheck sets the pyInterpError, if error message is not nil.
func pyInterpErrorCheck(msg *C.char) {
	if msg != nil {
		pyInterpError = errors.New(C.GoString(msg))
	}
}

// init starts a dedicated Python thread.
func init() {
	var initilized sync.WaitGroup

	initilized.Add(1)
	go pyInterpThread(&initilized)
	initilized.Wait()
}
