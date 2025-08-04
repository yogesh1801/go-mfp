// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol - unit tests
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions

package wsscan

import (
	"testing"

	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// Test Dimensions XML encoding/decoding
func TestDimensionsXML(t *testing.T) {
	// Create a test Dimensions structure
	dim := Dimensions{
		Width:  100,
		Height: 200,
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
	if decoded.Width != dim.Width {
		t.Errorf("Width mismatch: got %d, want %d",
			decoded.Width, dim.Width)
	}

	if decoded.Height != dim.Height {
		t.Errorf("Height mismatch: got %d, want %d",
			decoded.Height, dim.Height)
	}
}

// Test Dimensions with invalid values
func TestDimensionsInvalid(t *testing.T) {
	// Test with invalid width
	encoded := xmldoc.Element{
		Name: "Dimensions",
		Children: []xmldoc.Element{
			{Name: NsWSCN + ":Width", Text: "invalid"},
			{Name: NsWSCN + ":Height", Text: "200"},
		},
	}
	if _, err := decodeDimensions(encoded); err == nil {
		t.Error("Expected error with invalid width")
	}

	// Test with invalid height
	encoded = xmldoc.Element{
		Name: "Dimensions",
		Children: []xmldoc.Element{
			{Name: NsWSCN + ":Width", Text: "100"},
			{Name: NsWSCN + ":Height", Text: "invalid"},
		},
	}
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
