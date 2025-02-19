// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scanner resolutions test

package escl

import (
	"errors"
	"reflect"
	"testing"

	"github.com/alexpevzner/mfp/optional"
	"github.com/alexpevzner/mfp/xmldoc"
)

// testSupportedResolutions contains example of the initialized
// SupportedResolutions structure
var testSupportedResolutions = SupportedResolutions{
	DiscreteResolutions: []DiscreteResolution{
		{100, 100},
		{200, 200},
	},
}

// TestResolutionRange tests [ResolutionRange] to/from the XML conversion
func TestResolutionRange(t *testing.T) {
	type testData struct {
		rng ResolutionRange
		xml xmldoc.Element
	}

	tests := []testData{
		{
			rng: ResolutionRange{
				Range{300, 1200, 450, optional.New(150)},
				Range{100, 600, 300, optional.New(100)},
			},
			xml: xmldoc.WithChildren(
				NsScan+":ResolutionRange",
				xmldoc.WithChildren(
					NsScan+":XResolutionRange",
					xmldoc.WithText(NsScan+":Min", "300"),
					xmldoc.WithText(NsScan+":Max", "1200"),
					xmldoc.WithText(NsScan+":Normal", "450"),
					xmldoc.WithText(NsScan+":Step", "150"),
				),
				xmldoc.WithChildren(
					NsScan+":YResolutionRange",
					xmldoc.WithText(NsScan+":Min", "100"),
					xmldoc.WithText(NsScan+":Max", "600"),
					xmldoc.WithText(NsScan+":Normal", "300"),
					xmldoc.WithText(NsScan+":Step", "100"),
				),
			),
		},
	}

	for _, test := range tests {
		xml := test.rng.toXML(NsScan + ":ResolutionRange")
		if !reflect.DeepEqual(xml, test.xml) {
			t.Errorf("encode mismatch:\n"+
				"expected: %s\n"+
				"present:  %s\n",
				test.xml.EncodeString(nil),
				xml.EncodeString(nil))
		}

		rng, err := decodeResolutionRange(test.xml)
		if err != nil {
			t.Errorf("decode error:\n"+
				"input: %s\n"+
				"error:  %s\n",
				test.xml.EncodeString(nil), err)
			continue
		}

		if !reflect.DeepEqual(rng, test.rng) {
			t.Errorf("decode mismatch:\n"+
				"expected: %#v\n"+
				"present:  %#v\n",
				test.rng, rng)
		}
	}
}

