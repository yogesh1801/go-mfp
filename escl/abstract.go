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
	"github.com/alexpevzner/mfp/util/optional"
)

// fromAbstractScannerCapabilities decodes [ScannerCapabilities]
// from the [abstract.ScannerCapabilities].
//
// The version parameters affects how some fields are converted.
func fromAbstractScannerCapabilities(
	version Version,
	abscaps abstract.ScannerCapabilities) ScannerCapabilities {

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

	return scancaps
}

// fromAbstractInputSourceCaps decodes [InputSourceCaps] from
// the abstract.InputCapabilities.
//
// The version parameters affects how some fields are converted.
func fromAbstractInputSourceCaps(
	version Version,
	abscaps *abstract.InputCapabilities) InputSourceCaps {

	caps := InputSourceCaps{
		MinWidth:       abscaps.MinWidth.Dots(300),
		MaxWidth:       abscaps.MaxWidth.Dots(300),
		MinHeight:      abscaps.MinHeight.Dots(300),
		MaxHeight:      abscaps.MaxHeight.Dots(300),
		MaxXOffset:     optional.New(abscaps.MaxXOffset.Dots(300)),
		MaxYOffset:     optional.New(abscaps.MaxYOffset.Dots(300)),
		MaxScanRegions: optional.New(1),
	}

	return caps
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
