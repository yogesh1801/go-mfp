// MFP - Miulti-Function Printers and scanners toolkit
// WSD core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Hello message body

package wsd

import (
	"github.com/alexpevzner/mfp/xmldoc"
)

// Hello represents body of the protocol Hello message.
// Each device must multicast this message when it enters the network.
type Hello struct {
	EndpointReference EndpointReference // Stable identifier of the device
	Types             Types             // Device types
	XAddrs            XAddrs            // Transport addresses (URLs)
	MetadataVersion   uint64            // Incremented when metadata changes
}

// DecodeHello decodes [Hello] from the XML tree
func DecodeHello(root xmldoc.Element) (hello Hello, err error) {
	ann, err := decodeAnnounce(root)
	if err == nil {
		hello = Hello(ann)
	}

	return
}

// Action returns [Action] to be used with the [Hello] message
func (Hello) Action() Action {
	return ActHello
}

// FillRequestHeader fills required [Header] fields for [Hello]
// request message.
func (Hello) FillRequestHeader(hdr *Header) {
	hdr.To = ToDiscovery
}

// ToXML generates XML tree for the message body
func (hello Hello) ToXML() xmldoc.Element {
	return announce(hello).ToXML(NsDiscovery + ":Hello")
}

// MarkUsedNamespace marks [xmldoc.Namespace] entries used by
// data elements within the message body, if any.
//
// This function should not care about Namespace entries, used
// by XML tags: they are handled automatically.
func (hello Hello) MarkUsedNamespace(ns xmldoc.Namespace) {
	announce(hello).MarkUsedNamespace(ns)
}
