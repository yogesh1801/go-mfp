// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scan intent

package escl

// Intent represents a scan intent, which implies automatic
// choosing of the appropriate scan parameters by the scanner.
type Intent int

// Known intents
const (
	IntentUnknown        Intent = iota - 1 // Unknown intent
	IntentDocument                         // Scanning optimized for text
	IntentTextAndGraphic                   // Doc with text and graphics
	IntentPhoto                            // Scanning optimized for photo
	IntentPreview                          // Preview scanning
	IntentObject                           // 3d object
	IntentBusinessCard                     // Scanning of business card
)

// String returns a string representation of the [Intent]
func (intent Intent) String() string {
	switch intent {
	case IntentDocument:
		return "Document"
	case IntentTextAndGraphic:
		return "TextAndGraphic"
	case IntentPhoto:
		return "Photo"
	case IntentPreview:
		return "Preview"
	case IntentObject:
		return "Object"
	case IntentBusinessCard:
		return "BusinessCard"
	}

	return "Unknown"
}

// DecodeIntent decodes [Intent] out of its XML string representation.
func DecodeIntent(s string) Intent {
	switch s {
	case "Document":
		return IntentDocument
	case "TextAndGraphic":
		return IntentTextAndGraphic
	case "Photo":
		return IntentPhoto
	case "Preview":
		return IntentPreview
	case "Object":
		return IntentObject
	case "BusinessCard":
		return IntentBusinessCard
	}

	return IntentUnknown
}
