// MFP - Miulti-Function Printers and scanners toolkit
// CPython binding.
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Python->Go callbacks test

package cpython

import (
	"fmt"
	"testing"

	"github.com/OpenPrinting/go-mfp/internal/assert"
)

// TestCallback tests calling Go from Python
func TestCallback(t *testing.T) {
	py, err := NewPython()
	assert.NoError(err)
	defer py.Close()

	call := py.NewObject(func() {
		fmt.Println("==== callbackCall ====")
	})
	println(call.Str())
	ret := call.Call(5)
	if err := ret.Err(); err != nil {
		println(err.Error())
	} else {
		println(ret.Str())
	}

}
