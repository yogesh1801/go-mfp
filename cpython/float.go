// MFP - Miulti-Function Printers and scanners toolkit
// CPython binding.
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Floating point limits.

package cpython

import "math"

// The following variables define ranges of the float64
// values that can be correctly converted into int64/uint64
//
// These ranges are machine-specific and guessed dynamically
// at the initialization time
var (
	// maxInt64Float is the maximum float64 value that can be
	// correctly converted into the int64
	maxInt64Float float64

	// minInt64Float is the minimum float64 value that can be
	// correctly converted into the int64
	minInt64Float float64

	// maxUint64Float is the maximum float64 value that can be
	// correctly converted into the uint64
	maxUint64Float float64
)

// init initializes the maxInt64Float, minInt64Float and maxUint64Float
// variables
func init() {
	var i64 int64
	var u64 uint64

	for i64 = int64(math.MaxInt64); int64(float64(i64)) != i64; i64-- {
	}

	maxInt64Float = float64(i64)

	for i64 = int64(math.MinInt64); int64(float64(i64)) != i64; i64++ {
	}

	minInt64Float = float64(i64)

	for u64 = uint64(math.MaxUint64); uint64(float64(u64)) != u64; u64-- {
	}

	maxUint64Float = float64(u64)
}
