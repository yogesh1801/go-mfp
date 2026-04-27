// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// ScanElemData: data returned for a scanner-related schema request

package wsscan

import (
	"fmt"
	"strings"

	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// ScanElemName identifies which scanner schema element is
// carried in a [ScanElemData].
type ScanElemName int

// Known ScanElemName values:
const (
	UnknownScanElem           ScanElemName = iota
	ScanElemDefaultScanTicket              // wscn:DefaultScanTicket
	ScanElemConfiguration                  // wscn:ScannerConfiguration
	ScanElemDescription                    // wscn:ScannerDescription
	ScanElemStatus                         // wscn:ScannerStatus
	ScanElemVendorSection                  // wscn:VendorSection
)

// String returns the local name for a [ScanElemName].
func (n ScanElemName) String() string {
	switch n {
	case ScanElemDefaultScanTicket:
		return "DefaultScanTicket"
	case ScanElemConfiguration:
		return "ScannerConfiguration"
	case ScanElemDescription:
		return "ScannerDescription"
	case ScanElemStatus:
		return "ScannerStatus"
	case ScanElemVendorSection:
		return "VendorSection"
	default:
		return "Unknown"
	}
}

// Encode returns the QName string for XML encoding of the
// [ScanElemName], used both as the value of the Name attribute on
// [ScanElemData] and as the text content of a wscn:Name element inside
// a GetScannerElementsRequest.
func (n ScanElemName) Encode() string {
	return NsWSCN + ":" + n.String()
}

// toXML generates an XML element whose text content is the QName for
// the [ScanElemName]. Used by [GetScannerElementsRequest] to encode
// each requested element name.
func (n ScanElemName) toXML(name string) xmldoc.Element {
	return xmldoc.Element{
		Name: name,
		Text: n.Encode(),
	}
}

// decodeScanElemName decodes a [ScanElemName] from an XML element
// whose text content is the QName form. Returns an error if the value is
// not a known name.
func decodeScanElemName(root xmldoc.Element) (ScanElemName, error) {
	return decodeEnum(root, DecodeScanElemName)
}

// DecodeScanElemName decodes a [ScanElemName] from its
// QName string. The prefix is stripped before matching because devices may
// use a different namespace prefix than we do for the same WS-Scan
// namespace URL.
func DecodeScanElemName(s string) ScanElemName {
	if i := strings.LastIndex(s, ":"); i >= 0 {
		s = s[i+1:]
	}
	switch s {
	case "DefaultScanTicket":
		return ScanElemDefaultScanTicket
	case "ScannerConfiguration":
		return ScanElemConfiguration
	case "ScannerDescription":
		return ScanElemDescription
	case "ScannerStatus":
		return ScanElemStatus
	case "VendorSection":
		return ScanElemVendorSection
	default:
		return UnknownScanElem
	}
}

// ScanElemData contains the data returned for a scanner-related
// schema request. The Name attribute identifies which schema element is
// present and Valid indicates whether the returned data is valid.
// Exactly one child element matching Name is expected to be present.
type ScanElemData struct {
	Name                 ScanElemName
	Valid                BooleanElement
	DefaultScanTicket    optional.Val[ScanTicket]
	ScannerConfiguration optional.Val[ScannerConfiguration]
	ScannerDescription   optional.Val[ScannerDescription]
	ScannerStatus        optional.Val[ScannerStatus]
}

// toXML creates an XML element for ScanElemData.
func (ed ScanElemData) toXML(name string) xmldoc.Element {
	elm := xmldoc.Element{
		Name: name,
		Attrs: []xmldoc.Attr{
			{Name: "Name", Value: NsWSCN + ":" + ed.Name.String()},
			{Name: "Valid", Value: string(ed.Valid)},
		},
	}

	if ed.DefaultScanTicket != nil {
		elm.Children = append(elm.Children,
			optional.Get(ed.DefaultScanTicket).toXML(
				NsWSCN+":DefaultScanTicket"))
	}
	if ed.ScannerConfiguration != nil {
		elm.Children = append(elm.Children,
			optional.Get(ed.ScannerConfiguration).toXML(
				NsWSCN+":ScannerConfiguration"))
	}
	if ed.ScannerDescription != nil {
		elm.Children = append(elm.Children,
			optional.Get(ed.ScannerDescription).toXML(
				NsWSCN+":ScannerDescription"))
	}
	if ed.ScannerStatus != nil {
		elm.Children = append(elm.Children,
			optional.Get(ed.ScannerStatus).toXML(
				NsWSCN+":ScannerStatus"))
	}

	return elm
}

// decodeScanElemData decodes a [ScanElemData] from an XML element.
func decodeScanElemData(root xmldoc.Element) (ScanElemData, error) {
	var ed ScanElemData

	nameAttr := xmldoc.LookupAttr{Name: "Name", Required: true}
	validAttr := xmldoc.LookupAttr{Name: "Valid", Required: true}

	if missed := root.LookupAttrs(&nameAttr, &validAttr); missed != nil {
		return ed, xmldoc.XMLErrMissed(missed.Name)
	}

	ed.Name = DecodeScanElemName(nameAttr.Attr.Value)
	if ed.Name == UnknownScanElem {
		return ed, fmt.Errorf("ScanElemData: unknown Name %q",
			nameAttr.Attr.Value)
	}

	ed.Valid = BooleanElement(validAttr.Attr.Value)
	if err := ed.Valid.Validate(); err != nil {
		return ed, fmt.Errorf("ScanElemData: Valid: %w", err)
	}

	defaultScanTicket := xmldoc.Lookup{
		Name: NsWSCN + ":DefaultScanTicket"}
	scannerConfiguration := xmldoc.Lookup{
		Name: NsWSCN + ":ScannerConfiguration"}
	scannerDescription := xmldoc.Lookup{
		Name: NsWSCN + ":ScannerDescription"}
	scannerStatus := xmldoc.Lookup{
		Name: NsWSCN + ":ScannerStatus"}

	root.Lookup(&defaultScanTicket, &scannerConfiguration,
		&scannerDescription, &scannerStatus)

	if defaultScanTicket.Found {
		st, err := decodeScanTicket(defaultScanTicket.Elem)
		if err != nil {
			return ed, fmt.Errorf("DefaultScanTicket: %w", err)
		}
		ed.DefaultScanTicket = optional.New(st)
	}
	if scannerConfiguration.Found {
		sc, err := decodeScannerConfiguration(scannerConfiguration.Elem)
		if err != nil {
			return ed, fmt.Errorf("ScannerConfiguration: %w", err)
		}
		ed.ScannerConfiguration = optional.New(sc)
	}
	if scannerDescription.Found {
		sd, err := decodeScannerDescription(scannerDescription.Elem)
		if err != nil {
			return ed, fmt.Errorf("ScannerDescription: %w", err)
		}
		ed.ScannerDescription = optional.New(sd)
	}
	if scannerStatus.Found {
		ss, err := decodeScannerStatus(scannerStatus.Elem)
		if err != nil {
			return ed, fmt.Errorf("ScannerStatus: %w", err)
		}
		ed.ScannerStatus = optional.New(ss)
	}

	return ed, nil
}
