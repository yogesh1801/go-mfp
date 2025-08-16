// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Image depth

package abstract

import "fmt"

// ColorDepth specified image depth, in bits per channel.
type ColorDepth int

// Known color modes:
const (
	ColorDepthUnset ColorDepth = iota // Not set
	ColorDepth8                       // 8 bit (24 bit RGB)
	ColorDepth16                      // 16 bit (48 bit RGB)
	colorDepthMax
)

// String returns the string representation of the [ColorDepth], for logging.
func (depth ColorDepth) String() string {
	switch depth {
	case ColorDepthUnset:
		return "Unset"
	case ColorDepth8:
		return "8"
	case ColorDepth16:
		return "16"
	}

	return fmt.Sprintf("Unknown (%d)", int(depth))
}
