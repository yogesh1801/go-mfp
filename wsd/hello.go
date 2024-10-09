// MFP - Miulti-Function Printers and scanners toolkit
// WSD core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Hello message body

package wsd

import (
	"strconv"
	"strings"

	"github.com/alexpevzner/mfp/xmldoc"
)

// Hello represents body of the protocol Hello message.
// Each device must multicast this message when it enters the network.
type Hello struct {
	EndpointReference EndpointReference // Stable identifier of the device
	Types             []string          // Service types
	XAddrs            XAddrs            // Transport addresses (URLs)
	MetadataVersion   uint64            // Incremented when metadata changes
}

// DecodeHello decodes [Hello] from the XML tree
func DecodeHello(root xmldoc.Element) (hello Hello, err error) {
	defer func() { err = xmlErrWrap(root, err) }()

	// Lookup message elements
	EndpointReference := xmldoc.Lookup{
		Name: NsAddressing + ":EndpointReference", Required: true}
	Types := xmldoc.Lookup{
		Name: NsDiscovery + ":" + "Types"}
	XAddrs := xmldoc.Lookup{
		Name: NsDiscovery + ":" + "XAddrs"}
	MetadataVersion := xmldoc.Lookup{
		Name: NsDiscovery + ":" + "MetadataVersion", Required: true}

	missed := root.Lookup(&EndpointReference, &Types,
		&XAddrs, &MetadataVersion)

	if missed != nil {
		err = xmlErrMissed(missed.Name)
		return
	}

	// Decode elements
	hello.EndpointReference, err = DecodeEndpointReference(
		EndpointReference.Elem)

	if err == nil && Types.Found {
		hello.Types = strings.Fields(Types.Elem.Text)
	}

	if err == nil && XAddrs.Found {
		hello.XAddrs, err = DecodeXAddrs(XAddrs.Elem)
	}

	if err == nil {
		hello.MetadataVersion, err = decodeUint64(MetadataVersion.Elem)
	}

	return
}

// ToXML generates XML tree for the message body
func (hello Hello) ToXML() xmldoc.Element {
	elm := xmldoc.Element{
		Name: NsDiscovery + ":" + "Hello",
		Children: []xmldoc.Element{
			hello.EndpointReference.ToXML(
				NsAddressing + ":EndpointReference"),
			{
				Name: NsDiscovery + ":" + "MetadataVersion",
				Text: strconv.FormatUint(hello.MetadataVersion, 10),
			},
		},
	}

	if len(hello.Types) != 0 {
		chld := xmldoc.Element{
			Name: NsDiscovery + ":" + "Types",
			Text: strings.Join(hello.Types, " "),
		}

		elm.Children = append(elm.Children, chld)
	}

	if len(hello.XAddrs) != 0 {
		elm.Children = append(elm.Children, hello.XAddrs.ToXML())
	}

	return elm
}

// MarkUsedNamespace marks [xmldoc.Namespace] entries used by
// data elements within the message body, if any.
//
// This function should not care about Namespace entries, used
// by XML tags: they are handled automatically.
func (hello Hello) MarkUsedNamespace(ns xmldoc.Namespace) {
	for _, name := range hello.Types {
		ns.MarkUsedName(name)
	}
}
