// IPPX - High-level implementation of IPP printing protocol on Go
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// IPP codec test

package ippx

import (
	"errors"
	"reflect"
	"testing"

	"github.com/OpenPrinting/goipp"
)

// ippTestStruct is the structure, intended for testing
// of the IPP codec
type ippTestStruct struct {
	FldBooleanF     bool   `ipp:"fld-boolean-f,boolean"`
	FldBooleanSlice []bool `ipp:"fld-boolean-slice,boolean"`
	FldBooleanT     bool   `ipp:"fld-boolean-t,boolean"`

	FldEnum      int   `ipp:"fld-enum,enum"`
	FldEnumSlice []int `ipp:"fld-enum-slice,enum"`

	FldInteger      int   `ipp:"fld-integer,integer"`
	FldIntegerSlice []int `ipp:"fld-integer-slice,integer"`
}

// ----- IPP encode/decode test -----

// ippEncodeDecodeTest represents a single IPP encode/decode test
type ippEncodeDecodeTest struct {
	name  string       // Test name, for logging
	t     reflect.Type // Input type
	data  interface{}  // Input data
	panic error        // Expected panic
	err   error        // Expected error
}

// ippEncodeDecodeTestData is the test data for the IPP encode/decode test
var ippEncodeDecodeTestData = []ippEncodeDecodeTest{
	{
		name:  "panic expected: ippCodecGenerate() with invalid type",
		t:     reflect.TypeOf(int(0)),
		data:  1,
		panic: errors.New(`int: is not struct`),
	},
	{
		name:  "panic expected: ippCodec applied to wrong type",
		t:     reflect.TypeOf(PrinterAttributes{}),
		data:  1,
		panic: errors.New(`Encoder for "*PrinterAttributes" applied to "int"`),
	},
	{
		name: "success expected",
		t:    reflect.TypeOf(PrinterAttributes{}),
		data: &testdataPrinterAttributes,
	},
}

func (test ippEncodeDecodeTest) exec(t *testing.T) {
	// This function catches the possible panic
	defer func() {
		// Panic not expected - let it go its way
		if test.panic == nil {
			return
		}

		p := recover()
		if p == nil && test.panic != nil {
			t.Errorf("in test %q:", test.name)
			t.Errorf("panic expected but didn't happen: %s",
				test.panic)
			return
		}

		if p != nil {
			err, ok := p.(error)
			if !ok {
				panic(p)
			}

			if err.Error() != test.panic.Error() {
				t.Errorf("in test %q:", test.name)
				t.Errorf("panic expected: %s, got: %s",
					test.panic, err)
			}
		}
	}()

	// Generate codec
	codec := ippCodecMustGenerate(test.t)

	// Test encoding
	var attrs goipp.Attributes
	codec.encode(test.data, &attrs)

	// Test decoding
	out := reflect.New(test.t).Interface()
	err := codec.decode(out, attrs)

	checkError(t, test.name, err, test.err)
	if err != nil {
		return
	}

	if !reflect.DeepEqual(test.data, out) {
		t.Errorf("in test %q:", test.name)
		t.Errorf("input/output mismatch")
	}
}

// IPP encode/decode test
func TestIppEncodeDecode(t *testing.T) {
	for _, test := range ippEncodeDecodeTestData {
		test.exec(t)
	}
}

// ----- IPP decode test -----

type ippDecodeTest struct {
	name  string
	t     reflect.Type
	err   error
	attrs goipp.Attributes
	data  interface{}
}

