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

// JobElemDataName identifies which job schema element is carried
// in a [JobElemData].
type JobElemDataName int

// Known JobElemDataName values:
const (
	UnknownJobElemDataName   JobElemDataName = iota
	JobElemDataJobStatus                        // xmlns:JobStatus
	JobElemDataScanTicket                       // xmlns:ScanTicket
	JobElemDataDocuments                        // xmlns:Documents
	JobElemDataVendorSection                    // xmlns:VendorSection
)

// String returns the QName string for a [JobElemDataName].
func (n JobElemDataName) String() string {
	switch n {
	case JobElemDataJobStatus:
		return NsWSCN + ":JobStatus"
	case JobElemDataScanTicket:
		return NsWSCN + ":ScanTicket"
	case JobElemDataDocuments:
		return NsWSCN + ":Documents"
	case JobElemDataVendorSection:
		return NsWSCN + ":VendorSection"
	default:
		return "Unknown"
	}
}

// decodeJobElemDataName decodes a [JobElemDataName] from its QName
// string. The prefix is stripped before matching because devices may use
// a different namespace prefix than we do for the same WS-Scan namespace
// URL.
func decodeJobElemDataName(s string) JobElemDataName {
	if i := strings.LastIndex(s, ":"); i >= 0 {
		s = s[i+1:]
	}
	switch s {
	case "JobStatus":
		return JobElemDataJobStatus
	case "ScanTicket":
		return JobElemDataScanTicket
	case "Documents":
		return JobElemDataDocuments
	case "VendorSection":
		return JobElemDataVendorSection
	default:
		return UnknownJobElemDataName
	}
}

// JobElemData contains the data returned for a job-related schema
// request. The Name attribute identifies which schema element is present
// and Valid indicates whether the returned data is valid. Exactly one
// child element matching Name is expected to be present.
type JobElemData struct {
	Name       JobElemDataName
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
			{Name: "Name", Value: ed.Name.String()},
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

	ed.Name = decodeJobElemDataName(nameAttr.Attr.Value)
	if ed.Name == UnknownJobElemDataName {
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
