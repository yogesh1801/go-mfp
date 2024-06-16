// MFP - Miulti-Function Printers and scanners toolkit
// Utility functions and data BLOBs for testing
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Package documentation

package testutils

import "testing"

// TestIppMustParse tests ippMustParse function
func TestIppMustParse(t *testing.T) {
	defer func() {
		p := recover()
		if err, ok := p.(error); ok {
			if err.Error() == "Message truncated at 0x0" {
				return
			}
		}

		t.Errorf("%s", p)
	}()

	// Must panic on invalid input
	ippMustParse([]byte{})
	panic("ippMustParse must panic on invalid input")
}
