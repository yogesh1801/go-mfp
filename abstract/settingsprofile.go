// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Settings profile

package abstract

import "github.com/alexpevzner/mfp/util/generic"

// SettingsProfile defines a valid combination of scan parameters setting.
//
// Each [InputCapabilities] may include multiple instances of
// SettingsProfile.
//
// SettingsProfile may be constrained by [ColorMode], [Depth]
// and [BinaryRendering].
type SettingsProfile struct {
	ColorModes       generic.Bitset[ColorMode]       // Allowed color modes
	Depts            generic.Bitset[Depth]           // Allowed depths
	BinaryRenderings generic.Bitset[BinaryRendering] // For 1-bit B&W
	Resolutions      []Resolution                    // Allowed resolutions
	ResolutionRanges []ResolutionRange               // Resolution ranges
}
