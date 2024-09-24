// MFP - Miulti-Function Printers and scanners toolkit
// Logging facilities
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Log Marshaler

package log

// Marshaler is the interface implemented by objects that can
// be written to the log, using [Object], [Logger.Object] and
// [Record.Object] functions.
type Marshaler interface {
	// LogMarshal returns a string representation of the
	// object, for logging. The returned string may be
	// multi-line, delimited with the '\n' characters.
	MarshalLog() []byte
}
