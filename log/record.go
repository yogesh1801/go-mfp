// MFP - Miulti-Function Printers and scanners toolkit
// Logging facilities
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Log records

package log

import (
	"bytes"
	"encoding"
	"fmt"
	"os"
	"slices"
	"sync"
)

// Record allows to build a multi-line log message, which
// will be written atomically, and will not be intermixed
// with logging activity from other goroutines and/or split
// between different files during log rotation.
type Record struct {
	parent *Logger    // Parent logger
	lines  [][]byte   // Collected lines
	levels []Level    // Corresponding levels
	mutex  sync.Mutex // Access lock
}

// Commit writes Record to the parent [Logger].
func (rec *Record) Commit() *Logger {
	rec.parent.send(rec.levels, rec.lines)
	return rec.parent
}

// Trace writes a Trace-level message to the Record.
func (rec *Record) Trace(format string, v ...any) *Record {
	return rec.format(LevelTrace, format, v...)
}

// Debug writes a Debug-level message to the Record.
func (rec *Record) Debug(format string, v ...any) *Record {
	return rec.format(LevelDebug, format, v...)
}

// Info writes a Info-level message to the Record.
func (rec *Record) Info(format string, v ...any) *Record {
	return rec.format(LevelInfo, format, v...)
}

// Warning writes a Warning-level message to the Record.
func (rec *Record) Warning(format string, v ...any) *Record {
	return rec.format(LevelWarning, format, v...)
}

// Error writes a Error-level message to the Record.
func (rec *Record) Error(format string, v ...any) *Record {
	return rec.format(LevelError, format, v...)
}

// Fatal writes a Fatal-level message to the Record.
//
// It calls os.Exit(1) and never returns.
func (rec *Record) Fatal(format string, v ...any) {
	rec.format(LevelFatal, format, v...)
	rec.Commit()
	os.Exit(1)
}

// Object writes any object that implements [encoding.TextMarshaler]
// interface to the Record.
//
// If [encoding.TextMarshaler.MarshalText] returns an error, it
// will be written to log with the [Error] log level, regardless
// of the level specified by the first parameter.
func (rec *Record) Object(level Level, obj encoding.TextMarshaler) *Record {
	text, err := obj.MarshalText()
	if err != nil {
		return rec.Error("%s", err)
	}

	return rec.text(level, text)
}

// format writes a single formatted message to the Record
func (rec *Record) format(level Level, format string, v ...any) *Record {
	buf := bufAlloc()
	defer bufFree(buf)

	fmt.Fprintf(buf, format, v...)
	return rec.text(level, slices.Clone(buf.Bytes()))
}

// text writes a text message to the Record
func (rec *Record) text(level Level, text []byte) *Record {
	lines := bytes.Split(text, []byte("\n"))
	levels := make([]Level, len(lines))

	for i := range lines {
		levels[i] = level
	}

	rec.mutex.Lock()
	rec.lines = append(rec.lines, lines...)
	rec.levels = append(rec.levels, levels...)
	rec.mutex.Unlock()

	return rec
}
