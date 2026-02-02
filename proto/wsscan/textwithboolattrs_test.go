// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// Test for TextWithBoolAttrs

package wsscan

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"

	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// Helper functions for string type
func stringDecoder(s string) (string, error) {
	return s, nil
}

func stringEncoder(s string) string {
	return s
}

// Helper functions for int type
func intDecoder(s string) (int, error) {
	return strconv.Atoi(s)
}

func intEncoder(i int) string {
	return strconv.Itoa(i)
}

func TestTextWithBoolAttrs_String_RoundTrip_AllAttributes(t *testing.T) {
	orig := TextWithBoolAttrs[string]{
		Text:        "100",
		MustHonor:   optional.New(BooleanElement("true")),
		Override:    optional.New(BooleanElement("false")),
		UsedDefault: optional.New(BooleanElement("1")),
	}
	elm := orig.toXML("wscn:CompressionQualityFactor", stringEncoder)
	if elm.Name != "wscn:CompressionQualityFactor" {
		t.Errorf("expected element name 'wscn:CompressionQualityFactor', got '%s'", elm.Name)
	}
	if elm.Text != orig.Text {
		t.Errorf("expected text '%s', got '%s'", orig.Text, elm.Text)
	}
	if len(elm.Attrs) != 3 {
		t.Errorf("expected 3 attributes, got %d: %+v", len(elm.Attrs), elm.Attrs)
	}

	var decoded TextWithBoolAttrs[string]
	decoded, err := decoded.decodeTextWithBoolAttrs(elm, stringDecoder)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}
}

func TestTextWithBoolAttrs_String_RoundTrip_NoAttributes(t *testing.T) {
	orig := TextWithBoolAttrs[string]{
		Text: "50",
	}
	elm := orig.toXML("wscn:CompressionQualityFactor", stringEncoder)
	if len(elm.Attrs) != 0 {
		t.Errorf("expected no attributes, got %+v", elm.Attrs)
	}

	var decoded TextWithBoolAttrs[string]
	decoded, err := decoded.decodeTextWithBoolAttrs(elm, stringDecoder)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}
}

func TestTextWithBoolAttrs_String_RoundTrip_PartialAttributes(t *testing.T) {
	orig := TextWithBoolAttrs[string]{
		Text:      "75",
		MustHonor: optional.New(BooleanElement("true")),
		Override:  optional.New(BooleanElement("0")),
	}
	elm := orig.toXML("wscn:CompressionQualityFactor", stringEncoder)
	if len(elm.Attrs) != 2 {
		t.Errorf("expected 2 attributes, got %d: %+v", len(elm.Attrs), elm.Attrs)
	}

	var decoded TextWithBoolAttrs[string]
	decoded, err := decoded.decodeTextWithBoolAttrs(elm, stringDecoder)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}
}

func TestTextWithBoolAttrs_Int_RoundTrip(t *testing.T) {
	orig := TextWithBoolAttrs[int]{
		Text:        100,
		MustHonor:   optional.New(BooleanElement("true")),
		UsedDefault: optional.New(BooleanElement("false")),
	}
	elm := orig.toXML("wscn:CompressionQualityFactor", intEncoder)
	if elm.Name != "wscn:CompressionQualityFactor" {
		t.Errorf("expected element name 'wscn:CompressionQualityFactor', got '%s'", elm.Name)
	}
	if elm.Text != "100" {
		t.Errorf("expected text '100', got '%s'", elm.Text)
	}

	var decoded TextWithBoolAttrs[int]
	decoded, err := decoded.decodeTextWithBoolAttrs(elm, intDecoder)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}
}

func TestTextWithBoolAttrs_InvalidBooleanAttribute(t *testing.T) {
	var elem TextWithBoolAttrs[string]
	root := elem.toXML("wscn:Test", stringEncoder)
	root.Attrs = []xmldoc.Attr{
		{Name: NsWSCN + ":MustHonor", Value: "invalid"},
	}

	_, err := elem.decodeTextWithBoolAttrs(root, stringDecoder)
	if err == nil {
		t.Error("expected error for invalid boolean attribute, got nil")
	}
}

func TestTextWithBoolAttrs_BooleanVariations(t *testing.T) {
	cases := []struct {
		name  string
		value BooleanElement
	}{
		{"true", "true"},
		{"false", "false"},
		{"1", "1"},
		{"0", "0"},
		{"TRUE", "TRUE"},
		{"FALSE", "FALSE"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			orig := TextWithBoolAttrs[string]{
				Text:      "test",
				MustHonor: optional.New(c.value),
			}
			elm := orig.toXML("wscn:Test", stringEncoder)
			var decoded TextWithBoolAttrs[string]
			decoded, err := decoded.decodeTextWithBoolAttrs(elm, stringDecoder)
			if err != nil {
				t.Fatalf("decode returned error: %v", err)
			}
			if !reflect.DeepEqual(orig, decoded) {
				t.Errorf("expected %+v, got %+v", orig, decoded)
			}
		})
	}
}

func TestTextWithBoolAttrs_Int_InvalidValue(t *testing.T) {
	var elem TextWithBoolAttrs[int]
	root := elem.toXML("wscn:Test", intEncoder)
	root.Text = "not-a-number"

	_, err := elem.decodeTextWithBoolAttrs(root, intDecoder)
	if err == nil {
		t.Error("expected error for invalid int value, got nil")
	}
}

// Example usage demonstrating the generic type
func ExampleTextWithBoolAttrs_string() {
	elem := TextWithBoolAttrs[string]{
		Text:      "high",
		MustHonor: optional.New(BooleanElement("true")),
	}
	xml := elem.toXML("wscn:Quality", stringEncoder)
	fmt.Printf("Name: %s, Text: %s\n", xml.Name, xml.Text)
	// Output: Name: wscn:Quality, Text: high
}

func ExampleTextWithBoolAttrs_int() {
	elem := TextWithBoolAttrs[int]{
		Text:     85,
		Override: optional.New(BooleanElement("false")),
	}
	xml := elem.toXML("wscn:CompressionQualityFactor", intEncoder)
	fmt.Printf("Name: %s, Text: %s\n", xml.Name, xml.Text)
	// Output: Name: wscn:CompressionQualityFactor, Text: 85
}
