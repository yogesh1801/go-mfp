// MFP - Miulti-Function Printers and scanners toolkit
// Useful generics
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Generic Min and Max functions

package generic

// Min returns minimum of two values, a and b
func Min[T Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Max returns maximum of two values, a and b
func Max[T Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}
