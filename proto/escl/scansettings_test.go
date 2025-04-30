// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scan settings test

package escl

import (
	"errors"
	"reflect"
	"testing"

	"github.com/alexpevzner/mfp/util/optional"
	"github.com/alexpevzner/mfp/util/xmldoc"
)

// testScanSettings contains example of the initialized
// ScanSettings structure
var testScanSettings = ScanSettings{
	Version:                      MakeVersion(2, 0),
	Intent:                       optional.New(Document),
	ScanRegions:                  []ScanRegion{testScanRegion},
	DocumentFormat:               optional.New("application/pdf"),
	DocumentFormatExt:            optional.New("application/pdf"),
	ContentType:                  optional.New(ContentTypeText),
	InputSource:                  optional.New(InputFeeder),
	XResolution:                  optional.New(200),
	YResolution:                  optional.New(400),
	ColorMode:                    optional.New(RGB24),
	ColorSpace:                   optional.New(SRGB),
	CCDChannel:                   optional.New(NTSC),
	BinaryRendering:              optional.New(Threshold),
	Duplex:                       optional.New(true),
	FeedDirection:                optional.New(ShortEdgeFeed),
	Brightness:                   optional.New(80),
	CompressionFactor:            optional.New(3),
	Contrast:                     optional.New(50),
	Gamma:                        optional.New(20),
	Highlight:                    optional.New(60),
	NoiseRemoval:                 optional.New(2),
	Shadow:                       optional.New(10),
	Sharpen:                      optional.New(15),
	Threshold:                    optional.New(50),
	BlankPageDetection:           optional.New(true),
	BlankPageDetectionAndRemoval: optional.New(false),
}

// TestScanSettings tests [ScanSettings] conversion
// to and from the XML
func TestScanSettings(t *testing.T) {
	type testData struct {
		ss  ScanSettings
		xml xmldoc.Element
	}

	tests := []testData{
		{
			// Full-data test
			ss: testScanSettings,
			xml: xmldoc.WithChildren(
				NsScan+":ScanSettings",
				xmldoc.WithText(NsPWG+":Version", "2.0"),
				Document.toXML(NsScan+":Intent"),
				xmldoc.WithChildren(NsPWG+":ScanRegions",
					testScanRegion.toXML(NsPWG+":ScanRegion")),
				xmldoc.WithText(NsPWG+":DocumentFormat", "application/pdf"),
				xmldoc.WithText(NsScan+":DocumentFormatExt", "application/pdf"),
				ContentTypeText.toXML(NsPWG+":ContentType"),
				InputFeeder.toXML(NsPWG+":InputSource"),
				xmldoc.WithText(NsScan+":XResolution", "200"),
				xmldoc.WithText(NsScan+":YResolution", "400"),
				RGB24.toXML(NsScan+":ColorMode"),
				SRGB.toXML(NsScan+":ColorSpace"),
				NTSC.toXML(NsScan+":CcdChannel"),
				Threshold.toXML(NsScan+":BinaryRendering"),
				xmldoc.WithText(NsScan+":Duplex", "true"),
				ShortEdgeFeed.toXML(NsScan+":FeedDirection"),
				xmldoc.WithText(NsScan+":Brightness", "80"),
				xmldoc.WithText(NsScan+":CompressionFactor", "3"),
				xmldoc.WithText(NsScan+":Contrast", "50"),
				xmldoc.WithText(NsScan+":Gamma", "20"),
				xmldoc.WithText(NsScan+":Highlight", "60"),
				xmldoc.WithText(NsScan+":NoiseRemoval", "2"),
				xmldoc.WithText(NsScan+":Shadow", "10"),
				xmldoc.WithText(NsScan+":Sharpen", "15"),
				xmldoc.WithText(NsScan+":Threshold", "50"),
				xmldoc.WithText(NsScan+":BlankPageDetection", "true"),
				xmldoc.WithText(NsScan+":BlankPageDetectionAndRemoval", "false"),
			),
		},
	}

	for _, test := range tests {
		xml := test.ss.ToXML()
		if !xml.Similar(test.xml) {
			t.Errorf("encode mismatch:\n"+
				"expected: %s\n"+
				"present:  %s\n",
				test.xml.EncodeString(nil),
				xml.EncodeString(nil))
		}

		ss, err := DecodeScanSettings(test.xml)
		if err != nil {
			t.Errorf("decode error:\n"+
				"input: %s\n"+
				"error:  %s\n",
				test.xml.EncodeString(nil), err)
			continue
		}

		if !reflect.DeepEqual(ss, test.ss) {
			t.Errorf("decode mismatch:\n"+
				"expected: %#v\n"+
				"present:  %#v\n",
				test.ss, ss)
		}
	}
}

