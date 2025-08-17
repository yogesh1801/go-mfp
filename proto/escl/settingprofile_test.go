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

	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// testSettingProfile contains example of the initialized
// SettingProfile structure
var testSettingProfile = SettingProfile{
	ColorModes: []ColorMode{
		Grayscale8, RGB24,
	},
	ContentTypes: []ContentType{
		ContentTypePhoto, ContentTypeText, ContentTypeTextAndPhoto,
	},
	DocumentFormats: []string{
		"image/jpeg", "application/pdf",
	},
	DocumentFormatsExt: []string{
		"image/jpeg", "application/pdf",
	},
	SupportedResolutions: []SupportedResolutions{
		testSupportedResolutions,
	},
	ColorSpaces: []ColorSpace{
		SRGB,
	},
	CCDChannels: []CCDChannel{
		Red, Green, Blue,
	},
	BinaryRenderings: []BinaryRendering{
		Halftone, Threshold,
	},
}

// TestSettingProfile tests [SettingProfile] conversion to and from the XML
func TestSettingProfile(t *testing.T) {
	type testData struct {
		prof SettingProfile
		xml  xmldoc.Element
	}

	tests := []testData{
		{
			// Full data test
			prof: testSettingProfile,
			xml: xmldoc.WithChildren(NsScan+":SettingProfile",
				xmldoc.WithChildren(NsScan+":ColorModes",
					Grayscale8.toXML(NsScan+":ColorMode"),
					RGB24.toXML(NsScan+":ColorMode"),
				),
				xmldoc.WithChildren(NsScan+":ContentTypes",
					ContentTypePhoto.toXML(NsPWG+":ContentType"),
					ContentTypeText.toXML(NsPWG+":ContentType"),
					ContentTypeTextAndPhoto.toXML(NsPWG+":ContentType"),
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
						"sRGB"),
				),
				xmldoc.WithChildren(NsScan+":CcdChannels",
					Red.toXML(NsScan+":CcdChannel"),
					Green.toXML(NsScan+":CcdChannel"),
					Blue.toXML(NsScan+":CcdChannel"),
				),
				xmldoc.WithChildren(NsScan+":BinaryRenderings",
					Halftone.toXML(NsScan+":BinaryRendering"),
					Threshold.toXML(NsScan+":BinaryRendering"),
				),
			),
		},

		// Missed optional elements
		{
			prof: SettingProfile{
				SupportedResolutions: []SupportedResolutions{
					testSupportedResolutions,
				},
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
				SupportedResolutions: []SupportedResolutions{
					testSupportedResolutions,
				},
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

		// Error in ContentType
		{
			xml: xmldoc.WithChildren(NsScan+":SettingProfile",
				xmldoc.WithChildren(NsScan+":ContentTypes",
					xmldoc.WithText(NsPWG+":ContentType",
						"Unknown"),
				),
				res.toXML(NsScan+":SupportedResolutions"),
			),
			err: `/scan:SettingProfile/scan:ContentTypes/pwg:ContentType: invalid ContentType: "Unknown"`,
		},

		// Error in ColorSpaces
		{
			xml: xmldoc.WithChildren(NsScan+":SettingProfile",
				xmldoc.WithChildren(NsScan+":ColorSpaces",
					xmldoc.WithText(NsScan+":ColorSpace",
						"bad val"),
				),
				res.toXML(NsScan+":SupportedResolutions"),
			),
			err: `/scan:SettingProfile/scan:ColorSpaces/scan:ColorSpace: invalid ColorSpace: "bad val"`,
		},

		// Error in CCDChannels
		{
			xml: xmldoc.WithChildren(NsScan+":SettingProfile",
				xmldoc.WithChildren(NsScan+":CcdChannels",
					xmldoc.WithText(NsScan+":CcdChannel",
						"Unknown"),
				),
				res.toXML(NsScan+":SupportedResolutions"),
			),
			err: `/scan:SettingProfile/scan:CcdChannels/scan:CcdChannel: invalid CCDChannel: "Unknown"`,
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
