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

// Test for ScalingRangeSupported
func TestScalingRangeSupported_XMLRoundTrip(t *testing.T) {
	cases := []struct {
		wmin, wmax, hmin, hmax int
	}{
		{100, 200, 150, 250},
		{1, 1000, 1, 1000},
		{100, 100, 100, 100},
	}
	for _, c := range cases {
		srs := ScalingRangeSupported{
			ScalingWidth:  ScalingWidth{RangeElement: RangeElement{MinValue: c.wmin, MaxValue: c.wmax}},
			ScalingHeight: ScalingHeight{RangeElement: RangeElement{MinValue: c.hmin, MaxValue: c.hmax}},
		}
		elm := srs.toXML(NsWSCN + ":ScalingRangeSupported")
		parsed, err := decodeScalingRangeSupported(elm)
		if err != nil {
			t.Errorf("decodeScalingRangeSupported: input %+v, unexpected error: %v",
				srs, err)
		}
		if !reflect.DeepEqual(parsed, srs) {
			t.Errorf("XML round-trip: expected %+v, got %+v", srs, parsed)
		}
	}
}

func TestScalingRangeSupported_Validation(t *testing.T) {
	badCases := []ScalingRangeSupported{
		{
			ScalingWidth:  ScalingWidth{RangeElement: RangeElement{MinValue: 0, MaxValue: 100}},
			ScalingHeight: ScalingHeight{RangeElement: RangeElement{MinValue: 100, MaxValue: 200}},
		},
		{
			ScalingWidth:  ScalingWidth{RangeElement: RangeElement{MinValue: 100, MaxValue: 200}},
			ScalingHeight: ScalingHeight{RangeElement: RangeElement{MinValue: 0, MaxValue: 100}},
		},
		{
			ScalingWidth:  ScalingWidth{RangeElement: RangeElement{MinValue: 100, MaxValue: 99}},
			ScalingHeight: ScalingHeight{RangeElement: RangeElement{MinValue: 100, MaxValue: 200}},
		},
		{
			ScalingWidth:  ScalingWidth{RangeElement: RangeElement{MinValue: 100, MaxValue: 200}},
			ScalingHeight: ScalingHeight{RangeElement: RangeElement{MinValue: 100, MaxValue: 99}},
		},
	}
	for _, srs := range badCases {
		if err := srs.Validate(); err == nil {
			t.Errorf("Validate: expected error for %+v, got nil", srs)
		}
	}
}

func TestScalingRangeSupported_DecodeErrors(t *testing.T) {
	// Missing ScalingWidth
	elm1 := xmldoc.Element{
		Name: NsWSCN + ":ScalingRangeSupported",
		Children: []xmldoc.Element{
			{Name: NsWSCN + ":ScalingHeight", Children: []xmldoc.Element{
				{Name: NsWSCN + ":MinValue", Text: "100"},
				{Name: NsWSCN + ":MaxValue", Text: "200"},
			}},
		},
	}
	if _, err := decodeScalingRangeSupported(elm1); err == nil {
		t.Error("decodeScalingRangeSupported: " +
			"expected error for missing ScalingWidth, got nil")
	}
	// Missing ScalingHeight
	elm2 := xmldoc.Element{
		Name: NsWSCN + ":ScalingRangeSupported",
		Children: []xmldoc.Element{
			{Name: NsWSCN + ":ScalingWidth", Children: []xmldoc.Element{
				{Name: NsWSCN + ":MinValue", Text: "100"},
				{Name: NsWSCN + ":MaxValue", Text: "200"},
			}},
		},
	}
	if _, err := decodeScalingRangeSupported(elm2); err == nil {
		t.Error("decodeScalingRangeSupported: " +
			"expected error for missing ScalingHeight, got nil")
	}
	// Invalid ScalingWidth
	elm3 := xmldoc.Element{
		Name: NsWSCN + ":ScalingRangeSupported",
		Children: []xmldoc.Element{
			{Name: NsWSCN + ":ScalingWidth", Children: []xmldoc.Element{
				{Name: NsWSCN + ":MinValue", Text: "abc"},
				{Name: NsWSCN + ":MaxValue", Text: "200"},
			}},
			{Name: NsWSCN + ":ScalingHeight", Children: []xmldoc.Element{
				{Name: NsWSCN + ":MinValue", Text: "100"},
				{Name: NsWSCN + ":MaxValue", Text: "200"},
			}},
		},
	}
	if _, err := decodeScalingRangeSupported(elm3); err == nil {
		t.Error("decodeScalingRangeSupported: " +
			"expected error for invalid ScalingWidth, got nil")
	}
	// Invalid ScalingHeight
	elm4 := xmldoc.Element{
		Name: NsWSCN + ":ScalingRangeSupported",
		Children: []xmldoc.Element{
			{Name: NsWSCN + ":ScalingWidth", Children: []xmldoc.Element{
				{Name: NsWSCN + ":MinValue", Text: "100"},
				{Name: NsWSCN + ":MaxValue", Text: "200"},
			}},
			{Name: NsWSCN + ":ScalingHeight", Children: []xmldoc.Element{
				{Name: NsWSCN + ":MinValue", Text: "abc"},
				{Name: NsWSCN + ":MaxValue", Text: "200"},
			}},
		},
	}
	if _, err := decodeScalingRangeSupported(elm4); err == nil {
		t.Error("decodeScalingRangeSupported: " +
			"expected error for invalid ScalingHeight, got nil")
	}
}
