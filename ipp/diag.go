// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Miscellaneous functions for debugging

package ipp

import "reflect"

// diagTypeName formats type name in a form, suitable
// for diagnostic purposes
func diagTypeName(t reflect.Type) string {
	if t.Kind() == reflect.Struct && t.Name() == "" {
		if t.Size() == 0 {
			return "struct {}"
		}

		return "struct {...}"
	}

	return t.String()
}
