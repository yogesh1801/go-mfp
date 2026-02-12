// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// Test for InputMediaSize

package wsscan

import (
	"reflect"
	"testing"

	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

func TestInputMediaSize_RoundTrip(t *testing.T) {
	orig := InputMediaSize{
		Height: ValWithOptions[int]{
			Text:     1200,
			Override: optional.New(BooleanElement("true")),
		},
		Width: ValWithOptions[int]{
			Text:        850,
			UsedDefault: optional.New(BooleanElement("false")),
		},
	}

	elm := orig.toXML("wscn:InputMediaSize")
	if elm.Name != "wscn:InputMediaSize" {
		t.Errorf("expected element name 'wscn:InputMediaSize', got '%s'", elm.Name)
	}
	if len(elm.Children) != 2 {
		t.Errorf("expected 2 children, got %d", len(elm.Children))
	}

	decoded, err := decodeInputMediaSize(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}
}

func TestInputMediaSize_MissingHeight(t *testing.T) {
	elm := xmldoc.Element{
		Name: "wscn:InputMediaSize",
		Children: []xmldoc.Element{
			{
				Name: NsWSCN + ":Width",
				Text: "850",
			},
		},
	}

	_, err := decodeInputMediaSize(elm)
	if err == nil {
		t.Error("expected error for missing Height, got nil")
	}
}

func TestInputMediaSize_MissingWidth(t *testing.T) {
	elm := xmldoc.Element{
		Name: "wscn:InputMediaSize",
		Children: []xmldoc.Element{
			{
				Name: NsWSCN + ":Height",
				Text: "1200",
			},
		},
	}

	_, err := decodeInputMediaSize(elm)
	if err == nil {
		t.Error("expected error for missing Width, got nil")
	}
}

func TestInputMediaSize_InvalidHeight(t *testing.T) {
	elm := xmldoc.Element{
		Name: "wscn:InputMediaSize",
		Children: []xmldoc.Element{
			{
				Name: NsWSCN + ":Height",
				Text: "not-a-number",
			},
			{
				Name: NsWSCN + ":Width",
				Text: "850",
			},
		},
	}

	_, err := decodeInputMediaSize(elm)
	if err == nil {
		t.Error("expected error for invalid Height value, got nil")
	}
}

func TestInputMediaSize_InvalidWidth(t *testing.T) {
	elm := xmldoc.Element{
		Name: "wscn:InputMediaSize",
		Children: []xmldoc.Element{
			{
				Name: NsWSCN + ":Height",
				Text: "1200",
			},
			{
				Name: NsWSCN + ":Width",
				Text: "invalid",
			},
		},
	}

	_, err := decodeInputMediaSize(elm)
	if err == nil {
		t.Error("expected error for invalid Width value, got nil")
	}
}
