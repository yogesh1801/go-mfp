// MFP - Miulti-Function Printers and scanners toolkit
// CPython binding.
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Python methods tests

package cpython

import (
	"testing"

	"github.com/OpenPrinting/go-mfp/internal/assert"
)

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
