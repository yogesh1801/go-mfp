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
	CcdChannel:                   optional.New(NTSC),
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

	tests := []testData{}

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
