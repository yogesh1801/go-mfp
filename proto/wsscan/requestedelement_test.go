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

// TestScannerRequestedElement_String checks that each known
// ScannerRequestedElement produces the correct QName string.
func TestScannerRequestedElement_String(t *testing.T) {
	tests := []struct {
		name     string
		re       ScannerRequestedElement
		expected string
	}{
		{"Unknown", UnknownScannerElem, "Unknown"},
		{"DefaultScanTicket", ScannerElemDefaultScanTicket,
			NsWSCN + ":DefaultScanTicket"},
		{"Description", ScannerElemDescription,
			NsWSCN + ":ScannerDescription"},
		{"Configuration", ScannerElemConfiguration,
			NsWSCN + ":ScannerConfiguration"},
		{"Status", ScannerElemStatus,
			NsWSCN + ":ScannerStatus"},
		{"VendorSection", ScannerElemVendorSection,
			NsXML + ":VendorSection"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.re.String(); got != tt.expected {
				t.Errorf("ScannerRequestedElement.String() = %v, want %v",
					got, tt.expected)
			}
		})
	}
}

// TestDecodeScannerRequestedElement checks that valid QName strings decode to
// the correct constant and invalid ones return Unknown.
func TestDecodeScannerRequestedElement(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected ScannerRequestedElement
	}{
		{"DefaultScanTicket", NsWSCN + ":DefaultScanTicket",
			ScannerElemDefaultScanTicket},
		{"Description", NsWSCN + ":ScannerDescription",
			ScannerElemDescription},
		{"Configuration", NsWSCN + ":ScannerConfiguration",
			ScannerElemConfiguration},
		{"Status", NsWSCN + ":ScannerStatus",
			ScannerElemStatus},
		{"VendorSection", NsXML + ":VendorSection",
			ScannerElemVendorSection},
		{"Empty", "", UnknownScannerElem},
		{"Invalid", "InvalidName", UnknownScannerElem},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecodeScannerRequestedElement(tt.input); got != tt.expected {
				t.Errorf("DecodeScannerRequestedElement() = %v, want %v",
					got, tt.expected)
			}
		})
	}
}

