// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scan intent

package escl

import "github.com/alexpevzner/mfp/xmldoc"

// Intent represents a scan intent, which implies automatic
// choosing of the appropriate scan parameters by the scanner.
type Intent int

// Known intents
const (
	UnknownIntent  Intent = iota // Unknown intent
	Document                     // Scanning optimized for text
	TextAndGraphic               // Document with text and graphics
	Photo                        // Scanning optimized for photo
	Preview                      // Preview scanning
	Object                       // 3d object
	BusinessCard                 // Scanning optimized for business card
)

// decodeIntent decodes [Intent] from the XML tree.
func decodeIntent(root xmldoc.Element) (intent Intent, err error) {
	return decodeEnum(root, DecodeIntent, NsScan)
}

// toXML generates XML tree for the [Intent].
func (intent Intent) toXML(name string) xmldoc.Element {
	return xmldoc.Element{
		Name: name,
		Text: NsScan + ":" + intent.String(),
	}
}

// String returns a string representation of the [Intent]
func (intent Intent) String() string {
	switch intent {
	case Document:
		return "Document"
	case TextAndGraphic:
		return "TextAndGraphic"
	case Photo:
		return "Photo"
	case Preview:
		return "Preview"
	case Object:
		return "Object"
	case BusinessCard:
		return "BusinessCard"
	}

	return "Unknown"
}

// DecodeIntent decodes [Intent] out of its XML string representation.
func DecodeIntent(s string) Intent {
	switch s {
	case "Document":
		return Document
	case "TextAndGraphic":
		return TextAndGraphic
	case "Photo":
		return Photo
	case "Preview":
		return Preview
	case "Object":
		return Object
	case "BusinessCard":
		return BusinessCard
	}

	return UnknownIntent
}
