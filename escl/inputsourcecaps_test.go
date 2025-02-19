// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Input source capabilities test

package escl

import (
	"reflect"
	"testing"

	"github.com/alexpevzner/mfp/optional"
	"github.com/alexpevzner/mfp/xmldoc"
)

// testInputSourceCaps contains example of the initialized
// InputSourceCaps structure
var testInputSourceCaps = InputSourceCaps{
	MaxWidth:              2551,
	MinWidth:              591,
	MaxHeight:             4205,
	MinHeight:             600,
	MaxXOffset:            optional.New(50),
	MaxYOffset:            optional.New(75),
	MaxOpticalXResolution: optional.New(2400),
	MaxOpticalYResolution: optional.New(1200),
	MaxScanRegions:        optional.New(1),
	RiskyLeftMargins:      optional.New(15),
	RiskyRightMargins:     optional.New(25),
	RiskyTopMargins:       optional.New(20),
	RiskyBottomMargins:    optional.New(35),
	MaxPhysicalWidth:      optional.New(2600),
	MaxPhysicalHeight:     optional.New(600),
	SupportedIntents:      []Intent{Document, TextAndGraphic, Photo},
	EdgeAutoDetection:     []SupportedEdge{TopEdge, LeftEdge},
	SettingProfiles:       []SettingProfile{testSettingProfile},
	FeedDirections:        []FeedDirection{LongEdgeFeed},
}

// TestInputSourceCaps tests [InputSourceCaps] conversion to and from the XML
func TestInputSourceCaps(t *testing.T) {
	type testData struct {
		caps InputSourceCaps
		xml  xmldoc.Element
	}

	tests := []testData{
		{
			caps: testInputSourceCaps,
			xml: xmldoc.WithChildren(
				NsScan+":PlatenInputCaps",
				xmldoc.WithText(NsScan+":MaxWidth", "2551"),
				xmldoc.WithText(NsScan+":MinWidth", "591"),
				xmldoc.WithText(NsScan+":MaxHeight", "4205"),
				xmldoc.WithText(NsScan+":MinHeight", "600"),
				xmldoc.WithText(NsScan+":MaxXOffset", "50"),
				xmldoc.WithText(NsScan+":MaxYOffset", "75"),
				xmldoc.WithText(NsScan+":MaxOpticalXResolution", "2400"),
				xmldoc.WithText(NsScan+":MaxOpticalYResolution", "1200"),
				xmldoc.WithText(NsScan+":MaxScanRegions", "1"),
				xmldoc.WithText(NsScan+":RiskyLeftMargins", "15"),
				xmldoc.WithText(NsScan+":RiskyRightMargins", "25"),
				xmldoc.WithText(NsScan+":RiskyTopMargins", "20"),
				xmldoc.WithText(NsScan+":RiskyBottomMargins", "35"),
				xmldoc.WithText(NsScan+":MaxPhysicalWidth", "2600"),
				xmldoc.WithText(NsScan+":MaxPhysicalHeight", "600"),
				xmldoc.WithChildren(
					NsScan+":SupportedIntents",
					Document.toXML(NsScan+":SupportedIntent"),
					TextAndGraphic.toXML(NsScan+":SupportedIntent"),
					Photo.toXML(NsScan+":SupportedIntent"),
				),
				xmldoc.WithChildren(
					NsScan+":EdgeAutoDetection",
					TopEdge.toXML(NsScan+":SupportedEdge"),
					LeftEdge.toXML(NsScan+":SupportedEdge"),
				),
				xmldoc.WithChildren(
					NsScan+":SettingProfiles",
					testSettingProfile.toXML(NsScan+":SettingProfile"),
				),
				xmldoc.WithChildren(
					NsScan+":FeedDirections",
					LongEdgeFeed.toXML(NsScan+":FeedDirection"),
				),
			),
		},
	}

	for _, test := range tests {
		xml := test.caps.toXML(NsScan + ":PlatenInputCaps")
		if !reflect.DeepEqual(xml, test.xml) {
			t.Errorf("encode mismatch:\n"+
				"expected: %s\n"+
				"present:  %s\n",
				test.xml.EncodeString(nil),
				xml.EncodeString(nil))
		}

		caps, err := decodeInputSourceCaps(test.xml)
		if err != nil {
			t.Errorf("decode error:\n"+
				"input: %s\n"+
				"error:  %s\n",
				test.xml.EncodeString(nil), err)
			continue
		}

		if !reflect.DeepEqual(caps, test.caps) {
			t.Errorf("decode mismatch:\n"+
				"expected: %#v\n"+
				"present:  %#v\n",
				test.caps, caps)
		}
	}
}
