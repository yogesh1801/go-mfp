// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scan Image Info test

package escl

import (
	"errors"
	"reflect"
	"testing"

	"github.com/alexpevzner/mfp/util/optional"
	"github.com/alexpevzner/mfp/util/xmldoc"
)

// testScanImageInfo contains example of the initialized
// ScanImageInfo structure
var testScanImageInfo = ScanImageInfo{
	JobURI:             "/eSCL/ScanJobs/urn:uuid:4509a320-00a0-008f-00b6-00559a327d32",
	JobUUID:            optional.New("urn:uuid:4509a320-00a0-008f-00b6-00559a327d32"),
	ActualWidth:        2551,
	ActualHeight:       3508,
	ActualBytesPerLine: 7653, // 2551 * 3
	BlankPageDetected:  optional.New(false),
}

// TestScanImageInfo tests [ScanImageInfo] conversion
// to and from the XML
func TestScanImageInfo(t *testing.T) {
	type testData struct {
		info ScanImageInfo
		xml  xmldoc.Element
	}

	tests := []testData{
		{
			// Full data test
			info: testScanImageInfo,
			xml: xmldoc.WithChildren(
				NsScan+":ScanImageInfo",
				xmldoc.WithText(NsPWG+":JobUri",
					"/eSCL/ScanJobs/urn:uuid:4509a320-00a0-008f-00b6-00559a327d32"),
				xmldoc.WithText(NsPWG+":JobUuid",
					"urn:uuid:4509a320-00a0-008f-00b6-00559a327d32"),
				xmldoc.WithText(NsScan+":ActualWidth",
					"2551"),
				xmldoc.WithText(NsScan+":ActualHeight",
					"3508"),
				xmldoc.WithText(NsScan+":ActualBytesPerLine",
					"7653"),
				xmldoc.WithText(NsScan+":BlankPageDetected",
					"false"),
			),
		},

		{
			// Missed optional elements
			info: ScanImageInfo{
				JobURI:             "/eSCL/ScanJobs/urn:uuid:4509a320-00a0-008f-00b6-00559a327d32",
				ActualWidth:        2551,
				ActualHeight:       3508,
				ActualBytesPerLine: 7653, // 2551 * 3
			},
			xml: xmldoc.WithChildren(
				NsScan+":ScanImageInfo",
				xmldoc.WithText(NsPWG+":JobUri",
					"/eSCL/ScanJobs/urn:uuid:4509a320-00a0-008f-00b6-00559a327d32"),
				xmldoc.WithText(NsScan+":ActualWidth",
					"2551"),
				xmldoc.WithText(NsScan+":ActualHeight",
					"3508"),
				xmldoc.WithText(NsScan+":ActualBytesPerLine",
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

		info, err := DecodeScanImageInfo(test.xml)
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

// TestScanImageInfoDecodeErrors tests [ScanImageInfo] XML decode
// errors handling
func TestScanImageInfoDecodeErrors(t *testing.T) {
	type testData struct {
		xml xmldoc.Element // Input XML
		err string         // Expected error
	}

	tests := []testData{
		// Missed required elements
		{
			// Missed JobURI
			xml: xmldoc.WithChildren(
				NsScan+":ScanImageInfo",
				xmldoc.WithText(NsScan+":ActualWidth",
					"2551"),
				xmldoc.WithText(NsScan+":ActualHeight",
					"3508"),
				xmldoc.WithText(NsScan+":ActualBytesPerLine",
					"7653"),
			),
			err: `/scan:ScanImageInfo/pwg:JobUri: missed`,
		},

		{
			// Missed ActualWidth
			xml: xmldoc.WithChildren(
				NsScan+":ScanImageInfo",
				xmldoc.WithText(NsPWG+":JobUri",
					"/eSCL/ScanJobs/urn:uuid:4509a320-00a0-008f-00b6-00559a327d32"),
				xmldoc.WithText(NsScan+":ActualHeight",
					"3508"),
				xmldoc.WithText(NsScan+":ActualBytesPerLine",
					"7653"),
			),
			err: `/scan:ScanImageInfo/scan:ActualWidth: missed`,
		},

		{
			// Missed ActualHeight
			xml: xmldoc.WithChildren(
				NsScan+":ScanImageInfo",
				xmldoc.WithText(NsPWG+":JobUri",
					"/eSCL/ScanJobs/urn:uuid:4509a320-00a0-008f-00b6-00559a327d32"),
				xmldoc.WithText(NsScan+":ActualWidth",
					"2551"),
				xmldoc.WithText(NsScan+":ActualBytesPerLine",
					"7653"),
			),
			err: `/scan:ScanImageInfo/scan:ActualHeight: missed`,
		},

		{
			// Missed ActualBytesPerLine
			xml: xmldoc.WithChildren(
				NsScan+":ScanImageInfo",
				xmldoc.WithText(NsPWG+":JobUri",
					"/eSCL/ScanJobs/urn:uuid:4509a320-00a0-008f-00b6-00559a327d32"),
				xmldoc.WithText(NsScan+":ActualWidth",
					"2551"),
				xmldoc.WithText(NsScan+":ActualHeight",
					"3508"),
			),
			err: `/scan:ScanImageInfo/scan:ActualBytesPerLine: missed`,
		},

		// Invalid elements
		{
			// Invalid ActualWidth
			xml: xmldoc.WithChildren(
				NsScan+":ScanImageInfo",
				xmldoc.WithText(NsPWG+":JobUri",
					"/eSCL/ScanJobs/urn:uuid:4509a320-00a0-008f-00b6-00559a327d32"),
				xmldoc.WithText(NsPWG+":JobUuid",
					"urn:uuid:4509a320-00a0-008f-00b6-00559a327d32"),
				xmldoc.WithText(NsScan+":ActualWidth",
					"bad"),
				xmldoc.WithText(NsScan+":ActualHeight",
					"3508"),
				xmldoc.WithText(NsScan+":ActualBytesPerLine",
					"7653"),
				xmldoc.WithText(NsScan+":BlankPageDetected",
					"false"),
			),
			err: `/scan:ScanImageInfo/scan:ActualWidth: invalid int: "bad"`,
		},
		{

			// Invalid ActualHeight
			xml: xmldoc.WithChildren(
				NsScan+":ScanImageInfo",
				xmldoc.WithText(NsPWG+":JobUri",
					"/eSCL/ScanJobs/urn:uuid:4509a320-00a0-008f-00b6-00559a327d32"),
				xmldoc.WithText(NsPWG+":JobUuid",
					"urn:uuid:4509a320-00a0-008f-00b6-00559a327d32"),
				xmldoc.WithText(NsScan+":ActualWidth",
					"2551"),
				xmldoc.WithText(NsScan+":ActualHeight",
					"bad"),
				xmldoc.WithText(NsScan+":ActualBytesPerLine",
					"7653"),
				xmldoc.WithText(NsScan+":BlankPageDetected",
					"false"),
			),
			err: `/scan:ScanImageInfo/scan:ActualHeight: invalid int: "bad"`,
		},
		{

			// Invalid ActualBytesPerLine
			xml: xmldoc.WithChildren(
				NsScan+":ScanImageInfo",
				xmldoc.WithText(NsPWG+":JobUri",
					"/eSCL/ScanJobs/urn:uuid:4509a320-00a0-008f-00b6-00559a327d32"),
				xmldoc.WithText(NsPWG+":JobUuid",
					"urn:uuid:4509a320-00a0-008f-00b6-00559a327d32"),
				xmldoc.WithText(NsScan+":ActualWidth",
					"2551"),
				xmldoc.WithText(NsScan+":ActualHeight",
					"3508"),
				xmldoc.WithText(NsScan+":ActualBytesPerLine",
					"bad"),
				xmldoc.WithText(NsScan+":BlankPageDetected",
					"false"),
			),
			err: `/scan:ScanImageInfo/scan:ActualBytesPerLine: invalid int: "bad"`,
		},
		{

			// Invalid BlankPageDetected
			xml: xmldoc.WithChildren(
				NsScan+":ScanImageInfo",
				xmldoc.WithText(NsPWG+":JobUri",
					"/eSCL/ScanJobs/urn:uuid:4509a320-00a0-008f-00b6-00559a327d32"),
				xmldoc.WithText(NsPWG+":JobUuid",
					"urn:uuid:4509a320-00a0-008f-00b6-00559a327d32"),
				xmldoc.WithText(NsScan+":ActualWidth",
					"2551"),
				xmldoc.WithText(NsScan+":ActualHeight",
					"3508"),
				xmldoc.WithText(NsScan+":ActualBytesPerLine",
					"7653"),
				xmldoc.WithText(NsScan+":BlankPageDetected",
					"bad"),
			),
			err: `/scan:ScanImageInfo/scan:BlankPageDetected: invalid bool: "bad"`,
		},
	}

	for _, test := range tests {
		_, err := DecodeScanImageInfo(test.xml)
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
