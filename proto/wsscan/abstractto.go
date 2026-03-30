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
		switch is.Text {
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
				colorEntryToAbstract(cp.Text)
		}

		// Resolution → XResolution, YResolution
		if front.Resolution != nil {
			res := optional.Get(front.Resolution)
			absreq.Resolution.XResolution = res.Width.Text
			absreq.Resolution.YResolution = res.Height.Text
		}

		// ScanRegion → Region
		if front.ScanRegion != nil {
			sr := optional.Get(front.ScanRegion)
			absreq.Region = abstract.Region{
				Width: abstract.DimensionFromDots(
					1000, sr.ScanRegionWidth.Text),
				Height: abstract.DimensionFromDots(
					1000, sr.ScanRegionHeight.Text),
			}

			if sr.ScanRegionXOffset != nil {
				absreq.Region.XOffset = abstract.DimensionFromDots(
					1000, optional.Get(sr.ScanRegionXOffset).Text)
			}
			if sr.ScanRegionYOffset != nil {
				absreq.Region.YOffset = abstract.DimensionFromDots(
					1000, optional.Get(sr.ScanRegionYOffset).Text)
			}
		}
	}

	// Translate Format → DocumentFormat
	if dp.Format != nil {
		fmtVal := optional.Get(dp.Format)
		absreq.DocumentFormat = fmtVal.Text.MimeType()
	}

	// Translate CompressionQualityFactor → Compression
	if dp.CompressionQualityFactor != nil {
		cqf := optional.Get(dp.CompressionQualityFactor)
		absreq.Compression = optional.New(cqf.Text)
	}

	// Translate ContentType → Intent
	if dp.ContentType != nil {
		ct := optional.Get(dp.ContentType)
		switch ct.Text {
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
				absreq.Brightness = optional.New(b.Text)
			}
			if es.Contrast != nil {
				c := optional.Get(es.Contrast)
				absreq.Contrast = optional.New(c.Text)
			}
			if es.Sharpness != nil {
				s := optional.Get(es.Sharpness)
				absreq.Sharpen = optional.New(s.Text)
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
