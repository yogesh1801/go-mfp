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
	UnknownScannerRequestedElement                ScannerRequestedElement = iota
	ScannerRequestedElementDefaultScanTicket                               // wscn:DefaultScanTicket
	ScannerRequestedElementDescription                                     // wscn:ScannerDescription
	ScannerRequestedElementConfiguration                                   // wscn:ScannerConfiguration
	ScannerRequestedElementStatus                                          // wscn:ScannerStatus
	ScannerRequestedElementVendorSection                                   // xmlns:VendorSection
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
	case ScannerRequestedElementDefaultScanTicket:
		return NsWSCN + ":DefaultScanTicket"
	case ScannerRequestedElementDescription:
		return NsWSCN + ":ScannerDescription"
	case ScannerRequestedElementConfiguration:
		return NsWSCN + ":ScannerConfiguration"
	case ScannerRequestedElementStatus:
		return NsWSCN + ":ScannerStatus"
	case ScannerRequestedElementVendorSection:
		return NsXML + ":VendorSection"
	}

	return "Unknown"
}

// DecodeScannerRequestedElement decodes [ScannerRequestedElement] out of its
// XML string representation.
func DecodeScannerRequestedElement(s string) ScannerRequestedElement {
	switch s {
	case NsWSCN + ":DefaultScanTicket":
		return ScannerRequestedElementDefaultScanTicket
	case NsWSCN + ":ScannerDescription":
		return ScannerRequestedElementDescription
	case NsWSCN + ":ScannerConfiguration":
		return ScannerRequestedElementConfiguration
	case NsWSCN + ":ScannerStatus":
		return ScannerRequestedElementStatus
	case NsXML + ":VendorSection":
		return ScannerRequestedElementVendorSection
	}

	return UnknownScannerRequestedElement
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
	UnknownJobRequestedElement       JobRequestedElement = iota
	JobRequestedElementJobStatus                         // wscn:JobStatus
	JobRequestedElementScanTicket                        // wscn:ScanTicket
	JobRequestedElementDocuments                         // wscn:Documents
	JobRequestedElementVendorSection                     // xmlns:VendorSection
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
	case JobRequestedElementJobStatus:
		return NsWSCN + ":JobStatus"
	case JobRequestedElementScanTicket:
		return NsWSCN + ":ScanTicket"
	case JobRequestedElementDocuments:
		return NsWSCN + ":Documents"
	case JobRequestedElementVendorSection:
		return NsXML + ":VendorSection"
	}

	return "Unknown"
}

// DecodeJobRequestedElement decodes [JobRequestedElement] out of its XML
// string representation.
func DecodeJobRequestedElement(s string) JobRequestedElement {
	switch s {
	case NsWSCN + ":JobStatus":
		return JobRequestedElementJobStatus
	case NsWSCN + ":ScanTicket":
		return JobRequestedElementScanTicket
	case NsWSCN + ":Documents":
		return JobRequestedElementDocuments
	case NsXML + ":VendorSection":
		return JobRequestedElementVendorSection
	}

	return UnknownJobRequestedElement
}
