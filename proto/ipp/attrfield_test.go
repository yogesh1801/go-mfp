// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Handling of IPP structure fields

package ipp

import (
	"reflect"
	"testing"

	"github.com/OpenPrinting/go-mfp/proto/ipp/iana"
	"github.com/OpenPrinting/goipp"
)

// TestAttrSyntaxTokenize tests the attrSyntaxTokenize function
func TestAttrSyntaxTokenize(t *testing.T) {
	type testData struct {
		in       string   // Input string
		expected []string // Output tokens
	}

	tests := []testData{
		{
			in:       "",
			expected: []string{},
		},
		{
			in:       "integer",
			expected: []string{"integer"},
		},

		{
			in:       "1setof integer",
			expected: []string{"1setof", "integer"},
		},

		{
			in:       "1setof integer(MIN)",
			expected: []string{"1setof", "integer", "(", "min", ")"},
		},

		{
			in:       "1setof integer(MIN:MAX)",
			expected: []string{"1setof", "integer", "(", "min", ":", "max", ")"},
		},

		{
			in:       "keyword | no-value",
			expected: []string{"keyword", "|", "no-value"},
		},

		{
			in:       "привет",
			expected: []string{"п", "р", "и", "в", "е", "т"},
		},
	}

	for _, test := range tests {
		tokens := attrSyntaxTokenize(test.in)
		if !reflect.DeepEqual(tokens, test.expected) {
			t.Errorf("attrSyntaxTokenize:\n"+
				"input:    %q\n"+
				"expected: %#v\n"+
				"present:  %#v\n",
				test.in, tokens, test.expected)
		}
	}
}

// TestAttrSyntaxParse tests the attrSyntaxParse function
func TestAttrSyntaxParse(t *testing.T) {
	type testData struct {
		in  string        // Input string
		out *iana.DefAttr // Expected output
		err string        // Expected error, "" if none
	}

	tests := []testData{
		{
			in: "integer",
			out: &iana.DefAttr{
				Min:  iana.MIN,
				Max:  iana.MAX,
				Tags: []goipp.Tag{goipp.TagInteger},
			},
		},

		{
			in: "integer(0:MAX)",
			out: &iana.DefAttr{
				Min:  0,
				Max:  iana.MAX,
				Tags: []goipp.Tag{goipp.TagInteger},
			},
		},

		{
			in: "name",
			out: &iana.DefAttr{
				Min:  0,
				Max:  255,
				Tags: []goipp.Tag{goipp.TagName},
			},
		},

		{
			in: "name(min:max)",
			out: &iana.DefAttr{
				Min:  0,
				Max:  255,
				Tags: []goipp.Tag{goipp.TagName},
			},
		},

		{
			in: "name(5:max)",
			out: &iana.DefAttr{
				Min:  5,
				Max:  255,
				Tags: []goipp.Tag{goipp.TagName},
			},
		},

		{
			in: "name(max)",
			out: &iana.DefAttr{
				Min:  0,
				Max:  255,
				Tags: []goipp.Tag{goipp.TagName},
			},
		},

		{
			in: "name(63)",
			out: &iana.DefAttr{
				Min:  0,
				Max:  63,
				Tags: []goipp.Tag{goipp.TagName},
			},
		},

		{
			in: "name(1:63)",
			out: &iana.DefAttr{
				Min:  1,
				Max:  63,
				Tags: []goipp.Tag{goipp.TagName},
			},
		},

		{
			in: "1setof name(1:63)",
			out: &iana.DefAttr{
				SetOf: true,
				Min:   1,
				Max:   63,
				Tags:  []goipp.Tag{goipp.TagName},
			},
		},

		{
			in: "name|unknown",
			out: &iana.DefAttr{
				Min:  0,
				Max:  255,
				Tags: []goipp.Tag{goipp.TagName, goipp.TagUnknown},
			},
		},

		{
			in: "name|keyword|unknown",
			out: &iana.DefAttr{
				Min:  1,
				Max:  255,
				Tags: []goipp.Tag{goipp.TagKeyword, goipp.TagName, goipp.TagUnknown},
			},
		},

		{
			in:  "invalid",
			err: `ipp:"invalid": unexpected token`,
		},

		{
			in:  "",
			err: `ipp:"": no tags defined`,
		},

		{
			in:  "1setOf",
			err: `ipp:"1setOf": no tags defined`,
		},

		{
			in:  "name(foo:max)",
			err: `ipp:"foo": invalid limit`,
		},

		{
			in:  "name(min:bar)",
			err: `ipp:"bar": invalid limit`,
		},

		{
			in:  "name(foobar)",
			err: `ipp:"foobar": invalid limit`,
		},
	}

	for _, test := range tests {
		def, err := attrSyntaxParse(test.in)
		errstr := ""
		if err != nil {
			errstr = err.Error()
		}

		if errstr != test.err {
			t.Errorf("attrSyntaxParse: error mismatch:\n"+
				"input:    %q\n"+
				"expected: %q\n"+
				"present:  %q",
				test.in, test.err, err)
			continue
		}

		if !reflect.DeepEqual(test.out, def) {
			t.Errorf("attrSyntaxParse: output mismatch:\n"+
				"input:    %q\n"+
				"expected: %#v\n"+
				"present:  %#v",
				test.in, test.out, def)
		}
	}
}

