// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Errors

package abstract

// Error is the error code.
type Error int

// Standard error codes:
const (
	_ Error = iota
	ErrInvalidInput
	ErrUnsupportedInput
	ErrInvalidADFMode
	ErrUnsupportedADFMode
)

func (Error) Error() string { return "" }
