// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// Test for Exposure

package wsscan

import (
	"reflect"
	"testing"

	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

func TestExposure_RoundTrip_WithAutoExposure(t *testing.T) {
	orig := Exposure{
		MustHonor:    optional.New(BooleanElement("true")),
		AutoExposure: optional.New(BooleanElement("1")),
	}

	elm := orig.toXML("wscn:Exposure")
	if elm.Name != "wscn:Exposure" {
		t.Errorf("expected element name 'wscn:Exposure', got '%s'", elm.Name)
	}
	if len(elm.Attrs) != 1 {
		t.Errorf("expected 1 attribute, got %d", len(elm.Attrs))
	}
	if len(elm.Children) != 1 {
		t.Errorf("expected 1 child, got %d", len(elm.Children))
	}

	decoded, err := decodeExposure(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}
}

func TestExposure_RoundTrip_WithExposureSettings(t *testing.T) {
	orig := Exposure{
		MustHonor: optional.New(BooleanElement("false")),
		ExposureSettings: optional.New(ExposureSettings{
			Brightness: optional.New(Brightness{
				TextWithBoolAttrs: TextWithBoolAttrs[int]{
					Text: 50,
				},
			}),
			Contrast: optional.New(Contrast{
				TextWithBoolAttrs: TextWithBoolAttrs[int]{
					Text: 75,
				},
			}),
		}),
	}

	elm := orig.toXML("wscn:Exposure")
	if len(elm.Children) != 1 {
		t.Errorf("expected 1 child, got %d", len(elm.Children))
	}

	decoded, err := decodeExposure(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}
}

func TestExposure_RoundTrip_NoMustHonor(t *testing.T) {
	orig := Exposure{
		AutoExposure: optional.New(BooleanElement("false")),
	}

	elm := orig.toXML("wscn:Exposure")
	if len(elm.Attrs) != 0 {
		t.Errorf("expected no attributes, got %d", len(elm.Attrs))
	}

	decoded, err := decodeExposure(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}
}

func TestExposure_RoundTrip_BothChildren(t *testing.T) {
	// Both AutoExposure and ExposureSettings can be present
	orig := Exposure{
		AutoExposure: optional.New(BooleanElement("true")),
		ExposureSettings: optional.New(ExposureSettings{
			Sharpness: optional.New(Sharpness{
				TextWithBoolAttrs: TextWithBoolAttrs[int]{
					Text: 90,
				},
			}),
		}),
	}

	elm := orig.toXML("wscn:Exposure")
	if len(elm.Children) != 2 {
		t.Errorf("expected 2 children, got %d", len(elm.Children))
	}

	decoded, err := decodeExposure(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}
}

func TestExposure_RoundTrip_NoChildren(t *testing.T) {
	orig := Exposure{
		MustHonor: optional.New(BooleanElement("1")),
	}

	elm := orig.toXML("wscn:Exposure")
	if len(elm.Children) != 0 {
		t.Errorf("expected no children, got %d", len(elm.Children))
	}

	decoded, err := decodeExposure(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}
}

func TestExposure_InvalidMustHonor(t *testing.T) {
	elm := xmldoc.Element{
		Name: "wscn:Exposure",
		Attrs: []xmldoc.Attr{
			{Name: NsWSCN + ":MustHonor", Value: "invalid"},
		},
	}

	_, err := decodeExposure(elm)
	if err == nil {
		t.Error("expected error for invalid MustHonor attribute, got nil")
	}
}

func TestExposure_InvalidAutoExposure(t *testing.T) {
	elm := xmldoc.Element{
		Name: "wscn:Exposure",
		Children: []xmldoc.Element{
			{
				Name: NsWSCN + ":AutoExposure",
				Text: "invalid",
			},
		},
	}

	_, err := decodeExposure(elm)
	if err == nil {
		t.Error("expected error for invalid AutoExposure value, got nil")
	}
}

func TestExposure_MustHonorTrue(t *testing.T) {
	orig := Exposure{
		MustHonor:    optional.New(BooleanElement("true")),
		AutoExposure: optional.New(BooleanElement("0")),
	}

	elm := orig.toXML("wscn:Exposure")
	decoded, err := decodeExposure(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}

	// Verify MustHonor attribute is present
	if decoded.MustHonor == nil {
		t.Error("expected MustHonor to be present")
	}
	if !optional.Get(decoded.MustHonor).Bool() {
		t.Error("expected MustHonor to be true")
	}
}
