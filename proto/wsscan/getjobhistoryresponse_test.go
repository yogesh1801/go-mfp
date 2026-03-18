// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// GetJobHistoryResponse tests

package wsscan

import (
	"reflect"
	"testing"

	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// TestGetJobHistoryResponse_RoundTrip_Empty verifies that a response with no
// completed jobs encodes to XML and decodes back identically.
func TestGetJobHistoryResponse_RoundTrip_Empty(t *testing.T) {
	orig := GetJobHistoryResponse{JobHistory: nil}
	elm := orig.toXML(NsWSCN + ":GetJobHistoryResponse")
	if elm.Name != NsWSCN+":GetJobHistoryResponse" {
		t.Errorf("expected element name %q, got %q",
			NsWSCN+":GetJobHistoryResponse", elm.Name)
	}

	parsed, err := decodeGetJobHistoryResponse(elm)
	if err != nil {
		t.Fatalf("decodeGetJobHistoryResponse returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, parsed) {
		t.Errorf("expected %+v, got %+v", orig, parsed)
	}
}

// TestGetJobHistoryResponse_RoundTrip_WithJobs verifies that a response with
// completed job summaries encodes to XML and decodes back identically.
func TestGetJobHistoryResponse_RoundTrip_WithJobs(t *testing.T) {
	orig := GetJobHistoryResponse{
		JobHistory: []JobSummary{
			{
				JobID:                  1,
				JobName:                "CompletedJob1",
				JobOriginatingUserName: "user1",
				JobState:               JobStateCompleted,
				ScansCompleted:         3,
			},
			{
				JobID:                  2,
				JobName:                "AbortedJob",
				JobOriginatingUserName: "user2",
				JobState:               JobStateAborted,
				JobStateReasons:        []JobStateReason{JobTimedOut},
				ScansCompleted:         1,
			},
		},
	}

	elm := orig.toXML(NsWSCN + ":GetJobHistoryResponse")

	parsed, err := decodeGetJobHistoryResponse(elm)
	if err != nil {
		t.Fatalf("decodeGetJobHistoryResponse returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, parsed) {
		t.Errorf("expected %+v, got %+v", orig, parsed)
	}
}

// TestGetJobHistoryResponse_MissingJobHistory verifies that decoding a
// response without the required JobHistory child returns an error.
func TestGetJobHistoryResponse_MissingJobHistory(t *testing.T) {
	elm := xmldoc.Element{Name: NsWSCN + ":GetJobHistoryResponse"}

	_, err := decodeGetJobHistoryResponse(elm)
	if err == nil {
		t.Error("expected error for missing JobHistory, got nil")
	}
}
