// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// Test for CompressionQualityFactor

package wsscan

import (
	"reflect"
	"testing"

	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

func TestCompressionQualityFactor_RoundTrip_AllAttributes(t *testing.T) {
	orig := CompressionQualityFactor(ValWithBoolAttr[int]{
		Text:        85,
		MustHonor:   optional.New(BooleanElement("true")),
		Override:    optional.New(BooleanElement("false")),
		UsedDefault: optional.New(BooleanElement("1")),
	})

	elm := orig.toXML("wscn:CompressionQualityFactor")
	if elm.Name != "wscn:CompressionQualityFactor" {
		t.Errorf(
			"expected element name 'wscn:CompressionQualityFactor', got '%s'",
			elm.Name)
	}
	if elm.Text != "85" {
		t.Errorf("expected text '85', got '%s'", elm.Text)
	}
	if len(elm.Attrs) != 3 {
		t.Errorf(
			"expected 3 attributes, got %d: %+v",
			len(elm.Attrs),
			elm.Attrs,
		)
	}

	decoded, err := decodeCompressionQualityFactor(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}
}

func TestCompressionQualityFactor_RoundTrip_NoAttributes(t *testing.T) {
	orig := CompressionQualityFactor(ValWithBoolAttr[int]{
		Text: 50,
	})

	elm := orig.toXML("wscn:CompressionQualityFactor")
	if len(elm.Attrs) != 0 {
		t.Errorf(
			"expected no attributes, got %+v",
			elm.Attrs,
		)
	}

	decoded, err := decodeCompressionQualityFactor(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}
}

func TestCompressionQualityFactor_BoundaryValues(t *testing.T) {
	cases := []struct {
		name  string
		value int
	}{
		{"minimum", 0},
		{"maximum", 100},
		{"middle", 50},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			cqf := CompressionQualityFactor(ValWithBoolAttr[int]{
				Text: c.value,
			})
			elm := cqf.toXML("wscn:CompressionQualityFactor")

			decoded, err := decodeCompressionQualityFactor(elm)
			if err != nil {
				t.Errorf(
					"expected no error for value %d, got: %v",
					c.value,
					err,
				)
			}
			if decoded.Text != c.value {
				t.Errorf(
					"expected value %d, got %d",
					c.value,
					decoded.Text,
				)
			}
		})
	}
}

func TestCompressionQualityFactor_InvalidText(t *testing.T) {
	elm := xmldoc.Element{
		Name: "wscn:CompressionQualityFactor",
		Text: "not-a-number",
	}

	_, err := decodeCompressionQualityFactor(elm)
	if err == nil {
		t.Error(
			"expected error for invalid integer text, got nil",
		)
	}
}

func TestCompressionQualityFactor_InvalidBooleanAttribute(t *testing.T) {
	elm := xmldoc.Element{
		Name: "wscn:CompressionQualityFactor",
		Text: "75",
		Attrs: []xmldoc.Attr{
			{Name: NsWSCN + ":MustHonor", Value: "invalid"},
		},
	}

	_, err := decodeCompressionQualityFactor(elm)
	if err == nil {
		t.Error(
			"expected error for invalid boolean attribute, got nil",
		)
	}
}

func TestCompressionQualityFactor_WithMustHonor(t *testing.T) {
	orig := CompressionQualityFactor(ValWithBoolAttr[int]{
		Text:      100,
		MustHonor: optional.New(BooleanElement("true")),
	})

	elm := orig.toXML("wscn:CompressionQualityFactor")
	decoded, err := decodeCompressionQualityFactor(elm)
	if err != nil {
		t.Fatalf(
			"decode returned error: %v",
			err,
		)
	}

	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf(
			"expected %+v, got %+v",
			orig,
			decoded,
		)
	}

	// Verify MustHonor attribute is present
	if decoded.MustHonor == nil {
		t.Error("expected MustHonor attribute to be present")
	}
	if !optional.Get(decoded.MustHonor).Bool() {
		t.Error("expected MustHonor to be true")
	}
}

func TestCompressionQualityFactor_WithOverride(t *testing.T) {
	orig := CompressionQualityFactor(ValWithBoolAttr[int]{
		Text:     25,
		Override: optional.New(BooleanElement("false")),
	})

	elm := orig.toXML("wscn:CompressionQualityFactor")
	decoded, err := decodeCompressionQualityFactor(elm)
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

func TestCompressionQualityFactor_WithUsedDefault(t *testing.T) {
	orig := CompressionQualityFactor(ValWithBoolAttr[int]{
		Text:        75,
		UsedDefault: optional.New(BooleanElement("1")),
	})

	elm := orig.toXML("wscn:CompressionQualityFactor")
	decoded, err := decodeCompressionQualityFactor(elm)
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
