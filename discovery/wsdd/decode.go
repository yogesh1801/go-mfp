// MFP - Miulti-Function Printers and scanners toolkit
// WSD device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// WSDD messages decoding from XML

package wsdd

import (
	"errors"
	"fmt"

	"github.com/alexpevzner/mfp/internal/xml"
)

// decodeMsg decodes message (msg) from the XML tree
func decodeMsg(root xml.Element) (m msg, err error) {
	const (
		rootName = msgNsSOAP + ":" + "Envelope"
		hdrName  = msgNsSOAP + ":" + "Header"
		bodyName = msgNsSOAP + ":" + "Body"
	)

	// Check root element
	if root.Name != rootName {
		err = fmt.Errorf("%s: missed", rootName)
		return
	}

	// Look for Header and Body elements
	hdr := xml.Lookup{Name: hdrName, Required: true}
	body := xml.Lookup{Name: bodyName, Required: true}

	missed := root.Lookup(&hdr, &body)
	if missed != nil {
		err = fmt.Errorf("%s: missed", missed.Name)
		return
	}

	// Decode message header
	m.Hdr, err = decodeHdr(hdr.Elem)
	if err != nil {
		err = xmlErrWrap(hdr.Elem, err)
		return
	}

	// Decode message body
	switch m.Hdr.Action {
	case actHello:
		m.Body, err = decodeHello(body.Elem)
	case actBye:
		m.Body, err = decodeBye(body.Elem)
	default:
		err = fmt.Errorf("%s: unhanded action ", m.Hdr.Action)
		return
	}

	if err != nil {
		err = xmlErrWrap(body.Elem, err)
	}

	return
}

// decodeHdr decodes message header (msgHdr) from the XML tree
func decodeHdr(root xml.Element) (hdr msgHdr, err error) {
	Action := xml.Lookup{Name: msgNsAddressing + ":Action", Required: true}
	MessageID := xml.Lookup{Name: msgNsAddressing + ":MessageID", Required: true}
	To := xml.Lookup{Name: msgNsAddressing + ":To", Required: true}
	RelatesTo := xml.Lookup{Name: msgNsAddressing + ":RelatesTo"}
	AppSequence := xml.Lookup{Name: msgNsAddressing + ":AppSequence", Required: true}

	missed := root.Lookup(&Action, &MessageID, &To, &RelatesTo, &AppSequence)
	if missed != nil {
		err = fmt.Errorf("%s: missed", missed.Name)
		return
	}

	err = errors.New("not implemented")
	return
}

// decodeHello decodes msgHello from the XML tree
func decodeHello(root xml.Element) (hello msgHello, err error) {
	err = errors.New("not implemented")
	return
}

// decodeBye decodes msgBye from the XML tree
func decodeBye(root xml.Element) (bye msgBye, err error) {
	err = errors.New("not implemented")
	return
}

// decodeAppSequence decodes AppSequence from the XML tree
func decodeAppSequence(root xml.Element) (seq msgAppSequence, err error) {
	InstanceID := xml.LookupAttr{
		Name: msgNsAddressing + ":InstanceID", Required: true,
	}
	MessageNumber := xml.LookupAttr{
		Name: msgNsAddressing + ":MessageNumber", Required: true,
	}
	SequenceID := xml.LookupAttr{
		Name: msgNsAddressing + ":SequenceID",
	}

	missed := root.LookupAttrs(&InstanceID, &MessageNumber, &SequenceID)
	if missed != nil {
		err = fmt.Errorf("%s: missed", missed.Name)
		return
	}

	err = errors.New("not implemented")
	return
}
