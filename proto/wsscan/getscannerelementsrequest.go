// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// GetScannerElementsRequest: requests scanner element information

package wsscan

import (
	"errors"

	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// GetScannerElementsRequest enables a client to request information
// about the scanner from the WSD Scan Service.
type GetScannerElementsRequest struct {
	RequestedElements []RequestedElement // At least one required
}

// Action returns the [Action] associated with this body.
func (*GetScannerElementsRequest) Action() Action {
	return ActGetScannerElements
}

// ToXML encodes the body into an XML tree.
func (gser *GetScannerElementsRequest) ToXML() xmldoc.Element {
	return gser.toXML(NsWSCN + ":GetScannerElementsRequest")
}

// toXML generates XML tree for the GetScannerElementsRequest.
func (gser GetScannerElementsRequest) toXML(name string) xmldoc.Element {
	nameElems := make([]xmldoc.Element, len(gser.RequestedElements))
	for i, re := range gser.RequestedElements {
		nameElems[i] = re.toXML(NsWSCN + ":Name")
	}

	return xmldoc.Element{
		Name: name,
		Children: []xmldoc.Element{
			{Name: NsWSCN + ":RequestedElements", Children: nameElems},
		},
	}
}

// decodeGetScannerElementsRequest decodes GetScannerElementsRequest from
// the XML tree.
func decodeGetScannerElementsRequest(root xmldoc.Element) (
	gser GetScannerElementsRequest,
	err error,
) {
	defer func() { err = xmldoc.XMLErrWrap(root, err) }()

	requestedElements := xmldoc.Lookup{
		Name:     NsWSCN + ":RequestedElements",
		Required: true,
	}

	missed := root.Lookup(&requestedElements)
	if missed != nil && missed.Required {
		return gser, xmldoc.XMLErrMissed(missed.Name)
	}

	for _, child := range requestedElements.Elem.Children {
		if child.Name == NsWSCN+":Name" {
			re, decErr := decodeRequestedElement(child)
			if decErr != nil {
				return gser, decErr
			}
			gser.RequestedElements = append(gser.RequestedElements, re)
		}
	}

	if len(gser.RequestedElements) == 0 {
		return gser, errors.New("at least one Name is required")
	}

	return gser, nil
}
