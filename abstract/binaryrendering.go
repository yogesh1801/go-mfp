// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Binary rendering for BlackAndWhite1 mode

package abstract

import "fmt"

// BinaryRendering specifies how to render black and white images
// in the BlackAndWhite1 mode.
type BinaryRendering int

// BinaryRendering modes:
const (
	BinaryRenderingUnset     BinaryRendering = iota // Not set
	BinaryRenderingHalftone                         // Simulate Halftone
	BinaryRenderingThreshold                        // Use Threshold
	binaryRenderingMax
)

// String returns the string representation of the [BinaryRendering],
// for logging.
func (rend BinaryRendering) String() string {
	switch rend {
	case BinaryRenderingUnset:
		return "Unset"
	case BinaryRenderingHalftone:
		return "Halftone"
	case BinaryRenderingThreshold:
		return "Threshold"
	}

	return fmt.Sprintf("Unknown (%d)", int(rend))
}
