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

// Document contains one or more [DocumentFile]s, each of which may include
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

	// Next returns the next [DocumentFile].
	//
	// Next implicitly closes the file, returned by the
	// previous call to the Next function.
	//
	// If there are no more files, Next returns [io.EOF].
	Next() (DocumentFile, error)

	// Close closes the Document. It implicitly closes the current
	// image being read.
	Close() error
}

// DocumentFile represents a single file, contained in the document.
//
// Essentially, it is the [io.Reader] that returns bytes from the
// file. In the most cases, the DocumentFile will not load all
// image bytes into the memory, but instead will read it on demand
// from the external source (say, network connection or a disk file).
type DocumentFile interface {
	// Format returns the MIME type of the image format used by
	// the document file.
	Format() string

	// Read reads the document file content as a sequence of bytes.
	// It implements the [io.Reader] interface.
	Read([]byte) (int, error)
}

// documentFromBytes is the [Document], constructed by the
// [NewDocumentFromBytes] from a sequence of byte slices.
type documentFromBytes struct {
	res    Resolution             // Returned by Document.Resolution
	files  [][]byte               // Remaining "files"
	file   *documentFromBytesFile // Current file
	closed bool                   // True if document is closed
	lock   sync.Mutex             // Access lock
}

// documentFromBytesFile implements the [DocumentFile] for reading from
// the documentFromBytes
type documentFromBytesFile struct {
	format string     // Returned by DocumentFile.Format
	data   []byte     // Remaining data bytes
	lock   sync.Mutex // Access lock
}

// newDocumentFromBytesFile returns new documentFromBytesFile
func newDocumentFromBytesFile(data []byte) *documentFromBytesFile {
	format := DocumentFormatDetect(data)
	if format == "" {
		format = DocumentFormatData
	}

	return &documentFromBytesFile{
		format: format,
		data:   data,
	}
}

// Format returns the MIME type of the image format used by
// the document file.
func (file *documentFromBytesFile) Format() string {
	return file.format
}

// Read reads data bytes from the [documentFromBytesFile].
func (file *documentFromBytesFile) Read(buf []byte) (n int, err error) {
	file.lock.Lock()

	switch {
	case file.data == nil:
		err = ErrDocumentClosed

	case len(file.data) == 0:
		err = io.EOF

	default:
		n = copy(buf, file.data)
		file.data = file.data[n:]
	}

	file.lock.Unlock()

	return
}

// close closes the documentFromBytesFile
func (file *documentFromBytesFile) close() {
	file.lock.Lock()
	file.data = nil
	file.lock.Unlock()
}

// Format returns Document's Resolution
func (doc *documentFromBytes) Resolution() Resolution {
	return doc.res
}

// Next returns the next file as [io.Reader]
func (doc *documentFromBytes) Next() (DocumentFile, error) {
	// Lock the lock
	doc.lock.Lock()
	defer doc.lock.Unlock()

	// Already closed?
	if doc.closed {
		return nil, ErrDocumentClosed
	}

	// Close the previously returned file
	if doc.file != nil {
		doc.file.close()
		doc.file = nil
	}

	// Return new file, if more data is available
	if len(doc.files) != 0 {
		doc.file = newDocumentFromBytesFile(doc.files[0])
		doc.files = doc.files[1:]
		return doc.file, nil
	}

	return nil, io.EOF
}

// Close closes the document.
func (doc *documentFromBytes) Close() error {
	// Lock the lock
	doc.lock.Lock()
	defer doc.lock.Unlock()

	// Close current file
	if doc.file != nil {
		doc.file.close()
		doc.file = nil
	}

	// Purge the document
	doc.files = nil
	doc.closed = true

	return nil
}

// NewDocumentFromBytes creates a new [Document], composed from
// the supplied byte slices. Each slice corresponds to the single
// file, returned by the Document.Next
func NewDocumentFromBytes(res Resolution, files ...[]byte) Document {
	return &documentFromBytes{
		res:   res,
		files: files,
	}
}
