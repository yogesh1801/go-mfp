// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// Name for RequestedElements element

package wsscan

import "github.com/OpenPrinting/go-mfp/util/xmldoc"

// ScannerRequestedElement identifies a section of the WSD Scan Service schema
// that a client wants data for in a GetScannerElementsRequest.
//
// Valid QName values:
//   - wscn:DefaultScanTicket
//   - wscn:ScannerDescription
//   - wscn:ScannerConfiguration
//   - wscn:ScannerStatus
//   - xmlns:VendorSection (vendor-defined extension)
type ScannerRequestedElement int

// Known ScannerRequestedElement values.
const (
	UnknownScannerElem           ScannerRequestedElement = iota
	ScannerElemDefaultScanTicket                         // wscn:DefaultScanTicket
	ScannerElemDescription                               // wscn:ScannerDescription
	ScannerElemConfiguration                             // wscn:ScannerConfiguration
	ScannerElemStatus                                    // wscn:ScannerStatus
	ScannerElemVendorSection                             // xmlns:VendorSection
)

// decodeScannerRequestedElement decodes [ScannerRequestedElement] from the XML tree.
func decodeScannerRequestedElement(root xmldoc.Element) (ScannerRequestedElement, error) {
	return decodeEnum(root, DecodeScannerRequestedElement)
}

// toXML generates XML tree for the [ScannerRequestedElement].
func (re ScannerRequestedElement) toXML(name string) xmldoc.Element {
	return xmldoc.Element{
		Name: name,
		Text: re.String(),
	}
}

// String returns the string representation of the [ScannerRequestedElement].
func (re ScannerRequestedElement) String() string {
	switch re {
	case ScannerElemDefaultScanTicket:
		return NsWSCN + ":DefaultScanTicket"
	case ScannerElemDescription:
		return NsWSCN + ":ScannerDescription"
	case ScannerElemConfiguration:
		return NsWSCN + ":ScannerConfiguration"
	case ScannerElemStatus:
		return NsWSCN + ":ScannerStatus"
	case ScannerElemVendorSection:
		return NsXML + ":VendorSection"
	}

	return "Unknown"
}

// DecodeScannerRequestedElement decodes [ScannerRequestedElement] out of its
// XML string representation.
func DecodeScannerRequestedElement(s string) ScannerRequestedElement {
	switch s {
	case NsWSCN + ":DefaultScanTicket":
		return ScannerElemDefaultScanTicket
	case NsWSCN + ":ScannerDescription":
		return ScannerElemDescription
	case NsWSCN + ":ScannerConfiguration":
		return ScannerElemConfiguration
	case NsWSCN + ":ScannerStatus":
		return ScannerElemStatus
	case NsXML + ":VendorSection":
		return ScannerElemVendorSection
	}

	return UnknownScannerElem
}

// JobRequestedElement identifies a section of the WSD Scan Service schema
// that a client wants data for in a GetJobElementsRequest.
//
// Valid QName values:
//   - wscn:JobStatus
//   - wscn:ScanTicket
//   - wscn:Documents
//   - xmlns:VendorSection (vendor-defined extension)
type JobRequestedElement int

// Known JobRequestedElement values.
const (
	UnknownJobElem       JobRequestedElement = iota
	JobElemStatus                            // wscn:JobStatus
	JobElemScanTicket                        // wscn:ScanTicket
	JobElemDocuments                         // wscn:Documents
	JobElemVendorSection                     // xmlns:VendorSection
)

// decodeJobRequestedElement decodes [JobRequestedElement] from the XML tree.
func decodeJobRequestedElement(root xmldoc.Element) (JobRequestedElement, error) {
	return decodeEnum(root, DecodeJobRequestedElement)
}

// toXML generates XML tree for the [JobRequestedElement].
func (re JobRequestedElement) toXML(name string) xmldoc.Element {
	return xmldoc.Element{
		Name: name,
		Text: re.String(),
	}
}

// String returns the string representation of the [JobRequestedElement].
func (re JobRequestedElement) String() string {
	switch re {
	case JobElemStatus:
		return NsWSCN + ":JobStatus"
	case JobElemScanTicket:
		return NsWSCN + ":ScanTicket"
	case JobElemDocuments:
		return NsWSCN + ":Documents"
	case JobElemVendorSection:
		return NsXML + ":VendorSection"
	}

	return "Unknown"
}

// DecodeJobRequestedElement decodes [JobRequestedElement] out of its XML
// string representation.
func DecodeJobRequestedElement(s string) JobRequestedElement {
	switch s {
	case NsWSCN + ":JobStatus":
		return JobElemStatus
	case NsWSCN + ":ScanTicket":
		return JobElemScanTicket
	case NsWSCN + ":Documents":
		return JobElemDocuments
	case NsXML + ":VendorSection":
		return JobElemVendorSection
	}

	return UnknownJobElem
}
