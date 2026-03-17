// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// GetJobElementsResponse tests

package wsscan

import (
	"reflect"
	"testing"

	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// TestGetJobElementsResponse_RoundTrip verifies that a GetJobElementsResponse
// encodes to XML and decodes back to an identical value.
func TestGetJobElementsResponse_RoundTrip(t *testing.T) {
	orig := GetJobElementsResponse{
		JobElements: []ElementData{
			{
				Name:               ElementDataScannerDescription,
				Valid:              BooleanElement("true"),
				ScannerDescription: optional.New(createValidScannerDescription()),
			},
		},
	}
	elm := orig.toXML(NsWSCN + ":GetJobElementsResponse")
	if elm.Name != NsWSCN+":GetJobElementsResponse" {
		t.Errorf("expected element name %q, got %q",
			NsWSCN+":GetJobElementsResponse", elm.Name)
	}

	parsed, err := decodeGetJobElementsResponse(elm)
	if err != nil {
		t.Fatalf("decodeGetJobElementsResponse returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, parsed) {
		t.Errorf("expected %+v, got %+v", orig, parsed)
	}
}

// TestGetJobElementsResponse_MissingJobElements verifies that decoding a
// response without the required JobElements child returns an error.
func TestGetJobElementsResponse_MissingJobElements(t *testing.T) {
	elm := xmldoc.Element{Name: NsWSCN + ":GetJobElementsResponse"}
	_, err := decodeGetJobElementsResponse(elm)
	if err == nil {
		t.Error("expected error for missing JobElements, got nil")
	}
}

// TestGetJobElementsResponse_EmptyJobElements verifies that a response with
// an empty JobElements (no ElementData children) returns an error.
func TestGetJobElementsResponse_EmptyJobElements(t *testing.T) {
	elm := xmldoc.Element{
		Name: NsWSCN + ":GetJobElementsResponse",
		Children: []xmldoc.Element{
			{Name: NsWSCN + ":JobElements"},
		},
	}
	_, err := decodeGetJobElementsResponse(elm)
	if err == nil {
		t.Error("expected error for empty JobElements, got nil")
	}
}
