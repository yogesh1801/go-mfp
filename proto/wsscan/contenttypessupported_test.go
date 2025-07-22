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

// Test for ContentTypesSupported
func TestContentTypesSupported_XMLRoundTrip(t *testing.T) {
	cases := [][]ContentTypeValue{
		{Auto},
		{Auto, Text},
		{Auto, Text, Photo, Halftone, Mixed},
	}
	for _, vals := range cases {
		cts := ContentTypesSupported{Values: vals}
		elm := cts.toXML(NsWSCN + ":ContentTypesSupported")
		parsed, err := decodeContentTypesSupported(elm)
		if err != nil {
			t.Errorf("decodeContentTypesSupported: input %+v, unexpected error: %v",
				cts, err)
		}
		if !reflect.DeepEqual(parsed, cts) {
			t.Errorf("XML round-trip: expected %+v, got %+v", cts, parsed)
		}
	}
}

func TestContentTypesSupported_DecodeErrors(t *testing.T) {
	// Missing ContentTypeValue
	elm1 := xmldoc.Element{
		Name: NsWSCN + ":ContentTypesSupported",
	}
	if _, err := decodeContentTypesSupported(elm1); err == nil {
		t.Error("decodeContentTypesSupported: " +
			"expected error for missing ContentTypeValue, got nil")
	}
	// Invalid ContentTypeValue
	elm2 := xmldoc.Element{
		Name: NsWSCN + ":ContentTypesSupported",
		Children: []xmldoc.Element{
			{Name: NsWSCN + ":ContentTypeValue", Text: "bad"},
		},
	}
	if _, err := decodeContentTypesSupported(elm2); err == nil {
		t.Error("decodeContentTypesSupported: " +
			"expected error for invalid ContentTypeValue, got nil")
	}
}