var ippDecodeTestData = []ippDecodeTest{
	{
		name: "success expected",
		t:    reflect.TypeOf(ippTestStruct{}),
		attrs: goipp.Attributes{
			goipp.MakeAttribute("fld-boolean-f",
				goipp.TagBoolean, goipp.Boolean(false)),
			goipp.MakeAttribute("fld-boolean-t",
				goipp.TagBoolean, goipp.Boolean(true)),
			goipp.Attribute{
				Name: "fld-boolean-slice",
				Values: goipp.Values{
					{goipp.TagBoolean, goipp.Boolean(true)},
					{goipp.TagBoolean, goipp.Boolean(false)},
				},
			},

			goipp.MakeAttribute("fld-enum",
				goipp.TagEnum, goipp.Integer(4321)),
			goipp.Attribute{
				Name: "fld-enum-slice",
				Values: goipp.Values{
					{goipp.TagEnum, goipp.Integer(3)},
					{goipp.TagEnum, goipp.Integer(2)},
					{goipp.TagEnum, goipp.Integer(1)},
				},
			},

			goipp.MakeAttribute("fld-integer",
				goipp.TagInteger, goipp.Integer(1234)),
			goipp.Attribute{
				Name: "fld-integer-slice",
				Values: goipp.Values{
					{goipp.TagInteger, goipp.Integer(1)},
					{goipp.TagInteger, goipp.Integer(2)},
					{goipp.TagInteger, goipp.Integer(3)},
				},
			},
		},
		data: &ippTestStruct{
			FldBooleanF:     false,
			FldBooleanT:     true,
			FldBooleanSlice: []bool{true, false},

			FldEnum:      4321,
			FldEnumSlice: []int{3, 2, 1},

			FldInteger:      1234,
			FldIntegerSlice: []int{1, 2, 3},
		},
	},
	{
		name: "success expected",
		t:    reflect.TypeOf(PrinterAttributes{}),
		attrs: goipp.Attributes{
			goipp.Attribute{
				Name: "charset-configured",
				Values: goipp.Values{
					{
						goipp.TagString,
						goipp.String("utf-8"),
					},
				},
			},
		},
		data: &PrinterAttributes{
			CharsetConfigured: DefaultCharsetConfigured,
		},
	},
	{
		name: "string field: Integer passed",
		t:    reflect.TypeOf(PrinterAttributes{}),
		err:  errors.New(`IPP decode PrinterAttributes: "charset-configured": can't convert Integer to String`),
		attrs: goipp.Attributes{
			goipp.Attribute{
				Name: "charset-configured",
				Values: goipp.Values{
					{
						goipp.TagInteger,
						goipp.Integer(0),
					},
				},
			},
		},
	},
	{
		name: "string field: no values passed",
		t:    reflect.TypeOf(PrinterAttributes{}),
		err:  errors.New(`IPP decode PrinterAttributes: "charset-configured": at least 1 value required`),
		attrs: goipp.Attributes{
			goipp.Attribute{
				Name: "charset-configured",
			},
		},
	},
	{
		name: "[]string field: Integer passed",
		t:    reflect.TypeOf(PrinterAttributes{}),
		err:  errors.New(`IPP decode PrinterAttributes: "charset-supported": can't convert Integer to String`),
		attrs: goipp.Attributes{
			goipp.Attribute{
				Name: "charset-supported",
				Values: goipp.Values{
					{
						goipp.TagInteger,
						goipp.Integer(0),
					},
				},
			},
		},
	},
}

func (test ippDecodeTest) exec(t *testing.T) {
	codec := ippCodecMustGenerate(test.t)

	out := reflect.New(test.t).Interface()
	err := codec.decode(out, test.attrs)

	checkError(t, test.name, err, test.err)
	if err != nil {
		return
	}

	if !reflect.DeepEqual(test.data, out) {
		t.Errorf("in test %q:", test.name)
		t.Errorf("input/output mismatch")
	}
}

func TestIppDecode(t *testing.T) {
	for _, test := range ippDecodeTestData {
		test.exec(t)
	}
}

// ----- Common stuff -----

// Check error against expected
func checkError(t *testing.T, name string, err, expected error) {
	switch {
	case err == nil && expected != nil:
		t.Errorf("in test %q:", name)
		t.Errorf("error expected but didn't happen: %s", expected)
	case err != nil && expected == nil:
		t.Errorf("in test %q:", name)
		t.Errorf("error not expected: %s", err)
	case err != nil && expected != nil && err.Error() != expected.Error():
		t.Errorf("in test %q:", name)
		t.Errorf("error expected: %s, got: %s", expected, err)
	}
}

var testdataPrinterAttributes = PrinterAttributes{
	CharsetConfigured:    DefaultCharsetConfigured,
	CharsetSupported:     DefaultCharsetSupported,
	CompressionSupported: []string{"none"},
	IppFeaturesSupported: []string{
		"airprint-1.7",
		"airprint-1.6",
		"airprint-1.5",
		"airprint-1.4",
	},
	IppVersionsSupported: DefaultIppVersionsSupported,
	MediaSizeSupported: []PrinterMediaSizeSupported{
		{21590, 27940},
		{21000, 29700},
	},
	MediaSizeSupportedRange: PrinterMediaSizeSupportedRange{
		XDimension: goipp.Range{Lower: 10000, Upper: 14800},
		YDimension: goipp.Range{Lower: 21600, Upper: 35600},
	},
	OperationsSupported: []goipp.Op{
		goipp.OpGetPrinterAttributes,
	},
}
