// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// SettingProfile tests

package escl

import (
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

	res := SupportedResolutions{
		DiscreteResolutions: []DiscreteResolution{
			{100, 100},
			{200, 200},
		},
	}

	tests := []testData{
		{
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
				SupportedResolutions: res,
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
				res.toXML(NsScan+":SupportedResolutions"),
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
