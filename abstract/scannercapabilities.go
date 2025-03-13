// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scanner capabilities

package abstract

import (
	"github.com/alexpevzner/mfp/util/uuid"
)

// ScannerCapabilities defines the scanner capabilities.
type ScannerCapabilities struct {
	// General information
	UUID         uuid.UUID // Device UUID
	MakeAndModel string    // Device make and model
	SerialNumber string    // Device-unique serial number
	Manufacturer string    // Device manufacturer
	AdminURI     string    // Configuration mage URL
	IconURI      string    // Device icon URL

	// Common image processing parameters
	DocumentFormats  []string // Supported output formats
	CompressionRange Range    // Lower num, better image
	ADFCapacity      int      // 0 if unknown or no ADF

	// Exposure control parameters
	BrightnessRange   Range // Brightness
	ContrastRange     Range // Contrast
	GammaRange        Range // Gamma (y=x^(1/g)
	HighlightRange    Range // Image Highlight
	NoiseRemovalRange Range // Noise removal level
	ShadowRange       Range // The lower, the darger
	SharpenRange      Range // Image sharpen
	ThresholdRange    Range // ColorModeBinary+BinaryRenderingThreshold

	// Input capabilities (nil if input not suppored)
	Platen     *InputCapabilities // InputPlaten capabilities
	ADFSimplex *InputCapabilities // InputADF+ADFModeSimplex
	ADFDuplex  *InputCapabilities // InputADF+ADFModeDuplex
}
