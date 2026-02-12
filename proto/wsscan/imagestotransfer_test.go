// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// Test for ImagesToTransfer

package wsscan

import (
	"reflect"
	"testing"

	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

func TestImagesToTransfer_RoundTrip_AllAttributes(t *testing.T) {
	orig := ImagesToTransfer(
		ValWithOptions[int]{
			Text:        10,
			MustHonor:   optional.New(BooleanElement("true")),
			Override:    optional.New(BooleanElement("false")),
			UsedDefault: optional.New(BooleanElement("1")),
		},
	)

	elm := orig.toXML("wscn:ImagesToTransfer")
	if elm.Name != "wscn:ImagesToTransfer" {
		t.Errorf("expected element name 'wscn:ImagesToTransfer', got '%s'",
			elm.Name)
	}
	if elm.Text != "10" {
		t.Errorf("expected text '10', got '%s'", elm.Text)
	}
	if len(elm.Attrs) != 3 {
		t.Errorf("expected 3 attributes, got %d: %+v", len(elm.Attrs),
			elm.Attrs)
	}

	decoded, err := decodeImagesToTransfer(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}
}

func TestImagesToTransfer_RoundTrip_NoAttributes(t *testing.T) {
	orig := ImagesToTransfer(
		ValWithOptions[int]{
			Text: 5,
		},
	)

	elm := orig.toXML("wscn:ImagesToTransfer")
	if len(elm.Attrs) != 0 {
		t.Errorf("expected no attributes, got %+v", elm.Attrs)
	}

	decoded, err := decodeImagesToTransfer(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}
}

func TestImagesToTransfer_BoundaryValues(t *testing.T) {
	cases := []struct {
		name  string
		value int
	}{
		{"minimum", 0},
		{"small", 1},
		{"medium", 100},
		{"large", 2147483648},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			itt := ImagesToTransfer(
				ValWithOptions[int]{
					Text: c.value,
				},
			)
			elm := itt.toXML("wscn:ImagesToTransfer")

			decoded, err := decodeImagesToTransfer(elm)
			if err != nil {
				t.Errorf("expected no error for value %d, got: %v",
					c.value, err)
			}
			if decoded.Text != c.value {
				t.Errorf("expected value %d, got %d", c.value, decoded.Text)
			}
		})
	}
}

func TestImagesToTransfer_InvalidText(t *testing.T) {
	elm := xmldoc.Element{
		Name: "wscn:ImagesToTransfer",
		Text: "not-a-number",
	}

	_, err := decodeImagesToTransfer(elm)
	if err == nil {
		t.Error("expected error for invalid integer text, got nil")
	}
}

func TestImagesToTransfer_InvalidBooleanAttribute(t *testing.T) {
	elm := xmldoc.Element{
		Name: "wscn:ImagesToTransfer",
		Text: "25",
		Attrs: []xmldoc.Attr{
			{Name: NsWSCN + ":MustHonor", Value: "invalid"},
		},
	}

	_, err := decodeImagesToTransfer(elm)
	if err == nil {
		t.Error("expected error for invalid boolean attribute, got nil")
	}
}

func TestImagesToTransfer_WithMustHonor(t *testing.T) {
	orig := ImagesToTransfer(
		ValWithOptions[int]{
			Text:      20,
			MustHonor: optional.New(BooleanElement("true")),
		},
	)

	elm := orig.toXML("wscn:ImagesToTransfer")
	decoded, err := decodeImagesToTransfer(elm)
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

func TestImagesToTransfer_WithOverride(t *testing.T) {
	orig := ImagesToTransfer(
		ValWithOptions[int]{
			Text:     15,
			Override: optional.New(BooleanElement("false")),
		},
	)

	elm := orig.toXML("wscn:ImagesToTransfer")
	decoded, err := decodeImagesToTransfer(elm)
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

func TestImagesToTransfer_WithUsedDefault(t *testing.T) {
	orig := ImagesToTransfer(
		ValWithOptions[int]{
			Text:        30,
			UsedDefault: optional.New(BooleanElement("1")),
		},
	)

	elm := orig.toXML("wscn:ImagesToTransfer")
	decoded, err := decodeImagesToTransfer(elm)
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
