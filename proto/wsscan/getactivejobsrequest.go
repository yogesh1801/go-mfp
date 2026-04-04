// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// GetActiveJobsRequest: requests a summary of all currently active scan jobs

package wsscan

import "github.com/OpenPrinting/go-mfp/util/xmldoc"

// GetActiveJobsRequest requests a summary of all currently active scan jobs.
// It carries no parameters.
type GetActiveJobsRequest struct{}

// Action returns the [Action] associated with this body.
func (GetActiveJobsRequest) Action() Action { return ActGetActiveJobs }

// ToXML encodes the body into an XML tree.
func (r GetActiveJobsRequest) ToXML() xmldoc.Element {
	return xmldoc.Element{Name: NsWSCN + ":GetActiveJobsRequest"}
}
