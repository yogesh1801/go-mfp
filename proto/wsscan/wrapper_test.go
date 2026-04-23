// MFP - Miulti-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by go-mfp authors.
// See LICENSE for license terms and conditions
//
// Wrapper tests

package wsscan

var (
	// These test verifies (at compile time) that
	// the following types implement the Wrapper
	// interface.
	_ = Wrapper(ValWithOptions[int]{})
	_ = Wrapper(TextWithLangElement{})
)
