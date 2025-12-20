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
	io.Reader           // Underlying io.TeeReader
	rCloser   io.Closer // Closer part of original io.ReadCloser
	wCloser   io.Closer // Closer part of original io.Writer
	closefunc func()    // Called by tee.Close() under sync.OnceFunc
	closeerr  error     // Error that tee.Close() returns
}

// TeeReadCloser is like [io.TeeReader] but for [io.ReadCloser]s.
//
// If the supplied [io.Writer] supports Close method, it will
// be closed as well.
func TeeReadCloser(r io.ReadCloser, w io.Writer) io.ReadCloser {
	tee := &teeReadCloser{
		Reader:  io.TeeReader(r, w),
		rCloser: r,
	}

	tee.closefunc = sync.OnceFunc(tee.doClose)

	if wCloser, ok := w.(io.Closer); ok {
		tee.wCloser = wCloser
	}

	return tee
}

// Close closes the teeReadCloser.
func (tee *teeReadCloser) Close() error {
	tee.closefunc()
	return tee.closeerr
}

func (tee *teeReadCloser) doClose() {
	tee.closeerr = tee.rCloser.Close()
	if tee.wCloser != nil {
		tee.wCloser.Close()
	}
}
