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

// Test for FormatsSupported
func TestFormatsSupported_XMLRoundTrip(t *testing.T) {
	cases := [][]FormatValue{
		{DIB},
		{DIB, PNG},
		{DIB, PNG, XPS},
	}
	for _, vals := range cases {
		fs := FormatsSupported{Values: vals}
		elm := fs.toXML(NsWSCN + ":FormatsSupported")
		parsed, err := decodeFormatsSupported(elm)
		if err != nil {
			t.Errorf("decodeFormatsSupported: input %+v, unexpected error: %v",
				fs, err)
		}
		if !reflect.DeepEqual(parsed, fs) {
			t.Errorf("XML round-trip: expected %+v, got %+v", fs, parsed)
		}
	}
}

func TestFormatsSupported_DecodeErrors(t *testing.T) {
	// Missing FormatValue
	elm1 := xmldoc.Element{
		Name: NsWSCN + ":FormatsSupported",
	}
	if _, err := decodeFormatsSupported(elm1); err == nil {
		t.Error("decodeFormatsSupported: " +
			"expected error for missing FormatValue, got nil")
	}
	// Invalid FormatValue
	elm2 := xmldoc.Element{
		Name: NsWSCN + ":FormatsSupported",
		Children: []xmldoc.Element{
			{Name: NsWSCN + ":FormatValue", Text: "bad"},
		},
	}
	if _, err := decodeFormatsSupported(elm2); err == nil {
		t.Error("decodeFormatsSupported: " +
			"expected error for invalid FormatValue, got nil")
	}
}
