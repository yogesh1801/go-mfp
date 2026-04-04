// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// GetJobHistoryRequest: requests a summary of completed scan jobs

package wsscan

import "github.com/OpenPrinting/go-mfp/util/xmldoc"

// GetJobHistoryRequest requests a summary of completed scan jobs.
// It carries no parameters.
type GetJobHistoryRequest struct{}

// Action returns the [Action] associated with this body.
func (GetJobHistoryRequest) Action() Action { return ActGetJobHistory }

// ToXML encodes the body into an XML tree.
func (r GetJobHistoryRequest) ToXML() xmldoc.Element {
	return xmldoc.Element{Name: NsWSCN + ":GetJobHistoryRequest"}
}
