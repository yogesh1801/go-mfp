// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scan request

package abstract

// ScannerRequest specified scan request parameters
type ScannerRequest struct {
	// General parameters
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

	// Image processing parameters
	Brightness   int // Brightness
	Contrast     int // Contrast
	Gamma        int // Gamma (y=x^(1/g)
	Highlight    int // Image Highlight
	NoiseRemoval int // Noise removal level
	Shadow       int // The lower, the darger
	Sharpen      int // Image sharpen
	Threshold    int // ColorModeBinary+BinaryRenderingThreshold
	Compression  int // Lower num, better image
}
