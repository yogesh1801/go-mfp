// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Platen capabilities tests.

package escl

import (
	"errors"
	"reflect"
	"testing"

	"github.com/alexpevzner/mfp/optional"
	"github.com/alexpevzner/mfp/xmldoc"
)

// TestPlaten tests [Platen] to/from the XML conversion
func TestPlaten(t *testing.T) {
	type testData struct {
		platen Platen
		xml    xmldoc.Element
	}

	tests := []testData{
		{
			platen: Platen{nil},
			xml:    xmldoc.WithChildren(NsScan + ":Platen"),
		},

		{
			platen: Platen{
				PlatenInputCaps: optional.New(testInputSourceCaps),
			},
			xml: xmldoc.WithChildren(
				NsScan+":Platen",
				testInputSourceCaps.toXML(NsScan+":PlatenInputCaps"),
			),
		},
	}

	for _, test := range tests {
		xml := test.platen.toXML(NsScan + ":Platen")
		if !reflect.DeepEqual(xml, test.xml) {
			t.Errorf("encode mismatch:\n"+
				"expected: %s\n"+
				"present:  %s\n",
				test.xml.EncodeString(nil),
				xml.EncodeString(nil))
		}

		platen, err := decodePlaten(test.xml)
		if err != nil {
			t.Errorf("decode error:\n"+
				"input: %s\n"+
				"error:  %s\n",
				test.xml.EncodeString(nil), err)
			continue
		}

		if !reflect.DeepEqual(platen, test.platen) {
			t.Errorf("decode mismatch:\n"+
				"expected: %#v\n"+
				"present:  %#v\n",
				test.platen, platen)
		}
	}
}

// TestPlatenDecodeErrors tests [Platen] XML decode
// errors handling
func TestPlatenDecodeErrors(t *testing.T) {
	type testData struct {
		xml xmldoc.Element
		err string
	}

	tests := []testData{
		{
			xml: xmldoc.WithChildren(
				NsScan+":Platen",
				xmldoc.WithChildren(NsScan+":PlatenInputCaps"),
			),
			err: `/scan:Platen/scan:PlatenInputCaps/scan:MinWidth: missed`,
		},
	}

	for _, test := range tests {
		_, err := decodePlaten(test.xml)
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
