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

// Within reports if reg2 is within reg.
func (reg Region) Within(reg2 Region) bool {
	return reg.XOffset <= reg2.XOffset &&
		reg2.XOffset+reg2.Width <= reg.XOffset+reg.Width &&
		reg.YOffset <= reg2.YOffset &&
		reg2.YOffset+reg2.Height <= reg.YOffset+reg.Height
}
