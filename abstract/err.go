// MFP - Miulti-Function Printers and scanners toolkit
// Abstract definition for printer and scanner interfaces
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Errors

package abstract

import "fmt"

// ErrCode is the error code.
type ErrCode int

// Standard error codes:
const (
	_ ErrCode = iota
	ErrInvalidParam
	ErrUnsupportedParam
	ErrDocumentClosed
)

// Error returns error string. It implements the [error] interface.
func (e ErrCode) Error() string {
	switch e {
	case ErrInvalidParam:
		return "Invalid parameter"
	case ErrUnsupportedParam:
		return "Unsupported parameter"
	case ErrDocumentClosed:
		return "Document is closed"
	}
	return ""
}

// ErrParam used by functions like [ScannerRequest.Validate] to
// represent parameter error in the supplied request.
type ErrParam struct {
	Err   ErrCode // Underlying error
	Name  string  // Parameter name
	Value any     // Parameter value
}

// Error returns error string. It implements the [error] interface.
func (e ErrParam) Error() string {
	return fmt.Sprintf("%s %s: %v", e.Err, e.Name, e.Value)
}

// Unwrap unwraps the underlying parameter error.
func (e ErrParam) Unwrap() error {
	return e.Err
}
