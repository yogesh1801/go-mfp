// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// HTTP errors

package ipp

import (
	"fmt"
	"net/http"
)

// Common errors, reported as ErrHTTP:
var (
	ErrHTTPMethodNotAllowed = NewErrHTTP(http.StatusMethodNotAllowed, "")
)

// ErrHTTP represents HTTP error.
// It consist of the HTTP status and message text.
// Implements error interface.
type ErrHTTP struct {
	Status  int    // HTTP status
	Message string // Error message
}

// NewErrHTTP creates a new HTTP error.
// If msg is "", [http.StatusText] used instead.
func NewErrHTTP(code int, msg string) *ErrHTTP {
	if msg == "" {
		msg = http.StatusText(code)
	}

	return &ErrHTTP{
		Status:  code,
		Message: msg,
	}
}

// Error returns an error string. It implements [error] interface.
func (e *ErrHTTP) Error() string {
	return fmt.Sprintf("HTTP %3.3d %s", e.Status, e.Message)
}
