// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Dimension (Coordinates)

package abstract

// Dimension represents image coordinates and sizes.
//
// Coordinates are 0-based, with the (0,0) point placed at the
// image top-left corner:
//
//	(0,0)
//	+-----> X/Width
//	|
//	|
//	V
//	Y/Height
//
// Units are 1/100 of millimeters (1/2540th of an inch).
// This is consistent with the coordinate system, used by IPP.
type Dimension int

// Common units, represented as [Dimension]
const (
	Millimeter Dimension = 100
	Centimeter           = 10 * Millimeter
	Inch       Dimension = 2540
)

// Dots converts Dimension value into number of image dots,
// assuming specified DPI (dots per inch).
func (dim Dimension) Dots(dpi int) int {
	tmp := uint64(dim)
	tmp *= uint64(dpi)
	tmp += uint64(Inch / 2)
	tmp /= uint64(Inch)

	return int(tmp)
}
