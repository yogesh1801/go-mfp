// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Python interpreter.

package cpython

// Python represents a Python interpreter.
// There are may be many interpreters within a single process.
// Each has its own namespace and isolated from others.
type Python struct {
	interp pyInterp
}

// NewPython creates a new Python interpreter.
func NewPython() (py *Python, err error) {
	interp, err := pyNewInterp()
	if err == nil {
		py = &Python{interp}
	}

	return
}

// Close closes the [Python] interpreter and releases all
// resources it holds.
func (py *Python) Close() {
	pyInterpDelete(py.interp)
}

// Eval evaluates string as a Python statement and returns its value.
func (py *Python) Eval(s string) *Object {
	return pyInterpEval(py.interp, s)
}