// TestResolutionRangeDecodeErrors tests [ResolutionRange] XML decode
// errors handling
func TestResolutionRangeDecodeErrors(t *testing.T) {
	type testData struct {
		xml xmldoc.Element
		err string
	}

	tests := []testData{
		{
			xml: xmldoc.WithChildren(
				NsScan+":ResolutionRange",
				xmldoc.WithChildren(
					NsScan+":XResolutionRange",
					xmldoc.WithText(NsScan+":Min", "300"),
					xmldoc.WithText(NsScan+":Max", "1200"),
					xmldoc.WithText(NsScan+":Normal", "450"),
					xmldoc.WithText(NsScan+":Step", "150"),
				),
			),
			err: `/scan:ResolutionRange/scan:YResolutionRange: missed`,
		},
		{
			xml: xmldoc.WithChildren(
				NsScan+":ResolutionRange",
				xmldoc.WithChildren(
					NsScan+":XResolutionRange",
					xmldoc.WithText(NsScan+":Min", "300"),
					xmldoc.WithText(NsScan+":Max", "1200"),
					xmldoc.WithText(NsScan+":Normal", "450"),
					xmldoc.WithText(NsScan+":Step", "150"),
				),
				xmldoc.WithChildren(
					NsScan+":YResolutionRange",
					xmldoc.WithText(NsScan+":Max", "600"),
					xmldoc.WithText(NsScan+":Normal", "300"),
					xmldoc.WithText(NsScan+":Step", "100"),
				),
			),
			err: `/scan:ResolutionRange/scan:YResolutionRange/scan:Min: missed`,
		},
		{
			xml: xmldoc.WithChildren(
				NsScan+":ResolutionRange",
				xmldoc.WithChildren(
					NsScan+":XResolutionRange",
					xmldoc.WithText(NsScan+":Min", "300"),
					xmldoc.WithText(NsScan+":Max", "1200"),
					xmldoc.WithText(NsScan+":Normal", "450"),
					xmldoc.WithText(NsScan+":Step", "150"),
				),
				xmldoc.WithChildren(
					NsScan+":YResolutionRange",
					xmldoc.WithText(NsScan+":Min", "100"),
					xmldoc.WithText(NsScan+":Max", "aaa"),
					xmldoc.WithText(NsScan+":Normal", "300"),
					xmldoc.WithText(NsScan+":Step", "100"),
				),
			),
			err: `/scan:ResolutionRange/scan:YResolutionRange/scan:Max: invalid int: "aaa"`,
		},
	}

	for _, test := range tests {
		_, err := decodeResolutionRange(test.xml)
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

// TestDiscreteResolution tests [DiscreteResolution] to/from the XML conversion
func TestDiscreteResolution(t *testing.T) {
	type testData struct {
		res DiscreteResolution
		xml xmldoc.Element
	}

	tests := []testData{
		{
			res: DiscreteResolution{200, 400},
			xml: xmldoc.WithChildren(
				NsScan+":DiscreteResolution",
				xmldoc.WithText(NsScan+":XResolution", "200"),
				xmldoc.WithText(NsScan+":YResolution", "400"),
			),
		},
	}

	for _, test := range tests {
		xml := test.res.toXML(NsScan + ":DiscreteResolution")
		if !reflect.DeepEqual(xml, test.xml) {
			t.Errorf("encode mismatch:\n"+
				"expected: %s\n"+
				"present:  %s\n",
				test.xml.EncodeString(nil),
				xml.EncodeString(nil))
		}

		res, err := decodeDiscreteResolution(test.xml)
		if err != nil {
			t.Errorf("decode error:\n"+
				"input: %s\n"+
				"error:  %s\n",
				test.xml.EncodeString(nil), err)
			continue
		}

		if !reflect.DeepEqual(res, test.res) {
			t.Errorf("decode mismatch:\n"+
				"expected: %#v\n"+
				"present:  %#v\n",
				test.res, res)
		}
	}
}

// TestDiscreteResolutionDecodeErrors tests [DiscreteResolution] XML decode
// errors handling
func TestDiscreteResolutionDecodeErrors(t *testing.T) {
	type testData struct {
		xml xmldoc.Element
		err string
	}

	tests := []testData{
		{
			xml: xmldoc.WithChildren(
				NsScan+":DiscreteResolution",
				xmldoc.WithText(NsScan+":YResolution", "400"),
			),
			err: `/scan:DiscreteResolution/scan:XResolution: missed`,
		},
		{
			xml: xmldoc.WithChildren(
				NsScan+":DiscreteResolution",
				xmldoc.WithText(NsScan+":XResolution", "200"),
			),
			err: `/scan:DiscreteResolution/scan:YResolution: missed`,
		},
		{
			xml: xmldoc.WithChildren(
				NsScan+":DiscreteResolution",
				xmldoc.WithText(NsScan+":XResolution", "200"),
				xmldoc.WithText(NsScan+":YResolution", "aaa"),
			),
			err: `/scan:DiscreteResolution/scan:YResolution: invalid int: "aaa"`,
		},
	}

	for _, test := range tests {
		_, err := decodeDiscreteResolution(test.xml)
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

// TestSupportedResolutions tests [SupportedResolutions] to/from the XML
// conversion
func TestSupportedResolutions(t *testing.T) {
	type testData struct {
		supp SupportedResolutions
		xml  xmldoc.Element
	}

	tests := []testData{
		{
			// testSupportedResolutions
			supp: testSupportedResolutions,
			xml: xmldoc.WithChildren(
				NsScan+":SupportedResolutions",
				xmldoc.WithChildren(
					NsScan+":DiscreteResolutions",
					xmldoc.WithChildren(
						NsScan+":DiscreteResolution",
						xmldoc.WithText(NsScan+":XResolution", "100"),
						xmldoc.WithText(NsScan+":YResolution", "100"),
					),
					xmldoc.WithChildren(
						NsScan+":DiscreteResolution",
						xmldoc.WithText(NsScan+":XResolution", "200"),
						xmldoc.WithText(NsScan+":YResolution", "200"),
					),
				),
			),
		},

		{
			// DiscreteResolutions + ResolutionRange
			supp: SupportedResolutions{
				DiscreteResolutions: []DiscreteResolution{
					{100, 100},
					{200, 200},
				},
				ResolutionRange: optional.New(ResolutionRange{
					XResolutionRange: Range{
						100, 300, 150, nil,
					},
					YResolutionRange: Range{
						100, 300, 150, nil,
					},
				}),
			},
			xml: xmldoc.WithChildren(
				NsScan+":SupportedResolutions",
				xmldoc.WithChildren(
					NsScan+":DiscreteResolutions",
					xmldoc.WithChildren(
						NsScan+":DiscreteResolution",
						xmldoc.WithText(NsScan+":XResolution", "100"),
						xmldoc.WithText(NsScan+":YResolution", "100"),
					),
					xmldoc.WithChildren(
						NsScan+":DiscreteResolution",
						xmldoc.WithText(NsScan+":XResolution", "200"),
						xmldoc.WithText(NsScan+":YResolution", "200"),
					),
				),
				xmldoc.WithChildren(
					NsScan+":ResolutionRange",
					xmldoc.WithChildren(
						NsScan+":XResolutionRange",
						xmldoc.WithText(NsScan+":Min", "100"),
						xmldoc.WithText(NsScan+":Max", "300"),
						xmldoc.WithText(NsScan+":Normal", "150"),
					),
					xmldoc.WithChildren(
						NsScan+":YResolutionRange",
						xmldoc.WithText(NsScan+":Min", "100"),
						xmldoc.WithText(NsScan+":Max", "300"),
						xmldoc.WithText(NsScan+":Normal", "150"),
					),
				),
			),
		},
		{
			// ColorMode + DiscreteResolutions
			supp: SupportedResolutions{
				ColorMode: optional.New(RGB24),
				DiscreteResolutions: []DiscreteResolution{
					{100, 100},
					{200, 200},
				},
			},
			xml: xmldoc.WithChildren(
				NsScan+":SupportedResolutions",
				xmldoc.WithText(NsScan+":ColorMode", NsScan+":RGB24"),
				xmldoc.WithChildren(
					NsScan+":DiscreteResolutions",
					xmldoc.WithChildren(
						NsScan+":DiscreteResolution",
						xmldoc.WithText(NsScan+":XResolution", "100"),
						xmldoc.WithText(NsScan+":YResolution", "100"),
					),
					xmldoc.WithChildren(
						NsScan+":DiscreteResolution",
						xmldoc.WithText(NsScan+":XResolution", "200"),
						xmldoc.WithText(NsScan+":YResolution", "200"),
					),
				),
			),
		},
	}

	for _, test := range tests {
		xml := test.supp.toXML(NsScan + ":SupportedResolutions")
		if !reflect.DeepEqual(xml, test.xml) {
			t.Errorf("encode mismatch:\n"+
				"expected: %s\n"+
				"present:  %s\n",
				test.xml.EncodeString(nil),
				xml.EncodeString(nil))
		}

		supp, err := decodeSupportedResolutions(test.xml)
		if err != nil {
			t.Errorf("decode error:\n"+
				"input: %s\n"+
				"error:  %s\n",
				test.xml.EncodeString(nil), err)
			continue
		}

		if !reflect.DeepEqual(supp, test.supp) {
			t.Errorf("decode mismatch:\n"+
				"expected: %#v\n"+
				"present:  %#v\n",
				test.supp, supp)
		}
	}
}

// TestSupportedResolutionsDecodeErrors tests [SupportedResolutions] XML decode
// errors handling
func TestSupportedResolutionsDecodeErrors(t *testing.T) {
	type testData struct {
		xml xmldoc.Element
		err string
	}

	tests := []testData{
		{
			// Missed DiscreteResolutions and ResolutionRange
			xml: xmldoc.WithChildren(
				NsScan+":SupportedResolutions",
				xmldoc.WithText(NsScan+":ColorMode", NsScan+":RGB24"),
			),
			err: `/scan:SupportedResolutions: missed scan:DiscreteResolutions and scan:ResolutionRange`,
		},
		{
			// Invalid color mode
			xml: xmldoc.WithChildren(
				NsScan+":SupportedResolutions",
				xmldoc.WithText(NsScan+":ColorMode", "AAA"),
				xmldoc.WithChildren(
					NsScan+":DiscreteResolutions",
					xmldoc.WithChildren(
						NsScan+":DiscreteResolution",
						xmldoc.WithText(NsScan+":XResolution", "100"),
						xmldoc.WithText(NsScan+":YResolution", "100"),
					),
				),
			),
			err: `/scan:SupportedResolutions/scan:ColorMode/scan:ColorMode: invalid ColorMode: "AAA"`,
		},
		{
			// Error in DiscreteResolutions
			xml: xmldoc.WithChildren(
				NsScan+":SupportedResolutions",
				xmldoc.WithChildren(
					NsScan+":DiscreteResolutions",
					xmldoc.WithChildren(
						NsScan+":DiscreteResolution",
						xmldoc.WithText(NsScan+":XResolution", "AAA"),
						xmldoc.WithText(NsScan+":YResolution", "100"),
					),
				),
			),
			err: `/scan:SupportedResolutions/scan:DiscreteResolutions/scan:DiscreteResolution/scan:XResolution: invalid int: "AAA"`,
		},
		{
			// Error in ResolutionRange
			xml: xmldoc.WithChildren(
				NsScan+":SupportedResolutions",
				xmldoc.WithChildren(
					NsScan+":ResolutionRange",
					xmldoc.WithChildren(
						NsScan+":XResolutionRange",
						xmldoc.WithText(NsScan+":Min", "100"),
						xmldoc.WithText(NsScan+":Max", "300"),
						xmldoc.WithText(NsScan+":Normal", "150"),
					),
				),
			),
			err: `/scan:SupportedResolutions/scan:ResolutionRange/scan:YResolutionRange: missed`,
		},
	}

	for _, test := range tests {
		_, err := decodeSupportedResolutions(test.xml)
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
