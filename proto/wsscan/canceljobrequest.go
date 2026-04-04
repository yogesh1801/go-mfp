// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// CancelJobRequest: client request to cancel a scan job

package wsscan

import (
	"fmt"
	"strconv"

	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// CancelJobRequest enables a client to cancel the scan job identified by
// JobID.
type CancelJobRequest struct {
	JobID int
}

// Action returns the [Action] associated with this body.
func (CancelJobRequest) Action() Action { return ActCancelJob }

// ToXML encodes the body into an XML tree.
func (r CancelJobRequest) ToXML() xmldoc.Element {
	return r.toXML(NsWSCN + ":CancelJobRequest")
}

// toXML generates XML tree for the [CancelJobRequest].
func (r CancelJobRequest) toXML(name string) xmldoc.Element {
	return xmldoc.Element{
		Name: name,
		Children: []xmldoc.Element{
			{Name: NsWSCN + ":JobId", Text: strconv.Itoa(r.JobID)},
		},
	}
}

// decodeCancelJobRequest decodes [CancelJobRequest] from the XML tree.
func decodeCancelJobRequest(root xmldoc.Element) (CancelJobRequest, error) {
	var r CancelJobRequest

	jobID := xmldoc.Lookup{
		Name:     NsWSCN + ":JobId",
		Required: true,
	}

	if missed := root.Lookup(&jobID); missed != nil {
		return r, xmldoc.XMLErrMissed(missed.Name)
	}

	var err error
	if r.JobID, err = decodeNonNegativeInt(jobID.Elem); err != nil {
		return r, fmt.Errorf("JobId: %w", err)
	}
	if r.JobID < 1 {
		return r, fmt.Errorf("JobId: must be at least 1, got %d", r.JobID)
	}

	return r, nil
}
