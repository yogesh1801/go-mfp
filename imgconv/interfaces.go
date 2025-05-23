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
)

// Decoder implements streaming image decoder. It allows to read
// image line by line and doesn't provide a whole image internal
// buffer, so even the large images can be processed using very
// small memory amount.
//
// There are two kinds of decoders:
//   - image format decoder (for example, PNG decoder). They work
//     on a top of [io.Reader].
//   - image filter (for example, image resizer). They work on a top
//     of existent Decoder and  also implement the Decoder interface,
//     which allows filters to be chained.
//
// Decoder may own some resources associated with it, which will
// not be automatically garbage-collected. So Decoder needs to
// be explicitly closed after use (with the Decoder.Close call).
//
// By convention, when Decoder is closed:
//   - image decoder will close its source io.Reader, if it
//     implements the [io.Closer] interface.
//   - image filter will close its underlying Decoder.
//
// Decoder doesn't guarantee that its input stream ([io.Reader] or
// underlying Decoder) will be consumed till the end or even that
// at least some data will be consumed from it.
type Decoder interface {
	// ColorModel returns the [color.Model] of image being decoded.
	ColorModel() color.Model

	// Size returns the image size.
	Size() (wid, hei int)

	// NewRow allocates a [Row] of the appropriate type and width for
	// use with the [Decoder.Read] function.
	NewRow() Row

	// Read returns the next image [Row].
	// The Row type must match the [Decoder]'s [color.Model].
	//
	// It returns the resulting row length, in pixels, or an error.
	Read(Row) (int, error)

	// Close closes the decoder.
	Close()
}

// Encoder implements streaming image encoder.
//
// It writes image row by row into the supplied [io.Writer].
//
// Encoder may own some resources associated with it, which will
// not be automatically garbage-collected. So Encoder needs to
// be explicitly closed after use (with the Encoder.Close call).
//
// By convention, Encoder.Close will close the output io.Writer,
// if it implements the [io.Closer] interface.
type Encoder interface {
	// Write writes the next image [Row].
	Write(Row) error

	// Close flushes the buffered data and then closes the Encoder
	Close() error
}
