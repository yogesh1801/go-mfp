// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scan color mode

package abstract

// ColorMode specifies the color space capabilities of printer or scanner.
type ColorMode int

// Known color modes:
const (
	ColorModeUnset  ColorMode = iota // Not set
	ColorModeBinary                  // 1-bit monochrome
	ColorModeMono                    // Gray scale monochrome
	ColorModeColor                   // Full-color mode
	colorModeMax
)
