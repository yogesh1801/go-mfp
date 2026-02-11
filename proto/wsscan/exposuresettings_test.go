// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// Test for ExposureSettings

package wsscan

import (
	"reflect"
	"testing"

	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

func TestExposureSettings_RoundTrip_AllChildren(t *testing.T) {
	orig := ExposureSettings{
		Brightness: optional.New(Brightness(
			ValWithOptions[int]{
				Text:     50,
				Override: optional.New(BooleanElement("true")),
			},
		)),
		Contrast: optional.New(Contrast(
			ValWithOptions[int]{
				Text:        75,
				UsedDefault: optional.New(BooleanElement("false")),
			},
		)),
		Sharpness: optional.New(Sharpness(
			ValWithOptions[int]{
				Text:        90,
				Override:    optional.New(BooleanElement("1")),
				UsedDefault: optional.New(BooleanElement("0")),
			},
		)),
	}

	elm := orig.toXML("wscn:ExposureSettings")
	if elm.Name != "wscn:ExposureSettings" {
		t.Errorf("expected element name 'wscn:ExposureSettings', got '%s'",
			elm.Name)
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
		Brightness: optional.New(Brightness(
			ValWithOptions[int]{
				Text: 25,
			},
		)),
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
