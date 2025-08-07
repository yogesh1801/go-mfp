// MFP - Miulti-Function Printers and scanners toolkit
// Logging facilities
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Console backend

package log

import (
	"os"
	"sync"
)

// backendStderr is the Backend that writes logs to stderr
type backendStderr struct {
	mutex sync.Mutex // Send lock
}

// Line implements [Backend.Send] method
func (bk *backendStderr) Send(levels []Level, lines [][]byte) {
	// Build the entire message in the buffer
	buf := bufAlloc()
	defer bufFree(buf)

	for _, line := range lines {
		buf.Write(line)
		buf.WriteByte('\n')
	}

	// Now send buffer to the os.Stderr
	buf.WriteTo(os.Stderr)
}
