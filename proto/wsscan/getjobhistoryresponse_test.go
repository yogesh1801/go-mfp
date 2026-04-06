// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// GetJobHistoryResponse tests

package wsscan

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/uuid"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// TestGetJobHistoryResponse_Action verifies that Action returns
// ActGetJobHistoryResponse.
func TestGetJobHistoryResponse_Action(t *testing.T) {
	r := GetJobHistoryResponse{}
	if r.Action() != ActGetJobHistoryResponse {
		t.Errorf("expected ActGetJobHistoryResponse, got %v", r.Action())
	}
}

// TestGetJobHistoryResponse_ToXML verifies that ToXML produces the correct
// root element name.
func TestGetJobHistoryResponse_ToXML(t *testing.T) {
	r := GetJobHistoryResponse{}
	elm := r.ToXML()
	if elm.Name != NsWSCN+":GetJobHistoryResponse" {
		t.Errorf("expected element name %q, got %q",
			NsWSCN+":GetJobHistoryResponse", elm.Name)
	}
}

// TestGetJobHistoryResponse_MessageRoundTrip verifies that a
// GetJobHistoryResponse survives a full Message encode/decode cycle.
func TestGetJobHistoryResponse_MessageRoundTrip(t *testing.T) {
	body := GetJobHistoryResponse{JobHistory: nil}
	msg := Message{
		Header: Header{
			Action:    body.Action(),
			MessageID: AnyURI(uuid.Random().URN()),
			To:        optional.New(AnyURI(AddrAnonymous)),
		},
		Body: body,
	}

	data := msg.Encode()

	root, err := xmldoc.Decode(NsMap, bytes.NewReader(data))
	if err != nil {
		t.Fatalf("xmldoc.Decode returned error: %v", err)
	}

	decoded, err := DecodeMessage(root)
	if err != nil {
		t.Fatalf("DecodeMessage returned error: %v", err)
	}

	if _, ok := decoded.Body.(GetJobHistoryResponse); !ok {
		t.Errorf("expected body type GetJobHistoryResponse, got %T", decoded.Body)
	}
}

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
