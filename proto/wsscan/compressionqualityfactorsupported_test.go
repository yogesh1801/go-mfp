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

// Test for CompressionQualityFactorSupported
func TestCompressionQualityFactorSupported_XMLRoundTrip(t *testing.T) {
	cases := []struct {
		min, max int
	}{
		{1, 100},
		{50, 75},
		{100, 100},
	}
	for _, c := range cases {
		cqfs := CompressionQualityFactorSupported{MinValue: MinValue(c.min), MaxValue: MaxValue(c.max)}
		elm := cqfs.toXML(NsWSCN + ":CompressionQualityFactorSupported")
		parsed, err := decodeCompressionQualityFactorSupported(elm)
		if err != nil {
			t.Errorf("decodeCompressionQualityFactorSupported: "+
				"input %+v, unexpected error: %v",
				cqfs, err)
		}
		if !reflect.DeepEqual(parsed, cqfs) {
			t.Errorf("XML round-trip: expected %+v, got %+v", cqfs, parsed)
		}
	}
}

func TestCompressionQualityFactorSupported_Validation(t *testing.T) {
	badCases := []CompressionQualityFactorSupported{
		{MinValue: 0, MaxValue: 100},   // MinValue < 1
		{MinValue: 1, MaxValue: 0},     // MaxValue < 1
		{MinValue: 101, MaxValue: 101}, // MinValue > 100
		{MinValue: 100, MaxValue: 101}, // MaxValue > 100
		{MinValue: 50, MaxValue: 49},   // MinValue > MaxValue
	}
	for _, cqfs := range badCases {
		if err := cqfs.Validate(); err == nil {
			t.Errorf("Validate: expected error for %+v, got nil", cqfs)
		}
	}
}

func TestCompressionQualityFactorSupported_DecodeErrors(t *testing.T) {
	// Missing MinValue
	elm1 := xmldoc.Element{
		Name: NsWSCN + ":CompressionQualityFactorSupported",
		Children: []xmldoc.Element{
			{Name: NsWSCN + ":MaxValue", Text: "100"},
		},
	}
	if _, err := decodeCompressionQualityFactorSupported(elm1); err == nil {
		t.Error("decodeCompressionQualityFactorSupported: " +
			"expected error for missing MinValue, got nil")
	}
	// Missing MaxValue
	elm2 := xmldoc.Element{
		Name: NsWSCN + ":CompressionQualityFactorSupported",
		Children: []xmldoc.Element{
			{Name: NsWSCN + ":MinValue", Text: "1"},
		},
	}
	if _, err := decodeCompressionQualityFactorSupported(elm2); err == nil {
		t.Error("decodeCompressionQualityFactorSupported: " +
			"expected error for missing MaxValue, got nil")
	}
	// Invalid MinValue
	elm3 := xmldoc.Element{
		Name: NsWSCN + ":CompressionQualityFactorSupported",
		Children: []xmldoc.Element{
			{Name: NsWSCN + ":MinValue", Text: "abc"},
			{Name: NsWSCN + ":MaxValue", Text: "100"},
		},
	}
	if _, err := decodeCompressionQualityFactorSupported(elm3); err == nil {
		t.Error("decodeCompressionQualityFactorSupported: " +
			"expected error for invalid MinValue, got nil")
	}
	// Invalid MaxValue
	elm4 := xmldoc.Element{
		Name: NsWSCN + ":CompressionQualityFactorSupported",
		Children: []xmldoc.Element{
			{Name: NsWSCN + ":MinValue", Text: "1"},
			{Name: NsWSCN + ":MaxValue", Text: "abc"},
		},
	}
	if _, err := decodeCompressionQualityFactorSupported(elm4); err == nil {
		t.Error("decodeCompressionQualityFactorSupported: " +
			"expected error for invalid MaxValue, got nil")
	}
}
