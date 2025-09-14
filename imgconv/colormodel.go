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

// colorModelFilter implements the image filter that converts image
// images into the different color.Model
type colorModelFilter struct {
	input Reader      // Image source
	model color.Model // Target color model
	wid   int         // Image width
	tmp   Row         // Conversion buffer
}

// NewColorModelFilter creates a new image filter on a top of the existent
// [Reader].
//
// This filter converts image's [color.Model].
func NewColorModelFilter(in Reader, model color.Model) Reader {
	if model == in.ColorModel() {
		return in
	}

	wid, _ := in.Size()

	return &colorModelFilter{
		input: in,
		model: model,
		wid:   wid,
		tmp:   NewRow(in.ColorModel(), wid),
	}
}

// ColorModel returns the [color.Model] of image being decoded.
func (cm *colorModelFilter) ColorModel() color.Model {
	return cm.model
}

// Size returns the image size.
func (cm *colorModelFilter) Size() (wid, hei int) {
	return cm.input.Size()
}

// NewRow allocates a [Row] of the appropriate type and width for
// use with the [Reader.Read] function.
func (cm *colorModelFilter) NewRow() Row {
	return NewRow(cm.model, cm.wid)
}

// Read returns the next image [Row].
// It returns the resulting row length, in pixels, or an error.
func (cm *colorModelFilter) Read(row Row) (n int, err error) {
	// Read the next row from the underlying Reader
	n, err = cm.input.Read(cm.tmp)
	if err == nil {
		row.Copy(cm.tmp)
	}

	return
}

// Close closes the reader.
func (cm *colorModelFilter) Close() {
	cm.input.Close()
}
