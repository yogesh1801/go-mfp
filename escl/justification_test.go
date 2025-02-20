// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// ADF image justification tests

package escl

import (
	"errors"
	"reflect"
	"testing"

	"github.com/alexpevzner/mfp/xmldoc"
)

// TestJustification tests [Justification] conversion to and from the XML
func TestJustification(t *testing.T) {
	type testData struct {
		jst Justification
		xml xmldoc.Element
	}

	tests := []testData{
		{
			jst: Justification{Left, Top},
			xml: xmldoc.WithChildren(
				NsScan+":Justification",
				Left.toXML(NsScan+":XImagePosition"),
				Top.toXML(NsScan+":YImagePosition"),
			),
		},
	}

	for _, test := range tests {
		xml := test.jst.toXML(NsScan + ":Justification")
		if !xml.Similar(test.xml) {
			t.Errorf("encode mismatch:\n"+
				"expected: %s\n"+
				"present:  %s\n",
				test.xml.EncodeString(nil),
				xml.EncodeString(nil))
		}

		jst, err := decodeJustification(test.xml)
		if err != nil {
			t.Errorf("decode error:\n"+
				"input: %s\n"+
				"error:  %s\n",
				test.xml.EncodeString(nil), err)
			continue
		}

		if !reflect.DeepEqual(jst, test.jst) {
			t.Errorf("decode mismatch:\n"+
				"expected: %#v\n"+
				"present:  %#v\n",
				test.jst, jst)
		}
	}
}

// TestJustificationDecodeErrors tests [Justification] XML decode
// errors handling
func TestJustificationDecodeErrors(t *testing.T) {
	type testData struct {
		xml xmldoc.Element
		err string
	}

	tests := []testData{
		// Test for missed elements
		{
			xml: xmldoc.WithChildren(
				NsScan + ":Justification",
			),
			err: `/scan:Justification/scan:XImagePosition: missed`,
		},

		{
			xml: xmldoc.WithChildren(
				NsScan+":Justification",
				Left.toXML(NsScan+":XImagePosition"),
			),
			err: `/scan:Justification/scan:YImagePosition: missed`,
		},

		// Test for invalid elements
		{
			xml: xmldoc.WithChildren(
				NsScan+":Justification",
				xmldoc.WithText(NsScan+":XImagePosition", "bad"),
				Top.toXML(NsScan+":YImagePosition"),
			),
			err: `/scan:Justification/scan:XImagePosition: invalid ImagePosition: "bad"`,
		},

		{
			xml: xmldoc.WithChildren(
				NsScan+":Justification",
				Left.toXML(NsScan+":XImagePosition"),
				xmldoc.WithText(NsScan+":YImagePosition", "bad"),
			),
			err: `/scan:Justification/scan:YImagePosition: invalid ImagePosition: "bad"`,
		},
	}

	for _, test := range tests {
		_, err := decodeJustification(test.xml)
		if err == nil {
			err = errors.New("")
		}

		if err.Error() != test.err {
			t.Errorf("error mismatch:\n"+
				"input:    %s\n"+
				"expected: %q\n"+
				"present:  %q\n",
				test.xml.EncodeString(nil), test.err, err)
		}
	}
}
