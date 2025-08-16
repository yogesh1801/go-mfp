// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scan intent

package abstract

import "fmt"

// Intent represents a scan intent, which implies automatic
// choosing of the appropriate scan parameters by the scanner.
type Intent int

// Known intents
const (
	IntentUnset          Intent = iota // Not set
	IntentDocument                     // Scanning optimized for text
	IntentTextAndGraphic               // Document with text and graphics
	IntentPhoto                        // Scanning optimized for photo
	IntentPreview                      // Preview scanning
	IntentObject                       // 3d object
	IntentBusinessCard                 // Business card
	intentMax
)

// String returns the string representation of the [Intent], for logging.
func (intent Intent) String() string {
	switch intent {
	case IntentUnset:
		return "Unset"
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

	return fmt.Sprintf("Unknown (%d)", int(intent))
}
