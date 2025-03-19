// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Conversions from abstract.Scanner to eSCL data structures

package escl

import (
	"github.com/alexpevzner/mfp/abstract"
	"github.com/alexpevzner/mfp/util/generic"
	"github.com/alexpevzner/mfp/util/optional"
)

// fromAbstractScanSettings translates [abstract.ScannerRequest
// into the [ScanSettings].
//
// The version parameters affects how some fields are converted.
func fromAbstractScanSettings(
	version Version,
	absreq *abstract.ScannerRequest) ScanSettings {

	ss := ScanSettings{
		Version:           version,
		Brightness:        fromAbstractOptionalInt(absreq.Brightness),
		Contrast:          fromAbstractOptionalInt(absreq.Contrast),
		Gamma:             fromAbstractOptionalInt(absreq.Gamma),
		Highlight:         fromAbstractOptionalInt(absreq.Highlight),
		NoiseRemoval:      fromAbstractOptionalInt(absreq.NoiseRemoval),
		Shadow:            fromAbstractOptionalInt(absreq.Shadow),
		Sharpen:           fromAbstractOptionalInt(absreq.Sharpen),
		CompressionFactor: fromAbstractOptionalInt(absreq.Compression),
	}

	// Translate intent
	intent := fromAbstractIntent(absreq.Intent)
	if intent != UnknownIntent {
		ss.Intent = optional.New(intent)
	}

	// Translate ScanRegions
	if !absreq.Region.IsZero() {
		reg := ScanRegion{
			XOffset:            absreq.Region.XOffset.Dots(300),
			YOffset:            absreq.Region.YOffset.Dots(300),
			Width:              absreq.Region.Width.Dots(300),
			Height:             absreq.Region.Height.Dots(300),
			ContentRegionUnits: ThreeHundredthsOfInches,
		}
		ss.ScanRegions = []ScanRegion{reg}
	}

	// Translate DocumentFormat/DocumentFormatExt
	if absreq.DocumentFormat != "" {
		ss.DocumentFormat = optional.New(absreq.DocumentFormat)
		if version >= MakeVersion(2, 1) {
			// eSCL 2.1+
			ss.DocumentFormatExt = ss.DocumentFormat
		}
	}

	// Translate InputSource and Duplex
	switch absreq.Input {
	case abstract.InputPlaten:
		ss.InputSource = optional.New(InputPlaten)
	case abstract.InputADF:
		ss.InputSource = optional.New(InputFeeder)
		if absreq.ADFMode == abstract.ADFModeDuplex {
			ss.Duplex = optional.New(true)
		}
	}

	// Translate XResolution and YResolution
	if !absreq.Resolution.IsZero() {
		ss.XResolution = optional.New(absreq.Resolution.XResolution)
		ss.YResolution = optional.New(absreq.Resolution.YResolution)
	}

	// Translate ColorMode, BinaryRendering and Threshold
	switch absreq.ColorMode {
	case abstract.ColorModeBinary:
		ss.ColorMode = optional.New(BlackAndWhite1)
		switch absreq.BinaryRendering {
		case abstract.BinaryRenderingHalftone:
			ss.BinaryRendering = optional.New(Halftone)
		case abstract.BinaryRenderingThreshold:
			ss.BinaryRendering = optional.New(Threshold)
			ss.Threshold = fromAbstractOptionalInt(absreq.Threshold)
		}
	case abstract.ColorModeMono:
		if absreq.Depth == abstract.Depth16 {
			ss.ColorMode = optional.New(Grayscale16)
		} else {
			ss.ColorMode = optional.New(Grayscale8)
		}
	case abstract.ColorModeColor:
		if absreq.Depth == abstract.Depth16 {
			ss.ColorMode = optional.New(RGB48)
		} else {
			ss.ColorMode = optional.New(RGB24)
		}
	}

	// Translate CCDChannel
	ccd := fromAbstractCCDChannel(absreq.CCDChannel)
	if ccd != UnknownCCDChannel {
		ss.CCDChannel = optional.New(ccd)
	}

	return ss
}

