// MFP - Miulti-Function Printers and scanners toolkit
// WSD core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Resolve message body

package wsd

import (
	"github.com/alexpevzner/mfp/xmldoc"
)

// Resolve represents a protocol Resolve message.
//
// This message usually sent as UDP multicast to the
// 239.255.255.250:3702 or [ff02::c]:3702 in order to solicit
// devices, that match the [Resolve.EndpointReference] to respond
// with the [ResolveMatches] message.
//
// The typical use case is to obtain XAddrs, if XAddrs are
// missed in the [Hello] or [ProbeMatches] message.
type Resolve struct {
	EndpointReference EndpointReference // Target device
}

// DecodeResolve decodes [Resolve] from the XML tree
func DecodeResolve(root xmldoc.Element) (resolve Resolve, err error) {
	defer func() { err = xmlErrWrap(root, err) }()

	// Lookup message elements
	endpointReference := xmldoc.Lookup{
		Name: NsAddressing + ":EndpointReference", Required: true}

	missed := root.Lookup(&endpointReference)

	if missed != nil {
		err = xmlErrMissed(missed.Name)
		return
	}

	// Decode elements
	resolve.EndpointReference, err = DecodeEndpointReference(
		endpointReference.Elem)

	return
}

// Action returns [Action] to be used with the [Resolve] message
func (Resolve) Action() Action {
	return ActResolve
}

// ToXML generates XML tree for the message body
func (resolve Resolve) ToXML() xmldoc.Element {
	elm := xmldoc.Element{
		Name: NsDiscovery + ":" + "Resolve",
		Children: []xmldoc.Element{
			resolve.EndpointReference.ToXML(
				NsAddressing + ":EndpointReference"),
		},
	}

	return elm

}

// MarkUsedNamespace marks [xmldoc.Namespace] entries used by
// data elements within the message body, if any.
//
// This function should not care about Namespace entries, used
// by XML tags: they are handled automatically.
func (resolve Resolve) MarkUsedNamespace(ns xmldoc.Namespace) {
	// Nothing to mark for Resolve
}
