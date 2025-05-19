// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Package documentation

package imgconv

import (
	"image"
	"image/color"
	"io"
)

// resizer implements an image resizer.
type resizer struct {
	input    Decoder         // Image source
	rect     image.Rectangle // Clipping region
	fill     color.Color     // Filler color
	wid, hei int             // Resized image size
	y        int             // Current y coordinate
}

// NewResizer creates a new image resize filter on a top of the
// existent [Decoder].
//
// Resizer works by either clipping or expanding image to fit
// the specified region.
//
// Resizer implements the [Decoder] interface, which allows
// to build a chain of image filters.
//
// When resizer is closed, its input Decoder is also closed.
func NewResizer(in Decoder, rect image.Rectangle) Decoder {
	rect = rect.Canon()
	wid, hei := in.Size()

	if rect.Min.X == 0 && rect.Min.Y == 0 &&
		rect.Dx() == wid && rect.Dy() == hei {
		return in
	}

	model := in.ColorModel()
	rsz := &resizer{
		input: in,
		rect:  rect.Canon(),
		wid:   rect.Dx(),
		hei:   rect.Dy(),
		fill:  model.Convert(color.White),
	}

	return rsz
}

// ColorModel returns the [color.Model] of image being decoded.
func (rsz *resizer) ColorModel() color.Model {
	return rsz.input.ColorModel()
}

// Size returns the image size.
func (rsz *resizer) Size() (wid, hei int) {
	return rsz.rect.Dx(), rsz.rect.Dy()
}

// NewRow allocates a [Row] of the appropriate type and width for
// use with the [Decoder.Read] function.
func (rsz *resizer) NewRow() Row {
	return NewRow(rsz.ColorModel(), rsz.rect.Dx())
}

// Read returns the next image [Row].
// The Row type must match the [Decoder]'s [color.Model].
//
// It returns the resulting row length, in pixels, or an error.
func (rsz *resizer) Read(row Row) (int, error) {
	_, end := rsz.Size()

	switch {
	case rsz.y == end:
		return 0, io.EOF
	case rsz.y < rsz.rect.Min.Y || rsz.y >= rsz.rect.Max.Y:
		row.Fill(rsz.fill)
	}

	return 0, io.EOF
}

// Close closes the decoder
func (rsz *resizer) Close() {
	rsz.input.Close()
}
