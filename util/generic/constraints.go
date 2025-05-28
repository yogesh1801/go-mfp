// MFP - Miulti-Function Printers and scanners toolkit
// Useful generics
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Useful constraints for generics

package generic

// Ordered is a constraint that permits any ordered type, i.e.,
// any type that supports <, <=, >= and > operations.
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}
