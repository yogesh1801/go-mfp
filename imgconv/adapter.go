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

	"github.com/OpenPrinting/go-mfp/util/generic"
)

// AdapterHistorySize specifies how many latest image rows
// [DecoderImageAdapter] keeps on its history.
const AdapterHistorySize = 8

// DecoderImageAdapter implements [image.Image] interface on a top of Decoder.
//
// In comparison to the normal image.Image it has the following limitations:
//   - image rows (y-coordinate when calling the [image.Image.At] method)
//     needs to be read in sequence.
//   - only the latest [AdapterHistorySize] lines are available.
//
// If these conditions are not met, [image.Image.At] simply returns
// fully transparent color converted to the underlying color model.
//
// Despite these limitations, DecoderImageAdapter can be used as a source
// image for many image transformation algorithms, like image scaling.
type DecoderImageAdapter struct {
	decoder Decoder                     // Underlying decoder
	rows    [AdapterHistorySize + 1]Row // Last some rows
	y       int                         // Latest Row's y-coordinate
	err     error                       // Sticky error
}

// NewDecoderImageAdapter creates a new DecoderImageAdapter on a top
// of existent [Decoder].
func NewDecoderImageAdapter(decoder Decoder) *DecoderImageAdapter {
	adapter := &DecoderImageAdapter{decoder: decoder}
	for i := 0; i < AdapterHistorySize; i++ {
		adapter.rows[i] = decoder.NewRow()
	}
	return adapter
}

// Close closes the adapter and underlying [Decoder].
func (adapter *DecoderImageAdapter) Close() {
	adapter.decoder.Close()
}

// Error returns the Decoder's error, if any.
func (adapter *DecoderImageAdapter) Error() error {
	return adapter.err
}

// ColorModel returns the Image's color model.
func (adapter *DecoderImageAdapter) ColorModel() color.Model {
	return adapter.decoder.ColorModel()
}

// Bounds returns image bounds (always 0-based).
func (adapter *DecoderImageAdapter) Bounds() image.Rectangle {
	wid, hei := adapter.decoder.Size()
	return image.Rect(0, 0, wid, hei)
}

// At returns the color of the pixel at (x, y).
func (adapter *DecoderImageAdapter) At(x, y int) color.Color {
	// The fast path: hope pixel already in the buffer
	off := adapter.y - y
	if off >= 0 && off < AdapterHistorySize {
		return adapter.rows[off].At(x)
	}

	// Read more rows.
	adapter.seek(y)
	if adapter.y == y {
		return adapter.rows[0].At(x)
	}

	// Fail. Return the default color.
	return adapter.decoder.ColorModel().Convert(color.Transparent)
}

// Read rows from the underlying decoder until y is reached or error
func (adapter *DecoderImageAdapter) seek(y int) {
	for adapter.err == nil && adapter.y < y {
		_, err := adapter.decoder.Read(adapter.rows[AdapterHistorySize])
		if err != nil {
			adapter.err = err
		} else {
			row := adapter.rows[AdapterHistorySize]
			copy(adapter.rows[1:], adapter.rows[0:AdapterHistorySize])
			adapter.rows[0] = row
			adapter.y++
		}
	}
}

// EncoderImageAdapter implements [draw.Image] interface on a top of Encoder.
//
// In comparison to the normal draw.Image it has the following limitations:
//   - the [draw.Image.At] method is meaningless and exists simply
//     for compatibility.
//   - image rows (y-coordinate when calling the [draw.Image.Set] method)
//     needs to be filled in sequence.
//   - only the latest image row is available for draw.Image.Set (i.e.,
//     once some y is Set, y-1 and so becomes unavailable).
//
// If these conditions are not met, [draw.Image.Set] silently discards
// the pixel.
//
// EncoderImageAdapter needs to be explicitly closed. Otherwise, the
// latest image rows can be lost.
//
// Despite these limitations, EncoderImageAdapter can be used as a destination
// image for many image transformation algorithms, like image scaling.
type EncoderImageAdapter struct {
	encoder Encoder         // Underlying encoder
	bounds  image.Rectangle // Image bounds
	y       int             // Latest Row's y-coordinate
	row     Row             // The latest image row
	err     error           // Sticky error
}

// NewEncoderImageAdapter creates a new EncoderImageAdapter on a top
// of existent [Encoder].
func NewEncoderImageAdapter(encoder Encoder) *EncoderImageAdapter {
	wid, hei := encoder.Size()
	model := encoder.ColorModel()

	adapter := &EncoderImageAdapter{
		encoder: encoder,
		bounds:  image.Rect(0, 0, wid, hei),
		row:     NewRow(model, wid),
	}

	return adapter
}

// Close closes the adapter and underlying [Encoder].
func (adapter *EncoderImageAdapter) Close() {
	adapter.advance(adapter.bounds.Max.Y)
	adapter.encoder.Close()
}

// Error returns the Encoder's error, if any.
func (adapter *EncoderImageAdapter) Error() error {
	return adapter.err
}

// ColorModel returns the Image's color model.
func (adapter *EncoderImageAdapter) ColorModel() color.Model {
	return adapter.encoder.ColorModel()
}

// Bounds returns image bounds (always 0-based).
func (adapter *EncoderImageAdapter) Bounds() image.Rectangle {
	wid, hei := adapter.encoder.Size()
	return image.Rect(0, 0, wid, hei)
}

// At returns the color of the pixel at (x, y).
func (adapter *EncoderImageAdapter) At(x, y int) color.Color {
	return adapter.encoder.ColorModel().Convert(color.Transparent)
}

// Set sets the color of the pixel at (x, y).
func (adapter *EncoderImageAdapter) Set(x, y int, c color.Color) {
	// Ignore Set outside of the image bounds
	if !(image.Point{x, y}.In(adapter.bounds)) {
		return
	}

	// The fast path: just update the current row
	if y == adapter.y {
		adapter.row.Set(x, c)
		return
	}

	// Advance the current y. Fill possible gaps.
	adapter.advance(y)

	// Update the current row, if everything is OK so far.
	if y == adapter.y {
		adapter.row.Set(x, c)
	}
}

// Flush sends out image parts not has been written yet into the
// underlying encoder.
func (adapter *EncoderImageAdapter) Flush() error {
	adapter.advance(adapter.y + 1)
	return adapter.err
}

// advance advances encoder's current y-position, until requested
// row is reached or error occurs.
func (adapter *EncoderImageAdapter) advance(y int) {
	row := adapter.row
	lim := generic.Min(y, adapter.bounds.Max.Y)
	for adapter.y < lim && adapter.err == nil {
		adapter.err = adapter.encoder.Write(row)
		if adapter.err == nil {
			adapter.y++
			row = RowEmpty{}
		}
	}
}
