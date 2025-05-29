// MFP - Miulti-Function Printers and scanners toolkit
// Assertions checks.
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Assertions checks.

package assert

// Must panics if condition is not true.
func Must(cond bool) {
	if !cond {
		panic("internal error")
	}
}

// NoError panics if error is not nil.
func NoError(err error) {
	if err != nil {
		panic(err)
	}
}
