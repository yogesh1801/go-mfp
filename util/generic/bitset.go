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
	"strings"
	"sync/atomic"
)

// Bitset represents a bitset of instances of some integer type T.
// Operations with Bitset are goroutine-safe.
type Bitset[T BitsetElem] uint32

// BitsetElem is the member type for Bitset[T].
type BitsetElem interface {
	~int
	fmt.Stringer
}

// MakeBitset makes [Bitset] from the list of elements of type T.
func MakeBitset[T BitsetElem](list ...T) Bitset[T] {
	var set Bitset[T]

	for _, elem := range list {
		set.Add(elem)
	}

	return set
}

// String returns a string representation of the [Bitset],
// for debugging.
func (set Bitset[T]) String() string {
	s := make([]string, 0, 31)

	for elem := T(0); set != 0; elem++ {
		if set.Contains(elem) {
			s = append(s, elem.String())
			set.Del(elem)
		}
	}

	return strings.Join(s, ",")
}

// Add adds element to the set.
// It returns true if element was actually added.
func (set *Bitset[T]) Add(elem T) bool {
	mask := uint32(1) << elem
	old := atomic.OrUint32((*uint32)(set), mask)
	return old&mask == 0
}

// Del deletes element from the set.
// It returns true if element was actually deleted.
func (set *Bitset[T]) Del(elem T) bool {
	mask := uint32(1) << elem
	old := atomic.AndUint32((*uint32)(set), ^mask)
	return old&mask != 0
}

// Contains reports if element exists in the set.
func (set Bitset[T]) Contains(elem T) bool {
	mask := uint32(1) << elem
	return uint32(set)&mask != 0
}
