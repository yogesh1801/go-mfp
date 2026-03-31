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
	"io"
	"mime/multipart"
	"net/textproto"

	"github.com/OpenPrinting/go-mfp/abstract"
	"github.com/OpenPrinting/go-mfp/util/generic"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// Message represents a WS-Scan SOAP message, consisting of
// a [Header] and a body XML element.
//
// When File is set, the message carries a binary attachment
// and must be encoded as MTOM/XOP multipart using
// [Message.WriteMTOM]. The FileCID links the xop:Include
// reference in the SOAP body to the file part.
type Message struct {
	Header  Header
	Body    xmldoc.Element
	File    abstract.DocumentFile
	FileCID string
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

// MTOMContentType returns the Content-Type header value for the
// MTOM/XOP multipart encoding using the given boundary and
// envelope Content-ID.
func MTOMContentType(boundary, envelopeCID string) string {
	return fmt.Sprintf(
		`multipart/related;`+
			` type="application/xop+xml";`+
			` boundary="%s";`+
			` start="<%s>";`+
			` start-info="application/soap+xml"`,
		boundary, envelopeCID)
}

// WriteMTOM encodes the message as an MTOM/XOP multipart response,
// with the SOAP envelope as the first part and the binary data from
// [Message.File] as the second part.
//
// The caller must set HTTP headers (using [MTOMContentType]) and
// call WriteHeader before invoking WriteMTOM, because WriteMTOM
// writes directly to w.
//
// The boundary and envelopeCID must match the values used in the
// Content-Type header.
func (msg Message) WriteMTOM(w io.Writer, boundary, envelopeCID string) error {
	// Encode the SOAP envelope
	soapData := msg.Encode()

	// Create multipart writer with the pre-determined boundary
	mw := multipart.NewWriter(w)
	if err := mw.SetBoundary(boundary); err != nil {
		return err
	}

	// Part 1: SOAP envelope
	soapPart, err := mw.CreatePart(textproto.MIMEHeader{
		"Content-Type": {`application/xop+xml;` +
			` charset=UTF-8;` +
			` type="application/soap+xml"`},
		"Content-Transfer-Encoding": {"binary"},
		"Content-Id":                {"<" + envelopeCID + ">"},
	})
	if err != nil {
		return err
	}
	if _, err = soapPart.Write(soapData); err != nil {
		return err
	}

	// Part 2: File data (streamed)
	filePart, err := mw.CreatePart(textproto.MIMEHeader{
		"Content-Type":              {msg.File.Format()},
		"Content-Transfer-Encoding": {"binary"},
		"Content-Id":                {"<" + msg.FileCID + ">"},
	})
	if err != nil {
		return err
	}
	if _, err = io.Copy(filePart, msg.File); err != nil {
		return err
	}

	return mw.Close()
}
