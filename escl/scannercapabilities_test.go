// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scanner capabilities test

package escl

import (
	"reflect"
	"testing"

	"github.com/alexpevzner/mfp/optional"
	"github.com/alexpevzner/mfp/uuid"
	"github.com/alexpevzner/mfp/xmldoc"
)

// testScannerCapabilities contains example of the initialized
// ScannerCapabilities structure
var testScannerCapabilities = ScannerCapabilities{
	Version:      MakeVersion(2, 0),
	MakeAndModel: optional.New("Example scanner"),
	SerialNumber: optional.New("00-123456"),
	Manufacturer: optional.New("Example printers, LTD"),
	UUID: optional.New(
		uuid.Must(uuid.Parse("00000000-1111-2222-3333-333333333333"))),
	AdminURI:                     optional.New("http://example.com"),
	IconURI:                      optional.New("http://example.com/p.jpg"),
	SettingProfiles:              []SettingProfile{testSettingProfile},
	Platen:                       optional.New(testPlaten),
	Camera:                       optional.New(testCamera),
	ADF:                          optional.New(testADF),
	BrightnessSupport:            optional.New(Range{0, 100, 80, nil}),
	CompressionFactorSupport:     optional.New(Range{1, 5, 3, nil}),
	ContrastSupport:              optional.New(Range{0, 100, 50, nil}),
	GammaSupport:                 optional.New(Range{1, 40, 20, nil}),
	HighlightSupport:             optional.New(Range{0, 100, 60, nil}),
	NoiseRemovalSupport:          optional.New(Range{0, 10, 2, nil}),
	ShadowSupport:                optional.New(Range{0, 100, 10, nil}),
	SharpenSupport:               optional.New(Range{0, 100, 15, nil}),
	ThresholdSupport:             optional.New(Range{0, 100, 50, nil}),
	BlankPageDetection:           optional.New(true),
	BlankPageDetectionAndRemoval: optional.New(true),
}

// TestScannerCapabilities tests [ScannerCapabilities] conversion to and from the XML
func TestScannerCapabilities(t *testing.T) {
	type testData struct {
		scancaps ScannerCapabilities
		xml      xmldoc.Element
	}

	tests := []testData{
		{
			scancaps: testScannerCapabilities,
			xml: xmldoc.WithChildren(
				NsScan+":ScannerCapabilities",
				xmldoc.WithText(NsPWG+":Version", "2.0"),
				xmldoc.WithText(NsPWG+":MakeAndModel",
					"Example scanner"),
				xmldoc.WithText(NsPWG+":SerialNumber",
					"00-123456"),
				xmldoc.WithText(NsPWG+":Manufacturer",
					"Example printers, LTD"),
				xmldoc.WithText(NsScan+":UUID",
					"00000000-1111-2222-3333-333333333333"),
				xmldoc.WithText(NsScan+":AdminURI",
					"http://example.com"),
				xmldoc.WithText(NsScan+":IconURI",
					"http://example.com/p.jpg"),
				xmldoc.WithChildren(NsScan+":SettingProfiles",
					testSettingProfile.toXML(
						NsScan+":SettingProfile"),
				),
				testPlaten.toXML(NsScan+":Platen"),
				testCamera.toXML(NsScan+":Camera"),
				testADF.toXML(NsScan+":Adf"),
				Range{0, 100, 80, nil}.toXML(
					NsScan+":BrightnessSupport"),
				Range{1, 5, 3, nil}.toXML(
					NsScan+":CompressionFactorSupport"),
				Range{0, 100, 50, nil}.toXML(
					NsScan+":ContrastSupport"),
				Range{1, 40, 20, nil}.toXML(
					NsScan+":GammaSupport"),
				Range{0, 100, 60, nil}.toXML(
					NsScan+":HighlightSupport"),
				Range{0, 10, 2, nil}.toXML(
					NsScan+":NoiseRemovalSupport"),
				Range{0, 100, 10, nil}.toXML(
					NsScan+":ShadowSupport"),
				Range{0, 100, 15, nil}.toXML(
					NsScan+":SharpenSupport"),
				Range{0, 100, 50, nil}.toXML(
					NsScan+":ThresholdSupport"),
				xmldoc.WithText(NsScan+":BlankPageDetection",
					"true"),
				xmldoc.WithText(NsScan+":BlankPageDetectionAndRemoval",
					"true"),
			),
		},
	}

	for _, test := range tests {
		xml := test.scancaps.ToXML()
		if !xml.Similar(test.xml) {
			t.Errorf("encode mismatch:\n"+
				"expected: %s\n"+
				"present:  %s\n",
				test.xml.EncodeString(nil),
				xml.EncodeString(nil))
		}

		scancaps, err := DecodeScannerCapabilities(test.xml)
		if err != nil {
			t.Errorf("decode error:\n"+
				"input: %s\n"+
				"error:  %s\n",
				test.xml.EncodeString(nil), err)
			continue
		}

		if !reflect.DeepEqual(scancaps, test.scancaps) {
			t.Errorf("decode mismatch:\n"+
				"expected: %#v\n"+
				"present:  %#v\n",
				test.scancaps, scancaps)
		}
	}
}
