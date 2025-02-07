// MFP       - Miulti-Function Printers and scanners toolkit
// TRANSPORT - Transport protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Peeker allows to prefetch data from io.ReadCloser

package transport

import (
	"bytes"
	"io"
)

// Peeker wraps [io.ReadCloser] object and allows to peek some
// data, then rewind the stream to the beginning or replace
// already consumed bytes with some other bytes and continue
// reading.
//
// It can be used, for example, to prefetch IPP Message body
// from the [http.Request.Body] or [http.Response.Body] and
// then forward HTTP message body either unmodified (with [Peeker.Rewind])
// or rewrote (with [Peeker.Replace]).
//
// This is important to call Peeker.Rewind or Peeker.Replace when
// enough data is prefetched and more prefetching is not planned,
// as calling these functions stops recording of the returned data,
// so avoiding excessive memory usage.
type Peeker struct {
	in  io.ReadCloser // Underlying io.ReadCloser
	out io.Reader     // Output stream
	buf bytes.Buffer  // Keeps consumed bytes for rewind
}

// NewPeeker creates a new [Peeker] that wraps existing [io.ReadCloser].
func NewPeeker(in io.ReadCloser) *Peeker {
	p := &Peeker{
		in: in,
	}
	p.out = io.TeeReader(in, &p.buf)
	return p
}

// Read reads up to len(b) bytes into b.
//
// It returns the number of bytes read (0 <= n <= len(b))
// and any error encountered.
func (p *Peeker) Read(b []byte) (int, error) {
	return p.out.Read(b)
}

// Close closes the Peeker and its underlying io.ReadCloser.
func (p *Peeker) Close() error {
	return p.in.Close()
}

// Rewind rewinds the output stream to the beginning, making
// already consumed bytes available again.
func (p *Peeker) Rewind() {
	p.out = io.MultiReader(&p.buf, p.in)
}

// Replace works like [Peeker.Rewind], but consumed data will be
// replaced with the new content.
func (p *Peeker) Replace(data []byte) {
	p.buf.Reset()
	p.buf.Write(data)
	p.out = io.MultiReader(&p.buf, p.in)
}
