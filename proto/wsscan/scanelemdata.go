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

// ScanElemDataName identifies which scanner schema element is
// carried in a [ScanElemData].
type ScanElemDataName int

// Known ScanElemDataName values:
const (
	UnknownScanElemDataName          ScanElemDataName = iota
	ScanElemDataDefaultScanTicket                           // xmlns:DefaultScanTicket
	ScanElemDataScannerConfiguration                        // xmlns:ScannerConfiguration
	ScanElemDataScannerDescription                          // xmlns:ScannerDescription
	ScanElemDataScannerStatus                               // xmlns:ScannerStatus
	ScanElemDataVendorSection                               // xmlns:VendorSection
)

// String returns the local name for a [ScanElemDataName].
func (n ScanElemDataName) String() string {
	switch n {
	case ScanElemDataDefaultScanTicket:
		return "DefaultScanTicket"
	case ScanElemDataScannerConfiguration:
		return "ScannerConfiguration"
	case ScanElemDataScannerDescription:
		return "ScannerDescription"
	case ScanElemDataScannerStatus:
		return "ScannerStatus"
	case ScanElemDataVendorSection:
		return "VendorSection"
	default:
		return "Unknown"
	}
}

// Encode returns the QName string for XML encoding of the
// [ScanElemDataName], used both as the value of the Name attribute on
// [ScanElemData] and as the text content of a wscn:Name element inside
// a GetScannerElementsRequest.
func (n ScanElemDataName) Encode() string {
	return NsWSCN + ":" + n.String()
}

// toXML generates an XML element whose text content is the QName for
// the [ScanElemDataName]. Used by [GetScannerElementsRequest] to encode
// each requested element name.
func (n ScanElemDataName) toXML(name string) xmldoc.Element {
	return xmldoc.Element{
		Name: name,
		Text: n.Encode(),
	}
}

// decodeScanElemDataName decodes a [ScanElemDataName] from an XML element
// whose text content is the QName form. Returns an error if the value is
// not a known name.
func decodeScanElemDataName(root xmldoc.Element) (ScanElemDataName, error) {
	return decodeEnum(root, DecodeScanElemDataName)
}

// DecodeScanElemDataName decodes a [ScanElemDataName] from its
// QName string. The prefix is stripped before matching because devices may
// use a different namespace prefix than we do for the same WS-Scan
// namespace URL.
func DecodeScanElemDataName(s string) ScanElemDataName {
	if i := strings.LastIndex(s, ":"); i >= 0 {
		s = s[i+1:]
	}
	switch s {
	case "DefaultScanTicket":
		return ScanElemDataDefaultScanTicket
	case "ScannerConfiguration":
		return ScanElemDataScannerConfiguration
	case "ScannerDescription":
		return ScanElemDataScannerDescription
	case "ScannerStatus":
		return ScanElemDataScannerStatus
	case "VendorSection":
		return ScanElemDataVendorSection
	default:
		return UnknownScanElemDataName
	}
}

// ScanElemData contains the data returned for a scanner-related
// schema request. The Name attribute identifies which schema element is
// present and Valid indicates whether the returned data is valid.
// Exactly one child element matching Name is expected to be present.
type ScanElemData struct {
	Name                 ScanElemDataName
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

	ed.Name = DecodeScanElemDataName(nameAttr.Attr.Value)
	if ed.Name == UnknownScanElemDataName {
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
