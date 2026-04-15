// MFP - Miulti-Function Printers and scanners toolkit
// CPython binding.
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Binary operations on objects tests

package cpython

import (
	"fmt"
	"testing"

	"github.com/OpenPrinting/go-mfp/internal/assert"
)

// TestBinops tests binary operations on Objects
func TestBinops(t *testing.T) {
	type testData struct {
		name     string                     // Operation name
		in1, in2 any                        // Operands
		op       func(*Object, any) *Object // Operation
		out      any                        // Expected output
		err      string                     // Expected error
	}

	tests := []testData{
		{
			name: "+",
			in1:  1,
			in2:  2,
			op:   (*Object).Add,
			out:  3,
		},

		{
			name: "-",
			in1:  1,
			in2:  2,
			op:   (*Object).Sub,
			out:  -1,
		},

		{
			name: "*",
			in1:  2,
			in2:  3,
			op:   (*Object).Mul,
			out:  6,
		},

		{
			name: "*",
			in1:  "x",
			in2:  5,
			op:   (*Object).Mul,
			out:  "xxxxx",
		},

		{
			name: "*",
			in1:  5,
			in2:  "x",
			op:   (*Object).Mul,
			err:  "NotImplemented",
		},

		{
			name: "/",
			in1:  2,
			in2:  4,
			op:   (*Object).TrueDiv,
			out:  0.5,
		},

		{
			name: "//",
			in1:  2,
			in2:  4,
			op:   (*Object).FloorDiv,
			out:  0,
		},

		{
			name: "%",
			in1:  2,
			in2:  10,
			op:   (*Object).Pow,
			out:  1024,
		},

		{
			name: "<",
			in1:  2,
			in2:  4,
			op:   (*Object).Lt,
			out:  "True",
		},

		{
			name: "<",
			in1:  2,
			in2:  2,
			op:   (*Object).Lt,
			out:  "False",
		},

		{
			name: "<",
			in1:  4,
			in2:  2,
			op:   (*Object).Lt,
			out:  "False",
		},

		{
			name: ">",
			in1:  2,
			in2:  4,
			op:   (*Object).Gt,
			out:  "False",
		},

		{
			name: ">",
			in1:  2,
			in2:  2,
			op:   (*Object).Gt,
			out:  "False",
		},

		{
			name: ">",
			in1:  4,
			in2:  2,
			op:   (*Object).Gt,
			out:  "True",
		},

		{
			name: "<=",
			in1:  2,
			in2:  4,
			op:   (*Object).Le,
			out:  "True",
		},

		{
			name: "<=",
			in1:  2,
			in2:  2,
			op:   (*Object).Le,
			out:  "True",
		},

		{
			name: "<=",
			in1:  4,
			in2:  2,
			op:   (*Object).Le,
			out:  "False",
		},

		{
			name: ">=",
			in1:  2,
			in2:  4,
			op:   (*Object).Ge,
			out:  "False",
		},

		{
			name: ">=",
			in1:  2,
			in2:  2,
			op:   (*Object).Ge,
			out:  "True",
		},

		{
			name: ">=",
			in1:  4,
			in2:  2,
			op:   (*Object).Ge,
			out:  "True",
		},

		{
			name: "==",
			in1:  2,
			in2:  2,
			op:   (*Object).Eq,
			out:  "True",
		},

		{
			name: "==",
			in1:  2,
			in2:  4,
			op:   (*Object).Eq,
			out:  "False",
		},

		{
			name: "!=",
			in1:  2,
			in2:  2,
			op:   (*Object).Ne,
			out:  "False",
		},

		{
			name: "!=",
			in1:  2,
			in2:  4,
			op:   (*Object).Ne,
			out:  "True",
		},
	}

	py, err := NewPython()
	assert.NoError(err)
	defer py.Close()

	for _, test := range tests {
		obj1 := py.NewObject(test.in1)
		assert.NoError(obj1.Err())

		obj2 := py.NewObject(test.in2)
		assert.NoError(obj2.Err())

		res := test.op(obj1, obj2)
		exp := fmt.Sprintf("%v", test.out)
		if test.err != "" {
			exp = test.err
		}

		s, err := res.Str()
		if err != nil || s != exp {
			pres := s
			if err != nil {
				pres = err.Error()
			}

			t.Errorf("%v %s %v:\n"+
				"expected: %v\n"+
				"present:  %v\n",
				test.in1,
				test.name,
				test.in2,
				exp, pres)
		}
	}
}
