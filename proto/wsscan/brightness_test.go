// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// Test for Brightness

package wsscan

import (
	"reflect"
	"testing"

	"github.com/OpenPrinting/go-mfp/util/optional"
)

func TestBrightness_RoundTrip(t *testing.T) {
	orig := Brightness(
		ValWithOptions[int]{
			Text:        50,
			Override:    optional.New(BooleanElement("true")),
			UsedDefault: optional.New(BooleanElement("0")),
		},
	)

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
	b := Brightness(
		ValWithOptions[int]{
			Text:     25,
			Override: optional.New(BooleanElement("false")),
		},
	)

	elm := b.toXML("wscn:Brightness")

	for _, attr := range elm.Attrs {
		if attr.Name == NsWSCN+":MustHonor" {
			t.Error("MustHonor attribute should not be present in Brightness")
		}
	}
}
