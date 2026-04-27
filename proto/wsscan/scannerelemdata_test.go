// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// Test for ScannerElemData

package wsscan

import (
	"reflect"
	"testing"
	"time"

	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

func createValidScannerDescription() ScannerDescription {
	return ScannerDescription{
		ScannerName: TextWithLangList{
			{
				Text: "Test Scanner",
				Lang: optional.New("en-US"),
			},
		},
	}
}

func createValidScannerStatus() ScannerStatus {
	t, _ := time.Parse(time.RFC3339, "2024-01-01T12:00:00Z")
	return ScannerStatus{
		ActiveConditions: []DeviceCondition{
			{
				Component: PlatenComponent,
				Name:      CoverOpen,
				Severity:  Warning,
				Time:      t,
			},
		},
		ScannerCurrentTime:  t,
		ScannerState:        Idle,
		ScannerStateReasons: []ScannerStateReason{StateNone},
	}
}

func createValidScanTicket() ScanTicket {
	return ScanTicket{
		JobDescription: JobDescription{
			JobName:                "TestJob",
			JobOriginatingUserName: "user",
		},
	}
}

// TestScannerElemData_RoundTrip_ScannerConfiguration verifies that an ScannerElemData
// carrying a ScannerConfiguration child encodes to XML and decodes back
// to an identical value.
func TestScannerElemData_RoundTrip_ScannerConfiguration(t *testing.T) {
	orig := ScannerElemData{
		Name:  ScannerElemConfiguration,
		Valid: BooleanElement("true"),
		ScannerConfiguration: optional.New(ScannerConfiguration{
			DeviceSettings: createValidDeviceSettings(),
		}),
	}
	elm := orig.toXML(NsWSCN + ":ElementData")
	if elm.Name != NsWSCN+":ElementData" {
		t.Errorf("expected element name %q, got %q",
			NsWSCN+":ElementData", elm.Name)
	}

	parsed, err := decodeScannerElemData(elm)
	if err != nil {
		t.Fatalf("decodeScannerElemData returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, parsed) {
		t.Errorf("expected %+v, got %+v", orig, parsed)
	}
}

// TestScannerElemData_RoundTrip_ScannerDescription verifies that an ScannerElemData
// carrying a ScannerDescription child encodes to XML and decodes back
// to an identical value.
func TestScannerElemData_RoundTrip_ScannerDescription(t *testing.T) {
	orig := ScannerElemData{
		Name:               ScannerElemDescription,
		Valid:              BooleanElement("true"),
		ScannerDescription: optional.New(createValidScannerDescription()),
	}
	elm := orig.toXML(NsWSCN + ":ElementData")

	parsed, err := decodeScannerElemData(elm)
	if err != nil {
		t.Fatalf("decodeScannerElemData returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, parsed) {
		t.Errorf("expected %+v, got %+v", orig, parsed)
	}
}

// TestScannerElemData_RoundTrip_ScannerStatus verifies that an ScannerElemData
// carrying a ScannerStatus child encodes to XML and decodes back
// to an identical value.
func TestScannerElemData_RoundTrip_ScannerStatus(t *testing.T) {
	orig := ScannerElemData{
		Name:          ScannerElemStatus,
		Valid:         BooleanElement("true"),
		ScannerStatus: optional.New(createValidScannerStatus()),
	}
	elm := orig.toXML(NsWSCN + ":ElementData")

	parsed, err := decodeScannerElemData(elm)
	if err != nil {
		t.Fatalf("decodeScannerElemData returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, parsed) {
		t.Errorf("expected %+v, got %+v", orig, parsed)
	}
}

// TestScannerElemData_RoundTrip_DefaultScanTicket verifies that an ScannerElemData
// carrying a DefaultScanTicket child encodes to XML and decodes back
// to an identical value.
func TestScannerElemData_RoundTrip_DefaultScanTicket(t *testing.T) {
	orig := ScannerElemData{
		Name:              ScannerElemDefaultScanTicket,
		Valid:             BooleanElement("true"),
		DefaultScanTicket: optional.New(createValidScanTicket()),
	}
	elm := orig.toXML(NsWSCN + ":ElementData")

	parsed, err := decodeScannerElemData(elm)
	if err != nil {
		t.Fatalf("decodeScannerElemData returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, parsed) {
		t.Errorf("expected %+v, got %+v", orig, parsed)
	}
}

// TestScannerElemData_MissingNameAttr verifies that decoding an ScannerElemData
// element without the required Name attribute returns an error.
func TestScannerElemData_MissingNameAttr(t *testing.T) {
	orig := ScannerElemData{
		Name:  ScannerElemStatus,
		Valid: BooleanElement("true"),
	}
	elm := orig.toXML(NsWSCN + ":ElementData")
	var attrs []xmldoc.Attr
	for _, a := range elm.Attrs {
		if a.Name != "Name" {
			attrs = append(attrs, a)
		}
	}
	elm.Attrs = attrs

	_, err := decodeScannerElemData(elm)
	if err == nil {
		t.Error("expected error for missing Name attribute, got nil")
	}
}

// TestScannerElemData_InvalidValidAttr verifies that decoding an ScannerElemData
// element with an invalid Valid attribute value returns an error.
func TestScannerElemData_InvalidValidAttr(t *testing.T) {
	orig := ScannerElemData{
		Name:  ScannerElemStatus,
		Valid: BooleanElement("maybe"),
	}
	elm := orig.toXML(NsWSCN + ":ElementData")

	_, err := decodeScannerElemData(elm)
	if err == nil {
		t.Error("expected error for invalid Valid value, got nil")
	}
}

// TestScannerElemName_String checks that each known ScannerElemName
// produces the correct local name.
func TestScannerElemName_String(t *testing.T) {
	tests := []struct {
		name     string
		n        ScannerElemName
		expected string
	}{
		{"Unknown", UnknownScannerElem, "Unknown"},
		{"DefaultScanTicket", ScannerElemDefaultScanTicket,
			"DefaultScanTicket"},
		{"ScannerConfiguration", ScannerElemConfiguration,
			"ScannerConfiguration"},
		{"ScannerDescription", ScannerElemDescription,
			"ScannerDescription"},
		{"ScannerStatus", ScannerElemStatus, "ScannerStatus"},
		{"VendorSection", ScannerElemVendorSection, "VendorSection"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.String(); got != tt.expected {
				t.Errorf("String() = %q, want %q", got, tt.expected)
			}
		})
	}
}

// TestScannerElemName_Encode checks that Encode produces the QName form
// used as element text content in GetScannerElementsRequest.
func TestScannerElemName_Encode(t *testing.T) {
	if got := ScannerElemDescription.Encode(); got !=
		NsWSCN+":ScannerDescription" {
		t.Errorf("Encode() = %q, want %q",
			got, NsWSCN+":ScannerDescription")
	}
	if got := ScannerElemVendorSection.Encode(); got !=
		NsWSCN+":VendorSection" {
		t.Errorf("Encode() = %q, want %q",
			got, NsWSCN+":VendorSection")
	}
}

// TestDecodeScannerElemName checks that valid QName strings decode to
// the correct constant and invalid ones return Unknown. The namespace
// prefix is intentionally ignored (devices may use a different one for
// the same namespace URL).
func TestDecodeScannerElemName(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected ScannerElemName
	}{
		{"DefaultScanTicket", NsWSCN + ":DefaultScanTicket",
			ScannerElemDefaultScanTicket},
		{"ScannerConfiguration", NsWSCN + ":ScannerConfiguration",
			ScannerElemConfiguration},
		{"ScannerDescription", NsWSCN + ":ScannerDescription",
			ScannerElemDescription},
		{"ScannerStatus", NsWSCN + ":ScannerStatus",
			ScannerElemStatus},
		{"VendorSection different prefix",
			"vendor:VendorSection", ScannerElemVendorSection},
		{"Empty", "", UnknownScannerElem},
		{"Invalid", "InvalidName", UnknownScannerElem},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecodeScannerElemName(tt.input); got !=
				tt.expected {
				t.Errorf("DecodeScannerElemName(%q) = %v, want %v",
					tt.input, got, tt.expected)
			}
		})
	}
}

// TestScannerElemName_toXML_RoundTrip checks that toXML and
// decodeScannerElemName round-trip the Name-as-text-element form used
// by GetScannerElementsRequest.
func TestScannerElemName_toXML_RoundTrip(t *testing.T) {
	values := []ScannerElemName{
		ScannerElemDefaultScanTicket,
		ScannerElemDescription,
		ScannerElemConfiguration,
		ScannerElemStatus,
		ScannerElemVendorSection,
	}
	for _, v := range values {
		t.Run(v.String(), func(t *testing.T) {
			elm := v.toXML(NsWSCN + ":Name")
			if elm.Name != NsWSCN+":Name" {
				t.Errorf("toXML().Name = %q, want %q",
					elm.Name, NsWSCN+":Name")
			}
			if elm.Text != v.Encode() {
				t.Errorf("toXML().Text = %q, want %q",
					elm.Text, v.Encode())
			}
			got, err := decodeScannerElemName(elm)
			if err != nil {
				t.Fatalf("decodeScannerElemName: %v", err)
			}
			if got != v {
				t.Errorf("round-trip = %v, want %v", got, v)
			}
		})
	}
}

// Test_decodeScannerElemName_Invalid verifies that an XML element with
// an unrecognised Name value returns an error.
func Test_decodeScannerElemName_Invalid(t *testing.T) {
	elm := xmldoc.Element{
		Name: NsWSCN + ":Name",
		Text: "InvalidName",
	}
	if _, err := decodeScannerElemName(elm); err == nil {
		t.Error("expected error for invalid Name text, got nil")
	}
}

// TestScannerElemData_UnknownName verifies that decoding an ScannerElemData element
// with an unrecognised Name attribute value returns an error.
func TestScannerElemData_UnknownName(t *testing.T) {
	orig := ScannerElemData{
		Name:  ScannerElemStatus,
		Valid: BooleanElement("true"),
	}
	elm := orig.toXML(NsWSCN + ":ElementData")
	for i, a := range elm.Attrs {
		if a.Name == "Name" {
			elm.Attrs[i].Value = "wscn:UnknownElement"
		}
	}

	_, err := decodeScannerElemData(elm)
	if err == nil {
		t.Error("expected error for unknown Name value, got nil")
	}
}
