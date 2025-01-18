// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scan settings (scan request parameters)

package escl

import "github.com/alexpevzner/mfp/optional"

// ScanSettings defines the set of parameters for scan request.
//
// eSCL Technical Specification, 7.
type ScanSettings struct {
	// Version is the only required parameter
	Version Version // eSCL protocol version

	// General parameters
	Intent            optional.Val[Intent]          // Scan intent
	ScanRegioons      []ScanRegion                  // List of scan regions
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
