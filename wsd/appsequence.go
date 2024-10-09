// MFP - Miulti-Function Printers and scanners toolkit
// WSD core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Package documentation

package wsd

import (
	"fmt"
	"strconv"

	"github.com/alexpevzner/mfp/xmldoc"
)

// AppSequence provides a mechanism that allows a receiver
// to order messages that may have been received out of order.
//
// It is included into the announcement and response messages
// ([Hello], [Bye], [ProbeMatches], and [ResolveMatches]).
type AppSequence struct {
	InstanceID    uint64 // MUST increment on each reboot
	MessageNumber uint64 // MUST increment on each message
	SequenceID    AnyURI // Optional: sequence within instance
	Skip          bool   // Skip when sending
}

// DecodeAppSequence decodes AppSequence from the XML tree
func DecodeAppSequence(root xmldoc.Element) (seq AppSequence, err error) {
	defer func() { err = xmlErrWrap(root, err) }()

	InstanceID := xmldoc.LookupAttr{
		Name: NsAddressing + ":InstanceID", Required: true,
	}
	MessageNumber := xmldoc.LookupAttr{
		Name: NsAddressing + ":MessageNumber", Required: true,
	}
	SequenceID := xmldoc.LookupAttr{
		Name: NsAddressing + ":SequenceID",
	}

	missed := root.LookupAttrs(&InstanceID, &MessageNumber, &SequenceID)
	if missed != nil {
		err = fmt.Errorf("%s: missed", missed.Name)
		return
	}

	seq.InstanceID, err = decodeUint64Attr(InstanceID.Attr)
	if err == nil {
		seq.MessageNumber, err = decodeUint64Attr(MessageNumber.Attr)
	}

	if err == nil && SequenceID.Found {
		seq.SequenceID, err = DecodeAnyURIAttr(SequenceID.Attr)
	}

	return
}

// ToXML generates XML tree for the AppSequence
func (seq AppSequence) ToXML() xmldoc.Element {
	elm := xmldoc.Element{
		Name: NsDiscovery + ":" + "AppSequence",
		Attrs: []xmldoc.Attr{
			{
				Name:  "InstanceId",
				Value: strconv.FormatUint(seq.InstanceID, 10),
			},
			{
				Name:  "MessageNumber",
				Value: strconv.FormatUint(seq.MessageNumber, 10),
			},
		},
	}

	if seq.SequenceID != "" {
		elm.Attrs = append(elm.Attrs, xmldoc.Attr{
			Name:  "SequenceID",
			Value: string(seq.SequenceID),
		})
	}

	return elm
}
