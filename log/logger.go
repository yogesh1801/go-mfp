// MFP - Miulti-Function Printers and scanners toolkit
// Logging facilities
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Package documentation

package log

import (
	"bytes"
	"encoding"
	"fmt"
	"os"
)

// DefaultLogger is the default logging destination.
var DefaultLogger = NewLogger("", LevelAll, Console)

// Logger is the logging destination.
// It can be connected to console, to the disk file etc...
type Logger struct {
	prefix string
	out    []loggerDest
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

// Trace writes a Trace-level message to the Logger.
func (lgr *Logger) Trace(format string, v ...any) *Logger {
	return lgr.format(LevelTrace, format, v...)
}

// Debug writes a Debug-level message to the Logger.
func (lgr *Logger) Debug(format string, v ...any) *Logger {
	return lgr.format(LevelDebug, format, v...)
}

// Info writes a Info-level message to the Logger.
func (lgr *Logger) Info(format string, v ...any) *Logger {
	return lgr.format(LevelInfo, format, v...)
}

// Error writes a Error-level message to the Logger.
func (lgr *Logger) Error(format string, v ...any) *Logger {
	return lgr.format(LevelError, format, v...)
}

// Fatal writes a Fatal-level message to the Logger.
//
// It calls os.Exit(1) and never returns.
func (lgr *Logger) Fatal(format string, v ...any) {
	lgr.format(LevelFatal, format, v...)
	os.Exit(1)
}

// Object writes any object that implements [encoding.TextMarshaler]
// interface to the Logger.
//
// If [encoding.TextMarshaler.MarshalText] returns an error, it
// will be written to log with the [Error] log level, regardless
// of the level specified by the first parameter.
func (lgr *Logger) Object(level Level, obj encoding.TextMarshaler) *Logger {
	text, err := obj.MarshalText()
	if err != nil {
		return lgr.Error("%s", err)
	}

	return lgr.text(level, text)
}

// format writes a single formatted message to the Logger
func (lgr *Logger) format(level Level, format string, v ...any) *Logger {
	buf := bufAlloc()
	defer bufFree(buf)

	fmt.Fprintf(buf, format, v...)
	return lgr.text(level, buf.Bytes())
}

// text writes a text message to the Logger
func (lgr *Logger) text(level Level, text []byte) *Logger {
	lines := bytes.Split(text, []byte("\n"))
	levels := make([]Level, len(lines))

	if lgr.prefix != "" {
		for i := range lines {
			lines[i] = []byte(lgr.prefix + string(lines[i]))
		}
	}

	for i := range lines {
		levels[i] = level
	}

	for _, dest := range lgr.out {
		if level >= dest.level {
			dest.backend.Send(levels, lines)
		}
	}

	return lgr
}
