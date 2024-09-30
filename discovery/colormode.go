// MFP - Miulti-Function Printers and scanners toolkit
// Device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Supported Color Mode

package discovery

import "strings"

// ColorMode defines the color space capabilities of printer or scanner.
type ColorMode int

// ColorMode bits:
const (
	ColorOther     ColorMode = 1 << iota // Other color mode
	ColorGrayscale                       // Gray scale print and scan
	ColorRGB                             // RGB color
	ColorBW                              // 1-bit monochrome
)

// String formats ColorMode as string, for printing and logging
func (cm ColorMode) String() string {
	s := []string{}

	if cm&ColorOther != 0 {
		s = append(s, "other")
	}
	if cm&ColorGrayscale != 0 {
		s = append(s, "grayscale")
	}
	if cm&ColorRGB != 0 {
		s = append(s, "RGB")
	}
	if cm&ColorBW != 0 {
		s = append(s, "BW")
	}

	return strings.Join(s, ",")
}