// fromAbstractScannerCapabilities translates [abstract.ScannerCapabilities]
// into the [ScannerCapabilities].
//
// The version parameters affects how some fields are converted.
func fromAbstractScannerCapabilities(
	version Version,
	abscaps *abstract.ScannerCapabilities) ScannerCapabilities {

	scancaps := ScannerCapabilities{
		Version: version,
		UUID:    optional.New(abscaps.UUID),
	}

	// Translate general parameters
	if abscaps.MakeAndModel != "" {
		scancaps.MakeAndModel = optional.New(abscaps.MakeAndModel)
	}
	if abscaps.SerialNumber != "" {
		scancaps.SerialNumber = optional.New(abscaps.SerialNumber)
	}
	if abscaps.Manufacturer != "" {
		scancaps.Manufacturer = optional.New(abscaps.Manufacturer)
	}
	if abscaps.AdminURI != "" {
		scancaps.AdminURI = optional.New(abscaps.AdminURI)
	}
	if abscaps.IconURI != "" {
		scancaps.IconURI = optional.New(abscaps.IconURI)
	}

	// Translate image transform parameters
	scancaps.BrightnessSupport = fromAbstractOptionalRange(
		abscaps.BrightnessRange)
	scancaps.ContrastSupport = fromAbstractOptionalRange(
		abscaps.ContrastRange)
	scancaps.GammaSupport = fromAbstractOptionalRange(
		abscaps.GammaRange)
	scancaps.HighlightSupport = fromAbstractOptionalRange(
		abscaps.HighlightRange)
	scancaps.NoiseRemovalSupport = fromAbstractOptionalRange(
		abscaps.NoiseRemovalRange)
	scancaps.ShadowSupport = fromAbstractOptionalRange(
		abscaps.ShadowRange)
	scancaps.SharpenSupport = fromAbstractOptionalRange(
		abscaps.SharpenRange)
	scancaps.ThresholdSupport = fromAbstractOptionalRange(
		abscaps.ThresholdRange)
	scancaps.CompressionFactorSupport = fromAbstractOptionalRange(
		abscaps.CompressionRange)

	// Translate input capabilities
	if abscaps.Platen != nil {
		caps := fromAbstractInputSourceCaps(version,
			abscaps.DocumentFormats, abscaps.Platen)
		scancaps.Platen = optional.New(Platen{optional.New(caps)})
	}

	if abscaps.ADFSimplex != nil || abscaps.ADFDuplex != nil {
		adf := ADF{
			FeederCapacity: fromAbstractOptionalInt(
				abscaps.ADFCapacity),
		}

		if abscaps.ADFSimplex != nil {
			caps := fromAbstractInputSourceCaps(version,
				abscaps.DocumentFormats, abscaps.ADFSimplex)
			adf.ADFSimplexInputCaps = optional.New(caps)
		}

		if abscaps.ADFDuplex != nil {
			caps := fromAbstractInputSourceCaps(version,
				abscaps.DocumentFormats, abscaps.ADFDuplex)
			adf.ADFDuplexInputCaps = optional.New(caps)
			adf.ADFOptions = append(adf.ADFOptions, Duplex)
		}

		scancaps.ADF = optional.New(adf)
	}

	return scancaps
}

// fromAbstractInputSourceCaps translates [abstract.InputCapabilities]
// into the [InputSourceCaps].
//
// The version parameters affects how some fields are converted.
func fromAbstractInputSourceCaps(
	version Version, docFormats []string,
	abscaps *abstract.InputCapabilities) InputSourceCaps {

	// Fill InputSourceCaps structure
	caps := InputSourceCaps{
		MinWidth:       abscaps.MinWidth.Dots(300),
		MaxWidth:       abscaps.MaxWidth.Dots(300),
		MinHeight:      abscaps.MinHeight.Dots(300),
		MaxHeight:      abscaps.MaxHeight.Dots(300),
		MaxXOffset:     optional.New(abscaps.MaxXOffset.Dots(300)),
		MaxYOffset:     optional.New(abscaps.MaxYOffset.Dots(300)),
		MaxScanRegions: optional.New(1),

		MaxOpticalXResolution: fromAbstractOptionalInt(
			abscaps.MaxOpticalXResolution),
		MaxOpticalYResolution: fromAbstractOptionalInt(
			abscaps.MaxOpticalYResolution),

		RiskyLeftMargins: fromAbstractOptionalInt(
			abscaps.RiskyLeftMargins.Dots(300)),
		RiskyRightMargins: fromAbstractOptionalInt(
			abscaps.RiskyRightMargins.Dots(300)),
		RiskyTopMargins: fromAbstractOptionalInt(
			abscaps.RiskyTopMargins.Dots(300)),
		RiskyBottomMargins: fromAbstractOptionalInt(
			abscaps.RiskyBottomMargins.Dots(300)),
	}

	// Translate intents
	caps.SupportedIntents = fromAbstractIntents(abscaps.Intents)

	// Translate setting profiles
	caps.SettingProfiles = fromAbstractSettingsProfiles(
		version, docFormats, abscaps.Profiles)

	return caps
}

