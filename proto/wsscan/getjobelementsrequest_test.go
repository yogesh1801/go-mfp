// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// GetJobElementsRequest tests

package wsscan

import (
	"reflect"
	"testing"

	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// TestGetJobElementsRequest_RoundTrip verifies that a GetJobElementsRequest
// encodes to XML and decodes back to an identical value.
func TestGetJobElementsRequest_RoundTrip(t *testing.T) {
	orig := GetJobElementsRequest{
		JobID: 5,
		RequestedElements: []RequestedElement{
			RequestedElementDescription,
			RequestedElementStatus,
		},
	}
	elm := orig.toXML(NsWSCN + ":GetJobElementsRequest")
	if elm.Name != NsWSCN+":GetJobElementsRequest" {
		t.Errorf("expected element name %q, got %q",
			NsWSCN+":GetJobElementsRequest", elm.Name)
	}

	parsed, err := decodeGetJobElementsRequest(elm)
	if err != nil {
		t.Fatalf("decodeGetJobElementsRequest returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, parsed) {
		t.Errorf("expected %+v, got %+v", orig, parsed)
	}
}

// TestGetJobElementsRequest_MissingJobId verifies that decoding a request
// without JobId returns an error.
func TestGetJobElementsRequest_MissingJobId(t *testing.T) {
	elm := xmldoc.Element{
		Name: NsWSCN + ":GetJobElementsRequest",
		Children: []xmldoc.Element{
			{
				Name: NsWSCN + ":RequestedElements",
				Children: []xmldoc.Element{
					{Name: NsWSCN + ":Name", Text: NsWSCN + ":ScannerDescription"},
				},
			},
		},
	}
	_, err := decodeGetJobElementsRequest(elm)
	if err == nil {
		t.Error("expected error for missing JobId, got nil")
	}
}

// TestGetJobElementsRequest_MissingRequestedElements verifies that decoding
// a request without RequestedElements returns an error.
func TestGetJobElementsRequest_MissingRequestedElements(t *testing.T) {
	elm := xmldoc.Element{
		Name: NsWSCN + ":GetJobElementsRequest",
		Children: []xmldoc.Element{
			{Name: NsWSCN + ":JobId", Text: "1"},
		},
	}
	_, err := decodeGetJobElementsRequest(elm)
	if err == nil {
		t.Error("expected error for missing RequestedElements, got nil")
	}
}

// TestGetJobElementsRequest_EmptyRequestedElements verifies that an empty
// RequestedElements (no Name children) returns an error.
func TestGetJobElementsRequest_EmptyRequestedElements(t *testing.T) {
	elm := xmldoc.Element{
		Name: NsWSCN + ":GetJobElementsRequest",
		Children: []xmldoc.Element{
			{Name: NsWSCN + ":JobId", Text: "1"},
			{Name: NsWSCN + ":RequestedElements"},
		},
	}
	_, err := decodeGetJobElementsRequest(elm)
	if err == nil {
		t.Error("expected error for empty RequestedElements, got nil")
	}
}

// TestGetJobElementsRequest_ZeroJobId verifies that JobId=0 is rejected.
func TestGetJobElementsRequest_ZeroJobId(t *testing.T) {
	orig := GetJobElementsRequest{
		JobID:             0,
		RequestedElements: []RequestedElement{RequestedElementStatus},
	}
	elm := orig.toXML(NsWSCN + ":GetJobElementsRequest")
	_, err := decodeGetJobElementsRequest(elm)
	if err == nil {
		t.Error("expected error for JobId=0, got nil")
	}
}
