// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// ScanData: xop:Include reference to binary image in MTOM response

package wsscan

import "github.com/OpenPrinting/go-mfp/util/xmldoc"

// ScanData represents the <wscn:ScanData> element containing an
// xop:Include reference to the binary image attachment in the
// MTOM multipart response.
type ScanData struct {
	ContentID string
}

// toXML generates XML tree for the [ScanData].
func (sd ScanData) toXML(name string) xmldoc.Element {
	return xmldoc.Element{
		Name: name,
		Children: []xmldoc.Element{
			{
				Name: NsXOP + ":Include",
				Attrs: []xmldoc.Attr{
					{Name: "href", Value: "cid:" + sd.ContentID},
				},
			},
		},
	}
}
