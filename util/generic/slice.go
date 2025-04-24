// MFP - Miulti-Function Printers and scanners toolkit
// Useful generics
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Package documentation

package generic

import "sort"

// Ordered is a constraint that permits any ordered type, i.e.,
// any type that supports <, <=, >= and > operations
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

// CopySlice returns a shallow copy of the slice
func CopySlice[S ~[]T, T any](s S) (clone S) {
	if s != nil {
		clone = make(S, len(s))
		copy(clone, s)
	}
	return
}

// SortSlice sorts slice of [Ordered] elements.
func SortSlice[S ~[]T, T Ordered](s S) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}

// ConcatSlices returns the concatenation of some slices.
func ConcatSlices[S ~[]T, T any](slices ...S) S {
	size := 0

	for _, s := range slices {
		size += len(s)
		if size < 0 {
			panic("len out of range")
		}
	}

	out := make(S, 0, size)
	for _, s := range slices {
		out = append(out, s...)
	}

	return out
}

// EqualSlices tells if two slices are equal (the same length and
// all elements are equal)
func EqualSlices[S ~[]T, T comparable](s1, s2 S) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}
