// MFP - Miulti-Function Printers and scanners toolkit
// Logging facilities
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Discard backend

package log

// backendDiscard is the Backend that discards any output
type backendDiscard struct{}

// Line implements [Backend.Line] method
func (bk *backendDiscard) Send(levels []Level, lines [][]byte) {
}
