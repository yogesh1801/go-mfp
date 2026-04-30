// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// Conversions from WS-Scan to abstract.Scanner data structures

package wsscan

import (
	"github.com/OpenPrinting/go-mfp/abstract"
	"github.com/OpenPrinting/go-mfp/util/optional"
)

// ToAbstract converts [GetScannerElementsResponse] to
// *[abstract.ScannerCapabilities] by extracting the first
// [ScannerConfiguration] and [ScannerDescription] from ScannerElements.
func (r *GetScannerElementsResponse) ToAbstract() *abstract.ScannerCapabilities {
	var sc ScannerConfiguration
	var sd ScannerDescription

	for _, elem := range r.ScannerElements {
		if elem.ScannerConfiguration != nil {
			sc = optional.Get(elem.ScannerConfiguration)
		}
		if elem.ScannerDescription != nil {
			sd = optional.Get(elem.ScannerDescription)
		}
	}

	caps := &abstract.ScannerCapabilities{}

	caps.MakeAndModel = sd.ScannerName.NeutralLang().Text

	ds := sc.DeviceSettings
	for _, fv := range ds.FormatsSupported {
		if mime := fv.MimeType(); mime != "" {
			caps.DocumentFormats = append(caps.DocumentFormats, mime)
		}
	}

	cqf := ds.CompressionQualityFactorSupported
	caps.CompressionRange = abstract.Range{
		Min:    cqf.MinValue,
		Max:    cqf.MaxValue,
		Normal: (cqf.MinValue + cqf.MaxValue) / 2,
	}

	if ds.BrightnessSupported.Bool() {
		caps.BrightnessRange = abstract.Range{Min: 0, Max: 100, Normal: 50}
	}
	if ds.ContrastSupported.Bool() {
		caps.ContrastRange = abstract.Range{Min: 0, Max: 100, Normal: 50}
	}

	if sc.Platen != nil {
		p := optional.Get(sc.Platen)
		caps.Platen = platenToAbstract(p)
	}

	if sc.ADF != nil {
		adf := optional.Get(sc.ADF)
		if adf.ADFFront != nil {
			caps.ADFSimplex = adfSideToAbstract(optional.Get(adf.ADFFront))
		}
		if adf.ADFSupportsDuplex.Bool() && adf.ADFBack != nil {
			caps.ADFDuplex = adfSideToAbstract(optional.Get(adf.ADFBack))
		}
	}

	return caps
}

// platenToAbstract converts [Platen] to *[abstract.InputCapabilities].
func platenToAbstract(p Platen) *abstract.InputCapabilities {
	return &abstract.InputCapabilities{
		MinWidth: abstract.DimensionFromDots(
			wsscanDPI, p.PlatenMinimumSize.Width),
		MaxWidth: abstract.DimensionFromDots(
			wsscanDPI, p.PlatenMaximumSize.Width),
		MinHeight: abstract.DimensionFromDots(
			wsscanDPI, p.PlatenMinimumSize.Height),
		MaxHeight: abstract.DimensionFromDots(
			wsscanDPI, p.PlatenMaximumSize.Height),
		MaxOpticalXResolution: p.PlatenOpticalResolution.Width,
		MaxOpticalYResolution: p.PlatenOpticalResolution.Height,
		Profiles: colorResolutionsToProfiles(
			p.PlatenColor, p.PlatenResolutions),
	}
}

// adfSideToAbstract converts [ADFSide] to *[abstract.InputCapabilities].
func adfSideToAbstract(s ADFSide) *abstract.InputCapabilities {
	return &abstract.InputCapabilities{
		MinWidth: abstract.DimensionFromDots(
			wsscanDPI, s.ADFMinimumSize.Width),
		MaxWidth: abstract.DimensionFromDots(
			wsscanDPI, s.ADFMaximumSize.Width),
		MinHeight: abstract.DimensionFromDots(
			wsscanDPI, s.ADFMinimumSize.Height),
		MaxHeight: abstract.DimensionFromDots(
			wsscanDPI, s.ADFMaximumSize.Height),
		MaxOpticalXResolution: s.ADFOpticalResolution.Width,
		MaxOpticalYResolution: s.ADFOpticalResolution.Height,
		Profiles: colorResolutionsToProfiles(
			s.ADFColor, s.ADFResolutions),
	}
}

// colorResolutionsToProfiles builds a single [abstract.SettingsProfile]
// from WS-Scan color entries and resolution lists. Resolutions are the
// cross-product of all widths × heights.
func colorResolutionsToProfiles(
	colors []ColorEntry, res Resolutions) []abstract.SettingsProfile {

	prof := abstract.SettingsProfile{}

	for _, ce := range colors {
		mode, depth := colorEntryToAbstract(ce)
		if mode == abstract.ColorModeUnset {
			continue
		}
		prof.ColorModes.Add(mode)
		if depth != abstract.ColorDepthUnset {
			prof.Depths.Add(depth)
		}
		if mode == abstract.ColorModeBinary {
			prof.BinaryRenderings.Add(abstract.BinaryRenderingThreshold)
		}
	}

	for _, w := range res.Widths {
		for _, h := range res.Heights {
			prof.Resolutions = append(prof.Resolutions,
				abstract.Resolution{XResolution: w, YResolution: h})
		}
	}

	return []abstract.SettingsProfile{prof}
}

