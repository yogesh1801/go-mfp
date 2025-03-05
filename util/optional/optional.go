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
