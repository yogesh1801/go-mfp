// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// Test for JobElemData

package wsscan

import (
	"reflect"
	"testing"
	"time"

	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

func createValidJobStatus() JobStatus {
	created, _ := time.Parse(time.RFC3339, "2024-01-01T12:00:00Z")
	return JobStatus{
		JobCreatedTime: optional.New(created),
		JobID:          42,
		JobState:       JobStateProcessing,
		ScansCompleted: 1,
	}
}

func createValidDocuments() Documents {
	return Documents{
		DocumentFinalParameters: DocumentParameters{},
	}
}

// TestJobElemData_RoundTrip_JobStatus verifies that a JobElemData
// carrying a JobStatus child encodes to XML and decodes back to an
// identical value.
func TestJobElemData_RoundTrip_JobStatus(t *testing.T) {
	orig := JobElemData{
		Name:      JobElemDataJobStatus,
		Valid:     BooleanElement("true"),
		JobStatus: optional.New(createValidJobStatus()),
	}
	elm := orig.toXML(NsWSCN + ":ElementData")
	if elm.Name != NsWSCN+":ElementData" {
		t.Errorf("expected element name %q, got %q",
			NsWSCN+":ElementData", elm.Name)
	}

	parsed, err := decodeJobElemData(elm)
	if err != nil {
		t.Fatalf("decodeJobElemData returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, parsed) {
		t.Errorf("expected %+v, got %+v", orig, parsed)
	}
}

// TestJobElemData_RoundTrip_ScanTicket verifies that a JobElemData
// carrying a ScanTicket child encodes to XML and decodes back to an
// identical value.
func TestJobElemData_RoundTrip_ScanTicket(t *testing.T) {
	orig := JobElemData{
		Name:       JobElemDataScanTicket,
		Valid:      BooleanElement("true"),
		ScanTicket: optional.New(createValidScanTicket()),
	}
	elm := orig.toXML(NsWSCN + ":ElementData")

	parsed, err := decodeJobElemData(elm)
	if err != nil {
		t.Fatalf("decodeJobElemData returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, parsed) {
		t.Errorf("expected %+v, got %+v", orig, parsed)
	}
}

// TestJobElemData_RoundTrip_Documents verifies that a JobElemData
// carrying a Documents child encodes to XML and decodes back to an
// identical value.
func TestJobElemData_RoundTrip_Documents(t *testing.T) {
	orig := JobElemData{
		Name:      JobElemDataDocuments,
		Valid:     BooleanElement("true"),
		Documents: optional.New(createValidDocuments()),
	}
	elm := orig.toXML(NsWSCN + ":ElementData")

	parsed, err := decodeJobElemData(elm)
	if err != nil {
		t.Fatalf("decodeJobElemData returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, parsed) {
		t.Errorf("expected %+v, got %+v", orig, parsed)
	}
}

// TestJobElemData_MissingNameAttr verifies that decoding a JobElemData
// element without the required Name attribute returns an error.
func TestJobElemData_MissingNameAttr(t *testing.T) {
	orig := JobElemData{
		Name:      JobElemDataJobStatus,
		Valid:     BooleanElement("true"),
		JobStatus: optional.New(createValidJobStatus()),
	}
	elm := orig.toXML(NsWSCN + ":ElementData")
	var attrs []xmldoc.Attr
	for _, a := range elm.Attrs {
		if a.Name != "Name" {
			attrs = append(attrs, a)
		}
	}
	elm.Attrs = attrs

	_, err := decodeJobElemData(elm)
	if err == nil {
		t.Error("expected error for missing Name attribute, got nil")
	}
}

// TestJobElemData_InvalidValidAttr verifies that decoding a JobElemData
// element with an invalid Valid attribute value returns an error.
func TestJobElemData_InvalidValidAttr(t *testing.T) {
	orig := JobElemData{
		Name:  JobElemDataJobStatus,
		Valid: BooleanElement("maybe"),
	}
	elm := orig.toXML(NsWSCN + ":ElementData")

	_, err := decodeJobElemData(elm)
	if err == nil {
		t.Error("expected error for invalid Valid value, got nil")
	}
}

// TestJobElemData_UnknownName verifies that decoding a JobElemData
// element with an unrecognised Name attribute value returns an error.
func TestJobElemData_UnknownName(t *testing.T) {
	orig := JobElemData{
		Name:  JobElemDataJobStatus,
		Valid: BooleanElement("true"),
	}
	elm := orig.toXML(NsWSCN + ":ElementData")
	for i, a := range elm.Attrs {
		if a.Name == "Name" {
			elm.Attrs[i].Value = "wscn:UnknownElement"
		}
	}

	_, err := decodeJobElemData(elm)
	if err == nil {
		t.Error("expected error for unknown Name value, got nil")
	}
}

// TestJobElemDataName_String checks that each known JobElemDataName
// produces the correct local name.
func TestJobElemDataName_String(t *testing.T) {
	tests := []struct {
		name     string
		n        JobElemDataName
		expected string
	}{
		{"Unknown", UnknownJobElemDataName, "Unknown"},
		{"JobStatus", JobElemDataJobStatus, "JobStatus"},
		{"ScanTicket", JobElemDataScanTicket, "ScanTicket"},
		{"Documents", JobElemDataDocuments, "Documents"},
		{"VendorSection", JobElemDataVendorSection, "VendorSection"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.String(); got != tt.expected {
				t.Errorf("String() = %q, want %q", got, tt.expected)
			}
		})
	}
}

// TestJobElemDataName_Encode checks that Encode produces the QName form
// used as element text content in GetJobElementsRequest.
func TestJobElemDataName_Encode(t *testing.T) {
	if got := JobElemDataJobStatus.Encode(); got != NsWSCN+":JobStatus" {
		t.Errorf("Encode() = %q, want %q",
			got, NsWSCN+":JobStatus")
	}
	if got := JobElemDataVendorSection.Encode(); got !=
		NsWSCN+":VendorSection" {
		t.Errorf("Encode() = %q, want %q",
			got, NsWSCN+":VendorSection")
	}
}

// TestDecodeJobElemDataName checks that valid QName strings decode to the
// correct constant and invalid ones return Unknown. The namespace prefix
// is intentionally ignored.
func TestDecodeJobElemDataName(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected JobElemDataName
	}{
		{"JobStatus", NsWSCN + ":JobStatus", JobElemDataJobStatus},
		{"ScanTicket", NsWSCN + ":ScanTicket", JobElemDataScanTicket},
		{"Documents", NsWSCN + ":Documents", JobElemDataDocuments},
		{"VendorSection different prefix",
			"vendor:VendorSection", JobElemDataVendorSection},
		{"Empty", "", UnknownJobElemDataName},
		{"Invalid", "InvalidName", UnknownJobElemDataName},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecodeJobElemDataName(tt.input); got !=
				tt.expected {
				t.Errorf("DecodeJobElemDataName(%q) = %v, want %v",
					tt.input, got, tt.expected)
			}
		})
	}
}

