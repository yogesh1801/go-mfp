// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scanner state

package escl

import "github.com/alexpevzner/mfp/xmldoc"

// ScannerState represents the overall scanner state
type ScannerState int

// Known scanner states
const (
	UnknownScannerState ScannerState = iota // Unknown scanner state
	ScannerIdle                             // The scanner is idle
	ScannerProcessing                       // The scanner is doing some work
	ScannerTesting                          // Calibrating, preparing the unit
	ScannerStopped                          // Stopped: error condition occured
	ScannerDown                             // Down: unit is unavailable
)

// decodeScannerState decodes [ScannerState] from the XML tree.
func decodeScannerState(root xmldoc.Element) (state ScannerState, err error) {
	return decodeEnum(root, DecodeScannerState)
}

// toXML generates XML tree for the [ScannerState].
func (state ScannerState) toXML(name string) xmldoc.Element {
	return xmldoc.Element{
		Name: name,
		Text: state.String(),
	}
}

// String returns a string representation of the [ScannerState]
func (state ScannerState) String() string {
	switch state {
	case ScannerIdle:
		return "Idle"
	case ScannerProcessing:
		return "Processing"
	case ScannerTesting:
		return "Testing"
	case ScannerStopped:
		return "Stopped"
	case ScannerDown:
		return "Down"
	}

	return "Unknown"
}

// DecodeScannerState decodes [ScannerState] out of its XML string
// representation.
func DecodeScannerState(s string) ScannerState {
	switch s {
	case "Idle":
		return ScannerIdle
	case "Processing":
		return ScannerProcessing
	case "Testing":
		return ScannerTesting
	case "Stopped":
		return ScannerStopped
	case "Down":
		return ScannerDown
	}

	return UnknownScannerState
}
