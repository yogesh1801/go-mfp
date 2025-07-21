// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// content type value

package wsscan

import (
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// ContentTypeValue defines the document content type
// supported by the scan device.
type ContentTypeValue int

// Known content type values:
const (
	UnknownContentTypeValue ContentTypeValue = iota
	Auto
	Text
	Photo
	Halftone
	Mixed
)

// decodeContentTypeValue decodes [ContentTypeValue] from the XML tree.
func decodeContentTypeValue(root xmldoc.Element) (
	ctv ContentTypeValue, err error) {
	return decodeEnum(root, DecodeContentTypeValue)
}

// toXML generates XML tree for the [ContentTypeValue].
func (ctv ContentTypeValue) toXML(name string) xmldoc.Element {
	return xmldoc.Element{
		Name: name,
		Text: ctv.String(),
	}
}

// String returns a string representation of the [ContentTypeValue].
func (ctv ContentTypeValue) String() string {
	switch ctv {
	case Auto:
		return "Auto"
	case Text:
		return "Text"
	case Photo:
		return "Photo"
	case Halftone:
		return "Halftone"
	case Mixed:
		return "Mixed"
	}
	return "Unknown"
}

// DecodeContentTypeValue decodes [ContentTypeValue]
// out of its XML string representation.
func DecodeContentTypeValue(s string) ContentTypeValue {
	switch s {
	case "Auto":
		return Auto
	case "Text":
		return Text
	case "Photo":
		return Photo
	case "Halftone":
		return Halftone
	case "Mixed":
		return Mixed
	}
	return UnknownContentTypeValue
}
