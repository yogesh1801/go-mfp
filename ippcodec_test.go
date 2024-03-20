// IPPX - High-level implementation of IPP printing protocol on Go
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// IPP codec test

package ippx

import (
	"bytes"
	"errors"
	"reflect"
	"sort"
	"testing"
	"time"

	"github.com/OpenPrinting/goipp"
)

// ippTestStruct is the structure, intended for testing
// of the IPP codec
type ippTestStruct struct {
	FldBooleanF     bool   `ipp:"fld-boolean-f,boolean"`
	FldBooleanSlice []bool `ipp:"fld-boolean-slice,boolean"`
	FldBooleanT     bool   `ipp:"fld-boolean-t,boolean"`

	FldCharset      string   `ipp:"fld-charset,charset"`
	FldCharsetSlice []string `ipp:"fld-charset-slice,charset"`

	FldDateTime      time.Time   `ipp:"fld-datetime,datetime"`
	FldDateTimeSlice []time.Time `ipp:"fld-datetime-slice,datetime"`

	FldEnum      int   `ipp:"fld-enum,enum"`
	FldEnumSlice []int `ipp:"fld-enum-slice,enum"`

	FldInteger      int   `ipp:"fld-integer,integer"`
	FldIntegerSlice []int `ipp:"fld-integer-slice,integer"`

	FldKeyword      string   `ipp:"fld-keyword,keyword"`
	FldKeywordSlice []string `ipp:"fld-keyword-slice,keyword"`

	FldLanguage      string   `ipp:"fld-language,naturalLanguage"`
	FldLanguageSlice []string `ipp:"fld-language-slice,naturalLanguage"`

	FldMime      string   `ipp:"fld-mime,mimemediatype"`
	FldMimeSlice []string `ipp:"fld-mime-slice,mimemediatype"`

	FldName      string   `ipp:"fld-name,name"`
	FldNameSlice []string `ipp:"fld-name-slice,name"`

	FldString      string   `ipp:"fld-string,string"`
	FldStringSlice []string `ipp:"fld-string-slice,string"`

	FldText      string   `ipp:"fld-text,text"`
	FldTextSlice []string `ipp:"fld-text-slice,text"`

	FldURI      string   `ipp:"fld-uri,uri"`
	FldURISlice []string `ipp:"fld-uri-slice,uri"`

	FldURIScheme      string   `ipp:"fld-urischeme,urischeme"`
	FldURISchemeSlice []string `ipp:"fld-urischeme-slice,urischeme"`

	FldVersion      goipp.Version   `ipp:"fld-version"`
	FldVersionSlice []goipp.Version `ipp:"fld-version-slice"`
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
		panic: errors.New(`Encoder for "*ippx.PrinterAttributes" applied to "int"`),
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
		t.Errorf("expected: %#v\n", test.data)
		t.Errorf("present: %#v\n", out)
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
	name       string
	t          reflect.Type
	err        error
	attrs      goipp.Attributes
	data       interface{}
	skipEncode bool
}

