// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Range of "analog" parameter.

package abstract

import "github.com/alexpevzner/mfp/util/optional"

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

// IsZero reports if Range has a zero value.
func (r Range) IsZero() bool {
	return r == Range{}
}

// Within reports if value is within the [Range].
func (r Range) Within(v int) bool {
	return r.Min <= v && v <= r.Max
}

// validate returns ErrParam errir if parameter is not within the Range.
func (r Range) validate(name string, param optional.Val[int]) error {
	if param != nil {
		v := *param
		if !r.Within(v) {
			return ErrParam{ErrUnsupportedParam, name, v}
		}
	}

	return nil
}
