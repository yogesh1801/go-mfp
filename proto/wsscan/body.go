// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// Message Body interface

package wsscan

import "github.com/OpenPrinting/go-mfp/util/xmldoc"

// Body represents a WS-Scan SOAP message body.
//
// Body can be one of the following types:
//   - [GetScannerElementsRequest]
//   - [GetScannerElementsResponse]
//   - [CreateScanJobRequest]
//   - [CreateScanJobResponse]
//   - [RetrieveImageRequest]
//   - [RetrieveImageResponse]
type Body interface {
	// Action returns the [Action] associated with this body.
	Action() Action

	// ToXML encodes the body into an XML tree.
	ToXML() xmldoc.Element
}
