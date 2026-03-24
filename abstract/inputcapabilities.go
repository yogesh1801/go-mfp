// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Input capabilities

package abstract

import (
	"sort"

	"github.com/OpenPrinting/go-mfp/util/generic"
)

// InputCapabilities defines scanning capabilities of the
// particular [Input].
type InputCapabilities struct {
	// Input geometry
	MinWidth              Dimension // Min scan width
	MaxWidth              Dimension // Max scan width
	MinHeight             Dimension // Min scan height
	MaxHeight             Dimension // Max scan height
	MaxXOffset            Dimension // Max XOffset, 0 - unset
	MaxYOffset            Dimension // Max YOffset, 0 - unset
	MaxOpticalXResolution int       // DPI, 0 - unknown
	MaxOpticalYResolution int       // DPI, 0 - unknown
	RiskyLeftMargins      Dimension // Risky left margins, 0 - unknown
	RiskyRightMargins     Dimension // Risky right margins, 0 - unknown
	RiskyTopMargins       Dimension // Risky top margins, 0 - unknown
	RiskyBottomMargins    Dimension // Risky bottom margins, 0 - unknown

	// Scanning parameters
	Intents generic.Bitset[Intent] // Supported intents

	// Supported setting profiles
	Profiles []SettingsProfile // List of supported profiles
}

// Clone makes a shallow copy of the [InputCapabilities]
func (inpcaps *InputCapabilities) Clone() *InputCapabilities {
	clone := *inpcaps
	return &clone
}

// Resolutions returns all supported resolutions, taking
// in account all [InputCapabilities.Profiles].
//
// Returned resolutions are sorted in some deterministic but
// unspecified order, duplicates are removed.
func (inpcaps *InputCapabilities) Resolutions() []Resolution {
	var resolutions []Resolution

	// Gather all resolutions. Remove duplicates
	seen := make(map[Resolution]struct{})
	for _, prof := range inpcaps.Profiles {
		for _, res := range prof.Resolutions {
			if _, dup := seen[res]; !dup {
				resolutions = append(resolutions, res)
				seen[res] = struct{}{}
			}
		}
	}

	// Sort the response
	sort.Slice(resolutions, func(i, j int) bool {
		r1 := resolutions[i]
		r2 := resolutions[j]

		switch {
		case r1.XResolution < r2.XResolution:
			return true
		case r1.XResolution > r2.XResolution:
			return false
		}

		return r1.YResolution < r2.YResolution
	})

	return resolutions
}

// SquareResolutions is like [InputCapabilities.Resolutions],
// but it reports only square resolutions (i.e., resolutions
// with XResolution == YResolution)
func (inpcaps *InputCapabilities) SquareResolutions() []Resolution {
	resolutions := inpcaps.Resolutions()

	var i, o int
	for i = range resolutions {
		res := resolutions[i]
		if res.XResolution == res.YResolution {
			resolutions[o] = res
			o++
		}
	}

	resolutions = resolutions[:o]
	return resolutions
}
