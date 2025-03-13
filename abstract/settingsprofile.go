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
type SettingsProfile struct {
	ColorModes       generic.Bitset[ColorMode]       // Allowed color modes
	Depths           generic.Bitset[Depth]           // Allowed depths
	BinaryRenderings generic.Bitset[BinaryRendering] // For 1-bit B&W
	CCDChannels      generic.Bitset[CCDChannel]      // Allowed CCD channel
	Resolutions      []Resolution                    // Allowed resolutions
	ResolutionRange  ResolutionRange                 // Zero if unset
}
