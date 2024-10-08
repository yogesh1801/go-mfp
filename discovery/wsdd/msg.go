// MFP - Miulti-Function Printers and scanners toolkit
// WSD device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Protocol messages

package wsdd

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/alexpevzner/mfp/internal/xml"
)

// Well-known destinations
const (
	msgToDiscovery = "urn:schemas-xmlsoap-org:ws:2005:04:discovery"
	msgToAnonymous = "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous"
)

// Namespace prefixes
//
// These prefixes used internally and on the generated messages.
const (
	msgNsSOAP       = "s"
	msgNsAddressing = "a"
	msgNsDiscovery  = "d"
	msgNsDevprof    = "devprof"
	msgNsMex        = "mex"
	msgNsPNPX       = "pnpx"
	msgNsScan       = "scan"
	msgNsPrint      = "print"
)

// msgNsMap maps namespace prefixes to URL
var msgNsMap = xml.Namespace{
	// SOAP 1.1
	{Prefix: msgNsSOAP, URL: "http://schemas.xmlsoap.org/soap/envelope"},

	// SOAP 1.2
	{Prefix: msgNsSOAP, URL: "http://www.w3.org/2003/05/soap-envelope"},

	// WSD prefixes
	{Prefix: msgNsAddressing, URL: "http://schemas.xmlsoap.org/ws/2004/08/addressing"},
	{Prefix: msgNsDiscovery, URL: "http://schemas.xmlsoap.org/ws/2005/04/discovery"},
	{Prefix: msgNsDevprof, URL: "http://schemas.xmlsoap.org/ws/2006/02/devprof"},
	{Prefix: msgNsMex, URL: "http://schemas.xmlsoap.org/ws/2004/09/mex"},
	{Prefix: msgNsPNPX, URL: "http://schemas.microsoft.com/windows/pnpx/2005/10"},
	{Prefix: msgNsScan, URL: "http://schemas.microsoft.com/windows/2006/08/wdp/scan"},
	{Prefix: msgNsPrint, URL: "http://schemas.microsoft.com/windows/2006/08/wdp/print"},
}

// msg represents a WSDD protocol message.
type msg struct {
	Hdr  msgHdr  // Message header
	Body msgBody // Message body
}

// ToXML generates XML tree for the message
func (m msg) ToXML() xml.Element {
	elm := xml.Element{
		Name: msgNsSOAP + ":" + "Envelope",
		Children: []xml.Element{
			m.Hdr.ToXML(),
			xml.Element{
				Name:     msgNsSOAP + ":" + "Body",
				Children: []xml.Element{m.Body.ToXML()},
			},
		},
	}

	return elm
}

// msgDecode decodes message from the XML tree
func msgDecode(root xml.Element) (m msg, err error) {
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
	m.Hdr, err = msgHdrDecode(hdr.Elem)
	if err != nil {
		return
	}

	// Decode message body
	switch m.Hdr.Action {
	case actHello:
		m.Body, err = msgHelloDecode(body.Elem)
	case actBye:
		m.Body, err = msgByeDecode(body.Elem)
	default:
		err = fmt.Errorf("%s: unhanded action ", m.Hdr.Action)
	}

	return
}

// msgBody represents a message body.
type msgBody interface {
	ToXML() xml.Element
}

// msgHdr represents a common WSDD message header.
type msgHdr struct {
	Action      action         // Message action
	MessageID   anyURI         // Required: message identifier
	To          anyURI         // Required: message destination
	RelatesTo   anyURI         // Reply-To or similar
	AppSequence msgAppSequence // Message sequence (recv only)
}

// ToXML generates XML tree for the message header
func (hdr msgHdr) ToXML() xml.Element {
	elm := xml.Element{
		Name: msgNsSOAP + ":" + "Header",
		Children: []xml.Element{
			{
				Name: msgNsAddressing + ":" + "Action",
				Text: hdr.Action.Encode(),
			},
			{
				Name: msgNsAddressing + ":" + "MessageID",
				Text: string(hdr.MessageID),
			},
			{
				Name: msgNsAddressing + ":" + "To",
				Text: string(hdr.To),
			},
		},
	}

	if hdr.RelatesTo != "" {
		elm.Children = append(elm.Children,
			xml.Element{
				Name: msgNsAddressing + ":" + "RelatesTo",
				Text: string(hdr.RelatesTo),
			})
	}

	if !hdr.AppSequence.Skip {
		elm.Children = append(elm.Children, hdr.AppSequence.ToXML())
	}

	return elm
}

