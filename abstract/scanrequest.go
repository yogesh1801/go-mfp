// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scan request

package abstract

import (
	"github.com/alexpevzner/mfp/util/generic"
	"github.com/alexpevzner/mfp/util/optional"
)

// ScannerRequest specified scan request parameters
type ScannerRequest struct {
	// General parameters
	//
	// All parameters are optional. Use zero value to to indicate
	// that parameter is missed.
	Input           Input           // Input source (ADF/Platen etc)
	ADFMode         ADFMode         // For InputADF: Duplex/Simplex
	ColorMode       ColorMode       // Color mode (mono/color etc)
	Depth           Depth           // Image depth (8-bit/16-bit etc)
	BinaryRendering BinaryRendering // For 1-bit B&W (halftone/threshold
	CCDChannel      CCDChannel      // CCD channel to use
	DocumentFormat  string          // Requested document format
	Region          Region          // Scan region
	Resolution      Resolution      // Scanner resolution
	Intent          Intent          // Scan intent hint

	// Image processing parameters.
	//
	// As zero value is the legal value of these parameters,
	// we have to use optional.Val[int] to distinguish between
	// missed parameter and 0.
	Brightness   optional.Val[int] // Brightness
	Contrast     optional.Val[int] // Contrast
	Gamma        optional.Val[int] // Gamma (y=x^(1/g)
	Highlight    optional.Val[int] // Image Highlight
	NoiseRemoval optional.Val[int] // Noise removal level
	Shadow       optional.Val[int] // The lower, the darger
	Sharpen      optional.Val[int] // Image sharpen
	Threshold    optional.Val[int] // ColorModeBinary+BinaryRenderingThreshold
	Compression  optional.Val[int] // Lower num, better image
}

// Validate checks request validity against the [ScannerCapabilities]
// and reports found error, if any.
func (req *ScannerRequest) Validate(scancaps *ScannerCapabilities) error {
	// Gather overall scanner parameters
	var inputs generic.Bitset[Input]
	var adfmodes generic.Bitset[ADFMode]
	var intents generic.Bitset[Intent]
	var colorModes generic.Bitset[ColorMode]
	var depths generic.Bitset[Depth]
	var binrend generic.Bitset[BinaryRendering]
	var ccdChannels generic.Bitset[CCDChannel]

	if scancaps.Platen != nil {
		inputs.Add(InputPlaten)
	}

	if scancaps.ADFSimplex != nil || scancaps.ADFDuplex != nil {
		inputs.Add(InputADF)
		if scancaps.ADFSimplex != nil {
			adfmodes.Add(ADFModeSimplex)
		}
		if scancaps.ADFDuplex != nil {
			adfmodes.Add(ADFModeDuplex)
		}
	}

	for _, inpcaps := range []*InputCapabilities{
		scancaps.Platen, scancaps.ADFSimplex, scancaps.ADFDuplex} {
		if inpcaps == nil {
			continue
		}

		intents = intents.Union(inpcaps.Intents)
		for _, prof := range inpcaps.Profiles {
			colorModes = colorModes.Union(prof.ColorModes)
			depths = depths.Union(prof.Depths)
			binrend = binrend.Union(prof.BinaryRenderings)
			ccdChannels = ccdChannels.Union(prof.CCDChannels)
		}
	}

	// Check Input
	switch {
	case req.Input == InputUnset:
	case req.Input < 0 || req.Input >= inputMax:
		return ErrInvalidInput
	case !inputs.Contains(req.Input):
		return ErrUnsupportedInput
	}

	// Check ADFMode
	switch {
	case req.Input != InputADF:
	case req.ADFMode == ADFModeUnset:
	case req.ADFMode < 0 || req.ADFMode >= adfModeMax:
		return ErrInvalidADFMode
	case !adfmodes.Contains(req.ADFMode):
		return ErrUnsupportedADFMode
	}

	return nil
}
