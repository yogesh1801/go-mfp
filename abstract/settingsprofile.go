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

// AllowsColorMode reports if the [SettingsProfile] compatible
// with the given combination of the [ColorMode], [Depth],
// and [BinaryRendering] combination.
func (prof SettingsProfile) AllowsColorMode(
	cm ColorMode, d Depth, r BinaryRendering) bool {
	switch {
	case cm == ColorModeUnset:
		return true
	case !prof.ColorModes.Contains(cm):
		return false
	case cm == ColorModeBinary:
		return r == BinaryRenderingUnset ||
			prof.BinaryRenderings.Contains(r)
	}

	return d == DepthUnset || prof.Depths.Contains(d)
}

// AllowsCCDChannel reports if the [SettingsProfile] compatible
// with the given [CCDChannel].
func (prof SettingsProfile) AllowsCCDChannel(ccd CCDChannel) bool {
	if ccd == CCDChannelUnset {
		return true
	}

	return prof.CCDChannels.Contains(ccd)
}

// AllowsResolution reports if the [SettingsProfile] compatible
// with the given [Resolution].
func (prof SettingsProfile) AllowsResolution(res Resolution) bool {
	if res.IsZero() {
		return true
	}

	for _, supported := range prof.Resolutions {
		if res == supported {
			return true
		}
	}

	// FIXME -- chech ResolutionRange

	return false
}
