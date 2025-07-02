// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Conversions from eSCL to abstract.Scanner data structures

package escl

import (
	"github.com/OpenPrinting/go-mfp/abstract"
	"github.com/OpenPrinting/go-mfp/util/generic"
	"github.com/OpenPrinting/go-mfp/util/optional"
)

// ToAbstract converts [ScannerCapabilities] to [abstract.ScannerCapabilities].
func (scancaps ScannerCapabilities) ToAbstract() abstract.ScannerCapabilities {
	abscaps := abstract.ScannerCapabilities{
		UUID:         optional.Get(scancaps.UUID),
		MakeAndModel: optional.Get(scancaps.MakeAndModel),
		SerialNumber: optional.Get(scancaps.SerialNumber),
		Manufacturer: optional.Get(scancaps.Manufacturer),
		AdminURI:     optional.Get(scancaps.AdminURI),
		IconURI:      optional.Get(scancaps.IconURI),

		DocumentFormats:  scancaps.DocumentFormats(),
		CompressionRange: optional.Get(scancaps.CompressionFactorSupport).toAbstract(),

		BrightnessRange:   optional.Get(scancaps.BrightnessSupport).toAbstract(),
		ContrastRange:     optional.Get(scancaps.ContrastSupport).toAbstract(),
		GammaRange:        optional.Get(scancaps.GammaSupport).toAbstract(),
		HighlightRange:    optional.Get(scancaps.HighlightSupport).toAbstract(),
		NoiseRemovalRange: optional.Get(scancaps.NoiseRemovalSupport).toAbstract(),
		ShadowRange:       optional.Get(scancaps.ShadowSupport).toAbstract(),
		SharpenRange:      optional.Get(scancaps.SharpenSupport).toAbstract(),
		ThresholdRange:    optional.Get(scancaps.ThresholdSupport).toAbstract(),
	}

	if scancaps.Platen != nil {
		abscaps.Platen = (*scancaps.Platen.PlatenInputCaps).toAbstract()
	}

	if scancaps.ADF != nil {
		abscaps.ADFCapacity = optional.Get(scancaps.ADF.FeederCapacity)

		if scancaps.ADF.ADFSimplexInputCaps != nil {
			abscaps.ADFSimplex =
				(*scancaps.ADF.ADFSimplexInputCaps).toAbstract()
		}

		if scancaps.ADF.ADFDuplexInputCaps != nil {
			abscaps.ADFDuplex =
				(*scancaps.ADF.ADFDuplexInputCaps).toAbstract()
		}
	}

	return abscaps
}

// ToAbstract converts [ScanSettings] to [abstract.ScannerRequest]
func (ss ScanSettings) ToAbstract() abstract.ScannerRequest {
	absreq := abstract.ScannerRequest{
		Brightness:   ss.Brightness,
		Contrast:     ss.Contrast,
		Gamma:        ss.Gamma,
		Highlight:    ss.Highlight,
		NoiseRemoval: ss.NoiseRemoval,
		Shadow:       ss.Shadow,
		Sharpen:      ss.Sharpen,
		Compression:  ss.CompressionFactor,
	}

	// Translate Input and ADFMode
	if ss.InputSource != nil {
		switch *ss.InputSource {
		case InputPlaten:
			absreq.Input = abstract.InputPlaten
		case InputFeeder:
			absreq.Input = abstract.InputADF
			absreq.ADFMode = abstract.ADFModeSimplex
			if ss.Duplex != nil && *ss.Duplex {
				absreq.ADFMode = abstract.ADFModeDuplex
			}
		}
	}

	// Translate ColorMode, ColorDepth, BinaryRendering and Threshold
	if ss.ColorMode != nil {
		absreq.ColorMode, absreq.ColorDepth = (*ss.ColorMode).toAbstract()
		if absreq.ColorMode == abstract.ColorModeBinary {
			if ss.BinaryRendering != nil {
				absreq.BinaryRendering =
					(*ss.BinaryRendering).toAbstract()

				if *ss.BinaryRendering == Threshold {
					if ss.Threshold != nil {
						absreq.Threshold = ss.Threshold
					}
				}
			}
		}
	}

	// Translate CCDChannel
	if ss.CCDChannel != nil {
		absreq.CCDChannel = (*ss.CCDChannel).toAbstract()
	}

	// Translate DocumentFormat
	//
	// Although DocumentFormatExt was introduced by the eSCL 2.1+,
	// we always prefer DocumentFormatExt with fallback to DocumentFormat,
	// regardless of the eSCL version.
	switch {
	case ss.DocumentFormatExt != nil:
		absreq.DocumentFormat = *ss.DocumentFormatExt
	case ss.DocumentFormat != nil:
		absreq.DocumentFormat = *ss.DocumentFormat
	}

	// Translate ScanRegions. Note, we only handle the first one.
	if len(ss.ScanRegions) > 0 {
		reg := ss.ScanRegions[0]
		absreq.Region = abstract.Region{
			XOffset: abstract.DimensionFromDots(300, reg.XOffset),
			YOffset: abstract.DimensionFromDots(300, reg.YOffset),
			Width:   abstract.DimensionFromDots(300, reg.Width),
			Height:  abstract.DimensionFromDots(300, reg.Height),
		}
	}

	// Translate Resolution
	if ss.XResolution != nil {
		absreq.Resolution.XResolution = *ss.XResolution
	}
	if ss.YResolution != nil {
		absreq.Resolution.YResolution = *ss.YResolution
	}

	// Translate Intent
	if ss.Intent != nil {
		switch *ss.Intent {
		case Document:
			absreq.Intent = abstract.IntentDocument
		case TextAndGraphic:
			absreq.Intent = abstract.IntentTextAndGraphic
		case Photo:
			absreq.Intent = abstract.IntentPhoto
		case Preview:
			absreq.Intent = abstract.IntentPreview
		case Object:
			absreq.Intent = abstract.IntentObject
		case BusinessCard:
			absreq.Intent = abstract.IntentBusinessCard
		}
	}

	return absreq
}

