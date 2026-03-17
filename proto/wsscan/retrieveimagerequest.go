// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// RetrieveImageRequest: client request to retrieve scan data from the device

package wsscan

import (
	"fmt"
	"strconv"

	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// RetrieveImageRequest contains the client's request to retrieve scan data
// from the device after a scan job has been created. All three child elements
// are required.
type RetrieveImageRequest struct {
	DocumentDescription DocumentDescription
	JobID               int
	JobToken            string
}

// toXML generates XML tree for the [RetrieveImageRequest].
func (r RetrieveImageRequest) toXML(name string) xmldoc.Element {
	return xmldoc.Element{
		Name: name,
		Children: []xmldoc.Element{
			r.DocumentDescription.toXML(NsWSCN + ":DocumentDescription"),
			{Name: NsWSCN + ":JobId", Text: strconv.Itoa(r.JobID)},
			{Name: NsWSCN + ":JobToken", Text: r.JobToken},
		},
	}
}

// decodeRetrieveImageRequest decodes [RetrieveImageRequest] from the XML tree.
func decodeRetrieveImageRequest(root xmldoc.Element) (
	RetrieveImageRequest, error,
) {
	var r RetrieveImageRequest

	documentDescription := xmldoc.Lookup{
		Name:     NsWSCN + ":DocumentDescription",
		Required: true,
	}
	jobID := xmldoc.Lookup{
		Name:     NsWSCN + ":JobId",
		Required: true,
	}
	jobToken := xmldoc.Lookup{
		Name:     NsWSCN + ":JobToken",
		Required: true,
	}

	if missed := root.Lookup(
		&documentDescription, &jobID, &jobToken,
	); missed != nil {
		return r, xmldoc.XMLErrMissed(missed.Name)
	}

	var err error

	if r.DocumentDescription, err = decodeDocumentDescription(
		documentDescription.Elem); err != nil {
		return r, fmt.Errorf("DocumentDescription: %w", err)
	}

	if r.JobID, err = decodeNonNegativeInt(jobID.Elem); err != nil {
		return r, fmt.Errorf("JobId: %w", err)
	}
	if r.JobID < 1 {
		return r, fmt.Errorf("JobId: must be at least 1, got %d", r.JobID)
	}

	r.JobToken = jobToken.Elem.Text

	return r, nil
}