// fromAbstractSettingsProfiles translates a []abstract.SettingsProfile
// slice into the eSCL []SettingProfile.
//
// The version parameters affects how some fields are converted.
func fromAbstractSettingsProfiles(
	version Version, docFormats []string,
	absprof []abstract.SettingsProfile) []SettingProfile {

	var profiles []SettingProfile
	for _, ap := range absprof {
		// Create SettingProfile structure
		prof := SettingProfile{
			ColorModes: fromAbstractColorModes(
				ap.ColorModes,
				ap.Depths),
			ColorSpaces:     []ColorSpace{SRGB},
			DocumentFormats: docFormats,
			CCDChannels: fromAbstractCCDChannels(
				ap.CCDChannels),
			BinaryRenderings: fromAbstractBinaryRenderings(
				ap.BinaryRenderings),
			SupportedResolutions: fromAbstractResolutions(
				ap.Resolutions,
				ap.ResolutionRange),
		}

		if version >= MakeVersion(2, 1) {
			// Since eSCL 2.1...
			prof.DocumentFormatsExt = prof.DocumentFormats
		}

		// Append to the output
		profiles = append(profiles, prof)
	}

	return profiles
}

// fromAbstractResolutions translates []abstract.Resolution and
// abstract.ResolutionRange into the [SupportedResolutions].
func fromAbstractResolutions(
	absres []abstract.Resolution,
	absresrange abstract.ResolutionRange) SupportedResolutions {

	res := SupportedResolutions{}
	for _, ar := range absres {
		res.DiscreteResolutions = append(
			res.DiscreteResolutions,
			DiscreteResolution{
				ar.XResolution, ar.YResolution,
			})
	}

	if !absresrange.IsZero() {
		xstep := absresrange.XStep
		if xstep < 1 {
			xstep = 1
		}

		ystep := absresrange.YStep
		if ystep < 1 {
			ystep = 1
		}

		rng := ResolutionRange{
			XResolutionRange: Range{
				Min:    absresrange.XMin,
				Max:    absresrange.XMax,
				Normal: absresrange.XNormal,
				Step:   optional.New(xstep),
			},
			YResolutionRange: Range{
				Min:    absresrange.YMin,
				Max:    absresrange.YMax,
				Normal: absresrange.YNormal,
				Step:   optional.New(ystep),
			},
		}

		res.ResolutionRange = optional.New(rng)
	}

	return res
}

// fromAbstractColorModes translates abstract color modes into
// the []ColorMode slice
//
// It returns nil if resulting slice is empty.
func fromAbstractColorModes(
	absmodes generic.Bitset[abstract.ColorMode],
	absdepths generic.Bitset[abstract.Depth]) []ColorMode {

	modes := make([]ColorMode, 0, 5)

	if absmodes.Contains(abstract.ColorModeBinary) {
		modes = append(modes, BlackAndWhite1)
	}

	if absmodes.Contains(abstract.ColorModeMono) {
		if absdepths.Contains(abstract.Depth8) {
			modes = append(modes, Grayscale8)
		}

		if absdepths.Contains(abstract.Depth16) {
			modes = append(modes, Grayscale16)
		}
	}

	if absmodes.Contains(abstract.ColorModeColor) {
		if absdepths.Contains(abstract.Depth8) {
			modes = append(modes, RGB24)
		}

		if absdepths.Contains(abstract.Depth16) {
			modes = append(modes, RGB48)
		}
	}

	if len(modes) == 0 {
		return nil
	}

	return modes
}

