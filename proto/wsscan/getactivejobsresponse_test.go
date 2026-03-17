// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// GetActiveJobsResponse tests

package wsscan

import (
	"reflect"
	"testing"

	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// TestGetActiveJobsResponse_RoundTrip_Empty verifies that a response with an
// empty ActiveJobs (no current jobs) encodes to XML and decodes back
// identically.
func TestGetActiveJobsResponse_RoundTrip_Empty(t *testing.T) {
	orig := GetActiveJobsResponse{
		ActiveJobs: ActiveJobs{},
	}
	elm := orig.toXML(NsWSCN + ":GetActiveJobsResponse")
	if elm.Name != NsWSCN+":GetActiveJobsResponse" {
		t.Errorf("expected element name %q, got %q",
			NsWSCN+":GetActiveJobsResponse", elm.Name)
	}

	parsed, err := decodeGetActiveJobsResponse(elm)
	if err != nil {
		t.Fatalf("decodeGetActiveJobsResponse returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, parsed) {
		t.Errorf("expected %+v, got %+v", orig, parsed)
	}
}

// TestGetActiveJobsResponse_RoundTrip_WithJobs verifies that a response
// carrying active job summaries encodes to XML and decodes back identically.
func TestGetActiveJobsResponse_RoundTrip_WithJobs(t *testing.T) {
	orig := GetActiveJobsResponse{
		ActiveJobs: ActiveJobs{
			JobSummary: []JobSummary{
				{
					JobID:                  1,
					JobName:                "ScanJob1",
					JobOriginatingUserName: "user1",
					JobState:               JobStatePending,
					ScansCompleted:         0,
				},
				{
					JobID:                  2,
					JobName:                "ScanJob2",
					JobOriginatingUserName: "user2",
					JobState:               JobStateProcessing,
					ScansCompleted:         3,
				},
			},
		},
	}
	elm := orig.toXML(NsWSCN + ":GetActiveJobsResponse")

	parsed, err := decodeGetActiveJobsResponse(elm)
	if err != nil {
		t.Fatalf("decodeGetActiveJobsResponse returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, parsed) {
		t.Errorf("expected %+v, got %+v", orig, parsed)
	}
}

// TestGetActiveJobsResponse_MissingActiveJobs verifies that decoding a
// response without the required ActiveJobs child returns an error.
func TestGetActiveJobsResponse_MissingActiveJobs(t *testing.T) {
	elm := xmldoc.Element{Name: NsWSCN + ":GetActiveJobsResponse"}

	_, err := decodeGetActiveJobsResponse(elm)
	if err == nil {
		t.Error("expected error for missing ActiveJobs, got nil")
	}
}
