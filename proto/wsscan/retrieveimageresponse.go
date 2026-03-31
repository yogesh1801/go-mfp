// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// RetrieveImageResponse: scan data response with MTOM/XOP reference

package wsscan

import "github.com/OpenPrinting/go-mfp/util/xmldoc"

// RetrieveImageResponse contains the WSD Scan Service's response
// to a client's RetrieveImage request. The ScanData element carries
// an xop:Include reference to the binary image part in the MTOM
// multipart response.
type RetrieveImageResponse struct {
	ScanData ScanData
}

// toXML generates XML tree for the [RetrieveImageResponse].
func (r RetrieveImageResponse) toXML(name string) xmldoc.Element {
	return xmldoc.Element{
		Name: name,
		Children: []xmldoc.Element{
			r.ScanData.toXML(NsWSCN + ":ScanData"),
		},
	}
}
