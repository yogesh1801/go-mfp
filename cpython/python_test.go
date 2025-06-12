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

		res, err := py.Eval(`"hello"`)
		if err != nil {
			t.Errorf("py.Eval: unexpected error: %s", err)
			py.Close()
			return
		}

		if res.Unbox() != "hello" {
			t.Errorf("Returned value mismatch:\n"+
				"expected: %v\n"+
				"present:  %v\n",
				"hello", res.Unbox())
		}

		py.Close()
	}

	// Basic test for Python.Exec
	py, err := NewPython()
	assert.NoError(err)

	script := "" +
		"x = 5\n" +
		"x *= 2\n"

	err = py.Exec(script, "")
	if err != nil {
		t.Errorf("py.Eval: unexpected error: %s", err)
		py.Close()
		return
	}

	res, err := py.Eval(`x`)
	if err != nil {
		t.Errorf("py.Eval: unexpected error: %s", err)
		py.Close()
		return
	}

	if res.Unbox() != 10 {
		t.Errorf("Returned value mismatch:\n"+
			"expected: %v\n"+
			"present:  %v\n",
			"hello", res.Unbox())
	}

	py.Close()
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
