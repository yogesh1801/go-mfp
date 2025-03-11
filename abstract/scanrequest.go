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
	Brightness   float64 // Brightness
	Contrast     float64 // Contrast
	Gamma        float64 // Gamma (y=x^(1/g)
	Highlight    float64 // Image Highlight
	NoiseRemoval float64 // Noise removal level
	Shadow       float64 // The lower, the darger
	Sharpen      float64 // Image sharpen
	Threshold    float64 // ColorModeBinary+BinaryRenderingThreshold
	Compression  float64 // Lower num, better image
}
