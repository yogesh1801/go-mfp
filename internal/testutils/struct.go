// MFP - Miulti-Function Printers and scanners toolkit
// Utility functions and data BLOBs for testing
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Utilities for test with complex structures

package testutils

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Diff returns a pretty-printing report of the difference between
// two values. If values are equal, it returns empty string.
func Diff(x, y any) string {
	return cmp.Diff(x, y)
}

// CheckConvertionTest checks results of the data conversion test
// and in the case of mismatch, reports error into in the readable
// form into the [testing.T] and returns false.
//
// If test is passed, it returns true.
func CheckConvertionTest(t *testing.T,
	title, comment string,
	expected, present any) bool {

	diff := Diff(present, expected)
	if diff != "" {
		t.Errorf("\n"+
			"testing: %s\n"+
			"comment: %s\n"+
			"output mismatch:\n"+
			"%s", title, comment, diff)
		return false
	}

	return true
}
