// MFP   - Miulti-Function Printers and scanners toolkit
// MAYBE - Go Maybe type for IPP values
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// void type (represents maybe types with absent values)

package maybe

import (
	"time"

	"github.com/OpenPrinting/goipp"
)

// This package defines the following values to represent
// a "missed value" with appropriate IPP tag.
//
// Each of these values is assignable to any value type
// defined in this package (Binary, Boolean, ...).
//
// Any attempt to get its value returns a zero value of the
// particular type accompanied with the appropriate (ErrNoValue,
// ErrUnknown or ErrUnsupported) error.
var (
	// NoValue is the missed value marked with goipp.TagNoValue tag.
	NoValue = void{err: ErrNoValue, tag: goipp.TagNoValue}

	// Unknown is the missed value marked with goipp.TagUnknown tag.
	Unknown = void{err: ErrUnknown, tag: goipp.TagUnknown}

	// Unsupported is the missed value marked with
	// goipp.TagUnsupportedValue tag.
	Unsupported = void{err: ErrUnsupported, tag: goipp.TagUnsupportedValue}
)

// void type used to represent value-less maybe value
type void struct {
	err error
	tag goipp.Tag
}

func (v void) Binary() ([]byte, error) {
	return nil, v.err
}

func (v void) Boolean() (bool, error) {
	return false, v.err
}

func (v void) Collection() (goipp.Collection, error) {
	return nil, v.err
}

func (v void) Integer() (int32, error) {
	return 0, v.err
}

func (v void) Range() (goipp.Range, error) {
	return goipp.Range{}, v.err
}

func (v void) Resolution() (goipp.Resolution, error) {
	return goipp.Resolution{}, v.err
}

func (v void) String() (string, error) {
	return "", v.err
}

func (v void) TextWithLang() (goipp.TextWithLang, error) {
	return goipp.TextWithLang{}, v.err
}

func (v void) Time() (time.Time, error) {
	return time.Time{}, v.err
}
