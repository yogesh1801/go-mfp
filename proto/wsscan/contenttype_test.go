// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// Test for ContentType

package wsscan

import (
	"reflect"
	"testing"

	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

func TestContentType_RoundTrip_AllAttributes(t *testing.T) {
	orig := ContentType{
		TextWithBoolAttrs: TextWithBoolAttrs[ContentTypeValue]{
			Text:        Photo,
			MustHonor:   optional.New(BooleanElement("true")),
			Override:    optional.New(BooleanElement("false")),
			UsedDefault: optional.New(BooleanElement("1")),
		},
	}

	elm := orig.toXML("wscn:ContentType")
	if elm.Name != "wscn:ContentType" {
		t.Errorf("expected element name 'wscn:ContentType', got '%s'", elm.Name)
	}
	if elm.Text != "Photo" {
		t.Errorf("expected text 'Photo', got '%s'", elm.Text)
	}
	if len(elm.Attrs) != 3 {
		t.Errorf("expected 3 attributes, got %d: %+v", len(elm.Attrs), elm.Attrs)
	}

	decoded, err := decodeContentType(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}
}

func TestContentType_RoundTrip_NoAttributes(t *testing.T) {
	orig := ContentType{
		TextWithBoolAttrs: TextWithBoolAttrs[ContentTypeValue]{
			Text: Auto,
		},
	}

	elm := orig.toXML("wscn:ContentType")
	if len(elm.Attrs) != 0 {
		t.Errorf("expected no attributes, got %+v", elm.Attrs)
	}

	decoded, err := decodeContentType(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}
}

func TestContentType_AllValues(t *testing.T) {
	cases := []struct {
		name     string
		value    ContentTypeValue
		expected string
	}{
		{"Auto", Auto, "Auto"},
		{"Text", Text, "Text"},
		{"Photo", Photo, "Photo"},
		{"Halftone", Halftone, "Halftone"},
		{"Mixed", Mixed, "Mixed"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ct := ContentType{
				TextWithBoolAttrs: TextWithBoolAttrs[ContentTypeValue]{
					Text: c.value,
				},
			}
			elm := ct.toXML("wscn:ContentType")

			if elm.Text != c.expected {
				t.Errorf("expected text '%s', got '%s'", c.expected, elm.Text)
			}

			decoded, err := decodeContentType(elm)
			if err != nil {
				t.Fatalf("decode returned error: %v", err)
			}
			if decoded.Text != c.value {
				t.Errorf("expected value %v, got %v", c.value, decoded.Text)
			}
		})
	}
}

func TestContentType_UnknownValue(t *testing.T) {
	elm := xmldoc.Element{
		Name: "wscn:ContentType",
		Text: "UnknownType",
	}

	decoded, err := decodeContentType(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}
	if decoded.Text != UnknownContentTypeValue {
		t.Errorf("expected UnknownContentTypeValue, got %v", decoded.Text)
	}
}

func TestContentType_WithMustHonor(t *testing.T) {
	orig := ContentType{
		TextWithBoolAttrs: TextWithBoolAttrs[ContentTypeValue]{
			Text:      Text,
			MustHonor: optional.New(BooleanElement("true")),
		},
	}

	elm := orig.toXML("wscn:ContentType")
	decoded, err := decodeContentType(elm)
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

func TestContentType_WithOverride(t *testing.T) {
	orig := ContentType{
		TextWithBoolAttrs: TextWithBoolAttrs[ContentTypeValue]{
			Text:     Halftone,
			Override: optional.New(BooleanElement("0")),
		},
	}

	elm := orig.toXML("wscn:ContentType")
	decoded, err := decodeContentType(elm)
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

func TestContentType_WithUsedDefault(t *testing.T) {
	orig := ContentType{
		TextWithBoolAttrs: TextWithBoolAttrs[ContentTypeValue]{
			Text:        Mixed,
			UsedDefault: optional.New(BooleanElement("true")),
		},
	}

	elm := orig.toXML("wscn:ContentType")
	decoded, err := decodeContentType(elm)
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

func TestContentType_InvalidBooleanAttribute(t *testing.T) {
	elm := xmldoc.Element{
		Name: "wscn:ContentType",
		Text: "Photo",
		Attrs: []xmldoc.Attr{
			{Name: NsWSCN + ":MustHonor", Value: "invalid"},
		},
	}

	_, err := decodeContentType(elm)
	if err == nil {
		t.Error("expected error for invalid boolean attribute, got nil")
	}
}
