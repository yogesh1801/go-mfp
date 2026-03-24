// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Discrete resolution

package abstract

import (
	"fmt"
	"math"
)

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

// euclideanDistance returns Euclidean distance (actually, the
// square of the Euclidean distance) between two resolutions
func (res Resolution) euclideanDistance(res2 Resolution) int {
	dx := res.XResolution - res2.XResolution
	dy := res.YResolution - res2.YResolution

	return dx*dx + dy*dy
}

// lessSquare returns true if res is less square than res2.
func (res Resolution) lessSquare(res2 Resolution) bool {
	ratio1 := res.aspectRatio()
	ratio2 := res2.aspectRatio()

	return ratio1 < ratio2
}

// aspectRatio returns the ratio of the larger side to the smaller side
// Higher value means less square resolution.
//
// For invalid resolutions (zero values), returns maximum possible value.
func (res Resolution) aspectRatio() float64 {
	// Edge cases: if either side is <= 0, treat as maximally non-square
	if res.XResolution <= 0 || res.YResolution <= 0 {
		return math.MaxFloat64
	}

	x := float64(res.XResolution)
	y := float64(res.YResolution)

	if x >= y {
		return x / y
	}
	return y / x
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
