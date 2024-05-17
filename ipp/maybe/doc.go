// MFP   - Miulti-Function Printers and scanners toolkit
// MAYBE - Go Maybe type for IPP values
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Package documentation

// Package maybe implements a kind of [Maybe (a.k.a. Option) type]
// in Go for IPP values.
//
// Some IPP attributes may either have a particular value, or be absent
// with a reason. IPP represents such a values with Tag that provides
// some information why value is missed.
//
// RFC 8011 calls this mechanism "Out-of-Band Values", see section
// 5.1.1 for details.
//
// Three IPP tags are defined for this purpose:
//
//   'unknown':     The attribute is supported by the IPP object, but the
//                  value is unknown to the IPP object for some reason.
//
//   'unsupported': The attribute is unsupported by the IPP object.
//
//   'no-value':    The attribute is supported by the IPP object, but the
//                  Administrator has not yet configured a value.
//
// This package implements a kind of Maybe (a.k.a. Option) type
// for all IPP values, that allows to represent such an optional
// values which either have a particular value or reason why value
// is missed.
//
// For example, "date-time-at-processing" attribute, which RFC 8011
// defines as (dateTime|unknown|no-value), can be represented in Go
// with the following variable (or structure field):
//
//   DateTimeAtProcessing maybe.Time
//
// It's value can be easily assigned:
//
//   DateTimeAtProcessing = maybe.NewTime(time.Now()) // Assign particular value
//   DateTimeAtProcessing = maybe.NoValye             // IPP 'no-value'
//
// And can be easily accessed:
//
//   tm, err = DateTimeAtProcessing.Time()
//
// The Time() method returns either particular value, if value is set,
// of one of the three following errors:
//
//   ErrNoValue     - value absent with IPP 'no-value' tag
//   ErrUnknown     - value absent with IPP 'unknown' tag
//   ErrUnsupported - value absent with IPP 'unsupported' tag
//
// The following variants of the Maybe type are implemented: [Binary],
// [Boolean], [Collection], [Integer], [Range], [Resolution], [String],
// [TextWithLang] and [Time].
//
// For each of these types, named Typename, the following functions
// exist:
//
//   NewTypename(value sometype) Typename    - the constructur
//   (Typename) Typename() (sometype, error) - returns underlying value
//
// The predefined values [NoValue], [Unknown] and [Unsupported] are
// assignable to any of these types.
//
// [Maybe (a.k.a. Option) type]: https://en.wikipedia.org/wiki/Option_type
package maybe
