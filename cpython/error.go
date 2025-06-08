// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Package documentation

package cpython

// Error represents a Python Error
type Error struct {
	msg string
}

// Error returns error message. It implements error interface.
func (e Error) Error() string {
	return e.msg
}
