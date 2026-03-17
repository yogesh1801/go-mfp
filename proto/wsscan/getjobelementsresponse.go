// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// GetJobElementsResponse: returns job-related information requested by client

package wsscan

import (
	"errors"

	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// GetJobElementsResponse returns the job-related information that a client
// requested via GetJobElementsRequest. JobElements holds one [ElementData]
// entry for each job schema element that was requested.
type GetJobElementsResponse struct {
	JobElements []ElementData
}

// toXML generates XML tree for the [GetJobElementsResponse].
func (r GetJobElementsResponse) toXML(name string) xmldoc.Element {
	children := make([]xmldoc.Element, len(r.JobElements))
	for i, ed := range r.JobElements {
		children[i] = ed.toXML(NsWSCN + ":ElementData")
	}

	return xmldoc.Element{
		Name: name,
		Children: []xmldoc.Element{
			{Name: NsWSCN + ":JobElements", Children: children},
		},
	}
}

// decodeGetJobElementsResponse decodes [GetJobElementsResponse] from the XML
// tree.
func decodeGetJobElementsResponse(root xmldoc.Element) (
	GetJobElementsResponse, error,
) {
	var r GetJobElementsResponse

	jobElements := xmldoc.Lookup{
		Name:     NsWSCN + ":JobElements",
		Required: true,
	}

	if missed := root.Lookup(&jobElements); missed != nil {
		return r, xmldoc.XMLErrMissed(missed.Name)
	}

	for _, child := range jobElements.Elem.Children {
		if child.Name == NsWSCN+":ElementData" {
			ed, err := decodeElementData(child)
			if err != nil {
				return r, err
			}

			r.JobElements = append(r.JobElements, ed)
		}
	}

	if len(r.JobElements) == 0 {
		return r, errors.New(
			"JobElements: at least one ElementData is required")
	}

	return r, nil
}
