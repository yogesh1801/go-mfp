// MFP - Miulti-Function Printers and scanners toolkit
// WSD core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// WSD Message header

package wsd

import (
	"github.com/alexpevzner/mfp/xmldoc"
)

// Header represents a common WSD message header.
type Header struct {
	Action      Action            // Required: Message action
	MessageID   AnyURI            // Required: message identifier
	To          AnyURI            // Required: message destination
	ReplyTo     EndpointReference // Optional: address to reply to
	RelatesTo   AnyURI            // Optional: ID of related message
	AppSequence *AppSequence      // Optional: Message sequence
}

// DecodeHeader decodes message header [Header] from the XML tree
func DecodeHeader(root xmldoc.Element) (hdr Header, err error) {
	defer func() { err = xmlErrWrap(root, err) }()

	// Lookup header elements
	Action := xmldoc.Lookup{Name: NsAddressing + ":Action", Required: true}
	MessageID := xmldoc.Lookup{Name: NsAddressing + ":MessageID", Required: true}
	To := xmldoc.Lookup{Name: NsAddressing + ":To", Required: true}
	ReplyTo := xmldoc.Lookup{Name: NsAddressing + ":ReplyTo"}
	RelatesTo := xmldoc.Lookup{Name: NsAddressing + ":RelatesTo"}
	AppSequence := xmldoc.Lookup{Name: NsDiscovery + ":AppSequence"}

	missed := root.Lookup(&Action, &MessageID, &To, &ReplyTo,
		&RelatesTo, &AppSequence)
	if missed != nil {
		err = xmlErrMissed(missed.Name)
		return
	}

	// Decode header elements
	hdr.Action, err = DecodeAction(Action.Elem)
	if err == nil {
		hdr.MessageID, err = DecodeAnyURI(MessageID.Elem)
	}
	if err == nil {
		hdr.To, err = DecodeAnyURI(To.Elem)
	}
	if err == nil && ReplyTo.Found {
		hdr.ReplyTo, err = DecodeEndpointReference(ReplyTo.Elem)
	}
	if err == nil && RelatesTo.Found {
		hdr.RelatesTo, err = DecodeAnyURI(RelatesTo.Elem)
	}

	if err == nil && AppSequence.Found {
		hdr.AppSequence, err = DecodeAppSequence(AppSequence.Elem)
	}

	return
}

// ToXML generates XML tree for the message header
func (hdr Header) ToXML() xmldoc.Element {
	elm := xmldoc.Element{
		Name: NsSOAP + ":" + "Header",
		Children: []xmldoc.Element{
			{
				Name: NsAddressing + ":" + "Action",
				Text: hdr.Action.Encode(),
			},
			{
				Name: NsAddressing + ":" + "MessageID",
				Text: string(hdr.MessageID),
			},
			{
				Name: NsAddressing + ":" + "To",
				Text: string(hdr.To),
			},
		},
	}

	if hdr.ReplyTo.Address != "" {
		elm.Children = append(elm.Children,
			hdr.ReplyTo.ToXML(NsAddressing+":ReplyTo"))
	}

	if hdr.RelatesTo != "" {
		elm.Children = append(elm.Children,
			xmldoc.Element{
				Name: NsAddressing + ":" + "RelatesTo",
				Text: string(hdr.RelatesTo),
			})
	}

	if hdr.AppSequence != nil {
		elm.Children = append(elm.Children, hdr.AppSequence.ToXML())
	}

	return elm
}
