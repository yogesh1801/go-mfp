// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Image depth

package abstract

// Depth specified image depth, in bits per channel.
type Depth int

// Known color modes:
const (
	DepthUnset Depth = 0  // Not set
	Depth8     Depth = 8  // 8 bit (24 bit RGB)
	Depth16    Depth = 16 // 16 bit (48 bit RGB)
	depthMax
)
