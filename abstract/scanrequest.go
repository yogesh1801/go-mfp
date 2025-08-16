// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scan request

package abstract

import (
	"bytes"
	"fmt"

	"github.com/OpenPrinting/go-mfp/util/generic"
	"github.com/OpenPrinting/go-mfp/util/optional"
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
	ColorDepth      ColorDepth      // Image depth (8-bit/16-bit etc)
	BinaryRendering BinaryRendering // For 1-bit B&W (halftone/threshold)
	CCDChannel      CCDChannel      // CCD channel to use (mono/gray only)
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

// MarshalLog formats [ScannerRequest] for logging.
// It implements the [log.Marshaler] interface.
func (req *ScannerRequest) MarshalLog() []byte {
	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "Input:           %s\n", req.Input)
	if req.ADFMode != ADFModeUnset {
		fmt.Fprintf(buf, "ADFMode:         %s\n", req.ADFMode)
	}
	if req.ColorMode != ColorModeUnset {
		fmt.Fprintf(buf, "ColorMode:       %s\n", req.ColorMode)
	}
	if req.ColorDepth != ColorDepthUnset {
		fmt.Fprintf(buf, "ColorDepth:      %s\n", req.ColorDepth)
	}
	if req.BinaryRendering != BinaryRenderingUnset {
		fmt.Fprintf(buf, "BinaryRendering: %s\n", req.BinaryRendering)
	}
	if req.CCDChannel != CCDChannelUnset {
		fmt.Fprintf(buf, "CCDChannel:      %s\n", req.CCDChannel)
	}
	fmt.Fprintf(buf, "DocumentFormat:  %q\n", req.DocumentFormat)
	if !req.Region.IsZero() {
		fmt.Fprintf(buf, "Region:          %s\n", req.Region)
	}
	if !req.Resolution.IsZero() {
		fmt.Fprintf(buf, "Resolution:      %s\n", req.Resolution)
	}
	fmt.Fprintf(buf, "Intent:          %s\n", req.Intent)
	if req.Brightness != nil {
		fmt.Fprintf(buf, "Brightness:      %d\n", *req.Brightness)
	}
	if req.Contrast != nil {
		fmt.Fprintf(buf, "Contrast:        %d\n", *req.Contrast)
	}
	if req.Gamma != nil {
		fmt.Fprintf(buf, "Gamma:           %d\n", *req.Gamma)
	}
	if req.Highlight != nil {
		fmt.Fprintf(buf, "Highlight:       %d\n", *req.Highlight)
	}
	if req.NoiseRemoval != nil {
		fmt.Fprintf(buf, "NoiseRemoval:    %d\n", *req.NoiseRemoval)
	}
	if req.Shadow != nil {
		fmt.Fprintf(buf, "Shadow:          %d\n", *req.Shadow)
	}
	if req.Sharpen != nil {
		fmt.Fprintf(buf, "Sharpen:         %d\n", *req.Sharpen)
	}
	if req.Threshold != nil {
		fmt.Fprintf(buf, "Threshold:       %d\n", *req.Threshold)
	}
	if req.Compression != nil {
		fmt.Fprintf(buf, "Compression:     %d\n", *req.Compression)
	}

	return buf.Bytes()
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

	// Check DocumentFormat
	if req.DocumentFormat != "" {
		ok := false
		for _, fmt := range scancaps.DocumentFormats {
			if req.DocumentFormat == fmt {
				ok = true
				break
			}
		}
		if !ok {
			return ErrParam{ErrInvalidParam,
				"DocumentFormat", req.DocumentFormat}
		}
	}

	// Gather overall scanner parameters
	var intents generic.Bitset[Intent]
	var colorModes generic.Bitset[ColorMode]
	var depths generic.Bitset[ColorDepth]
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
		return ErrParam{ErrInvalidParam, "ColorMode", req.ColorMode}
	case !colorModes.Contains(req.ColorMode):
		return ErrParam{ErrUnsupportedParam,
			"ColorMode", req.ColorMode}
	}

	switch req.ColorMode {
	case ColorModeUnset, ColorModeBinary:
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
	}

	switch req.ColorMode {
	case ColorModeUnset, ColorModeMono, ColorModeColor:
		switch {
		case req.ColorDepth == ColorDepthUnset:
		case req.ColorDepth < 0 || req.ColorDepth >= colorDepthMax:
			return ErrParam{ErrInvalidParam,
				"ColorDepth", req.ColorDepth}
		case !depths.Contains(req.ColorDepth):
			return ErrParam{ErrUnsupportedParam,
				"ColorDepth", req.ColorDepth}
		}
	}

	// Check CCDChannel
	switch req.ColorMode {
	case ColorModeUnset, ColorModeBinary, ColorModeMono:
		switch {
		case req.CCDChannel == CCDChannelUnset:
		case req.CCDChannel < 0 || req.CCDChannel >= ccdChannelMax:
			return ErrParam{ErrInvalidParam, "CCDChannel",
				req.CCDChannel}
		case !ccdChannels.Contains(req.CCDChannel):
			return ErrParam{ErrUnsupportedParam,
				"CCDChannel", req.CCDChannel}
		}
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
		for i := 0; i < len(inputs) && !ok; i++ {
			inp := inputs[i]
			for j := 0; j < len(inp.Profiles) && !ok; j++ {
				prof := inp.Profiles[j]
				ok = prof.AllowsColorMode(req.ColorMode,
					req.ColorDepth, req.BinaryRendering)
				ok = ok && prof.AllowsCCDChannel(req.CCDChannel)
				ok = ok && prof.AllowsResolution(req.Resolution)
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

	return err
}
