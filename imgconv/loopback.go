// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Decoder/Encoder loopback

package imgconv

import (
	"image/color"
	"io"
	"sync/atomic"
)

// loopbackQueueSize is the queue size of loopback, in image rows
const loopbackQueueSize = 8

// loopback implements the Decoder/Encoder loopback
type loopback struct {
	wid, hei int         // Image size
	model    color.Model // Image color model
	queue    chan Row    // Row queue
}

// loopbackDecoder implements the Decoder side of the loopback
type loopbackDecoder struct {
	*loopback     // Underlying loopback
	readcnt   int // Count of read rows
}

// loopbackEncoder implements the Encoder side of the loopback
type loopbackEncoder struct {
	*loopback             // Underlying loopback
	writecnt  int         // Count of written rows
	closed    atomic.Bool // Encoder is closed
}

// NewLoopback creates a pair of [Decoder] and [Encoder], connected
// via internal loopback.
//
// Every image [Row], written into the [Encoder] appears as a [Row]
// that can be read from the [Decoder].
//
// Encoder allows up to hei image Rows to be written, all subsequent
// rows are silently discarded.
//
// Decoder allows up to hei image Rows to be read. All subsequent Read
// attempt will return io.EOF.
//
// If Encoder is closed before hei Rows were written, Decoder will
// return all actually written Rows, then io.ErrUnexpectedEOF.
//
// If Decoder is closed before all Rows are consumed, Encoder.Write
// will not be blocked. However, if Decoder is not closed and nobody
// reads from it, Encoder.Write will block when queue becomes full.
//
// This is safe to call [Encoder.Write] and [Decoder.Read] simultaneously
// from the different goroutines.
//
// [Decoder.ColorModel], [Decoder.Size], [Encoder.ColorMode] and
// [Encoder.Colormode] are completely goroutine-safe.
//
// All other methods need explicit synchronization and not reentrant.
func NewLoopback(wid, hei int, model color.Model) (Decoder, Encoder) {
	l := &loopback{
		wid:   wid,
		hei:   hei,
		model: model,
		queue: make(chan Row, loopbackQueueSize),
	}
	return &loopbackDecoder{loopback: l}, &loopbackEncoder{loopback: l}
}

// ColorModel returns the [color.Model] of image being decoded.
func (l *loopback) ColorModel() color.Model {
	return l.model
}

// Size returns the image size.
func (l *loopback) Size() (wid, hei int) {
	return l.wid, l.hei
}

// NewRow allocates a [Row] of the appropriate type and width for
// use with the [Decoder.Read] function.
func (decoder *loopbackDecoder) NewRow() Row {
	return NewRow(decoder.model, decoder.wid)
}

// Read returns the next image [Row].
func (decoder *loopbackDecoder) Read(row Row) (int, error) {
	if decoder.readcnt == decoder.hei {
		return 0, io.EOF
	}

	next := <-decoder.queue

	if next == nil {
		return 0, io.ErrUnexpectedEOF
	}

	decoder.readcnt++
	row.Copy(next)

	return next.Width(), nil
}

// Write writes the next image [Row].
func (encoder *loopbackEncoder) Write(row Row) error {
	if encoder.writecnt != encoder.hei && !encoder.closed.Load() {
		encoder.queue <- row
		encoder.writecnt++
	}
	return nil
}

// Close closes the Decoder side of loopback
func (decoder loopbackDecoder) Close() {
	// Drain the queue, unblock Encoder
	go func() {
		for range decoder.queue {
		}
	}()
}

// Close closes the Encoder side of loopback
func (encoder *loopbackEncoder) Close() error {
	if encoder.closed.CompareAndSwap(false, true) {
		close(encoder.queue)
	}
	return nil
}
