// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// GetScannerElementsResponse tests

package wsscan

import (
	"reflect"
	"testing"

	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// TestGetScannerElementsResponse_RoundTrip_Single verifies that a response
// with a single ElementData entry encodes to XML and decodes back identically.
func TestGetScannerElementsResponse_RoundTrip_Single(t *testing.T) {
	orig := GetScannerElementsResponse{
		ScannerElements: []ElementData{
			{
				Name:               ElementDataScannerDescription,
				Valid:              BooleanElement("true"),
				ScannerDescription: optional.New(createValidScannerDescription()),
			},
		},
	}

	elm := orig.toXML(NsWSCN + ":GetScannerElementsResponse")
	if elm.Name != NsWSCN+":GetScannerElementsResponse" {
		t.Errorf("expected element name %q, got %q",
			NsWSCN+":GetScannerElementsResponse", elm.Name)
	}

	parsed, err := decodeGetScannerElementsResponse(elm)
	if err != nil {
		t.Fatalf("decodeGetScannerElementsResponse returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, parsed) {
		t.Errorf("expected %+v, got %+v", orig, parsed)
	}
}

// TestGetScannerElementsResponse_RoundTrip_Multiple verifies that a response
// carrying multiple ElementData entries (description + status) round-trips
// correctly and preserves order.
func TestGetScannerElementsResponse_RoundTrip_Multiple(t *testing.T) {
	orig := GetScannerElementsResponse{
		ScannerElements: []ElementData{
			{
				Name:               ElementDataScannerDescription,
				Valid:              BooleanElement("true"),
				ScannerDescription: optional.New(createValidScannerDescription()),
			},
			{
				Name:          ElementDataScannerStatus,
				Valid:         BooleanElement("true"),
				ScannerStatus: optional.New(createValidScannerStatus()),
			},
		},
	}

	elm := orig.toXML(NsWSCN + ":GetScannerElementsResponse")

	parsed, err := decodeGetScannerElementsResponse(elm)
	if err != nil {
		t.Fatalf("decodeGetScannerElementsResponse returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, parsed) {
		t.Errorf("expected %+v, got %+v", orig, parsed)
	}
}

// TestGetScannerElementsResponse_MissingScannerElements verifies that
// decoding a response without the required ScannerElements child returns
// an error.
func TestGetScannerElementsResponse_MissingScannerElements(t *testing.T) {
	elm := xmldoc.Element{Name: NsWSCN + ":GetScannerElementsResponse"}

	_, err := decodeGetScannerElementsResponse(elm)
	if err == nil {
		t.Error("expected error for missing ScannerElements, got nil")
	}
}

// TestGetScannerElementsResponse_EmptyScannerElements verifies that
// decoding a response with an empty ScannerElements (no ElementData children)
// returns an error.
func TestGetScannerElementsResponse_EmptyScannerElements(t *testing.T) {
	elm := xmldoc.Element{
		Name: NsWSCN + ":GetScannerElementsResponse",
		Children: []xmldoc.Element{
			{Name: NsWSCN + ":ScannerElements"},
		},
	}

	_, err := decodeGetScannerElementsResponse(elm)
	if err == nil {
		t.Error("expected error for empty ScannerElements, got nil")
	}
}
