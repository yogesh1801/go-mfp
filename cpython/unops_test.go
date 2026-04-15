// MFP - Miulti-Function Printers and scanners toolkit
// CPython binding.
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Unary operations on objects tests

package cpython

import (
	"fmt"
	"testing"

	"github.com/OpenPrinting/go-mfp/internal/assert"
)

// TestBinops tests binary operations on Objects
func TestUnops(t *testing.T) {
	type testData struct {
		name string                // Operation name
		in   any                   // Operand
		op   func(*Object) *Object // Operation
		out  any                   // Expected output
		err  string                // Expected error
	}

	tests := []testData{
		{
			name: "-",
			in:   1,
			op:   (*Object).Neg,
			out:  -1,
		},

		{
			name: "-",
			in:   "s",
			op:   (*Object).Neg,
			err:  `AttributeError: 'str' object has no attribute '__neg__'`,
		},

		{
			name: "+",
			in:   1,
			op:   (*Object).Pos,
			out:  1,
		},

		{
			name: "~",
			in:   1,
			op:   (*Object).Invert,
			out:  -2,
		},
	}

	py, err := NewPython()
	assert.NoError(err)
	defer py.Close()

	for _, test := range tests {
		obj := py.NewObject(test.in)
		assert.NoError(obj.Err())

		res := test.op(obj)
		exp := fmt.Sprintf("%v", test.out)
		if test.err != "" {
			exp = test.err
		}

		pres, err := res.Str()
		if err != nil {
			pres = err.Error()
		}

		if exp != pres {
			t.Errorf("%s %v:\n"+
				"expected: %v\n"+
				"present:  %v\n",
				test.name,
				test.in,
				exp, pres)
		}
	}
}
