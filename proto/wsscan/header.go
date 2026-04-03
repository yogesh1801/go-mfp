// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// WS-Scan SOAP message header

package wsscan

import (
	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// Header represents the SOAP header for a WS-Scan message.
//
// Unlike WS-Discovery, WS-Scan uses HTTP as transport, so
// fields like From, AppSequence and ReplyTo are not needed.
type Header struct {
	Action    Action               // Message action (type)
	MessageID AnyURI               // Unique message identifier
	To        optional.Val[AnyURI] // Message destination
	RelatesTo optional.Val[AnyURI] // ID of the request being replied to
}

// decodeHeader decodes a [Header] from the XML tree.
func decodeHeader(root xmldoc.Element) (hdr Header, err error) {
	defer func() { err = xmldoc.XMLErrWrap(root, err) }()

	action := xmldoc.Lookup{
		Name:     NsAddressing + ":Action",
		Required: true,
	}
	messageID := xmldoc.Lookup{
		Name:     NsAddressing + ":MessageID",
		Required: true,
	}
	to := xmldoc.Lookup{
		Name:     NsAddressing + ":To",
		Required: false,
	}
	relatesTo := xmldoc.Lookup{
		Name:     NsAddressing + ":RelatesTo",
		Required: false,
	}

	missed := root.Lookup(&action, &messageID, &to, &relatesTo)
	if missed != nil {
		err = xmldoc.XMLErrMissed(missed.Name)
		return
	}

	// Decode Action
	hdr.Action, err = decodeAction(action.Elem)
	if err != nil {
		return
	}

	// Decode MessageID
	hdr.MessageID, err = DecodeAnyURI(messageID.Elem)
	if err != nil {
		return
	}

	// Decode To (optional)
	if to.Found {
		var tmp AnyURI
		tmp, err = DecodeAnyURI(to.Elem)
		if err != nil {
			return
		}
		hdr.To = optional.New(tmp)
	}

	// Decode RelatesTo (optional)
	if relatesTo.Found {
		var tmp AnyURI
		tmp, err = DecodeAnyURI(relatesTo.Elem)
		if err != nil {
			return
		}
		hdr.RelatesTo = optional.New(tmp)
	}

	return
}

// toXML generates the XML tree for the SOAP header.
func (hdr Header) toXML() xmldoc.Element {
	elm := xmldoc.Element{
		Name: NsSOAP + ":Header",
		Children: []xmldoc.Element{
			{
				Name: NsAddressing + ":Action",
				Text: hdr.Action.Encode(),
			},
			{
				Name: NsAddressing + ":MessageID",
				Text: string(hdr.MessageID),
			},
		},
	}

	if hdr.To != nil {
		elm.Children = append(elm.Children, xmldoc.Element{
			Name: NsAddressing + ":To",
			Text: string(*hdr.To),
		})
	}

	if hdr.RelatesTo != nil {
		elm.Children = append(elm.Children, xmldoc.Element{
			Name: NsAddressing + ":RelatesTo",
			Text: string(*hdr.RelatesTo),
		})
	}

	return elm
}
