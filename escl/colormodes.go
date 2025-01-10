// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Set of scan color modes

package escl

import "github.com/alexpevzner/mfp/internal/generic"

// ColorModes contains a set of [ColorMode]s.
type ColorModes struct {
	generic.Bitset[ColorMode]           // Set of modes
	Default                   ColorMode // Default color mode
}

// MakeColorModes makes [ColorModes] from the list of [ColorMode]s.
func MakeColorModes(list ...ColorMode) ColorModes {
	return ColorModes{
		generic.MakeBitset(list...),
		UnknownColorMode,
	}
}
