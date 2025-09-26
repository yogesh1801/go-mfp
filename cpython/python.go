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
	interp   pyInterp // Underlying *C.PyInterpreterState
	objects  *objmap  // Objects owned by the interpreter
	pyNone   pyObject // Cached None pyObject
	pyTrue   pyObject // Cached True pyObject
	pyFalse  pyObject // Cached False pyObject
	objNone  *Object  // Cached None Object
	objTrue  *Object  // Cached True Object
	objFalse *Object  // Cached False Object
	globals  *Object  // Global dictionary
}

// NewPython creates a new Python interpreter.
func NewPython() (py *Python, err error) {
	interp, err := pyNewInterp()
	if err == nil {
		py = &Python{
			interp:  interp,
			objects: newObjmap(),
		}

		// Load None, True and False singletons
		gate := py.gate()

		py.pyNone, err = gate.eval("None", "", true)
		assert.NoError(err)

		py.pyTrue, err = gate.eval("True", "", true)
		assert.NoError(err)

		py.pyFalse, err = gate.eval("False", "", true)
		assert.NoError(err)

		py.objNone = newObjectFromPython(py, gate, py.pyNone)
		py.objTrue = newObjectFromPython(py, gate, py.pyTrue)
		py.objFalse = newObjectFromPython(py, gate, py.pyFalse)

		gate.release()

		// Load global dictionary.
		// It is more convenient to use it as the Object
		py.globals, err = py.Eval("globals()")
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

// GetGlobal returns item from the interpreter's global dictionary.
//
// In Python:
//
//	globals()[name]
//
// It returns:
//   - (*Object, nil) if item was found
//   - (nil, nil) if item was not found
//   - (nil, error) in a case of error
func (py *Python) GetGlobal(name string) (*Object, error) {
	return py.globals.Get(name)
}

// SetGlobal sets item in the interpreter's global dictionary.
//
// In Python:
//
//	globals()[name] = val
//
// The val may be any value that [Python.NewObject] accepts.
func (py *Python) SetGlobal(name string, val any) error {
	return py.globals.Set(name, val)
}

// DelGlobal deletes item the interpreter's global dictionary.
//
// In Python:
//
//	del(globals(), name)
//
// It returns:
//   - (true, nil) if item was found and deleted
//   - (false, nil) if item was not found
//   - (false, error) in a case of error
func (py *Python) DelGlobal(name string) (bool, error) {
	return py.globals.Del(name)
}

// ContainsGlobal reports if the interpreter's global dictionary
// contains the named item.
//
// In Python:
//
//	name in globals()
func (py *Python) ContainsGlobal(name string) (bool, error) {
	return py.globals.Contains(name)
}

// None returns the None Object
func (py *Python) None() *Object {
	return py.objNone
}

// Bool returns the boolean Object
func (py *Python) Bool(v bool) *Object {
	if v {
		return py.objTrue
	}

	return py.objFalse
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
//	[cmp.Ordered or bool]any        PyDict_Type
func (py *Python) NewObject(val any) (*Object, error) {
	gate := py.gate()
	defer gate.release()

	pyobj, err := py.newPyObject(gate, val)
	if err != nil {
		return nil, err
	}

	obj := newObjectFromPython(py, gate, pyobj)

	return obj, nil
}

// newPyObject is the internal function behind the [Python.NewObject]
// which does all the dirty work of conversion value into the
// Python Object.
func (py *Python) newPyObject(gate pyGate, val any) (pyObject, error) {
	// Handle special cases
	if val == nil {
		return py.pyNone, nil
	}

	switch v := val.(type) {
	case *big.Int:
		return gate.makeBigint(v)
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
		return gate.makeBool(rv.Bool())

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Int64:
		return gate.makeInt(rv.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
		reflect.Uint64:
		return gate.makeUint(rv.Uint())

	case reflect.Float32, reflect.Float64:
		return gate.makeFloat(rv.Float())

	case reflect.Complex64, reflect.Complex128:
		return gate.makeComplex(rv.Complex())

	case reflect.String:
		return gate.makeString(rv.String())

	case reflect.Array, reflect.Slice:
		if rt.Elem().Kind() == reflect.Uint8 {
			if rt.Kind() == reflect.Array && !rv.CanAddr() {
				tmp := reflect.New(rt).Elem()
				reflect.Copy(tmp, rv)
				rv = tmp
			}

			data := rv.Bytes()
			return gate.makeBytes(data)
		}
		return py.newPyList(gate, val)

	case reflect.Map:
		return py.newPyDict(gate, val)

	case reflect.Chan:
	case reflect.Func:
	case reflect.Interface:
	case reflect.Pointer:
	case reflect.Struct:
	case reflect.Uintptr:
	case reflect.UnsafePointer:
	}

	err := ErrTypeConversion{
		from: rt.String(),
		to:   "Python object",
	}

	return nil, err
}

// newPyList creates PyObject from array of slice of values.
func (py *Python) newPyList(gate pyGate, val any) (pyObject, error) {
	rv := reflect.ValueOf(val)
	sz := rv.Len()

	list, err := gate.makeList(sz)
	if err != nil {
		return nil, err
	}

	defer gate.unref(list)

	for i := 0; i < sz; i++ {
		item := rv.Index(i).Interface()
		pyobj, err := py.newPyObject(gate, item)
		if pyobj == nil {
			return nil, err
		}

		err = gate.setListItem(list, pyobj, i)
		gate.unref(pyobj) // Now owned by the list

		if err != nil {
			return nil, err
		}
	}

	gate.ref(list)
	return list, nil
}

// newPyList creates PyObject from the go map.
func (py *Python) newPyDict(gate pyGate, val any) (pyObject, error) {
	rv := reflect.ValueOf(val)

	// Go maps are not ordered, while Python dictionaries
	// are ordered. So sort the map keys, to have reproducible
	// results.
	//
	// If keys cannot be sorted due to their type, map cannot
	// be converted.
	keys := rv.MapKeys()
	ok := reflectSort(keys)
	if !ok {
		err := ErrTypeConversion{
			from: reflect.TypeOf(val).String(),
			to:   "Python dict",
		}
		return nil, err
	}

	// Create a PyDict_Type object
	pydict, err := gate.makeDict()
	if err != nil {
		return nil, err
	}

	// Populate the dictionary
	for _, key := range keys {
		pykey, err := py.newPyObject(gate, key.Interface())
		if err != nil {
			gate.unref(pydict)
			return nil, err
		}

		item := rv.MapIndex(key)
		pyitem, err := py.newPyObject(gate, item.Interface())
		if err != nil {
			gate.unref(pydict)
			gate.unref(pykey)
			return nil, err
		}

		err = gate.setitem(pydict, pykey, pyitem)
		gate.unref(pykey)
		gate.unref(pyitem)

		if err != nil {
			gate.unref(pydict)
			return nil, err
		}
	}

	return pydict, nil
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

// Load loads (imports) string as a Python module with name 'name' as if
// it was loaded from the file 'file'.
//
// On success it returns the module [Object].
func (py *Python) Load(s, name, file string) (*Object, error) {
	gate := py.gate()
	defer gate.release()

	pyobj, err := gate.load(s, name, file)
	if err != nil {
		return nil, err
	}

	obj := newObjectFromPython(py, gate, pyobj)
	return obj, nil
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

	obj := newObjectFromPython(py, gate, pyobj)
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
func (py *Python) countObjID() int {
	return py.objects.count()
}
