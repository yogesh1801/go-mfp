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

// Test for PlatenColor
func TestPlatenColor_XMLRoundTrip(t *testing.T) {
	cases := [][]ColorEntry{
		{BlackAndWhite1},
		{BlackAndWhite1, Grayscale8},
		{BlackAndWhite1, Grayscale8, RGB24},
	}
	for _, vals := range cases {
		pc := PlatenColor{Values: vals}
		elm := pc.toXML(NsWSCN + ":PlatenColor")
		parsed, err := decodePlatenColor(elm)
		if err != nil {
			t.Errorf("decodePlatenColor: input %+v, unexpected error: %v",
				pc, err)
		}
		if !reflect.DeepEqual(parsed, pc) {
			t.Errorf("XML round-trip: expected %+v, got %+v", pc, parsed)
		}
	}
}

func TestPlatenColor_DecodeErrors(t *testing.T) {
	// Missing ColorEntry
	elm1 := xmldoc.Element{
		Name: NsWSCN + ":PlatenColor",
	}
	if _, err := decodePlatenColor(elm1); err == nil {
		t.Error("decodePlatenColor: " +
			"expected error for missing ColorEntry, got nil")
	}
	// Invalid ColorEntry
	elm2 := xmldoc.Element{
		Name: NsWSCN + ":PlatenColor",
		Children: []xmldoc.Element{
			{Name: NsWSCN + ":ColorEntry", Text: "bad"},
		},
	}
	if _, err := decodePlatenColor(elm2); err == nil {
		t.Error("decodePlatenColor: " +
			"expected error for invalid ColorEntry, got nil")
	}
}
