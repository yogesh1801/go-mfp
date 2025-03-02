// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scanner status test

package escl

import (
	"reflect"
	"testing"
	"time"

	"github.com/alexpevzner/mfp/optional"
	"github.com/alexpevzner/mfp/uuid"
	"github.com/alexpevzner/mfp/xmldoc"
)

// testScannerStatus contains example of the initialized
// ScannerStatus structure
var testScannerStatus = ScannerStatus{
	Version:  MakeVersion(2, 0),
	State:    ScannerProcessing,
	ADFState: optional.New(ScannerAdfProcessing),
	Jobs: []JobInfo{
		{
			JobURI:           "/eSCL/ScanJobs/urn:uuid:4509a320-00a0-008f-00b6-00559a327d32",
			JobUUID:          optional.New(uuid.Must(uuid.Parse("4509a320-00a0-008f-00b6-00559a327d32"))),
			Age:              optional.New(20 * time.Second),
			ImagesCompleted:  optional.New(2),
			ImagesToTransfer: optional.New(1),
			JobState:         JobProcessing,
			JobStateReasons:  []JobStateReason{JobScanningAndTransferring},
		},
		{
			JobURI:           "/eSCL/ScanJobs/urn:uuid:4509a320-00a0-008f-00b6-00559a327d31",
			JobUUID:          optional.New(uuid.Must(uuid.Parse("4509a320-00a0-008f-00b6-00559a327d31"))),
			Age:              optional.New(39 * time.Second),
			ImagesCompleted:  optional.New(1),
			ImagesToTransfer: optional.New(1),
			JobState:         JobCompleted,
			JobStateReasons:  []JobStateReason{JobCompletedSuccessfully},
		},
	},
}

// TestScannerStatus tests [ScannerStatus] conversion
// to and from the XML
func TestScannerStatus(t *testing.T) {
	type testData struct {
		status ScannerStatus
		xml    xmldoc.Element
	}

	tests := []testData{}

	for _, test := range tests {
		xml := test.status.ToXML()
		if !xml.Similar(test.xml) {
			t.Errorf("encode mismatch:\n"+
				"expected: %s\n"+
				"present:  %s\n",
				test.xml.EncodeString(nil),
				xml.EncodeString(nil))
		}

		status, err := DecodeScannerStatus(test.xml)
		if err != nil {
			t.Errorf("decode error:\n"+
				"input: %s\n"+
				"error:  %s\n",
				test.xml.EncodeString(nil), err)
			continue
		}

		if !reflect.DeepEqual(status, test.status) {
			t.Errorf("decode mismatch:\n"+
				"expected: %#v\n"+
				"present:  %#v\n",
				test.status, status)
		}
	}

}
