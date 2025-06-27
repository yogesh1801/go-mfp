// MFP - Miulti-Function Printers and scanners toolkit
// WSD core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// WSD Message header

package wsd

import (
	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// Header represents a common WSD message header.
type Header struct {
	Action      Action                          // Message action
	MessageID   AnyURI                          // Message identifier
	To          optional.Val[AnyURI]            // Message destination
	ReplyTo     optional.Val[EndpointReference] // Address to reply to
	RelatesTo   optional.Val[AnyURI]            // ID of related message
	AppSequence optional.Val[AppSequence]       // Message sequence
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
		var tmp AnyURI
		tmp, err = DecodeAnyURI(to.Elem)
		if err == nil {
			hdr.To = optional.New(tmp)
		}
	}

	if err == nil && replyTo.Found {
		var tmp EndpointReference
		tmp, err = DecodeEndpointReference(replyTo.Elem)
		if err == nil {
			hdr.ReplyTo = optional.New(tmp)
		}
	}

	if err == nil && relatesTo.Found {
		var tmp AnyURI
		tmp, err = DecodeAnyURI(relatesTo.Elem)
		if err == nil {
			hdr.RelatesTo = optional.New(tmp)
		}
	}

	if err == nil && appSequence.Found {
		var seq AppSequence
		seq, err = DecodeAppSequence(appSequence.Elem)
		if err == nil {
			hdr.AppSequence = optional.New(seq)
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

	if hdr.To != nil {
		elm.Children = append(elm.Children,
			xmldoc.Element{
				Name: NsAddressing + ":" + "To",
				Text: string(*hdr.To),
			})
	}

	if hdr.ReplyTo != nil {
		elm.Children = append(elm.Children,
			(*hdr.ReplyTo).ToXML(NsAddressing+":ReplyTo"))
	}

	if hdr.RelatesTo != nil {
		elm.Children = append(elm.Children,
			xmldoc.Element{
				Name: NsAddressing + ":" + "RelatesTo",
				Text: string(*hdr.RelatesTo),
			})
	}

	if hdr.AppSequence != nil {
		elm.Children = append(elm.Children, (*hdr.AppSequence).ToXML())
	}

	return elm
}
