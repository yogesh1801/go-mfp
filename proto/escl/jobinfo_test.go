// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// JobInfo tests

package escl

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/alexpevzner/mfp/util/optional"
	"github.com/alexpevzner/mfp/util/xmldoc"
)

// testJobInfo contains example of the initialized
// JobInfo structure
var testJobInfo = JobInfo{
	JobURI:             "/eSCL/ScanJobs/urn:uuid:4509a320-00a0-008f-00b6-00559a327d32",
	JobUUID:            optional.New("urn:uuid:4509a320-00a0-008f-00b6-00559a327d32"),
	Age:                optional.New(20 * time.Second),
	ImagesCompleted:    optional.New(2),
	ImagesToTransfer:   optional.New(1),
	TransferRetryCount: optional.New(0),
	JobState:           JobCompleted,
	JobStateReasons:    []JobStateReason{JobCompletedSuccessfully},
}

// TestJobInfo tests [JobInfo] conversion to and from the XML
func TestJobInfo(t *testing.T) {
	type testData struct {
		info JobInfo
		xml  xmldoc.Element
	}

	tests := []testData{
		{
			// Full data test
			info: testJobInfo,
			xml: xmldoc.WithChildren(NsScan+":JobInfo",
				xmldoc.WithText(NsPWG+":JobUri",
					"/eSCL/ScanJobs/urn:uuid:4509a320-00a0-008f-00b6-00559a327d32"),
				xmldoc.WithText(NsPWG+":JobUuid",
					"urn:uuid:4509a320-00a0-008f-00b6-00559a327d32"),
				xmldoc.WithText(NsScan+":Age", "20"),
				xmldoc.WithText(NsPWG+":ImagesCompleted", "2"),
				xmldoc.WithText(NsPWG+":ImagesToTransfer", "1"),
				xmldoc.WithText(NsScan+":TransferRetryCount", "0"),
				JobCompleted.toXML(NsPWG+":JobState"),
				xmldoc.WithChildren(NsPWG+":JobStateReasons",
					JobCompletedSuccessfully.toXML(
						NsPWG+":JobStateReason")),
			),
		},

		{
			// Missed optional elements
			info: JobInfo{
				JobURI:   "/eSCL/ScanJobs/1",
				JobState: JobProcessing,
			},
			xml: xmldoc.WithChildren(NsScan+":JobInfo",
				xmldoc.WithText(NsPWG+":JobUri",
					"/eSCL/ScanJobs/1"),
				JobProcessing.toXML(NsPWG+":JobState"),
			),
		},
	}

	for _, test := range tests {
		xml := test.info.toXML(NsScan + ":JobInfo")
		if !xml.Similar(test.xml) {
			t.Errorf("encode mismatch:\n"+
				"expected: %s\n"+
				"present:  %s\n",
				test.xml.EncodeString(nil),
				xml.EncodeString(nil))
		}

		info, err := decodeJobInfo(test.xml)
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

// TestJobInfoDecodeErrors tests [JobInfo] XML decode
// errors handling
func TestJobInfoDecodeErrors(t *testing.T) {
	type testData struct {
		xml xmldoc.Element // Input XML
		err string         // Expected error
	}

	tests := []testData{
		// Missed required element (JobURI)
		{
			xml: xmldoc.WithChildren(NsScan + ":JobInfo"),
			err: `/scan:JobInfo/pwg:JobUri: missed`,
		},

		// Missed required element (JobState)
		{
			xml: xmldoc.WithChildren(NsScan+":JobInfo",
				xmldoc.WithText(NsPWG+":JobUri",
					"/eSCL/ScanJobs/1"),
			),
			err: `/scan:JobInfo/pwg:JobState: missed`,
		},

		// Error in Age
		{
			xml: xmldoc.WithChildren(NsScan+":JobInfo",
				xmldoc.WithText(NsPWG+":JobUri",
					"/eSCL/ScanJobs/1"),
				JobProcessing.toXML(NsPWG+":JobState"),
				xmldoc.WithText(NsScan+":Age", "bad"),
			),
			err: `/scan:JobInfo/scan:Age: invalid int: "bad"`,
		},

		// Error in ImagesCompleted
		{
			xml: xmldoc.WithChildren(NsScan+":JobInfo",
				xmldoc.WithText(NsPWG+":JobUri",
					"/eSCL/ScanJobs/1"),
				JobProcessing.toXML(NsPWG+":JobState"),
				xmldoc.WithText(NsPWG+":ImagesCompleted", "bad"),
			),
			err: `/scan:JobInfo/pwg:ImagesCompleted: invalid int: "bad"`,
		},

		// Error in ImagesToTransfer
		{
			xml: xmldoc.WithChildren(NsScan+":JobInfo",
				xmldoc.WithText(NsPWG+":JobUri",
					"/eSCL/ScanJobs/1"),
				JobProcessing.toXML(NsPWG+":JobState"),
				xmldoc.WithText(NsPWG+":ImagesToTransfer", "bad"),
			),
			err: `/scan:JobInfo/pwg:ImagesToTransfer: invalid int: "bad"`,
		},

		// Error in TransferRetryCount
		{
			xml: xmldoc.WithChildren(NsScan+":JobInfo",
				xmldoc.WithText(NsPWG+":JobUri",
					"/eSCL/ScanJobs/1"),
				JobProcessing.toXML(NsPWG+":JobState"),
				xmldoc.WithText(NsScan+":TransferRetryCount", "bad"),
			),
			err: `/scan:JobInfo/scan:TransferRetryCount: invalid int: "bad"`,
		},

		// Error in JobState
		{
			xml: xmldoc.WithChildren(NsScan+":JobInfo",
				xmldoc.WithText(NsPWG+":JobUri",
					"/eSCL/ScanJobs/1"),
				xmldoc.WithText(NsPWG+":JobState", "bad"),
			),
			err: `/scan:JobInfo/pwg:JobState: invalid JobState: "bad"`,
		},

		// Error in JobStateReasons
		{
			xml: xmldoc.WithChildren(NsScan+":JobInfo",
				xmldoc.WithText(NsPWG+":JobUri",
					"/eSCL/ScanJobs/1"),
				JobProcessing.toXML(NsPWG+":JobState"),
				xmldoc.WithChildren(NsPWG+":JobStateReasons",
					xmldoc.WithText(NsPWG+":JobStateReason", "")),
			),
			err: `/scan:JobInfo/pwg:JobStateReasons/pwg:JobStateReason: invalid JobStateReason: ""`,
		},
	}

	for _, test := range tests {
		_, err := decodeJobInfo(test.xml)
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
