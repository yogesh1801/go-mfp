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

// Bitset represents a bitset of instances of some integer type T
type Bitset[T BitsetElem] uint32

// BitsetElem is the member type for Bitset[T].
type BitsetElem interface {
	~int
	fmt.Stringer
}

// MakeBitset makes [Bitset] from the list of elements of type T.
func MakeBitset[T BitsetElem](list ...T) Bitset[T] {
	var bits Bitset[T]

	for _, elem := range list {
		bits.Add(elem)
	}

	return bits
}

// String returns a string representation of the [Bitset],
// for debugging.
func (bits Bitset[T]) String() string {
	s := make([]string, 0, 31)

	for elem := T(0); bits != 0; elem++ {
		if bits.Contains(elem) {
			s = append(s, elem.String())
			bits.Del(elem)
		}
	}

	return strings.Join(s, ",")
}

// Add adds element to the set.
// It returns true if element was actually added.
func (bits *Bitset[T]) Add(elem T) bool {
	mask := uint32(1) << elem
	old := atomic.OrUint32((*uint32)(bits), mask)
	return old&mask == 0
}

// Del deletes element from the set.
// It returns true if element was actually deleted.
func (bits *Bitset[T]) Del(elem T) bool {
	mask := uint32(1) << elem
	old := atomic.AndUint32((*uint32)(bits), ^mask)
	return old&mask != 0
}

// Contains reports if element exists in the set.
func (bits *Bitset[T]) Contains(elem T) bool {
	mask := uint32(1) << elem
	load := atomic.LoadUint32((*uint32)(bits))
	return load&mask != 0
}
