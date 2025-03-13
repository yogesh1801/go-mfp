// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scanned or printable document

package abstract

import "io"

// Document contains one or more image pages.
//
// The document interface is optimized for streaming images, eliminating
// the need to maintain a full-page image buffer in memory.
type Document interface {
	// Format returns the MIME type of the image format used by
	// the document.
	Format() string

	// Next returns a next image page, represented as
	// a byte stream.
	//
	// Next implicitly closes the page, returned by the
	// previous call to the Next function.
	//
	// If there are no more pages, Next returns [io.EOF].
	Next() (io.Reader, error)

	// Close closes the Document. It implicitly closes the current
	// image being read, unless EOF is reached.
	Close() error
}
