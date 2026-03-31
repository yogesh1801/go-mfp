// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scanner capabilities

package abstract

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/OpenPrinting/go-mfp/internal/assert"
	"github.com/OpenPrinting/go-mfp/util/uuid"
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
	ShadowRange       Range // The lower, the darker
	SharpenRange      Range // Image sharpen
	ThresholdRange    Range // ColorModeBinary+BinaryRenderingThreshold

	// Input capabilities (nil if input not suppored)
	Platen     *InputCapabilities // InputPlaten capabilities
	ADFSimplex *InputCapabilities // InputADF+ADFModeSimplex
	ADFDuplex  *InputCapabilities // InputADF+ADFModeDuplex
}

// Clone makes a shallow copy of the [ScannerCapabilities].
func (scancaps *ScannerCapabilities) Clone() *ScannerCapabilities {
	clone := *scancaps
	return &clone
}

// FillRequest returns a fully populated copy of req, filling in any
// missing fields with reasonable defaults.
//
// The function accepts a partially populated ScannerRequest (which may
// be nil, treated as a zero-value request) and validates all fields
// against the ScannerCapabilities:
//   - Unsupported parameters are silently dropped from the result
//   - Supported parameters are validated if present, or populated with
//     compatible concrete defaults if missing
//
// The input request is never modified.
func (scancaps *ScannerCapabilities) FillRequest(
	req *ScannerRequest) (*ScannerRequest, error) {

	// Clone the request. We don't want to touch the input.
	if req != nil {
		req2 := *req
		req = &req2
	} else {
		req = &ScannerRequest{}
	}

	// Check Input
	if !req.Input.Valid() {
		err := ErrParam{ErrInvalidParam, "Input", req.Input}
		return nil, err
	}

	switch req.Input {
	case InputUnset:
		switch {
		case scancaps.Platen != nil:
			req.Input = InputPlaten
		case scancaps.ADFSimplex != nil || scancaps.ADFDuplex != nil:
			req.Input = InputADF
		default:
			err := errors.New("ScannerCapabilities: no inputs")
			return nil, err
		}

	case InputPlaten:
		if scancaps.Platen == nil {
			err := ErrParam{ErrUnsupportedParam, "Input", req.Input}
			return nil, err
		}

		req.ADFMode = ADFModeUnset

	case InputADF:
		if scancaps.ADFSimplex == nil && scancaps.ADFDuplex == nil {
			err := ErrParam{ErrUnsupportedParam, "Input", req.Input}
			return nil, err
		}
	}

	// Check ADFMode -- it depends on req.Input
	if !req.ADFMode.Valid() {
		err := ErrParam{ErrInvalidParam, "ADFMode", req.ADFMode}
		return nil, err
	}

	switch req.Input {
	case InputPlaten:
		req.ADFMode = ADFModeUnset

	case InputADF:
		switch req.ADFMode {
		case ADFModeUnset:
			// Prefer ADFSimplex with fallback to the ADFDuplex.
			if scancaps.ADFSimplex != nil {
				req.ADFMode = ADFModeSimplex
			} else {
				req.ADFMode = ADFModeDuplex
			}

		case ADFModeSimplex:
			if scancaps.ADFSimplex == nil {
				err := ErrParam{ErrUnsupportedParam,
					"ADFMode", req.ADFMode}
				return nil, err
			}

		case ADFModeDuplex:
			if scancaps.ADFDuplex == nil {
				err := ErrParam{ErrUnsupportedParam,
					"ADFMode", req.ADFMode}
				return nil, err
			}
		}
	}

	// Check DocumentFormat
	if req.DocumentFormat == "" {
		req.DocumentFormat = scancaps.preferredFormat()
		if req.DocumentFormat == "" {
			err := errors.New(
				"ScannerCapabilities.DocumentFormats: empty")
			return nil, err
		}
	} else {
		ok := false
		for _, fmt := range scancaps.DocumentFormats {
			if req.DocumentFormat == fmt {
				ok = true
				break
			}
		}
		if !ok {
			err := ErrParam{ErrUnsupportedParam,
				"DocumentFormat", req.DocumentFormat}
			return nil, err
		}
	}

	// Now fetch chosen InputCapabilities
	var input *InputCapabilities
	switch {
	case req.Input == InputPlaten:
		input = scancaps.Platen
	case req.Input == InputADF && req.ADFMode == ADFModeSimplex:
		input = scancaps.ADFSimplex
	case req.Input == InputADF && req.ADFMode == ADFModeDuplex:
		input = scancaps.ADFDuplex
	}

	assert.Must(input != nil)

	// Check Intent
	if !req.Intent.Valid() {
		err := ErrParam{ErrInvalidParam, "Intent", req.Intent}
		return nil, err
	}

	switch {
	case req.Intent == IntentUnset:
		req.Intent = input.defaultIntent()

	case input.Intents.IsEmpty():
		// Scanner doesn't support intents, so ignore req.Intent
		req.Intent = IntentUnset

	case !input.Intents.Contains(req.Intent):
		err := ErrParam{ErrUnsupportedParam, "Intent", req.Intent}
		return nil, err
	}

	// Check ColorMode and filter matching profiles
	if !req.ColorMode.Valid() {
		err := ErrParam{ErrInvalidParam, "ColorMode", req.ColorMode}
		return nil, err
	}

	profiles := slices.Clone(input.Profiles)
	if req.ColorMode == ColorModeUnset {
		// Prefer Color, then Mono, then Binary
		colorModes := summarizeColorModes(profiles)
		switch {
		case colorModes.Contains(ColorModeColor):
			req.ColorMode = ColorModeColor
		case colorModes.Contains(ColorModeMono):
			req.ColorMode = ColorModeMono
		case colorModes.Contains(ColorModeBinary):
			req.ColorMode = ColorModeBinary
		default:
			err := fmt.Errorf(
				"ScannerCapabilities.%s: empty color modes",
				req.Input)
			return nil, err
		}

		profiles = profilesByColorMode(profiles, req.ColorMode)
		assert.Must(len(profiles) > 0)
	} else {
		profiles = profilesByColorMode(profiles, req.ColorMode)
		if len(profiles) == 0 {
			err := ErrParam{ErrUnsupportedParam,
				"ColorMode", req.ColorMode}
			return nil, err
		}
	}

	// Check ColorDepts and BinaryRendering -- they depends on ColorMode
	if !req.ColorDepth.Valid() {
		err := ErrParam{ErrInvalidParam, "ColorDepth", req.ColorDepth}
		return nil, err
	}

	if !req.BinaryRendering.Valid() {
		err := ErrParam{ErrInvalidParam,
			"BinaryRendering", req.BinaryRendering}
		return nil, err
	}

	switch req.ColorMode {
	case ColorModeColor, ColorModeMono:
		req.BinaryRendering = BinaryRenderingUnset

		if req.ColorDepth == ColorDepthUnset {
			depths := summarizeColorDepths(profiles)
			switch {
			case depths.Contains(ColorDepth8):
				req.ColorDepth = ColorDepth8
			case depths.Contains(ColorDepth16):
				req.ColorDepth = ColorDepth16
			default:
				err := fmt.Errorf(
					"ScannerCapabilities.%s: empty depth",
					req.Input)
				return nil, err
			}

			profiles = profilesByColorDepth(profiles, req.ColorDepth)
			assert.Must(len(profiles) > 0)
		} else {
			profiles = profilesByColorDepth(profiles, req.ColorDepth)
			if len(profiles) == 0 {
				err := ErrParam{ErrUnsupportedParam,
					"ColorDepth", req.ColorDepth}
				return nil, err
			}
		}

	case ColorModeBinary:
		req.ColorDepth = ColorDepthUnset

		if req.BinaryRendering == BinaryRenderingUnset {
			binrend := summarizeBinaryRenderings(profiles)
			switch {
			case binrend.Contains(BinaryRenderingThreshold):
				req.BinaryRendering = BinaryRenderingThreshold
			case binrend.Contains(BinaryRenderingHalftone):
				req.BinaryRendering = BinaryRenderingHalftone
			default:
				err := fmt.Errorf(
					"ScannerCapabilities.%s: empty binary rendering",
					req.Input)
				return nil, err
			}

			profiles = profilesByBinaryRendering(profiles,
				req.BinaryRendering)
			assert.Must(len(profiles) > 0)
		} else {
			profiles = profilesByBinaryRendering(profiles,
				req.BinaryRendering)
			if len(profiles) == 0 {
				err := ErrParam{ErrUnsupportedParam,
					"BinaryRendering", req.BinaryRendering}
				return nil, err
			}
		}
	}

	// Check Threshold
	if req.BinaryRendering == BinaryRenderingThreshold {
		var err error
		req.Threshold, err = scancaps.ThresholdRange.resolve(
			"Threshold", req.Threshold)
		if err != nil {
			return nil, err
		}
	} else {
		req.Threshold = nil
	}

	// Check CCDChannel
	if !req.CCDChannel.Valid() {
		err := ErrParam{ErrInvalidParam, "CCDChannel", req.CCDChannel}
		return nil, err
	}

	switch req.ColorMode {
	case ColorModeBinary, ColorModeMono:
		if req.CCDChannel != CCDChannelUnset {
			profiles = profilesByCCDChannel(profiles,
				req.CCDChannel)
			if len(profiles) == 0 {
				err := ErrParam{ErrUnsupportedParam,
					"CCDChannel", req.CCDChannel}
				return nil, err
			}
		}

	default:
		req.CCDChannel = CCDChannelUnset
	}

	// Check Region
	switch {
	case req.Region.IsZero():
		req.Region = Region{
			Width:  input.MaxWidth,
			Height: input.MaxHeight,
		}

	case !req.Region.Valid():
		err := ErrParam{ErrInvalidParam, "Region", req.Region}
		return nil, err

	case !req.Region.FitsCapabilities(input):
		err := ErrParam{ErrUnsupportedParam, "Region", req.Region}
		return nil, err
	}

	// Check Resolution
	switch {
	case req.Resolution.IsZero():
		req.Resolution = defaultResolution(profiles)
		if req.Resolution.IsZero() {
			err := fmt.Errorf(
				"ScannerCapabilities.%s: empty resolutions",
				req.Input)
			return nil, err
		}

	case !req.Resolution.Valid():
		err := ErrParam{ErrInvalidParam, "Resolution", req.Resolution}
		return nil, err

	default:
		ok := false
		for i := 0; i < len(profiles) && !ok; i++ {
			prof := profiles[i]
			ok = prof.AllowsResolution(req.Resolution)
		}

		if !ok {
			return nil, ErrParam{ErrUnsupportedParam,
				"Resolution", req.Resolution}
		}
	}

	// Check image processing parameters.
	var err error
	req.Brightness, err = scancaps.BrightnessRange.resolve(
		"Brightness", req.Brightness)
	if err == nil {
		req.Contrast, err = scancaps.ContrastRange.resolve(
			"Contrast", req.Contrast)
	}
	if err == nil {
		req.Gamma, err = scancaps.GammaRange.resolve(
			"Gamma", req.Gamma)
	}
	if err == nil {
		req.Highlight, err = scancaps.HighlightRange.resolve(
			"Highlight", req.Highlight)
	}
	if err == nil {
		req.NoiseRemoval, err = scancaps.NoiseRemovalRange.resolve(
			"NoiseRemoval", req.NoiseRemoval)
	}
	if err == nil {
		req.Shadow, err = scancaps.ShadowRange.resolve(
			"Shadow", req.Shadow)
	}
	if err == nil {
		req.Sharpen, err = scancaps.SharpenRange.resolve(
			"Sharpen", req.Sharpen)
	}
	if err == nil {
		req.Compression, err = scancaps.CompressionRange.resolve(
			"Compression", req.Compression)
	}

	if err != nil {
		return nil, err
	}

	// Fixup final parameters
	prof := profiles[0]

	switch req.ColorMode {
	case ColorModeBinary, ColorModeMono:
		if req.CCDChannel == CCDChannelUnset &&
			!prof.CCDChannels.IsEmpty() {
			switch {
			case prof.CCDChannels.Contains(CCDChannelGrayCcd):
				req.CCDChannel = CCDChannelGrayCcd
			case prof.CCDChannels.Contains(CCDChannelGrayCcdEmulated):
				req.CCDChannel = CCDChannelGrayCcdEmulated
			case prof.CCDChannels.Contains(CCDChannelNTSC):
				req.CCDChannel = CCDChannelNTSC

			case prof.CCDChannels.Contains(CCDChannelRed):
				req.CCDChannel = CCDChannelRed
			case prof.CCDChannels.Contains(CCDChannelGreen):
				req.CCDChannel = CCDChannelGreen
			case prof.CCDChannels.Contains(CCDChannelBlue):
				req.CCDChannel = CCDChannelBlue
			}

			assert.Must(req.CCDChannel != CCDChannelUnset)
		}
	}

	return req, nil
}

