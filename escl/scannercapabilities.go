// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scanner capabilities

package escl

import (
	"github.com/alexpevzner/mfp/optional"
	"github.com/alexpevzner/mfp/uuid"
)

// ScannerCapabilities defines the scanner capabilities.
//
// eSCL Technical Specification, 8.1.4.
type ScannerCapabilities struct {
	// General options
	Version         Version                 // eSCL protocol version
	MakeAndModel    optional.Val[string]    // Device make and model
	SerialNumber    optional.Val[string]    // Device-unique serial number
	Manufacturer    optional.Val[string]    // Device manufacturer
	UUID            optional.Val[uuid.UUID] // Device UUID
	AdminURI        optional.Val[string]    // Configuration mage URL
	IconURI         optional.Val[string]    // Device icon URL
	SettingProfiles []SettingProfile        // Common settings profs

	// Inputs capabilities
	Platen optional.Val[Platen] // Platen capabilities
	Camera optional.Val[Camera] // Camera capabilities
	ADF    optional.Val[ADF]    // ADF capabilities

	// Image transform ranges
	BrightnessSupport        optional.Val[Range] // Brightness
	CompressionFactorSupport optional.Val[Range] // Lower num, better image
	ContrastSupport          optional.Val[Range] // Contrast
	GammaSupport             optional.Val[Range] // Gamma (y = x^(1/g))
	HighlightSupport         optional.Val[Range] // Image Highlight
	NoiseRemovalSupport      optional.Val[Range] // Noise removal level
	ShadowSupport            optional.Val[Range] // The lower, the darger
	SharpenSupport           optional.Val[Range] // Image sharpen
	ThresholdSupport         optional.Val[Range] // For BlackAndWhite1

	// Automatic detection and removal of the blank pages
	BlankPageDetection           optional.Val[bool] // Detection supported
	BlankPageDetectionAndRemoval optional.Val[bool] // Auto-remove supported
}