// ToAbstract converts [ScanTicket] to [abstract.ScannerRequest].
func (ticket ScanTicket) ToAbstract() abstract.ScannerRequest {
	absreq := abstract.ScannerRequest{}

	if ticket.DocumentParameters == nil {
		return absreq
	}
	dp := optional.Get(ticket.DocumentParameters)

	// Translate InputSource → Input + ADFMode
	if dp.InputSource != nil {
		is := optional.Get(dp.InputSource)
		switch is.Val {
		case InputSourcePlaten:
			absreq.Input = abstract.InputPlaten
		case InputSourceADF:
			absreq.Input = abstract.InputADF
			absreq.ADFMode = abstract.ADFModeSimplex
		case InputSourceADFDuplex:
			absreq.Input = abstract.InputADF
			absreq.ADFMode = abstract.ADFModeDuplex
		}
	}

	// Translate MediaSides → ColorMode, ColorDepth, Resolution, Region
	if dp.MediaSides != nil {
		ms := optional.Get(dp.MediaSides)
		front := ms.MediaFront

		// ColorProcessing → ColorMode + ColorDepth
		if front.ColorProcessing != nil {
			cp := optional.Get(front.ColorProcessing)
			absreq.ColorMode, absreq.ColorDepth =
				colorEntryToAbstract(cp.Val)
		}

		// Resolution → XResolution, YResolution
		if front.Resolution != nil {
			res := optional.Get(front.Resolution)
			absreq.Resolution.XResolution = res.Width.Val
			absreq.Resolution.YResolution = res.Height.Val
		}

		// ScanRegion → Region
		if front.ScanRegion != nil {
			sr := optional.Get(front.ScanRegion)
			absreq.Region = abstract.Region{
				Width: abstract.DimensionFromDots(
					1000, sr.ScanRegionWidth.Val),
				Height: abstract.DimensionFromDots(
					1000, sr.ScanRegionHeight.Val),
			}

			if sr.ScanRegionXOffset != nil {
				absreq.Region.XOffset = abstract.DimensionFromDots(
					1000, optional.Get(sr.ScanRegionXOffset).Val)
			}
			if sr.ScanRegionYOffset != nil {
				absreq.Region.YOffset = abstract.DimensionFromDots(
					1000, optional.Get(sr.ScanRegionYOffset).Val)
			}
		}
	}

	// Translate Format → DocumentFormat
	if dp.Format != nil {
		fmtVal := optional.Get(dp.Format)
		absreq.DocumentFormat = fmtVal.Val.MimeType()
	}

	// Translate CompressionQualityFactor → Compression
	if dp.CompressionQualityFactor != nil {
		cqf := optional.Get(dp.CompressionQualityFactor)
		absreq.Compression = optional.New(cqf.Val)
	}

	// Translate ContentType → Intent
	if dp.ContentType != nil {
		ct := optional.Get(dp.ContentType)
		switch ct.Val {
		case Text:
			absreq.Intent = abstract.IntentDocument
		case Photo:
			absreq.Intent = abstract.IntentPhoto
		case Mixed:
			absreq.Intent = abstract.IntentTextAndGraphic
		case Halftone:
			absreq.Intent = abstract.IntentDocument
		}
	}

	// Translate Exposure → Brightness, Contrast, Sharpen
	if dp.Exposure != nil {
		exp := optional.Get(dp.Exposure)
		if exp.ExposureSettings != nil {
			es := optional.Get(exp.ExposureSettings)

			if es.Brightness != nil {
				b := optional.Get(es.Brightness)
				absreq.Brightness = optional.New(b.Val)
			}
			if es.Contrast != nil {
				c := optional.Get(es.Contrast)
				absreq.Contrast = optional.New(c.Val)
			}
			if es.Sharpness != nil {
				s := optional.Get(es.Sharpness)
				absreq.Sharpen = optional.New(s.Val)
			}
		}
	}

	return absreq
}

// colorEntryToAbstract converts [ColorEntry] into the combination of
// [abstract.ColorMode] and [abstract.ColorDepth].
func colorEntryToAbstract(ce ColorEntry) (
	abstract.ColorMode, abstract.ColorDepth) {
	switch ce {
	case BlackAndWhite1:
		return abstract.ColorModeBinary, abstract.ColorDepthUnset
	case Grayscale4:
		return abstract.ColorModeMono, abstract.ColorDepth8
	case Grayscale8:
		return abstract.ColorModeMono, abstract.ColorDepth8
	case Grayscale16:
		return abstract.ColorModeMono, abstract.ColorDepth16
	case RGB24:
		return abstract.ColorModeColor, abstract.ColorDepth8
	case RGB48:
		return abstract.ColorModeColor, abstract.ColorDepth16
	case RGBA32:
		return abstract.ColorModeColor, abstract.ColorDepth8
	case RGBA64:
		return abstract.ColorModeColor, abstract.ColorDepth16
	}

	return abstract.ColorModeUnset, abstract.ColorDepthUnset
}
