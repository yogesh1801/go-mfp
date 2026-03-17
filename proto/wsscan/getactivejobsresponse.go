// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// GetActiveJobsResponse: returns a summary of all currently active scan jobs

package wsscan

import (
	"fmt"

	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// GetActiveJobsResponse returns a summary of job-related variables for all
// currently active scan jobs.
type GetActiveJobsResponse struct {
	ActiveJobs ActiveJobs
}

// toXML generates XML tree for the [GetActiveJobsResponse].
func (r GetActiveJobsResponse) toXML(name string) xmldoc.Element {
	return xmldoc.Element{
		Name:     name,
		Children: []xmldoc.Element{r.ActiveJobs.toXML(NsWSCN + ":ActiveJobs")},
	}
}

// decodeGetActiveJobsResponse decodes [GetActiveJobsResponse] from the XML
// tree.
func decodeGetActiveJobsResponse(root xmldoc.Element) (
	GetActiveJobsResponse, error,
) {
	var r GetActiveJobsResponse

	activeJobs := xmldoc.Lookup{
		Name:     NsWSCN + ":ActiveJobs",
		Required: true,
	}

	if missed := root.Lookup(&activeJobs); missed != nil {
		return r, xmldoc.XMLErrMissed(missed.Name)
	}

	var err error
	if r.ActiveJobs, err = decodeActiveJobs(activeJobs.Elem); err != nil {
		return r, fmt.Errorf("ActiveJobs: %w", err)
	}

	return r, nil
}
