// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// GetJobElementsRequest: requests information about the job identified by JobId

package wsscan

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// GetJobElementsRequest requests information related to the job identified
// by JobID. RequestedElements specifies which job schema elements to return.
type GetJobElementsRequest struct {
	JobID             int
	RequestedElements []JobRequestedElement
}

// toXML generates XML tree for the [GetJobElementsRequest].
func (r GetJobElementsRequest) toXML(name string) xmldoc.Element {
	nameElems := make([]xmldoc.Element, len(r.RequestedElements))
	for i, re := range r.RequestedElements {
		nameElems[i] = re.toXML(NsWSCN + ":Name")
	}

	return xmldoc.Element{
		Name: name,
		Children: []xmldoc.Element{
			{Name: NsWSCN + ":JobId", Text: strconv.Itoa(r.JobID)},
			{Name: NsWSCN + ":RequestedElements", Children: nameElems},
		},
	}
}

// decodeGetJobElementsRequest decodes [GetJobElementsRequest] from the XML
// tree.
func decodeGetJobElementsRequest(root xmldoc.Element) (
	GetJobElementsRequest, error,
) {
	var r GetJobElementsRequest

	jobID := xmldoc.Lookup{
		Name:     NsWSCN + ":JobId",
		Required: true,
	}
	requestedElements := xmldoc.Lookup{
		Name:     NsWSCN + ":RequestedElements",
		Required: true,
	}

	if missed := root.Lookup(&jobID, &requestedElements); missed != nil {
		return r, xmldoc.XMLErrMissed(missed.Name)
	}

	var err error
	if r.JobID, err = decodeNonNegativeInt(jobID.Elem); err != nil {
		return r, fmt.Errorf("JobId: %w", err)
	}
	if r.JobID < 1 {
		return r, fmt.Errorf("JobId: must be at least 1, got %d", r.JobID)
	}

	for _, child := range requestedElements.Elem.Children {
		if child.Name == NsWSCN+":Name" {
			re, err := decodeJobRequestedElement(child)
			if err != nil {
				return r, err
			}
			r.RequestedElements = append(r.RequestedElements, re)
		}
	}

	if len(r.RequestedElements) == 0 {
		return r, errors.New("RequestedElements: at least one Name is required")
	}

	return r, nil
}
