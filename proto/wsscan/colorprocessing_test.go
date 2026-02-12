// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// Test for ColorProcessing

package wsscan

import (
	"reflect"
	"testing"

	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

func TestColorProcessing_RoundTrip_AllAttributes(t *testing.T) {
	orig := ColorProcessing(
		ValWithOptions[ColorEntry]{
			Text:        RGB24,
			MustHonor:   optional.New(BooleanElement("true")),
			Override:    optional.New(BooleanElement("false")),
			UsedDefault: optional.New(BooleanElement("1")),
		},
	)

	elm := orig.toXML("wscn:ColorProcessing")
	if elm.Name != "wscn:ColorProcessing" {
		t.Errorf("expected element name 'wscn:ColorProcessing', got '%s'",
			elm.Name)
	}
	if elm.Text != "RGB24" {
		t.Errorf("expected text 'RGB24', got '%s'", elm.Text)
	}
	if len(elm.Attrs) != 3 {
		t.Errorf("expected 3 attributes, got %d: %+v", len(elm.Attrs),
			elm.Attrs)
	}

	decoded, err := decodeColorProcessing(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}
}

func TestColorProcessing_RoundTrip_NoAttributes(t *testing.T) {
	orig := ColorProcessing(
		ValWithOptions[ColorEntry]{
			Text: Grayscale8,
		},
	)

	elm := orig.toXML("wscn:ColorProcessing")
	if len(elm.Attrs) != 0 {
		t.Errorf("expected no attributes, got %+v", elm.Attrs)
	}
	if elm.Text != "Grayscale8" {
		t.Errorf("expected text 'Grayscale8', got '%s'", elm.Text)
	}

	decoded, err := decodeColorProcessing(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}
}

func TestColorProcessing_RoundTrip_PartialAttributes(t *testing.T) {
	orig := ColorProcessing(
		ValWithOptions[ColorEntry]{
			Text:      BlackAndWhite1,
			MustHonor: optional.New(BooleanElement("true")),
			Override:  optional.New(BooleanElement("0")),
		},
	)

	elm := orig.toXML("wscn:ColorProcessing")
	if len(elm.Attrs) != 2 {
		t.Errorf("expected 2 attributes, got %d: %+v", len(elm.Attrs),
			elm.Attrs)
	}
	if elm.Text != "BlackAndWhite1" {
		t.Errorf("expected text 'BlackAndWhite1', got '%s'", elm.Text)
	}

	decoded, err := decodeColorProcessing(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}
}

func TestColorProcessing_AllColorEntries(t *testing.T) {
	cases := []struct {
		name  string
		value ColorEntry
		str   string
	}{
		{"BlackAndWhite1", BlackAndWhite1, "BlackAndWhite1"},
		{"Grayscale4", Grayscale4, "Grayscale4"},
		{"Grayscale8", Grayscale8, "Grayscale8"},
		{"Grayscale16", Grayscale16, "Grayscale16"},
		{"RGB24", RGB24, "RGB24"},
		{"RGB48", RGB48, "RGB48"},
		{"RGBA32", RGBA32, "RGBa32"},
		{"RGBA64", RGBA64, "RGBa64"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			orig := ColorProcessing(
				ValWithOptions[ColorEntry]{
					Text: c.value,
				},
			)

			elm := orig.toXML("wscn:ColorProcessing")
			if elm.Text != c.str {
				t.Errorf("expected text '%s', got '%s'", c.str, elm.Text)
			}

			decoded, err := decodeColorProcessing(elm)
			if err != nil {
				t.Fatalf("decode returned error: %v", err)
			}
			if decoded.Text != c.value {
				t.Errorf("expected value %v, got %v", c.value, decoded.Text)
			}
		})
	}
}

func TestColorProcessing_InvalidValue(t *testing.T) {
	elm := xmldoc.Element{
		Name: "wscn:ColorProcessing",
		Text: "InvalidColorMode",
	}

	_, err := decodeColorProcessing(elm)
	if err == nil {
		t.Error("expected error for invalid ColorProcessing value, got nil")
	}
}

func TestColorProcessing_InvalidBooleanAttribute(t *testing.T) {
	elm := xmldoc.Element{
		Name: "wscn:ColorProcessing",
		Text: "RGB24",
		Attrs: []xmldoc.Attr{
			{Name: NsWSCN + ":MustHonor", Value: "invalid"},
		},
	}

	_, err := decodeColorProcessing(elm)
	if err == nil {
		t.Error("expected error for invalid boolean attribute, got nil")
	}
}

func TestColorProcessing_BooleanVariations(t *testing.T) {
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
			orig := ColorProcessing(
				ValWithOptions[ColorEntry]{
					Text:      Grayscale16,
					MustHonor: optional.New(c.value),
				},
			)

			elm := orig.toXML("wscn:ColorProcessing")
			decoded, err := decodeColorProcessing(elm)
			if err != nil {
				t.Fatalf("decode returned error: %v", err)
			}
			if !reflect.DeepEqual(orig, decoded) {
				t.Errorf("expected %+v, got %+v", orig, decoded)
			}
		})
	}
}
