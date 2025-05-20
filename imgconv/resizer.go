// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Image resizer

package imgconv

import (
	"image"
	"image/color"
	"io"

	"github.com/OpenPrinting/go-mfp/util/generic"
)

// resizer implements an image resizer.
type resizer struct {
	input     Decoder         // Image source
	err       error           // I/O error
	rect      image.Rectangle // Clipping region
	fill      color.Color     // Filler color
	nooverlap bool            // Source and destination don't overlap
	y         int             // y-coordinate, relative to destination
	skip      int             // Lines to skip
	tmp       Row             // For horizontal shifting, nil if not needed
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
	bounds := image.Rect(0, 0, wid, hei)

	// Bypass filter, if target rectangle is equal to source bounds
	if rect.Eq(bounds) {
		return in
	}

	// Create resize filter
	model := in.ColorModel()
	rsz := &resizer{
		input:     in,
		rect:      rect.Canon(),
		fill:      model.Convert(color.White),
		nooverlap: !rect.Overlaps(bounds),
		skip:      generic.Max(0, rect.Min.Y),
	}

	// Allocate temporary row for the horizontal shifting,
	// if we really need it.
	if !rsz.nooverlap && rect.Min.X > 0 {
		sz := generic.Min(wid, rect.Max.X)
		rsz.tmp = NewRow(model, sz)
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
	wid, hei := rsz.Size()
	wid = generic.Min(wid, row.Width())
	row = row.Slice(0, wid)

	srcY := rsz.y + rsz.rect.Min.Y
	srcWid, srcHei := rsz.input.Size()

	// Handle special cases
	switch {
	// We've already got an I/O error, just return it.
	case rsz.err != nil:
		return 0, rsz.err

	// Image consumed till the end. Return io.EOF
	case rsz.y == hei:
		return 0, io.EOF

	// Target row doesn't overlap with source. Just fill it.
	case srcY < 0 || srcY >= srcHei || rsz.nooverlap:
		row.Fill(rsz.fill)
		rsz.y++
		return wid, nil

	// Target shifted down vertically. Consume unneeded source lines.
	case rsz.skip > 0:
		tmp := NewRow(rsz.ColorModel(), 0)
		for rsz.skip > 0 {
			_, err := rsz.input.Read(tmp)
			if err != nil {
				rsz.err = err
				return 0, err
			}

			rsz.skip--
		}
	}

	// Read the next row.
	// We do it either directly, or via intermediate buffer.
	//
	// Note, the need of double buffering already checked by
	// constructor and the fact that source and destination rows
	// do overlap already checked by the preceding lines of code.
	if rsz.tmp == nil {
		l := -rsz.rect.Min.X
		r := generic.Min(srcWid, rsz.rect.Max.X) - rsz.rect.Min.X

		_, err := rsz.input.Read(row.Slice(l, r))
		if err != nil {
			rsz.err = err
			return 0, err
		}

		row.Slice(0, l).Fill(rsz.fill)
		row.Slice(r, wid).Fill(rsz.fill)
	} else {
		_, err := rsz.input.Read(rsz.tmp)
		if err != nil {
			rsz.err = err
			return 0, err
		}

		n := row.Copy(rsz.tmp.Slice(rsz.rect.Min.X, rsz.tmp.Width()))
		row.Slice(n, wid).Fill(rsz.fill)
	}

	rsz.y++

	return wid, nil
}

// Close closes the decoder
func (rsz *resizer) Close() {
	rsz.input.Close()
}
