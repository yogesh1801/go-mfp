// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// ADF capabilities tests.

package escl

import (
	"errors"
	"reflect"
	"testing"

	"github.com/alexpevzner/mfp/optional"
	"github.com/alexpevzner/mfp/xmldoc"
)

// testADF contains example of the initialized
// ADF structure
var testADF = ADF{
	ADFSimplexInputCaps: optional.New(testInputSourceCaps),
	ADFDuplexInputCaps:  optional.New(testInputSourceCaps),
	FeederCapacity:      optional.New(50),
	ADFOptions: []ADFOption{
		DetectPaperLoaded,
		Duplex,
	},
	Justification: optional.New(Justification{
		XImagePosition: Center,
		YImagePosition: Top,
	}),
}

// TestADF tests [ADF] to/from the XML conversion
func TestADF(t *testing.T) {
	type testData struct {
		adf ADF
		xml xmldoc.Element
	}

	tests := []testData{
		{
			adf: ADF{},
			xml: xmldoc.WithChildren(NsScan + ":Adf"),
		},

		{
			adf: ADF{
				ADFSimplexInputCaps: optional.New(testInputSourceCaps),
			},
			xml: xmldoc.WithChildren(
				NsScan+":Adf",
				testInputSourceCaps.toXML(NsScan+":AdfSimplexInputCaps"),
			),
		},

		{
			adf: testADF,
			xml: xmldoc.WithChildren(
				NsScan+":Adf",
				testInputSourceCaps.toXML(NsScan+":AdfSimplexInputCaps"),
				testInputSourceCaps.toXML(NsScan+":AdfDuplexInputCaps"),
				xmldoc.WithText(NsScan+":FeederCapacity", "50"),
				xmldoc.WithChildren(
					NsScan+":AdfOptions",
					DetectPaperLoaded.toXML(NsScan+":AdfOption"),
					Duplex.toXML(NsScan+":AdfOption"),
				),
				xmldoc.WithChildren(
					NsScan+":Justification",
					Center.toXML(NsScan+":XImagePosition"),
					Top.toXML(NsScan+":YImagePosition"),
				),
			),
		},
	}

	for _, test := range tests {
		xml := test.adf.toXML(NsScan + ":Adf")
		if !xml.Similar(test.xml) {
			t.Errorf("encode mismatch:\n"+
				"expected: %s\n"+
				"present:  %s\n",
				test.xml.EncodeString(nil),
				xml.EncodeString(nil))
		}

		adf, err := decodeADF(test.xml)
		if err != nil {
			t.Errorf("decode error:\n"+
				"input: %s\n"+
				"error:  %s\n",
				test.xml.EncodeString(nil), err)
			continue
		}

		if !reflect.DeepEqual(adf, test.adf) {
			t.Errorf("decode mismatch:\n"+
				"expected: %#v\n"+
				"present:  %#v\n",
				test.adf, adf)
		}
	}
}

// TestADFDecodeErrors tests [ADF] XML decode
// errors handling
func TestADFDecodeErrors(t *testing.T) {
	type testData struct {
		xml xmldoc.Element
		err string
	}

	tests := []testData{
		{
			xml: xmldoc.WithChildren(
				NsScan+":Adf",
				xmldoc.WithText(NsScan+":AdfSimplexInputCaps", "bad"),
			),
			err: `/scan:Adf/scan:AdfSimplexInputCaps/scan:MinWidth: missed`,
		},

		{
			xml: xmldoc.WithChildren(
				NsScan+":Adf",
				xmldoc.WithText(NsScan+":AdfDuplexInputCaps", "bad"),
			),
			err: `/scan:Adf/scan:AdfDuplexInputCaps/scan:MinWidth: missed`,
		},

		{
			xml: xmldoc.WithChildren(
				NsScan+":Adf",
				xmldoc.WithText(NsScan+":FeederCapacity", "bad"),
			),
			err: `/scan:Adf/scan:FeederCapacity: invalid int: "bad"`,
		},

		{
			xml: xmldoc.WithChildren(
				NsScan+":Adf",
				xmldoc.WithChildren(
					NsScan+":AdfOptions",
					xmldoc.WithText(NsScan+":AdfOption", "bad"),
				),
			),
			err: `/scan:Adf/scan:AdfOptions/scan:AdfOption: invalid ADFOption: "bad"`,
		},

		{
			xml: xmldoc.WithChildren(
				NsScan+":Adf",
				xmldoc.WithChildren(NsScan+":Justification"),
			),
			err: `/scan:Adf/scan:Justification/scan:XImagePosition: missed`,
		},
	}

	for _, test := range tests {
		_, err := decodeADF(test.xml)
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
