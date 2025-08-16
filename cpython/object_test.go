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
	"github.com/OpenPrinting/go-mfp/internal/testutils"
)

// TestObjectFromPython tests objectFromPython
func TestObjectFromPython(t *testing.T) {
	type testData struct {
		expr     string                     // Python expression
		val      any                        // Expected value
		unbox    func(*Object) (any, error) // Value unboxing function
		mustfail bool                       // Error expected
	}

	const verybig = "21267647892944572736998860269687930881"

	bigint := func(s string) *big.Int {
		v := big.NewInt(0)
		v.SetString(s, 0)
		return v
	}

	unboxNone := func(obj *Object) (any, error) { return obj.IsNone(), nil }
	unboxBigint := func(obj *Object) (any, error) { return obj.Bigint() }
	unboxBool := func(obj *Object) (any, error) { return obj.Bool() }
	unboxBytes := func(obj *Object) (any, error) { return obj.Bytes() }
	unboxComplex := func(obj *Object) (any, error) { return obj.Complex() }
	unboxFloat := func(obj *Object) (any, error) { return obj.Float() }
	unboxInt := func(obj *Object) (any, error) { return obj.Int() }
	unboxUint := func(obj *Object) (any, error) { return obj.Uint() }
	unboxUnicode := func(obj *Object) (any, error) { return obj.Unicode() }

	tests := []testData{
		{expr: `None`, val: true, unbox: unboxNone},
		{expr: `True`, val: true, unbox: unboxBool},
		{expr: `0`, unbox: unboxBool, mustfail: true},
		{expr: `False`, val: false, unbox: unboxBool},
		{expr: `"hello"`, val: "hello", unbox: unboxUnicode},
		{expr: `"привет"`, val: "привет", unbox: unboxUnicode},
		{expr: `0`, val: "привет", unbox: unboxUnicode, mustfail: true},
		{expr: `""`, val: "", unbox: unboxUnicode},
		{expr: `0`, val: int64(0), unbox: unboxInt},
		{expr: `0`, val: uint64(0), unbox: unboxUint},
		{expr: `0x7fffffff`, val: int64(0x7fffffff), unbox: unboxInt},
		{expr: `0x7fffffff`, val: uint64(0x7fffffff), unbox: unboxUint},
		{expr: `-0x7fffffff`, val: int64(-0x7fffffff), unbox: unboxInt},
		{expr: `-0x7fffffff`, unbox: unboxUint, mustfail: true},
		{expr: `0xffffffff`, val: int64(0xffffffff), unbox: unboxInt},
		{expr: `0xffffffff`, val: uint64(0xffffffff), unbox: unboxUint},
		{expr: `0xffffffffffffffff`, val: uint64(0xffffffffffffffff), unbox: unboxUint},
		{expr: verybig, unbox: unboxInt, mustfail: true},
		{expr: verybig, unbox: unboxUint, mustfail: true},
		{expr: `0`, val: bigint("0"), unbox: unboxBigint},
		{expr: `1`, val: bigint("1"), unbox: unboxBigint},
		{expr: `-1`, val: bigint("-1"), unbox: unboxBigint},
		{expr: `0xffffffffffffffff`, val: bigint("0xffffffffffffffff"),
			unbox: unboxBigint},
		{expr: verybig, val: bigint(verybig), unbox: unboxBigint},
		{expr: `None`, unbox: unboxInt, mustfail: true},
		{expr: `None`, unbox: unboxUint, mustfail: true},
		{expr: `None`, unbox: unboxBigint, mustfail: true},
		{expr: `b'\x01\x02\x03'`, val: []byte{0x1, 0x2, 0x3},
			unbox: unboxBytes},
		{expr: `bytearray(b'\x01\x02\x03')`, val: []byte{0x1, 0x2, 0x3},
			unbox: unboxBytes},
		{expr: `None`, unbox: unboxBytes, mustfail: true},
		{expr: `0.5`, val: 0.5, unbox: unboxFloat},
		{expr: `None`, unbox: unboxFloat, mustfail: true},
		{expr: `complex(1,2)`, val: complex(1, 2), unbox: unboxComplex},
		{expr: `None`, unbox: unboxComplex, mustfail: true},

		// Corner cases for integers
		{expr: `-1`, val: int64(-1), unbox: unboxInt},
		{expr: `-9223372036854775808`, val: int64(-9223372036854775808), unbox: unboxInt},
		{expr: `9223372036854775807`, val: int64(9223372036854775807), unbox: unboxInt},
		{expr: `0xffffffffffffffff`, val: uint64(0xffffffffffffffff), unbox: unboxUint},
		{expr: `18446744073709551615`, val: uint64(18446744073709551615), unbox: unboxUint},

		// Numerical conversions -> complex
		{expr: `0`, val: complex(0, 0), unbox: unboxComplex},
		{expr: `1`, val: complex(1, 0), unbox: unboxComplex},
		{expr: `1.5`, val: complex(1.5, 0), unbox: unboxComplex},
		{expr: `"hello"`, unbox: unboxComplex, mustfail: true},

		// Numerical conversions -> float
		{expr: `0`, val: 0.0, unbox: unboxFloat},
		{expr: `1`, val: 1.0, unbox: unboxFloat},
		{expr: `1.5`, val: 1.5, unbox: unboxFloat},
		{expr: `0xffffffffffffffff`, val: 18446744073709551615., unbox: unboxFloat},
		{expr: `0xfffffffffffffffff`, unbox: unboxFloat, mustfail: true},
		{expr: `"hello"`, unbox: unboxFloat, mustfail: true},

		// Numerical conversions -> int64
		{expr: `0`, val: int64(0), unbox: unboxInt},
		{expr: `0.0`, val: int64(0), unbox: unboxInt},
		{expr: `2.0`, val: int64(2), unbox: unboxInt},
		{expr: `1.5`, val: int64(1), unbox: unboxInt},
		{expr: `-1.5`, val: int64(-1), unbox: unboxInt},
		{expr: fmt.Sprintf("%f", maxInt64Float), val: int64(maxInt64Float), unbox: unboxInt},
		{expr: fmt.Sprintf("%f", minInt64Float), val: int64(minInt64Float), unbox: unboxInt},
		{expr: fmt.Sprintf("%f", maxUint64Float), mustfail: true, unbox: unboxInt},

		// Numerical conversions -> uint64
		{expr: `0`, val: uint64(0), unbox: unboxUint},
		{expr: `0.0`, val: uint64(0), unbox: unboxUint},
		{expr: `2.0`, val: uint64(2), unbox: unboxUint},
		{expr: `1.5`, val: uint64(1), unbox: unboxUint},
		{expr: `-1.5`, mustfail: true, unbox: unboxUint},
		{expr: fmt.Sprintf("%f", maxInt64Float), val: uint64(maxInt64Float), unbox: unboxUint},
		{expr: fmt.Sprintf("%f", minInt64Float), mustfail: true, unbox: unboxUint},
		{expr: fmt.Sprintf("%f", maxUint64Float), val: uint64(maxUint64Float), unbox: unboxUint},
		{expr: fmt.Sprintf("%f", maxUint64Float+2000), mustfail: true, unbox: unboxUint},
	}

	py, err := NewPython()
	assert.NoError(err)
	defer py.Close()

	for _, test := range tests {
		obj, err := py.Eval(test.expr)

		// Check errors vs expectations
		if err != nil {
			t.Errorf("%s: unexpected error: %s",
				test.expr, err)
			continue
		}

		// Check returned value
		val, err := test.unbox(obj)
		if err != nil {
			if !test.mustfail {
				t.Errorf("%s: object value error: %s",
					test.expr, err)
			}
			continue
		}

		if test.mustfail {
			t.Errorf("%s: error expected but didn't occur",
				test.expr)
			continue
		}

		if !reflect.DeepEqual(val, test.val) {
			t.Errorf("%s: object value mismatch:\n"+
				"expected: %#v\n"+
				"present:  %#v\n",
				test.expr, test.val, val)
		}
	}
}

