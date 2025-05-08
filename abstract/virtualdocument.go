// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// VirtualDocument implements Document interface on a top of in-memory images

package abstract

import (
	"io"
	"sync"
)

// virtualDocument is the [Document], constructed by the
// [NewVirtualDocument] from a sequence of byte slices.
type virtualDocument struct {
	res    Resolution           // Returned by Document.Resolution
	files  [][]byte             // Bodies of not yet consumed "files"
	file   *virtualDocumentFile // Current file
	closed bool                 // True if document is closed
	lock   sync.Mutex           // Access lock
}

// NewVirtualDocument creates a new [Document], composed from
// the supplied byte slices. Each slice corresponds to the single
// file, returned by the Document.Next
func NewVirtualDocument(res Resolution, files ...[]byte) Document {
	return &virtualDocument{
		res:   res,
		files: files,
	}
}

// virtualDocumentFile implements the [DocumentFile] for reading from
// the virtualDocument
type virtualDocumentFile struct {
	format string     // Returned by DocumentFile.Format
	data   []byte     // Remaining data bytes
	lock   sync.Mutex // Access lock
}

// newVirtualDocumentFile returns new virtualDocumentFile
func newVirtualDocumentFile(data []byte) *virtualDocumentFile {
	format := DocumentFormatDetect(data)
	if format == "" {
		format = DocumentFormatData
	}

	return &virtualDocumentFile{
		format: format,
		data:   data,
	}
}

// Format returns the MIME type of the image format used by
// the document file.
func (file *virtualDocumentFile) Format() string {
	return file.format
}

// Read reads data bytes from the [virtualDocumentFile].
func (file *virtualDocumentFile) Read(buf []byte) (n int, err error) {
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

// close closes the virtualDocumentFile
func (file *virtualDocumentFile) close() {
	file.lock.Lock()
	file.data = nil
	file.lock.Unlock()
}

// Format returns Document's Resolution
func (doc *virtualDocument) Resolution() Resolution {
	return doc.res
}

// Next returns the next file as [io.Reader]
func (doc *virtualDocument) Next() (DocumentFile, error) {
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
		doc.file = newVirtualDocumentFile(doc.files[0])
		doc.files = doc.files[1:]
		return doc.file, nil
	}

	return nil, io.EOF
}

// Close closes the document.
func (doc *virtualDocument) Close() error {
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
