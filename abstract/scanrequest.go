// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scan request

package abstract

import (
	"bytes"
	"fmt"

	"github.com/OpenPrinting/go-mfp/util/optional"
)

// ScannerRequest specified scan request parameters
type ScannerRequest struct {
	// General parameters
	//
	// All parameters are optional. Use zero value to to indicate
	// that parameter is missed.
	Input           Input           // Input source (ADF/Platen etc)
	ADFMode         ADFMode         // For InputADF: Duplex/Simplex
	ColorMode       ColorMode       // Color mode (mono/color etc)
	ColorDepth      ColorDepth      // Image depth (8-bit/16-bit etc)
	BinaryRendering BinaryRendering // For 1-bit B&W (halftone/threshold)
	CCDChannel      CCDChannel      // CCD channel to use (mono/gray only)
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

// MarshalLog formats [ScannerRequest] for logging.
// It implements the [log.Marshaler] interface.
func (req *ScannerRequest) MarshalLog() []byte {
	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "Input:           %s\n", req.Input)
	if req.ADFMode != ADFModeUnset {
		fmt.Fprintf(buf, "ADFMode:         %s\n", req.ADFMode)
	}
	if req.ColorMode != ColorModeUnset {
		fmt.Fprintf(buf, "ColorMode:       %s\n", req.ColorMode)
	}
	if req.ColorDepth != ColorDepthUnset {
		fmt.Fprintf(buf, "ColorDepth:      %s\n", req.ColorDepth)
	}
	if req.BinaryRendering != BinaryRenderingUnset {
		fmt.Fprintf(buf, "BinaryRendering: %s\n", req.BinaryRendering)
	}
	if req.CCDChannel != CCDChannelUnset {
		fmt.Fprintf(buf, "CCDChannel:      %s\n", req.CCDChannel)
	}
	fmt.Fprintf(buf, "DocumentFormat:  %q\n", req.DocumentFormat)
	if !req.Region.IsZero() {
		fmt.Fprintf(buf, "Region:          %s\n", req.Region)
	}
	if !req.Resolution.IsZero() {
		fmt.Fprintf(buf, "Resolution:      %s\n", req.Resolution)
	}
	fmt.Fprintf(buf, "Intent:          %s\n", req.Intent)
	if req.Brightness != nil {
		fmt.Fprintf(buf, "Brightness:      %d\n", *req.Brightness)
	}
	if req.Contrast != nil {
		fmt.Fprintf(buf, "Contrast:        %d\n", *req.Contrast)
	}
	if req.Gamma != nil {
		fmt.Fprintf(buf, "Gamma:           %d\n", *req.Gamma)
	}
	if req.Highlight != nil {
		fmt.Fprintf(buf, "Highlight:       %d\n", *req.Highlight)
	}
	if req.NoiseRemoval != nil {
		fmt.Fprintf(buf, "NoiseRemoval:    %d\n", *req.NoiseRemoval)
	}
	if req.Shadow != nil {
		fmt.Fprintf(buf, "Shadow:          %d\n", *req.Shadow)
	}
	if req.Sharpen != nil {
		fmt.Fprintf(buf, "Sharpen:         %d\n", *req.Sharpen)
	}
	if req.Threshold != nil {
		fmt.Fprintf(buf, "Threshold:       %d\n", *req.Threshold)
	}
	if req.Compression != nil {
		fmt.Fprintf(buf, "Compression:     %d\n", *req.Compression)
	}

	return buf.Bytes()
}
