// MFP   - Miulti-Function Printers and scanners toolkit
// MAYBE - Go Maybe type for IPP values
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// maybe types implementation

package maybe

import (
	"errors"
	"time"

	"github.com/OpenPrinting/goipp"
)

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

// ----- Binary -----

// Binary represents a binary value.
//
// Use (Binary) Binary() to obtain underlying value.
type Binary interface{ Binary() ([]byte, error) }

// NewBinary creates a new Binary value.
func NewBinary(v []byte) Binary {
	return implBinary(v)
}

// implBinary implements Binary interface.
type implBinary []byte

// Binary returns underlying value.
func (impl implBinary) Binary() ([]byte, error) {
	return []byte(impl), nil
}

// ----- Boolean -----

// Boolean represents a bool value.
//
// Use (Boolean) Boolean() to obtain underlying value.
type Boolean interface{ Boolean() (bool, error) }

// NewBoolean creates a new Boolean value.
func NewBoolean(v bool) Boolean {
	return implBoolean(v)
}

// implBoolean implements Boolean interface.
type implBoolean bool

// Boolean returns underlying value.
func (impl implBoolean) Boolean() (bool, error) {
	return bool(impl), nil
}

// ----- Collection -----

// Collection represents a [goipp.Collection] value.
//
// Use (Collection) Collection() to obtain underlying value.
type Collection interface {
	Collection() (goipp.Collection, error)
}

// NewCollection creates a new Collection value.
func NewCollection(v goipp.Collection) Collection {
	return implCollection(v)
}

// implCollection implements Collection interface.
type implCollection goipp.Collection

// Collection returns underlying value.
func (impl implCollection) Collection() (goipp.Collection, error) {
	return goipp.Collection(impl), nil
}

// ----- Integer -----

// Integer represents a int32 value.
//
// Use (Integer) Integer() to obtain underlying value.
type Integer interface{ Integer() (int32, error) }

// NewInteger creates a new Integer value.
func NewInteger(v int32) Integer {
	return implInteger(v)
}

// implInteger implements Integer interface.
type implInteger int32

// Integer returns underlying value.
func (impl implInteger) Integer() (int32, error) {
	return int32(impl), nil
}

// ----- Range -----

// Range represents a [goipp.Range] value.
//
// Use (Range) Range() to obtain underlying value.
type Range interface{ Range() (goipp.Range, error) }

// NewRange creates a new Range value.
func NewRange(v goipp.Range) Range {
	return implRange(v)
}

// implRange implements Range interface.
type implRange goipp.Range

// Range returns underlying value.
func (impl implRange) Range() (goipp.Range, error) {
	return goipp.Range(impl), nil
}

// ----- Resolution -----

// Resolution represents a [goipp.Resolution] value.
//
// Use (Resolution) Resolution() to obtain underlying value.
type Resolution interface {
	Resolution() (goipp.Resolution, error)
}

// NewResolution creates a new Resolution value.
func NewResolution(v goipp.Resolution) Resolution {
	return implResolution(v)
}

// implResolution implements Resolution interface.
type implResolution goipp.Resolution

// Resolution returns underlying value.
func (impl implResolution) Resolution() (goipp.Resolution, error) {
	return goipp.Resolution(impl), nil
}

// ----- String -----

// String represents a string value.
//
// Use (String) String() to obtain underlying value.
type String interface{ String() (string, error) }

// NewString creates a new String value.
func NewString(v string) String { return implString(v) }

// implString implements String interface.
type implString string

// String returns underlying value.
func (impl implString) String() (string, error) {
	return string(impl), nil
}

// ----- TextWithLang -----

// TextWithLang represents a [goipp.TextWithLang] value.
//
// Use (TextWithLang) TextWithLang() to obtain underlying value.
type TextWithLang interface {
	TextWithLang() (goipp.TextWithLang, error)
}

// NewTextWithLang creates a new TextWithLang value.
func NewTextWithLang(v goipp.TextWithLang) TextWithLang {
	return implTextWithLang(v)
}

