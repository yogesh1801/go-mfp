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
// The following constraint MUST be true:
//
//	Min <= Normal <= Max
//
// If Min == Max, the corresponding parameter considered unsupported
type Range struct {
	Min, Max, Normal int
}

// Supported tells if the parameter, defined by the [Range] is "supported"
// (i.e., can be changed).
//
// Range considered supported, if r.Min != r.Max
func (r Range) Supported() bool {
	return r.Min != r.Max
}
