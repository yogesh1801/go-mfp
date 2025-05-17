// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Common interfaces

package imgconv

import (
	"image/color"

	"golang.org/x/image/draw"
)

// Decoder implements streaming image decoder.
//
// It implements [image.Image] interface with the following
// limitations:
//   - image rows must be read in sequence ([image.Image.At]
//     method must be called with non-increasing y-coordinates).
//
// Decoder is suitable as a source image for the [draw.Scaler.Scale]
// function.
type Decoder interface {
	// ColorModel returns the [color.Model] of image being decoded.
	ColorModel() color.Model

	// Size returns the image size.
	Size() (wid, hei int)

	// Next returns the next image [Row].
	Next() (Row, error)

	// Close closes the decoder
	Close()
}

// Encoder implements streaming image encoder.
//
// It implements [draw.Image] interface with the following
// limitations:
//   - image rows must be written in sequence ([image.Image.At
//     method must be called with non-increasing y-coordinates).
//   - [image.Image.At] is not functional
//
// Encoder is bound to [io.WriteCloser], specified during Encoder
// creation where it writes the generated image.
//
// Encoder is suitable as a destination image for the [draw.Scaler.Scale]
// function.
type Encoder interface {
	draw.Image

	// Error returns the latest I/O error, encountered during
	// the Encoder operations.
	Error() error

	// Flush writes out the buffered data.
	Flush() error

	// Close flushes the buffered data and then closes
	// the destination [io.WriteCloser].
	Close() error
}
