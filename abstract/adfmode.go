// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// ADF mode

package abstract

import "fmt"

// ADFMode specifies the ADF scanning mode.
type ADFMode int

// Known intents
const (
	ADFModeUnset   ADFMode = iota // Not set
	ADFModeSimplex                // AFD simplex mode
	ADFModeDuplex                 // AFD duplex mode
	adfModeMax
)

// String returns the string representation of the [ADFMode], for logging.
func (mode ADFMode) String() string {
	switch mode {
	case ADFModeUnset:
		return "Unset"
	case ADFModeSimplex:
		return "Simplex"
	case ADFModeDuplex:
		return "Duplex"
	}

	return fmt.Sprintf("Unknown (%d)", int(mode))
}
