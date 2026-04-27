// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// JobElemData: data returned for a job-related schema request

package wsscan

import (
	"fmt"
	"strings"

	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// JobElemName identifies which job schema element is carried
// in a [JobElemData].
type JobElemName int

// Known JobElemName values:
const (
	UnknownJobElem       JobElemName = iota
	JobElemStatus                    // wscn:JobStatus
	JobElemScanTicket                // wscn:ScanTicket
	JobElemDocuments                 // wscn:Documents
	JobElemVendorSection             // wscn:VendorSection
)

// String returns the local name for a [JobElemName].
func (n JobElemName) String() string {
	switch n {
	case JobElemStatus:
		return "JobStatus"
	case JobElemScanTicket:
		return "ScanTicket"
	case JobElemDocuments:
		return "Documents"
	case JobElemVendorSection:
		return "VendorSection"
	default:
		return "Unknown"
	}
}

// Encode returns the QName string for XML encoding of the
// [JobElemName], used both as the value of the Name attribute on
// [JobElemData] and as the text content of a wscn:Name element inside
// a GetJobElementsRequest.
func (n JobElemName) Encode() string {
	return NsWSCN + ":" + n.String()
}

// toXML generates an XML element whose text content is the QName for
// the [JobElemName]. Used by [GetJobElementsRequest] to encode each
// requested element name.
func (n JobElemName) toXML(name string) xmldoc.Element {
	return xmldoc.Element{
		Name: name,
		Text: n.Encode(),
	}
}

// decodeJobElemName decodes a [JobElemName] from an XML element
// whose text content is the QName form. Returns an error if the value is
// not a known name.
func decodeJobElemName(root xmldoc.Element) (JobElemName, error) {
	return decodeEnum(root, DecodeJobElemName)
}

// DecodeJobElemName decodes a [JobElemName] from its QName
// string. The prefix is stripped before matching because devices may use
// a different namespace prefix than we do for the same WS-Scan namespace
// URL.
func DecodeJobElemName(s string) JobElemName {
	if i := strings.LastIndex(s, ":"); i >= 0 {
		s = s[i+1:]
	}
	switch s {
	case "JobStatus":
		return JobElemStatus
	case "ScanTicket":
		return JobElemScanTicket
	case "Documents":
		return JobElemDocuments
	case "VendorSection":
		return JobElemVendorSection
	default:
		return UnknownJobElem
	}
}

// JobElemData contains the data returned for a job-related schema
// request. The Name attribute identifies which schema element is present
// and Valid indicates whether the returned data is valid. Exactly one
// child element matching Name is expected to be present.
type JobElemData struct {
	Name       JobElemName
	Valid      BooleanElement
	JobStatus  optional.Val[JobStatus]
	ScanTicket optional.Val[ScanTicket]
	Documents  optional.Val[Documents]
}

// toXML creates an XML element for JobElemData.
func (ed JobElemData) toXML(name string) xmldoc.Element {
	elm := xmldoc.Element{
		Name: name,
		Attrs: []xmldoc.Attr{
			{Name: "Name", Value: NsWSCN + ":" + ed.Name.String()},
			{Name: "Valid", Value: string(ed.Valid)},
		},
	}

	if ed.JobStatus != nil {
		elm.Children = append(elm.Children,
			optional.Get(ed.JobStatus).toXML(
				NsWSCN+":JobStatus"))
	}
	if ed.ScanTicket != nil {
		elm.Children = append(elm.Children,
			optional.Get(ed.ScanTicket).toXML(
				NsWSCN+":ScanTicket"))
	}
	if ed.Documents != nil {
		elm.Children = append(elm.Children,
			optional.Get(ed.Documents).toXML(
				NsWSCN+":Documents"))
	}

	return elm
}

// decodeJobElemData decodes a [JobElemData] from an XML element.
func decodeJobElemData(root xmldoc.Element) (JobElemData, error) {
	var ed JobElemData

	nameAttr := xmldoc.LookupAttr{Name: "Name", Required: true}
	validAttr := xmldoc.LookupAttr{Name: "Valid", Required: true}

	if missed := root.LookupAttrs(&nameAttr, &validAttr); missed != nil {
		return ed, xmldoc.XMLErrMissed(missed.Name)
	}

	ed.Name = DecodeJobElemName(nameAttr.Attr.Value)
	if ed.Name == UnknownJobElem {
		return ed, fmt.Errorf("JobElemData: unknown Name %q",
			nameAttr.Attr.Value)
	}

	ed.Valid = BooleanElement(validAttr.Attr.Value)
	if err := ed.Valid.Validate(); err != nil {
		return ed, fmt.Errorf("JobElemData: Valid: %w", err)
	}

	jobStatus := xmldoc.Lookup{Name: NsWSCN + ":JobStatus"}
	scanTicket := xmldoc.Lookup{Name: NsWSCN + ":ScanTicket"}
	documents := xmldoc.Lookup{Name: NsWSCN + ":Documents"}

	root.Lookup(&jobStatus, &scanTicket, &documents)

	if jobStatus.Found {
		js, err := decodeJobStatus(jobStatus.Elem)
		if err != nil {
			return ed, fmt.Errorf("JobStatus: %w", err)
		}
		ed.JobStatus = optional.New(js)
	}
	if scanTicket.Found {
		st, err := decodeScanTicket(scanTicket.Elem)
		if err != nil {
			return ed, fmt.Errorf("ScanTicket: %w", err)
		}
		ed.ScanTicket = optional.New(st)
	}
	if documents.Found {
		docs, err := decodeDocuments(documents.Elem)
		if err != nil {
			return ed, fmt.Errorf("Documents: %w", err)
		}
		ed.Documents = optional.New(docs)
	}

	return ed, nil
}
