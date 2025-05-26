// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Decoder/Encoder image adapter

package imgconv

import (
	"image"
	"image/color"
	"io"
	"sync"

	"github.com/OpenPrinting/go-mfp/util/generic"
)

// AdapterHistorySize specifies how many latest image rows
// [SourceImageAdapter] keeps on its history.
const AdapterHistorySize = 8

// SourceImageAdapter reads image rows sequentially from the provided
// [Decoder], retaining the last [AdapterHistorySize] rows in its internal
// history buffer. It implements the [image.Image] interface on top of this
// buffer.
//
// Although its [image.Image.At] method can only return meaningful values for
// the most recent rows, this is sufficient for many standard image processing
// operations (such as scaling) when using SourceImageAdapter as a source
// image.
type SourceImageAdapter struct {
	decoder Decoder                     // Underlying decoder
	rows    [AdapterHistorySize + 1]Row // Last some rows
	y       int                         // Latest Row's y-coordinate
	err     error                       // Sticky error
}

// NewSourceImageAdapter creates a new SourceImageAdapter on a top
// of existent [Decoder].
func NewSourceImageAdapter(decoder Decoder) *SourceImageAdapter {
	source := &SourceImageAdapter{decoder: decoder, y: -1}
	for i := 0; i <= AdapterHistorySize; i++ {
		source.rows[i] = decoder.NewRow()
	}
	return source
}

// Close closes the source and underlying [Decoder].
func (source *SourceImageAdapter) Close() {
	source.decoder.Close()
}

// Error returns the Decoder's error, if any.
func (source *SourceImageAdapter) Error() error {
	return source.err
}

// ColorModel returns the Image's color model.
func (source *SourceImageAdapter) ColorModel() color.Model {
	return source.decoder.ColorModel()
}

// Bounds returns image bounds (always 0-based).
func (source *SourceImageAdapter) Bounds() image.Rectangle {
	wid, hei := source.decoder.Size()
	return image.Rect(0, 0, wid, hei)
}

// At returns the color of the pixel at (x, y).
func (source *SourceImageAdapter) At(x, y int) color.Color {
	// The fast path: hope pixel already in the buffer
	off := source.y - y
	if off >= 0 && off < AdapterHistorySize {
		return source.rows[off].At(x)
	}

	// Read more rows.
	source.seek(y)
	if source.y == y {
		return source.rows[0].At(x)
	}

	// Fail. Return the default color.
	return source.decoder.ColorModel().Convert(color.Transparent)
}

// Read rows from the underlying decoder until y is reached or error
func (source *SourceImageAdapter) seek(y int) {
	for source.err == nil && source.y < y {
		_, err := source.decoder.Read(source.rows[AdapterHistorySize])
		if err != nil {
			source.err = err
		} else {
			row := source.rows[AdapterHistorySize]
			copy(source.rows[1:], source.rows[0:AdapterHistorySize])
			source.rows[0] = row
			source.y++
		}
	}
}

// TargetImageAdapter generates image sequentially as a series of Rows.
// It implements the [draw.Image] interface for image construction and
// [Decoder] interface for image Rows consumption.
//
// While its [draw.Image.Set] method only allows changes to pixels in the most
// recently accessed row (previous rows are automatically flushed when the
// y-coordinate advances), this restriction still supports many standard
// image operations - such as scaling - when using TargetImageAdapter as a
// destination image.
//
// TargetImageAdapter needs to be explicitly flushed. Otherwise, the
// latest image rows can be lost.
type TargetImageAdapter struct {
	model  color.Model     // Image color model
	bounds image.Rectangle // Image bounds
	y      int             // Latest Row's y-coordinate
	row    Row             // The latest image row
	queue  chan Row        // Decoder's read queue
	pool   sync.Pool       // Pool of empty rows
}

// NewTargetImageAdapter creates a new [TargetImageAdapter]
func NewTargetImageAdapter(wid, hei int, mdl color.Model) *TargetImageAdapter {
	target := &TargetImageAdapter{
		model:  mdl,
		bounds: image.Rect(0, 0, wid, hei),
		queue:  make(chan Row, AdapterHistorySize),
		pool:   sync.Pool{New: func() any { return NewRow(mdl, wid) }},
	}

	target.row = target.NewRow()

	return target
}

// ColorModel returns the [color.Model] of image being produces.
func (target *TargetImageAdapter) ColorModel() color.Model {
	return target.model
}

// Size returns the image size.
// It implements the [Decoder] interface.
func (target *TargetImageAdapter) Size() (wid, hei int) {
	return target.bounds.Max.X, target.bounds.Max.Y
}

// NewRow allocates a [Row] of the appropriate type and width for
// use with the [Decoder.Read] function.
//
// It implements the [Decoder] interface.
func (target *TargetImageAdapter) NewRow() Row {
	return NewRow(target.model, target.bounds.Max.X)
}

// Read returns the next image [Row].
// It implements the [Decoder] interface.
func (target *TargetImageAdapter) Read(row Row) (int, error) {
	next := <-target.queue
	if next == nil {
		return 0, io.EOF
	}

	row.Copy(next)
	target.pool.Put(next)

	return row.Width(), nil
}

// Close closes the decoder side of the TargetImageAdapter.
// It implements the [Decoder] interface.
func (target *TargetImageAdapter) Close() {
	// Drain the queue, unblock Encoder
	go func() {
		for range target.queue {
		}
	}()

}

// Bounds returns image bounds (always 0-based).
// It implements the [draw.Image] interface.
func (target *TargetImageAdapter) Bounds() image.Rectangle {
	return target.bounds
}

// At returns the color of the pixel at (x, y).
func (target *TargetImageAdapter) At(x, y int) color.Color {
	return target.model.Convert(color.Transparent)
}

// Set sets the color of the pixel at (x, y).
func (target *TargetImageAdapter) Set(x, y int, c color.Color) {
	// Ignore Set outside of the image bounds
	if (image.Point{x, y}).In(target.bounds) {
		// The fast path: just update the current row
		if y == target.y {
			target.row.Set(x, c)
			return
		}

		// Advance the current y. Fill possible gaps.
		target.advance(y)

		// Update the current row.
		target.row.Set(x, c)
	}
}

// Flush flushes still buffered image parts to the [Decoder] side
// of the [TargetImageAdapter].
func (target *TargetImageAdapter) Flush() {
	target.advance(target.y + 1)
	close(target.queue)
}

// advance advances target's current y-position, until requested
// row is reached.
func (target *TargetImageAdapter) advance(y int) {
	lim := generic.Min(y, target.bounds.Max.Y)
	for target.y < lim {
		target.queue <- target.row
		target.row = target.pool.Get().(Row)
		target.y++
	}
}
