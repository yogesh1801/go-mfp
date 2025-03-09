// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scan region

package abstract

// Region defines a scanning region.
// See [Coord] description for definition of the coordinate system.
type Region struct {
	XOffset Coord // Horizontal offset, 0-based
	YOffset Coord // Vertical offset, 0-based
	Width   Coord // Region width
	Height  Coord // Region height
}
