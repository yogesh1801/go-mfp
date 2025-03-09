// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Discrete resolution

package abstract

import "github.com/alexpevzner/mfp/util/optional"

// Resolution specifies a discrete scanner resolution.
type Resolution struct {
	XResolution optional.Val[int] // X resolution, DPI
	YResolution optional.Val[int] // Y resolution, DPI
}
