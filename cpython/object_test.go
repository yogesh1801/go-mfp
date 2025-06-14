// MFP - Miulti-Function Printers and scanners toolkit
// CPython binding.
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Python objects test

package cpython

import (
	"fmt"
	"math"
	"math/big"
	"reflect"
	"runtime"
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

	const verybig = "21267647892944572736998860269687930881"

	bigint := func(s string) *big.Int {
		v := big.NewInt(0)
		v.SetString(s, 0)
		return v
	}

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
		{expr: `0xffffffff`, val: 0xffffffff},
		{expr: `0xffffffffffffffff`, val: bigint("0xffffffffffffffff")},
		{expr: verybig, val: bigint(verybig)},
		{expr: `1/0`, mustfail: true},
		{expr: `b'\x01\x02\x03'`, val: []byte{0x1, 0x2, 0x3}},
		{expr: `bytearray(b'\x01\x02\x03')`, val: []byte{0x1, 0x2, 0x3}},
		{expr: `0.5`, val: 0.5},
		{expr: `complex(1,2)`, val: complex(1, 2)},
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

// TestNewObject tests Python.NewObject
func TestNewObject(t *testing.T) {
	type testData struct {
		in  any    // Input value
		out string // Expected output
	}

	py, err := NewPython()
	assert.NoError(err)
	defer py.Close()

	bigint := func(s string) *big.Int {
		v := big.NewInt(0)
		v.SetString(s, 0)
		return v
	}

	eval := func(s string) *Object {
		obj, err := py.Eval(s)
		assert.NoError(err)
		return obj
	}

	tests := []testData{
		// nil value
		{in: nil, out: `None`},

		// bool
		{in: true, out: `True`},
		{in: false, out: `False`},

		// Signed int
		{in: 0, out: `0`},
		{in: 1, out: `1`},
		{in: -1, out: `-1`},
		{in: int8(0), out: `0`},
		{in: int8(math.MaxInt8), out: fmt.Sprintf("%d", math.MaxInt8)},
		{in: int8(math.MinInt8), out: fmt.Sprintf("%d", math.MinInt8)},
		{in: int16(0), out: `0`},
		{in: int16(math.MaxInt16),
			out: fmt.Sprintf("%d", math.MaxInt16)},
		{in: int16(math.MinInt16),
			out: fmt.Sprintf("%d", math.MinInt16)},
		{in: int32(0), out: `0`},
		{in: int32(math.MaxInt32),
			out: fmt.Sprintf("%d", math.MaxInt32)},
		{in: int32(math.MinInt32),
			out: fmt.Sprintf("%d", math.MinInt32)},
		{in: int64(0), out: `0`},
		{in: int64(math.MaxInt64),
			out: fmt.Sprintf("%d", math.MaxInt64)},
		{in: int64(math.MinInt64),
			out: fmt.Sprintf("%d", math.MinInt64)},

		// Unsigned int
		{in: uint(0), out: `0`},
		{in: uint(1), out: `1`},
		{in: uint8(0), out: `0`},
		{in: uint8(math.MaxUint8),
			out: fmt.Sprintf("%d", math.MaxUint8)},
		{in: uint16(0), out: `0`},
		{in: uint16(math.MaxUint16),
			out: fmt.Sprintf("%d", math.MaxUint16)},
		{in: uint32(0), out: `0`},
		{in: uint32(math.MaxUint32),
			out: fmt.Sprintf("%d", math.MaxUint32)},
		{in: uint64(0), out: `0`},
		{in: uint64(math.MaxUint64),
			out: fmt.Sprintf("%d", uint64(math.MaxUint64))},

		// Big int
		{in: bigint("0"), out: `0`},
		{in: bigint("1"), out: `1`},
		{in: bigint("-1"), out: `-1`},
		{in: bigint("340282366920938463426481119284349108225"),
			out: `340282366920938463426481119284349108225`},
		{in: bigint("-340282366920938463426481119284349108225"),
			out: `-340282366920938463426481119284349108225`},

		// Float
		{in: 0.5, out: `0.5`},
		{in: -0.5, out: `-0.5`},

		// Complex
		{in: 0.25 + 0.25i, out: `(0.25+0.25j)`},
		{in: -0.25 - 0.25i, out: `(-0.25-0.25j)`},

		// String
		{in: "", out: ``},
		{in: "Hello, world!", out: `Hello, world!`},
		{in: "Здравствуй, мир!", out: `Здравствуй, мир!`},

		// Object
		{in: eval("12345"), out: `12345`},
	}

	for _, test := range tests {
		obj, err := py.NewObject(test.in)
		if err != nil {
			t.Errorf("%v: Python.NewObject: %s", test.in, err)
			continue
		}

		s, err := obj.Str()
		if err != nil {
			t.Errorf("%v: Object.Str: %s", test.in, err)
			continue
		}

		if s != test.out {
			t.Errorf("%v: Python.NewObject:\n"+
				"expected: %s\n"+
				"present:  %s\n",
				test.in, test.out, s)
		}
	}
}

// TestObjectAttributes tests operations with object attributes.
func TestObjectAttributes(t *testing.T) {
	script := `
class Dog:
    # Class attribute
    species = "Canis familiaris"

    def __init__(self, name, age):
        # Instance attributes
        self.name = name
        self.age = age

# Creating instances of the Dog class
dog = Dog("Archi", 4)
`

	py, err := NewPython()
	assert.NoError(err)
	defer py.Close()

	err = py.Exec(script, "")
	assert.NoError(err)

	obj, err := py.Eval("dog")
	assert.NoError(err)

	found, err := obj.HasAttr("name")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if !found {
		t.Errorf("Attribute %q expected but not found", "name")
	}

	found, err = obj.HasAttr("unknown")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if found {
		t.Errorf("Attribute not %q expected but found", "unknown")
	}

	attr, err := obj.GetAttr("name")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	} else {
		if attr.Unbox() != "Archi" {
			t.Errorf("obj.GetAttr mismatch:\n"+
				"expected: %s\n"+
				"present:  %s\n", "Archi", attr.Unbox())
		}
	}

	deleted, err := obj.DelAttr("name")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if !deleted {
		t.Errorf("obj.DelAttr: attribute %s not deleted", "name")
	}

	found, err = obj.HasAttr("name")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if found {
		t.Errorf("Deleted attribute %q still present", "name")
	}
}

// TestObjectGC tests how objects are garbage-collected
func TestObjectGC(t *testing.T) {
	py, err := NewPython()
	assert.NoError(err)

	_, err = py.Eval("5")
	assert.NoError(err)

	if len(py.objects.mapped) != 1 {
		t.Errorf("TestObjectGC: looks object is not properly mapped")
	}

	runtime.GC()
	runtime.GC()

	if len(py.objects.mapped) != 0 {
		t.Errorf("TestObjectGC: looks object GS doesn't work")
	}
}
