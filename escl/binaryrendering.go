// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Binary rendering for BlackAndWhite1 mode

package escl

// BinaryRendering specifies how to render black and white images
// in the BlackAndWhite1 mode.
type BinaryRendering int

// BinaryRendering modes:
const (
	UnknownBinaryRendering BinaryRendering = iota - 1 // Unknown CCD
	Halftone                                          // Simulate Halftone
	Threshold                                         // Use Threshold
)

// String returns a string representation of the [BinaryRendering]
func (rnd BinaryRendering) String() string {
	switch rnd {
	case Halftone:
		return "Halftone"
	case Threshold:
		return "Threshold"
	}

	return "Unknown"
}

// DecodeBinaryRendering decodes [BinaryRendering] out of its XML string representation.
func DecodeBinaryRendering(s string) BinaryRendering {
	switch s {
	case "Halftone":
		return Halftone
	case "Threshold":
		return Threshold
	}

	return UnknownBinaryRendering
}
