// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// Name for RequestedElements element

package wsscan

import "github.com/OpenPrinting/go-mfp/util/xmldoc"

// RequestedElement identifies a section of the WSD Scan Service schema
// that a client wants data for in a GetScannerElementsRequest.
//
// For GetScannerElementsRequest, one of the following QName values:
//   - wscn:ScannerDescription
//   - wscn:ScannerConfiguration
//   - wscn:ScannerStatus
//   - xmlns:VendorSection (vendor-defined extension)
type RequestedElement int

// Known RequestedElement values.
const (
	UnknownRequestedElement    RequestedElement = iota
	RequestedElementDescription                 // wscn:ScannerDescription
	RequestedElementConfiguration               // wscn:ScannerConfiguration
	RequestedElementStatus                      // wscn:ScannerStatus
	RequestedElementVendorSection               // xmlns:VendorSection
)

// decodeRequestedElement decodes [RequestedElement] from the XML tree.
func decodeRequestedElement(root xmldoc.Element) (RequestedElement, error) {
	return decodeEnum(root, DecodeRequestedElement)
}

// toXML generates XML tree for the [RequestedElement].
func (re RequestedElement) toXML(name string) xmldoc.Element {
	return xmldoc.Element{
		Name: name,
		Text: re.String(),
	}
}

// String returns the string representation of the [RequestedElement].
func (re RequestedElement) String() string {
	switch re {
	case RequestedElementDescription:
		return NsWSCN + ":ScannerDescription"
	case RequestedElementConfiguration:
		return NsWSCN + ":ScannerConfiguration"
	case RequestedElementStatus:
		return NsWSCN + ":ScannerStatus"
	case RequestedElementVendorSection:
		return NsXML + ":VendorSection"
	}

	return "Unknown"
}

// DecodeRequestedElement decodes [RequestedElement] out of its XML string
// representation.
func DecodeRequestedElement(s string) RequestedElement {
	switch s {
	case NsWSCN + ":ScannerDescription":
		return RequestedElementDescription
	case NsWSCN + ":ScannerConfiguration":
		return RequestedElementConfiguration
	case NsWSCN + ":ScannerStatus":
		return RequestedElementStatus
	case NsXML + ":VendorSection":
		return RequestedElementVendorSection
	}

	return UnknownRequestedElement
}
