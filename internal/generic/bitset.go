// MFP - Miulti-Function Printers and scanners toolkit
// Useful generic types
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Generic bitsets

package generic

import (
	"fmt"
	"strings"
)

// Bitset represents a bitset of instances of some integer type T
type Bitset[T BitsetElem] int

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

	bits &= 0x7fffffff
	for elem := T(0); bits != 0; elem++ {
		if bits.Contains(elem) {
			s = append(s, elem.String())
			bits.Del(elem)
		}
	}

	return strings.Join(s, ",")
}

// Add adds [ColorMode] to the set.
func (bits *Bitset[T]) Add(elem T) {
	*bits |= 1 << elem
}

// Del deletes [ColorMode] from the set.
func (bits *Bitset[T]) Del(elem T) {
	*bits &^= 1 << elem
}

// Contains reports if [ColorMode] exists in the set.
func (bits Bitset[T]) Contains(elem T) bool {
	return (bits & (1 << elem)) != 0
}
