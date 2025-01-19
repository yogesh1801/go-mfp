// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Common type for Range of some value test.

package escl

import (
	"reflect"
	"testing"

	"github.com/alexpevzner/mfp/optional"
	"github.com/alexpevzner/mfp/xmldoc"
)

// TestRange tests Range
func TestRange(t *testing.T) {
	type testData struct {
		rng Range
		xml xmldoc.Element
	}

	tests := []testData{
		{
			rng: Range{
				Min:    0,
				Max:    100,
				Step:   optional.New(10),
				Normal: 50,
			},
			xml: xmldoc.Element{
				Name: NsScan + ":Range",
				Children: []xmldoc.Element{
					{Name: NsScan + ":Min", Text: "0"},
					{Name: NsScan + ":Max", Text: "100"},
					{Name: NsScan + ":Normal", Text: "50"},
					{Name: NsScan + ":Step", Text: "10"},
				},
			},
		},

		{
			rng: Range{
				Min:    100,
				Max:    200,
				Normal: 150,
			},
			xml: xmldoc.Element{
				Name: NsScan + ":Range",
				Children: []xmldoc.Element{
					{Name: NsScan + ":Min", Text: "100"},
					{Name: NsScan + ":Max", Text: "200"},
					{Name: NsScan + ":Normal", Text: "150"},
				},
			},
		},
	}

	for _, test := range tests {
		xml := test.rng.ToXML(NsScan + ":Range")
		if !reflect.DeepEqual(xml, test.xml) {
			t.Errorf("ToXML:\nexpected: %s\npresent: %s\n",
				test.xml.EncodeString(NsMap),
				xml.EncodeString(NsMap))
		}

		rng, err := decodeRange(xml)
		if err != nil {
			t.Errorf("decodeRange: %s", err)
			continue
		}

		if !reflect.DeepEqual(rng, test.rng) {
			t.Errorf("DecodeAppSequence:\n"+
				"expected: %#v\npresent:  %#v\n",
				test.rng, rng)
		}
	}
}

// TestRangeErrors tests Range decode errors
func TestRangeErrors(t *testing.T) {
	type testData struct {
		xml  xmldoc.Element
		estr string
	}

	tests := []testData{
		{
			xml: xmldoc.Element{
				Name: NsScan + ":Range",
				Children: []xmldoc.Element{
					{Name: NsScan + ":Min", Text: "0"},
					{Name: NsScan + ":Normal", Text: "50"},
					{Name: NsScan + ":Step", Text: "10"},
				},
			},
			estr: `/scan:Range/scan:Max: missed`,
		},
		{
			xml: xmldoc.Element{
				Name: NsScan + ":Range",
				Children: []xmldoc.Element{
					{Name: NsScan + ":Min", Text: "0"},
					{Name: NsScan + ":Max", Text: "hello"},
					{Name: NsScan + ":Normal", Text: "50"},
					{Name: NsScan + ":Step", Text: "10"},
				},
			},
			estr: `/scan:Range/scan:Max: invalid int: "hello"`,
		},
	}

	for _, test := range tests {
		_, err := decodeRange(test.xml)
		estr := ""
		if err != nil {
			estr = err.Error()
		}

		if estr != test.estr {
			t.Errorf("%s\nexpected: %q\npresent:  %q",
				test.xml.EncodeString(NsMap),
				test.estr, estr)
		}
	}
}
