// MFP - Miulti-Function Printers and scanners toolkit
// Useful generics
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Generic bitsets

package generic

import (
	"fmt"
	"math/bits"
	"strings"
)

// Bitset represents a bitset of instances of some integer type T.
// Operations with Bitset are NOT goroutine-safe.
type Bitset[T ~int | ~uint] uint32

// MakeBitset makes [Bitset] from the list of elements of type T.
func MakeBitset[T ~int](list ...T) Bitset[T] {
	var set Bitset[T]

	for _, elem := range list {
		set.Add(elem)
	}

	return set
}

// String returns a string representation of the [Bitset], for debugging.
func (set Bitset[T]) String() string {
	s := make([]string, 0, 31)

	for elem := T(0); set != 0; elem++ {
		if set.Contains(elem) {
			s = append(s, fmt.Sprintf("%v", elem))
			set.Del(elem)
		}
	}

	return strings.Join(s, ",")
}

// Elements returns slice of the Bitset elements
func (set Bitset[T]) Elements() []T {
	out := make([]T, 0, bits.OnesCount32(uint32(set)))

	for elem := T(0); set != 0; elem++ {
		if set.Contains(elem) {
			out = append(out, elem)
			set.Del(elem)
		}
	}

	return out
}

// IsEmpty reports if set is empty.
func (set Bitset[T]) IsEmpty() bool {
	return set == 0
}

// Add adds element to the set.
// It returns true if element was actually added.
func (set *Bitset[T]) Add(elem T) bool {
	mask := Bitset[T](1) << elem
	old := *set
	*set |= mask
	return old&mask == 0
}

// Del deletes element from the set.
// It returns true if element was actually deleted.
func (set *Bitset[T]) Del(elem T) bool {
	mask := Bitset[T](1) << elem
	old := *set
	*set &= ^mask
	return old&mask != 0
}

// Contains reports if element exists in the set.
func (set Bitset[T]) Contains(elem T) bool {
	mask := uint32(1) << elem
	return uint32(set)&mask != 0
}
