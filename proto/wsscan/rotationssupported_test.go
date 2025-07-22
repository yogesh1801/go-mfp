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

// Test for RotationsSupported
func TestRotationsSupported_XMLRoundTrip(t *testing.T) {
	cases := [][]RotationValue{
		{Rotation0},
		{Rotation0, Rotation90},
		{Rotation0, Rotation90, Rotation180, Rotation270},
	}
	for _, vals := range cases {
		rs := RotationsSupported{Values: vals}
		elm := rs.toXML(NsWSCN + ":RotationsSupported")
		parsed, err := decodeRotationsSupported(elm)
		if err != nil {
			t.Errorf("decodeRotationsSupported: input %+v, unexpected error: %v",
				rs, err)
		}
		if !reflect.DeepEqual(parsed, rs) {
			t.Errorf("XML round-trip: expected %+v, got %+v", rs, parsed)
		}
	}
}

func TestRotationsSupported_Validation(t *testing.T) {
	badCases := []RotationsSupported{
		{}, // empty
	}
	for _, rs := range badCases {
		// No Validate method; test decodeRotationsSupported error instead
		elm := RotationsSupported{Values: rs.Values}.toXML(NsWSCN +
			":RotationsSupported")
		_, err := decodeRotationsSupported(elm)
		if err == nil {
			t.Errorf("decodeRotationsSupported: "+
				"expected error for %+v, got nil", rs)
		}
	}
}

func TestRotationsSupported_DecodeErrors(t *testing.T) {
	// Missing RotationValue
	elm1 := xmldoc.Element{
		Name: NsWSCN + ":RotationsSupported",
	}
	if _, err := decodeRotationsSupported(elm1); err == nil {
		t.Error("decodeRotationsSupported: " +
			"expected error for missing RotationValue, got nil")
	}
	// Invalid RotationValue
	elm2 := xmldoc.Element{
		Name: NsWSCN + ":RotationsSupported",
		Children: []xmldoc.Element{
			{Name: NsWSCN + ":RotationValue", Text: "bad"},
		},
	}
	if _, err := decodeRotationsSupported(elm2); err == nil {
		t.Error("decodeRotationsSupported: " +
			"expected error for invalid RotationValue, got nil")
	}
}
