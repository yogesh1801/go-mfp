// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// Test for ExposureSettings and its child elements

package wsscan

import (
	"reflect"
	"testing"

	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// Tests for Brightness
func TestBrightness_RoundTrip(t *testing.T) {
	orig := Brightness{
		TextWithBoolAttrs: TextWithBoolAttrs[int]{
			Text:        50,
			Override:    optional.New(BooleanElement("true")),
			UsedDefault: optional.New(BooleanElement("0")),
		},
	}

	elm := orig.toXML("wscn:Brightness")
	if elm.Name != "wscn:Brightness" {
		t.Errorf("expected element name 'wscn:Brightness', got '%s'", elm.Name)
	}
	if elm.Text != "50" {
		t.Errorf("expected text '50', got '%s'", elm.Text)
	}

	decoded, err := decodeBrightness(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}
}

func TestBrightness_NoMustHonor(t *testing.T) {
	b := Brightness{
		TextWithBoolAttrs: TextWithBoolAttrs[int]{
			Text:     25,
			Override: optional.New(BooleanElement("false")),
		},
	}

	elm := b.toXML("wscn:Brightness")
	
	for _, attr := range elm.Attrs {
		if attr.Name == NsWSCN+":MustHonor" {
			t.Error("MustHonor attribute should not be present in Brightness")
		}
	}
}

// Tests for Contrast
func TestContrast_RoundTrip(t *testing.T) {
	orig := Contrast{
		TextWithBoolAttrs: TextWithBoolAttrs[int]{
			Text:        75,
			Override:    optional.New(BooleanElement("false")),
			UsedDefault: optional.New(BooleanElement("1")),
		},
	}

	elm := orig.toXML("wscn:Contrast")
	decoded, err := decodeContrast(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}
}

// Tests for Sharpness
func TestSharpness_RoundTrip(t *testing.T) {
	orig := Sharpness{
		TextWithBoolAttrs: TextWithBoolAttrs[int]{
			Text:        90,
			Override:    optional.New(BooleanElement("1")),
			UsedDefault: optional.New(BooleanElement("false")),
		},
	}

	elm := orig.toXML("wscn:Sharpness")
	decoded, err := decodeSharpness(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}
}

// Tests for ExposureSettings
func TestExposureSettings_RoundTrip_AllChildren(t *testing.T) {
	orig := ExposureSettings{
		Brightness: optional.New(Brightness{
			TextWithBoolAttrs: TextWithBoolAttrs[int]{
				Text:     50,
				Override: optional.New(BooleanElement("true")),
			},
		}),
		Contrast: optional.New(Contrast{
			TextWithBoolAttrs: TextWithBoolAttrs[int]{
				Text:        75,
				UsedDefault: optional.New(BooleanElement("false")),
			},
		}),
		Sharpness: optional.New(Sharpness{
			TextWithBoolAttrs: TextWithBoolAttrs[int]{
				Text:        90,
				Override:    optional.New(BooleanElement("1")),
				UsedDefault: optional.New(BooleanElement("0")),
			},
		}),
	}

	elm := orig.toXML("wscn:ExposureSettings")
	if elm.Name != "wscn:ExposureSettings" {
		t.Errorf("expected element name 'wscn:ExposureSettings', got '%s'", elm.Name)
	}
	if len(elm.Children) != 3 {
		t.Errorf("expected 3 children, got %d", len(elm.Children))
	}

	decoded, err := decodeExposureSettings(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}
}

func TestExposureSettings_RoundTrip_NoChildren(t *testing.T) {
	orig := ExposureSettings{}

	elm := orig.toXML("wscn:ExposureSettings")
	if len(elm.Children) != 0 {
		t.Errorf("expected no children, got %d", len(elm.Children))
	}

	decoded, err := decodeExposureSettings(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}
}

func TestExposureSettings_RoundTrip_OnlyBrightness(t *testing.T) {
	orig := ExposureSettings{
		Brightness: optional.New(Brightness{
			TextWithBoolAttrs: TextWithBoolAttrs[int]{
				Text: 25,
			},
		}),
	}

	elm := orig.toXML("wscn:ExposureSettings")
	if len(elm.Children) != 1 {
		t.Errorf("expected 1 child, got %d", len(elm.Children))
	}

	decoded, err := decodeExposureSettings(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}
}

func TestExposureSettings_InvalidChildValue(t *testing.T) {
	elm := xmldoc.Element{
		Name: "wscn:ExposureSettings",
		Children: []xmldoc.Element{
			{
				Name: "wscn:Brightness",
				Text: "not-a-number",
			},
		},
	}

	_, err := decodeExposureSettings(elm)
	if err == nil {
		t.Error("expected error for invalid brightness value, got nil")
	}
}
