// MFP - Miulti-Function Printers and scanners toolkit
// XML mini library
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// XML-related errors

package xmldoc

import (
	"errors"
	"strings"
)

// XMLErr represents error, related to the XML processing.
//
// It consist of the arbitrary underlying error and XML path to the
// problematic Element or Attr.
type XMLErr struct {
	path []string // Path from XML root to the relevant element
	err  error    // Underlying error
}

// Error returns error string.
func (xe XMLErr) Error() string {
	return "/" + strings.Join(xe.path, "/") + ": " + xe.err.Error()
}

// Unwrap "unwraps" the error.
//
// It returns the underlying error (of its original type), undoing effect
// of all preceding wrapping with the [XMLErrWrap], [XMLErrWrapAttr] and
// [XMLErrWrapName] functions.
func (xe XMLErr) Unwrap() error {
	return xe.err
}

// XMLErrWrap "wraps" the error in the context of the [Element].
func XMLErrWrap(elem Element, err error) error {
	return XMLErrWrapName(elem.Name, err)
}

// XMLErrWrapAttr "wraps" the error in the context of the [Attr].
func XMLErrWrapAttr(attr Attr, err error) error {
	return XMLErrWrapName("@"+attr.Name, err)
}

// XMLErrWrapName "wraps" the error in the context of the [Element].
// or [Attr] with the specified name
func XMLErrWrapName(name string, err error) error {
	if err == nil {
		return nil
	}

	path := []string{name}

	if xe, ok := err.(XMLErr); ok {
		path = append(path, xe.path...)
		err = xe.err
	}

	return XMLErr{path: path, err: err}
}

// XMLErrNew is equal to XMLErrWrap(elem,errors.New(text)).
func XMLErrNew(elem Element, text string) error {
	return XMLErrWrap(elem, errors.New(text))
}

// XMLErrMissed creates an error that happens when some
// required child element is missed
func XMLErrMissed(name string) error {
	return XMLErrWrapName(name, errors.New("missed"))
}
