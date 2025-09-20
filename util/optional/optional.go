// MFP - Miulti-Function Printers and scanners toolkit
// Optional value
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Optional values

package optional

// Val represents optional value of any type.
//
// The optional value is represented as a pointer to the corresponding
// type. If pointer is nil, value is missed.
type Val[T any] *T

// New returns a new optional value.
func New[T any](v T) *T {
	return &v
}

// NotZero returns [New](v) for any non-zero values of v, nil otherwise.
func NotZero[T comparable](v T) *T {
	var zero T
	if v != zero {
		return New(v)
	}
	return nil
}

// Get returns [Val]'s value. If Val is nil, T's zero value is returned.
func Get[T any](opt Val[T]) T {
	var v T
	if opt != nil {
		v = *opt
	}
	return v
}
