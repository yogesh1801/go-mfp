// MFP - Miulti-Function Printers and scanners toolkit
// Functions, missed from the older versions of the Go stdlib
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Missed functions

package missed

import "strings"

// StringsCutSuffix returns s without the suffix and reports whether
// the suffix was found in the string.
// If s doesn't end with the suffix, StringCutSuffix returns s, false.
//
// Newer version of Go exports this function as [strings.CutSuffix]
func StringsCutSuffix(s, suffix string) (string, bool) {
	if strings.HasSuffix(s, suffix) {
		return s[:len(s)-len(suffix)], true
	}
	return s, false
}
