// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Package documentation

package abstract

import (
	"io"
	"sync"
)

// documentReader implements the [Document] interface on a top of
// [io.ReadCloser].
type documentReader struct {
	in       io.ReadCloser // Input stream
	res      Resolution    // Document resolution
	format   string        // Document format
	consumed bool          // All document files consumed
	closed   bool          // documentReader.in closed
	closeErr error         // Close error, if any
	lock     sync.Mutex    // Access lock
}

// documentReaderFile implements the [DocumentFile] interface for
// the documentReader.
type documentReaderFile struct {
	doc *documentReader // Underlying documentReader
}

// NewDocumentReader creates a new [Document] interface adapter
// on a top of existent [io.ReadCloser].
//
// It doesn't interpret the content of the io.ReadCloser by itself.
// The [Resolution] and format information must be provided, to
// implement the [Document.Resolution] and [DocumentFile.Format]
// interfaces.
//
// Reading Documents from the existent [io.ReadCloser] may be
// useful for applying the [Filter] to the document content.
func NewDocumentReader(in io.ReadCloser,
	res Resolution, format string) Document {
	doc := &documentReader{
		in:     in,
		res:    res,
		format: format,
	}
	return doc
}

// Resolution returns the document's rendering resolution in DPI
// (dots per inch).
func (doc *documentReader) Resolution() Resolution {
	return doc.res
}

// Next returns the next [DocumentFile].
//
// If there are no more files, Next returns [io.EOF].
func (doc *documentReader) Next() (DocumentFile, error) {
	doc.lock.Lock()
	defer doc.lock.Unlock()

	if doc.consumed {
		doc.doClose()
		return nil, io.EOF
	}

	doc.consumed = true
	return &documentReaderFile{doc}, nil
}

// Close closes the Document. It implicitly closes the current
// [DocumentFile] being read.
func (doc *documentReader) Close() error {
	doc.lock.Lock()
	doc.doClose()
	doc.lock.Unlock()

	return doc.closeErr
}

// doClose actually closes the documentReader.
// It must be called under the documentReader.lock
func (doc *documentReader) doClose() {
	if !doc.closed {
		doc.closeErr = doc.in.Close()
		doc.closed = true
	}
}

// Format returns the MIME type of the image format used by
// the document file.
func (file *documentReaderFile) Format() string {
	return file.doc.format
}

// Read reads the document file content as a sequence of bytes.
// It implements the [io.Reader] interface.
func (file *documentReaderFile) Read(buf []byte) (int, error) {
	return file.doc.in.Read(buf)
}
