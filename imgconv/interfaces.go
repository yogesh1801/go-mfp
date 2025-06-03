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

// Reader implements streaming image reader. It allows to read
// image line by line and doesn't use the whole image internal
// buffer, so even the large images can be processed using very
// small amount of memory.
//
// There are two kinds of readers:
//   - image format reader (for example, PNG reader). They work
//     on a top of [io.Reader].
//   - image filter (for example, image resizer). They work on a top
//     of existent Reader and  also implement the Reader interface,
//     which allows filters to be stacked.
//
// [Reader.Read] may block or fail, if its source [io.Reader]
// blocks of fails, and this condition propagates over the entire
// reader stack.
//
// Once Reader.Read fails, all subsequent calls to Reader.Read
// will return the same error.
//
// Attempt to Read after the last line returns [io.EOF]. However,
// if underlying io.Reader unexpectedly returns io.EOF (i.e.,
// if image is truncated), Reader.Read will return some other
// error code (typically, [io.ErrUnexpectedEOF]).
//
// Reader may own some resources associated with it, which will
// not be automatically garbage-collected. So Reader needs to
// be explicitly closed after use (with the Reader.Close call).
//
// Closing the input format reader doesn't close the underlying
// [io.Reader]. However, closing the image filter closes its
// underlying input [Reader]. So the entire filter chain can
// be closed by closing the top-level filter, but the bottom
// [io.Reader] still needs to be closed explicitly.
//
// Reader doesn't guarantee that its input stream ([io.Reader] or
// underlying Reader) will be consumed till the end or even that
// at least some data will be consumed from it.
type Reader interface {
	// ColorModel returns the [color.Model] of image being decoded.
	ColorModel() color.Model

	// Size returns the image size.
	Size() (wid, hei int)

	// NewRow allocates a [Row] of the appropriate type and width for
	// use with the [Reader.Read] function.
	NewRow() Row

	// Read returns the next image [Row].
	// It returns the resulting row length, in pixels, or an error.
	Read(Row) (int, error)

	// Close closes the reader.
	Close()
}

// Writer implements streaming image writer.
//
// It writes image row by row into the supplied [io.Writer].
//
// Writer may own some resources associated with it, which will
// not be automatically garbage-collected. So Writer needs to
// be explicitly closed after use (with the Writer.Close call).
type Writer interface {
	// ColorModel returns the [color.Model] of image being written.
	ColorModel() color.Model

	// Size returns the image size.
	Size() (wid, hei int)

	// Write writes the next image [Row].
	Write(Row) error

	// Close flushes the buffered data and then closes the Writer
	Close() error
}
