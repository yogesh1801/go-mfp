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
	py, err := NewPython()
	assert.NoError(err)
	py.Eval(`print("hello, world")`)
	py.Close()

	py, err = NewPython()
	assert.NoError(err)
	py.Eval(`print("hello, world")`)
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
