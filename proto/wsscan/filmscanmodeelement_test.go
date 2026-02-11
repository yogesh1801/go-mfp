// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// Test for FilmScanModeElement

package wsscan

import (
	"reflect"
	"testing"

	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

func TestFilmScanModeElement_RoundTrip_AllAttributes(t *testing.T) {
	orig := FilmScanModeElement(
		ValWithOptions[string]{
			Text:        "ColorSlideFilm",
			MustHonor:   optional.New(BooleanElement("true")),
			Override:    optional.New(BooleanElement("false")),
			UsedDefault: optional.New(BooleanElement("1")),
		},
	)

	elm := orig.toXML("wscn:FilmScanMode")
	if elm.Name != "wscn:FilmScanMode" {
		t.Errorf("expected element name 'wscn:FilmScanMode', got '%s'",
			elm.Name)
	}
	if elm.Text != "ColorSlideFilm" {
		t.Errorf("expected text 'ColorSlideFilm', got '%s'", elm.Text)
	}
	if len(elm.Attrs) != 3 {
		t.Errorf("expected 3 attributes, got %d: %+v", len(elm.Attrs),
			elm.Attrs)
	}

	decoded, err := decodeFilmScanModeElement(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}
}

func TestFilmScanModeElement_RoundTrip_NoAttributes(t *testing.T) {
	orig := FilmScanModeElement(
		ValWithOptions[string]{
			Text: "NotApplicable",
		},
	)

	elm := orig.toXML("wscn:FilmScanMode")
	if len(elm.Attrs) != 0 {
		t.Errorf("expected no attributes, got %+v", elm.Attrs)
	}

	decoded, err := decodeFilmScanModeElement(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}
}

func TestFilmScanModeElement_StandardValues(t *testing.T) {
	cases := []string{
		"NotApplicable",
		"ColorSlideFilm",
		"ColorNegativeFilm",
		"BlackandWhiteNegativeFilm",
	}

	for _, value := range cases {
		t.Run(value, func(t *testing.T) {
			fsm := FilmScanModeElement(
				ValWithOptions[string]{
					Text: value,
				},
			)
			elm := fsm.toXML("wscn:FilmScanMode")

			if elm.Text != value {
				t.Errorf("expected text '%s', got '%s'", value, elm.Text)
			}

			decoded, err := decodeFilmScanModeElement(elm)
			if err != nil {
				t.Fatalf("decode returned error: %v", err)
			}
			if decoded.Text != value {
				t.Errorf("expected value %s, got %s", value, decoded.Text)
			}
		})
	}
}

func TestFilmScanModeElement_CustomValue(t *testing.T) {
	// Test that custom/extended values are supported
	orig := FilmScanModeElement(
		ValWithOptions[string]{
			Text: "CustomFilmType",
		},
	)

	elm := orig.toXML("wscn:FilmScanMode")
	decoded, err := decodeFilmScanModeElement(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}
	if decoded.Text != "CustomFilmType" {
		t.Errorf("expected custom value 'CustomFilmType', got %s",
			decoded.Text)
	}
}

func TestFilmScanModeElement_WithMustHonor(t *testing.T) {
	orig := FilmScanModeElement(
		ValWithOptions[string]{
			Text:      "ColorNegativeFilm",
			MustHonor: optional.New(BooleanElement("true")),
		},
	)

	elm := orig.toXML("wscn:FilmScanMode")
	decoded, err := decodeFilmScanModeElement(elm)
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

func TestFilmScanModeElement_WithOverride(t *testing.T) {
	orig := FilmScanModeElement(
		ValWithOptions[string]{
			Text:     "BlackandWhiteNegativeFilm",
			Override: optional.New(BooleanElement("0")),
		},
	)

	elm := orig.toXML("wscn:FilmScanMode")
	decoded, err := decodeFilmScanModeElement(elm)
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

func TestFilmScanModeElement_InvalidBooleanAttribute(t *testing.T) {
	elm := xmldoc.Element{
		Name: "wscn:FilmScanMode",
		Text: "ColorSlideFilm",
		Attrs: []xmldoc.Attr{
			{Name: NsWSCN + ":MustHonor", Value: "invalid"},
		},
	}

	_, err := decodeFilmScanModeElement(elm)
	if err == nil {
		t.Error("expected error for invalid boolean attribute, got nil")
	}
}
