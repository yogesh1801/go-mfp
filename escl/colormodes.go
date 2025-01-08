// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Set of scan color modes

package escl

import "github.com/alexpevzner/mfp/internal/generic"

// ColorModes contains a set (bitmask) of [ColorMode]s.
type ColorModes = generic.Bitset[ColorMode]

// MakeColorModes makes [ColorModes] from the list of [ColorMode]s.
func MakeColorModes(list ...ColorMode) ColorModes {
	return generic.MakeBitset(list...)
}
