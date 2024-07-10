// MFP - Miulti-Function Printers and scanners toolkit
// Logging facilities
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Package documentation

package log

import (
	"encoding"
	"sync"
)

// Standard loggers:
var (
	// DefaultLogger is the default logging destination.
	DefaultLogger = NewLogger("", LevelAll, Console)

	// DiscardLogger discards all logs written to.
	DiscardLogger = NewLogger("", LevelNone, Discard)
)

// Logger is the logging destination.
// It can be connected to console, to the disk file etc...
type Logger struct {
	prefix  string       // Log prefix
	out     []loggerDest // Attached destinations
	outLock sync.Mutex   // Destinations modification lock
}

// loggerDest represents logging destination
type loggerDest struct {
	level   Level
	backend Backend
}

// NewLogger returns a new logger, attached to the specified backend
func NewLogger(prefix string, lvl Level, b Backend) *Logger {
	if prefix != "" {
		prefix = prefix + ": "
	}

	return &Logger{
		prefix: prefix,
		out: []loggerDest{
			{
				level:   lvl,
				backend: b,
			},
		},
	}
}

// Attach adds an additional [Backend] to send logs to.
//
// If this backend already attached to this logger, it
// only updates the log level.
func (lgr *Logger) Attach(lvl Level, b Backend) {
	// Must do under the lock
	lgr.outLock.Lock()
	defer lgr.outLock.Unlock()

	// If Backend already attached just update a Level
	for i := range lgr.out {
		if lgr.out[i].backend == b {
			lgr.out[i].level = lvl
			return
		}
	}

	// Create new attachment
	lgr.out = append(lgr.out, loggerDest{level: lvl, backend: b})
}

// Begin initiates creation of a new multi-line log [Record].
// Records are always written atomically. Records written from
// the concurrently running goroutines are never intermixed at
// output. During log rotation, Records are not split between
// different log files.
func (lgr *Logger) Begin() *Record {
	return &Record{parent: lgr}
}

// Trace writes a Trace-level message to the Logger.
func (lgr *Logger) Trace(format string, v ...any) *Logger {
	return lgr.Begin().Trace(format, v...).Commit()
}

// Debug writes a Debug-level message to the Logger.
func (lgr *Logger) Debug(format string, v ...any) *Logger {
	return lgr.Begin().Debug(format, v...).Commit()
}

// Info writes a Info-level message to the Logger.
func (lgr *Logger) Info(format string, v ...any) *Logger {
	return lgr.Begin().Info(format, v...).Commit()
}

// Error writes a Error-level message to the Logger.
func (lgr *Logger) Error(format string, v ...any) *Logger {
	return lgr.Begin().Error(format, v...).Commit()
}

// Fatal writes a Fatal-level message to the Logger.
//
// It calls os.Exit(1) and never returns.
func (lgr *Logger) Fatal(format string, v ...any) {
	lgr.Begin().Fatal(format, v...)
}

// Object writes any object that implements [encoding.TextMarshaler]
// interface to the Logger.
//
// If [encoding.TextMarshaler.MarshalText] returns an error, it
// will be written to log with the [Error] log level, regardless
// of the level specified by the first parameter.
func (lgr *Logger) Object(level Level, obj encoding.TextMarshaler) *Logger {
	return lgr.Begin().Object(level, obj).Commit()
}

// send writes some lines to the Logger.
func (lgr *Logger) send(levels []Level, lines [][]byte) *Logger {
	// Prepend prefix
	if lgr.prefix != "" {
		prefixed := make([][]byte, len(lines))
		for i := range lines {
			prefixed[i] = []byte(lgr.prefix + string(lines[i]))
		}
		lines = prefixed
	}

	// Send message to all destinations
	for _, dest := range lgr.out {
		// Filter lines by level
		filteredLevels := make([]Level, 0, len(lines))
		filteredLines := make([][]byte, 0, len(lines))

		for i := range lines {
			lvl := levels[i]
			if lvl >= dest.level {
				filteredLevels = append(filteredLevels, lvl)
				filteredLines = append(filteredLines, lines[i])
			}
		}

		// Send to destination
		if len(filteredLines) > 0 {
			dest.backend.Send(levels, lines)
		}
	}

	return lgr
}
