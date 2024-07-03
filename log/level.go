// MFP - Miulti-Function Printers and scanners toolkit
// Logging facilities
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Log levels

package log

// Level enumerates possible log levels
type Level int

// Log levels:
const (
	LevelTrace Level = iota // Protocol trace
	LevelDebug              // Debug messages
	LevelInfo               // Informational messages
	LevelError              // Error messages
	LevelFatal              // Fatal errors
)
