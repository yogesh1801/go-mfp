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
	"reflect"
)

// ErrPython represents a Python exception.
type ErrPython struct {
	msg string
}

// Error returns error message. It implements the [error] interface.
func (e ErrPython) Error() string {
	return e.msg
}

// ErrTypeConversion represents Go->Python type conversion error.
type ErrTypeConversion struct {
	from reflect.Type // Go type that can't be converted
}

// Error returns error message. It implements the [error] interface.
func (e ErrTypeConversion) Error() string {
	return fmt.Sprintf("%T cannot be converted to PyObject", e.from)
}
