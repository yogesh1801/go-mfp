// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scan settings (scan request parameters)

package escl

import (
	"strconv"

	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// ScanSettings is the client request, that defines the set of scan parameters.
//
// eSCL Technical Specification, 7.
//
// POST /{root}/ScanJobs       - to start scanning
// PUT  /{root}/ScanBufferInfo - to estimate actual scanning parameters
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
	CCDChannel        optional.Val[CCDChannel]      // Desired CCD channel
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
	ret *ScanSettings, err error) {

	defer func() { err = xmldoc.XMLErrWrap(root, err) }()

	var ss ScanSettings

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
	mode := xmldoc.Lookup{Name: NsScan + ":ColorMode"}
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
		ss.CCDChannel, err = decodeOptional(ccd.Elem, decodeCCDChannel)
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

	ret = &ss
	return
}

// ToXML generates XML tree for the [ScanSettings].
func (ss *ScanSettings) ToXML() xmldoc.Element {
	elm := xmldoc.Element{
		Name: NsScan + ":ScanSettings",
		Children: []xmldoc.Element{
			ss.Version.toXML(NsPWG + ":Version"),
		},
	}

	if ss.Intent != nil {
		chld := (*ss.Intent).toXML(NsScan + ":Intent")
		elm.Children = append(elm.Children, chld)
	}

	if ss.ScanRegions != nil {
		chld := xmldoc.Element{Name: NsPWG + ":ScanRegions"}
		for _, reg := range ss.ScanRegions {
			chld2 := reg.toXML(NsPWG + ":ScanRegion")
			chld.Children = append(chld.Children, chld2)
		}
		elm.Children = append(elm.Children, chld)
	}

	if ss.DocumentFormat != nil {
		chld := xmldoc.WithText(NsPWG+":DocumentFormat",
			*ss.DocumentFormat)
		elm.Children = append(elm.Children, chld)
	}

	if ss.DocumentFormatExt != nil {
		chld := xmldoc.WithText(NsScan+":DocumentFormatExt",
			*ss.DocumentFormatExt)
		elm.Children = append(elm.Children, chld)
	}

	if ss.ContentType != nil {
		chld := (*ss.ContentType).toXML(NsPWG + ":ContentType")
		elm.Children = append(elm.Children, chld)
	}

	if ss.InputSource != nil {
		chld := (*ss.InputSource).toXML(NsPWG + ":InputSource")
		elm.Children = append(elm.Children, chld)
	}

	if ss.XResolution != nil {
		chld := xmldoc.WithText(NsScan+":XResolution",
			strconv.FormatUint(uint64(*ss.XResolution), 10))
		elm.Children = append(elm.Children, chld)
	}

	if ss.YResolution != nil {
		chld := xmldoc.WithText(NsScan+":YResolution",
			strconv.FormatUint(uint64(*ss.YResolution), 10))
		elm.Children = append(elm.Children, chld)
	}

	if ss.ColorMode != nil {
		chld := (*ss.ColorMode).toXML(NsScan + ":ColorMode")
		elm.Children = append(elm.Children, chld)
	}

	if ss.ColorSpace != nil {
		chld := (*ss.ColorSpace).toXML(NsScan + ":ColorSpace")
		elm.Children = append(elm.Children, chld)
	}

	if ss.CCDChannel != nil {
		chld := (*ss.CCDChannel).toXML(NsScan + ":CcdChannel")
		elm.Children = append(elm.Children, chld)
	}

	if ss.BinaryRendering != nil {
		chld := (*ss.BinaryRendering).toXML(NsScan + ":BinaryRendering")
		elm.Children = append(elm.Children, chld)
	}

	if ss.Duplex != nil {
		chld := xmldoc.WithText(NsScan+":Duplex",
			strconv.FormatBool(*ss.Duplex))
		elm.Children = append(elm.Children, chld)
	}

	if ss.FeedDirection != nil {
		chld := (*ss.FeedDirection).toXML(NsScan + ":FeedDirection")
		elm.Children = append(elm.Children, chld)
	}

	if ss.Brightness != nil {
		chld := xmldoc.WithText(NsScan+":Brightness",
			strconv.FormatUint(uint64(*ss.Brightness), 10))
		elm.Children = append(elm.Children, chld)
	}

	if ss.CompressionFactor != nil {
		chld := xmldoc.WithText(NsScan+":CompressionFactor",
			strconv.FormatUint(uint64(*ss.CompressionFactor), 10))
		elm.Children = append(elm.Children, chld)
	}

	if ss.Contrast != nil {
		chld := xmldoc.WithText(NsScan+":Contrast",
			strconv.FormatUint(uint64(*ss.Contrast), 10))
		elm.Children = append(elm.Children, chld)
	}

	if ss.Gamma != nil {
		chld := xmldoc.WithText(NsScan+":Gamma",
			strconv.FormatUint(uint64(*ss.Gamma), 10))
		elm.Children = append(elm.Children, chld)
	}

	if ss.Highlight != nil {
		chld := xmldoc.WithText(NsScan+":Highlight",
			strconv.FormatUint(uint64(*ss.Highlight), 10))
		elm.Children = append(elm.Children, chld)
	}

	if ss.NoiseRemoval != nil {
		chld := xmldoc.WithText(NsScan+":NoiseRemoval",
			strconv.FormatUint(uint64(*ss.NoiseRemoval), 10))
		elm.Children = append(elm.Children, chld)
	}

	if ss.Shadow != nil {
		chld := xmldoc.WithText(NsScan+":Shadow",
			strconv.FormatUint(uint64(*ss.Shadow), 10))
		elm.Children = append(elm.Children, chld)
	}

	if ss.Sharpen != nil {
		chld := xmldoc.WithText(NsScan+":Sharpen",
			strconv.FormatUint(uint64(*ss.Sharpen), 10))
		elm.Children = append(elm.Children, chld)
	}

	if ss.Threshold != nil {
		chld := xmldoc.WithText(NsScan+":Threshold",
			strconv.FormatUint(uint64(*ss.Threshold), 10))
		elm.Children = append(elm.Children, chld)
	}

	if ss.BlankPageDetection != nil {
		chld := xmldoc.WithText(NsScan+":BlankPageDetection",
			strconv.FormatBool(*ss.BlankPageDetection))
		elm.Children = append(elm.Children, chld)
	}

	if ss.BlankPageDetectionAndRemoval != nil {
		chld := xmldoc.WithText(NsScan+":BlankPageDetectionAndRemoval",
			strconv.FormatBool(*ss.BlankPageDetectionAndRemoval))
		elm.Children = append(elm.Children, chld)
	}

	return elm
}
