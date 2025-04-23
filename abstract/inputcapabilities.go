// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Input capabilities

package abstract

import "github.com/alexpevzner/mfp/util/generic"

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
