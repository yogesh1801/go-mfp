// MFP - Miulti-Function Printers and scanners toolkit
// Assertions checks.
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Assertions checks.

package assert

// True panics if condition is not true.
func True(cond bool) {
	if !cond {
		panic("internal error")
	}
}

// True panics if error is not nil.
func NoError(err error) {
	if err != nil {
		panic(err)
	}
}
