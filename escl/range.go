// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Common type for Range of some value.

package escl

import "github.com/alexpevzner/mfp/optional"

// Range commonly used to specify the range of some parameter, like
// brightness, contrast etc.
type Range struct {
	Min    int               // Minimal supported value
	Max    int               // Maximal supported value
	Normal int               // Normal value
	Step   optional.Val[int] // Step between the subsequent values
}
