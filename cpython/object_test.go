// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Python objects test

package cpython

import (
	"testing"

	"github.com/OpenPrinting/go-mfp/internal/assert"
)

// TestObjectFromPython tests objectFromPython
func TestObjectFromPython(t *testing.T) {
	type testData struct {
		expr string // Python expression
		val  any    // Expected value
	}

	tests := []testData{
		{expr: `None`, val: nil},
		{expr: `True`, val: true},
		{expr: `False`, val: false},
		{expr: `"hello"`, val: "hello"},
		{expr: `"привет"`, val: "привет"},
	}

	py, err := NewPython()
	assert.NoError(err)
	defer py.Close()

	for _, test := range tests {
		obj := py.Eval(test.expr)
		val := obj.Unbox()

		if val != test.val {
			t.Errorf("%s: object value mismatch:\n"+
				"expected: %v\n"+
				"present:  %v\n",
				test.expr, test.val, val)
		}
	}
}
