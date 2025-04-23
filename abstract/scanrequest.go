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
	// Check Input and ADFMode. Choose relevant input capabilities.
	var inputs []*InputCapabilities

	switch req.Input {
	case InputUnset:
		if scancaps.Platen != nil {
			inputs = append(inputs, scancaps.Platen)
		}
		if scancaps.ADFSimplex != nil {
			inputs = append(inputs, scancaps.ADFSimplex)
		}
		if scancaps.ADFDuplex != nil {
			inputs = append(inputs, scancaps.ADFDuplex)
		}
		if len(inputs) == 0 {
			return ErrParam{ErrUnsupportedParam, "Input", req.Input}
		}

	case InputPlaten:
		if scancaps.Platen == nil {
			return ErrParam{ErrUnsupportedParam, "Input", req.Input}
		}
		inputs = []*InputCapabilities{scancaps.Platen}

	case InputADF:
		if scancaps.ADFSimplex == nil && scancaps.ADFDuplex == nil {
			return ErrParam{ErrUnsupportedParam, "Input", req.Input}
		}

		switch req.ADFMode {
		case ADFModeUnset:
			// If ADF mode is not set, we prefer ADFSimplex
			// with fallback to the ADFDuplex.
			if scancaps.ADFSimplex != nil {
				inputs = append(inputs, scancaps.ADFSimplex)
			} else {
				inputs = append(inputs, scancaps.ADFDuplex)
			}

		case ADFModeSimplex:
			if scancaps.ADFSimplex == nil {
				return ErrParam{ErrUnsupportedParam,
					"ADFMode", req.ADFMode}
			}
			inputs = append(inputs, scancaps.ADFSimplex)

		case ADFModeDuplex:
			if scancaps.ADFDuplex == nil {
				return ErrParam{ErrUnsupportedParam,
					"ADFMode", req.ADFMode}
			}
			inputs = append(inputs, scancaps.ADFDuplex)

		default:
			return ErrParam{ErrInvalidParam, "ADFMode", req.ADFMode}
		}

	default:
		return ErrParam{ErrInvalidParam, "Input", req.Input}
	}

	// Gather overall scanner parameters
	var intents generic.Bitset[Intent]
	var colorModes generic.Bitset[ColorMode]
	var depths generic.Bitset[Depth]
	var binrend generic.Bitset[BinaryRendering]
	var ccdChannels generic.Bitset[CCDChannel]

	for _, inp := range inputs {
		intents = intents.Union(inp.Intents)
		for _, prof := range inp.Profiles {
			colorModes = colorModes.Union(prof.ColorModes)
			depths = depths.Union(prof.Depths)
			binrend = binrend.Union(prof.BinaryRenderings)
			ccdChannels = ccdChannels.Union(prof.CCDChannels)
		}
	}

	// Check ColorMode, Depth, BinaryRendering and Threshold
	switch {
	case req.ColorMode == ColorModeUnset:
	case req.ColorMode < 0 || req.ColorMode >= colorModeMax:
		return ErrParam{ErrInvalidParam, "ColorMode,", req.ColorMode}
	case !colorModes.Contains(req.ColorMode):
		return ErrParam{ErrUnsupportedParam,
			"ColorMode,", req.ColorMode}
	}

	switch req.ColorMode {
	case ColorModeBinary:
		switch {
		case req.BinaryRendering == BinaryRenderingUnset:
		case req.BinaryRendering < 0 || req.BinaryRendering >= binaryRenderingMax:
			return ErrParam{ErrInvalidParam,
				"BinaryRendering", req.BinaryRendering}
		case !binrend.Contains(req.BinaryRendering):
			return ErrParam{ErrUnsupportedParam,
				"BinaryRendering", req.BinaryRendering}
		}

		err := scancaps.ThresholdRange.validate(
			"Threshold", req.Threshold)
		if err != nil {
			return err
		}

	case ColorModeMono, ColorModeColor:
		switch {
		case req.Depth == DepthUnset:
		case req.Depth < 0 || req.Depth >= depthMax:
			return ErrParam{ErrInvalidParam, "Depth", req.Depth}
		case !depths.Contains(req.Depth):
			return ErrParam{ErrUnsupportedParam, "Depth", req.Depth}
		}
	}

	// Check CCDChannel
	switch {
	case req.CCDChannel == CCDChannelUnset:
	case req.CCDChannel < 0 || req.CCDChannel >= ccdChannelMax:
		return ErrParam{ErrInvalidParam, "CCDChannel", req.CCDChannel}
	case !ccdChannels.Contains(req.CCDChannel):
		return ErrParam{ErrUnsupportedParam,
			"CCDChannel", req.CCDChannel}
	}

	// Check Intent
	switch {
	case req.Intent == IntentUnset:
	case req.Intent < 0 || req.Intent >= intentMax:
		return ErrParam{ErrInvalidParam, "Intent", req.Intent}
	case !intents.Contains(req.Intent):
		return ErrParam{ErrUnsupportedParam, "Intent", req.Intent}
	}

	// Check Region
	if !req.Region.IsZero() {
		if !req.Region.Valid() {
			return ErrParam{ErrInvalidParam, "Region", req.Region}
		}

		ok := false
		for _, inp := range inputs {
			if req.Region.FitsCapabilities(inp) {
				ok = true
				break
			}
		}

		if !ok {
			return ErrParam{ErrUnsupportedParam,
				"Region", req.Region}
		}
	}

	// Check Resolution
	if !req.Resolution.IsZero() {
		if !req.Resolution.Valid() {
			return ErrParam{ErrInvalidParam,
				"Resolution", req.Resolution}
		}

		ok := false
	loop:
		for _, inp := range inputs {
			for _, prof := range inp.Profiles {
				if !prof.AllowsColorMode(req.ColorMode,
					req.Depth, req.BinaryRendering) {
					continue
				}

				if !prof.AllowsCCDChannel(req.CCDChannel) {
					continue
				}

				if !prof.AllowsResolution(
					req.Resolution) {
					continue
				}

				ok = true
				break loop
			}
		}

		if !ok {
			return ErrParam{ErrUnsupportedParam,
				"Resolution", req.Resolution}
		}
	}

	// Check image processing parameters.
	err := scancaps.BrightnessRange.validate("Brightness", req.Brightness)
	if err == nil {
		err = scancaps.ContrastRange.validate("Contrast", req.Contrast)
	}
	if err == nil {
		err = scancaps.GammaRange.validate("Gamma", req.Gamma)
	}
	if err == nil {
		err = scancaps.HighlightRange.validate(
			"Highlight", req.Highlight)
	}
	if err == nil {
		err = scancaps.NoiseRemovalRange.validate(
			"NoiseRemoval", req.NoiseRemoval)
	}
	if err == nil {
		err = scancaps.ShadowRange.validate("Shadow", req.Shadow)
	}
	if err == nil {
		err = scancaps.SharpenRange.validate("Sharpen", req.Sharpen)
	}
	if err == nil {
		err = scancaps.CompressionRange.validate(
			"Compression", req.Compression)
	}

	if err != nil {
		return err
	}

	return nil
}
