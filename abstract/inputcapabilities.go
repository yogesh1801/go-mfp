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
	MinWidth   Dimension // Min scan width
	MaxWidth   Dimension // Max scan width
	MinHeight  Dimension // Min scan height
	MaxHeight  Dimension // Max scan height
	MaxXOffset Dimension // Max XOffset
	MaxYOffset Dimension // Max YOffset

	// Scanning parameters
	Intents generic.Bitset[Intent] // Supported intents

	// Supported setting profiles
	Profiles []SettingsProfile // List of supported profiles
}
