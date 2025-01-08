// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Set of CCD binary rendering modes

package escl

import "github.com/alexpevzner/mfp/internal/generic"

// BinaryRenderings contains a set (bitmask) of [BinaryRendering]s.
type BinaryRenderings = generic.Bitset[BinaryRendering]

// MakeBinaryRenderings makes [BinaryRenderings] from the list of
// [BinaryRendering]s.
func MakeBinaryRenderings(list ...BinaryRendering) BinaryRenderings {
	return generic.MakeBitset(list...)
}
