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
// image line by line and doesn't use the whole image internal
// buffer, so even the large images can be processed using very
// small amount of memory.
//
// There are two kinds of decoders:
//   - image format decoder (for example, PNG decoder). They work
//     on a top of [io.Reader].
//   - image filter (for example, image resizer). They work on a top
//     of existent Decoder and  also implement the Decoder interface,
//     which allows filters to be stacked.
//
// [Decoder.Read] may block or fail, if its source [io.Reader]
// blocks of fails, and this condition propagates over the entire
// decoder stack.
//
// Once Decoder.Read fails, all subsequent calls to Decoder.Read
// will return the same error.
//
// Attempt to Read after the last line returns [io.EOF]. However,
// if underlying io.Reader unexpectedly returns io.EOF (i.e.,
// if image is truncated), Decoder.Read will return some other
// error code (typically, [io.ErrUnexpectedEOF]).
//
// Decoder may own some resources associated with it, which will
// not be automatically garbage-collected. So Decoder needs to
// be explicitly closed after use (with the Decoder.Close call).
//
// Closing the input format decoder doesn't close the underlying
// [io.Reader]. However, closing the image filter closes its
// underlying input [Decoder]. So the entire filter chain can
// be closed by closing the top-level filter, but the bottom
// [io.Reader] still needs to be closed explicitly.
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
type Encoder interface {
	// ColorModel returns the [color.Model] of image being encoded.
	ColorModel() color.Model

	// Size returns the image size.
	Size() (wid, hei int)

	// Write writes the next image [Row].
	Write(Row) error

	// Close flushes the buffered data and then closes the Encoder
	Close() error
}
