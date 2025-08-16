// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Discrete resolution

package abstract

import "fmt"

// Resolution specifies a discrete scanner resolution.
type Resolution struct {
	XResolution int // X resolution, DPI
	YResolution int // Y resolution, DPI
}

// String returns string representation of [Resolution], for logging.
func (res Resolution) String() string {
	return fmt.Sprintf("%dx%d", res.XResolution, res.YResolution)
}

// IsZero reports if Resolution has zero value.
func (res Resolution) IsZero() bool {
	return res == Resolution{}
}

// Valid reports if Resolution is valid.
func (res Resolution) Valid() bool {
	if res.IsZero() {
		return true
	}

	return res.XResolution > 0 && res.YResolution > 0
}

// ResolutionRange specifies a range of scanner resolutions.
type ResolutionRange struct {
	XMin, XMax, XStep, XNormal int // X resolution range, DPI
	YMin, YMax, YStep, YNormal int // Y resolution range, DPI
}

// IsZero reports if ResolutionRange has zero value.
func (rr ResolutionRange) IsZero() bool {
	return rr == ResolutionRange{}
}
