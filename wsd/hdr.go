// MFP - Miulti-Function Printers and scanners toolkit
// WSD core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Package documentation

package wsd

import (
	"errors"
	"fmt"

	"github.com/alexpevzner/mfp/xmldoc"
)

// Header represents a common WSD message header.
type Header struct {
	Action      Action      // Message action
	MessageID   AnyURI      // Required: message identifier
	To          AnyURI      // Required: message destination
	RelatesTo   AnyURI      // Reply-To or similar
	AppSequence AppSequence // Message sequence (recv only)
}

// DecodeHdr decodes message header [Header] from the XML tree
func DecodeHdr(root xmldoc.Element) (hdr Header, err error) {
	defer func() { err = xmlErrWrap(root, err) }()

	// Lookup header elements
	Action := xmldoc.Lookup{Name: NsAddressing + ":Action", Required: true}
	MessageID := xmldoc.Lookup{Name: NsAddressing + ":MessageID", Required: true}
	To := xmldoc.Lookup{Name: NsAddressing + ":To", Required: true}
	RelatesTo := xmldoc.Lookup{Name: NsAddressing + ":RelatesTo"}
	AppSequence := xmldoc.Lookup{Name: NsAddressing + ":AppSequence", Required: true}

	missed := root.Lookup(&Action, &MessageID, &To, &RelatesTo, &AppSequence)
	if missed != nil {
		err = fmt.Errorf("%s: missed", missed.Name)
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
	if err == nil && RelatesTo.Found {
		hdr.To, err = DecodeAnyURI(RelatesTo.Elem)
	}
	if err == nil {
		hdr.AppSequence, err = DecodeAppSequence(AppSequence.Elem)
	}

	err = errors.New("not implemented")
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

	if hdr.RelatesTo != "" {
		elm.Children = append(elm.Children,
			xmldoc.Element{
				Name: NsAddressing + ":" + "RelatesTo",
				Text: string(hdr.RelatesTo),
			})
	}

	if !hdr.AppSequence.Skip {
		elm.Children = append(elm.Children, hdr.AppSequence.ToXML())
	}

	return elm
}
