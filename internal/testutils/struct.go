// MFP - Miulti-Function Printers and scanners toolkit
// Utility functions and data BLOBs for testing
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Utilities for test with complex structures

package testutils

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/kr/pretty"
)

// Diff returns a pretty-printing report of the difference between
// two values. If values are equal, it returns empty string.
func Diff(x, y any) string {
	return cmp.Diff(x, y, compareErrorsOption())
}

// compareErrorsOption returns an cmp.Option that tells
// Diff to use compareErrors for comparing error values
func compareErrorsOption() cmp.Option {
	return cmp.FilterValues(
		func(x, y interface{}) bool {
			_, ok1 := x.(error)
			_, ok2 := y.(error)
			return ok1 && ok2
		},
		cmp.Comparer(compareErrors),
	)
}

// compareErrors tells if two errors are equal.
//
// Errors that render to the same text message
// considered equal, even if underlying values
// are different.
func compareErrors(e1, e2 interface{}) bool {
	switch {
	case e1 == nil && e2 == nil:
		return true
	case e1 == nil || e2 == nil:
		return false
	}

	err1 := e1.(error)
	err2 := e2.(error)

	if errors.Is(err1, err2) || errors.Is(err2, err1) {
		return true
	}

	return err1.Error() == err2.Error()
}

// Format formats (pretty-prints) an arbitrary Go value.
func Format(v any) string {
	return pretty.Sprint(v)
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