// toAbstract converts [InputSourceCaps] to *[anstract.InputCapabilities].
func (caps InputSourceCaps) toAbstract() *abstract.InputCapabilities {
	abscaps := &abstract.InputCapabilities{
		MinWidth:  abstract.DimensionFromDots(300, caps.MinWidth),
		MaxWidth:  abstract.DimensionFromDots(300, caps.MaxWidth),
		MinHeight: abstract.DimensionFromDots(300, caps.MinHeight),
		MaxHeight: abstract.DimensionFromDots(300, caps.MaxHeight),

		MaxXOffset: abstract.DimensionFromDots(300, optional.Get(caps.MaxXOffset)),
		MaxYOffset: abstract.DimensionFromDots(300, optional.Get(caps.MaxYOffset)),

		MaxOpticalXResolution: optional.Get(caps.MaxOpticalXResolution),
		MaxOpticalYResolution: optional.Get(caps.MaxOpticalYResolution),

		RiskyLeftMargins:   abstract.DimensionFromDots(300, optional.Get(caps.RiskyLeftMargins)),
		RiskyRightMargins:  abstract.DimensionFromDots(300, optional.Get(caps.RiskyRightMargins)),
		RiskyTopMargins:    abstract.DimensionFromDots(300, optional.Get(caps.RiskyTopMargins)),
		RiskyBottomMargins: abstract.DimensionFromDots(300, optional.Get(caps.RiskyBottomMargins)),
	}

	for _, intent := range caps.SupportedIntents {
		switch intent {
		case Document:
			abscaps.Intents.Add(abstract.IntentDocument)
		case TextAndGraphic:
			abscaps.Intents.Add(abstract.IntentTextAndGraphic)
		case Photo:
			abscaps.Intents.Add(abstract.IntentPhoto)
		case Preview:
			abscaps.Intents.Add(abstract.IntentPreview)
		case Object:
			abscaps.Intents.Add(abstract.IntentObject)
		case BusinessCard:
			abscaps.Intents.Add(abstract.IntentBusinessCard)
		}
	}

	for _, prof := range caps.SettingProfiles {
		abscaps.Profiles = append(abscaps.Profiles,
			prof.toAbstract()...)
	}

	return abscaps
}

