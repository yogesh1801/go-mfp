// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// GetJobHistoryRequest tests

package wsscan

import (
	"bytes"
	"testing"

	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/uuid"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// TestGetJobHistoryRequest_Action verifies that Action returns ActGetJobHistory.
func TestGetJobHistoryRequest_Action(t *testing.T) {
	r := GetJobHistoryRequest{}
	if r.Action() != ActGetJobHistory {
		t.Errorf("expected ActGetJobHistory, got %v", r.Action())
	}
}

// TestGetJobHistoryRequest_ToXML verifies that ToXML produces the correct root
// element name with no children.
func TestGetJobHistoryRequest_ToXML(t *testing.T) {
	r := GetJobHistoryRequest{}
	elm := r.ToXML()
	if elm.Name != NsWSCN+":GetJobHistoryRequest" {
		t.Errorf("expected element name %q, got %q",
			NsWSCN+":GetJobHistoryRequest", elm.Name)
	}
	if len(elm.Children) != 0 {
		t.Errorf("expected no children, got %d", len(elm.Children))
	}
}

// TestGetJobHistoryRequest_MessageRoundTrip verifies that a GetJobHistoryRequest
// survives a full Message encode/decode cycle.
func TestGetJobHistoryRequest_MessageRoundTrip(t *testing.T) {
	body := GetJobHistoryRequest{}
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

	if _, ok := decoded.Body.(GetJobHistoryRequest); !ok {
		t.Errorf("expected body type GetJobHistoryRequest, got %T", decoded.Body)
	}
}
