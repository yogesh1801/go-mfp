// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scanner capabilities

package abstract

import (
	"github.com/alexpevzner/mfp/util/generic"
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

	// Overall capabilities
	InputSupported     generic.Bitset[Input]     // Supported inputs
	ADFModesSupported  generic.Bitset[ADFMode]   // Supported ADF modes
	ColorModeSupported generic.Bitset[ColorMode] // Supported color modes
	DepthSupported     generic.Bitset[Depth]     // Supported image depths

	// Common image processing parameters
	DocumentFormats  []string // Supported output formats
	CompressionRange Range    // Lower num, better image

	// Input capabilities (nil if input not suppored)
	Platen     *InputCapabilities // Capabilities of InputPlaten
	ADFSimplex *InputCapabilities // InputADF+ADFModeSimplex
	ADFDuplex  *InputCapabilities // InputADF+ADFModeDuplex
}
