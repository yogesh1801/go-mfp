// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// CancelJobRequest tests

package wsscan

import (
	"reflect"
	"testing"

	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// TestCancelJobRequest_RoundTrip verifies that a CancelJobRequest encodes to
// XML and decodes back to an identical value.
func TestCancelJobRequest_RoundTrip(t *testing.T) {
	orig := CancelJobRequest{JobID: 7}

	elm := orig.toXML(NsWSCN + ":CancelJobRequest")
	if elm.Name != NsWSCN+":CancelJobRequest" {
		t.Errorf("expected element name %q, got %q",
			NsWSCN+":CancelJobRequest", elm.Name)
	}

	parsed, err := decodeCancelJobRequest(elm)
	if err != nil {
		t.Fatalf("decodeCancelJobRequest returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, parsed) {
		t.Errorf("expected %+v, got %+v", orig, parsed)
	}
}

// TestCancelJobRequest_MissingJobId verifies that decoding a request without
// JobId returns an error.
func TestCancelJobRequest_MissingJobId(t *testing.T) {
	elm := xmldoc.Element{Name: NsWSCN + ":CancelJobRequest"}
	_, err := decodeCancelJobRequest(elm)
	if err == nil {
		t.Error("expected error for missing JobId, got nil")
	}
}

// TestCancelJobRequest_ZeroJobId verifies that a JobId of 0 is rejected.
func TestCancelJobRequest_ZeroJobId(t *testing.T) {
	orig := CancelJobRequest{JobID: 0}
	elm := orig.toXML(NsWSCN + ":CancelJobRequest")
	_, err := decodeCancelJobRequest(elm)
	if err == nil {
		t.Error("expected error for JobId=0, got nil")
	}
}
