// MFP   - Miulti-Function Printers and scanners toolkit
// MAYBE - Go Maybe type for IPP values
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Predefined errors

package maybe

import "errors"

// This package defines the following errors:
var (
	// ErrNoValue returned when value with the goipp.TagNoValue IPP tag
	// is accessed.
	ErrNoValue = errors.New("no-value")

	// ErrUnknown returned when value with the goipp.TagUnknown IPP tag
	// is accessed.
	ErrUnknown = errors.New("unknown")

	// ErrUnknown returned when value with the goipp.TagUnsupportedValue
	// IPP tag is accessed.
	ErrUnsupported = errors.New("unsupported")
)
