// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Conversions from eSCL to abstract.Scanner data structures

package escl

import "github.com/OpenPrinting/go-mfp/abstract"

// toAbstract converts [ScanSettings] to [abstract.ScannerRequest]
func (ss ScanSettings) toAbstract() abstract.ScannerRequest {
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
		switch *ss.ColorMode {
		case BlackAndWhite1:
			absreq.ColorMode = abstract.ColorModeBinary
			if ss.BinaryRendering != nil {
				switch *ss.BinaryRendering {
				case Halftone:
					absreq.BinaryRendering = abstract.
						BinaryRenderingHalftone
				case Threshold:
					absreq.BinaryRendering = abstract.
						BinaryRenderingThreshold
					if ss.Threshold != nil {
						absreq.Threshold = ss.Threshold
					}
				}
			}
		case Grayscale8:
			absreq.ColorMode = abstract.ColorModeMono
			absreq.ColorDepth = abstract.ColorDepth8
		case Grayscale16:
			absreq.ColorMode = abstract.ColorModeMono
			absreq.ColorDepth = abstract.ColorDepth16
		case RGB24:
			absreq.ColorMode = abstract.ColorModeColor
			absreq.ColorDepth = abstract.ColorDepth8
		case RGB48:
			absreq.ColorMode = abstract.ColorModeColor
			absreq.ColorDepth = abstract.ColorDepth16
		}
	}

	// Translate CCDChannel
	if ss.CCDChannel != nil {
		switch *ss.CCDChannel {
		case Red:
			absreq.CCDChannel = abstract.CCDChannelRed
		case Green:
			absreq.CCDChannel = abstract.CCDChannelGreen
		case Blue:
			absreq.CCDChannel = abstract.CCDChannelBlue
		case NTSC:
			absreq.CCDChannel = abstract.CCDChannelNTSC
		case GrayCcd:
			absreq.CCDChannel = abstract.CCDChannelGrayCcd
		case GrayCcdEmulated:
			absreq.CCDChannel = abstract.CCDChannelGrayCcdEmulated
		}
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
