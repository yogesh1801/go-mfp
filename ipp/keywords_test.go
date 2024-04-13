// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Tests for IPP keywords

package ipp

import "testing"

// TestKwPrinterStateReasons tests KwPrinterStateReasons methods
func TestKwPrinterStateReasons(t *testing.T) {
	testData := []struct{ input, reason, severity KwPrinterStateReasons }{
		{"media-low-warning", "media-low", "-warning"},
		{"media-jam-error", "media-jam", "-error"},
		{"media-needed-report", "media-needed", "-report"},
		{"unknown-unknown", "unknown-unknown", ""},
		{"unknown", "unknown", ""},
		{"", "", ""},
	}

	for _, data := range testData {
		reason := data.input.Reason()
		severity := data.input.Severity()

		if reason != data.reason || severity != data.severity {
			t.Errorf("%q: expected (%q,%q), present (%q, %q)",
				data.input,
				data.reason, data.severity,
				reason, severity,
			)

		}
	}
}
