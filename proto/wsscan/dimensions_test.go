// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol - unit tests
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions

package wsscan

import (
	"testing"

	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// Test Dimensions XML encoding/decoding
func TestDimensionsXML(t *testing.T) {
	// Create a test Dimensions structure
	dim := Dimensions{
		Width: TextWithOverrideAndDefault{
			Text:        "100",
			Override:    optional.New(BooleanElement("true")),
			UsedDefault: optional.New(BooleanElement("false")),
		},
		Height: TextWithOverrideAndDefault{
			Text:        "200",
			Override:    optional.New(BooleanElement("false")),
			UsedDefault: optional.New(BooleanElement("true")),
		},
	}

	// Convert to XML
	encoded := dim.toXML("Dimensions")

	// Verify structure
	if encoded.Name != "Dimensions" {
		t.Errorf("Expected root element name 'Dimensions', got %q", encoded.Name)
	}

	// Decode back
	decoded, err := decodeDimensions(encoded)
	if err != nil {
		t.Errorf("Failed to decode Dimensions: %v", err)
	}

	// Compare original and decoded
	if decoded.Width.Text != dim.Width.Text {
		t.Errorf("Width text mismatch: got %q, want %q",
			decoded.Width.Text, dim.Width.Text)
	}
	if optional.Get(decoded.Width.Override) != optional.Get(dim.Width.Override) {
		t.Errorf("Width Override mismatch: got %v, want %v",
			optional.Get(decoded.Width.Override), optional.Get(dim.Width.Override))
	}
	if optional.Get(decoded.Width.UsedDefault) != optional.Get(dim.Width.UsedDefault) {
		t.Errorf("Width UsedDefault mismatch: got %v, want %v",
			optional.Get(decoded.Width.UsedDefault), optional.Get(dim.Width.UsedDefault))
	}

	if decoded.Height.Text != dim.Height.Text {
		t.Errorf("Height text mismatch: got %q, want %q",
			decoded.Height.Text, dim.Height.Text)
	}
	if optional.Get(decoded.Height.Override) != optional.Get(dim.Height.Override) {
		t.Errorf("Height Override mismatch: got %v, want %v",
			optional.Get(decoded.Height.Override), optional.Get(dim.Height.Override))
	}
	if optional.Get(decoded.Height.UsedDefault) != optional.Get(dim.Height.UsedDefault) {
		t.Errorf("Height UsedDefault mismatch: got %v, want %v",
			optional.Get(decoded.Height.UsedDefault), optional.Get(dim.Height.UsedDefault))
	}
}

// Test Dimensions with invalid values
func TestDimensionsInvalid(t *testing.T) {
	// Test with invalid width
	dim := Dimensions{
		Width:  TextWithOverrideAndDefault{Text: "invalid"},
		Height: TextWithOverrideAndDefault{Text: "200"},
	}
	encoded := dim.toXML("Dimensions")
	if _, err := decodeDimensions(encoded); err == nil {
		t.Error("Expected error with invalid width")
	}

	// Test with invalid height
	dim = Dimensions{
		Width:  TextWithOverrideAndDefault{Text: "100"},
		Height: TextWithOverrideAndDefault{Text: "invalid"},
	}
	encoded = dim.toXML("Dimensions")
	if _, err := decodeDimensions(encoded); err == nil {
		t.Error("Expected error with invalid height")
	}

	// Test with missing width
	encoded = xmldoc.Element{
		Name: "Dimensions",
		Children: []xmldoc.Element{
			{Name: NsWSCN + ":Height", Text: "200"},
		},
	}
	if _, err := decodeDimensions(encoded); err == nil {
		t.Error("Expected error with missing width")
	}

	// Test with missing height
	encoded = xmldoc.Element{
		Name: "Dimensions",
		Children: []xmldoc.Element{
			{Name: NsWSCN + ":Width", Text: "100"},
		},
	}
	if _, err := decodeDimensions(encoded); err == nil {
		t.Error("Expected error with missing height")
	}
}
