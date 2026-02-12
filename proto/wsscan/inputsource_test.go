// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// Test for InputSource

package wsscan

import (
	"reflect"
	"testing"

	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

func TestInputSource_RoundTrip_AllAttributes(t *testing.T) {
	orig := InputSource(
		ValWithOptions[InputSourceValue]{
			Text:        InputSourceADF,
			MustHonor:   optional.New(BooleanElement("true")),
			Override:    optional.New(BooleanElement("false")),
			UsedDefault: optional.New(BooleanElement("1")),
		},
	)

	elm := orig.toXML("wscn:InputSource")
	if elm.Name != "wscn:InputSource" {
		t.Errorf("expected element name 'wscn:InputSource', got '%s'",
			elm.Name)
	}
	if elm.Text != "ADF" {
		t.Errorf("expected text 'ADF', got '%s'", elm.Text)
	}
	if len(elm.Attrs) != 3 {
		t.Errorf("expected 3 attributes, got %d: %+v", len(elm.Attrs),
			elm.Attrs)
	}

	decoded, err := decodeInputSource(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}
}

func TestInputSource_RoundTrip_NoAttributes(t *testing.T) {
	orig := InputSource(
		ValWithOptions[InputSourceValue]{
			Text: InputSourcePlaten,
		},
	)

	elm := orig.toXML("wscn:InputSource")
	if len(elm.Attrs) != 0 {
		t.Errorf("expected no attributes, got %+v", elm.Attrs)
	}
	if elm.Text != "Platen" {
		t.Errorf("expected text 'Platen', got '%s'", elm.Text)
	}

	decoded, err := decodeInputSource(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}
}

func TestInputSource_RoundTrip_PartialAttributes(t *testing.T) {
	orig := InputSource(
		ValWithOptions[InputSourceValue]{
			Text:      InputSourceADFDuplex,
			MustHonor: optional.New(BooleanElement("true")),
			Override:  optional.New(BooleanElement("0")),
		},
	)

	elm := orig.toXML("wscn:InputSource")
	if len(elm.Attrs) != 2 {
		t.Errorf("expected 2 attributes, got %d: %+v", len(elm.Attrs),
			elm.Attrs)
	}
	if elm.Text != "ADFDuplex" {
		t.Errorf("expected text 'ADFDuplex', got '%s'", elm.Text)
	}

	decoded, err := decodeInputSource(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}
}

func TestInputSource_AllValues(t *testing.T) {
	cases := []struct {
		name  string
		value InputSourceValue
		str   string
	}{
		{"ADF", InputSourceADF, "ADF"},
		{"ADFDuplex", InputSourceADFDuplex, "ADFDuplex"},
		{"Film", InputSourceFilm, "Film"},
		{"Platen", InputSourcePlaten, "Platen"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			orig := InputSource(
				ValWithOptions[InputSourceValue]{
					Text: c.value,
				},
			)

			elm := orig.toXML("wscn:InputSource")
			if elm.Text != c.str {
				t.Errorf("expected text '%s', got '%s'", c.str, elm.Text)
			}

			decoded, err := decodeInputSource(elm)
			if err != nil {
				t.Fatalf("decode returned error: %v", err)
			}
			if decoded.Text != c.value {
				t.Errorf("expected value %v, got %v", c.value, decoded.Text)
			}
		})
	}
}

func TestInputSource_InvalidValue(t *testing.T) {
	elm := xmldoc.Element{
		Name: "wscn:InputSource",
		Text: "InvalidSource",
	}

	_, err := decodeInputSource(elm)
	if err == nil {
		t.Error("expected error for invalid InputSource value, got nil")
	}
}

func TestInputSource_InvalidBooleanAttribute(t *testing.T) {
	elm := xmldoc.Element{
		Name: "wscn:InputSource",
		Text: "ADF",
		Attrs: []xmldoc.Attr{
			{Name: NsWSCN + ":MustHonor", Value: "invalid"},
		},
	}

	_, err := decodeInputSource(elm)
	if err == nil {
		t.Error("expected error for invalid boolean attribute, got nil")
	}
}

func TestInputSource_BooleanVariations(t *testing.T) {
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
			orig := InputSource(
				ValWithOptions[InputSourceValue]{
					Text:      InputSourceFilm,
					MustHonor: optional.New(c.value),
				},
			)

			elm := orig.toXML("wscn:InputSource")
			decoded, err := decodeInputSource(elm)
			if err != nil {
				t.Fatalf("decode returned error: %v", err)
			}
			if !reflect.DeepEqual(orig, decoded) {
				t.Errorf("expected %+v, got %+v", orig, decoded)
			}
		})
	}
}

func TestInputSourceValue_String(t *testing.T) {
	cases := []struct {
		value    InputSourceValue
		expected string
	}{
		{InputSourceADF, "ADF"},
		{InputSourceADFDuplex, "ADFDuplex"},
		{InputSourceFilm, "Film"},
		{InputSourcePlaten, "Platen"},
		{UnknownInputSource, "Unknown"},
		{InputSourceValue(999), "Unknown"},
	}

	for _, c := range cases {
		t.Run(c.expected, func(t *testing.T) {
			if got := c.value.String(); got != c.expected {
				t.Errorf("expected '%s', got '%s'", c.expected, got)
			}
		})
	}
}

func TestDecodeInputSourceValue(t *testing.T) {
	cases := []struct {
		input    string
		expected InputSourceValue
	}{
		{"ADF", InputSourceADF},
		{"ADFDuplex", InputSourceADFDuplex},
		{"Film", InputSourceFilm},
		{"Platen", InputSourcePlaten},
		{"Unknown", UnknownInputSource},
		{"invalid", UnknownInputSource},
		{"", UnknownInputSource},
	}

	for _, c := range cases {
		t.Run(c.input, func(t *testing.T) {
			if got := DecodeInputSourceValue(c.input); got != c.expected {
				t.Errorf("expected %v, got %v", c.expected, got)
			}
		})
	}
}
