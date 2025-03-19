// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scan region

package abstract

// Region defines a scanning region.
// See [Dimension] description for definition of the coordinate system.
type Region struct {
	XOffset Dimension // Horizontal offset, 0-based
	YOffset Dimension // Vertical offset, 0-based
	Width   Dimension // Region width
	Height  Dimension // Region height
}

// IsZero reports if Region has a zero value.
func (reg Region) IsZero() bool {
	return reg == Region{}
}
