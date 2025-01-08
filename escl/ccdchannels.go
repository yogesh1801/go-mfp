// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Set of CCD channels

package escl

import "github.com/alexpevzner/mfp/internal/generic"

// CcdChannels contains a set (bitmask) of [CcdChannel]s.
type CcdChannels = generic.Bitset[CcdChannel]

// MakeCcdChannels makes [CcdChannels] from the list of [CcdChannel]s.
func MakeCcdChannels(list ...CcdChannel) CcdChannels {
	return generic.MakeBitset(list...)
}
