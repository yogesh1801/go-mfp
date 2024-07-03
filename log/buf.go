// MFP - Miulti-Function Printers and scanners toolkit
// Logging facilities
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Buffer pool

package log

import (
	"bytes"
	"sync"
)

// bufPoolLargeBufferSize defines the capacity threshold for buffer
// to be reused.
//
// Buffers with capacity that exceeds this limits will not be returned
// into the bufPool, but instead left for GC.
const bufPoolLargeBufferSize = 1024

// bufPool is the pool of temporary buffers.
var bufPool = sync.Pool{
	New: func() any { return new(bytes.Buffer) },
}

// bufAlloc allocates a temporary buffer
func bufAlloc() *bytes.Buffer {
	return bufPool.Get().(*bytes.Buffer)
}

// bufFree releases a buffer, previously allocated by bufAlloc()
func bufFree(buf *bytes.Buffer) {
	if buf.Cap() <= bufPoolLargeBufferSize {
		buf.Reset()
		bufPool.Put(buf)
	}
}
