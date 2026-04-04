// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// CancelJobResponse: server response to a CancelJob request

package wsscan

import "github.com/OpenPrinting/go-mfp/util/xmldoc"

// CancelJobResponse is the server's response to a [CancelJobRequest].
// It carries no parameters.
type CancelJobResponse struct{}

// Action returns the [Action] associated with this body.
func (CancelJobResponse) Action() Action { return ActCancelJobResponse }

// ToXML encodes the body into an XML tree.
func (r CancelJobResponse) ToXML() xmldoc.Element {
	return xmldoc.Element{Name: NsWSCN + ":CancelJobResponse"}
}
