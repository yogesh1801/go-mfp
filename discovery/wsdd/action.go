// MFP - Miulti-Function Printers and scanners toolkit
// WSD device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// WSDD message actions (message types)

package wsdd

// action represents a message action (or message type).
//
// Each action represented on the wire by appropriate URL
// string (e.g., http://schemas.xmlsoap.org/ws/2004/09/transfer/Get
// for probe).
type action int

// Message actions:
const (
	actUnknown = iota // Other (unknown) action
	actHello
	actBye
	actProbe
	actProbeMatches
	actResolve
	actResolveMatches
	actGet
	actGetResponse
)

// String represents action as a short string, for debugging.
func (act action) String() string {
	switch act {
	case actHello:
		return "Hello"
	case actBye:
		return "Bye"
	case actProbe:
		return "Probe"
	case actProbeMatches:
		return "ProbeMatches"
	case actResolve:
		return "Resolve"
	case actResolveMatches:
		return "ResolveMatches"
	case actGet:
		return "Get"
	case actGetResponse:
		return "GetResponse"
	}

	return "Unknown"
}

// Encode represents action as a string for wire encoding.
// For unknown action it returns "".
func (act action) Encode() string {
	switch act {
	case actHello:
		return "http://schemas.xmlsoap.org/ws/2005/04/discovery/Hello"
	case actBye:
		return "http://schemas.xmlsoap.org/ws/2005/04/discovery/Bye"
	case actProbe:
		return "http://schemas.xmlsoap.org/ws/2005/04/discovery/Probe"
	case actProbeMatches:
		return "http://schemas.xmlsoap.org/ws/2005/04/discovery/ProbeMatches"
	case actResolve:
		return "http://schemas.xmlsoap.org/ws/2005/04/discovery/Resolve"
	case actResolveMatches:
		return "http://schemas.xmlsoap.org/ws/2005/04/discovery/ResolveMatches"
	case actGet:
		return "http://schemas.xmlsoap.org/ws/2004/09/transfer/Get"
	case actGetResponse:
		return "http://schemas.xmlsoap.org/ws/2004/09/transfer/GetResponse"
	}

	return ""
}

// actDecode decodes wire representation of action into the action number.
// For unknown actions it returns actOther
func actDecode(s string) action {
	switch s {
	case "http://schemas.xmlsoap.org/ws/2005/04/discovery/Hello":
		return actHello
	case "http://schemas.xmlsoap.org/ws/2005/04/discovery/Bye":
		return actBye
	case "http://schemas.xmlsoap.org/ws/2005/04/discovery/Probe":
		return actProbe
	case "http://schemas.xmlsoap.org/ws/2005/04/discovery/ProbeMatches":
		return actProbeMatches
	case "http://schemas.xmlsoap.org/ws/2005/04/discovery/Resolve":
		return actResolve
	case "http://schemas.xmlsoap.org/ws/2005/04/discovery/ResolveMatches":
		return actResolveMatches
	case "http://schemas.xmlsoap.org/ws/2004/09/transfer/Get":
		return actGet
	case "http://schemas.xmlsoap.org/ws/2004/09/transfer/GetResponse":
		return actGetResponse
	}

	return actUnknown
}
