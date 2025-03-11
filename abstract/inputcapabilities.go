// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Input capabilities

package abstract

// InputCapabilities defines scanning capabilities of the
// particular [Input].
type InputCapabilities struct {
	MinWidth   int // Min scan width
	MaxWidth   int // Max scan width
	MinHeight  int // Min scan height
	MaxHeight  int // Max scan height
	MaxXOffset int // Max XOffset
	MaxYOffset int // Max YOffset

	// Supported setting profiles
	Profiles []SettingsProfile // List of supported profiles
}
