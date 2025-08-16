// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Input source

package abstract

import "fmt"

// Input specifies the desired input source.
type Input int

// Known intents
const (
	InputUnset  Input = iota // Not set
	InputPlaten              // Scan from platen
	InputADF                 // Automatic Document Feeder
	InputCamera              // Scan from camera
	inputMax
)

// String returns the string representation of the [Input], for logging
func (in Input) String() string {
	switch in {
	case InputUnset:
		return "Unset"
	case InputPlaten:
		return "Platen"
	case InputADF:
		return "ADF"
	case InputCamera:
		return "Camera"
	}

	return fmt.Sprintf("Unknown (%d)", int(in))
}
