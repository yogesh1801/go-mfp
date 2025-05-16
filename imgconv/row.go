// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Package documentation

package imgconv

import "image/color"

// Row represents a single row of image
type Row interface {
	// Width returns the row width, in pixels.
	Width() int

	// At returns pixel at the specified position.
	At(x int) color.Color
}

// RowGray8 represents a row of 8-bit grayscale image.
type RowGray8 []color.Gray

// Width returns the row width, in pixels.
func (r RowGray8) Width() int {
	return len(r)
}

// At returns pixel at the specified position as [color.Color].
func (r RowGray8) At(x int) color.Color {
	return r.GrayAt(x)
}

// GrayAt returns pixel at the specified position as [color.Gray].
func (r RowGray8) GrayAt(x int) color.Gray {
	return r[x]
}

// RowGray16 represents a row of 16-bit grayscale image.
type RowGray16 []color.Gray16

// Width returns the row width, in pixels.
func (r RowGray16) Width() int {
	return len(r)
}

// At returns pixel at the specified position as [color.Color].
func (r RowGray16) At(x int) color.Color {
	return r.Gray16At(x)
}

// Gray16At returns pixel at the specified position as [color.Gray16].
func (r RowGray16) Gray16At(x int) color.Gray16 {
	return r[x]
}

// RowRGBA32 represents a row of 8-bit RGBA image.
type RowRGBA32 []color.RGBA

// Width returns the row width, in pixels.
func (r RowRGBA32) Width() int {
	return len(r)
}

// At returns pixel at the specified position as [color.Color].
func (r RowRGBA32) At(x int) color.Color {
	return r.RGBAAt(x)
}

// RGBAAt returns pixel at the specified position as [color.RGBA].
func (r RowRGBA32) RGBAAt(x int) color.RGBA {
	return r[x]
}

// RowRGBA64 represents a row of 16-bit RGBA image.
type RowRGBA64 []color.RGBA64

// Width returns the row width, in pixels.
func (r RowRGBA64) Width() int {
	return len(r)
}

// At returns pixel at the specified position as [color.Color].
func (r RowRGBA64) At(x int) color.Color {
	return r.RGBA64At(x)
}

// RGBA64At returns pixel at the specified position as [color.RGBA64].
func (r RowRGBA64) RGBA64At(x int) color.RGBA64 {
	return r[x]
}
