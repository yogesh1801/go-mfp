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

// Test for ScalingHeight
func TestScalingHeight_XMLRoundTrip(t *testing.T) {
	cases := []struct {
		min, max int
	}{
		{100, 200},
		{1, 1000},
		{100, 100}, // min == max
	}
	for _, c := range cases {
		sh := ScalingHeight{RangeElement: RangeElement{MinValue: c.min, MaxValue: c.max}}
		elm := sh.toXML(NsWSCN + ":ScalingHeight")
		parsed, err := decodeScalingHeight(elm)
		if err != nil {
			t.Errorf("decodeScalingHeight: input %+v, unexpected error: %v",
				sh, err)
		}
		if !reflect.DeepEqual(parsed, sh) {
			t.Errorf("XML round-trip: expected %+v, got %+v", sh, parsed)
		}
	}
}

func TestScalingHeight_Validation(t *testing.T) {
	badCases := []ScalingHeight{
		{RangeElement: RangeElement{MinValue: 0, MaxValue: 100}},     // MinValue < 1
		{RangeElement: RangeElement{MinValue: 1, MaxValue: 0}},       // MaxValue < 1
		{RangeElement: RangeElement{MinValue: 1001, MaxValue: 1002}}, // MinValue > 1000
		{RangeElement: RangeElement{MinValue: 100, MaxValue: 99}},    // MinValue > MaxValue
	}
	for _, sh := range badCases {
		if err := sh.Validate(); err == nil {
			t.Errorf("Validate: expected error for %+v, got nil", sh)
		}
	}
}

func TestScalingHeight_DecodeErrors(t *testing.T) {
	// Missing MinValue
	elm1 := xmldoc.Element{
		Name: NsWSCN + ":ScalingHeight",
		Children: []xmldoc.Element{
			{Name: NsWSCN + ":MaxValue", Text: "100"},
		},
	}
	if _, err := decodeScalingHeight(elm1); err == nil {
		t.Error("decodeScalingHeight: " +
			"expected error for missing MinValue, got nil")
	}
	// Missing MaxValue
	elm2 := xmldoc.Element{
		Name: NsWSCN + ":ScalingHeight",
		Children: []xmldoc.Element{
			{Name: NsWSCN + ":MinValue", Text: "100"},
		},
	}
	if _, err := decodeScalingHeight(elm2); err == nil {
		t.Error("decodeScalingHeight: " +
			"expected error for missing MaxValue, got nil")
	}
	// Invalid MinValue
	elm3 := xmldoc.Element{
		Name: NsWSCN + ":ScalingHeight",
		Children: []xmldoc.Element{
			{Name: NsWSCN + ":MinValue", Text: "abc"},
			{Name: NsWSCN + ":MaxValue", Text: "100"},
		},
	}
	if _, err := decodeScalingHeight(elm3); err == nil {
		t.Error("decodeScalingHeight: " +
			"expected error for invalid MinValue, got nil")
	}
	// Invalid MaxValue
	elm4 := xmldoc.Element{
		Name: NsWSCN + ":ScalingHeight",
		Children: []xmldoc.Element{
			{Name: NsWSCN + ":MinValue", Text: "100"},
			{Name: NsWSCN + ":MaxValue", Text: "abc"},
		},
	}
	if _, err := decodeScalingHeight(elm4); err == nil {
		t.Error("decodeScalingHeight: " +
			"expected error for invalid MaxValue, got nil")
	}
}
