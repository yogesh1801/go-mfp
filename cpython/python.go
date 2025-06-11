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

// Eval evaluates string as a Python expression and returns its value.
func (py *Python) Eval(s string) (*Object, error) {
	filename := "<Python.Eval>"
	pc := make([]uintptr, 1)
	if n := runtime.Callers(2, pc); n > 0 {
		frames := runtime.CallersFrames(pc)
		frame, _ := frames.Next()
		filename = fmt.Sprintf("<%s:%d>", frame.File, frame.Line)
	}

	return pyInterpEval(py.interp, s, filename, true)
}

// Exec evaluates string as a Python script.
// The filename parameter specifies the Python source file name
// and used only for diagnostic.
func (py *Python) Exec(s, filename string) (*Object, error) {
	return pyInterpEval(py.interp, s, filename, false)
}
