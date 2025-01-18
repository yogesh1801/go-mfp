// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Package documentation

package escl

// ScanRegion defines the desired scan region
type ScanRegion struct {
	XOffset int   // Horizontal offset, 0-based
	YOffset int   // Vertical offset, 0-based
	Width   int   // Region width
	Height  int   // Region height
	Units   Units // Always ThreeHundredthsOfInches
}
