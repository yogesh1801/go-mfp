// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// Test for min value

package wsscan

import (
	"testing"
)

func TestMinValue_StringAndDecode(t *testing.T) {
	cases := []struct {
		val MinValue
		str string
	}{
		{-100, "-100"},
		{0, "0"},
		{42, "42"},
	}
	for _, c := range cases {
		if got := c.val.String(); got != c.str {
			t.Errorf("String: expected %q, got %q", c.str, got)
		}
		parsed, err := decodeMinValueStr(c.str)
		if err != nil {
			t.Errorf(
				"decodeMinValueStr: input %q, unexpected error: %v", c.str, err)
		}
		if parsed != c.val {
			t.Errorf(
				"decodeMinValueStr: input %q, expected %v, got %v",
				c.str, c.val, parsed)
		}
	}
	// Error case
	if _, err := decodeMinValueStr("abc"); err == nil {
		t.Errorf("decodeMinValueStr: expected error for invalid input, got nil")
	}
}

func TestMinValue_XMLRoundTrip(t *testing.T) {
	values := []MinValue{-100, 0, 42}
	for _, val := range values {
		elm := val.toXML("wscn:MinValue")
		parsed, err := decodeMinValue(elm)
		if err != nil {
			t.Errorf(
				"decodeMinValue: input %q, unexpected error: %v",
				val.String(), err)
		}
		if parsed != val {
			t.Errorf(
				"XML round-trip: expected %v, got %v",
				val, parsed)
		}
	}
}
