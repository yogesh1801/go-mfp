// MFP - Miulti-Function Printers and scanners toolkit
// WSD core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// XML-related errors

package wsd

import (
	"errors"
	"strings"

	"github.com/alexpevzner/mfp/xmldoc"
)

// xmlErr represents error, related to the XML processing
type xmlErr struct {
	path []string // Path from XML root to the relevant element
	err  error    // Underlying error
}

// Error returns error string.
func (xe xmlErr) Error() string {
	return strings.Join(xe.path, "/") + ": " + xe.err.Error()
}

// Unwrap "unwraps" the error
func (xe xmlErr) Unwrap() error {
	return xe.err
}

// xmlErrWrap "wraps" the error in the context of the xmldoc.Element
func xmlErrWrap(elem xmldoc.Element, err error) error {
	return xmlErrWrapName(elem.Name, err)
}

// xmlErrWrap "wraps" the error in the context of the xmldoc.Attr
func xmlErrWrapAttr(attr xmldoc.Attr, err error) error {
	return xmlErrWrapName(attr.Name, err)
}

// xmlErrWrap "wraps" the error in the context of the xmldoc.Element
// or xmldoc.Attr with the specified name
func xmlErrWrapName(name string, err error) error {
	if err == nil {
		return nil
	}

	path := []string{name}

	if xe, ok := err.(xmlErr); ok {
		path = append(path, xe.path...)
		err = xe.err
	}

	return xmlErr{path: path, err: err}
}

// xmlErrNew is equal to xmlErrWrap(elem,errors.New(text))
func xmlErrNew(elem xmldoc.Element, text string) error {
	return xmlErrWrap(elem, errors.New(text))
}

// xmlErrMissed creates an error that happens when some
// required child element is missed
func xmlErrMissed(name string) error {
	return xmlErrWrapName(name, errors.New("missed"))
}
