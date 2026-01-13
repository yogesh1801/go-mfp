// MFP - Miulti-Function Printers and scanners toolkit
// Assertions checks.
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Assertions checks.

package assert

import "fmt"

// Must panics if condition is not true.
func Must(cond bool) {
	MustMsg(cond, "internal error")
}

// NoError panics if error is not nil.
func NoError(err error) {
	if err != nil {
		panic(err)
	}
}

// MustMsg panics if condition is not true.
// If it panics, format and optional args define its final message.
func MustMsg(cond bool, format string, args ...any) {
	if !cond {
		panic(fmt.Sprintf(format, args...))
	}
}
