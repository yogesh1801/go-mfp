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

	"github.com/OpenPrinting/go-mfp/util/generic"
)

// Row represents a single row of image
type Row interface {
	// Width returns the row width, in pixels.
	Width() int

	// At returns the pixel at the specified position.
	At(x int) color.Color

	// Set sets the pixel at the specified position.
	Set(x int, c color.Color)

	// Slice returns a [low:high] sub-slice of the original Row.
	Slice(low, high int) Row

	// Fill fills Row with the pixels of the specified color
	Fill(c color.Color)

	// Copy copies content of the r2 into the receiver Row.
	// If they have a different color model, pixels are converted.
	// Rows may be of the different size and may overlap.
	//
	// It returns number of pixels copied.
	Copy(r2 Row) int
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

// RowFP extends the [Row] interface with the operations, required
// for image processing.
type RowFP interface {
	Row

	// MultiplyAccumulate performs the multiply-accumulate operation on
	// the entire row:
	//
	//	r += r2 * w
	//
	// Rows must be of the same type and size.
	MultiplyAccumulate(r2 Row, w float32)
}

// NewRowFP returns the new [Row] with the [RowGrayFP32] or [RowRGBAFP32],
// compatible with the [color.Model] (grayscale or RGBA).
//
// The following color models are supported:
//   - color.GrayModel
//   - color.Gray16Model
//   - color.RGBAModel
//   - color.RGBA64Model
//
// For unknown (unsupported) model nil is returned.
func NewRowFP(model color.Model, width int) (row RowFP) {
	switch model {
	case color.GrayModel, color.Gray16Model:
		row = make(RowGrayFP32, width)
	case color.RGBAModel, color.RGBA64Model:
		row = make(RowRGBAFP32, width)
	}

	return
}

// RowEmpty represents an empty row
type RowEmpty struct{}

// Width returns the row width, in pixels.
func (r RowEmpty) Width() int {
	return 0
}

// At returns the pixel at the specified position as [color.Color].
func (r RowEmpty) At(x int) color.Color {
	return color.White
}

// Set sets the pixel at the specified position.
func (r RowEmpty) Set(x int, c color.Color) {
}

// Slice returns a [low:high] sub-slice of the original Row.
func (r RowEmpty) Slice(low, high int) Row {
	return RowEmpty{}
}

// Fill fills Row with the pixels of the specified color
func (r RowEmpty) Fill(c color.Color) {
}

// Copy copies content of the r2 into the receiver Row.
func (r RowEmpty) Copy(r2 Row) int {
	return 0
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

// Set sets the pixel at the specified position.
func (r RowGray8) Set(x int, c color.Color) {
	c2, ok := c.(color.Gray)
	if !ok {
		c2 = color.GrayModel.Convert(c).(color.Gray)
	}

	r[x] = c2
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

// Copy copies content of the r2 into the receiver Row.
func (r RowGray8) Copy(r2 Row) int {
	if r2, ok := r2.(RowGray8); ok {
		return copy(r, r2)
	}

	wid := generic.Min(r.Width(), r2.Width())
	for x := 0; x < wid; x++ {
		r[x] = color.GrayModel.Convert(r2.At(x)).(color.Gray)
	}

	return wid
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

// Set sets the pixel at the specified position.
func (r RowGray16) Set(x int, c color.Color) {
	c2, ok := c.(color.Gray16)
	if !ok {
		c2 = color.Gray16Model.Convert(c).(color.Gray16)
	}

	r[x] = c2
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

// Copy copies content of the r2 into the receiver Row.
func (r RowGray16) Copy(r2 Row) int {
	if r2, ok := r2.(RowGray16); ok {
		return copy(r, r2)
	}

	wid := generic.Min(r.Width(), r2.Width())
	for x := 0; x < wid; x++ {
		r[x] = color.Gray16Model.Convert(r2.At(x)).(color.Gray16)
	}

	return wid
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

// Set sets the pixel at the specified position.
func (r RowRGBA32) Set(x int, c color.Color) {
	c2, ok := c.(color.RGBA)
	if !ok {
		c2 = color.RGBAModel.Convert(c).(color.RGBA)
	}

	r[x] = c2
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

// Copy copies content of the r2 into the receiver Row.
func (r RowRGBA32) Copy(r2 Row) int {
	if r2, ok := r2.(RowRGBA32); ok {
		return copy(r, r2)
	}

	wid := generic.Min(r.Width(), r2.Width())
	for x := 0; x < wid; x++ {
		r[x] = color.RGBAModel.Convert(r2.At(x)).(color.RGBA)
	}

	return wid
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

// Set sets the pixel at the specified position.
func (r RowRGBA64) Set(x int, c color.Color) {
	c2, ok := c.(color.RGBA64)
	if !ok {
		c2 = color.RGBA64Model.Convert(c).(color.RGBA64)
	}

	r[x] = c2
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

// Copy copies content of the r2 into the receiver Row.
func (r RowRGBA64) Copy(r2 Row) int {
	if r2, ok := r2.(RowRGBA64); ok {
		return copy(r, r2)
	}

	wid := generic.Min(r.Width(), r2.Width())
	for x := 0; x < wid; x++ {
		r[x] = color.RGBA64Model.Convert(r2.At(x)).(color.RGBA64)
	}

	return wid
}

// RowGrayFP32 represents a row of the grayscale image as a sequence of
// the 32-bit floating point numbers in range [0...1.0].
type RowGrayFP32 []float32

// Width returns the row width, in pixels.
func (r RowGrayFP32) Width() int {
	return len(r)
}

// At returns pixel at the specified position as [color.Color].
func (r RowGrayFP32) At(x int) color.Color {
	return r.Gray16At(x)
}

// GrayAt returns pixel at the specified position as [color.Gray].
func (r RowGrayFP32) GrayAt(x int) color.Gray {
	c := r[x]
	if c > 1.0 {
		return color.Gray{Y: 0xff}
	}
	return color.Gray{Y: uint8(c * 0xff)}
}

// Gray16At returns pixel at the specified position as [color.Gray16].
func (r RowGrayFP32) Gray16At(x int) color.Gray16 {
	c := r[x]
	if c > 1.0 {
		return color.Gray16{Y: 0xffff}
	}
	return color.Gray16{Y: uint16(c * 0xffff)}
}

// Set sets the pixel at the specified position.
func (r RowGrayFP32) Set(x int, c color.Color) {
	switch c := c.(type) {
	case color.Gray:
		r[x] = float32(c.Y) / 0xff
	case color.Gray16:
		r[x] = float32(c.Y) / 0xffff
	}
	c2 := color.Gray16Model.Convert(c).(color.Gray16)
	r[x] = float32(c2.Y) / 0xffff
}

// Slice returns a [low:high] sub-slice of the original Row.
func (r RowGrayFP32) Slice(low, high int) Row {
	return r[low:high]
}

// Fill fills Row with the pixels of the specified color
func (r RowGrayFP32) Fill(c color.Color) {
	var fill float32
	switch c := c.(type) {
	case color.Gray:
		fill = float32(c.Y) / 0xff
	case color.Gray16:
		fill = float32(c.Y) / 0xffff
	default:
		c2 := color.Gray16Model.Convert(c).(color.Gray16)
		fill = float32(c2.Y) / 0xffff
	}

	for i := range r {
		r[i] = fill
	}
}

// Copy copies content of the r2 into the receiver Row.
func (r RowGrayFP32) Copy(r2 Row) int {
	wid := generic.Min(r.Width(), r2.Width())
	switch r2 := r2.(type) {
	case RowGrayFP32:
		return copy(r, r2)
	case RowGray8:
		for x := 0; x < wid; x++ {
			r[x] = float32(r2[x].Y) / 0xff
		}
	case RowGray16:
		for x := 0; x < wid; x++ {
			r[x] = float32(r2[x].Y) / 0xffff
		}
	default:
		for x := 0; x < wid; x++ {
			c2 := color.RGBA64Model.Convert(r2.At(x)).(color.Gray16)
			r[x] = float32(c2.Y) / 0xffff
		}
	}

	return wid
}

// MultiplyAccumulate performs the multiply-accumulate operation on
// the entire row:
//
//	r += r2 * w
//
// Rows must be of the same type and size.
func (r RowGrayFP32) MultiplyAccumulate(r2 Row, w float32) {
	src := r2.(RowGrayFP32)
	for i := range r {
		r[i] += src[i] * w
	}
}

// RowRGBAFP32 represents a row of the RGBA image as a sequence of
// the 32-bit floating point numbers in range [0...1.0].
// The sequence is ordered as follows: R-G-B-A-R-G-B-A-...
type RowRGBAFP32 []float32

// Width returns the row width, in pixels.
func (r RowRGBAFP32) Width() int {
	return len(r) / 4
}

// At returns pixel at the specified position as [color.Color].
func (r RowRGBAFP32) At(x int) color.Color {
	return r.RGBAAt(x)
}

// RGBAAt returns pixel at the specified position as [color.RGBA].
func (r RowRGBAFP32) RGBAAt(x int) color.RGBA {
	off := x * 4
	s := r[off : off+4]

	return color.RGBA{
		R: uint8(generic.Min(1.0, s[0]) * 0xff),
		G: uint8(generic.Min(1.0, s[1]) * 0xff),
		B: uint8(generic.Min(1.0, s[2]) * 0xff),
		A: uint8(generic.Min(1.0, s[3]) * 0xff),
	}
}

// RGBA16At returns pixel at the specified position as [color.RGBA64].
func (r RowRGBAFP32) RGBA16At(x int) color.RGBA64 {
	off := x * 4
	s := r[off : off+4]

	return color.RGBA64{
		R: uint16(generic.Min(1.0, s[0]) * 0xffff),
		G: uint16(generic.Min(1.0, s[1]) * 0xffff),
		B: uint16(generic.Min(1.0, s[2]) * 0xffff),
		A: uint16(generic.Min(1.0, s[3]) * 0xffff),
	}
}

// Set sets the pixel at the specified position.
func (r RowRGBAFP32) Set(x int, c color.Color) {
	off := x * 4
	s := r[off : off+4]

	switch c := c.(type) {
	case color.RGBA:
		s[0] = float32(c.R) / 0xff
		s[1] = float32(c.G) / 0xff
		s[2] = float32(c.B) / 0xff
		s[3] = float32(c.A) / 0xff
	case color.RGBA64:
		s[0] = float32(c.R) / 0xffff
		s[1] = float32(c.G) / 0xffff
		s[2] = float32(c.B) / 0xffff
		s[3] = float32(c.A) / 0xffff
	}

	c2 := color.RGBA64Model.Convert(c).(color.RGBA64)
	s[0] = float32(c2.R) / 0xffff
	s[1] = float32(c2.G) / 0xffff
	s[2] = float32(c2.B) / 0xffff
	s[3] = float32(c2.A) / 0xffff
}

// Slice returns a [low:high] sub-slice of the original Row.
func (r RowRGBAFP32) Slice(low, high int) Row {
	return r[low*4 : high*4]
}

// Fill fills Row with the pixels of the specified color
func (r RowRGBAFP32) Fill(c color.Color) {
	var fill [4]float32
	switch c := c.(type) {
	case color.RGBA:
		fill[0] = float32(c.R) / 0xff
		fill[1] = float32(c.G) / 0xff
		fill[2] = float32(c.B) / 0xff
		fill[3] = float32(c.A) / 0xff
	case color.RGBA64:
		fill[0] = float32(c.R) / 0xffff
		fill[1] = float32(c.G) / 0xffff
		fill[2] = float32(c.B) / 0xffff
		fill[3] = float32(c.A) / 0xffff
	default:
		R, G, B, A := c.RGBA()
		fill[0] = float32(R) / 0xffff
		fill[1] = float32(G) / 0xffff
		fill[2] = float32(B) / 0xffff
		fill[3] = float32(A) / 0xffff
	}

	for i := 0; i < len(r); i += 4 {
		copy(r[i:], fill[:])
	}
}

// Copy copies content of the r2 into the receiver Row.
func (r RowRGBAFP32) Copy(r2 Row) int {
	if r2, ok := r2.(RowRGBAFP32); ok {
		return copy(r, r2)
	}

	wid := generic.Min(r.Width(), r2.Width())
	switch r2 := r2.(type) {
	case RowRGBAFP32:
		return copy(r, r2)
	case RowRGBA32:
		for x := 0; x < wid; x++ {
			off := x * 4
			s := r[off : off+4]
			c := r2[x]
			s[0] = float32(c.R) / 0xff
			s[1] = float32(c.G) / 0xff
			s[2] = float32(c.B) / 0xff
			s[3] = float32(c.A) / 0xff
		}
	case RowRGBA64:
		for x := 0; x < wid; x++ {
			off := x * 4
			s := r[off : off+4]
			c := r2[x]
			s[0] = float32(c.R) / 0xffff
			s[1] = float32(c.G) / 0xffff
			s[2] = float32(c.B) / 0xffff
			s[3] = float32(c.A) / 0xffff
		}
	default:
		for x := 0; x < wid; x++ {
			off := x * 4
			s := r[off : off+4]
			c := color.RGBA64Model.Convert(r2.At(x)).(color.RGBA64)
			s[0] = float32(c.R) / 0xffff
			s[1] = float32(c.G) / 0xffff
			s[2] = float32(c.B) / 0xffff
			s[3] = float32(c.A) / 0xffff
		}
	}

	return wid
}

// MultiplyAccumulate performs the multiply-accumulate operation on
// the entire row:
//
//	r += r2 * w
//
// Rows must be of the same type and size.
func (r RowRGBAFP32) MultiplyAccumulate(r2 Row, w float32) {
	src := r2.(RowRGBAFP32)
	for i := range r {
		r[i] += src[i] * w
	}
}
