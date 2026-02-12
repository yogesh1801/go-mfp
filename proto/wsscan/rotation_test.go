// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// Test for Rotation

package wsscan

import (
	"reflect"
	"testing"

	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

func TestRotation_RoundTrip_AllAttributes(t *testing.T) {
	orig := Rotation{
		ValWithOptions: ValWithOptions[RotationValue]{
			Text:        Rotation90,
			MustHonor:   optional.New(BooleanElement("true")),
			Override:    optional.New(BooleanElement("false")),
			UsedDefault: optional.New(BooleanElement("1")),
		},
	}

	elm := orig.toXML("wscn:Rotation")
	if elm.Name != "wscn:Rotation" {
		t.Errorf("expected element name 'wscn:Rotation', got '%s'", elm.Name)
	}
	if elm.Text != "90" {
		t.Errorf("expected text '90', got '%s'", elm.Text)
	}
	if len(elm.Attrs) != 3 {
		t.Errorf("expected 3 attributes, got %d: %+v", len(elm.Attrs),
			elm.Attrs)
	}

	decoded, err := decodeRotation(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}
}

func TestRotation_RoundTrip_NoAttributes(t *testing.T) {
	orig := Rotation{
		ValWithOptions: ValWithOptions[RotationValue]{
			Text: Rotation180,
		},
	}

	elm := orig.toXML("wscn:Rotation")
	if len(elm.Attrs) != 0 {
		t.Errorf("expected no attributes, got %+v", elm.Attrs)
	}

	decoded, err := decodeRotation(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}
}

func TestRotation_AllRotationValues(t *testing.T) {
	cases := []struct {
		name  string
		value RotationValue
		text  string
	}{
		{"0 degrees", Rotation0, "0"},
		{"90 degrees", Rotation90, "90"},
		{"180 degrees", Rotation180, "180"},
		{"270 degrees", Rotation270, "270"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			r := Rotation{
				ValWithOptions: ValWithOptions[RotationValue]{
					Text: c.value,
				},
			}
			elm := r.toXML("wscn:Rotation")

			if elm.Text != c.text {
				t.Errorf("expected text '%s', got '%s'", c.text, elm.Text)
			}

			decoded, err := decodeRotation(elm)
			if err != nil {
				t.Errorf("expected no error for value %s, got: %v", c.text, err)
			}
			if decoded.Text != c.value {
				t.Errorf("expected value %v, got %v", c.value, decoded.Text)
			}
		})
	}
}

func TestRotation_InvalidText(t *testing.T) {
	elm := xmldoc.Element{
		Name: "wscn:Rotation",
		Text: "45",
	}

	_, err := decodeRotation(elm)
	if err == nil {
		t.Error("expected error for invalid rotation value, got nil")
	}
}

func TestRotation_InvalidBooleanAttribute(t *testing.T) {
	elm := xmldoc.Element{
		Name: "wscn:Rotation",
		Text: "90",
		Attrs: []xmldoc.Attr{
			{Name: NsWSCN + ":MustHonor", Value: "invalid"},
		},
	}

	_, err := decodeRotation(elm)
	if err == nil {
		t.Error("expected error for invalid boolean attribute, got nil")
	}
}

func TestRotation_WithMustHonor(t *testing.T) {
	orig := Rotation{
		ValWithOptions: ValWithOptions[RotationValue]{
			Text:      Rotation270,
			MustHonor: optional.New(BooleanElement("true")),
		},
	}

	elm := orig.toXML("wscn:Rotation")
	decoded, err := decodeRotation(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}

	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}

	// Verify MustHonor attribute is present
	if decoded.MustHonor == nil {
		t.Error("expected MustHonor attribute to be present")
	}
	if !optional.Get(decoded.MustHonor).Bool() {
		t.Error("expected MustHonor to be true")
	}
}

func TestRotation_WithOverride(t *testing.T) {
	orig := Rotation{
		ValWithOptions: ValWithOptions[RotationValue]{
			Text:     Rotation0,
			Override: optional.New(BooleanElement("false")),
		},
	}

	elm := orig.toXML("wscn:Rotation")
	decoded, err := decodeRotation(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}

	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}

	// Verify Override attribute is present
	if decoded.Override == nil {
		t.Error("expected Override attribute to be present")
	}
	if optional.Get(decoded.Override).Bool() {
		t.Error("expected Override to be false")
	}
}

func TestRotation_WithUsedDefault(t *testing.T) {
	orig := Rotation{
		ValWithOptions: ValWithOptions[RotationValue]{
			Text:        Rotation180,
			UsedDefault: optional.New(BooleanElement("1")),
		},
	}

	elm := orig.toXML("wscn:Rotation")
	decoded, err := decodeRotation(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}

	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}

	// Verify UsedDefault attribute is present
	if decoded.UsedDefault == nil {
		t.Error("expected UsedDefault attribute to be present")
	}
	if !optional.Get(decoded.UsedDefault).Bool() {
		t.Error("expected UsedDefault to be true")
	}
}

func TestRotation_BooleanVariations(t *testing.T) {
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
			orig := Rotation{
				ValWithOptions: ValWithOptions[RotationValue]{
					Text:        Rotation90,
					MustHonor:   optional.New(c.value),
					Override:    optional.New(c.value),
					UsedDefault: optional.New(c.value),
				},
			}

			elm := orig.toXML("wscn:Rotation")
			decoded, err := decodeRotation(elm)
			if err != nil {
				t.Fatalf("decode returned error: %v", err)
			}
			if !reflect.DeepEqual(orig, decoded) {
				t.Errorf("expected %+v, got %+v", orig, decoded)
			}
		})
	}
}
