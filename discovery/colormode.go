// MFP - Miulti-Function Printers and scanners toolkit
// Device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Supported Color Mode

package discovery

// ColorMode defines the color space capabilities of printer or scanner.
type ColorMode int

// ColorMode bits:
const (
	ColorGrayscale ColorMode = 1 << iota // Gray scale print and scan
	ColorRGB                             // RGB color
	ColorBW                              // 1-bit monochrome
)