// TestScannerRequestedElement_toXML checks that toXML produces an element with
// the correct name and QName text value.
func TestScannerRequestedElement_toXML(t *testing.T) {
	tests := []struct {
		name     string
		re       ScannerRequestedElement
		xmlName  string
		expected xmldoc.Element
	}{
		{
			name:    "DefaultScanTicket",
			re:      ScannerElemDefaultScanTicket,
			xmlName: NsWSCN + ":Name",
			expected: xmldoc.Element{
				Name: NsWSCN + ":Name",
				Text: NsWSCN + ":DefaultScanTicket",
			},
		},
		{
			name:    "Description",
			re:      ScannerElemDescription,
			xmlName: NsWSCN + ":Name",
			expected: xmldoc.Element{
				Name: NsWSCN + ":Name",
				Text: NsWSCN + ":ScannerDescription",
			},
		},
		{
			name:    "Configuration",
			re:      ScannerElemConfiguration,
			xmlName: NsWSCN + ":Name",
			expected: xmldoc.Element{
				Name: NsWSCN + ":Name",
				Text: NsWSCN + ":ScannerConfiguration",
			},
		},
		{
			name:    "Status",
			re:      ScannerElemStatus,
			xmlName: NsWSCN + ":Name",
			expected: xmldoc.Element{
				Name: NsWSCN + ":Name",
				Text: NsWSCN + ":ScannerStatus",
			},
		},
		{
			name:    "VendorSection",
			re:      ScannerElemVendorSection,
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

// Test_decodeScannerRequestedElement checks that decoding from XML elements
// works for valid values and returns an error for invalid ones.
func Test_decodeScannerRequestedElement(t *testing.T) {
	tests := []struct {
		name     string
		element  xmldoc.Element
		expected ScannerRequestedElement
		wantErr  bool
	}{
		{
			name: "Valid DefaultScanTicket",
			element: xmldoc.Element{
				Name: NsWSCN + ":Name",
				Text: NsWSCN + ":DefaultScanTicket",
			},
			expected: ScannerElemDefaultScanTicket,
		},
		{
			name: "Valid Description",
			element: xmldoc.Element{
				Name: NsWSCN + ":Name",
				Text: NsWSCN + ":ScannerDescription",
			},
			expected: ScannerElemDescription,
		},
		{
			name: "Valid Configuration",
			element: xmldoc.Element{
				Name: NsWSCN + ":Name",
				Text: NsWSCN + ":ScannerConfiguration",
			},
			expected: ScannerElemConfiguration,
		},
		{
			name: "Valid Status",
			element: xmldoc.Element{
				Name: NsWSCN + ":Name",
				Text: NsWSCN + ":ScannerStatus",
			},
			expected: ScannerElemStatus,
		},
		{
			name: "Valid VendorSection",
			element: xmldoc.Element{
				Name: NsWSCN + ":Name",
				Text: NsXML + ":VendorSection",
			},
			expected: ScannerElemVendorSection,
		},
		{
			name: "Invalid value",
			element: xmldoc.Element{
				Name: NsWSCN + ":Name",
				Text: "InvalidName",
			},
			expected: UnknownScannerElem,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := decodeScannerRequestedElement(tt.element)
			if (err != nil) != tt.wantErr {
				t.Errorf("decodeScannerRequestedElement() error = %v, wantErr %v",
					err, tt.wantErr)
				return
			}
			if got != tt.expected {
				t.Errorf("decodeScannerRequestedElement() = %v, want %v",
					got, tt.expected)
			}
		})
	}
}

// TestJobRequestedElement_String checks that each known JobRequestedElement
// produces the correct QName string.
func TestJobRequestedElement_String(t *testing.T) {
	tests := []struct {
		name     string
		re       JobRequestedElement
		expected string
	}{
		{"Unknown", UnknownJobElem, "Unknown"},
		{"JobStatus", JobElemStatus,
			NsWSCN + ":JobStatus"},
		{"ScanTicket", JobElemScanTicket,
			NsWSCN + ":ScanTicket"},
		{"Documents", JobElemDocuments,
			NsWSCN + ":Documents"},
		{"VendorSection", JobElemVendorSection,
			NsXML + ":VendorSection"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.re.String(); got != tt.expected {
				t.Errorf("JobRequestedElement.String() = %v, want %v",
					got, tt.expected)
			}
		})
	}
}

// TestDecodeJobRequestedElement checks that valid QName strings decode to the
// correct constant and invalid ones return Unknown.
func TestDecodeJobRequestedElement(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected JobRequestedElement
	}{
		{"JobStatus", NsWSCN + ":JobStatus",
			JobElemStatus},
		{"ScanTicket", NsWSCN + ":ScanTicket",
			JobElemScanTicket},
		{"Documents", NsWSCN + ":Documents",
			JobElemDocuments},
		{"VendorSection", NsXML + ":VendorSection",
			JobElemVendorSection},
		{"Empty", "", UnknownJobElem},
		{"Invalid", "InvalidName", UnknownJobElem},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecodeJobRequestedElement(tt.input); got != tt.expected {
				t.Errorf("DecodeJobRequestedElement() = %v, want %v",
					got, tt.expected)
			}
		})
	}
}

// Test_decodeJobRequestedElement checks that decoding from XML elements works
// for valid values and returns an error for invalid ones.
func Test_decodeJobRequestedElement(t *testing.T) {
	tests := []struct {
		name     string
		element  xmldoc.Element
		expected JobRequestedElement
		wantErr  bool
	}{
		{
			name: "Valid JobStatus",
			element: xmldoc.Element{
				Name: NsWSCN + ":Name",
				Text: NsWSCN + ":JobStatus",
			},
			expected: JobElemStatus,
		},
		{
			name: "Valid ScanTicket",
			element: xmldoc.Element{
				Name: NsWSCN + ":Name",
				Text: NsWSCN + ":ScanTicket",
			},
			expected: JobElemScanTicket,
		},
		{
			name: "Valid Documents",
			element: xmldoc.Element{
				Name: NsWSCN + ":Name",
				Text: NsWSCN + ":Documents",
			},
			expected: JobElemDocuments,
		},
		{
			name: "Valid VendorSection",
			element: xmldoc.Element{
				Name: NsWSCN + ":Name",
				Text: NsXML + ":VendorSection",
			},
			expected: JobElemVendorSection,
		},
		{
			name: "Invalid value",
			element: xmldoc.Element{
				Name: NsWSCN + ":Name",
				Text: "InvalidName",
			},
			expected: UnknownJobElem,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := decodeJobRequestedElement(tt.element)
			if (err != nil) != tt.wantErr {
				t.Errorf("decodeJobRequestedElement() error = %v, wantErr %v",
					err, tt.wantErr)
				return
			}
			if got != tt.expected {
				t.Errorf("decodeJobRequestedElement() = %v, want %v",
					got, tt.expected)
			}
		})
	}
}
