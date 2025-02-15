// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Package documentation

package escl

// SettingProfile defines a valid combination of scanning parameters.
//
// eSCL Technical Specification, 8.1.2.
type SettingProfile struct {
	ColorModes           []ColorMode          // Supported color modes
	DocumentFormats      []string             // MIME types of supported formats
	DocumentFormatsExt   []string             // eSCL 2.1+
	SupportedResolutions SupportedResolutions // Supported resolutions
	ColorSpaces          []ColorSpace         // Supported color spaces
	CcdChannels          []CcdChannel         // Supported CCD channels
	BinaryRenderings     []BinaryRendering    // Supported bin renderings
}
