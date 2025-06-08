// MFP - Miulti-Function Printers and scanners toolkit
// CPython binding.
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Python objects test

package cpython

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/OpenPrinting/go-mfp/internal/assert"
)

// TestObjectFromPython tests objectFromPython
func TestObjectFromPython(t *testing.T) {
	type testData struct {
		expr     string // Python expression
		val      any    // Expected value
		mustfail bool   // Error expected
	}

	const valbigint = "21267647892944572736998860269687930881"

	bigint := big.NewInt(0)
	bigint.SetString(valbigint, 0)

	tests := []testData{
		{expr: `None`, val: nil},
		{expr: `True`, val: true},
		{expr: `False`, val: false},
		{expr: `"hello"`, val: "hello"},
		{expr: `"привет"`, val: "привет"},
		{expr: `""`, val: ""},
		{expr: `0`, val: 0},
		{expr: `0x7fffffff`, val: 0x7fffffff},
		{expr: `-0x7fffffff`, val: -0x7fffffff},
		{expr: valbigint, val: bigint},
		{expr: `1/0`, mustfail: true},
	}

	py, err := NewPython()
	assert.NoError(err)
	defer py.Close()

	for _, test := range tests {
		obj, err := py.Eval(test.expr)

		// Check errors vs expectations
		if err != nil {
			if !test.mustfail {
				t.Errorf("%s: unexpected error: %s",
					test.expr, err)
			}
			continue
		}

		if test.mustfail {
			t.Errorf("%s: error expected but didn't occur",
				test.expr)
			continue
		}

		// Check returned value
		val := obj.Unbox()
		if !reflect.DeepEqual(val, test.val) {
			t.Errorf("%s: object value mismatch:\n"+
				"expected: %#v\n"+
				"present:  %#v\n",
				test.expr, test.val, val)
		}
	}
}
