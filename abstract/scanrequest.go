// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scan request

package abstract

import "github.com/alexpevzner/mfp/util/optional"

// ScannerRequest specified scan request parameters
type ScannerRequest struct {
	// General parameters
	//
	// All parameters are optional. Use zero value to to indicate
	// that parameter is missed.
	Input           Input           // Input source (ADF/Platen etc)
	ADFMode         ADFMode         // For InputADF: Duplex/Simplex
	ColorMode       ColorMode       // Color mode (mono/color etc)
	Depth           Depth           // Image depth (8-bit/16-bit etc)
	BinaryRendering BinaryRendering // For 1-bit B&W (halftone/threshold
	CCDChannel      CCDChannel      // CCD channel to use
	DocumentFormat  string          // Requested document format
	Region          Region          // Scan region
	Resolution      Resolution      // Scanner resolution
	Intent          Intent          // Scan intent hint

	// Image processing parameters.
	//
	// As zero value is the legal value of these parameters,
	// we have to use optional.Val[int] to distinguish between
	// missed parameter and 0.
	Brightness   optional.Val[int] // Brightness
	Contrast     optional.Val[int] // Contrast
	Gamma        optional.Val[int] // Gamma (y=x^(1/g)
	Highlight    optional.Val[int] // Image Highlight
	NoiseRemoval optional.Val[int] // Noise removal level
	Shadow       optional.Val[int] // The lower, the darger
	Sharpen      optional.Val[int] // Image sharpen
	Threshold    optional.Val[int] // ColorModeBinary+BinaryRenderingThreshold
	Compression  optional.Val[int] // Lower num, better image
}
