// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Set of CCD binary rendering modes

package escl

import "github.com/alexpevzner/mfp/internal/generic"

// BinaryRenderings contains a set of [BinaryRendering]s.
type BinaryRenderings struct {
	generic.Bitset[BinaryRendering]                 // Set of modes
	Default                         BinaryRendering // Default mode
}

// MakeBinaryRenderings makes [BinaryRenderings] from the list of
// [BinaryRendering]s.
func MakeBinaryRenderings(list ...BinaryRendering) BinaryRenderings {
	return BinaryRenderings{
		generic.MakeBitset(list...),
		UnknownBinaryRendering,
	}
}