// TestAttrFieldAnalyze tests the attrFieldAnalyze function
func TestAttrFieldAnalyze(t *testing.T) {
	type testData struct {
		fld  reflect.StructField // Input structure field
		name string              // Expected name
		def  *iana.DefAttr       // Expected attribute definition
		err  string              // Expected error, "" if none
	}

	tests := []testData{
		{
			fld: reflect.StructField{
				Name: "CharsetSupported",
				Tag:  `ipp:"charset-configured,charset"`,
			},
			name: "charset-configured",
			def: &iana.DefAttr{
				Min:  0,
				Max:  63,
				Tags: []goipp.Tag{goipp.TagCharset},
			},
		},

		{
			fld: reflect.StructField{
				Name: "CharsetSupported",
				Tag:  `ipp:"charset-configured"`,
			},
			name: "charset-configured",
			def:  nil,
		},

		{
			fld: reflect.StructField{
				Name: "CharsetSupported",
				Tag:  `ipp:"charset-configured,"`,
			},
			name: "",
			def:  nil,
			err:  `ipp:"": no tags defined`,
		},

		{
			fld: reflect.StructField{
				Name: "CharsetSupported",
			},
			name: "",
			def:  nil,
		},

		{
			fld: reflect.StructField{
				Name:    "unexported",
				Tag:     `ipp:"charset-configured,charset"`,
				PkgPath: "foo", // marks field as unexported
			},
			name: "",
			def:  nil,
			err:  `ipp:tag used with unexported field`,
		},

		{
			fld: reflect.StructField{
				Name:      "CharsetSupported",
				Tag:       `ipp:"charset-configured,charset"`,
				Anonymous: true,
			},
			name: "",
			def:  nil,
			err:  `ipp:tag used with anonymous field`,
		},

		{
			fld: reflect.StructField{
				Name: "CharsetSupported",
				Tag:  `ipp:""`,
			},
			name: "",
			def:  nil,
			err:  `ipp:missed attribute name`,
		},
	}

	for _, test := range tests {
		name, def, err := attrFieldAnalyze(test.fld)
		errstr := ""
		if err != nil {
			errstr = err.Error()
		}

		if errstr != test.err {
			t.Errorf("attrSyntaxParse: error mismatch:\n"+
				"input:    %#v\n"+
				"expected: %q\n"+
				"present:  %q",
				test.fld, test.err, err)
			continue
		}

		if name != test.name {
			t.Errorf("attrSyntaxParse: output name mismatch:\n"+
				"input:    %#v\n"+
				"expected: %s\n"+
				"present:  %s",
				test.fld, test.name, name)
		}

		if !reflect.DeepEqual(test.def, def) {
			t.Errorf("attrSyntaxParse: output mismatch:\n"+
				"input:    %#v\n"+
				"expected: %#v\n"+
				"present:  %#v",
				test.fld, test.def, def)
		}
	}
}
