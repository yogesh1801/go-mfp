// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Image edges, for edge auto detection

package escl

import "github.com/alexpevzner/mfp/util/xmldoc"

// SupportedEdge represents an image edge, suitable for edge auto detection.
type SupportedEdge int

// Supported edges:
const (
	UnknownSupportedEdge SupportedEdge = iota // Unknown edge
	TopEdge
	LeftEdge
	BottomEdge
	RightEdge
)

// decodeSupportedEdge decodes [SupportedEdge] from the XML tree.
func decodeSupportedEdge(root xmldoc.Element) (edge SupportedEdge, err error) {
	return decodeEnum(root, DecodeSupportedEdge)
}

// toXML generates XML tree for the [SupportedEdge].
func (edge SupportedEdge) toXML(name string) xmldoc.Element {
	return xmldoc.Element{
		Name: name,
		Text: edge.String(),
	}
}

// String returns a string representation of the [SupportedEdge]
func (edge SupportedEdge) String() string {
	switch edge {
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
