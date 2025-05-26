// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Image scaler

package imgconv

import (
	"image"

	"golang.org/x/image/draw"
)

// NewScaler creates a new image resize filter on a top of the
// existent [Decoder].
//
// This filter scales the input image into the new dimensions,
// defined by the wid and hei parameters.
func NewScaler(in Decoder, wid, hei int) Decoder {
	oldwid, oldhei := in.Size()

	// Bypass the filter, if image dimensions doesn't change
	if oldwid == wid && oldhei == hei {
		return in
	}

	// Use transformer filter to do the job
	tranform := func(target draw.Image, source image.Image) {
		draw.BiLinear.Scale(
			target, target.Bounds(),
			source, source.Bounds(),
			draw.Over, nil)
	}

	return NewTransformer(in, wid, hei, in.ColorModel(), tranform)
}
