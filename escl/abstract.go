// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Conversion between eSCL and abstract.Scanner structures

package escl

import (
	"github.com/alexpevzner/mfp/abstract"
	"github.com/alexpevzner/mfp/util/generic"
	"github.com/alexpevzner/mfp/util/optional"
)

// fromAbstractScannerCapabilities decodes [ScannerCapabilities]
// from the [abstract.ScannerCapabilities].
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
	scancaps.BrightnessSupport = fromAbstractRange(abscaps.BrightnessRange)
	scancaps.ContrastSupport = fromAbstractRange(abscaps.ContrastRange)
	scancaps.GammaSupport = fromAbstractRange(abscaps.GammaRange)
	scancaps.HighlightSupport = fromAbstractRange(abscaps.HighlightRange)
	scancaps.NoiseRemovalSupport = fromAbstractRange(
		abscaps.NoiseRemovalRange)
	scancaps.ShadowSupport = fromAbstractRange(abscaps.ShadowRange)
	scancaps.SharpenSupport = fromAbstractRange(abscaps.SharpenRange)
	scancaps.ThresholdSupport = fromAbstractRange(abscaps.ThresholdRange)
	scancaps.CompressionFactorSupport = fromAbstractRange(
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

// fromAbstractInputSourceCaps decodes [InputSourceCaps] from
// the abstract.InputCapabilities.
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
	if !abscaps.Intents.IsEmpty() {
		caps.SupportedIntents = fromAbstractIntents(abscaps.Intents)
	}

	// Translate setting profiles
	for _, absprof := range abscaps.Profiles {
		// Create SettingProfile structure
		prof := SettingProfile{
			ColorModes: fromAbstractColorModes(
				absprof.ColorModes,
				absprof.Depths),
			ColorSpaces:     []ColorSpace{SRGB},
			DocumentFormats: docFormats,
			CCDChannels: fromAbstractCCDChannels(
				absprof.CCDChannels),
			BinaryRenderings: fromAbstractBinaryRenderings(
				absprof.BinaryRenderings),
		}

		if version >= MakeVersion(2, 1) {
			// Since eSCL 2.1...
			prof.DocumentFormatsExt = prof.DocumentFormats
		}

		// Translate resolutions
		res := SupportedResolutions{}
		for _, absres := range absprof.Resolutions {
			res.DiscreteResolutions = append(
				res.DiscreteResolutions,
				DiscreteResolution{
					absres.XResolution, absres.YResolution,
				})
		}

		if absrng := absprof.ResolutionRange; !absrng.IsZero() {
			xstep := absrng.XStep
			if xstep < 1 {
				xstep = 0
			}

			ystep := absrng.YStep
			if ystep < 1 {
				ystep = 0
			}

			rng := ResolutionRange{
				XResolutionRange: Range{
					Min:  absrng.XMin,
					Max:  absrng.XMax,
					Step: optional.New(xstep),
				},
				YResolutionRange: Range{
					Min:  absrng.YMin,
					Max:  absrng.YMax,
					Step: optional.New(ystep),
				},
			}

			res.ResolutionRange = optional.New(rng)
		}

		prof.SupportedResolutions = res

		// Append to capabilities
		caps.SettingProfiles = append(caps.SettingProfiles, prof)
	}

	return caps
}

// fromAbstractColorModes translates abstract color modes into
// the []ColorMode slice
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

	return modes
}

// fromAbstractCCDChannels translates generic.Bitset[abstract.CCDChannels]
// into the []CCDChannels slice.
func fromAbstractCCDChannels(
	absrend generic.Bitset[abstract.CCDChannel]) []CCDChannel {

	in := absrend.Elements()
	out := make([]CCDChannel, 0, len(in))

	for _, ccd := range in {
		switch ccd {
		case abstract.CCDChannelRed:
			out = append(out, Red)
		case abstract.CCDChannelGreen:
			out = append(out, Green)
		case abstract.CCDChannelBlue:
			out = append(out, Blue)
		case abstract.CCDChannelNTSC:
			out = append(out, NTSC)
		case abstract.CCDChannelGrayCcd:
			out = append(out, GrayCcd)
		case abstract.CCDChannelGrayCcdEmulated:
			out = append(out, GrayCcdEmulated)
		default:
			// Don't know how to translate to the eSCL.
			// Just skip it...
		}
	}

	return out
}

// fromAbstractBinaryRenderings translates
// generic.Bitset[abstract.BinaryRendering] into []BinaryRendering slice
func fromAbstractBinaryRenderings(
	absrend generic.Bitset[abstract.BinaryRendering]) []BinaryRendering {

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

	return out
}

// fromAbstractIntents translates generic.Bitset[abstract.Intent]
// into []Intent slice
func fromAbstractIntents(absintents generic.Bitset[abstract.Intent]) []Intent {
	in := absintents.Elements()
	out := make([]Intent, 0, len(in))

	for _, intent := range in {
		switch intent {
		case abstract.IntentDocument:
			out = append(out, Document)
		case abstract.IntentTextAndGraphic:
			out = append(out, TextAndGraphic)
		case abstract.IntentPhoto:
			out = append(out, Photo)
		case abstract.IntentPreview:
			out = append(out, Preview)
		case abstract.IntentObject:
			out = append(out, Object)
		case abstract.IntentBusinessCard:
			out = append(out, BusinessCard)
		default:
			// Don't know how to translate to the eSCL.
			// Just skip it...
		}
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

// fromAbstractRange converts abstract.Range into the escl Range
func fromAbstractRange(absrange abstract.Range) optional.Val[Range] {
	if absrange.Supported() {
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
