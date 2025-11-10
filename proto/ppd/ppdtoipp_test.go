// MFP - Multi-Function Printers and scanners toolkit
// PPD handling (libppd wrapper)
//
// Copyright (C) 2025 and up by GO-MFP authors.
// See LICENSE for license terms and conditions
//
// PPD->IPP conversion test

package ppd

import (
	_ "embed"
	"testing"
)

//go:embed "test.ppd"
var testPPD []byte

// TestPPDtoIPP tests PPD to IPP conversion
func TestPPDtoIPP(t *testing.T) {
	_, err := ToIPP(testPPD)
	if err != nil {
		t.Errorf("ToIPP: %s", err)
		return
	}
}
