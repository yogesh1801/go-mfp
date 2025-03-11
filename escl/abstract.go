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
// from the [abstract.ScannerCapabilities]
func fromAbstractScannerCapabilities(
	abscaps abstract.ScannerCapabilities) ScannerCapabilities {

	scancaps := ScannerCapabilities{
		Version: DefaultVersion,
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
