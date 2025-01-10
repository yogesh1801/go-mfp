// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Set of scan color modes

package escl

import "github.com/alexpevzner/mfp/internal/generic"

// ColorSpaces contains a set of [ColorSpace]s.
type ColorSpaces struct {
	generic.Bitset[ColorSpace]            // Set of modes
	Default                    ColorSpace // Default color mode
}

// MakeColorSpaces makes [ColorSpaces] from the list of [ColorSpace]s.
func MakeColorSpaces(list ...ColorSpace) ColorSpaces {
	return ColorSpaces{
		generic.MakeBitset(list...),
		UnknownColorSpace,
	}
}
