// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// WS-Scan message actions (message types)

package wsscan

import "github.com/OpenPrinting/go-mfp/util/xmldoc"

// Action represents a WS-Scan message action (message type).
//
// Each action is represented on the wire by a constant URL string
// in the SOAP header's wsa:Action element.
type Action int

// WS-Scan actions:
const (
	ActUnknown                    Action = iota
	ActGetScannerElements                // GetScannerElements request
	ActGetScannerElementsResponse        // GetScannerElements response
	ActCreateScanJob                     // CreateScanJob request
	ActCreateScanJobResponse             // CreateScanJob response
	ActRetrieveImage                     // RetrieveImage request
	ActRetrieveImageResponse             // RetrieveImage response
)

// actionBaseURL is the common prefix for all WS-Scan action URLs.
const actionBaseURL = "https://schemas.microsoft.com/windows/2006/01/wdp/scan/"

// String returns a short string representation for debugging.
func (act Action) String() string {
	switch act {
	case ActGetScannerElements:
		return "GetScannerElements"
	case ActGetScannerElementsResponse:
		return "GetScannerElementsResponse"
	case ActCreateScanJob:
		return "CreateScanJob"
	case ActCreateScanJobResponse:
		return "CreateScanJobResponse"
	case ActRetrieveImage:
		return "RetrieveImage"
	case ActRetrieveImageResponse:
		return "RetrieveImageResponse"
	}
	return "Unknown"
}

// Encode returns the wire representation (URL string) of the action.
func (act Action) Encode() string {
	s := act.String()
	if s == "Unknown" {
		return ""
	}
	return actionBaseURL + s
}

// decodeAction decodes an [Action] from an XML element's text.
func decodeAction(root xmldoc.Element) (Action, error) {
	act := actDecode(root.Text)
	if act != ActUnknown {
		return act, nil
	}
	return ActUnknown, xmldoc.XMLErrNew(root, "unknown action")
}

// actDecode decodes the wire representation of an action into
// the [Action] value.
func actDecode(s string) Action {
	switch s {
	case actionBaseURL + "GetScannerElements":
		return ActGetScannerElements
	case actionBaseURL + "GetScannerElementsResponse":
		return ActGetScannerElementsResponse
	case actionBaseURL + "CreateScanJob":
		return ActCreateScanJob
	case actionBaseURL + "CreateScanJobResponse":
		return ActCreateScanJobResponse
	case actionBaseURL + "RetrieveImage":
		return ActRetrieveImage
	case actionBaseURL + "RetrieveImageResponse":
		return ActRetrieveImageResponse
	}
	return ActUnknown
}
