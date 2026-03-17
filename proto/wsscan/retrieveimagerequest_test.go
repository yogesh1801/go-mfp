// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// RetrieveImageRequest tests

package wsscan

import (
	"reflect"
	"testing"

	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

func createValidRetrieveImageRequest() RetrieveImageRequest {
	return RetrieveImageRequest{
		DocumentDescription: DocumentDescription{
			DocumentName: "scan001.pdf",
		},
		JobID:    1,
		JobToken: "token-xyz",
	}
}

// TestRetrieveImageRequest_RoundTrip verifies that a RetrieveImageRequest
// encodes to XML and decodes back to an identical value.
func TestRetrieveImageRequest_RoundTrip(t *testing.T) {
	orig := createValidRetrieveImageRequest()

	elm := orig.toXML(NsWSCN + ":RetrieveImageRequest")
	if elm.Name != NsWSCN+":RetrieveImageRequest" {
		t.Errorf("expected element name %q, got %q",
			NsWSCN+":RetrieveImageRequest", elm.Name)
	}

	parsed, err := decodeRetrieveImageRequest(elm)
	if err != nil {
		t.Fatalf("decodeRetrieveImageRequest returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, parsed) {
		t.Errorf("expected %+v, got %+v", orig, parsed)
	}
}

// TestRetrieveImageRequest_MissingDocumentDescription verifies that decoding
// a request without DocumentDescription returns an error.
func TestRetrieveImageRequest_MissingDocumentDescription(t *testing.T) {
	elm := xmldoc.Element{
		Name: NsWSCN + ":RetrieveImageRequest",
		Children: []xmldoc.Element{
			{Name: NsWSCN + ":JobId", Text: "1"},
			{Name: NsWSCN + ":JobToken", Text: "token"},
		},
	}
	_, err := decodeRetrieveImageRequest(elm)
	if err == nil {
		t.Error("expected error for missing DocumentDescription, got nil")
	}
}

// TestRetrieveImageRequest_MissingJobId verifies that decoding a request
// without JobId returns an error.
func TestRetrieveImageRequest_MissingJobId(t *testing.T) {
	orig := createValidRetrieveImageRequest()
	elm := orig.toXML(NsWSCN + ":RetrieveImageRequest")
	var filtered []xmldoc.Element
	for _, child := range elm.Children {
		if child.Name != NsWSCN+":JobId" {
			filtered = append(filtered, child)
		}
	}
	elm.Children = filtered

	_, err := decodeRetrieveImageRequest(elm)
	if err == nil {
		t.Error("expected error for missing JobId, got nil")
	}
}

// TestRetrieveImageRequest_MissingJobToken verifies that decoding a request
// without JobToken returns an error.
func TestRetrieveImageRequest_MissingJobToken(t *testing.T) {
	orig := createValidRetrieveImageRequest()
	elm := orig.toXML(NsWSCN + ":RetrieveImageRequest")
	var filtered []xmldoc.Element
	for _, child := range elm.Children {
		if child.Name != NsWSCN+":JobToken" {
			filtered = append(filtered, child)
		}
	}
	elm.Children = filtered

	_, err := decodeRetrieveImageRequest(elm)
	if err == nil {
		t.Error("expected error for missing JobToken, got nil")
	}
}

// TestRetrieveImageRequest_ZeroJobId verifies that a JobId of 0 is rejected.
func TestRetrieveImageRequest_ZeroJobId(t *testing.T) {
	orig := createValidRetrieveImageRequest()
	orig.JobID = 0
	elm := orig.toXML(NsWSCN + ":RetrieveImageRequest")

	_, err := decodeRetrieveImageRequest(elm)
	if err == nil {
		t.Error("expected error for JobId=0, got nil")
	}
}
