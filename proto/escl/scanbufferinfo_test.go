// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scan Buffer Info test

package escl

import (
	"errors"
	"reflect"
	"testing"

	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// testScanBufferInfo contains example of the initialized
// ScanBufferInfo structure
var testScanBufferInfo = ScanBufferInfo{
	ScanSettings: testScanSettings,
	ImageWidth:   2551,
	ImageHeight:  3508,
	BytesPerLine: 7653, // 2551 * 3
}

// TestScanBufferInfo tests [ScanBufferInfo] conversion
// to and from the XML
func TestScanBufferInfo(t *testing.T) {
	type testData struct {
		info ScanBufferInfo
		xml  xmldoc.Element
	}

	tests := []testData{
		{
			// Full data test
			info: testScanBufferInfo,
			xml: xmldoc.WithChildren(
				NsScan+":ScanBufferInfo",
				testScanSettings.ToXML(),
				xmldoc.WithText(NsScan+":ImageWidth",
					"2551"),
				xmldoc.WithText(NsScan+":ImageHeight",
					"3508"),
				xmldoc.WithText(NsScan+":BytesPerLine",
					"7653"),
			),
		},
	}

	for _, test := range tests {
		xml := test.info.ToXML()
		if !xml.Similar(test.xml) {
			t.Errorf("encode mismatch:\n"+
				"expected: %s\n"+
				"present:  %s\n",
				test.xml.EncodeString(nil),
				xml.EncodeString(nil))
		}

		info, err := DecodeScanBufferInfo(test.xml)
		if err != nil {
			t.Errorf("decode error:\n"+
				"input: %s\n"+
				"error:  %s\n",
				test.xml.EncodeString(nil), err)
			continue
		}

		if !reflect.DeepEqual(info, test.info) {
			t.Errorf("decode mismatch:\n"+
				"expected: %#v\n"+
				"present:  %#v\n",
				test.info, info)
		}
	}
}

// TestScanBufferInfoDecodeErrors tests [ScanBufferInfo] XML decode
// errors handling
func TestScanBufferInfoDecodeErrors(t *testing.T) {
	type testData struct {
		xml xmldoc.Element
		err string
	}

	tests := []testData{
		// Tests for missed elements
		{
			// Missed ScanSettings
			xml: xmldoc.WithChildren(
				NsScan+":ScanBufferInfo",
				xmldoc.WithText(NsScan+":ImageWidth",
					"2551"),
				xmldoc.WithText(NsScan+":ImageHeight",
					"3508"),
				xmldoc.WithText(NsScan+":BytesPerLine",
					"7653"),
			),
			err: `/scan:ScanBufferInfo/scan:ScanSettings: missed`,
		},

		{
			// Missed ImageWidth
			xml: xmldoc.WithChildren(
				NsScan+":ScanBufferInfo",
				testScanSettings.ToXML(),
				xmldoc.WithText(NsScan+":ImageHeight",
					"3508"),
				xmldoc.WithText(NsScan+":BytesPerLine",
					"7653"),
			),
			err: `/scan:ScanBufferInfo/scan:ImageWidth: missed`,
		},

		{
			// Missed ImageHeight
			xml: xmldoc.WithChildren(
				NsScan+":ScanBufferInfo",
				testScanSettings.ToXML(),
				xmldoc.WithText(NsScan+":ImageWidth",
					"2551"),
				xmldoc.WithText(NsScan+":BytesPerLine",
					"7653"),
			),
			err: `/scan:ScanBufferInfo/scan:ImageHeight: missed`,
		},

		{
			// Missed BytesPerLine
			xml: xmldoc.WithChildren(
				NsScan+":ScanBufferInfo",
				testScanSettings.ToXML(),
				xmldoc.WithText(NsScan+":ImageWidth",
					"2551"),
				xmldoc.WithText(NsScan+":ImageHeight",
					"3508"),
			),
			err: `/scan:ScanBufferInfo/scan:BytesPerLine: missed`,
		},

		// Tests for invalid elements
		{
			// Invalid ScanSettings
			xml: xmldoc.WithChildren(
				NsScan+":ScanBufferInfo",
				xmldoc.WithText(NsScan+":ScanSettings", "bad"),
				xmldoc.WithText(NsScan+":ImageWidth",
					"2551"),
				xmldoc.WithText(NsScan+":ImageHeight",
					"3508"),
				xmldoc.WithText(NsScan+":BytesPerLine",
					"7653"),
			),
			err: `/scan:ScanBufferInfo/scan:ScanSettings/pwg:Version: missed`,
		},

		{
			// Invalid ImageWidth
			xml: xmldoc.WithChildren(
				NsScan+":ScanBufferInfo",
				testScanSettings.ToXML(),
				xmldoc.WithText(NsScan+":ImageWidth", "bad"),
				xmldoc.WithText(NsScan+":ImageHeight",
					"3508"),
				xmldoc.WithText(NsScan+":BytesPerLine",
					"7653"),
			),
			err: `/scan:ScanBufferInfo/scan:ImageWidth: invalid int: "bad"`,
		},

		{
			// Invalid ImageHeight
			xml: xmldoc.WithChildren(
				NsScan+":ScanBufferInfo",
				testScanSettings.ToXML(),
				xmldoc.WithText(NsScan+":ImageWidth",
					"2551"),
				xmldoc.WithText(NsScan+":ImageHeight", "bad"),
				xmldoc.WithText(NsScan+":BytesPerLine",
					"7653"),
			),
			err: `/scan:ScanBufferInfo/scan:ImageHeight: invalid int: "bad"`,
		},

		{
			// Invalid BytesPerLine
			xml: xmldoc.WithChildren(
				NsScan+":ScanBufferInfo",
				testScanSettings.ToXML(),
				xmldoc.WithText(NsScan+":ImageWidth",
					"2551"),
				xmldoc.WithText(NsScan+":ImageHeight",
					"3508"),
				xmldoc.WithText(NsScan+":BytesPerLine", "bad"),
			),
			err: `/scan:ScanBufferInfo/scan:BytesPerLine: invalid int: "bad"`,
		},
	}

	for _, test := range tests {
		_, err := DecodeScanBufferInfo(test.xml)
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
