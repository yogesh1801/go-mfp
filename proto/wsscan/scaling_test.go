// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// Test for Scaling

package wsscan

import (
	"reflect"
	"testing"

	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

func TestScaling_RoundTrip_AllAttributes(t *testing.T) {
	orig := Scaling{
		ScalingHeight: ValWithOptions[int]{
			Text:        100,
			Override:    optional.New(BooleanElement("true")),
			UsedDefault: optional.New(BooleanElement("false")),
		},
		ScalingWidth: ValWithOptions[int]{
			Text:        100,
			Override:    optional.New(BooleanElement("0")),
			UsedDefault: optional.New(BooleanElement("1")),
		},
		MustHonor: optional.New(BooleanElement("true")),
	}

	elm := orig.toXML("wscn:Scaling")
	if elm.Name != "wscn:Scaling" {
		t.Errorf("expected element name 'wscn:Scaling', got '%s'", elm.Name)
	}
	if len(elm.Attrs) != 1 {
		t.Errorf("expected 1 attribute, got %d: %+v", len(elm.Attrs), elm.Attrs)
	}
	if len(elm.Children) != 2 {
		t.Errorf("expected 2 children, got %d", len(elm.Children))
	}

	decoded, err := decodeScaling(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}
}

func TestScaling_RoundTrip_NoAttributes(t *testing.T) {
	orig := Scaling{
		ScalingHeight: ValWithOptions[int]{Text: 50},
		ScalingWidth:  ValWithOptions[int]{Text: 50},
	}

	elm := orig.toXML("wscn:Scaling")
	if len(elm.Attrs) != 0 {
		t.Errorf("expected no attributes, got %+v", elm.Attrs)
	}

	decoded, err := decodeScaling(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}
}

func TestScaling_RoundTrip_PartialAttributes(t *testing.T) {
	orig := Scaling{
		ScalingHeight: ValWithOptions[int]{
			Text:     200,
			Override: optional.New(BooleanElement("true")),
		},
		ScalingWidth: ValWithOptions[int]{
			Text:        150,
			UsedDefault: optional.New(BooleanElement("false")),
		},
		MustHonor: optional.New(BooleanElement("1")),
	}

	elm := orig.toXML("wscn:Scaling")
	decoded, err := decodeScaling(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}
}

func TestScaling_DifferentHeightWidth(t *testing.T) {
	orig := Scaling{
		ScalingHeight: ValWithOptions[int]{Text: 75},
		ScalingWidth:  ValWithOptions[int]{Text: 125},
	}

	elm := orig.toXML("wscn:Scaling")

	// Check ScalingHeight child
	heightElem := elm.Children[0]
	if heightElem.Name != "wscn:ScalingHeight" {
		t.Errorf("expected first child 'wscn:ScalingHeight', got '%s'", heightElem.Name)
	}
	if heightElem.Text != "75" {
		t.Errorf("expected ScalingHeight text '75', got '%s'", heightElem.Text)
	}

	// Check ScalingWidth child
	widthElem := elm.Children[1]
	if widthElem.Name != "wscn:ScalingWidth" {
		t.Errorf("expected second child 'wscn:ScalingWidth', got '%s'", widthElem.Name)
	}
	if widthElem.Text != "125" {
		t.Errorf("expected ScalingWidth text '125', got '%s'", widthElem.Text)
	}

	decoded, err := decodeScaling(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}
}

func TestScaling_MissingHeight(t *testing.T) {
	elm := xmldoc.Element{
		Name: "wscn:Scaling",
		Children: []xmldoc.Element{
			{Name: NsWSCN + ":ScalingWidth", Text: "100"},
		},
	}

	_, err := decodeScaling(elm)
	if err == nil {
		t.Error("expected error for missing ScalingHeight, got nil")
	}
}

func TestScaling_MissingWidth(t *testing.T) {
	elm := xmldoc.Element{
		Name: "wscn:Scaling",
		Children: []xmldoc.Element{
			{Name: NsWSCN + ":ScalingHeight", Text: "100"},
		},
	}

	_, err := decodeScaling(elm)
	if err == nil {
		t.Error("expected error for missing ScalingWidth, got nil")
	}
}

func TestScaling_InvalidMustHonor(t *testing.T) {
	elm := xmldoc.Element{
		Name: "wscn:Scaling",
		Attrs: []xmldoc.Attr{
			{Name: NsWSCN + ":MustHonor", Value: "invalid"},
		},
		Children: []xmldoc.Element{
			{Name: NsWSCN + ":ScalingHeight", Text: "100"},
			{Name: NsWSCN + ":ScalingWidth", Text: "100"},
		},
	}

	_, err := decodeScaling(elm)
	if err == nil {
		t.Error("expected error for invalid MustHonor attribute, got nil")
	}
}

func TestScaling_InvalidOverride(t *testing.T) {
	elm := xmldoc.Element{
		Name: "wscn:Scaling",
		Children: []xmldoc.Element{
			{
				Name: NsWSCN + ":ScalingHeight",
				Text: "100",
				Attrs: []xmldoc.Attr{
					{Name: NsWSCN + ":Override", Value: "invalid"},
				},
			},
			{Name: NsWSCN + ":ScalingWidth", Text: "100"},
		},
	}

	_, err := decodeScaling(elm)
	if err == nil {
		t.Error("expected error for invalid Override attribute, got nil")
	}
}

func TestScaling_InvalidValue(t *testing.T) {
	elm := xmldoc.Element{
		Name: "wscn:Scaling",
		Children: []xmldoc.Element{
			{Name: NsWSCN + ":ScalingHeight", Text: "not-a-number"},
			{Name: NsWSCN + ":ScalingWidth", Text: "100"},
		},
	}

	_, err := decodeScaling(elm)
	if err == nil {
		t.Error("expected error for invalid ScalingHeight value, got nil")
	}
}

func TestScaling_BooleanVariations(t *testing.T) {
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
			orig := Scaling{
				ScalingHeight: ValWithOptions[int]{
					Text:        100,
					Override:    optional.New(c.value),
					UsedDefault: optional.New(c.value),
				},
				ScalingWidth: ValWithOptions[int]{
					Text:        100,
					Override:    optional.New(c.value),
					UsedDefault: optional.New(c.value),
				},
				MustHonor: optional.New(c.value),
			}

			elm := orig.toXML("wscn:Scaling")
			decoded, err := decodeScaling(elm)
			if err != nil {
				t.Fatalf("decode returned error: %v", err)
			}
			if !reflect.DeepEqual(orig, decoded) {
				t.Errorf("expected %+v, got %+v", orig, decoded)
			}
		})
	}
}
