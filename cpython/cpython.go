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
	"os/exec"
	"runtime"
	"sync"
	"unsafe"
)

// #cgo pkg-config: python3
// #cgo LDFLAGS: -l dl
//
// #include "cpython.h"
import "C"

type (
	// pyThreadState is the Go name for the *C.PyThreadState
	pyThreadState = *C.PyThreadState

	// pyObject is the Go name for the *C.PyObject
	pyObject = *C.PyObject

	// pyObject is the Go name for the *C.PyTypeObject
	pyTypeObject = *C.PyTypeObject
)

var (
	// pyInitError holds Python initialization error, if any.
	pyInitError error
)

// pyInterpNewRequestChan is the channel where requests to create
// new Python sub-interpretes are sent to.
//
// These requests are handled by the dedicated thread. The request
// itself is the channel of pyThreadState, where response is sent.
var pyInterpNewRequestChan = make(chan chan pyThreadState)

// pyNewInterp creates a new Python sub-interpreter and returns
// pointer to its main thread state.
func pyNewInterp() (pyThreadState, error) {
	if pyInitError != nil {
		return nil, pyInitError
	}

	rsp := make(chan pyThreadState)
	pyInterpNewRequestChan <- rsp
	interp := <-rsp

	return interp, nil
}

// pyInterpDelete releases the Python sub-interpreter
func pyInterpDelete(interp pyThreadState) {
	C.py_interp_close(interp)
}

// pyLocateLibPython locates the full path to the libpython3.XX.so library
func pyLocateLibPython() (string, error) {
	script := ""
	script += "import sysconfig;"
	script += "import os;"
	script += "dir=sysconfig.get_config_var('LIBDIR');"
	script += "lib=sysconfig.get_config_var('LDLIBRARY');"
	script += "print(os.path.join(dir,lib),end='');"

	cmd := exec.Command("python3", "-c", script)
	out, err := cmd.CombinedOutput()

	return string(out), err
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

	// Locate and initialize Python library
	var lib string
	lib, pyInitError = pyLocateLibPython()
	if pyInitError == nil {
		clib := C.CString(lib)
		msg := C.py_init(clib)
		C.free(unsafe.Pointer(clib))

		pyInitErrorCheck(msg)
	}

	// Notify caller that initialization is completed
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
