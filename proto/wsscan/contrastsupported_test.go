// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// Test for contrast supported

package wsscan

import (
	"testing"

	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// Test for ContrastSupported
func TestContrastSupported(t *testing.T) {
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
		cs := ContrastSupported(c.xmlValue)
		if cs.IsValid() != c.expectedValid {
			t.Errorf(
				"IsValid: input %q, expected %v, got %v",
				c.xmlValue, c.expectedValid, cs.IsValid(),
			)
		}
		b, err := cs.Bool()
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
		cs := ContrastSupported(val)
		elm := cs.toXML("wscn:ContrastSupported")
		cs2, err := decodeContrastSupported(elm)
		if err != nil {
			t.Errorf(
				"decodeContrastSupported: input %q, unexpected error: %v",
				val, err,
			)
		}
		if !cs2.IsValid() {
			t.Errorf(
				"decodeContrastSupported: input %q, result not valid",
				val,
			)
		}
	}

	// Test decodeContrastSupported with invalid value
	elm := xmldoc.Element{
		Name: "wscn:ContrastSupported", Text: "maybe",
	}
	_, err := decodeContrastSupported(elm)
	if err == nil {
		t.Errorf(
			"decodeContrastSupported: " +
				"expected error for invalid value, got nil",
		)
	}
}
