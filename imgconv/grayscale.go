// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Package documentation

package imgconv

import (
	"image/color"
)

// grayscale implements the image filter that converts color
// images into the grayscale.
type grayscale struct {
	input Reader // Image source
}

// NewGrayScale creates a new image filter on a top of the existent
// [Reader].
//
// This filter converts color images into the Grayscale without
// changing the image [color.Model]. For example, the full-color
// RGB24 image will be converted into the RGB24 with all color
// channels set to the same value, effectively making it Grayscale.
func NewGrayScale(in Reader) Reader {
	// Bypass filter, if underlying color.Model is already Grayscale
	switch in.ColorModel() {
	case color.GrayModel, color.Gray16Model:
		return in
	}

	return &grayscale{input: in}
}

// ColorModel returns the [color.Model] of image being decoded.
func (gr *grayscale) ColorModel() color.Model {
	return gr.input.ColorModel()
}

// Size returns the image size.
func (gr *grayscale) Size() (wid, hei int) {
	return gr.input.Size()
}

// NewRow allocates a [Row] of the appropriate type and width for
// use with the [Reader.Read] function.
func (gr *grayscale) NewRow() Row {
	return gr.input.NewRow()
}

// Read returns the next image [Row].
// It returns the resulting row length, in pixels, or an error.
func (gr *grayscale) Read(row Row) (n int, err error) {
	// Read the next row from the underlying Reader
	n, err = gr.input.Read(row)
	if err != nil || n == 0 {
		return
	}

	// Convert to Grayscale, using the standard NTSC formula:
	//
	//   Y = R * 0.299 + G * 0.587 + B * 0.114
	switch row := row.(type) {
	case RowRGBA32:
		for i := range row {
			// Note:
			//
			//   19595 : 38470 : 7471 = 0.299 : 0.587 : 0.114
			//   19595 + 38470 + 7471 = 65536.
			c := row[i]
			r, g, b, _ := c.RGBA()
			y := uint8((19595*r + 38470*g + 7471*b + 1<<15) >> 24)

			c.R, c.G, c.B = y, y, y
			row[i] = c
		}
	case RowRGBA64:
		for i := range row {
			// Note:
			//
			//   19595 : 38470 : 7471 = 0.299 : 0.587 : 0.114
			//   19595 + 38470 + 7471 = 65536.
			c := row[i]
			r, g, b, _ := c.RGBA()
			y := uint16((19595*r + 38470*g + 7471*b + 1<<15) >> 16)

			c.R, c.G, c.B = y, y, y
			row[i] = c
		}

	case RowRGBAFP32:
		for i := range row {
			// RowRGBAFP32 uses the following layout:
			// R-G-B-A-R-G-B-A-...
			off := i * 4
			s := row[off : off+4]
			y := s[0]*0.299 + s[1]*0.587 + s[2]*0.114
			s[0], s[1], s[2] = y, y, y
		}
	}

	return
}

// Close closes the reader.
func (gr *grayscale) Close() {
	gr.input.Close()
}
