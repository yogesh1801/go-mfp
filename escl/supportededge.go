// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Image edges, for edge auto detection

package escl

// SupportedEdge represents an image edge, suitable for edge auto detection.
type SupportedEdge int

// Supported edges:
const (
	UnknownSupportedEdge SupportedEdge = iota // Unknown intent
	TopEdge
	LeftEdge
	BottomEdge
	RightEdge
)

// String returns a string representation of the [SupportedEdge]
func (intent SupportedEdge) String() string {
	switch intent {
	case TopEdge:
		return "TopEdge"
	case LeftEdge:
		return "LeftEdge"
	case BottomEdge:
		return "BottomEdge"
	case RightEdge:
		return "RightEdge"
	}

	return "Unknown"
}

// DecodeSupportedEdge decodes [SupportedEdge] out of its XML string representation.
func DecodeSupportedEdge(s string) SupportedEdge {
	switch s {
	case "TopEdge":
		return TopEdge
	case "LeftEdge":
		return LeftEdge
	case "BottomEdge":
		return BottomEdge
	case "RightEdge":
		return RightEdge
	}

	return UnknownSupportedEdge
}
