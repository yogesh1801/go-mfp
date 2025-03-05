// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scan settings (scan request parameters)

package escl

import (
	"github.com/alexpevzner/mfp/util/optional"
	"github.com/alexpevzner/mfp/util/xmldoc"
)

// ScanSettings defines the set of parameters for scan request.
//
// eSCL Technical Specification, 7.
type ScanSettings struct {
	// Version is the only required parameter
	Version Version // eSCL protocol version

	// General parameters
	Intent            optional.Val[Intent]          // Scan intent
	ScanRegions       []ScanRegion                  // List of scan regions
	DocumentFormat    optional.Val[string]          // Image fmt (MIME type)
	DocumentFormatExt optional.Val[string]          // Image fmt, eSCL 2.1+
	ContentType       optional.Val[ContentType]     // Content type
	InputSource       optional.Val[InputSource]     // Desired input source
	XResolution       optional.Val[int]             // X resolution, DPI
	YResolution       optional.Val[int]             // Y resolution, DPI
	ColorMode         optional.Val[ColorMode]       // Desired color mode
	ColorSpace        optional.Val[ColorSpace]      // Desired color space
	CcdChannel        optional.Val[CcdChannel]      // Desired CCD channel
	BinaryRendering   optional.Val[BinaryRendering] // For BlackAndWhite1
	Duplex            optional.Val[bool]            // For ADF
	FeedDirection     optional.Val[FeedDirection]   // Desired feed dir

	// Image transform parameters
	Brightness        optional.Val[int] // Brightness
	CompressionFactor optional.Val[int] // Lower num, better image
	Contrast          optional.Val[int] // Contrast
	Gamma             optional.Val[int] // Gamma (y=x^(1/g)
	Highlight         optional.Val[int] // Image Highlight
	NoiseRemoval      optional.Val[int] // Noise removal level
	Shadow            optional.Val[int] // The lower, the darger
	Sharpen           optional.Val[int] // Image sharpen
	Threshold         optional.Val[int] // For BlackAndWhite1

	// Blank page detection and removal (ADF only).
	//
	// If blank page detection is requested, device should set the
	// BlankPageDetected element of ScanImageInfo resource SHOULD be set
	// appropriately.
	//
	// If blank page removal is requested, device should skip the
	// skip the scanned blank pages.
	BlankPageDetection           optional.Val[bool] // Detection requested
	BlankPageDetectionAndRemoval optional.Val[bool] // Auto-remove requested
}

