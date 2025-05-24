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
)

// AdapterHistorySize specifies how many latest image rows
// [DecoderImageAdapter] or [EncoderImageAdapter] keep on
// their history.
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
