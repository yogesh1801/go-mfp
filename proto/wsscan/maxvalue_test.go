// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// Test for max value

package wsscan

import (
	"testing"
)

// Test for MaxValue
func TestMaxValue_StringAndDecode(t *testing.T) {
	cases := []struct {
		val MaxValue
		str string
	}{
		{100, "100"},
		{0, "0"},
	}
	for _, c := range cases {
		if got := c.val.String(); got != c.str {
			t.Errorf("String: expected %q, got %q", c.str, got)
		}
		parsed, err := decodeMaxValueStr(c.str)
		if err != nil {
			t.Errorf(
				"DecodeMaxValue: input %q, unexpected error: %v", c.str, err)
		}
		if parsed != c.val {
			t.Errorf(
				"DecodeMaxValue: input %q, expected %v, got %v",
				c.str, c.val, parsed)
		}
	}
	// Error case
	if _, err := decodeMaxValueStr("abc"); err == nil {
		t.Errorf("DecodeMaxValue: expected error for invalid input, got nil")
	}
}

func TestMaxValue_XMLRoundTrip(t *testing.T) {
	values := []MaxValue{100, 0, -42}
	for _, val := range values {
		elm := val.toXML("wscn:MaxValue")
		parsed, err := decodeMaxValue(elm)
		if err != nil {
			t.Errorf(
				"decodeMaxValue: input %q, unexpected error: %v",
				val.String(), err)
		}
		if parsed != val {
			t.Errorf("XML round-trip: expected %v, got %v", val, parsed)
		}
	}
}
