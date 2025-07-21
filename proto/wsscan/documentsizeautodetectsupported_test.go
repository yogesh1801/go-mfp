// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// Test for document size auto detect supported

package wsscan

import (
	"testing"

	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// Test for DocumentSizeAutoDetectSupported
func TestDocumentSizeAutoDetectSupported(t *testing.T) {
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
		dsads := DocumentSizeAutoDetectSupported(c.xmlValue)
		if dsads.IsValid() != c.expectedValid {
			t.Errorf(
				"IsValid: input %q, expected %v, got %v",
				c.xmlValue, c.expectedValid, dsads.IsValid(),
			)
		}
		b, err := dsads.Bool()
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
		dsads := DocumentSizeAutoDetectSupported(val)
		elm := dsads.toXML("wscn:DocumentSizeAutoDetectSupported")
		dsads2, err := decodeDocumentSizeAutoDetectSupported(elm)
		if err != nil {
			t.Errorf(
				"decodeDocumentSizeAutoDetectSupported: input %q, unexpected error: %v",
				val, err,
			)
		}
		if !dsads2.IsValid() {
			t.Errorf(
				"decodeDocumentSizeAutoDetectSupported: input %q, result not valid",
				val,
			)
		}
	}

	// Test decodeDocumentSizeAutoDetectSupported with invalid value
	elm := xmldoc.Element{
		Name: "wscn:DocumentSizeAutoDetectSupported", Text: "maybe",
	}
	_, err := decodeDocumentSizeAutoDetectSupported(elm)
	if err == nil {
		t.Errorf(
			"decodeDocumentSizeAutoDetectSupported: " +
				"expected error for invalid value, got nil",
		)
	}
}
