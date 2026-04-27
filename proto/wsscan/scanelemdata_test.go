// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// Test for ScanElemData

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

// TestScanElemData_RoundTrip_ScannerConfiguration verifies that an ScanElemData
// carrying a ScannerConfiguration child encodes to XML and decodes back
// to an identical value.
func TestScanElemData_RoundTrip_ScannerConfiguration(t *testing.T) {
	orig := ScanElemData{
		Name:  ScanElemDataScannerConfiguration,
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

	parsed, err := decodeScanElemData(elm)
	if err != nil {
		t.Fatalf("decodeScanElemData returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, parsed) {
		t.Errorf("expected %+v, got %+v", orig, parsed)
	}
}

// TestScanElemData_RoundTrip_ScannerDescription verifies that an ScanElemData
// carrying a ScannerDescription child encodes to XML and decodes back
// to an identical value.
func TestScanElemData_RoundTrip_ScannerDescription(t *testing.T) {
	orig := ScanElemData{
		Name:               ScanElemDataScannerDescription,
		Valid:              BooleanElement("true"),
		ScannerDescription: optional.New(createValidScannerDescription()),
	}
	elm := orig.toXML(NsWSCN + ":ElementData")

	parsed, err := decodeScanElemData(elm)
	if err != nil {
		t.Fatalf("decodeScanElemData returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, parsed) {
		t.Errorf("expected %+v, got %+v", orig, parsed)
	}
}

// TestScanElemData_RoundTrip_ScannerStatus verifies that an ScanElemData
// carrying a ScannerStatus child encodes to XML and decodes back
// to an identical value.
func TestScanElemData_RoundTrip_ScannerStatus(t *testing.T) {
	orig := ScanElemData{
		Name:          ScanElemDataScannerStatus,
		Valid:         BooleanElement("true"),
		ScannerStatus: optional.New(createValidScannerStatus()),
	}
	elm := orig.toXML(NsWSCN + ":ElementData")

	parsed, err := decodeScanElemData(elm)
	if err != nil {
		t.Fatalf("decodeScanElemData returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, parsed) {
		t.Errorf("expected %+v, got %+v", orig, parsed)
	}
}

// TestScanElemData_RoundTrip_DefaultScanTicket verifies that an ScanElemData
// carrying a DefaultScanTicket child encodes to XML and decodes back
// to an identical value.
func TestScanElemData_RoundTrip_DefaultScanTicket(t *testing.T) {
	orig := ScanElemData{
		Name:              ScanElemDataDefaultScanTicket,
		Valid:             BooleanElement("true"),
		DefaultScanTicket: optional.New(createValidScanTicket()),
	}
	elm := orig.toXML(NsWSCN + ":ElementData")

	parsed, err := decodeScanElemData(elm)
	if err != nil {
		t.Fatalf("decodeScanElemData returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, parsed) {
		t.Errorf("expected %+v, got %+v", orig, parsed)
	}
}

// TestScanElemData_MissingNameAttr verifies that decoding an ScanElemData
// element without the required Name attribute returns an error.
func TestScanElemData_MissingNameAttr(t *testing.T) {
	orig := ScanElemData{
		Name:  ScanElemDataScannerStatus,
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

	_, err := decodeScanElemData(elm)
	if err == nil {
		t.Error("expected error for missing Name attribute, got nil")
	}
}

// TestScanElemData_InvalidValidAttr verifies that decoding an ScanElemData
// element with an invalid Valid attribute value returns an error.
func TestScanElemData_InvalidValidAttr(t *testing.T) {
	orig := ScanElemData{
		Name:  ScanElemDataScannerStatus,
		Valid: BooleanElement("maybe"),
	}
	elm := orig.toXML(NsWSCN + ":ElementData")

	_, err := decodeScanElemData(elm)
	if err == nil {
		t.Error("expected error for invalid Valid value, got nil")
	}
}

// TestScanElemDataName_String checks that each known ScanElemDataName
// produces the correct local name.
func TestScanElemDataName_String(t *testing.T) {
	tests := []struct {
		name     string
		n        ScanElemDataName
		expected string
	}{
		{"Unknown", UnknownScanElemDataName, "Unknown"},
		{"DefaultScanTicket", ScanElemDataDefaultScanTicket,
			"DefaultScanTicket"},
		{"ScannerConfiguration", ScanElemDataScannerConfiguration,
			"ScannerConfiguration"},
		{"ScannerDescription", ScanElemDataScannerDescription,
			"ScannerDescription"},
		{"ScannerStatus", ScanElemDataScannerStatus, "ScannerStatus"},
		{"VendorSection", ScanElemDataVendorSection, "VendorSection"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.String(); got != tt.expected {
				t.Errorf("String() = %q, want %q", got, tt.expected)
			}
		})
	}
}

// TestScanElemDataName_Encode checks that Encode produces the QName form
// used as element text content in GetScannerElementsRequest.
func TestScanElemDataName_Encode(t *testing.T) {
	if got := ScanElemDataScannerDescription.Encode(); got !=
		NsWSCN+":ScannerDescription" {
		t.Errorf("Encode() = %q, want %q",
			got, NsWSCN+":ScannerDescription")
	}
	if got := ScanElemDataVendorSection.Encode(); got !=
		NsWSCN+":VendorSection" {
		t.Errorf("Encode() = %q, want %q",
			got, NsWSCN+":VendorSection")
	}
}

// TestDecodeScanElemDataName checks that valid QName strings decode to
// the correct constant and invalid ones return Unknown. The namespace
// prefix is intentionally ignored (devices may use a different one for
// the same namespace URL).
func TestDecodeScanElemDataName(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected ScanElemDataName
	}{
		{"DefaultScanTicket", NsWSCN + ":DefaultScanTicket",
			ScanElemDataDefaultScanTicket},
		{"ScannerConfiguration", NsWSCN + ":ScannerConfiguration",
			ScanElemDataScannerConfiguration},
		{"ScannerDescription", NsWSCN + ":ScannerDescription",
			ScanElemDataScannerDescription},
		{"ScannerStatus", NsWSCN + ":ScannerStatus",
			ScanElemDataScannerStatus},
		{"VendorSection different prefix",
			"vendor:VendorSection", ScanElemDataVendorSection},
		{"Empty", "", UnknownScanElemDataName},
		{"Invalid", "InvalidName", UnknownScanElemDataName},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecodeScanElemDataName(tt.input); got !=
				tt.expected {
				t.Errorf("DecodeScanElemDataName(%q) = %v, want %v",
					tt.input, got, tt.expected)
			}
		})
	}
}

// TestScanElemDataName_toXML_RoundTrip checks that toXML and
// decodeScanElemDataName round-trip the Name-as-text-element form used
// by GetScannerElementsRequest.
func TestScanElemDataName_toXML_RoundTrip(t *testing.T) {
	values := []ScanElemDataName{
		ScanElemDataDefaultScanTicket,
		ScanElemDataScannerDescription,
		ScanElemDataScannerConfiguration,
		ScanElemDataScannerStatus,
		ScanElemDataVendorSection,
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
			got, err := decodeScanElemDataName(elm)
			if err != nil {
				t.Fatalf("decodeScanElemDataName: %v", err)
			}
			if got != v {
				t.Errorf("round-trip = %v, want %v", got, v)
			}
		})
	}
}

// Test_decodeScanElemDataName_Invalid verifies that an XML element with
// an unrecognised Name value returns an error.
func Test_decodeScanElemDataName_Invalid(t *testing.T) {
	elm := xmldoc.Element{
		Name: NsWSCN + ":Name",
		Text: "InvalidName",
	}
	if _, err := decodeScanElemDataName(elm); err == nil {
		t.Error("expected error for invalid Name text, got nil")
	}
}

// TestScanElemData_UnknownName verifies that decoding an ScanElemData element
// with an unrecognised Name attribute value returns an error.
func TestScanElemData_UnknownName(t *testing.T) {
	orig := ScanElemData{
		Name:  ScanElemDataScannerStatus,
		Valid: BooleanElement("true"),
	}
	elm := orig.toXML(NsWSCN + ":ElementData")
	for i, a := range elm.Attrs {
		if a.Name == "Name" {
			elm.Attrs[i].Value = "wscn:UnknownElement"
		}
	}

	_, err := decodeScanElemData(elm)
	if err == nil {
		t.Error("expected error for unknown Name value, got nil")
	}
}
