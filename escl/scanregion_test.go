// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scan region test.

package escl

import (
	"errors"
	"reflect"
	"testing"

	"github.com/alexpevzner/mfp/util/xmldoc"
)

// testScanRegion contains example of the initialized
// ScanRegion structure
var testScanRegion = ScanRegion{
	XOffset:            5,
	YOffset:            10,
	Width:              2551,
	Height:             3508,
	ContentRegionUnits: ThreeHundredthsOfInches,
}

// TestScanRegion tests ScanRegion
func TestScanRegion(t *testing.T) {
	type testData struct {
		reg ScanRegion
		xml xmldoc.Element
	}

	tests := []testData{
		{
			// Full-data test
			reg: testScanRegion,
			xml: xmldoc.WithChildren(
				NsPWG+":ScanRegion",
				xmldoc.WithText(NsPWG+":XOffset", "5"),
				xmldoc.WithText(NsPWG+":YOffset", "10"),
				xmldoc.WithText(NsPWG+":Width", "2551"),
				xmldoc.WithText(NsPWG+":Height", "3508"),
				xmldoc.WithText(NsPWG+":ContentRegionUnits",
					"escl:ThreeHundredthsOfInches"),
			),
		},
	}

	for _, test := range tests {
		xml := test.reg.toXML(NsPWG + ":ScanRegion")
		if !reflect.DeepEqual(xml, test.xml) {
			t.Errorf("ToXML:\nexpected: %s\npresent: %s\n",
				test.xml.EncodeString(NsMap),
				xml.EncodeString(NsMap))
		}

		reg, err := decodeScanRegion(xml)
		if err != nil {
			t.Errorf("decodeScanRegion: %s", err)
			continue
		}

		if !reflect.DeepEqual(reg, test.reg) {
			t.Errorf("decodeScanRegion:\n"+
				"expected: %#v\npresent:  %#v\n",
				test.reg, reg)
		}
	}
}

