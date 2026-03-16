// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// Name for RequestedElements element tests

package wsscan

import (
	"testing"

	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// TestRequestedElement_String checks that each known RequestedElement
// produces the correct QName string.
func TestRequestedElement_String(t *testing.T) {
	tests := []struct {
		name     string
		re       RequestedElement
		expected string
	}{
		{"Unknown", UnknownRequestedElement, "Unknown"},
		{"Description", RequestedElementDescription,
			NsWSCN + ":ScannerDescription"},
		{"Configuration", RequestedElementConfiguration,
			NsWSCN + ":ScannerConfiguration"},
		{"Status", RequestedElementStatus,
			NsWSCN + ":ScannerStatus"},
		{"VendorSection", RequestedElementVendorSection,
			NsXML + ":VendorSection"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.re.String(); got != tt.expected {
				t.Errorf("RequestedElement.String() = %v, want %v",
					got, tt.expected)
			}
		})
	}
}

// TestDecodeRequestedElement checks that valid QName strings decode to the
// correct constant and invalid ones return Unknown.
func TestDecodeRequestedElement(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected RequestedElement
	}{
		{"Description", NsWSCN + ":ScannerDescription",
			RequestedElementDescription},
		{"Configuration", NsWSCN + ":ScannerConfiguration",
			RequestedElementConfiguration},
		{"Status", NsWSCN + ":ScannerStatus",
			RequestedElementStatus},
		{"VendorSection", NsXML + ":VendorSection",
			RequestedElementVendorSection},
		{"Empty", "", UnknownRequestedElement},
		{"Invalid", "InvalidName", UnknownRequestedElement},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecodeRequestedElement(tt.input); got != tt.expected {
				t.Errorf("DecodeRequestedElement() = %v, want %v",
					got, tt.expected)
			}
		})
	}
}

// TestRequestedElement_toXML checks that toXML produces an element with
// the correct name and QName text value.
func TestRequestedElement_toXML(t *testing.T) {
	tests := []struct {
		name     string
		re       RequestedElement
		xmlName  string
		expected xmldoc.Element
	}{
		{
			// Verify Description produces wscn:ScannerDescription text.
			name:    "Description",
			re:      RequestedElementDescription,
			xmlName: NsWSCN + ":Name",
			expected: xmldoc.Element{
				Name: NsWSCN + ":Name",
				Text: NsWSCN + ":ScannerDescription",
			},
		},
		{
			// Verify Configuration produces wscn:ScannerConfiguration text.
			name:    "Configuration",
			re:      RequestedElementConfiguration,
			xmlName: NsWSCN + ":Name",
			expected: xmldoc.Element{
				Name: NsWSCN + ":Name",
				Text: NsWSCN + ":ScannerConfiguration",
			},
		},
		{
			// Verify Status produces wscn:ScannerStatus text.
			name:    "Status",
			re:      RequestedElementStatus,
			xmlName: NsWSCN + ":Name",
			expected: xmldoc.Element{
				Name: NsWSCN + ":Name",
				Text: NsWSCN + ":ScannerStatus",
			},
		},
		{
			// Verify VendorSection produces xmlns:VendorSection text.
			name:    "VendorSection",
			re:      RequestedElementVendorSection,
			xmlName: NsWSCN + ":Name",
			expected: xmldoc.Element{
				Name: NsWSCN + ":Name",
				Text: NsXML + ":VendorSection",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.re.toXML(tt.xmlName)
			if got.Name != tt.expected.Name {
				t.Errorf("toXML().Name = %v, want %v",
					got.Name, tt.expected.Name)
			}
			if got.Text != tt.expected.Text {
				t.Errorf("toXML().Text = %v, want %v",
					got.Text, tt.expected.Text)
			}
		})
	}
}

// Test_decodeRequestedElement checks that decoding from XML elements works
// for valid values and returns an error for invalid ones.
func Test_decodeRequestedElement(t *testing.T) {
	tests := []struct {
		name     string
		element  xmldoc.Element
		expected RequestedElement
		wantErr  bool
	}{
		{
			// Valid wscn:ScannerDescription.
			name: "Valid Description",
			element: xmldoc.Element{
				Name: NsWSCN + ":Name",
				Text: NsWSCN + ":ScannerDescription",
			},
			expected: RequestedElementDescription,
			wantErr:  false,
		},
		{
			// Valid wscn:ScannerConfiguration.
			name: "Valid Configuration",
			element: xmldoc.Element{
				Name: NsWSCN + ":Name",
				Text: NsWSCN + ":ScannerConfiguration",
			},
			expected: RequestedElementConfiguration,
			wantErr:  false,
		},
		{
			// Valid wscn:ScannerStatus.
			name: "Valid Status",
			element: xmldoc.Element{
				Name: NsWSCN + ":Name",
				Text: NsWSCN + ":ScannerStatus",
			},
			expected: RequestedElementStatus,
			wantErr:  false,
		},
		{
			// Valid xmlns:VendorSection.
			name: "Valid VendorSection",
			element: xmldoc.Element{
				Name: NsWSCN + ":Name",
				Text: NsXML + ":VendorSection",
			},
			expected: RequestedElementVendorSection,
			wantErr:  false,
		},
		{
			// Invalid value must return an error.
			name: "Invalid value",
			element: xmldoc.Element{
				Name: NsWSCN + ":Name",
				Text: "InvalidName",
			},
			expected: UnknownRequestedElement,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := decodeRequestedElement(tt.element)
			if (err != nil) != tt.wantErr {
				t.Errorf("decodeRequestedElement() error = %v, wantErr %v",
					err, tt.wantErr)
				return
			}
			if got != tt.expected {
				t.Errorf("decodeRequestedElement() = %v, want %v",
					got, tt.expected)
			}
		})
	}
}
