// MFP - Miulti-Function Printers and scanners toolkit
// WSD core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// WSD Message

package wsd

import (
	"fmt"

	"github.com/alexpevzner/mfp/xmldoc"
)

// Msg represents a WSD protocol message.
type Msg struct {
	Header Header // Message header
	Body   Body   // Message body
}

// DecodeMsg decodes [msg] from the XML tree
func DecodeMsg(root xmldoc.Element) (m Msg, err error) {
	const (
		rootName = NsSOAP + ":" + "Envelope"
		hdrName  = NsSOAP + ":" + "Header"
		bodyName = NsSOAP + ":" + "Body"
	)

	defer func() { err = xmlErrWrap(root, err) }()

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
		err = fmt.Errorf("%s: missed", missed.Name)
		return
	}

	// Decode message header
	m.Header, err = DecodeHeader(hdr.Elem)
	if err != nil {
		return
	}

	// Decode message body
	switch m.Header.Action {
	case ActHello:
		m.Body, err = DecodeHello(body.Elem)
	case ActBye:
		m.Body, err = DecodeBye(body.Elem)
	default:
		err = fmt.Errorf("%s: unhanded action ", m.Header.Action)
		return
	}

	return
}

// ToXML generates XML tree for the message
func (m Msg) ToXML() xmldoc.Element {
	elm := xmldoc.Element{
		Name: NsSOAP + ":" + "Envelope",
		Children: []xmldoc.Element{
			m.Header.ToXML(),
			xmldoc.Element{
				Name:     NsSOAP + ":" + "Body",
				Children: []xmldoc.Element{m.Body.ToXML()},
			},
		},
	}

	return elm
}

// MarkUsedNamespace marks [xmldoc.Namespace] entries used by
// data elements within the message body, if any.
//
// This function should not care about Namespace entries, used
// by XML tags: they are handled automatically.
func (m Msg) MarkUsedNamespace(ns xmldoc.Namespace) {
	m.Body.MarkUsedNamespace(ns)
}

// Body represents a message body.
type Body interface {
	ToXML() xmldoc.Element
	MarkUsedNamespace(xmldoc.Namespace)
}
