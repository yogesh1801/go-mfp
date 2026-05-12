// MFP       - Miulti-Function Printers and scanners toolkit
// TRANSPORT - Transport protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// io.Writer with counter

package transport

// DiscardCounter is the [io.Writer] that discards all date
// being written to it (like [io.Discard]) but additionally
// counts amount of bytes being written.
type DiscardCounter struct {
	Count int64
}

// Write implements [io.Writer] interface for [DiscardCounter].
func (d *DiscardCounter) Write(data []byte) (int, error) {
	n := len(data)
	d.Count += int64(n)
	return n, nil
}
