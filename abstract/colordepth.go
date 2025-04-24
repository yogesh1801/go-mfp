// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Image depth

package abstract

// ColorDepth specified image depth, in bits per channel.
type ColorDepth int

// Known color modes:
const (
	ColorDepthUnset ColorDepth = iota // Not set
	ColorDepth8                       // 8 bit (24 bit RGB)
	ColorDepth16                      // 16 bit (48 bit RGB)
	colorDepthMax
)
