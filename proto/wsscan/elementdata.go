// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// ElementData: contains data returned for a scanner-related schema request

package wsscan

import (
	"fmt"

	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// ElementDataName identifies which scanner schema element is carried
// in an [ElementData].
type ElementDataName int

// Known ElementDataName values:
const (
	UnknownElementDataName          ElementDataName = iota
	ElementDataDefaultScanTicket                    // xmlns:DefaultScanTicket
	ElementDataScannerConfiguration                 // xmlns:ScannerConfiguration
	ElementDataScannerDescription                   // xmlns:ScannerDescription
	ElementDataScannerStatus                        // xmlns:ScannerStatus
	ElementDataVendorSection                        // xmlns:VendorSection
)

// String returns the QName string for an [ElementDataName].
func (n ElementDataName) String() string {
	switch n {
	case ElementDataDefaultScanTicket:
		return NsWSCN + ":DefaultScanTicket"
	case ElementDataScannerConfiguration:
		return NsWSCN + ":ScannerConfiguration"
	case ElementDataScannerDescription:
		return NsWSCN + ":ScannerDescription"
	case ElementDataScannerStatus:
		return NsWSCN + ":ScannerStatus"
	case ElementDataVendorSection:
		return NsWSCN + ":VendorSection"
	default:
		return "Unknown"
	}
}

// decodeElementDataName decodes an [ElementDataName] from its QName string.
func decodeElementDataName(s string) ElementDataName {
	switch s {
	case NsWSCN + ":DefaultScanTicket":
		return ElementDataDefaultScanTicket
	case NsWSCN + ":ScannerConfiguration":
		return ElementDataScannerConfiguration
	case NsWSCN + ":ScannerDescription":
		return ElementDataScannerDescription
	case NsWSCN + ":ScannerStatus":
		return ElementDataScannerStatus
	case NsXML + ":VendorSection":
		return ElementDataVendorSection
	default:
		return UnknownElementDataName
	}
}

// ElementData contains the data returned for a scanner-related schema request.
// The Name attribute identifies which schema element is present and Valid
// indicates whether the returned data is valid. Exactly one child element
// matching Name is expected to be present.
type ElementData struct {
	Name                 ElementDataName
	Valid                BooleanElement
	DefaultScanTicket    optional.Val[ScanTicket]
	ScannerConfiguration optional.Val[ScannerConfiguration]
	ScannerDescription   optional.Val[ScannerDescription]
	ScannerStatus        optional.Val[ScannerStatus]
}

// toXML creates an XML element for ElementData.
func (ed ElementData) toXML(name string) xmldoc.Element {
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

// decodeElementData decodes an ElementData from an XML element.
func decodeElementData(root xmldoc.Element) (ElementData, error) {
	var ed ElementData

	// Decode required attributes
	nameAttr := xmldoc.LookupAttr{Name: "Name", Required: true}
	validAttr := xmldoc.LookupAttr{Name: "Valid", Required: true}

	if missed := root.LookupAttrs(&nameAttr, &validAttr); missed != nil {
		return ed, xmldoc.XMLErrMissed(missed.Name)
	}

	ed.Name = decodeElementDataName(nameAttr.Attr.Value)
	if ed.Name == UnknownElementDataName {
		return ed, fmt.Errorf("ElementData: unknown Name %q",
			nameAttr.Attr.Value)
	}

	ed.Valid = BooleanElement(validAttr.Attr.Value)
	if err := ed.Valid.Validate(); err != nil {
		return ed, fmt.Errorf("ElementData: Valid: %w", err)
	}

	// Decode optional child elements
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
