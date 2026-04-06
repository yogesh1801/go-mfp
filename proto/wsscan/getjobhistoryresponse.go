// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// GetJobHistoryResponse: returns a summary of completed scan jobs

package wsscan

import (
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// GetJobHistoryResponse returns a summary of completed scan jobs.
// JobHistory holds one [JobSummary] entry per completed job.
type GetJobHistoryResponse struct {
	JobHistory []JobSummary
}

// Action returns the [Action] associated with this body.
func (*GetJobHistoryResponse) Action() Action { return ActGetJobHistoryResponse }

// ToXML encodes the body into an XML tree.
func (r *GetJobHistoryResponse) ToXML() xmldoc.Element {
	return r.toXML(NsWSCN + ":GetJobHistoryResponse")
}

// toXML generates XML tree for the [GetJobHistoryResponse].
func (r GetJobHistoryResponse) toXML(name string) xmldoc.Element {
	children := make([]xmldoc.Element, len(r.JobHistory))
	for i, js := range r.JobHistory {
		children[i] = js.toXML(NsWSCN + ":JobSummary")
	}

	return xmldoc.Element{
		Name: name,
		Children: []xmldoc.Element{
			{Name: NsWSCN + ":JobHistory", Children: children},
		},
	}
}

// decodeGetJobHistoryResponse decodes [GetJobHistoryResponse] from the XML
// tree.
func decodeGetJobHistoryResponse(root xmldoc.Element) (
	GetJobHistoryResponse, error,
) {
	var r GetJobHistoryResponse

	jobHistory := xmldoc.Lookup{
		Name:     NsWSCN + ":JobHistory",
		Required: true,
	}

	if missed := root.Lookup(&jobHistory); missed != nil {
		return r, xmldoc.XMLErrMissed(missed.Name)
	}

	for _, child := range jobHistory.Elem.Children {
		if child.Name == NsWSCN+":JobSummary" {
			js, err := decodeJobSummary(child)
			if err != nil {
				return r, err
			}
			r.JobHistory = append(r.JobHistory, js)
		}
	}

	return r, nil
}
