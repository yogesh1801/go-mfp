// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scanner resolutions

package escl

import "github.com/alexpevzner/mfp/optional"

// SupportedResolutions defines the set of resolutions,
// supported by the scanner.
type SupportedResolutions struct {
	DiscreteResolutions DiscreteResolutions           // Discrete res
	ResolutionRange     optional.Val[ResolutionRange] // Res range
}

// DiscreteResolutions define a set of discrete resolutions,
// supported by the scanner.
type DiscreteResolutions []DiscreteResolution

// ResolutionRange defines a set of resolutions range,
// supported by the scanner.
type ResolutionRange struct {
	XResolutionRange XYResolutionRange // Horizontal range
	YResolutionRange XYResolutionRange // Vertical range
}

// DiscreteResolution defines a discrete resolution, supported by the scanner.
type DiscreteResolution struct {
	XResolution int // Horizontal resolution, DPI
	YResolution int // Vertical resolution, DPI
}

// XYResolutionRange defines a range of horizontal or vertical resolutions,
// supported by the scanner, in DPI.
type XYResolutionRange = Range
