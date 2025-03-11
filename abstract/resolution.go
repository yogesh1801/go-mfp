// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Discrete resolution

package abstract

// Resolution specifies a discrete scanner resolution.
type Resolution struct {
	XResolution int // X resolution, DPI
	YResolution int // Y resolution, DPI
}

// ResolutionRange specifies a range of scanner resolutions.
type ResolutionRange struct {
	XMin, XMax, XStep int // X resolution range, DPI
	YMin, YMax, YStep int // Y resolution range, DPI
}
