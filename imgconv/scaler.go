// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Image scaler

package imgconv

import (
	"image/color"
	"io"
)

// scaler implements image scaler
type scaler struct {
	input    Decoder      // Image source
	err      error        // I/O error
	wid, hei int          // Scaled image size
	xcoeffs  []scaleCoeff // Horizontal scaling coefficients
	ycoeffs  []scaleCoeff // Remaining vertical scaling coefficients
	tmpin    RowFP        // Input buffer, for reading from scaler.input
	tmpout   RowFP        // Output buffer
	history  []RowFP      // Scaled source rows: 0 - latest, 1 - previous etc
	srcy     int          // Latest source row y-coordinate
}

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

	// Create a scaler
	model := in.ColorModel()
	scl := &scaler{
		input:   in,
		wid:     wid,
		hei:     hei,
		xcoeffs: makeScaleCoefficients(oldwid, wid),
		ycoeffs: makeScaleCoefficients(oldhei, hei),
		tmpin:   NewRowFP(model, oldwid),
		tmpout:  NewRowFP(model, wid),
		srcy:    -1,
	}

	// Populate scaler.rows
	hstsize := scaleCoefficientsHistorySize(scl.ycoeffs)
	scl.history = make([]RowFP, hstsize+1)

	for i := range scl.history {
		scl.history[i] = NewRowFP(model, wid)
	}

	return scl
}

// ColorModel returns the [color.Model] of image being decoded.
func (scl *scaler) ColorModel() color.Model {
	return scl.input.ColorModel()
}

// Size returns the image size.
func (scl *scaler) Size() (wid, hei int) {
	return scl.wid, scl.hei
}

// NewRow allocates a [Row] of the appropriate type and width for
// use with the [Decoder.Read] function.
func (scl *scaler) NewRow() Row {
	return NewRow(scl.ColorModel(), scl.wid)
}

// Read returns the next image [Row].
// The Row type must match the [Decoder]'s [color.Model].
//
// It returns the resulting row length, in pixels, or an error.
func (scl *scaler) Read(row Row) (int, error) {
	if len(scl.ycoeffs) == 0 {
		return 0, io.EOF // All rows processed
	}

	scl.tmpout.ZeroFill()

	D := scl.ycoeffs[0].D
	for len(scl.ycoeffs) > 0 && scl.ycoeffs[0].D == D {
		sc := scl.ycoeffs[0]
		srow := scl.hstrow(sc.S)
		scl.tmpout.MultiplyAccumulate(srow, sc.W)
		scl.ycoeffs = scl.ycoeffs[1:]
	}

	return row.Copy(scl.tmpout), nil
}

// Close closes the decoder
func (scl *scaler) Close() {
	scl.input.Close()
}

// hstrow returns the source row at y-position, relative to scl.input
func (scl *scaler) hstrow(y int) RowFP {
	for scl.err == nil && scl.srcy < y {
		_, scl.err = scl.input.Read(scl.tmpin)
		if scl.err == nil {
			scl.hstshift()
			scl.srcy++

			scl.history[0].ZeroFill()
			scl.history[0].scale(scl.tmpin, scl.xcoeffs)
		}
	}

	return scl.history[scl.srcy-y]
}

// hstshift shifts (rotates) scaler.history, so the next row can
// be read into the scaler.history[0]
func (scl *scaler) hstshift() {
	if l := len(scl.history); l > 1 {
		tmp := scl.history[l-1]
		copy(scl.history[1:], scl.history)
		scl.history[0] = tmp
	}
}