// msgHdrDecode decodes message header from the XML tree
func msgHdrDecode(root xml.Element) (hdr msgHdr, err error) {
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

// msgAppSequence provides a mechanism that allows a receiver
// to order messages that may have been received out of order.
//
// It is included into the announcement and response messages
// (Hello, Bye, ProbeMatches, and ResolveMatches).
type msgAppSequence struct {
	InstanceID    uint64 // MUST increment on each reboot
	MessageNumber uint64 // MUST increment on each message
	SequenceID    anyURI // Optional: sequence within instance
	Skip          bool   // Skip when sending
}

// ToXML generates XML tree for the AppSequence
func (seq msgAppSequence) ToXML() xml.Element {
	elm := xml.Element{
		Name: msgNsDiscovery + ":" + "AppSequence",
		Attrs: []xml.Attr{
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
		elm.Attrs = append(elm.Attrs, xml.Attr{
			Name:  "SequenceID",
			Value: string(seq.SequenceID),
		})
	}

	return elm
}

// msgAppSequenceDecode decodes AppSequence from the XML tree
func msgAppSequenceDecode(root xml.Element) (seq msgAppSequence, err error) {
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

// msgHello represents body of the protocol Hello message.
// Each device must multicast this message when it enters the network.
type msgHello struct {
	Address         anyURI   // Stable identifier of the device
	Types           []string // Service types
	XAddrs          []string // Transport addresses (URLs)
	MetadataVersion uint64   // Incremented when metadata changes on device
}

// ToXML generates XML tree for the message body
func (hello msgHello) ToXML() xml.Element {
	elm := xml.Element{
		Name: msgNsSOAP + ":" + "Hello",
		Children: []xml.Element{
			{
				Name: msgNsAddressing + ":" + "EndpointReference",
				Children: []xml.Element{
					{
						Name: msgNsAddressing + ":" +
							"Address",
						Text: string(hello.Address),
					},
				},
			},
			{
				Name: msgNsDiscovery + ":" + "MetadataVersion",
				Text: strconv.FormatUint(hello.MetadataVersion, 10),
			},
		},
	}

	if len(hello.Types) != 0 {
		chld := xml.Element{
			Name: msgNsDiscovery + ":" + "Types",
			Text: strings.Join(hello.Types, " "),
		}

		elm.Children = append(elm.Children, chld)
	}

	if len(hello.XAddrs) != 0 {
		chld := xml.Element{
			Name: msgNsDiscovery + ":" + "XAddrs",
			Text: strings.Join(hello.XAddrs, " "),
		}

		elm.Children = append(elm.Children, chld)
	}

	return elm
}

// msgHelloDecode decodes msgHello from the XML tree
func msgHelloDecode(root xml.Element) (hello msgHello, err error) {
	err = errors.New("not implemented")
	return
}

// msgBye represents a protocol Bye message.
// Each device must multicast this message before it enters the network.
type msgBye struct {
	Address anyURI // Stable identifier of the device
}

// ToXML generates XML tree for the message body
func (bye msgBye) ToXML() xml.Element {
	elm := xml.Element{
		Name: msgNsSOAP + ":" + "Bye",
		Children: []xml.Element{
			{
				Name: msgNsAddressing + ":" + "EndpointReference",
				Children: []xml.Element{
					{
						Name: msgNsAddressing + ":" +
							"Address",
						Text: string(bye.Address),
					},
				},
			},
		},
	}

	return elm
}

// msgByeDecode decodes msgBye from the XML tree
func msgByeDecode(root xml.Element) (bye msgBye, err error) {
	err = errors.New("not implemented")
	return
}
