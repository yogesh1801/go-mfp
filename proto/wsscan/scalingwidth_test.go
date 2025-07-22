// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions

package wsscan

import (
	"reflect"
	"testing"

	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// Test for ScalingWidth
func TestScalingWidth_XMLRoundTrip(t *testing.T) {
	cases := []struct {
		min, max int
	}{
		{100, 200},
		{1, 1000},
		{100, 100}, // min == max
	}
	for _, c := range cases {
		sw := ScalingWidth{RangeElement: RangeElement{MinValue: c.min, MaxValue: c.max}}
		elm := sw.toXML(NsWSCN + ":ScalingWidth")
		parsed, err := decodeScalingWidth(elm)
		if err != nil {
			t.Errorf("decodeScalingWidth: input %+v, unexpected error: %v",
				sw, err)
		}
		if !reflect.DeepEqual(parsed, sw) {
			t.Errorf("XML round-trip: expected %+v, got %+v", sw, parsed)
		}
	}
}

func TestScalingWidth_Validation(t *testing.T) {
	badCases := []ScalingWidth{
		{RangeElement: RangeElement{MinValue: 0, MaxValue: 100}},     // MinValue < 1
		{RangeElement: RangeElement{MinValue: 1, MaxValue: 0}},       // MaxValue < 1
		{RangeElement: RangeElement{MinValue: 1001, MaxValue: 1002}}, // MinValue > 1000
		{RangeElement: RangeElement{MinValue: 100, MaxValue: 99}},    // MinValue > MaxValue
	}
	for _, sw := range badCases {
		if err := sw.Validate(); err == nil {
			t.Errorf("Validate: expected error for %+v, got nil", sw)
		}
	}
}

func TestScalingWidth_DecodeErrors(t *testing.T) {
	// Missing MinValue
	elm1 := xmldoc.Element{
		Name: NsWSCN + ":ScalingWidth",
		Children: []xmldoc.Element{
			{Name: NsWSCN + ":MaxValue", Text: "100"},
		},
	}
	if _, err := decodeScalingWidth(elm1); err == nil {
		t.Error("decodeScalingWidth: " +
			"expected error for missing MinValue, got nil")
	}
	// Missing MaxValue
	elm2 := xmldoc.Element{
		Name: NsWSCN + ":ScalingWidth",
		Children: []xmldoc.Element{
			{Name: NsWSCN + ":MinValue", Text: "100"},
		},
	}
	if _, err := decodeScalingWidth(elm2); err == nil {
		t.Error("decodeScalingWidth: " +
			"expected error for missing MaxValue, got nil")
	}
	// Invalid MinValue
	elm3 := xmldoc.Element{
		Name: NsWSCN + ":ScalingWidth",
		Children: []xmldoc.Element{
			{Name: NsWSCN + ":MinValue", Text: "abc"},
			{Name: NsWSCN + ":MaxValue", Text: "100"},
		},
	}
	if _, err := decodeScalingWidth(elm3); err == nil {
		t.Error("decodeScalingWidth: " +
			"expected error for invalid MinValue, got nil")
	}
	// Invalid MaxValue
	elm4 := xmldoc.Element{
		Name: NsWSCN + ":ScalingWidth",
		Children: []xmldoc.Element{
			{Name: NsWSCN + ":MinValue", Text: "100"},
			{Name: NsWSCN + ":MaxValue", Text: "abc"},
		},
	}
	if _, err := decodeScalingWidth(elm4); err == nil {
		t.Error("decodeScalingWidth: " +
			"expected error for invalid MaxValue, got nil")
	}
}
