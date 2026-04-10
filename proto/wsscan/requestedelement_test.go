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
		{"Unknown", UnknownScannerRequestedElement, "Unknown"},
		{"DefaultScanTicket", ScannerRequestedElementDefaultScanTicket,
			NsWSCN + ":DefaultScanTicket"},
		{"Description", ScannerRequestedElementDescription,
			NsWSCN + ":ScannerDescription"},
		{"Configuration", ScannerRequestedElementConfiguration,
			NsWSCN + ":ScannerConfiguration"},
		{"Status", ScannerRequestedElementStatus,
			NsWSCN + ":ScannerStatus"},
		{"VendorSection", ScannerRequestedElementVendorSection,
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
			ScannerRequestedElementDefaultScanTicket},
		{"Description", NsWSCN + ":ScannerDescription",
			ScannerRequestedElementDescription},
		{"Configuration", NsWSCN + ":ScannerConfiguration",
			ScannerRequestedElementConfiguration},
		{"Status", NsWSCN + ":ScannerStatus",
			ScannerRequestedElementStatus},
		{"VendorSection", NsXML + ":VendorSection",
			ScannerRequestedElementVendorSection},
		{"Empty", "", UnknownScannerRequestedElement},
		{"Invalid", "InvalidName", UnknownScannerRequestedElement},
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
			re:      ScannerRequestedElementDefaultScanTicket,
			xmlName: NsWSCN + ":Name",
			expected: xmldoc.Element{
				Name: NsWSCN + ":Name",
				Text: NsWSCN + ":DefaultScanTicket",
			},
		},
		{
			name:    "Description",
			re:      ScannerRequestedElementDescription,
			xmlName: NsWSCN + ":Name",
			expected: xmldoc.Element{
				Name: NsWSCN + ":Name",
				Text: NsWSCN + ":ScannerDescription",
			},
		},
		{
			name:    "Configuration",
			re:      ScannerRequestedElementConfiguration,
			xmlName: NsWSCN + ":Name",
			expected: xmldoc.Element{
				Name: NsWSCN + ":Name",
				Text: NsWSCN + ":ScannerConfiguration",
			},
		},
		{
			name:    "Status",
			re:      ScannerRequestedElementStatus,
			xmlName: NsWSCN + ":Name",
			expected: xmldoc.Element{
				Name: NsWSCN + ":Name",
				Text: NsWSCN + ":ScannerStatus",
			},
		},
		{
			name:    "VendorSection",
			re:      ScannerRequestedElementVendorSection,
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
			expected: ScannerRequestedElementDefaultScanTicket,
		},
		{
			name: "Valid Description",
			element: xmldoc.Element{
				Name: NsWSCN + ":Name",
				Text: NsWSCN + ":ScannerDescription",
			},
			expected: ScannerRequestedElementDescription,
		},
		{
			name: "Valid Configuration",
			element: xmldoc.Element{
				Name: NsWSCN + ":Name",
				Text: NsWSCN + ":ScannerConfiguration",
			},
			expected: ScannerRequestedElementConfiguration,
		},
		{
			name: "Valid Status",
			element: xmldoc.Element{
				Name: NsWSCN + ":Name",
				Text: NsWSCN + ":ScannerStatus",
			},
			expected: ScannerRequestedElementStatus,
		},
		{
			name: "Valid VendorSection",
			element: xmldoc.Element{
				Name: NsWSCN + ":Name",
				Text: NsXML + ":VendorSection",
			},
			expected: ScannerRequestedElementVendorSection,
		},
		{
			name: "Invalid value",
			element: xmldoc.Element{
				Name: NsWSCN + ":Name",
				Text: "InvalidName",
			},
			expected: UnknownScannerRequestedElement,
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
		{"Unknown", UnknownJobRequestedElement, "Unknown"},
		{"JobStatus", JobRequestedElementJobStatus,
			NsWSCN + ":JobStatus"},
		{"ScanTicket", JobRequestedElementScanTicket,
			NsWSCN + ":ScanTicket"},
		{"Documents", JobRequestedElementDocuments,
			NsWSCN + ":Documents"},
		{"VendorSection", JobRequestedElementVendorSection,
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
			JobRequestedElementJobStatus},
		{"ScanTicket", NsWSCN + ":ScanTicket",
			JobRequestedElementScanTicket},
		{"Documents", NsWSCN + ":Documents",
			JobRequestedElementDocuments},
		{"VendorSection", NsXML + ":VendorSection",
			JobRequestedElementVendorSection},
		{"Empty", "", UnknownJobRequestedElement},
		{"Invalid", "InvalidName", UnknownJobRequestedElement},
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
			expected: JobRequestedElementJobStatus,
		},
		{
			name: "Valid ScanTicket",
			element: xmldoc.Element{
				Name: NsWSCN + ":Name",
				Text: NsWSCN + ":ScanTicket",
			},
			expected: JobRequestedElementScanTicket,
		},
		{
			name: "Valid Documents",
			element: xmldoc.Element{
				Name: NsWSCN + ":Name",
				Text: NsWSCN + ":Documents",
			},
			expected: JobRequestedElementDocuments,
		},
		{
			name: "Valid VendorSection",
			element: xmldoc.Element{
				Name: NsWSCN + ":Name",
				Text: NsXML + ":VendorSection",
			},
			expected: JobRequestedElementVendorSection,
		},
		{
			name: "Invalid value",
			element: xmldoc.Element{
				Name: NsWSCN + ":Name",
				Text: "InvalidName",
			},
			expected: UnknownJobRequestedElement,
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
