// MFP  - Miulti-Function Printers and scanners toolkit
// argv - Argv parsing mini-library
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Miscellaneous functions for names.

package argv

import (
	"unicode"
)

// nameCheck function verifies syntax of Options and
// Parameters names. Valid name starts with letter or
// digit and then consist of letters, digits and dash
// characters.
//
// It returns the first invalid character, if one is
// encountered, or -1 otherwise.
func nameCheck(name string) rune {
	for i, c := range name {
		switch {
		// Letters and digits always allowed
		case unicode.IsLetter(c) || unicode.IsDigit(c):

		// Dash allowed expect the very first character
		case i > 0 && c == '-':

		// Other characters not allowed
		default:
			return c
		}
	}

	return -1
}
