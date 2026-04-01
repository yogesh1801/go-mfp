// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// GetScannerElementsResponse: the Scan Service's response to a
// GetScannerElementsRequest

package wsscan

import (
	"errors"

	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// GetScannerElementsResponse contains the WSD Scan Service's response to a
// client's GetScannerElementsRequest. ScannerElements holds one [ElementData]
// entry for each schema element that was requested.
type GetScannerElementsResponse struct {
	ScannerElements []ElementData
}

// Action returns the [Action] associated with this body.
func (GetScannerElementsResponse) Action() Action { return ActGetScannerElementsResponse }

// ToXML encodes the body into an XML tree.
func (r GetScannerElementsResponse) ToXML() xmldoc.Element {
	return r.toXML(NsWSCN + ":GetScannerElementsResponse")
}

// toXML generates XML tree for the [GetScannerElementsResponse].
func (r GetScannerElementsResponse) toXML(name string) xmldoc.Element {
	edChildren := make([]xmldoc.Element, len(r.ScannerElements))
	for i, ed := range r.ScannerElements {
		edChildren[i] = ed.toXML(NsWSCN + ":ElementData")
	}

	return xmldoc.Element{
		Name: name,
		Children: []xmldoc.Element{
			{Name: NsWSCN + ":ScannerElements", Children: edChildren},
		},
	}
}

// decodeGetScannerElementsResponse decodes [GetScannerElementsResponse] from
// the XML tree.
func decodeGetScannerElementsResponse(root xmldoc.Element) (
	GetScannerElementsResponse, error,
) {
	var r GetScannerElementsResponse

	scannerElements := xmldoc.Lookup{
		Name:     NsWSCN + ":ScannerElements",
		Required: true,
	}

	if missed := root.Lookup(&scannerElements); missed != nil {
		return r, xmldoc.XMLErrMissed(missed.Name)
	}

	for _, child := range scannerElements.Elem.Children {
		if child.Name != NsWSCN+":ElementData" {
			continue
		}
		ed, err := decodeElementData(child)
		if err != nil {
			return r, err
		}
		r.ScannerElements = append(r.ScannerElements, ed)
	}

	if len(r.ScannerElements) == 0 {
		return r, errors.New(
			"ScannerElements: at least one ElementData is required")
	}

	return r, nil
}