// toAbstract converts [SettingProfile] to [anstract.SettingsProfile].
//
// Note, converting a single eSCL SettingProfile may result multiple
// abstract.SettingsProfile
func (prof SettingProfile) toAbstract() []abstract.SettingsProfile {
	absprof := abstract.SettingsProfile{}

	// Convert common part
	for _, cm := range prof.ColorModes {
		switch cm {
		case BlackAndWhite1:
			absprof.ColorModes.Add(abstract.ColorModeBinary)
		case Grayscale8:
			absprof.ColorModes.Add(abstract.ColorModeMono)
			absprof.Depths.Add(abstract.ColorDepth8)
		case Grayscale16:
			absprof.ColorModes.Add(abstract.ColorModeMono)
			absprof.Depths.Add(abstract.ColorDepth16)
		case RGB24:
			absprof.ColorModes.Add(abstract.ColorModeColor)
			absprof.Depths.Add(abstract.ColorDepth8)
		case RGB48:
			absprof.ColorModes.Add(abstract.ColorModeColor)
			absprof.Depths.Add(abstract.ColorDepth16)
		}
	}

	for _, rnd := range prof.BinaryRenderings {
		absrnd := rnd.toAbstract()
		if absrnd != abstract.BinaryRenderingUnset {
			absprof.BinaryRenderings.Add(absrnd)
		}
	}

	for _, ccd := range prof.CCDChannels {
		absccd := ccd.toAbstract()
		if absccd != abstract.CCDChannelUnset {
			absprof.CCDChannels.Add(absccd)
		}
	}

	// Convert resolutions.
	//
	// eSCL SettingProfile may contain multiple sets of resolutions,
	// optionally constrained by the color mode (this feature rarely
	// used in the wild, but defined by the specification).
	//
	// The anstract.SettingsProfile doesn't provide such a functionality.
	//
	// To address this problem, we generate multiple instances of
	// the anstract.SettingsProfile, one per resolutions set defined
	// at the eSCL side.
	absprofs := make([]abstract.SettingsProfile,
		len(prof.SupportedResolutions))

	for i := range prof.SupportedResolutions {
		absprofs[i] = absprof
		supported := prof.SupportedResolutions[i]

		// Adjust color mode, if constrained
		if supported.ColorMode != nil {
			cm := *prof.SupportedResolutions[i].ColorMode
			abscm, absdepth := cm.toAbstract()
			absprofs[i].ColorModes = generic.MakeBitset(abscm)
			absprofs[i].Depths = generic.MakeBitset(absdepth)
		}

		// Translate resolutions
		for _, res := range supported.DiscreteResolutions {
			absprofs[i].Resolutions = append(
				absprofs[i].Resolutions, res.toAbstract())
		}

		if supported.ResolutionRange != nil {
			absprofs[i].ResolutionRange =
				(*supported.ResolutionRange).toAbstract()
		}
	}

	return absprofs
}

// toAbstract converts [DiscreteResolution] to [abstract.Resolution]
func (res DiscreteResolution) toAbstract() abstract.Resolution {
	return abstract.Resolution{
		XResolution: res.XResolution,
		YResolution: res.YResolution,
	}
}

// toAbstract converts [ResolutionRange] to [abstract.ResolutionRange]
func (rng ResolutionRange) toAbstract() abstract.ResolutionRange {
	return abstract.ResolutionRange{
		XMin:    rng.XResolutionRange.Min,
		XMax:    rng.XResolutionRange.Max,
		XStep:   optional.Get(rng.XResolutionRange.Step),
		XNormal: rng.XResolutionRange.Normal,

		YMin:    rng.YResolutionRange.Min,
		YMax:    rng.YResolutionRange.Max,
		YStep:   optional.Get(rng.YResolutionRange.Step),
		YNormal: rng.YResolutionRange.Normal,
	}
}

// toAbstract converts [Range] to [abstract.Range]
func (r Range) toAbstract() abstract.Range {
	return abstract.Range{
		Min:    r.Min,
		Max:    r.Max,
		Normal: r.Normal,
		Step:   optional.Get(r.Step),
	}
}

// toAbstract converts [CCDChannel] to [abstract.CCDChannel]
func (rnd BinaryRendering) toAbstract() abstract.BinaryRendering {
	switch rnd {
	case Halftone:
		return abstract.BinaryRenderingHalftone
	case Threshold:
		return abstract.BinaryRenderingThreshold
	}

	return abstract.BinaryRenderingUnset
}

// toAbstract converts [CCDChannel] to [abstract.CCDChannel]
func (ccd CCDChannel) toAbstract() abstract.CCDChannel {
	switch ccd {
	case Red:
		return abstract.CCDChannelRed
	case Green:
		return abstract.CCDChannelGreen
	case Blue:
		return abstract.CCDChannelBlue
	case NTSC:
		return abstract.CCDChannelNTSC
	case GrayCcd:
		return abstract.CCDChannelGrayCcd
	case GrayCcdEmulated:
		return abstract.CCDChannelGrayCcdEmulated
	}

	return abstract.CCDChannelUnset
}

// toAbstract converts [ColorMode] into the combination of the
// [abstract.ColorMode] and [abstract.ColorDepth].
func (cm ColorMode) toAbstract() (abstract.ColorMode, abstract.ColorDepth) {
	switch cm {
	case BlackAndWhite1:
		return abstract.ColorModeBinary, abstract.ColorDepthUnset
	case Grayscale8:
		return abstract.ColorModeMono, abstract.ColorDepth8
	case Grayscale16:
		return abstract.ColorModeMono, abstract.ColorDepth16
	case RGB24:
		return abstract.ColorModeColor, abstract.ColorDepth8
	case RGB48:
		return abstract.ColorModeColor, abstract.ColorDepth16
	}

	return abstract.ColorModeUnset, abstract.ColorDepthUnset
}
