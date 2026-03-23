// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// WS-Scan message actions (message types)

package wsscan

import (
	"strings"

	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// ActionType identifies a WS-Scan message action (message type).
type ActionType int

// WS-Scan action types:
const (
	ActUnknown                    ActionType = iota
	ActGetScannerElements                    // GetScannerElements request
	ActGetScannerElementsResponse            // GetScannerElements response
)

// Action represents a WS-Scan message action, combining the
// action type with the base URL received on the wire.
type Action struct {
	Type    ActionType // The action type
	BaseURL string     // Base URL from the wire
}

// String returns a short string representation for debugging.
func (act Action) String() string {
	return act.Type.String()
}

// String returns the suffix string for the action type.
func (at ActionType) String() string {
	switch at {
	case ActGetScannerElements:
		return "GetScannerElements"
	case ActGetScannerElementsResponse:
		return "GetScannerElementsResponse"
	}
	return "Unknown"
}

// Encode returns the wire representation (URL string) of the action.
func (act Action) Encode() string {
	s := act.Type.String()
	if s == "Unknown" {
		return ""
	}
	return act.BaseURL + s
}

// decodeAction decodes an [Action] from an XML element's text.
func decodeAction(root xmldoc.Element) (Action, error) {
	act := actDecode(root.Text)
	if act.Type != ActUnknown {
		return act, nil
	}
	return act, xmldoc.XMLErrNew(root, "unknown action")
}

// actDecode decodes the wire representation of an action into
// an [Action] value. It matches by suffix so the base URL can
// vary between implementations.
func actDecode(s string) Action {
	i := strings.LastIndex(s, "/")
	if i < 0 {
		return Action{}
	}

	baseURL := s[:i+1]
	suffix := s[i+1:]

	switch suffix {
	case "GetScannerElementsRequest":
		return Action{Type: ActGetScannerElements, BaseURL: baseURL}
	case "GetScannerElementsResponse":
		return Action{Type: ActGetScannerElementsResponse, BaseURL: baseURL}
	}
	return Action{}
}
