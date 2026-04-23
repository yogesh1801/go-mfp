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

// String returns the QName string for a [ScanElemDataName].
func (n ScanElemDataName) String() string {
	switch n {
	case ScanElemDataDefaultScanTicket:
		return NsWSCN + ":DefaultScanTicket"
	case ScanElemDataScannerConfiguration:
		return NsWSCN + ":ScannerConfiguration"
	case ScanElemDataScannerDescription:
		return NsWSCN + ":ScannerDescription"
	case ScanElemDataScannerStatus:
		return NsWSCN + ":ScannerStatus"
	case ScanElemDataVendorSection:
		return NsWSCN + ":VendorSection"
	default:
		return "Unknown"
	}
}

// decodeScanElemDataName decodes a [ScanElemDataName] from its
// QName string. The prefix is stripped before matching because devices may
// use a different namespace prefix than we do for the same WS-Scan
// namespace URL.
func decodeScanElemDataName(s string) ScanElemDataName {
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
			{Name: "Name", Value: ed.Name.String()},
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

	ed.Name = decodeScanElemDataName(nameAttr.Attr.Value)
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
