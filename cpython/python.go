// MFP - Miulti-Function Printers and scanners toolkit
// CPython binding.
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Python interpreter.

package cpython

import (
	"fmt"
	"runtime"
)

// Python represents a Python interpreter.
// There are may be many interpreters within a single process.
// Each has its own namespace and isolated from others.
type Python struct {
	interp  pyInterp // Underlying *C.PyInterpreterState
	objects *objmap  // Objects owned by the interpreter
}

// NewPython creates a new Python interpreter.
func NewPython() (py *Python, err error) {
	interp, err := pyNewInterp()
	if err == nil {
		py = &Python{
			interp:  interp,
			objects: newObjmap(),
		}
	}

	return
}

// Close closes the [Python] interpreter and releases all
// resources it holds.
func (py *Python) Close() {
	gate := py.gate()
	py.objects.purge(gate)
	gate.release()

	pyInterpDelete(py.interp)
	py.interp = nil
}

// closed reports if interpreter is closed.
func (py *Python) closed() bool {
	return py.interp == nil
}

// Eval evaluates string as a Python expression and returns its value.
func (py *Python) Eval(s string) (*Object, error) {
	return py.eval(s, "", true)
}

// Exec evaluates string as a Python script.
//
// The filename parameter specifies the Python source file name
// and used only for diagnostic. If set to the empty string (""),
// the reasonable default is provided.
func (py *Python) Exec(s, filename string) error {
	_, err := py.eval(s, filename, false)
	return err
}

// eval is the common body for Python.Eval and Python.Exec
func (py *Python) eval(s, filename string, expr bool) (*Object, error) {
	// Adjust filename to point to the Go file:line position
	// of the called, if filename is not specified
	if filename == "" {
		pc := make([]uintptr, 1)
		if n := runtime.Callers(3, pc); n > 0 {
			frames := runtime.CallersFrames(pc)
			frame, _ := frames.Next()
			filename = fmt.Sprintf("%s:%d", frame.File, frame.Line)
		}
	}

	// Obtain pyGate
	gate := py.gate()
	defer gate.release()

	// Call interpreter
	pyobj, err := gate.eval(s, filename, expr)
	if pyobj == nil {
		return nil, err
	}

	// Decode the Object
	native, ok := gate.decodeObject(pyobj)
	if !ok {
		gate.unref(pyobj)
		return nil, gate.lastError()
	}

	oid := py.newObjID(gate, pyobj)
	obj := newObjectFromPython(py, oid, native)
	return obj, err
}

// gate is the convenience wrapper for pyGateAcquire(py.interp)
func (py *Python) gate() pyGate {
	return pyGateAcquire(py.interp)
}

// newObjID allocates new objiD for the *C.PyObject.
func (py *Python) newObjID(gate pyGate, obj pyObject) objid {
	return py.objects.put(gate, obj)
}

// delObjID deletes *C.PyObject by objid
func (py *Python) delObjID(gate pyGate, oid objid) {
	py.objects.del(gate, oid)
}

// lookupObjID return *C.PyObject by objid
func (py *Python) lookupObjID(gate pyGate, oid objid) pyObject {
	return py.objects.get(gate, oid)
}
