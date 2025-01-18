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
	defer func() { err = xmldoc.XMLErrWrap(root, err) }()

	// Lookup header elements
	action := xmldoc.Lookup{Name: NsAddressing + ":Action", Required: true}
	messageID := xmldoc.Lookup{Name: NsAddressing + ":MessageID", Required: true}
	to := xmldoc.Lookup{Name: NsAddressing + ":To"}
	replyTo := xmldoc.Lookup{Name: NsAddressing + ":ReplyTo"}
	relatesTo := xmldoc.Lookup{Name: NsAddressing + ":RelatesTo"}
	appSequence := xmldoc.Lookup{Name: NsDiscovery + ":AppSequence"}

	missed := root.Lookup(&action, &messageID, &to, &replyTo,
		&relatesTo, &appSequence)
	if missed != nil {
		err = xmldoc.XMLErrMissed(missed.Name)
		return
	}

	// Decode header elements
	hdr.Action, err = DecodeAction(action.Elem)
	if err == nil {
		hdr.MessageID, err = DecodeAnyURI(messageID.Elem)
	}
	if err == nil && to.Found {
		hdr.To, err = DecodeAnyURI(to.Elem)
	}
	if err == nil && replyTo.Found {
		hdr.ReplyTo, err = DecodeEndpointReference(replyTo.Elem)
	}
	if err == nil && relatesTo.Found {
		hdr.RelatesTo, err = DecodeAnyURI(relatesTo.Elem)
	}

	if err == nil && appSequence.Found {
		var seq AppSequence
		seq, err = DecodeAppSequence(appSequence.Elem)
		if err == nil {
			hdr.AppSequence = &seq
		}
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
		},
	}

	if hdr.To != "" {
		elm.Children = append(elm.Children,
			xmldoc.Element{
				Name: NsAddressing + ":" + "To",
				Text: string(hdr.To),
			})
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