// TestScanRegionDecodeErrors tests ScanRegion decode errors
func TestScanRegionDecodeErrors(t *testing.T) {
	type testData struct {
		xml xmldoc.Element
		err string
	}

	tests := []testData{
		// Tests for missed required elements
		{
			// Missed XOffset
			xml: xmldoc.WithChildren(
				NsPWG+":ScanRegion",
				xmldoc.WithText(NsPWG+":YOffset", "10"),
				xmldoc.WithText(NsPWG+":Width", "2551"),
				xmldoc.WithText(NsPWG+":Height", "3508"),
				xmldoc.WithText(NsPWG+":ContentRegionUnits",
					"escl:ThreeHundredthsOfInches"),
			),
			err: `/pwg:ScanRegion/pwg:XOffset: missed`,
		},

		{
			// Missed YOffset
			xml: xmldoc.WithChildren(
				NsPWG+":ScanRegion",
				xmldoc.WithText(NsPWG+":XOffset", "5"),
				xmldoc.WithText(NsPWG+":Width", "2551"),
				xmldoc.WithText(NsPWG+":Height", "3508"),
				xmldoc.WithText(NsPWG+":ContentRegionUnits",
					"escl:ThreeHundredthsOfInches"),
			),
			err: `/pwg:ScanRegion/pwg:YOffset: missed`,
		},

		{
			// Missed Width
			xml: xmldoc.WithChildren(
				NsPWG+":ScanRegion",
				xmldoc.WithText(NsPWG+":XOffset", "5"),
				xmldoc.WithText(NsPWG+":YOffset", "10"),
				xmldoc.WithText(NsPWG+":Height", "3508"),
				xmldoc.WithText(NsPWG+":ContentRegionUnits",
					"escl:ThreeHundredthsOfInches"),
			),
			err: `/pwg:ScanRegion/pwg:Width: missed`,
		},

		{
			// Missed Height
			xml: xmldoc.WithChildren(
				NsPWG+":ScanRegion",
				xmldoc.WithText(NsPWG+":XOffset", "5"),
				xmldoc.WithText(NsPWG+":YOffset", "10"),
				xmldoc.WithText(NsPWG+":Width", "2551"),
				xmldoc.WithText(NsPWG+":ContentRegionUnits",
					"escl:ThreeHundredthsOfInches"),
			),
			err: `/pwg:ScanRegion/pwg:Height: missed`,
		},

		{
			// Missed ContentRegionUnits
			xml: xmldoc.WithChildren(
				NsPWG+":ScanRegion",
				xmldoc.WithText(NsPWG+":XOffset", "5"),
				xmldoc.WithText(NsPWG+":YOffset", "10"),
				xmldoc.WithText(NsPWG+":Width", "2551"),
				xmldoc.WithText(NsPWG+":Height", "3508"),
			),
			err: `/pwg:ScanRegion/pwg:ContentRegionUnits: missed`,
		},

		// Tests for invalid elements data
		{
			// Invalid XOffset
			xml: xmldoc.WithChildren(
				NsPWG+":ScanRegion",
				xmldoc.WithText(NsPWG+":XOffset", "bad"),
				xmldoc.WithText(NsPWG+":YOffset", "10"),
				xmldoc.WithText(NsPWG+":Width", "2551"),
				xmldoc.WithText(NsPWG+":Height", "3508"),
				xmldoc.WithText(NsPWG+":ContentRegionUnits",
					"escl:ThreeHundredthsOfInches"),
			),
			err: `/pwg:ScanRegion/pwg:XOffset: invalid int: "bad"`,
		},

		{
			// Invalid YOffset
			xml: xmldoc.WithChildren(
				NsPWG+":ScanRegion",
				xmldoc.WithText(NsPWG+":XOffset", "5"),
				xmldoc.WithText(NsPWG+":YOffset", "bad"),
				xmldoc.WithText(NsPWG+":Width", "2551"),
				xmldoc.WithText(NsPWG+":Height", "3508"),
				xmldoc.WithText(NsPWG+":ContentRegionUnits",
					"escl:ThreeHundredthsOfInches"),
			),
			err: `/pwg:ScanRegion/pwg:YOffset: invalid int: "bad"`,
		},

		{
			// Invalid Width
			xml: xmldoc.WithChildren(
				NsPWG+":ScanRegion",
				xmldoc.WithText(NsPWG+":XOffset", "5"),
				xmldoc.WithText(NsPWG+":YOffset", "10"),
				xmldoc.WithText(NsPWG+":Width", "bad"),
				xmldoc.WithText(NsPWG+":Height", "3508"),
				xmldoc.WithText(NsPWG+":ContentRegionUnits",
					"escl:ThreeHundredthsOfInches"),
			),
			err: `/pwg:ScanRegion/pwg:Width: invalid int: "bad"`,
		},

		{
			// Invalid Height
			xml: xmldoc.WithChildren(
				NsPWG+":ScanRegion",
				xmldoc.WithText(NsPWG+":XOffset", "5"),
				xmldoc.WithText(NsPWG+":YOffset", "10"),
				xmldoc.WithText(NsPWG+":Width", "2551"),
				xmldoc.WithText(NsPWG+":Height", "bad"),
				xmldoc.WithText(NsPWG+":ContentRegionUnits",
					"escl:ThreeHundredthsOfInches"),
			),
			err: `/pwg:ScanRegion/pwg:Height: invalid int: "bad"`,
		},

		{
			// Invalid XOffset
			xml: xmldoc.WithChildren(
				NsPWG+":ScanRegion",
				xmldoc.WithText(NsPWG+":XOffset", "5"),
				xmldoc.WithText(NsPWG+":YOffset", "10"),
				xmldoc.WithText(NsPWG+":Width", "2551"),
				xmldoc.WithText(NsPWG+":Height", "3508"),
				xmldoc.WithText(NsPWG+":ContentRegionUnits",
					"bad"),
			),
			err: `/pwg:ScanRegion/pwg:ContentRegionUnits: invalid Units: "bad"`,
		},
	}

	for _, test := range tests {
		_, err := decodeScanRegion(test.xml)
		if err == nil {
			err = errors.New("")
		}

		if err.Error() != test.err {
			t.Errorf("%s\nexpected: %q\npresent:  %q",
				test.xml.EncodeString(NsMap),
				test.err, err)
		}
	}
}
