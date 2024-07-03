// MFP - Miulti-Function Printers and scanners toolkit
// Logging facilities
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Backend writes log messages to their final destination

package log

// Standard backends:
var (
	// Console writes output to console.
	Console Backend = &backendConsole{}

	// Discard silently discards any output.
	Discard Backend = &backendDiscard{}
)

// Backend represents a logging destination, which may be console,
// disk file, ...
type Backend interface {
	// Send writes some lines to the destination, represented
	// by the Backend.
	//
	// Backend may assume that levels and lines slices have
	// the same length and 1:1 correspondence each to other.
	//
	// Logically, these lines comprise a single log record.
	// If multiple goroutines write to log simultaneously,
	// this is Backend's responsibility to write them to the
	// destination in the sequential order and to avoid mixing
	// of unrelated records between each other.
	//
	// If Backend writes log to file and implements log rotation,
	// it should avoid rotation in the middle of some record.
	Send(levels []Level, lines [][]byte)
}
