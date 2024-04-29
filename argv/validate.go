// MFP  - Miulti-Function Printers and scanners toolkit
// argv - Argv parsing mini-library
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Value validators

package argv

import (
	"fmt"
	"math"
	"strconv"
)

// ValidateAny accepts any string
func ValidateAny(string) error {
	return nil
}

// ValidateStrings returns validator that accepts any of supplied strings
func ValidateStrings(s []string) func(string) error {
	// Create a copy of input, to protect from callers
	// that may change the slice after the call.
	set := make([]string, len(s))
	copy(set, s)

	// Create validator
	return func(in string) error {
		for _, member := range set {
			if in == member {
				return nil
			}
		}

		return fmt.Errorf("%q: invalid argument", in)
	}
}

// ValidateIntRange returns validator, that accepts signed integer
// in range (min <= x && x <= max). If base is 0, the actual base
// is implied automatically based on a string prefix following the
// sign (2 for "0b", 8 for "0" or "0o", 16 for "0x", and 10 otherwise).
func ValidateIntRange(base int, min, max int64) func(string) error {
	return func(in string) error {
		v, err := strconv.ParseInt(in, base, 64)
		if err != nil {
			return fmt.Errorf("%q: invalid integer", in)
		}

		if v < min || v > max {
			return fmt.Errorf(
				"%q: value out of range (%d...%d)",
				in, min, max)
		}

		return nil
	}
}

// ValidateUintRange returns validator, that accepts unsigned integer
// in range (min <= x && x <= max). If base is 0, the actual base
// is implied automatically based on a string prefix (2 for "0b",
// 8 for "0" or "0o", 16 for "0x", and 10 otherwise).
//
// It is like ValidateIntRange but for unsigned numbers. The sign
// prefix is not accepted.
func ValidateUintRange(base int, min, max uint64) func(string) error {
	return func(in string) error {
		v, err := strconv.ParseUint(in, base, 64)
		if err != nil {
			return fmt.Errorf("%q: invalid integer", in)
		}

		if v < min || v > max {
			return fmt.Errorf(
				"%q: value out of range (%d...%d)",
				in, min, max)
		}

		return nil
	}
}

// ValidateIntBits returns validator, that accepts signed integer
// that fits specified number of bits (1 to 64).  If base is 0, the
// actual base is implied automatically based on a string prefix
// (2 for "0b", 8 for "0" or "0o", 16 for "0x", and 10 otherwise).
//
// It panics, if number of bits is out of range
func ValidateIntBits(base, bits int) func(string) error {
	if bits < 1 || bits > 64 {
		err := fmt.Errorf(
			"ValidateIntBits: bits (%d) out of range (1...64)",
			bits)
		panic(err)
	}

	var min, max int64
	switch bits {
	case 1:
		min, max = 0, 1
	default:
		max = int64(1<<uint64(bits-1) - 1)
		min = -max - 1
	}

	return func(in string) error {
		v, err := strconv.ParseInt(in, base, 64)
		if err != nil {
			return fmt.Errorf("%q: invalid integer", in)
		}

		if v < min || v > max {
			return fmt.Errorf(
				"%q: value doesn't fit %d bits", in, bits)
		}

		return nil
	}
}

// ValidateUintBits returns validator, that accepts unsigned integer
// that fits specified number of bits (1 to 64).  If base is 0, the
// actual base is implied automatically based on a string prefix
// (2 for "0b", 8 for "0" or "0o", 16 for "0x", and 10 otherwise).
//
// It panics, if number of bits is out of range
func ValidateUintBits(base, bits int) func(string) error {
	if bits < 1 || bits > 64 {
		err := fmt.Errorf(
			"ValidateUintBits: bits (%d) out of range (1...64)",
			bits)
		panic(err)
	}

	var max uint64
	switch bits {
	case 64:
		max = math.MaxUint64

	default:
		max = (1 << uint64(bits)) - 1
	}

	return func(in string) error {
		v, err := strconv.ParseUint(in, base, 64)
		if err != nil {
			return fmt.Errorf("%q: invalid integer", in)
		}

		if v > max {
			return fmt.Errorf(
				"%q: value doesn't fit %d bits", in, bits)
		}

		return nil
	}
}
