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
	// The following parameters are per-source:
	BrightnessRange   Range // Brightness
	ContrastRange     Range // Contrast
	GammaRange        Range // Gamma (y=x^(1/g)
	HighlightRange    Range // Image Highlight
	NoiseRemovalRange Range // Noise removal level
	ShadowRange       Range // The lower, the darger
	SharpenRange      Range // Image sharpen
	ThresholdRange    Range // ColorModeBinary+BinaryRenderingThreshold

	// Supported setting profiles
	Profiles []SettingsProfile // List of supported profiles
}