// fromAbstractCCDChannel translates abstract.CCDChannel into the
// eSCL CCDChannel.
//
// For invalid or unknown CCDChannel, UnknownCCDChannel is returned.
func fromAbstractCCDChannel(absccd abstract.CCDChannel) CCDChannel {
	switch absccd {
	case abstract.CCDChannelRed:
		return Red
	case abstract.CCDChannelGreen:
		return Green
	case abstract.CCDChannelBlue:
		return Blue
	case abstract.CCDChannelNTSC:
		return NTSC
	case abstract.CCDChannelGrayCcd:
		return GrayCcd
	case abstract.CCDChannelGrayCcdEmulated:
		return GrayCcdEmulated
	}

	return UnknownCCDChannel
}

// fromAbstractCCDChannels translates generic.Bitset[abstract.CCDChannels]
// into the []CCDChannels slice.
//
// It returns nil if resulting slice is empty.
func fromAbstractCCDChannels(
	abschannels generic.Bitset[abstract.CCDChannel]) []CCDChannel {

	if abschannels.IsEmpty() {
		return nil
	}

	in := abschannels.Elements()
	out := make([]CCDChannel, 0, len(in))

	for _, absccd := range in {
		ccd := fromAbstractCCDChannel(absccd)
		if ccd != UnknownCCDChannel {
			out = append(out, ccd)
		}
	}

	if len(out) == 0 {
		return nil
	}

	return out
}

// fromAbstractBinaryRenderings translates
// generic.Bitset[abstract.BinaryRendering] into []BinaryRendering slice
//
// It returns nil if resulting slice is empty.
func fromAbstractBinaryRenderings(
	absrend generic.Bitset[abstract.BinaryRendering]) []BinaryRendering {

	if absrend.IsEmpty() {
		return nil
	}

	in := absrend.Elements()
	out := make([]BinaryRendering, 0, len(in))

	for _, rend := range in {
		switch rend {
		case abstract.BinaryRenderingHalftone:
			out = append(out, Halftone)
		case abstract.BinaryRenderingThreshold:
			out = append(out, Threshold)
		default:
			// Don't know how to translate to the eSCL.
			// Just skip it...
		}
	}

	if len(out) == 0 {
		return nil
	}

	return out
}

// fromAbstractIntent translates [abstract.Intent] into the
// Intent.
//
// For invalid or unknown Intent, UnknownIntent is returned.
func fromAbstractIntent(absintent abstract.Intent) Intent {
	switch absintent {
	case abstract.IntentDocument:
		return Document
	case abstract.IntentTextAndGraphic:
		return TextAndGraphic
	case abstract.IntentPhoto:
		return Photo
	case abstract.IntentPreview:
		return Preview
	case abstract.IntentObject:
		return Object
	case abstract.IntentBusinessCard:
		return BusinessCard
	}
	return UnknownIntent
}

// fromAbstractIntents translates generic.Bitset[abstract.Intent]
// into []Intent slice.
//
// It returns nil if resulting slice is empty.
func fromAbstractIntents(absintents generic.Bitset[abstract.Intent]) []Intent {
	if absintents.IsEmpty() {
		return nil
	}

	in := absintents.Elements()
	out := make([]Intent, 0, len(in))

	for _, absintent := range in {
		intent := fromAbstractIntent(absintent)
		if intent != UnknownIntent {
			out = append(out, intent)
		}
	}

	if len(out) == 0 {
		return nil
	}

	return out
}

// fromAbstractOptionalInt returns optional.New(v), if v != 0, nil otherwise
func fromAbstractOptionalInt(v int) optional.Val[int] {
	if v != 0 {
		return optional.New(v)
	}
	return nil
}

// fromAbstractRange converts abstract.Range into the escl Range if
// Range is not zero, nil otherwise
func fromAbstractOptionalRange(absrange abstract.Range) optional.Val[Range] {
	if !absrange.IsZero() {
		return optional.New(
			Range{
				Min:    absrange.Min,
				Max:    absrange.Max,
				Normal: absrange.Normal,
			},
		)
	}

	return nil
}
