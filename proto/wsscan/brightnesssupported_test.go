// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// Test for brightness supported

package wsscan

import (
	"testing"

	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// Test for BrightnessSupported
func TestBrightnessSupported(t *testing.T) {
	cases := []struct {
		xmlValue      string
		expectedBool  bool
		expectedValid bool
	}{
		{"1", true, true},
		{"0", false, true},
		{"true", true, true},
		{"false", false, true},
		{" TRUE ", true, true},
		{" False ", false, true},
		{"maybe", false, false},
		{"", false, false},
	}

	for _, c := range cases {
		bs := BrightnessSupported(c.xmlValue)
		if bs.IsValid() != c.expectedValid {
			t.Errorf(
				"IsValid: input %q, expected %v, got %v",
				c.xmlValue, c.expectedValid, bs.IsValid(),
			)
		}
		b, err := bs.Bool()
		if c.expectedValid {
			if err != nil {
				t.Errorf(
					"Bool: input %q, unexpected error: %v",
					c.xmlValue, err,
				)
			}
			if b != c.expectedBool {
				t.Errorf(
					"Bool: input %q, expected %v, got %v",
					c.xmlValue, c.expectedBool, b,
				)
			}
		} else {
			if err == nil {
				t.Errorf(
					"Bool: input %q, expected error, got nil",
					c.xmlValue,
				)
			}
		}
	}

	// Test XML round-trip for valid values
	validCases := []string{
		"1", "0", "true", "false", " TRUE ", " False ",
	}
	for _, val := range validCases {
		bs := BrightnessSupported(val)
		elm := bs.toXML("wscn:BrightnessSupported")
		bs2, err := decodeBrightnessSupported(elm)
		if err != nil {
			t.Errorf(
				"decodeBrightnessSupported: input %q, unexpected error: %v",
				val, err,
			)
		}
		if !bs2.IsValid() {
			t.Errorf(
				"decodeBrightnessSupported: input %q, result not valid",
				val,
			)
		}
	}

	// Test decodeBrightnessSupported with invalid value
	elm := xmldoc.Element{
		Name: "wscn:BrightnessSupported", Text: "maybe",
	}
	_, err := decodeBrightnessSupported(elm)
	if err == nil {
		t.Errorf(
			"decodeBrightnessSupported: " +
				"expected error for invalid value, got nil",
		)
	}
}
