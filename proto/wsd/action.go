// MFP - Miulti-Function Printers and scanners toolkit
// WSD core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// WSD message actions (message types)

package wsd

import "github.com/OpenPrinting/go-mfp/util/xmldoc"

// Action represents a message action (or message type).
//
// Each action represented on the wire by appropriate URL
// string (e.g., http://schemas.xmlsoap.org/ws/2004/09/transfer/Get
// for probe).
type Action int

// Message actions:
const (
	ActUnknown Action = iota // Other (unknown) action
	ActHello
	ActBye
	ActProbe
	ActProbeMatches
	ActResolve
	ActResolveMatches
	ActGet
	ActGetResponse
)

// String represents action as a short string, for debugging.
func (act Action) String() string {
	switch act {
	case ActHello:
		return "Hello"
	case ActBye:
		return "Bye"
	case ActProbe:
		return "Probe"
	case ActProbeMatches:
		return "ProbeMatches"
	case ActResolve:
		return "Resolve"
	case ActResolveMatches:
		return "ResolveMatches"
	case ActGet:
		return "Get"
	case ActGetResponse:
		return "GetResponse"
	}

	return "Unknown"
}

// bodyname returns name of the Body child element (namespace:name),
// depending on the Action.
//
// If child is not expected, "" is returned.
func (act Action) bodyname() string {
	switch act {
	case ActHello:
		return NsDiscovery + ":Hello"
	case ActBye:
		return NsDiscovery + ":Bye"
	case ActProbe:
		return NsDiscovery + ":Probe"
	case ActProbeMatches:
		return NsDiscovery + ":ProbeMatches"
	case ActResolve:
		return NsDiscovery + ":Resolve"
	case ActResolveMatches:
		return NsDiscovery + ":ResolveMatches"
	case ActGet:
		return ""
	case ActGetResponse:
		return NsMex + ":Metadata"
	}

	return ""
}

// Encode represents action as a string for wire encoding.
// For unknown action it returns "".
func (act Action) Encode() string {
	switch act {
	case ActHello:
		return "http://schemas.xmlsoap.org/ws/2005/04/discovery/Hello"
	case ActBye:
		return "http://schemas.xmlsoap.org/ws/2005/04/discovery/Bye"
	case ActProbe:
		return "http://schemas.xmlsoap.org/ws/2005/04/discovery/Probe"
	case ActProbeMatches:
		return "http://schemas.xmlsoap.org/ws/2005/04/discovery/ProbeMatches"
	case ActResolve:
		return "http://schemas.xmlsoap.org/ws/2005/04/discovery/Resolve"
	case ActResolveMatches:
		return "http://schemas.xmlsoap.org/ws/2005/04/discovery/ResolveMatches"
	case ActGet:
		return "http://schemas.xmlsoap.org/ws/2004/09/transfer/Get"
	case ActGetResponse:
		return "http://schemas.xmlsoap.org/ws/2004/09/transfer/GetResponse"
	}

	return ""
}

// ActDecode decodes wire representation of action into the action number.
// For unknown actions it returns actOther
func ActDecode(s string) Action {
	switch s {
	case "http://schemas.xmlsoap.org/ws/2005/04/discovery/Hello":
		return ActHello
	case "http://schemas.xmlsoap.org/ws/2005/04/discovery/Bye":
		return ActBye
	case "http://schemas.xmlsoap.org/ws/2005/04/discovery/Probe":
		return ActProbe
	case "http://schemas.xmlsoap.org/ws/2005/04/discovery/ProbeMatches":
		return ActProbeMatches
	case "http://schemas.xmlsoap.org/ws/2005/04/discovery/Resolve":
		return ActResolve
	case "http://schemas.xmlsoap.org/ws/2005/04/discovery/ResolveMatches":
		return ActResolveMatches
	case "http://schemas.xmlsoap.org/ws/2004/09/transfer/Get":
		return ActGet
	case "http://schemas.xmlsoap.org/ws/2004/09/transfer/GetResponse":
		return ActGetResponse
	}

	return ActUnknown
}

// DecodeAction decodes action, from the XML tree
func DecodeAction(root xmldoc.Element) (v Action, err error) {
	act := ActDecode(root.Text)
	if act != ActUnknown {
		return act, nil
	}

	return ActUnknown, xmldoc.XMLErrNew(root, "unknown action")
}
