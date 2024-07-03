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
	"sync/atomic"

	"golang.org/x/term"
)

// backendConsole is the Backend that writes logs to console
type backendConsole struct {
	color int32      // No: -1, Yes: +1, Unknown: 0
	mutex sync.Mutex // Send lock
}

// Line implements [Backend.Send] method
func (bk *backendConsole) Send(levels []Level, lines [][]byte) {
	// Color auto-detection
	if atomic.LoadInt32(&bk.color) == 0 {
		isatty := term.IsTerminal(int(os.Stdout.Fd()))
		if isatty {
			atomic.CompareAndSwapInt32(&bk.color, 0, +1)
		}
	}

	// Build the entire message in the buffer
	buf := bufAlloc()
	defer bufFree(buf)

	for i := range levels {
		level := levels[i]
		line := lines[i]

		var beg, end string

		if atomic.LoadInt32(&bk.color) > 0 {
			switch level {
			case LevelTrace:
				beg, end = "\033[37m", "\033[0m" // Gray
			case LevelDebug:
				beg, end = "\033[37;1m", "\033[0m" // White
			case LevelInfo:
				beg, end = "\033[32;1m", "\033[0m" // Green
			case LevelError, LevelFatal:
				beg, end = "\033[31;1m", "\033[0m" // Red
			}
		}

		buf.Write([]byte(beg))
		buf.Write(line)
		buf.Write([]byte(end + "\n"))
	}

	// Now send buffer to the os.Stdout
	bk.mutex.Lock()
	buf.WriteTo(os.Stdout)
	bk.mutex.Unlock()
}
