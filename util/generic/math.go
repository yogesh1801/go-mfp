// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Useful mathematical functions

package generic

// LowerDivisibleBy returns v rounded down to the nearest multiple of d.
func LowerDivisibleBy[T Integer](v, d T) T {
	return (v / d) * d
}

// UpperDivisibleBy returns v rounded up to the nearest multiple of d.
func UpperDivisibleBy[T Integer](v, d T) T {
	return ((v + d - 1) / d) * d
}
