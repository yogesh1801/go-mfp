// MFP - Miulti-Function Printers and scanners toolkit
// Logging facilities
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Package documentation

package log

// trim returns a subslice of line by slicing off all trailing
// white space
func trim(line []byte) []byte {
	var cut int
loop:
	for cut = len(line); cut > 0; cut-- {
		switch line[cut-1] {
		case ' ', '\t', '\r':
		default:
			break loop
		}
	}

	return line[:cut]
}
