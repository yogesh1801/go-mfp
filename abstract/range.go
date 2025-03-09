// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Range of "analog" parameter.

package abstract

// Range defines a range of "analog" parameter, like
// brightness, contrast and similar.
//
// The following constraints MUST be true:
//
//	-1.0 <= Min <= 0 && 0 <= Max <= +1.0
//	Min <= Max
//
// The 0 value of the appropriate parameter means the "normal" value.
// Note, if Min == Max, it means that the appropriate parameter
// cannot be adjusted by user
type Range struct {
	Min, Max float64
}
