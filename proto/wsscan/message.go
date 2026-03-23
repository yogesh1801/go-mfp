// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// WS-Scan SOAP message

package wsscan

import (
	"bytes"
	"fmt"

	"github.com/OpenPrinting/go-mfp/util/generic"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// Message represents a WS-Scan SOAP message, consisting of
// a [Header] and a body XML element.
type Message struct {
	Header Header
	Body   xmldoc.Element
}

// DecodeMessage decodes a [Message] from the wire representation.
func DecodeMessage(data []byte) (msg Message, err error) {
	root, err := xmldoc.Decode(NsMap, bytes.NewReader(data))
	if err == nil {
		msg, err = decodeMessageXML(root)
	}
	return
}

// decodeMessageXML decodes a [Message] from the XML tree.
func decodeMessageXML(root xmldoc.Element) (msg Message, err error) {
	const (
		rootName = NsSOAP + ":" + "Envelope"
		hdrName  = NsSOAP + ":" + "Header"
		bodyName = NsSOAP + ":" + "Body"
	)

	defer func() { err = xmldoc.XMLErrWrap(root, err) }()

	// Check root element
	if root.Name != rootName {
		err = fmt.Errorf("%s: missed", rootName)
		return
	}

	// Look for Header and Body elements
	hdr := xmldoc.Lookup{Name: hdrName, Required: true}
	body := xmldoc.Lookup{Name: bodyName, Required: true}

	missed := root.Lookup(&hdr, &body)
	if missed != nil {
		err = xmldoc.XMLErrMissed(missed.Name)
		return
	}

	// Decode message header
	msg.Header, err = decodeHeader(hdr.Elem)
	if err != nil {
		return
	}

	// Body is just the raw XML element
	msg.Body = body.Elem

	return
}

// Encode encodes the [Message] into its wire representation.
func (msg Message) Encode() []byte {
	buf := bytes.Buffer{}
	ns := generic.CopySlice(NsMap)
	msg.toXML().Encode(&buf, ns)
	return buf.Bytes()
}

// Format formats the [Message] for logging.
func (msg Message) Format() string {
	ns := generic.CopySlice(NsMap)
	return msg.toXML().EncodeIndentString(ns, "  ")
}

// toXML generates the XML tree for the SOAP envelope.
func (msg Message) toXML() xmldoc.Element {
	return xmldoc.Element{
		Name: NsSOAP + ":" + "Envelope",
		Children: []xmldoc.Element{
			msg.Header.toXML(),
			{
				Name:     NsSOAP + ":" + "Body",
				Children: []xmldoc.Element{msg.Body},
			},
		},
	}
}