var ippDecodeTestData = []ippDecodeTest{
	{
		name: "ippTestStruct: success expected",
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

			goipp.MakeAttribute("fld-charset",
				goipp.TagCharset, goipp.String("utf-8")),
			goipp.Attribute{
				Name: "fld-charset-slice",
				Values: goipp.Values{
					{goipp.TagCharset, goipp.String("ibm866")},
					{goipp.TagCharset, goipp.String("iso-8859-5")},
					{goipp.TagCharset, goipp.String("windows-1251")},
				},
			},

			goipp.MakeAttribute("fld-datetime",
				goipp.TagDateTime, goipp.Time{Time: testTime1}),
			goipp.Attribute{
				Name: "fld-datetime-slice",
				Values: goipp.Values{
					{goipp.TagDateTime, goipp.Time{Time: testTime2}},
					{goipp.TagDateTime, goipp.Time{Time: testTime3}},
					{goipp.TagDateTime, goipp.Time{Time: testTime4}},
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

			goipp.MakeAttribute("fld-keyword",
				goipp.TagKeyword, goipp.String("document")),
			goipp.Attribute{
				Name: "fld-keyword-slice",
				Values: goipp.Values{
					{goipp.TagKeyword, goipp.String("one-sided")},
					{goipp.TagKeyword, goipp.String("two-sided-short-edge")},
					{goipp.TagKeyword, goipp.String("two-sided-long-edge")},
				},
			},

			goipp.MakeAttribute("fld-language",
				goipp.TagLanguage, goipp.String("en-US")),
			goipp.Attribute{
				Name: "fld-language-slice",
				Values: goipp.Values{
					{goipp.TagLanguage, goipp.String("be-BY")},
					{goipp.TagLanguage, goipp.String("ru-RU")},
					{goipp.TagLanguage, goipp.String("uk-UA")},
				},
			},

			goipp.MakeAttribute("fld-mime",
				goipp.TagMimeType, goipp.String("application/pdf")),
			goipp.Attribute{
				Name: "fld-mime-slice",
				Values: goipp.Values{
					{goipp.TagMimeType, goipp.String("image/tiff")},
					{goipp.TagMimeType, goipp.String("image/jpeg")},
					{goipp.TagMimeType, goipp.String("image/urf")},
				},
			},

			goipp.MakeAttribute("fld-name",
				goipp.TagName, goipp.String("Printer in a classroom")),
			goipp.Attribute{
				Name: "fld-name-slice",
				Values: goipp.Values{
					{goipp.TagName, goipp.String("Job0001")},
					{goipp.TagName, goipp.String("Job0002")},
					{goipp.TagName, goipp.String("Job0003")},
				},
			},

			goipp.MakeAttribute("fld-string",
				goipp.TagString, goipp.String("hello, world")),
			goipp.Attribute{
				Name: "fld-string-slice",
				Values: goipp.Values{
					{goipp.TagString, goipp.String("A")},
					{goipp.TagString, goipp.String("B")},
					{goipp.TagString, goipp.String("C")},
				},
			},

			goipp.MakeAttribute("fld-text",
				goipp.TagText, goipp.String("ping pong")),
			goipp.Attribute{
				Name: "fld-text-slice",
				Values: goipp.Values{
					{goipp.TagText, goipp.String("X")},
					{goipp.TagText, goipp.String("Y")},
					{goipp.TagText, goipp.String("Z")},
				},
			},

			goipp.MakeAttribute("fld-uri",
				goipp.TagURI, goipp.String("http://example.com")),
			goipp.Attribute{
				Name: "fld-uri-slice",
				Values: goipp.Values{
					{goipp.TagURI, goipp.String("http://example.com/print")},
					{goipp.TagURI, goipp.String("http://example.com/scan")},
				},
			},

			goipp.MakeAttribute("fld-urischeme",
				goipp.TagURIScheme, goipp.String("http")),
			goipp.Attribute{
				Name: "fld-urischeme-slice",
				Values: goipp.Values{
					{goipp.TagURIScheme, goipp.String("tel")},
					{goipp.TagURIScheme, goipp.String("mailto")},
				},
			},

			goipp.MakeAttribute("fld-version",
				goipp.TagKeyword, goipp.String("2.0")),
			goipp.Attribute{
				Name: "fld-version-slice",
				Values: goipp.Values{
					{goipp.TagKeyword, goipp.String("2.0")},
					{goipp.TagKeyword, goipp.String("1.1")},
					{goipp.TagKeyword, goipp.String("1.0")},
				},
			},
		},
		data: &ippTestStruct{
			FldBooleanF:     false,
			FldBooleanT:     true,
			FldBooleanSlice: []bool{true, false},

			FldCharset:      "utf-8",
			FldCharsetSlice: []string{"ibm866", "iso-8859-5", "windows-1251"},

			FldDateTime: testTime1,
			FldDateTimeSlice: []time.Time{
				testTime2, testTime3, testTime4,
			},

			FldEnum:      4321,
			FldEnumSlice: []int{3, 2, 1},

			FldInteger:      1234,
			FldIntegerSlice: []int{1, 2, 3},

			FldKeyword: "document",
			FldKeywordSlice: []string{
				"one-sided",
				"two-sided-short-edge",
				"two-sided-long-edge"},

			FldLanguage:      "en-US",
			FldLanguageSlice: []string{"be-BY", "ru-RU", "uk-UA"},

			FldMime: "application/pdf",
			FldMimeSlice: []string{
				"image/tiff",
				"image/jpeg",
				"image/urf"},

			FldName: "Printer in a classroom",
			FldNameSlice: []string{
				"Job0001",
				"Job0002",
				"Job0003"},

			FldString:      "hello, world",
			FldStringSlice: []string{"A", "B", "C"},

			FldText:      "ping pong",
			FldTextSlice: []string{"X", "Y", "Z"},

			FldURI: "http://example.com",
			FldURISlice: []string{
				"http://example.com/print",
				"http://example.com/scan"},

			FldURIScheme:      "http",
			FldURISchemeSlice: []string{"tel", "mailto"},

			FldVersion: goipp.MakeVersion(2, 0),
			FldVersionSlice: []goipp.Version{
				goipp.MakeVersion(2, 0),
				goipp.MakeVersion(1, 1),
				goipp.MakeVersion(1, 0),
			},
		},
	},
	{
		name:       "PrinterAttributes: success expected",
		skipEncode: true,
		t:          reflect.TypeOf(PrinterAttributes{}),
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
		err:  errors.New(`IPP decode ippx.PrinterAttributes: "charset-configured": can't convert Integer to String`),
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
		err:  errors.New(`IPP decode ippx.PrinterAttributes: "charset-configured": at least 1 value required`),
		attrs: goipp.Attributes{
			goipp.Attribute{
				Name: "charset-configured",
			},
		},
	},
	{
		name: "[]string field: Integer passed",
		t:    reflect.TypeOf(PrinterAttributes{}),
		err:  errors.New(`IPP decode ippx.PrinterAttributes: "charset-supported": can't convert Integer to String`),
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
	// Compile the codec
	codec := ippCodecMustGenerate(test.t)

	// Decode IPP attributes
	out := reflect.New(test.t).Interface()
	err := codec.decode(out, test.attrs)

	checkError(t, test.name, err, test.err)
	if err != nil {
		return
	}

	// Compare result against expected
	if !reflect.DeepEqual(test.data, out) {
		t.Errorf("in test %q:", test.name)
		t.Errorf("decode: input/output mismatch")
		t.Errorf("expected: %#v\n", test.data)
		t.Errorf("present: %#v\n", out)
		return
	}

	// Now encode it back
	var attrs goipp.Attributes
	codec.encode(out, &attrs)

	// End compare encoded attributes
	if test.skipEncode {
		return
	}

	// Note, as decoding/encoding doesn't preserve
	// original order of attributes, we need to
	// sort them before comparison
	attrs2 := make(goipp.Attributes, len(test.attrs))
	copy(attrs2, test.attrs)

	sort.Slice(attrs, func(i, j int) bool {
		return attrs[i].Name < attrs[j].Name
	})

	sort.Slice(attrs2, func(i, j int) bool {
		return attrs2[i].Name < attrs2[j].Name
	})

	if !attrs.Equal(attrs2) {
		t.Errorf("in test %q:", test.name)
		t.Errorf("encode: input/output mismatch")

		var msg goipp.Message
		buf := &bytes.Buffer{}

		msg = goipp.Message{Printer: attrs2}
		msg.Print(buf, true)
		t.Errorf("expected:\n%s", buf)

		buf.Reset()
		msg = goipp.Message{Printer: attrs}
		msg.Print(buf, true)
		t.Errorf("present:\n%s", buf)
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

var (
	testTime1 = time.Date(1970, time.January, 9, 23, 0, 0, 0, time.UTC)
	testTime2 = time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	testTime3 = time.Date(2019, time.April, 12, 15, 0, 0, 0, time.UTC)
	testTime4 = time.Date(2025, time.May, 17, 45, 0, 0, 0, time.UTC)
)

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
