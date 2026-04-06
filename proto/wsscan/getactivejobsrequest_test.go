// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// GetActiveJobsRequest tests

package wsscan

import (
	"bytes"
	"testing"

	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/uuid"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// TestGetActiveJobsRequest_Action verifies that Action returns ActGetActiveJobs.
func TestGetActiveJobsRequest_Action(t *testing.T) {
	r := GetActiveJobsRequest{}
	if r.Action() != ActGetActiveJobs {
		t.Errorf("expected ActGetActiveJobs, got %v", r.Action())
	}
}

// TestGetActiveJobsRequest_ToXML verifies that ToXML produces the correct root
// element name with no children.
func TestGetActiveJobsRequest_ToXML(t *testing.T) {
	r := GetActiveJobsRequest{}
	elm := r.ToXML()
	if elm.Name != NsWSCN+":GetActiveJobsRequest" {
		t.Errorf("expected element name %q, got %q",
			NsWSCN+":GetActiveJobsRequest", elm.Name)
	}
	if len(elm.Children) != 0 {
		t.Errorf("expected no children, got %d", len(elm.Children))
	}
}

// TestGetActiveJobsRequest_MessageRoundTrip verifies that a GetActiveJobsRequest
// survives a full Message encode/decode cycle.
func TestGetActiveJobsRequest_MessageRoundTrip(t *testing.T) {
	body := GetActiveJobsRequest{}
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

	if _, ok := decoded.Body.(GetActiveJobsRequest); !ok {
		t.Errorf("expected body type GetActiveJobsRequest, got %T", decoded.Body)
	}
}
