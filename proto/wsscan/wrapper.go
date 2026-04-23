// MFP - Miulti-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by go-mfp authors.
// See LICENSE for license terms and conditions
//
// Wrapper for values with options

package wsscan

// Wrapper wraps some simple value and adds some optional parameters.
//
// The following wrappers currently exist:
//   - [ValWithOptions]
//   - [TextWithLangElement]
//
// The Wrapper interface provides set of helper methods
// to simplify serialization of these types.
type Wrapper interface {
	// HasOptions reports if value really has any options set.
	HasOptions() bool

	// Unwrap returns the wrapped simple value in the case
	// the Wrapper doesn't have any added options.
	//
	// If value cannot be unwrapped, it returns the original
	// value:
	//   TextWithLangElement{Text: "hello", Lang: nil} -> "hello"
	//   TextWithLangElement{Text: "hello", Lang: "en} -> TextWithLangElementP{...}
	Unwrap() any
}
