// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// Test for ElementData

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

// TestElementData_RoundTrip_ScannerConfiguration verifies that an ElementData
// carrying a ScannerConfiguration child encodes to XML and decodes back
// to an identical value.
func TestElementData_RoundTrip_ScannerConfiguration(t *testing.T) {
	orig := ElementData{
		Name:  ElementDataScannerConfiguration,
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

	parsed, err := decodeElementData(elm)
	if err != nil {
		t.Fatalf("decodeElementData returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, parsed) {
		t.Errorf("expected %+v, got %+v", orig, parsed)
	}
}

// TestElementData_RoundTrip_ScannerDescription verifies that an ElementData
// carrying a ScannerDescription child encodes to XML and decodes back
// to an identical value.
func TestElementData_RoundTrip_ScannerDescription(t *testing.T) {
	orig := ElementData{
		Name:               ElementDataScannerDescription,
		Valid:              BooleanElement("true"),
		ScannerDescription: optional.New(createValidScannerDescription()),
	}
	elm := orig.toXML(NsWSCN + ":ElementData")

	parsed, err := decodeElementData(elm)
	if err != nil {
		t.Fatalf("decodeElementData returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, parsed) {
		t.Errorf("expected %+v, got %+v", orig, parsed)
	}
}

// TestElementData_RoundTrip_ScannerStatus verifies that an ElementData
// carrying a ScannerStatus child encodes to XML and decodes back
// to an identical value.
func TestElementData_RoundTrip_ScannerStatus(t *testing.T) {
	orig := ElementData{
		Name:          ElementDataScannerStatus,
		Valid:         BooleanElement("true"),
		ScannerStatus: optional.New(createValidScannerStatus()),
	}
	elm := orig.toXML(NsWSCN + ":ElementData")

	parsed, err := decodeElementData(elm)
	if err != nil {
		t.Fatalf("decodeElementData returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, parsed) {
		t.Errorf("expected %+v, got %+v", orig, parsed)
	}
}

// TestElementData_RoundTrip_DefaultScanTicket verifies that an ElementData
// carrying a DefaultScanTicket child encodes to XML and decodes back
// to an identical value.
func TestElementData_RoundTrip_DefaultScanTicket(t *testing.T) {
	orig := ElementData{
		Name:              ElementDataDefaultScanTicket,
		Valid:             BooleanElement("true"),
		DefaultScanTicket: optional.New(createValidScanTicket()),
	}
	elm := orig.toXML(NsWSCN + ":ElementData")

	parsed, err := decodeElementData(elm)
	if err != nil {
		t.Fatalf("decodeElementData returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, parsed) {
		t.Errorf("expected %+v, got %+v", orig, parsed)
	}
}

// TestElementData_MissingNameAttr verifies that decoding an ElementData
// element without the required Name attribute returns an error.
func TestElementData_MissingNameAttr(t *testing.T) {
	orig := ElementData{
		Name:  ElementDataScannerStatus,
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

	_, err := decodeElementData(elm)
	if err == nil {
		t.Error("expected error for missing Name attribute, got nil")
	}
}

// TestElementData_InvalidValidAttr verifies that decoding an ElementData
// element with an invalid Valid attribute value returns an error.
func TestElementData_InvalidValidAttr(t *testing.T) {
	orig := ElementData{
		Name:  ElementDataScannerStatus,
		Valid: BooleanElement("maybe"),
	}
	elm := orig.toXML(NsWSCN + ":ElementData")

	_, err := decodeElementData(elm)
	if err == nil {
		t.Error("expected error for invalid Valid value, got nil")
	}
}

// TestElementData_UnknownName verifies that decoding an ElementData element
// with an unrecognised Name attribute value returns an error.
func TestElementData_UnknownName(t *testing.T) {
	orig := ElementData{
		Name:  ElementDataScannerStatus,
		Valid: BooleanElement("true"),
	}
	elm := orig.toXML(NsWSCN + ":ElementData")
	for i, a := range elm.Attrs {
		if a.Name == "Name" {
			elm.Attrs[i].Value = "wscn:UnknownElement"
		}
	}

	_, err := decodeElementData(elm)
	if err == nil {
		t.Error("expected error for unknown Name value, got nil")
	}
}
