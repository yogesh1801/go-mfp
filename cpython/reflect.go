// MFP - Miulti-Function Printers and scanners toolkit
// CPython binding.
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Helper functions for reflect.Value

package cpython

import (
	"reflect"
	"sort"
)

// reflectSort sorts slice of reflect.Value in acceding order.
//
// Slice elements must be [cmp.Ordered] or bool.
// All elements of the slice must be of the same type;
//
// It returns true if slice was sorted, false otherwise.
func reflectSort(slice []reflect.Value) bool {
	// If slice is empty, we have nothing to do
	if len(slice) == 0 {
		return true
	}

	// Make sure all elements have the same type
	typ := slice[0].Type()
	for i := 1; i < len(slice); i++ {
		if slice[i].Type() != typ {
			return false
		}
	}

	// Choose the proper less function
	var less func(i, j int) bool

	switch typ.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Int64:

		less = func(i, j int) bool {
			return slice[i].Int() < slice[j].Int()
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
		reflect.Uint64:

		less = func(i, j int) bool {
			return slice[i].Uint() < slice[j].Uint()
		}

	case reflect.Float32, reflect.Float64:
		less = func(i, j int) bool {
			return slice[i].Float() < slice[j].Float()
		}

	case reflect.String:
		less = func(i, j int) bool {
			return slice[i].String() < slice[j].String()
		}

	case reflect.Bool:
		less = func(i, j int) bool {
			return !slice[i].Bool() && slice[j].Bool()
		}

	default:
		return false
	}

	// And now sort the slice
	sort.Slice(slice, less)

	return true
}
