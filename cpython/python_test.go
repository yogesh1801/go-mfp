// MFP - Miulti-Function Printers and scanners toolkit
// CPython binding.
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Python methods tests

package cpython

import (
	"errors"
	"reflect"
	"strings"
	"testing"

	"github.com/OpenPrinting/go-mfp/internal/assert"
)

// TestPython tests basic functionality of the Python type
func TestPython(t *testing.T) {
	// Create the interpreter, try py.Eval. Do it several times.
	for i := 0; i < 5; i++ {
		py, err := NewPython()
		if err != nil {
			t.Errorf("NewPython: unexpected error: %s", err)
			return
		}

		res := py.Eval(`"hello"`)
		if err := res.Err(); err != nil {
			t.Errorf("py.Eval: unexpected error: %s", err)
			py.Close()
			return
		}

		s, err := res.Unicode()
		if err != nil {
			t.Errorf("Returned value decoding: %s", err)
			return
		}

		if s != "hello" {
			t.Errorf("Returned value mismatch:\n"+
				"expected: %v\n"+
				"present:  %v\n",
				"hello", s)
		}

		py.Close()
	}

	// Basic test for Python.Exec
	py, err := NewPython()
	assert.NoError(err)
	defer py.Close()

	script := "" +
		"x = 5\n" +
		"x *= 2\n"

	err = py.Exec(script, "")
	if err != nil {
		t.Errorf("py.Eval: unexpected error: %s", err)
		return
	}

	res := py.Eval(`x`)
	if err := res.Err(); err != nil {
		t.Errorf("py.Eval: unexpected error: %s", err)
		return
	}

	v, err := res.Int()
	if err != nil {
		t.Errorf("py.Eval: decode error: %s", err)
		return
	}

	if v != 10 {
		t.Errorf("Returned value mismatch:\n"+
			"expected: %v\n"+
			"present:  %v\n",
			"hello", v)
	}
}

// TestPythonInitError tests how Python initialization errors are handled.
func TestPythonInitError(t *testing.T) {
	initerr := errors.New("Initialization error")
	save := pyInitError
	defer func() { pyInitError = save }()

	pyInitError = nil // So we don't depend on a preceding errors
	pyInitErrorCheckTest(initerr.Error())

	py, err := NewPython()
	if !reflect.DeepEqual(err, initerr) {
		t.Errorf("Initialization error handling test failed:\n"+
			"error expected: %v\n"+
			"error present:  %v\n",
			initerr, err)
	}

	if py != nil {
		t.Errorf("Initialization error handling test failed:\n"+
			"Python expected: %v\n"+
			"Python present:  %v\n",
			nil, py)
		py.Close()
	}
}

// TestPythonLoad tests module loading.
func TestPythonLoad(t *testing.T) {
	py, err := NewPython()
	assert.NoError(err)

	mod := "\n" +
		"i = 5\n" +
		""

	err = py.Load(mod, "mymodule", "modulefile.py").Err()
	if err != nil {
		t.Errorf("Python.Import: %s", err)
		return
	}

	obj := py.Eval("mymodule.i")
	if err := obj.Err(); err != nil {
		t.Errorf("Python.Import: can't access module variable: %s", err)
		return
	}

	v, err := obj.Int()
	if err != nil {
		t.Errorf("Python.Import: can't decode module variable: %s", err)
		return
	}

	if v != 5 {
		t.Errorf("Python.Import: module variable value:\n"+
			"expected: %d\n"+
			"presend:  %d\n",
			5, v)
	}
}

// TestPythonErrorLocation tests accuracy of error location reporting
func TestPythonErrorLocation(t *testing.T) {
	// This script must raise an exception in the stdlib.
	// Here we verify that we can properly locate error line
	// that belongs to the our code, not to the stdlib.
	script := "" +
		"import base64\n" +
		"base64.b64encode(5)\n" +
		""

	py, err := NewPython()
	assert.NoError(err)

	err = py.Exec(script, "test.py")
	assert.MustMsg(err != nil, "Python exception MUST occur")
	expected := "(test.py, line 2)"
	present := err.Error()

	if !strings.HasSuffix(present, expected) {
		t.Errorf("invalid error location reported:\n"+
			"expected: %s\n"+
			"present:  %s\n",
			expected, present)
	}
}
