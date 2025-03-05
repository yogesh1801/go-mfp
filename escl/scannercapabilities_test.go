// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scanner capabilities test

package escl

import (
	"errors"
	"reflect"
	"testing"

	"github.com/alexpevzner/mfp/util/optional"
	"github.com/alexpevzner/mfp/util/uuid"
	"github.com/alexpevzner/mfp/util/xmldoc"
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

// TestScannerCapabilities tests [ScannerCapabilities] conversion
// to and from the XML
func TestScannerCapabilities(t *testing.T) {
	type testData struct {
		scancaps ScannerCapabilities
		xml      xmldoc.Element
	}

	tests := []testData{
		// Full-data test
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

		// Tests to see that Platen/Camera/ADF capabilites
		// are not messed up
		{
			scancaps: ScannerCapabilities{
				Version: MakeVersion(2, 0),
				Platen:  optional.New(testPlaten),
			},
			xml: xmldoc.WithChildren(
				NsScan+":ScannerCapabilities",
				xmldoc.WithText(NsPWG+":Version", "2.0"),
				testPlaten.toXML(NsScan+":Platen"),
			),
		},

		{
			scancaps: ScannerCapabilities{
				Version: MakeVersion(2, 0),
				Camera:  optional.New(testCamera),
			},
			xml: xmldoc.WithChildren(
				NsScan+":ScannerCapabilities",
				xmldoc.WithText(NsPWG+":Version", "2.0"),
				testCamera.toXML(NsScan+":Camera"),
			),
		},

		{
			scancaps: ScannerCapabilities{
				Version: MakeVersion(2, 0),
				ADF:     optional.New(testADF),
			},
			xml: xmldoc.WithChildren(
				NsScan+":ScannerCapabilities",
				xmldoc.WithText(NsPWG+":Version", "2.0"),
				testADF.toXML(NsScan+":Adf"),
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

// TestScannerCapabilitiesDecodeErrors tests [ScannerCapabilities] XML decode
// errors handling
func TestScannerCapabilitiesDecodeErrors(t *testing.T) {
	type testData struct {
		xml xmldoc.Element
		err string
	}

	tests := []testData{
		// Missed required fields
		{
			xml: xmldoc.WithChildren(
				NsScan + ":ScannerCapabilities",
			),
			err: `/scan:ScannerCapabilities/pwg:Version: missed`,
		},

		// Bad data, field by field
		{
			xml: xmldoc.WithChildren(
				NsScan+":ScannerCapabilities",
				xmldoc.WithText(NsPWG+":Version", "bad"),
			),
			err: `/scan:ScannerCapabilities/pwg:Version: "bad": invalid eSCL version`,
		},

		{
			xml: xmldoc.WithChildren(
				NsScan+":ScannerCapabilities",
				xmldoc.WithText(NsPWG+":Version", "2.0"),
				xmldoc.WithText(NsScan+":UUID", "bad"),
			),
			err: `/scan:ScannerCapabilities/scan:UUID: invalid UUID: "bad"`,
		},

		{
			xml: xmldoc.WithChildren(
				NsScan+":ScannerCapabilities",
				xmldoc.WithText(NsPWG+":Version", "2.0"),
				xmldoc.WithChildren(NsScan+":SettingProfiles",
					xmldoc.WithText(
						NsScan+":SettingProfile", "bad"),
				),
			),
			err: `/scan:ScannerCapabilities/scan:SettingProfiles/scan:SettingProfile/scan:SupportedResolutions: missed`,
		},

		{
			xml: xmldoc.WithChildren(
				NsScan+":ScannerCapabilities",
				xmldoc.WithText(NsPWG+":Version", "2.0"),
				xmldoc.WithChildren(NsScan+":Platen",
					xmldoc.WithText(
						NsScan+":PlatenInputCaps", "bad"),
				),
			),
			err: `/scan:ScannerCapabilities/scan:Platen/scan:PlatenInputCaps/scan:MinWidth: missed`,
		},

		{
			xml: xmldoc.WithChildren(
				NsScan+":ScannerCapabilities",
				xmldoc.WithText(NsPWG+":Version", "2.0"),
				xmldoc.WithChildren(NsScan+":Camera",
					xmldoc.WithText(
						NsScan+":CameraInputCaps", "bad"),
				),
			),
			err: `/scan:ScannerCapabilities/scan:Camera/scan:CameraInputCaps/scan:MinWidth: missed`,
		},

		{
			xml: xmldoc.WithChildren(
				NsScan+":ScannerCapabilities",
				xmldoc.WithText(NsPWG+":Version", "2.0"),
				xmldoc.WithChildren(NsScan+":Adf",
					xmldoc.WithText(
						NsScan+":AdfSimplexInputCaps", "bad"),
				),
			),
			err: `/scan:ScannerCapabilities/scan:Adf/scan:AdfSimplexInputCaps/scan:MinWidth: missed`,
		},

		{
			xml: xmldoc.WithChildren(
				NsScan+":ScannerCapabilities",
				xmldoc.WithText(NsPWG+":Version", "2.0"),
				xmldoc.WithText(NsScan+":BrightnessSupport", "bad"),
			),
			err: `/scan:ScannerCapabilities/scan:BrightnessSupport/scan:Min: missed`,
		},

		{
			xml: xmldoc.WithChildren(
				NsScan+":ScannerCapabilities",
				xmldoc.WithText(NsPWG+":Version", "2.0"),
				xmldoc.WithText(NsScan+":CompressionFactorSupport", "bad"),
			),
			err: `/scan:ScannerCapabilities/scan:CompressionFactorSupport/scan:Min: missed`,
		},

		{
			xml: xmldoc.WithChildren(
				NsScan+":ScannerCapabilities",
				xmldoc.WithText(NsPWG+":Version", "2.0"),
				xmldoc.WithText(NsScan+":ContrastSupport", "bad"),
			),
			err: `/scan:ScannerCapabilities/scan:ContrastSupport/scan:Min: missed`,
		},

		{
			xml: xmldoc.WithChildren(
				NsScan+":ScannerCapabilities",
				xmldoc.WithText(NsPWG+":Version", "2.0"),
				xmldoc.WithText(NsScan+":GammaSupport", "bad"),
			),
			err: `/scan:ScannerCapabilities/scan:GammaSupport/scan:Min: missed`,
		},

		{
			xml: xmldoc.WithChildren(
				NsScan+":ScannerCapabilities",
				xmldoc.WithText(NsPWG+":Version", "2.0"),
				xmldoc.WithText(NsScan+":HighlightSupport", "bad"),
			),
			err: `/scan:ScannerCapabilities/scan:HighlightSupport/scan:Min: missed`,
		},

		{
			xml: xmldoc.WithChildren(
				NsScan+":ScannerCapabilities",
				xmldoc.WithText(NsPWG+":Version", "2.0"),
				xmldoc.WithText(NsScan+":NoiseRemovalSupport", "bad"),
			),
			err: `/scan:ScannerCapabilities/scan:NoiseRemovalSupport/scan:Min: missed`,
		},

		{
			xml: xmldoc.WithChildren(
				NsScan+":ScannerCapabilities",
				xmldoc.WithText(NsPWG+":Version", "2.0"),
				xmldoc.WithText(NsScan+":ShadowSupport", "bad"),
			),
			err: `/scan:ScannerCapabilities/scan:ShadowSupport/scan:Min: missed`,
		},

		{
			xml: xmldoc.WithChildren(
				NsScan+":ScannerCapabilities",
				xmldoc.WithText(NsPWG+":Version", "2.0"),
				xmldoc.WithText(NsScan+":SharpenSupport", "bad"),
			),
			err: `/scan:ScannerCapabilities/scan:SharpenSupport/scan:Min: missed`,
		},

		{
			xml: xmldoc.WithChildren(
				NsScan+":ScannerCapabilities",
				xmldoc.WithText(NsPWG+":Version", "2.0"),
				xmldoc.WithText(NsScan+":ThresholdSupport", "bad"),
			),
			err: `/scan:ScannerCapabilities/scan:ThresholdSupport/scan:Min: missed`,
		},

		{
			xml: xmldoc.WithChildren(
				NsScan+":ScannerCapabilities",
				xmldoc.WithText(NsPWG+":Version", "2.0"),
				xmldoc.WithText(NsScan+":BlankPageDetection", "bad"),
			),
			err: `/scan:ScannerCapabilities/scan:BlankPageDetection: invalid bool: "bad"`,
		},

		{
			xml: xmldoc.WithChildren(
				NsScan+":ScannerCapabilities",
				xmldoc.WithText(NsPWG+":Version", "2.0"),
				xmldoc.WithText(NsScan+":BlankPageDetectionAndRemoval", "bad"),
			),
			err: `/scan:ScannerCapabilities/scan:BlankPageDetectionAndRemoval: invalid bool: "bad"`,
		},
	}

	for _, test := range tests {
		_, err := DecodeScannerCapabilities(test.xml)
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
