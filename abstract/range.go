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
	Min, Max, Normal int // Min, Max and Normal values
	Step             int // Step, ignored if <= 1
}

// IsZero reports if Range has a zero value.
func (r Range) IsZero() bool {
	return r == Range{}
}

// Within reports if value is within the [Range].
//
// It returns:
//   - false, if r.IsZero() is true
//   - false, if v is less that r.Min or greater that r.Max
//   - false, if r.Step is greater that 2 and there is not
//     integer number of steps between r.Min and v
//   - true otherwise
func (r Range) Within(v int) bool {
	switch {
	case r.IsZero():
		return false
	case v < r.Min || r.Max < v:
		return false
	case r.Step <= 1:
		return true
	default:
		return (v-r.Min)%r.Step == 0
	}
}

// validate returns ErrParam error if parameter is not within the Range.
func (r Range) validate(name string, param optional.Val[int]) error {
	if param != nil {
		v := *param
		if !r.Within(v) {
			return ErrParam{ErrUnsupportedParam, name, v}
		}
	}

	return nil
}
