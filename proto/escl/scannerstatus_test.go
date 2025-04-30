// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scanner status test

package escl

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/alexpevzner/mfp/util/optional"
	"github.com/alexpevzner/mfp/util/xmldoc"
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
			JobUUID:          optional.New("4509a320-00a0-008f-00b6-00559a327d32"),
			Age:              optional.New(20 * time.Second),
			ImagesCompleted:  optional.New(2),
			ImagesToTransfer: optional.New(1),
			JobState:         JobProcessing,
			JobStateReasons:  []JobStateReason{JobScanningAndTransferring},
		},
		{
			JobURI:           "/eSCL/ScanJobs/urn:uuid:4509a320-00a0-008f-00b6-00559a327d31",
			JobUUID:          optional.New("4509a320-00a0-008f-00b6-00559a327d31"),
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

	tests := []testData{
		{
			// Full data test
			status: testScannerStatus,
			xml: xmldoc.WithChildren(
				NsScan+":ScannerStatus",
				xmldoc.WithText(NsPWG+":Version", "2.0"),
				ScannerProcessing.toXML(NsPWG+":State"),
				ScannerAdfProcessing.toXML(NsScan+":AdfState"),
				xmldoc.WithChildren(
					NsScan+":Jobs",
					func() (jobs []xmldoc.Element) {
						for _, job := range testScannerStatus.Jobs {
							jobs = append(jobs,
								job.toXML(NsScan+":JobInfo"))
						}
						return
					}()...,
				),
			),
		},
	}

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

// TestScannerStatusDecodeErrors tests [ScannerStatus] XML decode
// errors handling
func TestScannerStatusDecodeErrors(t *testing.T) {
	type testData struct {
		xml xmldoc.Element
		err string
	}

	tests := []testData{
		// Missed ScannerStatus.Version
		{
			xml: xmldoc.WithChildren(
				NsScan + ":ScannerStatus",
			),
			err: `/scan:ScannerStatus/pwg:Version: missed`,
		},

		// Missed ScannerStatus.State
		{
			xml: xmldoc.WithChildren(
				NsScan+":ScannerStatus",
				xmldoc.WithText(NsPWG+":Version", "2.0"),
			),
			err: `/scan:ScannerStatus/pwg:State: missed`,
		},

		// Bad ScannerStatus.Version
		{
			xml: xmldoc.WithChildren(
				NsScan+":ScannerStatus",
				xmldoc.WithText(NsPWG+":Version", "Bad"),
				ScannerIdle.toXML(NsPWG+":State"),
			),
			err: `/scan:ScannerStatus/pwg:Version: "Bad": invalid eSCL version`,
		},

		// Bad ScannerStatus.State
		{
			xml: xmldoc.WithChildren(
				NsScan+":ScannerStatus",
				xmldoc.WithText(NsPWG+":Version", "2.0"),
				xmldoc.WithText(NsPWG+":State", "Bad"),
			),
			err: `/scan:ScannerStatus/pwg:State: invalid ScannerState: "Bad"`,
		},

		// Bad ScannerStatus.ADFState
		{
			xml: xmldoc.WithChildren(
				NsScan+":ScannerStatus",
				xmldoc.WithText(NsPWG+":Version", "2.0"),
				ScannerIdle.toXML(NsPWG+":State"),
				xmldoc.WithText(NsScan+":AdfState", "Bad"),
			),
			err: `/scan:ScannerStatus/scan:AdfState: invalid ADFState: "Bad"`,
		},

		// Bad ScannerStatus.Jobs
		{
			xml: xmldoc.WithChildren(
				NsScan+":ScannerStatus",
				xmldoc.WithText(NsPWG+":Version", "2.0"),
				ScannerIdle.toXML(NsPWG+":State"),
				xmldoc.WithChildren(
					NsScan+":Jobs",
					xmldoc.WithText(NsScan+":JobInfo", "Bad"),
				),
			),
			err: `/scan:ScannerStatus/scan:Jobs/scan:JobInfo/pwg:JobUri: missed`,
		},
	}

	for _, test := range tests {
		_, err := DecodeScannerStatus(test.xml)
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
