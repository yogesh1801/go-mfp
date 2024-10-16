// MFP - Miulti-Function Printers and scanners toolkit
// Utility functions for random numbers
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Utility functions for random numbers

package random

import (
	"crypto/rand"
	"fmt"
	"unsafe"
)

// Fill byte slice with random bytes
func Fill(b []byte) {
	_, err := rand.Read(b)
	if err != nil {
		panic(fmt.Sprintf("Failed to read %d rand bytes: %v",
			len(b), err))
	}
}

// Uint returns random unsigned integer
func Uint() uint {
	var v uint
	s := (*[unsafe.Sizeof(v)]byte)(unsafe.Pointer(&v))
	Fill(s[:])
	return v
}

// UintMax returns random unsigned integer in range [0...max], inclusively.
func UintMax(max uint) uint {
	if max == 0 {
		return 0
	}

	mask := max
	for tmp := mask >> 1; tmp != 0; tmp >>= 1 {
		mask |= tmp
	}

	for {
		tmp := Uint() & mask
		if tmp <= max {
			return tmp
		}
	}
}

// UintRange returns random unsigned integer in range [min...max-1],
// inclusively.
func UintRange(min, max uint) uint {
	// Normalize the range
	if min > max {
		min, max = max, min
	}

	// Generate random number
	return min + UintMax(max-min)
}