// TestObjectFromPython tests Object.IsDict and Objet.IsSeq functions
func TestObjectIsMapSeq(t *testing.T) {
	type testData struct {
		expr        string
		isbool      bool
		isbytearray bool
		isbytes     bool
		iscallable  bool
		iscomplex   bool
		isdict      bool
		isfloat     bool
		islong      bool
		isnone      bool
		isseq       bool
		isunicode   bool
	}

	tests := []testData{
		{expr: "True", isbool: true},
		{expr: "False", isbool: true},
		{expr: `bytearray(b'\x01\x02\x03')`, isbytearray: true},
		{expr: `bytes(b'\x01\x02\x03')`, isbytes: true},
		{expr: `min`, iscallable: true},
		{expr: `0.5 + 0.25j`, iscomplex: true},
		{expr: `{}`, isdict: true},
		{expr: `0.5`, isfloat: true},
		{expr: `5`, islong: true},
		{expr: `None`, isnone: true},
		{expr: `[]`, isseq: true},
		{expr: `()`, isseq: true},
		{expr: `"hello"`, isunicode: true},
	}

	py, err := NewPython()
	assert.NoError(err)
	defer py.Close()

	for _, test := range tests {
		obj, err := py.Eval(test.expr)
		assert.NoError(err)

		isbool := obj.IsBool()
		if isbool != test.isbool {
			t.Errorf("%#v: Object.IsBool:\n"+
				"expected: %v\n"+
				"present:  %v\n",
				test.expr, test.isbool, isbool)
		}

		isbytearray := obj.IsByteArray()
		if isbytearray != test.isbytearray {
			t.Errorf("%#v: Object.IsByteArray(:\n"+
				"expected: %v\n"+
				"present:  %v\n",
				test.expr, test.isbytearray, isbytearray)
		}

		isbytes := obj.IsBytes()
		if isbytes != test.isbytes {
			t.Errorf("%#v: Object.IsBytes:\n"+
				"expected: %v\n"+
				"present:  %v\n",
				test.expr, test.isbytes, isbytes)
		}

		iscallable := obj.IsCallable()
		if iscallable != test.iscallable {
			t.Errorf("%#v: Object.IsCallable:\n"+
				"expected: %v\n"+
				"present:  %v\n",
				test.expr, test.iscallable, iscallable)
		}

		iscomplex := obj.IsComplex()
		if iscomplex != test.iscomplex {
			t.Errorf("%#v: Object.IsComplex:\n"+
				"expected: %v\n"+
				"present:  %v\n",
				test.expr, test.iscomplex, iscomplex)
		}

		isdict := obj.IsDict()
		if isdict != test.isdict {
			t.Errorf("%#v: Object.IsDict:\n"+
				"expected: %v\n"+
				"present:  %v\n",
				test.expr, test.isdict, isdict)
		}

		isfloat := obj.IsFloat()
		if isfloat != test.isfloat {
			t.Errorf("%#v: Object.IsFloat:\n"+
				"expected: %v\n"+
				"present:  %v\n",
				test.expr, test.isfloat, isfloat)
		}

		islong := obj.IsLong()
		if islong != test.islong {
			t.Errorf("%#v: Object.IsLong:\n"+
				"expected: %v\n"+
				"present:  %v\n",
				test.expr, test.islong, islong)
		}

		isnone := obj.IsNone()
		if isnone != test.isnone {
			t.Errorf("%#v: Object.IsNone:\n"+
				"expected: %v\n"+
				"present:  %v\n",
				test.expr, test.isnone, isnone)
		}

		isseq := obj.IsSeq()
		if isseq != test.isseq {
			t.Errorf("%#v: Object.IsSeq:\n"+
				"expected: %v\n"+
				"present:  %v\n",
				test.expr, test.isseq, isseq)
		}

		isunicode := obj.IsUnicode()
		if isunicode != test.isunicode {
			t.Errorf("%#v: Object.IsuNicode:\n"+
				"expected: %v\n"+
				"present:  %v\n",
				test.expr, test.isunicode, isunicode)
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

		// Slices and arrays
		{in: []int{1, 2, 3}, out: `[1, 2, 3]`},
		{in: [3]int{1, 2, 3}, out: `[1, 2, 3]`},
		{in: []byte("ABC"), out: `b'ABC'`},
		{in: [3]byte{'A', 'B', 'C'}, out: `b'ABC'`},
		{in: []byte{}, out: `b''`},
		{in: [0]byte{}, out: `b''`},

		// Maps
		{in: map[int]int{}, out: `{}`},
		{in: map[bool]string{true: "T", false: "F"},
			out: `{False: 'F', True: 'T'}`},
		{in: map[int]string{1: "1", 2: "2", 3: "3"},
			out: `{1: '1', 2: '2', 3: '3'}`},
		{in: map[uint]string{1: "1", 2: "2", 3: "3"},
			out: `{1: '1', 2: '2', 3: '3'}`},
		{in: map[float64]string{0.25: "1/4", 0.5: "1/2"},
			out: `{0.25: '1/4', 0.5: '1/2'}`},
		{in: map[string]int{"one": 1, "two": 2, "three": 3},
			out: `{'one': 1, 'three': 3, 'two': 2}`},
		{in: map[string]int{"раз": 1, "два": 2, "три": 3},
			out: `{'два': 2, 'раз': 1, 'три': 3}`},
		{in: map[int]int(nil), out: `{}`},
		{in: map[any]any(nil), out: `{}`},
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

			fmt.Printf("%v\n", obj)
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
		s, err := attr.Unicode()
		assert.NoError(err)
		if s != "Archi" {
			t.Errorf("obj.GetAttr mismatch:\n"+
				"expected: %s\n"+
				"present:  %s\n", "Archi", s)
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

// TestObjectItems tests operations with object items.
func TestObjectItems(t *testing.T) {
	// Create an interpreter
	py, err := NewPython()
	assert.NoError(err)
	defer py.Close()

	// Create test object
	obj, err := py.Eval(`{1:"1", 2:"2", 3:"3"}`)
	assert.NoError(err)
	_ = obj

	// Items 1, 2, 3 must exist and have proper value
	for i := 1; i <= 3; i++ {
		found, err := obj.Contains(i)
		if err != nil {
			t.Errorf("Object.Contains(%v): %s", i, err)
			return
		}

		if !found {
			t.Errorf("Object.Contains(%v): item not found", i)
		}

		item, err := obj.Get(i)
		if err != nil {
			t.Errorf("Object.Get(%v): %s", i, err)
			return
		}

		s, err := item.Str()
		assert.NoError(err)

		if expected := fmt.Sprintf("%d", i); s != expected {
			t.Errorf("Object.Get(%v):\n"+
				"expected: %s\n"+
				"present:  %s\n",
				i, expected, s)
		}
	}

	// Items 4, 5, 6 must not exist.
	// Object.Get must return (nil,nil) for them.
	for i := 4; i <= 6; i++ {
		found, err := obj.Contains(i)
		if err != nil {
			t.Errorf("Object.Contains(%v): %s", i, err)
			return
		}

		if found {
			t.Errorf("Object.Contains(%v): item found", i)
		}

		item, err := obj.Get(i)
		if item != nil || err != nil {
			t.Errorf("Object.Get(%v):\n"+
				"expected: (%v, %v)\n"+
				"present:  (%v, %v)\n",
				i, nil, nil, item, err)
		}
	}

	// Add items 4, 5, 6
	for i := 4; i <= 6; i++ {
		val := fmt.Sprintf("%d", i)
		err := obj.Set(i, val)
		if err != nil {
			t.Errorf("Object.Set(%v): %s", i, err)
			return
		}
	}

	// Now all objects must be present
	for i := 1; i <= 6; i++ {
		item, err := obj.Get(i)
		if err != nil {
			t.Errorf("Object.Get(%v): %s", i, err)
			return
		}

		s, err := item.Str()
		assert.NoError(err)

		if expected := fmt.Sprintf("%d", i); s != expected {
			t.Errorf("Object.Get(%v):\n"+
				"expected: %s\n"+
				"present:  %s\n",
				i, expected, s)
		}
	}

	// Delete items 1, 2, 3
	for i := 1; i <= 3; i++ {
		// For the first time item must be found
		found, err := obj.Del(i)
		if err != nil {
			t.Errorf("Object.Del(%v): %s", i, err)
			return
		}

		if !found {
			t.Errorf("Object.Del(%v): item not found", i)
		}

		// And for the second time it must not be found
		found, err = obj.Del(i)
		if err != nil {
			t.Errorf("Object.Del(%v): %s", i, err)
			return
		}

		if found {
			t.Errorf("Object.Del(%v): item found", i)
		}
	}

	// Now items 1, 2, 3 must not exist.
	for i := 1; i <= 3; i++ {
		found, err := obj.Contains(i)
		if err != nil {
			t.Errorf("Object.Contains(%v): %s", i, err)
			return
		}

		if found {
			t.Errorf("Object.Contains(%v): item found", i)
		}
	}
}

// TestObjectLen tests Object.Len operation
func TestObjectLen(t *testing.T) {
	type testData struct {
		expr string // Python expression
		l    int    // Expected length
		err  bool   // Error expected
	}

	tests := []testData{
		{expr: `[1,2,3]`, l: 3},
		{expr: `(1,2,3)`, l: 3},
		{expr: `{1:"1", 2:"2", 3:"3"}`, l: 3},
		{expr: `[]`, l: 0},
		{expr: `()`, l: 0},
		{expr: `{}`, l: 0},
		{expr: `5`, l: 0, err: true},
		{expr: `"hello"`, l: 5},
		{expr: `"привет"`, l: 6},
	}

	// Create an interpreter
	py, err := NewPython()
	assert.NoError(err)
	defer py.Close()

	// Run tests
	for _, test := range tests {
		obj, err := py.Eval(test.expr)
		assert.NoError(err)

		l, err := obj.Len()
		switch {
		case err == nil && test.err:
			t.Errorf("Object.Len(%s): error not occurred",
				test.expr)
		case err != nil && !test.err:
			t.Errorf("Object.Len(%s): %s",
				test.expr, err)
		case l != test.l:
			t.Errorf("Object.Len(%s):\n"+
				"expected: %d\n"+
				"present:  %d\n",
				test.expr, test.l, l)
		}
	}
}

// TestObjectSlice tests Object.Slice operation
func TestObjectSlice(t *testing.T) {
	type testData struct {
		expr     string   // Python expression
		expected []string // Expected output
		mustfail bool     // Must not succeed
	}

	tests := []testData{
		{
			expr:     "()",
			expected: []string{},
		},

		{
			expr:     "(1,2,3)",
			expected: []string{"1", "2", "3"},
		},

		{
			expr:     "[]",
			expected: []string{},
		},

		{
			expr:     "[1,2,3]",
			expected: []string{"1", "2", "3"},
		},

		{
			expr:     "5",
			mustfail: true,
		},
	}

	// Create an interpreter
	py, err := NewPython()
	assert.NoError(err)
	defer py.Close()

	// Run tests
	for _, test := range tests {
		obj, err := py.Eval(test.expr)
		assert.NoError(err)

		slice, err := obj.Slice()
		if err != nil {
			if !test.mustfail {
				t.Errorf("%s: Object.Slice: %s",
					test.expr, err)
			}
			continue
		}

		if test.mustfail {
			t.Errorf("%s: Object.Slice: expected error didn't occur",
				test.expr)
		}

		result := make([]string, len(slice))
		for i := range slice {
			result[i], err = slice[i].Str()
			assert.NoError(err)
		}

		diff := testutils.Diff(test.expected, result)
		if diff != "" {
			t.Errorf("%s: Object.Slice:\n%s", test.expr, diff)
		}
	}
}

// TestObjectKeys tests Object.Keys operation
func TestObjectKeys(t *testing.T) {
	type testData struct {
		expr     string   // Python expression
		expected []string // Expected output
		mustfail bool     // Must not succeed
	}

	tests := []testData{
		{
			expr:     `{1:"one",2:"two",3:"three"}`,
			expected: []string{"1", "2", "3"},
		},

		{
			expr:     `{}`,
			expected: []string{},
		},

		{
			expr:     `{}`,
			expected: []string{},
		},

		{
			expr:     "()",
			mustfail: true,
		},

		{
			expr:     "5",
			mustfail: true,
		},
	}

	// Create an interpreter
	py, err := NewPython()
	assert.NoError(err)
	defer py.Close()

	// Run tests
	for _, test := range tests {
		obj, err := py.Eval(test.expr)
		assert.NoError(err)

		slice, err := obj.Keys()
		if err != nil {
			if !test.mustfail {
				t.Errorf("%s: Object.Keys: %s",
					test.expr, err)
			}
			continue
		}

		if test.mustfail {
			t.Errorf("%s: Object.Keys: expected error didn't occur",
				test.expr)
		}

		result := make([]string, len(slice))
		for i := range slice {
			result[i], err = slice[i].Str()
			assert.NoError(err)
		}

		diff := testutils.Diff(test.expected, result)
		if diff != "" {
			t.Errorf("%s: Object.Keys:\n%s", test.expr, diff)
		}
	}
}

// TestObjectCallable tests Object.Callable operation
func TestObjectCallable(t *testing.T) {
	// Create an interpreter
	py, err := NewPython()
	assert.NoError(err)
	defer py.Close()

	// Obtain non-callable object
	obj, err := py.Eval("5")
	assert.NoError(err)

	callable := obj.IsCallable()
	if callable {
		t.Errorf("Object.Callable: false positive response")
	}

	// Obtain callable object
	obj, err = py.Eval("min")
	assert.NoError(err)

	callable = obj.IsCallable()
	if !callable {
		t.Errorf("Object.Callable: false negative response")
	}
}

// TestObjectCall tests Object.Call operation
func TestObjectCall(t *testing.T) {
	// Create an interpreter
	py, err := NewPython()
	assert.NoError(err)
	defer py.Close()

	// Obtain callable object
	obj, err := py.Eval("min")
	assert.NoError(err)

	// Perform simple call
	res, err := obj.Call(1, 2)
	if err != nil {
		t.Errorf("Object.Call (positional args): %s", err)
		return
	}

	val, err := res.Int()
	if err != nil {
		t.Errorf("Object.Call (positional args): %s", err)
		return
	}
	if val != 1 {
		t.Errorf("Object.Call (positional args):\n"+
			"expected: %d\n"+
			"present:  %d\n",
			1, val)
		return
	}

	// Call with keyword arguments
	res, err = obj.CallKW(map[string]any{"default": 5}, []int{})
	if err != nil {
		t.Errorf("Object.Call (keyword args): %s", err)
		return
	}

	val, err = res.Int()
	if err != nil {
		t.Errorf("Object.Call (keyword args): %s", err)
		return
	}
	if val != 5 {
		t.Errorf("Object.Call (keyword args):\n"+
			"expected: %d\n"+
			"present:  %d\n",
			5, val)
		return
	}
}

// TestObjectGC tests how objects are garbage-collected
func TestObjectGC(t *testing.T) {
	py, err := NewPython()
	assert.NoError(err)

	base := py.countObjID()

	_, err = py.Eval("5")
	assert.NoError(err)

	if len(py.objects.mapped) != base+1 {
		t.Errorf("TestObjectGC: looks object is not properly mapped")
	}

	runtime.GC()
	runtime.GC()

	if py.countObjID() != base {
		t.Errorf("TestObjectGC: looks object GS doesn't work")
	}
}
