// MFP       - Miulti-Function Printers and scanners toolkit
// TRANSPORT - Transport protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// TeeReadCloser - io.TeeReader with Close

package transport

import (
	"io"
	"sync"
)

// teeReadCloser is like [io.TeeReader], but implements [io.ReadCloser]
// interface as well.
//
// It can be useful, for example, to wrap the [http.Request.Body] or
// [http.Response.Body] in order to shiff their content.
type teeReadCloser struct {
	r     io.ReadCloser // Source reader
	w     io.Writer     // Destination writer
	tr    io.Reader     // Underlying io.TeeReader(r, w)
	onceR sync.Once     // To close source only once
	onceW sync.Once     // To close destination only once
	rErr  error         // Stores the error from closing the reader
}

// TeeReadCloser is like [io.TeeReader] but for [io.ReadCloser]s.
//
// If the supplied [io.Writer] implements the [io.Closer] interface, it
// will be closed automatically when the returned [io.ReadCloser] is
// closed, or as soon as the underlying reader is exhausted (returns
// io.EOF) or returns an error.
func TeeReadCloser(r io.ReadCloser, w io.Writer) io.ReadCloser {
	return &teeReadCloser{
		r:  r,
		w:  w,
		tr: io.TeeReader(r, w),
	}
}

// TeeReadCloser2 is an alternative to [TeeReadCloser] that splits a single
// [io.ReadCloser] into two separate [io.ReadCloser] streams.
//
// Data read from r1 is internally copied to r2 using an [io.Pipe].
// Note that r1 and r2 are synchronized: a Read on r1 will block until
// the data is also accepted by a concurrent Read on r2.
//
// Both readers must be closed by the caller to ensure all resources
// (including the pipe) are released.
func TeeReadCloser2(r io.ReadCloser) (r1, r2 io.ReadCloser) {
	rpipe, wpipe := io.Pipe()
	r1 = TeeReadCloser(r, wpipe)
	return r1, rpipe
}

// closeReader closes t.r and saves its Close error
func (t *teeReadCloser) closeReader() {
	t.rErr = t.r.Close()
}

// closeWriter closes t.w, if it implements io.Closer interface.
func (t *teeReadCloser) closeWriter() {
	if closer, ok := t.w.(io.Closer); ok {
		closer.Close()
	}
}

// Read reads bytes io.ReadCloser, supplied to the [TeeReadCloser] or
// [TeeReadCloser2], and forwards them to the second parameter of
// these functions, either directly, as [io.TeeReader] does, or
// via the [io.Pipe].
func (t *teeReadCloser) Read(p []byte) (n int, err error) {
	n, err = t.tr.Read(p)
	if err != nil {
		// Close writer on error or EOF
		t.onceW.Do(t.closeWriter)
	}
	return n, err
}

// Close closes the reader.
func (t *teeReadCloser) Close() error {
	// Close reader exactly once
	t.onceR.Do(t.closeReader)

	// Close writer exactly once (might have been closed by Read already)
	t.onceW.Do(t.closeWriter)
	return t.rErr
}
