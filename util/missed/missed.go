// MFP - Miulti-Function Printers and scanners toolkit
// Functions, missed from the older versions of the Go stdlib
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Missed functions

package missed

import "strings"

// StringsCutPrefix returns s without the prefix and reports whether
// the prefix was found in the string.
// If s doesn't end with the suffix, StringCutPrefix returns s, false.
//
// Newer version of Go exports this function as [strings.CutPrefix]
func StringsCutPrefix(s, prefix string) (string, bool) {
	if strings.HasPrefix(s, prefix) {
		return s[len(prefix):], true
	}
	return s, false
}

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