// TestScanSettingsDecodeErrors tests [ScanSettings] XML decode
// errors handling
func TestScanSettingsDecodeErrors(t *testing.T) {
	type testData struct {
		xml xmldoc.Element
		err string
	}

	tests := []testData{
		{
			// Missed required elements
			xml: xmldoc.WithChildren(
				NsScan + ":ScanSettings",
			),
			err: `/scan:ScanSettings/pwg:Version: missed`,
		},

		{
			// Invalid Version
			xml: xmldoc.WithChildren(
				NsScan+":ScanSettings",
				xmldoc.WithText(NsPWG+":Version", "bad"),
			),
			err: `/scan:ScanSettings/pwg:Version: "bad": invalid eSCL version`,
		},

		{
			// Invalid Intent
			xml: xmldoc.WithChildren(
				NsScan+":ScanSettings",
				xmldoc.WithText(NsPWG+":Version", "2.0"),
				xmldoc.WithText(NsScan+":Intent", "bad"),
			),
			err: `/scan:ScanSettings/scan:Intent: invalid Intent: "bad"`,
		},

		{
			// Invalid ScanRegions
			xml: xmldoc.WithChildren(
				NsScan+":ScanSettings",
				xmldoc.WithText(NsPWG+":Version", "2.0"),
				xmldoc.WithChildren(NsPWG+":ScanRegions",
					xmldoc.WithText(NsPWG+":ScanRegion", "bad"),
				),
			),
			err: `/scan:ScanSettings/pwg:ScanRegions/pwg:ScanRegion/pwg:XOffset: missed`,
		},

		{
			// Invalid ContentType
			xml: xmldoc.WithChildren(
				NsScan+":ScanSettings",
				xmldoc.WithText(NsPWG+":Version", "2.0"),
				xmldoc.WithText(NsPWG+":ContentType", "bad"),
			),
			err: `/scan:ScanSettings/pwg:ContentType: invalid ContentType: "bad"`,
		},

		{
			// Invalid InputSource
			xml: xmldoc.WithChildren(
				NsScan+":ScanSettings",
				xmldoc.WithText(NsPWG+":Version", "2.0"),
				xmldoc.WithText(NsPWG+":InputSource", "bad"),
			),
			err: `/scan:ScanSettings/pwg:InputSource: invalid InputSource: "bad"`,
		},

		{
			// Invalid XResolution
			xml: xmldoc.WithChildren(
				NsScan+":ScanSettings",
				xmldoc.WithText(NsPWG+":Version", "2.0"),
				xmldoc.WithText(NsScan+":XResolution", "bad"),
			),
			err: `/scan:ScanSettings/scan:XResolution: invalid int: "bad"`,
		},

		{
			// Invalid YResolution
			xml: xmldoc.WithChildren(
				NsScan+":ScanSettings",
				xmldoc.WithText(NsPWG+":Version", "2.0"),
				xmldoc.WithText(NsScan+":YResolution", "bad"),
			),
			err: `/scan:ScanSettings/scan:YResolution: invalid int: "bad"`,
		},

		{
			// Invalid ColorMode
			xml: xmldoc.WithChildren(
				NsScan+":ScanSettings",
				xmldoc.WithText(NsPWG+":Version", "2.0"),
				xmldoc.WithText(NsScan+":ColorMode", "bad"),
			),
			err: `/scan:ScanSettings/scan:ColorMode: invalid ColorMode: "bad"`,
		},

		{
			// Invalid ColorSpace
			xml: xmldoc.WithChildren(
				NsScan+":ScanSettings",
				xmldoc.WithText(NsPWG+":Version", "2.0"),
				xmldoc.WithText(NsScan+":ColorSpace", "bad"),
			),
			err: `/scan:ScanSettings/scan:ColorSpace: invalid ColorSpace: "bad"`,
		},

		{
			// Invalid CCDChannel
			xml: xmldoc.WithChildren(
				NsScan+":ScanSettings",
				xmldoc.WithText(NsPWG+":Version", "2.0"),
				xmldoc.WithText(NsScan+":CcdChannel", "bad"),
			),
			err: `/scan:ScanSettings/scan:CcdChannel: invalid CCDChannel: "bad"`,
		},

		{
			// Invalid BinaryRendering
			xml: xmldoc.WithChildren(
				NsScan+":ScanSettings",
				xmldoc.WithText(NsPWG+":Version", "2.0"),
				xmldoc.WithText(NsScan+":BinaryRendering", "bad"),
			),
			err: `/scan:ScanSettings/scan:BinaryRendering: invalid BinaryRendering: "bad"`,
		},

		{
			// Invalid Duplex
			xml: xmldoc.WithChildren(
				NsScan+":ScanSettings",
				xmldoc.WithText(NsPWG+":Version", "2.0"),
				xmldoc.WithText(NsScan+":Duplex", "bad"),
			),
			err: `/scan:ScanSettings/scan:Duplex: invalid bool: "bad"`,
		},

		{
			// Invalid Duplex
			xml: xmldoc.WithChildren(
				NsScan+":ScanSettings",
				xmldoc.WithText(NsPWG+":Version", "2.0"),
				xmldoc.WithText(NsScan+":FeedDirection", "bad"),
			),
			err: `/scan:ScanSettings/scan:FeedDirection: invalid FeedDirection: "bad"`,
		},

		{
			// Invalid Brightness
			xml: xmldoc.WithChildren(
				NsScan+":ScanSettings",
				xmldoc.WithText(NsPWG+":Version", "2.0"),
				xmldoc.WithText(NsScan+":Brightness", "bad"),
			),
			err: `/scan:ScanSettings/scan:Brightness: invalid int: "bad"`,
		},

		{
			// Invalid CompressionFactor
			xml: xmldoc.WithChildren(
				NsScan+":ScanSettings",
				xmldoc.WithText(NsPWG+":Version", "2.0"),
				xmldoc.WithText(NsScan+":CompressionFactor", "bad"),
			),
			err: `/scan:ScanSettings/scan:CompressionFactor: invalid int: "bad"`,
		},

		{
			// Invalid Contrast
			xml: xmldoc.WithChildren(
				NsScan+":ScanSettings",
				xmldoc.WithText(NsPWG+":Version", "2.0"),
				xmldoc.WithText(NsScan+":Contrast", "bad"),
			),
			err: `/scan:ScanSettings/scan:Contrast: invalid int: "bad"`,
		},

		{
			// Invalid Gamma
			xml: xmldoc.WithChildren(
				NsScan+":ScanSettings",
				xmldoc.WithText(NsPWG+":Version", "2.0"),
				xmldoc.WithText(NsScan+":Gamma", "bad"),
			),
			err: `/scan:ScanSettings/scan:Gamma: invalid int: "bad"`,
		},

		{
			// Invalid Highlight
			xml: xmldoc.WithChildren(
				NsScan+":ScanSettings",
				xmldoc.WithText(NsPWG+":Version", "2.0"),
				xmldoc.WithText(NsScan+":Highlight", "bad"),
			),
			err: `/scan:ScanSettings/scan:Highlight: invalid int: "bad"`,
		},

		{
			// Invalid NoiseRemoval
			xml: xmldoc.WithChildren(
				NsScan+":ScanSettings",
				xmldoc.WithText(NsPWG+":Version", "2.0"),
				xmldoc.WithText(NsScan+":NoiseRemoval", "bad"),
			),
			err: `/scan:ScanSettings/scan:NoiseRemoval: invalid int: "bad"`,
		},

		{
			// Invalid Shadow
			xml: xmldoc.WithChildren(
				NsScan+":ScanSettings",
				xmldoc.WithText(NsPWG+":Version", "2.0"),
				xmldoc.WithText(NsScan+":Shadow", "bad"),
			),
			err: `/scan:ScanSettings/scan:Shadow: invalid int: "bad"`,
		},

		{
			// Invalid Sharpen
			xml: xmldoc.WithChildren(
				NsScan+":ScanSettings",
				xmldoc.WithText(NsPWG+":Version", "2.0"),
				xmldoc.WithText(NsScan+":Sharpen", "bad"),
			),
			err: `/scan:ScanSettings/scan:Sharpen: invalid int: "bad"`,
		},

		{
			// Invalid Threshold
			xml: xmldoc.WithChildren(
				NsScan+":ScanSettings",
				xmldoc.WithText(NsPWG+":Version", "2.0"),
				xmldoc.WithText(NsScan+":Threshold", "bad"),
			),
			err: `/scan:ScanSettings/scan:Threshold: invalid int: "bad"`,
		},

		{
			// Invalid BlankPageDetection
			xml: xmldoc.WithChildren(
				NsScan+":ScanSettings",
				xmldoc.WithText(NsPWG+":Version", "2.0"),
				xmldoc.WithText(NsScan+":BlankPageDetection", "bad"),
			),
			err: `/scan:ScanSettings/scan:BlankPageDetection: invalid bool: "bad"`,
		},

		{
			// Invalid BlankPageDetectionAndRemoval
			xml: xmldoc.WithChildren(
				NsScan+":ScanSettings",
				xmldoc.WithText(NsPWG+":Version", "2.0"),
				xmldoc.WithText(NsScan+":BlankPageDetectionAndRemoval", "bad"),
			),
			err: `/scan:ScanSettings/scan:BlankPageDetectionAndRemoval: invalid bool: "bad"`,
		},
	}

	for _, test := range tests {
		_, err := DecodeScanSettings(test.xml)
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
