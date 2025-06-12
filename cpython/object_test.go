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
