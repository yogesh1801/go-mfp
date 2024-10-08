// MFP - Miulti-Function Printers and scanners toolkit
// WSD device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// XML-related errors

package wsdd

import (
	"strings"

	"github.com/alexpevzner/mfp/internal/xml"
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

// xmlErrWrap "wraps" the error in the context of the xml.Element
func xmlErrWrap(elem xml.Element, err error) error {
	path := []string{elem.Name}

	if xe, ok := err.(xmlErr); ok {
		path = append(path, xe.path...)
		err = xe.err
	}

	return xmlErr{path: path, err: err}
}
