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

	// Slice returns a [low:high] sub-slice of the original Row.
	Slice(low, high int) Row

	// Fill fills Row with the pixels of the specified color
	Fill(c color.Color)
}

// NewRow returns the new [Row] of the specified width and [color.Model].
// The following color models are supported:
//   - color.GrayModel
//   - color.Gray16Model
//   - color.RGBAModel
//   - color.RGBA64Model
//
// For unknown (unsupported) model nil is returned.
func NewRow(model color.Model, width int) (row Row) {
	switch model {
	case color.GrayModel:
		row = make(RowGray8, width)
	case color.Gray16Model:
		row = make(RowGray16, width)
	case color.RGBAModel:
		row = make(RowRGBA32, width)
	case color.RGBA64Model:
		row = make(RowRGBA64, width)
	}

	return
}

// RowEmpty represents an empty row
type RowEmpty struct{}

// Width returns the row width, in pixels.
func (r RowEmpty) Width() int {
	return 0
}

// At returns pixel at the specified position as [color.Color].
func (r RowEmpty) At(x int) color.Color {
	return color.White
}

// Slice returns a [low:high] sub-slice of the original Row.
func (r RowEmpty) Slice(low, high int) Row {
	return RowEmpty{}
}

// Fill fills Row with the pixels of the specified color
func (r RowEmpty) Fill(c color.Color) {
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

// Slice returns a [low:high] sub-slice of the original Row.
func (r RowGray8) Slice(low, high int) Row {
	return r[low:high]
}

// Fill fills Row with the pixels of the specified color
func (r RowGray8) Fill(c color.Color) {
	c2, ok := c.(color.Gray)
	if !ok {
		c2 = color.GrayModel.Convert(c).(color.Gray)
	}

	for i := range r {
		r[i] = c2
	}
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

// Slice returns a [low:high] sub-slice of the original Row.
func (r RowGray16) Slice(low, high int) Row {
	return r[low:high]
}

// Fill fills Row with the pixels of the specified color
func (r RowGray16) Fill(c color.Color) {
	c2, ok := c.(color.Gray16)
	if !ok {
		c2 = color.Gray16Model.Convert(c).(color.Gray16)
	}

	for i := range r {
		r[i] = c2
	}
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

// Slice returns a [low:high] sub-slice of the original Row.
func (r RowRGBA32) Slice(low, high int) Row {
	return r[low:high]
}

// Fill fills Row with the pixels of the specified color
func (r RowRGBA32) Fill(c color.Color) {
	c2, ok := c.(color.RGBA)
	if !ok {
		c2 = color.RGBAModel.Convert(c).(color.RGBA)
	}

	for i := range r {
		r[i] = c2
	}
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

// Slice returns a [low:high] sub-slice of the original Row.
func (r RowRGBA64) Slice(low, high int) Row {
	return r[low:high]
}

// Fill fills Row with the pixels of the specified color
func (r RowRGBA64) Fill(c color.Color) {
	c2, ok := c.(color.RGBA64)
	if !ok {
		c2 = color.RGBA64Model.Convert(c).(color.RGBA64)
	}

	for i := range r {
		r[i] = c2
	}
}

// RGBA64At returns pixel at the specified position as [color.RGBA64].
func (r RowRGBA64) RGBA64At(x int) color.RGBA64 {
	return r[x]
}
