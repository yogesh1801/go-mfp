// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// CreateScanJobResponse tests

package wsscan

import (
	"reflect"
	"testing"

	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

func createValidDocumentFinalParameters() DocumentParameters {
	return DocumentParameters{
		Format: optional.New(ValWithOptions[FormatValue]{Val: PNG}),
	}
}

func createValidCreateScanJobResponse() CreateScanJobResponse {
	return CreateScanJobResponse{
		DocumentFinalParameters: createValidDocumentFinalParameters(),
		ImageInformation:        ImageInformation{},
		JobID:                   42,
		JobToken:                "token-abc-123",
	}
}

// TestCreateScanJobResponse_RoundTrip verifies that a fully populated
// CreateScanJobResponse encodes to XML and decodes back identically.
func TestCreateScanJobResponse_RoundTrip(t *testing.T) {
	orig := createValidCreateScanJobResponse()

	elm := orig.toXML(NsWSCN + ":CreateScanJobResponse")
	if elm.Name != NsWSCN+":CreateScanJobResponse" {
		t.Errorf("expected element name %q, got %q",
			NsWSCN+":CreateScanJobResponse", elm.Name)
	}

	parsed, err := decodeCreateScanJobResponse(elm)
	if err != nil {
		t.Fatalf("decodeCreateScanJobResponse returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, parsed) {
		t.Errorf("expected %+v, got %+v", orig, parsed)
	}
}

// TestCreateScanJobResponse_RoundTrip_WithImageInformation verifies that a
// response with ImageInformation containing side data encodes and decodes
// back identically.
func TestCreateScanJobResponse_RoundTrip_WithImageInformation(t *testing.T) {
	orig := CreateScanJobResponse{
		DocumentFinalParameters: createValidDocumentFinalParameters(),
		ImageInformation: ImageInformation{
			MediaFrontImageInfo: optional.New(createValidMediaSideImageInfo()),
		},
		JobID:    1,
		JobToken: "scan-job-token-xyz",
	}
	elm := orig.toXML(NsWSCN + ":CreateScanJobResponse")

	parsed, err := decodeCreateScanJobResponse(elm)
	if err != nil {
		t.Fatalf("decodeCreateScanJobResponse returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, parsed) {
		t.Errorf("expected %+v, got %+v", orig, parsed)
	}
}

// TestCreateScanJobResponse_MissingJobId verifies that decoding a response
// without the required JobId child returns an error.
func TestCreateScanJobResponse_MissingJobId(t *testing.T) {
	elm := xmldoc.Element{
		Name: NsWSCN + ":CreateScanJobResponse",
		Children: []xmldoc.Element{
			{Name: NsWSCN + ":JobToken", Text: "token"},
		},
	}
	_, err := decodeCreateScanJobResponse(elm)
	if err == nil {
		t.Error("expected error for missing JobId, got nil")
	}
}

// TestCreateScanJobResponse_MissingJobToken verifies that decoding a response
// without the required JobToken child returns an error.
func TestCreateScanJobResponse_MissingJobToken(t *testing.T) {
	elm := xmldoc.Element{
		Name: NsWSCN + ":CreateScanJobResponse",
		Children: []xmldoc.Element{
			{Name: NsWSCN + ":JobId", Text: "1"},
		},
	}
	_, err := decodeCreateScanJobResponse(elm)
	if err == nil {
		t.Error("expected error for missing JobToken, got nil")
	}
}

// TestCreateScanJobResponse_ZeroJobId verifies that a JobId of 0 (below the
// allowed minimum of 1) is rejected.
func TestCreateScanJobResponse_ZeroJobId(t *testing.T) {
	orig := createValidCreateScanJobResponse()
	orig.JobID = 0
	elm := orig.toXML(NsWSCN + ":CreateScanJobResponse")

	_, err := decodeCreateScanJobResponse(elm)
	if err == nil {
		t.Error("expected error for JobId=0, got nil")
	}
}
