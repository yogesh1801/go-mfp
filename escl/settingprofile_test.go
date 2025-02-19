// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// SettingProfile tests

package escl

import (
	"errors"
	"reflect"
	"testing"

	"github.com/alexpevzner/mfp/xmldoc"
)

// TestSettingProfile tests [SettingProfile] conversion to and from the XML
func TestSettingProfile(t *testing.T) {
	type testData struct {
		prof SettingProfile
		xml  xmldoc.Element
	}

	tests := []testData{
		{
			// Full data test
			prof: SettingProfile{
				ColorModes: []ColorMode{
					Grayscale8, RGB24,
				},
				DocumentFormats: []string{
					"image/jpeg", "application/pdf",
				},
				DocumentFormatsExt: []string{
					"image/jpeg", "application/pdf",
				},
				SupportedResolutions: testSupportedResolutions,
				ColorSpaces: []ColorSpace{
					SRGB,
				},
				CcdChannels: []CcdChannel{
					Red, Green, Blue,
				},
				BinaryRenderings: []BinaryRendering{
					Halftone, Threshold,
				},
			},
			xml: xmldoc.WithChildren(NsScan+":SettingProfile",
				xmldoc.WithChildren(NsScan+":ColorModes",
					xmldoc.WithText(NsScan+":ColorMode",
						NsScan+":Grayscale8"),
					xmldoc.WithText(NsScan+":ColorMode",
						NsScan+":RGB24"),
				),
				xmldoc.WithChildren(NsScan+":DocumentFormats",
					xmldoc.WithText(NsPWG+":DocumentFormat",
						"image/jpeg"),
					xmldoc.WithText(NsScan+":DocumentFormatExt",
						"image/jpeg"),
					xmldoc.WithText(NsPWG+":DocumentFormat",
						"application/pdf"),
					xmldoc.WithText(NsScan+":DocumentFormatExt",
						"application/pdf"),
				),
				testSupportedResolutions.toXML(NsScan+":SupportedResolutions"),
				xmldoc.WithChildren(NsScan+":ColorSpaces",
					xmldoc.WithText(NsScan+":ColorSpace",
						NsScan+":sRGB"),
				),
				xmldoc.WithChildren(NsScan+":CcdChannels",
					xmldoc.WithText(NsScan+":CcdChannel",
						NsScan+":Red"),
					xmldoc.WithText(NsScan+":CcdChannel",
						NsScan+":Green"),
					xmldoc.WithText(NsScan+":CcdChannel",
						NsScan+":Blue"),
				),
				xmldoc.WithChildren(NsScan+":BinaryRenderings",
					xmldoc.WithText(NsScan+":BinaryRendering",
						NsScan+":Halftone"),
					xmldoc.WithText(NsScan+":BinaryRendering",
						NsScan+":Threshold"),
				),
			),
		},

		// Missed optional elements
		{
			prof: SettingProfile{
				SupportedResolutions: testSupportedResolutions,
			},
			xml: xmldoc.WithChildren(NsScan+":SettingProfile",
				testSupportedResolutions.toXML(NsScan+":SupportedResolutions"),
			),
		},

		// Difference between DocumentFormats and DocumentFormatsExt
		{
			prof: SettingProfile{
				DocumentFormats: []string{
					"image/jpeg",
				},
				DocumentFormatsExt: []string{
					"application/pdf",
				},
				SupportedResolutions: testSupportedResolutions,
			},
			xml: xmldoc.WithChildren(NsScan+":SettingProfile",
				xmldoc.WithChildren(NsScan+":DocumentFormats",
					xmldoc.WithText(NsPWG+":DocumentFormat",
						"image/jpeg"),
					xmldoc.WithText(NsScan+":DocumentFormatExt",
						"application/pdf"),
				),
				testSupportedResolutions.toXML(NsScan+":SupportedResolutions"),
			),
		},
	}

	for _, test := range tests {
		xml := test.prof.toXML(NsScan + ":SettingProfile")
		if !xml.Similar(test.xml) {
			t.Errorf("encode mismatch:\n"+
				"expected: %s\n"+
				"present:  %s\n",
				test.xml.EncodeString(nil),
				xml.EncodeString(nil))
		}

		prof, err := decodeSettingProfile(test.xml)
		if err != nil {
			t.Errorf("decode error:\n"+
				"input: %s\n"+
				"error:  %s\n",
				test.xml.EncodeString(nil), err)
			continue
		}

		if !reflect.DeepEqual(prof, test.prof) {
			t.Errorf("decode mismatch:\n"+
				"expected: %#v\n"+
				"present:  %#v\n",
				test.prof, prof)
		}
	}
}

// TestSettingProfileDecodeErrors tests [SettingProfile] XML decode
// errors handling
func TestSettingProfileDecodeErrors(t *testing.T) {
	type testData struct {
		xml xmldoc.Element // Input XML
		err string         // Expected error
	}

	res := SupportedResolutions{
		DiscreteResolutions: []DiscreteResolution{
			{100, 100},
			{200, 200},
		},
	}

	tests := []testData{
		// Missed required element
		{
			xml: xmldoc.WithChildren(NsScan + ":SettingProfile"),
			err: `/scan:SettingProfile/scan:SupportedResolutions: missed`,
		},

		// Error in SupportedResolutions
		{
			xml: xmldoc.WithChildren(NsScan+":SettingProfile",
				xmldoc.WithChildren(NsScan+":SupportedResolutions"),
			),
			err: `/scan:SettingProfile/scan:SupportedResolutions: missed scan:DiscreteResolutions and scan:ResolutionRange`,
		},

		// Error in ColorMode
		{
			xml: xmldoc.WithChildren(NsScan+":SettingProfile",
				xmldoc.WithChildren(NsScan+":ColorModes",
					xmldoc.WithText(NsScan+":ColorMode",
						"Unknown"),
				),
				res.toXML(NsScan+":SupportedResolutions"),
			),
			err: `/scan:SettingProfile/scan:ColorModes/scan:ColorMode: invalid ColorMode: "Unknown"`,
		},

		// Error in ColorSpaces
		{
			xml: xmldoc.WithChildren(NsScan+":SettingProfile",
				xmldoc.WithChildren(NsScan+":ColorSpaces",
					xmldoc.WithText(NsScan+":ColorSpace",
						"Unknown"),
				),
				res.toXML(NsScan+":SupportedResolutions"),
			),
			err: `/scan:SettingProfile/scan:ColorSpaces/scan:ColorSpace: invalid ColorSpace: "Unknown"`,
		},

		// Error in CcdChannels
		{
			xml: xmldoc.WithChildren(NsScan+":SettingProfile",
				xmldoc.WithChildren(NsScan+":CcdChannels",
					xmldoc.WithText(NsScan+":CcdChannel",
						"Unknown"),
				),
				res.toXML(NsScan+":SupportedResolutions"),
			),
			err: `/scan:SettingProfile/scan:CcdChannels/scan:CcdChannel: invalid CcdChannel: "Unknown"`,
		},

		// Error in BinaryRenderings
		{
			xml: xmldoc.WithChildren(NsScan+":SettingProfile",
				xmldoc.WithChildren(NsScan+":BinaryRenderings",
					xmldoc.WithText(NsScan+":BinaryRendering",
						"Unknown"),
				),
				res.toXML(NsScan+":SupportedResolutions"),
			),
			err: `/scan:SettingProfile/scan:BinaryRenderings/scan:BinaryRendering: invalid BinaryRendering: "Unknown"`,
		},
	}

	for _, test := range tests {
		_, err := decodeSettingProfile(test.xml)
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
