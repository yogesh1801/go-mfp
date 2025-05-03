// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Input source capabilities test

package escl

import (
	"errors"
	"reflect"
	"testing"

	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
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
				xmldoc.WithText(NsScan+":MinWidth", "591"),
				xmldoc.WithText(NsScan+":MaxWidth", "2551"),
				xmldoc.WithText(NsScan+":MinHeight", "600"),
				xmldoc.WithText(NsScan+":MaxHeight", "4205"),
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
		if !xml.Similar(test.xml) {
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

// TestInputSourceCapsDecodeErrors tests [InputSourceCaps] XML decode
// errors handling
func TestInputSourceCapsDecodeErrors(t *testing.T) {
	type testData struct {
		xml xmldoc.Element
		err string
	}

	tests := []testData{
		// Test for missed elements handling
		{
			xml: xmldoc.WithChildren(
				NsScan + ":PlatenInputCaps",
			),
			err: `/scan:PlatenInputCaps/scan:MinWidth: missed`,
		},

		{
			xml: xmldoc.WithChildren(
				NsScan+":PlatenInputCaps",
				xmldoc.WithText(NsScan+":MinWidth", "591"),
			),
			err: `/scan:PlatenInputCaps/scan:MaxWidth: missed`,
		},
		{
			xml: xmldoc.WithChildren(
				NsScan+":PlatenInputCaps",
				xmldoc.WithText(NsScan+":MinWidth", "591"),
				xmldoc.WithText(NsScan+":MaxWidth", "2551"),
			),
			err: `/scan:PlatenInputCaps/scan:MinHeight: missed`,
		},
		{
			xml: xmldoc.WithChildren(
				NsScan+":PlatenInputCaps",
				xmldoc.WithText(NsScan+":MinWidth", "591"),
				xmldoc.WithText(NsScan+":MaxWidth", "2551"),
				xmldoc.WithText(NsScan+":MinHeight", "600"),
			),
			err: `/scan:PlatenInputCaps/scan:MaxHeight: missed`,
		},
		{
			xml: xmldoc.WithChildren(
				NsScan+":PlatenInputCaps",
				xmldoc.WithText(NsScan+":MinWidth", "591"),
				xmldoc.WithText(NsScan+":MaxWidth", "2551"),
				xmldoc.WithText(NsScan+":MinHeight", "600"),
				xmldoc.WithText(NsScan+":MaxHeight", "4205"),
			),
			err: ``,
		},

		// Errors handling within nested integer elements
		{
			xml: xmldoc.WithChildren(
				NsScan+":PlatenInputCaps",
				xmldoc.WithText(NsScan+":MinWidth", "bad"),
				xmldoc.WithText(NsScan+":MaxWidth", "2551"),
				xmldoc.WithText(NsScan+":MinHeight", "600"),
				xmldoc.WithText(NsScan+":MaxHeight", "4205"),
			),
			err: `/scan:PlatenInputCaps/scan:MinWidth: invalid int: "bad"`,
		},

		{
			xml: xmldoc.WithChildren(
				NsScan+":PlatenInputCaps",
				xmldoc.WithText(NsScan+":MinWidth", "591"),
				xmldoc.WithText(NsScan+":MaxWidth", "bad"),
				xmldoc.WithText(NsScan+":MinHeight", "600"),
				xmldoc.WithText(NsScan+":MaxHeight", "4205"),
			),
			err: `/scan:PlatenInputCaps/scan:MaxWidth: invalid int: "bad"`,
		},

		{
			xml: xmldoc.WithChildren(
				NsScan+":PlatenInputCaps",
				xmldoc.WithText(NsScan+":MinWidth", "591"),
				xmldoc.WithText(NsScan+":MaxWidth", "2551"),
				xmldoc.WithText(NsScan+":MinHeight", "bad"),
				xmldoc.WithText(NsScan+":MaxHeight", "4205"),
			),
			err: `/scan:PlatenInputCaps/scan:MinHeight: invalid int: "bad"`,
		},

		{
			xml: xmldoc.WithChildren(
				NsScan+":PlatenInputCaps",
				xmldoc.WithText(NsScan+":MinWidth", "591"),
				xmldoc.WithText(NsScan+":MaxWidth", "2551"),
				xmldoc.WithText(NsScan+":MinHeight", "600"),
				xmldoc.WithText(NsScan+":MaxHeight", "bad"),
			),
			err: `/scan:PlatenInputCaps/scan:MaxHeight: invalid int: "bad"`,
		},

		{
			xml: xmldoc.WithChildren(
				NsScan+":PlatenInputCaps",
				xmldoc.WithText(NsScan+":MinWidth", "591"),
				xmldoc.WithText(NsScan+":MaxWidth", "2551"),
				xmldoc.WithText(NsScan+":MinHeight", "600"),
				xmldoc.WithText(NsScan+":MaxHeight", "4205"),
				xmldoc.WithText(NsScan+":MaxXOffset", "bad"),
			),
			err: `/scan:PlatenInputCaps/scan:MaxXOffset: invalid int: "bad"`,
		},

		{
			xml: xmldoc.WithChildren(
				NsScan+":PlatenInputCaps",
				xmldoc.WithText(NsScan+":MinWidth", "591"),
				xmldoc.WithText(NsScan+":MaxWidth", "2551"),
				xmldoc.WithText(NsScan+":MinHeight", "600"),
				xmldoc.WithText(NsScan+":MaxHeight", "4205"),
				xmldoc.WithText(NsScan+":MaxYOffset", "bad"),
			),
			err: `/scan:PlatenInputCaps/scan:MaxYOffset: invalid int: "bad"`,
		},

		{
			xml: xmldoc.WithChildren(
				NsScan+":PlatenInputCaps",
				xmldoc.WithText(NsScan+":MinWidth", "591"),
				xmldoc.WithText(NsScan+":MaxWidth", "2551"),
				xmldoc.WithText(NsScan+":MinHeight", "600"),
				xmldoc.WithText(NsScan+":MaxHeight", "4205"),
				xmldoc.WithText(NsScan+":MaxOpticalXResolution", "bad"),
			),
			err: `/scan:PlatenInputCaps/scan:MaxOpticalXResolution: invalid int: "bad"`,
		},

		{
			xml: xmldoc.WithChildren(
				NsScan+":PlatenInputCaps",
				xmldoc.WithText(NsScan+":MinWidth", "591"),
				xmldoc.WithText(NsScan+":MaxWidth", "2551"),
				xmldoc.WithText(NsScan+":MinHeight", "600"),
				xmldoc.WithText(NsScan+":MaxHeight", "4205"),
				xmldoc.WithText(NsScan+":MaxOpticalYResolution", "bad"),
			),
			err: `/scan:PlatenInputCaps/scan:MaxOpticalYResolution: invalid int: "bad"`,
		},

		{
			xml: xmldoc.WithChildren(
				NsScan+":PlatenInputCaps",
				xmldoc.WithText(NsScan+":MinWidth", "591"),
				xmldoc.WithText(NsScan+":MaxWidth", "2551"),
				xmldoc.WithText(NsScan+":MinHeight", "600"),
				xmldoc.WithText(NsScan+":MaxHeight", "4205"),
				xmldoc.WithText(NsScan+":MaxScanRegions", "bad"),
			),
			err: `/scan:PlatenInputCaps/scan:MaxScanRegions: invalid int: "bad"`,
		},

		{
			xml: xmldoc.WithChildren(
				NsScan+":PlatenInputCaps",
				xmldoc.WithText(NsScan+":MinWidth", "591"),
				xmldoc.WithText(NsScan+":MaxWidth", "2551"),
				xmldoc.WithText(NsScan+":MinHeight", "600"),
				xmldoc.WithText(NsScan+":MaxHeight", "4205"),
				xmldoc.WithText(NsScan+":RiskyLeftMargins", "bad"),
			),
			err: `/scan:PlatenInputCaps/scan:RiskyLeftMargins: invalid int: "bad"`,
		},

		{
			xml: xmldoc.WithChildren(
				NsScan+":PlatenInputCaps",
				xmldoc.WithText(NsScan+":MinWidth", "591"),
				xmldoc.WithText(NsScan+":MaxWidth", "2551"),
				xmldoc.WithText(NsScan+":MinHeight", "600"),
				xmldoc.WithText(NsScan+":MaxHeight", "4205"),
				xmldoc.WithText(NsScan+":RiskyRightMargins", "bad"),
			),
			err: `/scan:PlatenInputCaps/scan:RiskyRightMargins: invalid int: "bad"`,
		},

		{
			xml: xmldoc.WithChildren(
				NsScan+":PlatenInputCaps",
				xmldoc.WithText(NsScan+":MinWidth", "591"),
				xmldoc.WithText(NsScan+":MaxWidth", "2551"),
				xmldoc.WithText(NsScan+":MinHeight", "600"),
				xmldoc.WithText(NsScan+":MaxHeight", "4205"),
				xmldoc.WithText(NsScan+":RiskyTopMargins", "bad"),
			),
			err: `/scan:PlatenInputCaps/scan:RiskyTopMargins: invalid int: "bad"`,
		},

		{
			xml: xmldoc.WithChildren(
				NsScan+":PlatenInputCaps",
				xmldoc.WithText(NsScan+":MinWidth", "591"),
				xmldoc.WithText(NsScan+":MaxWidth", "2551"),
				xmldoc.WithText(NsScan+":MinHeight", "600"),
				xmldoc.WithText(NsScan+":MaxHeight", "4205"),
				xmldoc.WithText(NsScan+":RiskyBottomMargins", "bad"),
			),
			err: `/scan:PlatenInputCaps/scan:RiskyBottomMargins: invalid int: "bad"`,
		},

		{
			xml: xmldoc.WithChildren(
				NsScan+":PlatenInputCaps",
				xmldoc.WithText(NsScan+":MinWidth", "591"),
				xmldoc.WithText(NsScan+":MaxWidth", "2551"),
				xmldoc.WithText(NsScan+":MinHeight", "600"),
				xmldoc.WithText(NsScan+":MaxHeight", "4205"),
				xmldoc.WithText(NsScan+":MaxPhysicalWidth", "bad"),
			),
			err: `/scan:PlatenInputCaps/scan:MaxPhysicalWidth: invalid int: "bad"`,
		},

		{
			xml: xmldoc.WithChildren(
				NsScan+":PlatenInputCaps",
				xmldoc.WithText(NsScan+":MinWidth", "591"),
				xmldoc.WithText(NsScan+":MaxWidth", "2551"),
				xmldoc.WithText(NsScan+":MinHeight", "600"),
				xmldoc.WithText(NsScan+":MaxHeight", "4205"),
				xmldoc.WithText(NsScan+":MaxPhysicalHeight", "bad"),
			),
			err: `/scan:PlatenInputCaps/scan:MaxPhysicalHeight: invalid int: "bad"`,
		},

		// Errors handling within more complex nested elements
		{
			xml: xmldoc.WithChildren(
				NsScan+":PlatenInputCaps",
				xmldoc.WithText(NsScan+":MinWidth", "591"),
				xmldoc.WithText(NsScan+":MaxWidth", "2551"),
				xmldoc.WithText(NsScan+":MinHeight", "600"),
				xmldoc.WithText(NsScan+":MaxHeight", "4205"),
				xmldoc.WithChildren(NsScan+":SupportedIntents",
					xmldoc.WithText(NsScan+":SupportedIntent", "bad"),
				),
			),
			err: `/scan:PlatenInputCaps/scan:SupportedIntents/scan:SupportedIntent: invalid Intent: "bad"`,
		},

		{
			xml: xmldoc.WithChildren(
				NsScan+":PlatenInputCaps",
				xmldoc.WithText(NsScan+":MinWidth", "591"),
				xmldoc.WithText(NsScan+":MaxWidth", "2551"),
				xmldoc.WithText(NsScan+":MinHeight", "600"),
				xmldoc.WithText(NsScan+":MaxHeight", "4205"),
				xmldoc.WithChildren(NsScan+":EdgeAutoDetection",
					xmldoc.WithText(NsScan+":SupportedEdge", "bad"),
				),
			),
			err: `/scan:PlatenInputCaps/scan:EdgeAutoDetection/scan:SupportedEdge: invalid SupportedEdge: "bad"`,
		},

		{
			xml: xmldoc.WithChildren(
				NsScan+":PlatenInputCaps",
				xmldoc.WithText(NsScan+":MinWidth", "591"),
				xmldoc.WithText(NsScan+":MaxWidth", "2551"),
				xmldoc.WithText(NsScan+":MinHeight", "600"),
				xmldoc.WithText(NsScan+":MaxHeight", "4205"),
				xmldoc.WithChildren(NsScan+":FeedDirections",
					xmldoc.WithText(NsScan+":FeedDirection", "bad"),
				),
			),
			err: `/scan:PlatenInputCaps/scan:FeedDirections/scan:FeedDirection: invalid FeedDirection: "bad"`,
		},

		{
			xml: xmldoc.WithChildren(
				NsScan+":PlatenInputCaps",
				xmldoc.WithText(NsScan+":MinWidth", "591"),
				xmldoc.WithText(NsScan+":MaxWidth", "2551"),
				xmldoc.WithText(NsScan+":MinHeight", "600"),
				xmldoc.WithText(NsScan+":MaxHeight", "4205"),
				xmldoc.WithChildren(NsScan+":SettingProfiles",
					xmldoc.WithText(NsScan+":SettingProfile", "bad"),
				),
			),
			err: `/scan:PlatenInputCaps/scan:SettingProfiles/scan:SettingProfile/scan:SupportedResolutions: missed`,
		},
	}

	for _, test := range tests {
		_, err := decodeInputSourceCaps(test.xml)
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
