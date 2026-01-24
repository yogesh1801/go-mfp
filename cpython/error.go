// MFP - Miulti-Function Printers and scanners toolkit
// CPython binding.
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Package documentation

package cpython

import (
	"fmt"
)

// ErrPython represents a Python exception.
type ErrPython struct {
	msg string
}

// Error returns error message. It implements the [error] interface.
func (e ErrPython) Error() string {
	return e.msg
}

// ErrTypeConversion represents Go<->Python type conversion error.
type ErrTypeConversion struct {
	from, to string // from/to types that can't be converted
}

// Error returns error message. It implements the [error] interface.
func (e ErrTypeConversion) Error() string {
	return fmt.Sprintf("can't convert %s to %s", e.from, e.to)
}

// ErrOverflow represents the integer overflow error.
type ErrOverflow struct {
	val string
}

// Error returns error message. It implements the [error] interface.
func (e ErrOverflow) Error() string {
	return fmt.Sprintf("integer overflow: %s", e.val)
}

// ErrClosed represent the error that occurs when [Python]
// interpreter or [Object] that it owns accessed after call
// to [Python.Close].
type ErrClosed struct{}

// Error returns error message. It implements the [error] interface.
func (e ErrClosed) Error() string {
	return "use of closed Python interpreter"
}
