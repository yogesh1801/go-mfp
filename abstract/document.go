// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scanned or printable document

package abstract

import "io"

// Document contains one or more image files, each of which may include
// (depending on the format) one or more image pages.
//
// All files in the document use the same format. The MIME type of
// the format can be obtained using the Document.Format method.
//
// For example, a document might contain three JPEG image files.
// Alternatively, it could include five multipage PDF files, each
// containing several image pages.
//
// This document can be used, for instance, to represent a sequence of
// images returned by a [Scanner]. Depending on the scanner's
// capabilities and the selected image format, this could result in
// either one image per file or a series of multipage files.
//
// The document interface is optimized for streaming images, eliminating
// the need to maintain a full-page image buffer in memory.
type Document interface {
	// Format returns the MIME type of the image format used by
	// the document.
	Format() string

	// Resolution returns the document's rendering resolution in DPI
	// (dots per inch).
	//
	// Knowing the resolution is essential for correctly rendering
	// document pages, especially if the file format does not include
	// geometric size information (e.g., JPEG).
	//
	// For formats that do include size details (e.g., PDF), the
	// embedded information will most likely be used instead.
	Resolution() Resolution

	// Next returns a next image page, represented as
	// a byte stream.
	//
	// Next implicitly closes the page, returned by the
	// previous call to the Next function.
	//
	// If there are no more pages, Next returns [io.EOF].
	Next() (io.Reader, error)

	// Close closes the Document. It implicitly closes the current
	// image being read.
	Close() error
}
