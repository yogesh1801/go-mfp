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
	"strconv"

	"github.com/alexpevzner/mfp/xmldoc"
)

// decodeMsg decodes message (msg) from the XML tree
func decodeMsg(root xmldoc.Element) (m msg, err error) {
	const (
		rootName = msgNsSOAP + ":" + "Envelope"
		hdrName  = msgNsSOAP + ":" + "Header"
		bodyName = msgNsSOAP + ":" + "Body"
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
	m.Hdr, err = decodeHdr(hdr.Elem)
	if err != nil {
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
	}

	return
}

// decodeHdr decodes message header (msgHdr) from the XML tree
func decodeHdr(root xmldoc.Element) (hdr msgHdr, err error) {
	defer func() { err = xmlErrWrap(root, err) }()

	// Lookup header elements
	Action := xmldoc.Lookup{Name: msgNsAddressing + ":Action", Required: true}
	MessageID := xmldoc.Lookup{Name: msgNsAddressing + ":MessageID", Required: true}
	To := xmldoc.Lookup{Name: msgNsAddressing + ":To", Required: true}
	RelatesTo := xmldoc.Lookup{Name: msgNsAddressing + ":RelatesTo"}
	AppSequence := xmldoc.Lookup{Name: msgNsAddressing + ":AppSequence", Required: true}

	missed := root.Lookup(&Action, &MessageID, &To, &RelatesTo, &AppSequence)
	if missed != nil {
		err = fmt.Errorf("%s: missed", missed.Name)
		return
	}

	// Decode header elements
	hdr.Action, err = decodeAction(Action.Elem)
	if err == nil {
		hdr.MessageID, err = decodeAnyURI(MessageID.Elem)
	}
	if err == nil {
		hdr.To, err = decodeAnyURI(To.Elem)
	}
	if err == nil && RelatesTo.Found {
		hdr.To, err = decodeAnyURI(RelatesTo.Elem)
	}
	if err == nil {
		hdr.AppSequence, err = decodeAppSequence(AppSequence.Elem)
	}

	err = errors.New("not implemented")
	return
}

// decodeHello decodes msgHello from the XML tree
func decodeHello(root xmldoc.Element) (hello msgHello, err error) {
	defer func() { err = xmlErrWrap(root, err) }()
	err = errors.New("not implemented")
	return
}

// decodeBye decodes msgBye from the XML tree
func decodeBye(root xmldoc.Element) (bye msgBye, err error) {
	defer func() { err = xmlErrWrap(root, err) }()
	err = errors.New("not implemented")
	return
}

// decodeAction decodes action, from the XML tree
func decodeAction(root xmldoc.Element) (v action, err error) {
	act := actDecode(root.Text)
	if act != actUnknown {
		return act, nil
	}

	return actUnknown, xmlErrNew(root, "unknown action")
}

// decodeAnyURI decodes anyURI from the XML tree
func decodeAnyURI(root xmldoc.Element) (v anyURI, err error) {
	if root.Text != "" {
		return anyURI(root.Text), nil
	}
	return "", xmlErrNew(root, "invalid URi")
}

// decodeAnyURIAttr decodes anyURI from the XML attribute
func decodeAnyURIAttr(attr xmldoc.Attr) (v anyURI, err error) {
	if attr.Value != "" {
		return anyURI(attr.Value), nil
	}
	return "", xmlErrWrapAttr(attr, errors.New("invalid URi"))
}

// decodeUint64 decodes uint64 from the XML tree
func decodeUint64(root xmldoc.Element) (v uint64, err error) {
	v, err = strconv.ParseUint(root.Text, 10, 64)
	err = xmlErrWrap(root, err)
	return
}

// decodeUint64 decodes uint64 from the XML attribute
func decodeUint64Attr(attr xmldoc.Attr) (v uint64, err error) {
	v, err = strconv.ParseUint(attr.Value, 10, 64)
	err = xmlErrWrapAttr(attr, err)
	return
}

// decodeAppSequence decodes AppSequence from the XML tree
func decodeAppSequence(root xmldoc.Element) (seq msgAppSequence, err error) {
	defer func() { err = xmlErrWrap(root, err) }()

	InstanceID := xmldoc.LookupAttr{
		Name: msgNsAddressing + ":InstanceID", Required: true,
	}
	MessageNumber := xmldoc.LookupAttr{
		Name: msgNsAddressing + ":MessageNumber", Required: true,
	}
	SequenceID := xmldoc.LookupAttr{
		Name: msgNsAddressing + ":SequenceID",
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
		seq.SequenceID, err = decodeAnyURIAttr(SequenceID.Attr)
	}

	return
}
