// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// ADF state

package escl

import "github.com/OpenPrinting/go-mfp/util/xmldoc"

// ADFState represents the ADF state
type ADFState int

// Known ADF states
const (
	UnknownADFState               ADFState = iota // Unknown ADF state
	ScannerAdfProcessing                          // This is the OK state
	ScannerAdfEmpty                               // ADF is empty
	ScannerAdfJam                                 // Paper jam in the ADF
	ScannerAdfLoaded                              // ADF is loaded with sheets
	ScannerAdfMispick                             // ADF can't peek a sheet
	ScannerAdfHatchOpen                           // ADF hatch is open
	ScannerAdfDuplexPageTooShort                  // Sheet is too short
	ScannerAdfDuplexPageTooLong                   // Sheet is too long
	ScannerAdfMultipickDetected                   // Multiple pick detected
	ScannerAdfInputTrayFailed                     // ADF tray failure
	ScannerAdfInputTrayOverloaded                 // Too many sheets in ADF
)

// decodeADFState decodes [ADFState] from the XML tree.
func decodeADFState(root xmldoc.Element) (state ADFState, err error) {
	return decodeEnum(root, DecodeADFState)
}

// toXML generates XML tree for the [ADFState].
func (state ADFState) toXML(name string) xmldoc.Element {
	return xmldoc.Element{
		Name: name,
		Text: state.String(),
	}
}

// String returns a string representation of the [ADFState]
func (state ADFState) String() string {
	switch state {
	case ScannerAdfProcessing:
		return "ScannerAdfProcessing"
	case ScannerAdfEmpty:
		return "ScannerAdfEmpty"
	case ScannerAdfJam:
		return "ScannerAdfJam"
	case ScannerAdfLoaded:
		return "ScannerAdfLoaded"
	case ScannerAdfMispick:
		return "ScannerAdfMispick"
	case ScannerAdfHatchOpen:
		return "ScannerAdfHatchOpen"
	case ScannerAdfDuplexPageTooShort:
		return "ScannerAdfDuplexPageTooShort"
	case ScannerAdfDuplexPageTooLong:
		return "ScannerAdfDuplexPageTooLong"
	case ScannerAdfMultipickDetected:
		return "ScannerAdfMultipickDetected"
	case ScannerAdfInputTrayFailed:
		return "ScannerAdfInputTrayFailed"
	case ScannerAdfInputTrayOverloaded:
		return "ScannerAdfInputTrayOverloaded"
	}

	return "Unknown"
}

// DecodeADFState decodes [ADFState] out of its XML string
// representation.
func DecodeADFState(s string) ADFState {
	switch s {
	case "ScannerAdfProcessing":
		return ScannerAdfProcessing
	case "ScannerAdfEmpty":
		return ScannerAdfEmpty
	case "ScannerAdfJam":
		return ScannerAdfJam
	case "ScannerAdfLoaded":
		return ScannerAdfLoaded
	case "ScannerAdfMispick":
		return ScannerAdfMispick
	case "ScannerAdfHatchOpen":
		return ScannerAdfHatchOpen
	case "ScannerAdfDuplexPageTooShort":
		return ScannerAdfDuplexPageTooShort
	case "ScannerAdfDuplexPageTooLong":
		return ScannerAdfDuplexPageTooLong
	case "ScannerAdfMultipickDetected":
		return ScannerAdfMultipickDetected
	case "ScannerAdfInputTrayFailed":
		return ScannerAdfInputTrayFailed
	case "ScannerAdfInputTrayOverloaded":
		return ScannerAdfInputTrayOverloaded
	}

	return UnknownADFState
}
