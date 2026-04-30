// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// RetrieveImageResponse: scan data response with MTOM/XOP reference

package wsscan

import (
	"io"
	"strings"

	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// RetrieveImageResponse contains the WSD Scan Service's response
// to a client's RetrieveImage request. The ScanData element carries
// an xop:Include reference to the binary image part in the MTOM
// multipart response.
//
// Image holds the binary image data as an [io.ReadCloser].
// On the server side, closing it is typically a no-op.
// On the client side, closing it releases the underlying
// HTTP response body.
type RetrieveImageResponse struct {
	ScanData    ScanData
	Image       io.ReadCloser
	ContentType string
}

// Action returns the [Action] associated with this body.
func (*RetrieveImageResponse) Action() Action { return ActRetrieveImageResponse }

// ToXML encodes the body into an XML tree.
func (r *RetrieveImageResponse) ToXML() xmldoc.Element {
	return r.toXML(NsWSCN + ":RetrieveImageResponse")
}

// decodeRetrieveImageResponse decodes [RetrieveImageResponse] from the XML tree,
// extracting the ContentID from the xop:Include href inside ScanData.
func decodeRetrieveImageResponse(root xmldoc.Element) (
	RetrieveImageResponse, error) {

	var r RetrieveImageResponse

	scanData := xmldoc.Lookup{Name: NsWSCN + ":ScanData", Required: true}
	if missed := root.Lookup(&scanData); missed != nil {
		return r, xmldoc.XMLErrMissed(missed.Name)
	}

	include := xmldoc.Lookup{Name: NsXOP + ":Include", Required: true}
	if missed := scanData.Elem.Lookup(&include); missed != nil {
		return r, xmldoc.XMLErrMissed(missed.Name)
	}

	href, _ := include.Elem.AttrByName("href")
	r.ScanData.ContentID = strings.TrimPrefix(href.Value, "cid:")

	return r, nil
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
