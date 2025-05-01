// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scanned or printable document

package abstract

import (
	"io"
	"sync"
)

// Document contains one or more image files, each of which may include
// (depending on the format) one or more image pages.
//
// For example, ADF scanner may return a Document that contains multiple
// files, one per page, and each file will contain a single page.
// Or, depending on a scanner model and format being used, Document may
// contain a single file with multiple pages, one per scanned physical page.
//
// All files in the document use the same format. The MIME type of
// the format can be obtained using the Document.Format method.
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

	// Next returns the next document file, represented as
	// a byte stream.
	//
	// Next implicitly closes the reader, returned by the
	// previous call to the Next function.
	//
	// If there are no more pages, Next returns [io.EOF].
	Next() (io.Reader, error)

	// Close closes the Document. It implicitly closes the current
	// image being read.
	Close() error
}

// documentFromBytes is the [Document], constructed by the
// [NewDocumentFromBytes] from a sequence of byte slices.
type documentFromBytes struct {
	format string               // Returned by Document.Format
	res    Resolution           // Returned by Document.Resolution
	files  [][]byte             // Remaining "files"
	reader *documentBytesReader // Current reader
	closed bool                 // True if document is closed
	lock   sync.Mutex           // Access lock
}

// documentBytesReader implements io.Reader for reading from
// the documentFromBytes
type documentBytesReader struct {
	data []byte     // Remaining data bytes
	lock sync.Mutex // Access lock
}

// Read reads data bytes from the [documentBytesReader].
func (rd *documentBytesReader) Read(buf []byte) (n int, err error) {
	rd.lock.Lock()

	switch {
	case rd.data == nil:
		err = ErrDocumentClosed

	case len(rd.data) == 0:
		err = io.EOF

	default:
		n = copy(buf, rd.data)
		rd.data = rd.data[n:]
	}

	rd.lock.Unlock()

	return
}

// close closes the documentBytesReader
func (rd *documentBytesReader) close() {
	rd.lock.Lock()
	rd.data = nil
	rd.lock.Unlock()
}

// Format returns Document's MIME type
func (doc *documentFromBytes) Format() string {
	return doc.format
}

// Format returns Document's Resolution
func (doc *documentFromBytes) Resolution() Resolution {
	return doc.res
}

// Next returns the next file as [io.Reader]
func (doc *documentFromBytes) Next() (io.Reader, error) {
	// Lock the lock
	doc.lock.Lock()
	defer doc.lock.Unlock()

	// Already closed?
	if doc.closed {
		return nil, ErrDocumentClosed
	}

	// Close the previously returned reader
	if doc.reader != nil {
		doc.reader.close()
		doc.reader = nil
	}

	// Return new reader, if more data is available
	if len(doc.files) != 0 {
		doc.reader = &documentBytesReader{data: doc.files[0]}
		doc.files = doc.files[1:]
		return doc.reader, nil
	}

	return nil, io.EOF
}

// Close closes the document.
func (doc *documentFromBytes) Close() error {
	// Lock the lock
	doc.lock.Lock()
	defer doc.lock.Unlock()

	// Close current reader
	if doc.reader != nil {
		doc.reader.close()
		doc.reader = nil
	}

	// Purge the document
	doc.files = nil
	doc.closed = true

	return nil
}

// NewDocumentFromBytes creates a new [Document], composed from
// the supplied byte slices. Each slice corresponds to the single
// file, returned by the Document.Next
func NewDocumentFromBytes(format string, res Resolution,
	files ...[]byte) Document {
	return &documentFromBytes{
		format: format,
		res:    res,
		files:  files,
	}
}