// TestJobElemDataName_toXML_RoundTrip checks that toXML and
// decodeJobElemDataName round-trip the Name-as-text-element form used by
// GetJobElementsRequest.
func TestJobElemDataName_toXML_RoundTrip(t *testing.T) {
	values := []JobElemDataName{
		JobElemDataJobStatus,
		JobElemDataScanTicket,
		JobElemDataDocuments,
		JobElemDataVendorSection,
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
			got, err := decodeJobElemDataName(elm)
			if err != nil {
				t.Fatalf("decodeJobElemDataName: %v", err)
			}
			if got != v {
				t.Errorf("round-trip = %v, want %v", got, v)
			}
		})
	}
}

// Test_decodeJobElemDataName_Invalid verifies that an XML element with
// an unrecognised Name value returns an error.
func Test_decodeJobElemDataName_Invalid(t *testing.T) {
	elm := xmldoc.Element{
		Name: NsWSCN + ":Name",
		Text: "InvalidName",
	}
	if _, err := decodeJobElemDataName(elm); err == nil {
		t.Error("expected error for invalid Name text, got nil")
	}
}

// TestJobElemData_RejectsScannerName verifies that a Name value that
// belongs to ScanElemData (e.g. ScannerDescription) is rejected when
// decoded as a JobElemData — enforcing the split between scanner and
// job element sets.
func TestJobElemData_RejectsScannerName(t *testing.T) {
	elm := xmldoc.Element{
		Name: NsWSCN + ":ElementData",
		Attrs: []xmldoc.Attr{
			{Name: "Name", Value: NsWSCN + ":ScannerDescription"},
			{Name: "Valid", Value: "true"},
		},
	}

	_, err := decodeJobElemData(elm)
	if err == nil {
		t.Error("expected error for scanner-only Name value, got nil")
	}
}
