// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Maybe (a.k.a. Optional) type implementation for IPP attributes.

package ipp

import (
	"reflect"
	"unsafe"

	"github.com/OpenPrinting/goipp"
)

// Maybe implements a [Maybe (a.k.a. Option) type] in Go for IPP values.
//
// Some IPP attributes may either have a particular value or be absent
// with a specific reason. IPP represents such values as Void values with Tags
// that provide information about why the value is missing.
//
// [RFC 8011] calls this mechanism "Out-of-Band Values"; see
// [RFC 8011, 5.1.1] for details.
//
// Three IPP tags are defined for this purpose:
//
//	'unknown':     The attribute is supported by the IPP object, but the
//	               value is unknown to the IPP object for some reason.
//
//	'unsupported': The attribute is unsupported by the IPP object.
//
//	'no-value':    The attribute is supported by the IPP object, but the
//	               Administrator has not yet configured a value.
//
// Maybe[T] can wrap any value that can be used with IPP and has the following
// states:
//
//  1. Value is not set and Tag is [goipp.TagZero]. In this case, the Attribute
//     will not be sent when encoding. Missing Attributes are represented
//     this way when IPP data is decoded. This is the initial (zero) state
//     of a Maybe[T] field.
//
//  2. Value is not set and Tag is not zero. In this case, the Attribute is encoded
//     as a Void value with the appropriate Tag.
//
//  3. Value is set. In this case, the Attribute is encoded normally, according to
//     the underlying value type and the "ipp:" struct tag.
//
// [RFC 8011]: https://datatracker.ietf.org/doc/html/rfc8011
// [RFC 8011, 5.1.1]: https://datatracker.ietf.org/doc/html/rfc8011#section-5.1.1
// [Maybe (a.k.a. Option) type]: https://en.wikipedia.org/wiki/Option_type
type Maybe[T any] struct {
	value  T         // Underlying value
	reason goipp.Tag // IPP tag to indicate absent value
	ok     bool      // True if value is set
}

// maybeCodecInterface implemented by any Maybe[T] and provided
// integration with the ippCodec.
type maybeCodecInterface interface {
	// typeof returns reflect.Type of the underlying value.
	typeof() reflect.Type

	// encode encodes Maybe[T] value as goipp.Values.
	//
	// Please note, it is called with nil receiver, and
	// actual Maybe[T] address is passed via 'p' parameter.
	encode(p unsafe.Pointer,
		encodeVal func(unsafe.Pointer) goipp.Values) goipp.Values

	// decode decodes Maybe[T] value from the goipp.Values.
	//
	// Please note, it is called with nil receiver, and
	// actual Maybe[T] address is passed via 'p' parameter.
	decode(p unsafe.Pointer, vals goipp.Values,
		decodeVal func(unsafe.Pointer, goipp.Values) error) error
}

// MaybeSet creates new non-empty Maybe[T] value.
func MaybeSet[T any](value T) (m Maybe[T]) {
	m.Set(value)
	return
}

// MaybeDel creates new empty Maybe[T] value.
func MaybeDel[T any](reason goipp.Tag) (m Maybe[T]) {
	m.Del(reason)
	return
}

// Get returns underlying value.
func (m Maybe[T]) Get() (value T, ok bool) {
	return m.value, m.ok
}

// Set sets underlying value.
func (m *Maybe[T]) Set(value T) {
	m.value = value
	m.reason = goipp.TagZero
	m.ok = true
}

// Del deletes underlying value and saves the reason.
func (m *Maybe[T]) Del(reason goipp.Tag) {
	var zero T
	m.value = zero
	m.reason = reason
	m.ok = false
}

// typeof returns reflect.Type of the underlying value.
func (m *Maybe[T]) typeof() reflect.Type {
	var v T
	return reflect.TypeOf(v)
}

// encode encodes Maybe[T] value as goipp.Values.
//
// It wraps encoder for the underlying value, which needs to
// be provided by caller.
//
// Please note, it is called with nil receiver, and
// actual Maybe[T] address is passed via 'p' parameter.
func (m *Maybe[T]) encode(p unsafe.Pointer,
	encodeVal func(unsafe.Pointer) goipp.Values) goipp.Values {

	m = (*Maybe[T])(p)

	switch {
	case m.ok:
		return encodeVal(unsafe.Pointer(&m.value))
	case m.reason != goipp.TagZero:
		return goipp.Values{{m.reason, goipp.Void{}}}
	}

	return nil
}

// decode decodes Maybe[T] value from the goipp.Values.
//
// It wraps decoder for the underlying value, which needs to
// be provided by caller.
//
// Please note, it is called with nil receiver, and
// actual Maybe[T] address is passed via 'p' parameter.
func (m *Maybe[T]) decode(p unsafe.Pointer, vals goipp.Values,
	decodeVal func(unsafe.Pointer, goipp.Values) error) error {

	m = (*Maybe[T])(p)

	switch tag := vals[0].T; tag {
	case goipp.TagNoValue, goipp.TagUnknown, goipp.TagUnsupportedValue:
		m.Del(tag)
		return nil
	}

	err := decodeVal(unsafe.Pointer(&m.value), vals)
	if err == nil {
		m.ok = true
	}

	return err
}
