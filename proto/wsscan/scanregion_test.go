// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// Test for ScanRegion

package wsscan

import (
	"reflect"
	"testing"

	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

func TestScanRegion_RoundTrip_AllAttributes(t *testing.T) {
	orig := ScanRegion{
		ScanRegionHeight: ValWithOptions[int]{
			Text:        1000,
			MustHonor:   optional.New(BooleanElement("true")),
			Override:    optional.New(BooleanElement("false")),
			UsedDefault: optional.New(BooleanElement("1")),
		},
		ScanRegionWidth: ValWithOptions[int]{
			Text:        800,
			MustHonor:   optional.New(BooleanElement("0")),
			Override:    optional.New(BooleanElement("true")),
			UsedDefault: optional.New(BooleanElement("false")),
		},
		ScanRegionXOffset: optional.New(ValWithOptions[int]{
			Text:        100,
			MustHonor:   optional.New(BooleanElement("1")),
			Override:    optional.New(BooleanElement("0")),
			UsedDefault: optional.New(BooleanElement("true")),
		}),
		ScanRegionYOffset: optional.New(ValWithOptions[int]{
			Text:        50,
			MustHonor:   optional.New(BooleanElement("false")),
			Override:    optional.New(BooleanElement("1")),
			UsedDefault: optional.New(BooleanElement("0")),
		}),
	}

	elm := orig.toXML("wscn:ScanRegion")
	if elm.Name != "wscn:ScanRegion" {
		t.Errorf("expected element name 'wscn:ScanRegion', got '%s'", elm.Name)
	}
	if len(elm.Children) != 4 {
		t.Errorf("expected 4 children (with offsets), got %d", len(elm.Children))
	}

	decoded, err := decodeScanRegion(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}
}

func TestScanRegion_RoundTrip_NoAttributes(t *testing.T) {
	orig := ScanRegion{
		ScanRegionHeight:  ValWithOptions[int]{Text: 2000},
		ScanRegionWidth:   ValWithOptions[int]{Text: 1500},
		ScanRegionXOffset: optional.New(ValWithOptions[int]{Text: 0}),
		ScanRegionYOffset: optional.New(ValWithOptions[int]{Text: 0}),
	}

	elm := orig.toXML("wscn:ScanRegion")

	// Verify no attributes on any child
	for _, child := range elm.Children {
		if len(child.Attrs) != 0 {
			t.Errorf("expected no attributes on %s, got %+v",
				child.Name, child.Attrs)
		}
	}

	decoded, err := decodeScanRegion(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}
}

func TestScanRegion_RoundTrip_PartialAttributes(t *testing.T) {
	orig := ScanRegion{
		ScanRegionHeight: ValWithOptions[int]{
			Text:     1200,
			Override: optional.New(BooleanElement("true")),
		},
		ScanRegionWidth: ValWithOptions[int]{
			Text:        850,
			UsedDefault: optional.New(BooleanElement("false")),
		},
		ScanRegionXOffset: optional.New(ValWithOptions[int]{
			Text:      200,
			MustHonor: optional.New(BooleanElement("1")),
		}),
		ScanRegionYOffset: optional.New(ValWithOptions[int]{
			Text: 100,
		}),
	}

	elm := orig.toXML("wscn:ScanRegion")
	decoded, err := decodeScanRegion(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}
}

func TestScanRegion_ChildElementOrder(t *testing.T) {
	orig := ScanRegion{
		ScanRegionHeight:  ValWithOptions[int]{Text: 500},
		ScanRegionWidth:   ValWithOptions[int]{Text: 400},
		ScanRegionXOffset: optional.New(ValWithOptions[int]{Text: 10}),
		ScanRegionYOffset: optional.New(ValWithOptions[int]{Text: 20}),
	}

	elm := orig.toXML("wscn:ScanRegion")

	// Verify child element order
	expectedOrder := []string{
		"wscn:ScanRegionHeight",
		"wscn:ScanRegionWidth",
		"wscn:ScanRegionXOffset",
		"wscn:ScanRegionYOffset",
	}

	if len(elm.Children) != len(expectedOrder) {
		t.Fatalf("expected %d children, got %d", len(expectedOrder), len(elm.Children))
	}

	for i, expected := range expectedOrder {
		if elm.Children[i].Name != expected {
			t.Errorf("child %d: expected name '%s', got '%s'",
				i, expected, elm.Children[i].Name)
		}
	}
}

func TestScanRegion_ChildElementValues(t *testing.T) {
	orig := ScanRegion{
		ScanRegionHeight:  ValWithOptions[int]{Text: 1234},
		ScanRegionWidth:   ValWithOptions[int]{Text: 5678},
		ScanRegionXOffset: optional.New(ValWithOptions[int]{Text: 42}),
		ScanRegionYOffset: optional.New(ValWithOptions[int]{Text: 99}),
	}

	elm := orig.toXML("wscn:ScanRegion")

	// Verify text values
	if elm.Children[0].Text != "1234" {
		t.Errorf("ScanRegionHeight: expected '1234', got '%s'",
			elm.Children[0].Text)
	}
	if elm.Children[1].Text != "5678" {
		t.Errorf("ScanRegionWidth: expected '5678', got '%s'",
			elm.Children[1].Text)
	}
	if elm.Children[2].Text != "42" {
		t.Errorf("ScanRegionXOffset: expected '42', got '%s'",
			elm.Children[2].Text)
	}
	if elm.Children[3].Text != "99" {
		t.Errorf("ScanRegionYOffset: expected '99', got '%s'",
			elm.Children[3].Text)
	}
}

func TestScanRegion_MissingHeight(t *testing.T) {
	elm := xmldoc.Element{
		Name: "wscn:ScanRegion",
		Children: []xmldoc.Element{
			{Name: NsWSCN + ":ScanRegionWidth", Text: "800"},
		},
	}

	_, err := decodeScanRegion(elm)
	if err == nil {
		t.Error("expected error for missing ScanRegionHeight, got nil")
	}
}

func TestScanRegion_MissingWidth(t *testing.T) {
	elm := xmldoc.Element{
		Name: "wscn:ScanRegion",
		Children: []xmldoc.Element{
			{Name: NsWSCN + ":ScanRegionHeight", Text: "1000"},
		},
	}

	_, err := decodeScanRegion(elm)
	if err == nil {
		t.Error("expected error for missing ScanRegionWidth, got nil")
	}
}

func TestScanRegion_MissingOffsets(t *testing.T) {
	// Test that offsets are optional - should decode successfully without them
	orig := ScanRegion{
		ScanRegionHeight: ValWithOptions[int]{Text: 1000},
		ScanRegionWidth:  ValWithOptions[int]{Text: 800},
	}

	elm := orig.toXML("wscn:ScanRegion")
	if len(elm.Children) != 2 {
		t.Errorf("expected 2 children (no offsets), got %d", len(elm.Children))
	}

	decoded, err := decodeScanRegion(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}

	// Verify offsets are nil
	if decoded.ScanRegionXOffset != nil {
		t.Error("expected ScanRegionXOffset to be nil")
	}
	if decoded.ScanRegionYOffset != nil {
		t.Error("expected ScanRegionYOffset to be nil")
	}
}

func TestScanRegion_InvalidHeightValue(t *testing.T) {
	elm := xmldoc.Element{
		Name: "wscn:ScanRegion",
		Children: []xmldoc.Element{
			{Name: NsWSCN + ":ScanRegionHeight", Text: "not-a-number"},
			{Name: NsWSCN + ":ScanRegionWidth", Text: "800"},
		},
	}

	_, err := decodeScanRegion(elm)
	if err == nil {
		t.Error("expected error for invalid ScanRegionHeight value, got nil")
	}
}

func TestScanRegion_InvalidBooleanAttribute(t *testing.T) {
	elm := xmldoc.Element{
		Name: "wscn:ScanRegion",
		Children: []xmldoc.Element{
			{
				Name: NsWSCN + ":ScanRegionHeight",
				Text: "1000",
				Attrs: []xmldoc.Attr{
					{Name: NsWSCN + ":MustHonor", Value: "invalid"},
				},
			},
			{Name: NsWSCN + ":ScanRegionWidth", Text: "800"},
		},
	}

	_, err := decodeScanRegion(elm)
	if err == nil {
		t.Error("expected error for invalid boolean attribute, got nil")
	}
}

func TestScanRegion_BooleanVariations(t *testing.T) {
	cases := []struct {
		name  string
		value BooleanElement
	}{
		{"true", "true"},
		{"false", "false"},
		{"1", "1"},
		{"0", "0"},
		{"TRUE", "TRUE"},
		{"FALSE", "FALSE"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			orig := ScanRegion{
				ScanRegionHeight: ValWithOptions[int]{
					Text:        300,
					MustHonor:   optional.New(c.value),
					Override:    optional.New(c.value),
					UsedDefault: optional.New(c.value),
				},
				ScanRegionWidth: ValWithOptions[int]{
					Text:        250,
					MustHonor:   optional.New(c.value),
					Override:    optional.New(c.value),
					UsedDefault: optional.New(c.value),
				},
				ScanRegionXOffset: optional.New(ValWithOptions[int]{
					Text:        10,
					MustHonor:   optional.New(c.value),
					Override:    optional.New(c.value),
					UsedDefault: optional.New(c.value),
				}),
				ScanRegionYOffset: optional.New(ValWithOptions[int]{
					Text:        20,
					MustHonor:   optional.New(c.value),
					Override:    optional.New(c.value),
					UsedDefault: optional.New(c.value),
				}),
			}

			elm := orig.toXML("wscn:ScanRegion")
			decoded, err := decodeScanRegion(elm)
			if err != nil {
				t.Fatalf("decode returned error: %v", err)
			}
			if !reflect.DeepEqual(orig, decoded) {
				t.Errorf("expected %+v, got %+v", orig, decoded)
			}
		})
	}
}

func TestScanRegion_ZeroOffsets(t *testing.T) {
	orig := ScanRegion{
		ScanRegionHeight:  ValWithOptions[int]{Text: 1000},
		ScanRegionWidth:   ValWithOptions[int]{Text: 800},
		ScanRegionXOffset: optional.New(ValWithOptions[int]{Text: 0}),
		ScanRegionYOffset: optional.New(ValWithOptions[int]{Text: 0}),
	}

	elm := orig.toXML("wscn:ScanRegion")

	// Verify zero values are encoded correctly
	if elm.Children[2].Text != "0" {
		t.Errorf("ScanRegionXOffset: expected '0', got '%s'",
			elm.Children[2].Text)
	}
	if elm.Children[3].Text != "0" {
		t.Errorf("ScanRegionYOffset: expected '0', got '%s'",
			elm.Children[3].Text)
	}

	decoded, err := decodeScanRegion(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}
}