// DecodeScanSettings decodes [ScanSettings] from the XML tree.
func DecodeScanSettings(root xmldoc.Element) (
	ss ScanSettings, err error) {

	defer func() { err = xmldoc.XMLErrWrap(root, err) }()

	// Lookup relevant XML elements
	ver := xmldoc.Lookup{Name: NsPWG + ":Version", Required: true}
	intent := xmldoc.Lookup{Name: NsScan + ":Intent"}
	regs := xmldoc.Lookup{Name: NsPWG + ":ScanRegions"}
	fmt := xmldoc.Lookup{Name: NsPWG + ":DocumentFormat"}
	fmtExt := xmldoc.Lookup{Name: NsScan + ":DocumentFormatExt"}
	content := xmldoc.Lookup{Name: NsPWG + ":ContentType"}
	input := xmldoc.Lookup{Name: NsPWG + ":InputSource"}
	xres := xmldoc.Lookup{Name: NsScan + ":XResolution"}
	yres := xmldoc.Lookup{Name: NsScan + ":YResolution"}
	mode := xmldoc.Lookup{Name: NsScan + ":ColorSpace"}
	space := xmldoc.Lookup{Name: NsScan + ":ColorSpace"}
	ccd := xmldoc.Lookup{Name: NsScan + ":CcdChannel"}
	binrend := xmldoc.Lookup{Name: NsScan + ":BinaryRendering"}
	duplex := xmldoc.Lookup{Name: NsScan + ":Duplex"}
	feed := xmldoc.Lookup{Name: NsScan + ":FeedDirection"}

	brightness := xmldoc.Lookup{Name: NsScan + ":Brightness"}
	compression := xmldoc.Lookup{Name: NsScan + ":CompressionFactor"}
	contrast := xmldoc.Lookup{Name: NsScan + ":Contrast"}
	gamma := xmldoc.Lookup{Name: NsScan + ":Gamma"}
	highlight := xmldoc.Lookup{Name: NsScan + ":Highlight"}
	noiseRemoval := xmldoc.Lookup{Name: NsScan + ":NoiseRemoval"}
	shadow := xmldoc.Lookup{Name: NsScan + ":Shadow"}
	sharpen := xmldoc.Lookup{Name: NsScan + ":Sharpen"}
	threshold := xmldoc.Lookup{Name: NsScan + ":Threshold"}

	blankDetect := xmldoc.Lookup{Name: NsScan + ":BlankPageDetection"}
	blankRemove := xmldoc.Lookup{Name: NsScan +
		":BlankPageDetectionAndRemoval"}

	missed := root.Lookup(&ver, &intent, &regs, &fmt, &fmtExt,
		&content, &input, &xres, &yres, &mode,
		&space, &ccd, &binrend, &duplex, &feed,
		&brightness, &compression, &contrast, &gamma, &highlight,
		&noiseRemoval, &shadow, &sharpen, &threshold,
		&blankDetect, &blankRemove)

	if missed != nil {
		err = xmldoc.XMLErrMissed(missed.Name)
		return
	}

	// Decode elements
	ss.Version, err = decodeVersion(ver.Elem)
	if err != nil {
		return
	}

	if intent.Found {
		ss.Intent, err = decodeOptional(intent.Elem, decodeIntent)
		if err != nil {
			return
		}
	}

	if regs.Found {
		for _, elem := range regs.Elem.Children {
			if elem.Name == NsPWG+":ScanRegion" {
				var reg ScanRegion
				reg, err = decodeScanRegion(elem)
				if err != nil {
					err = xmldoc.XMLErrWrap(regs.Elem, err)
					return
				}

				ss.ScanRegions = append(ss.ScanRegions, reg)
			}
		}
	}

	if fmt.Found {
		ss.DocumentFormat = optional.New(fmt.Elem.Text)
	}

	if fmtExt.Found {
		ss.DocumentFormatExt = optional.New(fmtExt.Elem.Text)
	}

	if content.Found {
		ss.ContentType, err = decodeOptional(
			content.Elem, decodeContentType)

		if err != nil {
			return
		}
	}

	if input.Found {
		ss.InputSource, err = decodeOptional(
			input.Elem, decodeInputSource)

		if err != nil {
			return
		}
	}

	if xres.Found {
		ss.XResolution, err = decodeOptional(
			xres.Elem, decodeNonNegativeInt)

		if err != nil {
			return
		}
	}

	if yres.Found {
		ss.YResolution, err = decodeOptional(
			yres.Elem, decodeNonNegativeInt)

		if err != nil {
			return
		}
	}

	if mode.Found {
		ss.ColorMode, err = decodeOptional(
			mode.Elem, decodeColorMode)

		if err != nil {
			return
		}
	}

	if space.Found {
		ss.ColorSpace, err = decodeOptional(
			space.Elem, decodeColorSpace)

		if err != nil {
			return
		}
	}

	if ccd.Found {
		ss.CcdChannel, err = decodeOptional(ccd.Elem, decodeCcdChannel)
		if err != nil {
			return
		}
	}

	if binrend.Found {
		ss.BinaryRendering, err = decodeOptional(
			binrend.Elem, decodeBinaryRendering)

		if err != nil {
			return
		}
	}

	if duplex.Found {
		ss.Duplex, err = decodeOptional(duplex.Elem, decodeBool)
		if err != nil {
			return
		}
	}

	if feed.Found {
		ss.FeedDirection, err = decodeOptional(
			feed.Elem, decodeFeedDirection)

		if err != nil {
			return
		}
	}

	if brightness.Found {
		ss.Brightness, err = decodeOptional(
			brightness.Elem, decodeNonNegativeInt)

		if err != nil {
			return
		}
	}

	if compression.Found {
		ss.CompressionFactor, err = decodeOptional(
			compression.Elem, decodeNonNegativeInt)

		if err != nil {
			return
		}
	}

	if contrast.Found {
		ss.Contrast, err = decodeOptional(
			contrast.Elem, decodeNonNegativeInt)

		if err != nil {
			return
		}
	}

	if gamma.Found {
		ss.Gamma, err = decodeOptional(gamma.Elem, decodeNonNegativeInt)
		if err != nil {
			return
		}
	}

	if highlight.Found {
		ss.Highlight, err = decodeOptional(
			highlight.Elem, decodeNonNegativeInt)

		if err != nil {
			return
		}
	}

	if noiseRemoval.Found {
		ss.NoiseRemoval, err = decodeOptional(
			noiseRemoval.Elem, decodeNonNegativeInt)

		if err != nil {
			return
		}
	}

	if shadow.Found {
		ss.Shadow, err = decodeOptional(
			shadow.Elem, decodeNonNegativeInt)

		if err != nil {
			return
		}
	}

	if sharpen.Found {
		ss.Sharpen, err = decodeOptional(
			sharpen.Elem, decodeNonNegativeInt)

		if err != nil {
			return
		}
	}

	if threshold.Found {
		ss.Threshold, err = decodeOptional(
			threshold.Elem, decodeNonNegativeInt)

		if err != nil {
			return
		}
	}

	if blankDetect.Found {
		ss.BlankPageDetection, err = decodeOptional(
			blankDetect.Elem, decodeBool)

		if err != nil {
			return
		}
	}

	if blankRemove.Found {
		ss.BlankPageDetectionAndRemoval, err = decodeOptional(
			blankRemove.Elem, decodeBool)

		if err != nil {
			return
		}
	}

	return
}
