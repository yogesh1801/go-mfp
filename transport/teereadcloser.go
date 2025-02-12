// MFP       - Miulti-Function Printers and scanners toolkit
// TRANSPORT - Transport protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Package documentation

package transport

import "io"

// teeReadCloser is like [io.TeeReader], but implements [io.ReadCloser]
// interface as well.
//
// It can be useful, for example, to wrap the [http.Request.Body] or
// [http.Response.Body] in order to shiff their content.
type teeReadCloser struct {
	io.Reader // Underlying io.TeeReader
	io.Closer // Closer part of original io.ReadCloser
}

// TeeReadCloser returns the new [TeeReadCloser].
func TeeReadCloser(r io.ReadCloser, w io.Writer) io.ReadCloser {
	return teeReadCloser{
		Reader: io.TeeReader(r, w),
		Closer: r,
	}
}