// DefaultRequest returns default [ScannerRequest] that
// matches the [ScannerCapabilities].
//
// Note, it may return nil, if ScannerCapabilities are ugly
// broken and doesn't allow any valid request.
func (scancaps *ScannerCapabilities) DefaultRequest() *ScannerRequest {
	req, _ := scancaps.FillRequest(nil)
	return req
}

// ValidateRequest validates the [ScannerRequest] against the
// [ScannerCapabilities].
func (scancaps *ScannerCapabilities) ValidateRequest(req *ScannerRequest) error {
	_, err := scancaps.FillRequest(req)
	return err
}

// defaultFormat returns the default document format,
// allowed by the ScannerCapabilities.
//
// It may return "", if ScannerCapabilities doesn't allow any
// document format at all.
func (scancaps *ScannerCapabilities) preferredFormat() string {
	// Check what we have
	pdf := ""
	png := ""
	jpg := ""

	for _, fmt := range scancaps.DocumentFormats {
		switch strings.ToLower(fmt) {
		case "application/pdf":
			if pdf == "" {
				pdf = fmt
			}
		case "image/png":
			if png == "" {
				png = fmt
			}
		case "image/jpeg":
			if jpg == "" {
				jpg = fmt
			}
		}
	}

	// Choose from known formats
	switch {
	case pdf != "":
		return pdf
	case png != "":
		return png
	case jpg != "":
		return jpg
	}

	// Fallback to the first available format, if any
	if len(scancaps.DocumentFormats) > 0 {
		return scancaps.DocumentFormats[0]
	}

	return ""
}
