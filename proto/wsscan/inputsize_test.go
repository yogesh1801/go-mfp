// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// Test for InputSize

package wsscan

import (
	"reflect"
	"testing"

	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

func TestInputSize_RoundTrip_WithDocumentSizeAutoDetect(t *testing.T) {
	orig := InputSize{
		MustHonor:              optional.New(BooleanElement("true")),
		DocumentSizeAutoDetect: optional.New(BooleanElement("1")),
		InputMediaSize: InputMediaSize{
			Height: ValWithOptions[int]{Text: 1200},
			Width:  ValWithOptions[int]{Text: 850},
		},
	}

	elm := orig.toXML("wscn:InputSize")
	if elm.Name != "wscn:InputSize" {
		t.Errorf("expected element name 'wscn:InputSize', got '%s'", elm.Name)
	}
	if len(elm.Attrs) != 1 {
		t.Errorf("expected 1 attribute, got %d", len(elm.Attrs))
	}
	if len(elm.Children) != 2 {
		t.Errorf("expected 2 children, got %d", len(elm.Children))
	}

	decoded, err := decodeInputSize(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}
}

func TestInputSize_RoundTrip_WithInputMediaSize(t *testing.T) {
	orig := InputSize{
		MustHonor: optional.New(BooleanElement("false")),
		InputMediaSize: InputMediaSize{
			Height: ValWithOptions[int]{Text: 1200},
			Width:  ValWithOptions[int]{Text: 850},
		},
	}

	elm := orig.toXML("wscn:InputSize")
	if len(elm.Children) != 1 {
		t.Errorf("expected 1 child, got %d", len(elm.Children))
	}

	decoded, err := decodeInputSize(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}
}

func TestInputSize_RoundTrip_NoMustHonor(t *testing.T) {
	orig := InputSize{
		DocumentSizeAutoDetect: optional.New(BooleanElement("false")),
		InputMediaSize: InputMediaSize{
			Height: ValWithOptions[int]{Text: 1200},
			Width:  ValWithOptions[int]{Text: 850},
		},
	}

	elm := orig.toXML("wscn:InputSize")
	if len(elm.Attrs) != 0 {
		t.Errorf("expected no attributes, got %d", len(elm.Attrs))
	}

	decoded, err := decodeInputSize(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}
}

func TestInputSize_RoundTrip_BothChildren(t *testing.T) {
	// Both DocumentSizeAutoDetect and InputMediaSize present
	orig := InputSize{
		DocumentSizeAutoDetect: optional.New(BooleanElement("true")),
		InputMediaSize: InputMediaSize{
			Height: ValWithOptions[int]{Text: 2000},
			Width:  ValWithOptions[int]{Text: 1500},
		},
	}

	elm := orig.toXML("wscn:InputSize")
	if len(elm.Children) != 2 {
		t.Errorf("expected 2 children, got %d", len(elm.Children))
	}

	decoded, err := decodeInputSize(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}
}

func TestInputSize_OnlyInputMediaSize(t *testing.T) {
	orig := InputSize{
		MustHonor: optional.New(BooleanElement("1")),
		InputMediaSize: InputMediaSize{
			Height: ValWithOptions[int]{Text: 1200},
			Width:  ValWithOptions[int]{Text: 850},
		},
	}

	elm := orig.toXML("wscn:InputSize")
	if len(elm.Children) != 1 {
		t.Errorf("expected 1 child, got %d", len(elm.Children))
	}

	decoded, err := decodeInputSize(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}
}

func TestInputSize_InvalidMustHonor(t *testing.T) {
	elm := xmldoc.Element{
		Name: "wscn:InputSize",
		Attrs: []xmldoc.Attr{
			{Name: NsWSCN + ":MustHonor", Value: "invalid"},
		},
	}

	_, err := decodeInputSize(elm)
	if err == nil {
		t.Error("expected error for invalid MustHonor attribute, got nil")
	}
}

func TestInputSize_InvalidDocumentSizeAutoDetect(t *testing.T) {
	elm := xmldoc.Element{
		Name: "wscn:InputSize",
		Children: []xmldoc.Element{
			{
				Name: NsWSCN + ":DocumentSizeAutoDetect",
				Text: "invalid",
			},
			{
				Name: NsWSCN + ":InputMediaSize",
				Children: []xmldoc.Element{
					{Name: NsWSCN + ":Height", Text: "1200"},
					{Name: NsWSCN + ":Width", Text: "850"},
				},
			},
		},
	}

	_, err := decodeInputSize(elm)
	if err == nil {
		t.Error("expected error for invalid DocumentSizeAutoDetect value, got nil")
	}
}

func TestInputSize_MustHonorTrue(t *testing.T) {
	orig := InputSize{
		MustHonor:              optional.New(BooleanElement("true")),
		DocumentSizeAutoDetect: optional.New(BooleanElement("0")),
		InputMediaSize: InputMediaSize{
			Height: ValWithOptions[int]{Text: 1200},
			Width:  ValWithOptions[int]{Text: 850},
		},
	}

	elm := orig.toXML("wscn:InputSize")
	decoded, err := decodeInputSize(elm)
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
