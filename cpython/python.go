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
	"math/big"
	"reflect"
	"runtime"

	"github.com/OpenPrinting/go-mfp/internal/assert"
)

// Python represents a Python interpreter.
// There are may be many interpreters within a single process.
// Each has its own namespace and isolated from others.
type Python struct {
	interp  pyInterp // Underlying *C.PyInterpreterState
	objects *objmap  // Objects owned by the interpreter
	none    pyObject // Cached None object
}

// NewPython creates a new Python interpreter.
func NewPython() (py *Python, err error) {
	interp, err := pyNewInterp()
	if err == nil {
		py = &Python{
			interp:  interp,
			objects: newObjmap(),
		}

		gate := py.gate()
		defer gate.release()

		py.none, err = gate.eval("None", "", true)
		assert.NoError(err)
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

// NewObject creates a new Python Object for the Go value.
//
// The following Go types are supported:
//
//	Go                              Python
//	==                              ======
//
//	nil                             None
//
//	bool and derivatives            PyBool_Type
//
//	int, int8, int16, int32,        PyLong_Type
//	int64, uint, uint8, uint16,
//	uint32, uint64 and derivatives
//
//	string and derivatives          PyUnicode_Type
//
//	[*big.Int]                      PyLong_Type
//
//	*Object                         new reference to the same PyObject
//
//	[]byte, [...]byte               PyBytes_Type
//
//	[]any, [...]any                 PyList_Type
//
//	[integer]any, [string]any       PyDict_Type
func (py *Python) NewObject(val any) (*Object, error) {
	gate := py.gate()
	defer gate.release()

	pyobj, err := py.newPyObject(gate, val)
	if pyobj == nil && err == nil {
		err = gate.lastError()
	}

	if err != nil {
		return nil, err
	}

	oid := py.newObjID(gate, pyobj)
	obj := newObjectFromPython(py, oid, val)

	return obj, nil
}

// newPyObject is the internal function behind the [Python.NewObject]
// which does all the dirty work of conversion value into the
// Python Object.
//
// It returns either pyObject or error. If it returns (nil, nil),
// gate.lastError needs to be consulted.
func (py *Python) newPyObject(gate pyGate, val any) (pyObject, error) {
	// Handle special cases
	if val == nil {
		return py.none, nil
	}

	switch v := val.(type) {
	case *big.Int:
		return gate.makeBigint(v), nil
	case *Object:
		pyobj := py.lookupObjID(gate, v.oid)
		gate.ref(pyobj)
		return pyobj, nil
	}

	// Generic conversions
	rv := reflect.ValueOf(val)
	rt := rv.Type()

	switch rt.Kind() {
	case reflect.Bool:
		return gate.makeBool(rv.Bool()), nil

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Int64:
		return gate.makeInt(rv.Int()), nil

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
		reflect.Uint64:
		return gate.makeUint(rv.Uint()), nil

	case reflect.Float32, reflect.Float64:
		return gate.makeFloat(rv.Float()), nil

	case reflect.Complex64, reflect.Complex128:
		return gate.makeComplex(rv.Complex()), nil

	case reflect.String:
		return gate.makeString(rv.String()), nil

	case reflect.Array, reflect.Slice:
		if rt.Elem().Kind() == reflect.Uint8 {
			if rt.Kind() == reflect.Array && !rv.CanAddr() {
				tmp := reflect.New(rt).Elem()
				reflect.Copy(tmp, rv)
				rv = tmp
			}

			data := rv.Bytes()
			return gate.makeBytes(data), nil
		}
		return py.newPyList(gate, val)

	case reflect.Chan:
	case reflect.Func:
	case reflect.Interface:
	case reflect.Map:
	case reflect.Pointer:
	case reflect.Struct:
	case reflect.Uintptr:
	case reflect.UnsafePointer:
	}

	err := ErrTypeConversion{from: rt}
	return nil, err
}

// newPyList creates PyObject from array of slice of values.
//
// It returns either pyObject or error. If it returns (nil, nil),
// gate.lastError needs to be consulted.
func (py *Python) newPyList(gate pyGate, val any) (pyObject, error) {
	rv := reflect.ValueOf(val)
	sz := rv.Len()

	list := gate.makeList(sz)
	if list == nil {
		return nil, nil
	}

	defer gate.unref(list)

	for i := 0; i < sz; i++ {
		item := rv.Index(i).Interface()
		pyobj, err := py.newPyObject(gate, item)
		if pyobj == nil {
			return nil, err
		}

		ok := gate.setListItem(list, pyobj, i)
		gate.unref(pyobj) // Now owned by list

		if !ok {
			return nil, nil
		}
	}

	gate.ref(list)
	return list, nil
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

// delObjID deletes *C.PyObject by objid.
func (py *Python) delObjID(gate pyGate, oid objid) {
	py.objects.del(gate, oid)
}

// lookupObjID return *C.PyObject by objid.
func (py *Python) lookupObjID(gate pyGate, oid objid) pyObject {
	return py.objects.get(gate, oid)
}

// countObjID returns count of active objid mappings.
// This is the testing interface
func (py *Python) countObjID(gate pyGate) int {
	return py.objects.count(gate)
}
