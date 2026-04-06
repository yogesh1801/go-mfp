// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// CancelJobResponse tests

package wsscan

import (
	"bytes"
	"testing"

	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/uuid"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// TestCancelJobResponse_Action verifies that Action returns ActCancelJobResponse.
func TestCancelJobResponse_Action(t *testing.T) {
	r := CancelJobResponse{}
	if r.Action() != ActCancelJobResponse {
		t.Errorf("expected ActCancelJobResponse, got %v", r.Action())
	}
}

// TestCancelJobResponse_ToXML verifies that ToXML produces the correct root
// element name with no children.
func TestCancelJobResponse_ToXML(t *testing.T) {
	r := CancelJobResponse{}
	elm := r.ToXML()
	if elm.Name != NsWSCN+":CancelJobResponse" {
		t.Errorf("expected element name %q, got %q",
			NsWSCN+":CancelJobResponse", elm.Name)
	}
	if len(elm.Children) != 0 {
		t.Errorf("expected no children, got %d", len(elm.Children))
	}
}

// TestCancelJobResponse_MessageRoundTrip verifies that a CancelJobResponse
// survives a full Message encode/decode cycle.
func TestCancelJobResponse_MessageRoundTrip(t *testing.T) {
	body := CancelJobResponse{}
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

	if _, ok := decoded.Body.(CancelJobResponse); !ok {
		t.Errorf("expected body type CancelJobResponse, got %T", decoded.Body)
	}
}
