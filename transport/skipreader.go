// MFP       - Miulti-Function Printers and scanners toolkit
// TRANSPORT - Transport protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// SkipReader - io.Reader that skips some data

package transport

import "io"

type skipReader struct {
	r io.Reader
	n int64
}

// SkipReader returns the [io.Reader] that skips the first n bytes
// from the underlying stream.
func SkipReader(r io.Reader, n int) io.Reader {
	return &skipReader{r: r, n: int64(n)}
}

// Read implements reading from the skipReader.
func (skip *skipReader) Read(buf []byte) (int, error) {
	if skip.n > 0 {
		n, err := io.CopyN(io.Discard, skip.r, skip.n)
		skip.n -= n
		if err != nil {
			return 0, err
		}
	}

	return skip.r.Read(buf)
}