// implTextWithLang implements TextWithLang interface.
type implTextWithLang goipp.TextWithLang

// TextWithLang returns underlying value.
func (impl implTextWithLang) TextWithLang() (goipp.TextWithLang, error) {
	return goipp.TextWithLang(impl), nil
}

// ----- Time -----

// Time represents a [time.Time] value.
//
// Use (Time) Time() to obtain underlying value.
type Time interface{ Time() (time.Time, error) }

// NewTime creates a new Time value.
func NewTime(v time.Time) Time {
	return implTime(v)
}

// implTime implements Time interface.
type implTime time.Time

// Time returns underlying value.
func (impl implTime) Time() (time.Time, error) {
	return time.Time(impl), nil
}

// ----- Void value types -----

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
	NoValue = noValue{}

	// Unknown is the missed value marked with goipp.TagUnknown tag.
	Unknown = unknown{}

	// Unsupported is the missed value marked with
	// goipp.TagUnsupportedValue tag.
	Unsupported = unsupported{}
)

// noValue type represents missed value marked with goipp.TagNoValue tag.
type noValue struct{}

func (noValue) Binary() ([]byte, error) {
	return nil, ErrNoValue
}

func (noValue) Boolean() (bool, error) {
	return false, ErrNoValue
}

func (noValue) Collection() (goipp.Collection, error) {
	return nil, ErrNoValue
}

func (noValue) Integer() (int32, error) {
	return 0, ErrNoValue
}

func (noValue) Range() (goipp.Range, error) {
	return goipp.Range{}, ErrNoValue
}

func (noValue) Resolution() (goipp.Resolution, error) {
	return goipp.Resolution{}, ErrNoValue
}

func (noValue) String() (string, error) {
	return "", ErrNoValue
}

func (noValue) TextWithLang() (goipp.TextWithLang, error) {
	return goipp.TextWithLang{}, ErrNoValue
}

func (noValue) Time() (time.Time, error) {
	return time.Time{}, ErrNoValue
}

// unknown type represents missed value marked with goipp.TagUnknown tag.
type unknown struct{}

func (unknown) Binary() ([]byte, error) {
	return nil, ErrUnknown
}

func (unknown) Boolean() (bool, error) {
	return false, ErrUnknown
}

func (unknown) Collection() (goipp.Collection, error) {
	return nil, ErrUnknown
}

func (unknown) Integer() (int32, error) {
	return 0, ErrUnknown
}

func (unknown) Range() (goipp.Range, error) {
	return goipp.Range{}, ErrUnknown
}

func (unknown) Resolution() (goipp.Resolution, error) {
	return goipp.Resolution{}, ErrUnknown
}

func (unknown) String() (string, error) {
	return "", ErrUnknown
}

func (unknown) TextWithLang() (goipp.TextWithLang, error) {
	return goipp.TextWithLang{}, ErrUnknown
}

func (unknown) Time() (time.Time, error) {
	return time.Time{}, ErrUnknown
}

// unsupported type represents missed value marked with goipp.TagNoValue tag.
type unsupported struct{}

func (unsupported) Binary() ([]byte, error) {
	return nil, ErrUnsupported
}

func (unsupported) Boolean() (bool, error) {
	return false, ErrUnsupported
}

func (unsupported) Collection() (goipp.Collection, error) {
	return nil, ErrUnsupported
}

func (unsupported) Integer() (int32, error) {
	return 0, ErrUnsupported
}

func (unsupported) Range() (goipp.Range, error) {
	return goipp.Range{}, ErrUnsupported
}

func (unsupported) Resolution() (goipp.Resolution, error) {
	return goipp.Resolution{}, ErrUnsupported
}

func (unsupported) String() (string, error) {
	return "", ErrUnsupported
}

func (unsupported) TextWithLang() (goipp.TextWithLang, error) {
	return goipp.TextWithLang{}, ErrUnsupported
}

func (unsupported) Time() (time.Time, error) {
	return time.Time{}, ErrUnsupported
}
