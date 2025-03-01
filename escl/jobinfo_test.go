// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// JobInfo tests

package escl

import (
	"reflect"
	"testing"
	"time"

	"github.com/alexpevzner/mfp/optional"
	"github.com/alexpevzner/mfp/uuid"
	"github.com/alexpevzner/mfp/xmldoc"
)

// testtestJobInfo=JobInfo contains example of the initialized
// testJobInfo=JobInfo structure
var testJobInfo = JobInfo{
	JobURI:           "/eSCL/ScanJobs/urn:uuid:4509a320-00a0-008f-00b6-00559a327d32",
	JobUUID:          optional.New(uuid.Must(uuid.Parse("4509a320-00a0-008f-00b6-00559a327d32"))),
	Age:              optional.New(20 * time.Second),
	ImagesCompleted:  optional.New(2),
	ImagesToTransfer: optional.New(1),
	JobState:         JobCompleted,
	JobStateReasons:  []JobStateReason{JobCompletedSuccessfully},
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
