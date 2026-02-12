// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// Test for Resolution

package wsscan

import (
	"reflect"
	"testing"

	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

func TestResolution_RoundTrip_AllAttributes(t *testing.T) {
	orig := Resolution{
		Height: ValWithOptions[int]{
			Text:        300,
			Override:    optional.New(BooleanElement("true")),
			UsedDefault: optional.New(BooleanElement("false")),
		},
		Width: ValWithOptions[int]{
			Text:        300,
			Override:    optional.New(BooleanElement("0")),
			UsedDefault: optional.New(BooleanElement("1")),
		},
		MustHonor: optional.New(BooleanElement("true")),
	}

	elm := orig.toXML("wscn:Resolution")
	if elm.Name != "wscn:Resolution" {
		t.Errorf("expected element name 'wscn:Resolution', got '%s'", elm.Name)
	}
	if len(elm.Attrs) != 1 {
		t.Errorf("expected 1 attribute, got %d: %+v", len(elm.Attrs), elm.Attrs)
	}
	if len(elm.Children) != 2 {
		t.Errorf("expected 2 children, got %d", len(elm.Children))
	}

	decoded, err := decodeResolution(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}
}

func TestResolution_RoundTrip_NoAttributes(t *testing.T) {
	orig := Resolution{
		Height: ValWithOptions[int]{Text: 600},
		Width:  ValWithOptions[int]{Text: 600},
	}

	elm := orig.toXML("wscn:Resolution")
	if len(elm.Attrs) != 0 {
		t.Errorf("expected no attributes, got %+v", elm.Attrs)
	}

	decoded, err := decodeResolution(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}
}

func TestResolution_RoundTrip_PartialAttributes(t *testing.T) {
	orig := Resolution{
		Height: ValWithOptions[int]{
			Text:     1200,
			Override: optional.New(BooleanElement("true")),
		},
		Width: ValWithOptions[int]{
			Text:        1200,
			UsedDefault: optional.New(BooleanElement("false")),
		},
		MustHonor: optional.New(BooleanElement("1")),
	}

	elm := orig.toXML("wscn:Resolution")
	decoded, err := decodeResolution(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}
}

func TestResolution_DifferentHeightWidth(t *testing.T) {
	orig := Resolution{
		Height: ValWithOptions[int]{Text: 300},
		Width:  ValWithOptions[int]{Text: 600},
	}

	elm := orig.toXML("wscn:Resolution")

	// Check Height child
	heightElem := elm.Children[0]
	if heightElem.Name != "wscn:Height" {
		t.Errorf("expected first child 'wscn:Height', got '%s'", heightElem.Name)
	}
	if heightElem.Text != "300" {
		t.Errorf("expected Height text '300', got '%s'", heightElem.Text)
	}

	// Check Width child
	widthElem := elm.Children[1]
	if widthElem.Name != "wscn:Width" {
		t.Errorf("expected second child 'wscn:Width', got '%s'", widthElem.Name)
	}
	if widthElem.Text != "600" {
		t.Errorf("expected Width text '600', got '%s'", widthElem.Text)
	}

	decoded, err := decodeResolution(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}
}

func TestResolution_MissingHeight(t *testing.T) {
	elm := xmldoc.Element{
		Name: "wscn:Resolution",
		Children: []xmldoc.Element{
			{Name: NsWSCN + ":Width", Text: "300"},
		},
	}

	_, err := decodeResolution(elm)
	if err == nil {
		t.Error("expected error for missing Height, got nil")
	}
}

func TestResolution_MissingWidth(t *testing.T) {
	elm := xmldoc.Element{
		Name: "wscn:Resolution",
		Children: []xmldoc.Element{
			{Name: NsWSCN + ":Height", Text: "300"},
		},
	}

	_, err := decodeResolution(elm)
	if err == nil {
		t.Error("expected error for missing Width, got nil")
	}
}

func TestResolution_InvalidMustHonor(t *testing.T) {
	elm := xmldoc.Element{
		Name: "wscn:Resolution",
		Attrs: []xmldoc.Attr{
			{Name: NsWSCN + ":MustHonor", Value: "invalid"},
		},
		Children: []xmldoc.Element{
			{Name: NsWSCN + ":Height", Text: "300"},
			{Name: NsWSCN + ":Width", Text: "300"},
		},
	}

	_, err := decodeResolution(elm)
	if err == nil {
		t.Error("expected error for invalid MustHonor attribute, got nil")
	}
}

func TestResolutionDimension_InvalidOverride(t *testing.T) {
	elm := xmldoc.Element{
		Name: "wscn:Resolution",
		Children: []xmldoc.Element{
			{
				Name: NsWSCN + ":Height",
				Text: "300",
				Attrs: []xmldoc.Attr{
					{Name: NsWSCN + ":Override", Value: "invalid"},
				},
			},
			{Name: NsWSCN + ":Width", Text: "300"},
		},
	}

	_, err := decodeResolution(elm)
	if err == nil {
		t.Error("expected error for invalid Override attribute, got nil")
	}
}

func TestResolutionDimension_InvalidValue(t *testing.T) {
	elm := xmldoc.Element{
		Name: "wscn:Resolution",
		Children: []xmldoc.Element{
			{Name: NsWSCN + ":Height", Text: "not-a-number"},
			{Name: NsWSCN + ":Width", Text: "300"},
		},
	}

	_, err := decodeResolution(elm)
	if err == nil {
		t.Error("expected error for invalid Height value, got nil")
	}
}

func TestResolution_BooleanVariations(t *testing.T) {
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
			orig := Resolution{
				Height: ValWithOptions[int]{
					Text:        150,
					Override:    optional.New(c.value),
					UsedDefault: optional.New(c.value),
				},
				Width: ValWithOptions[int]{
					Text:        150,
					Override:    optional.New(c.value),
					UsedDefault: optional.New(c.value),
				},
				MustHonor: optional.New(c.value),
			}

			elm := orig.toXML("wscn:Resolution")
			decoded, err := decodeResolution(elm)
			if err != nil {
				t.Fatalf("decode returned error: %v", err)
			}
			if !reflect.DeepEqual(orig, decoded) {
				t.Errorf("expected %+v, got %+v", orig, decoded)
			}
		})
	}
}
