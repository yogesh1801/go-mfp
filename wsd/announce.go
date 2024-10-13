// MFP - Miulti-Function Printers and scanners toolkit
// WSD core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// announce type (common for Hello, ProbeMatches and ResolveMatches)

package wsd

import (
	"strconv"

	"github.com/alexpevzner/mfp/xmldoc"
)

// announce is the common data structure, used for the [Hello],
// [ProbeMatches] and [ResolveMatches] message.
type announce struct {
	EndpointReference EndpointReference // Stable identifier of the device
	Types             Types             // Device types
	XAddrs            XAddrs            // Transport addresses (URLs)
	MetadataVersion   uint64            // Incremented when metadata changes
}

// decodeAnnounce decodes [announce] from the XML tree
func decodeAnnounce(root xmldoc.Element) (ann announce, err error) {
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
	ann.EndpointReference, err = DecodeEndpointReference(
		EndpointReference.Elem)

	if err == nil && Types.Found {
		ann.Types, err = DecodeTypes(Types.Elem)
	}

	if err == nil && XAddrs.Found {
		ann.XAddrs, err = DecodeXAddrs(XAddrs.Elem)
	}

	if err == nil {
		ann.MetadataVersion, err = decodeUint64(MetadataVersion.Elem)
	}

	return
}

// ToXML generates XML tree for the message body
func (ann announce) ToXML(name string) xmldoc.Element {
	elm := xmldoc.Element{
		Name: name,
		Children: []xmldoc.Element{
			ann.EndpointReference.ToXML(
				NsAddressing + ":EndpointReference"),
			{
				Name: NsDiscovery + ":" + "MetadataVersion",
				Text: strconv.FormatUint(ann.MetadataVersion, 10),
			},
		},
	}

	if ann.Types != 0 {
		elm.Children = append(elm.Children, ann.Types.ToXML())
	}

	if len(ann.XAddrs) != 0 {
		elm.Children = append(elm.Children, ann.XAddrs.ToXML())
	}

	return elm
}

// MarkUsedNamespace marks [xmldoc.Namespace] entries used by
// data elements within the message body, if any.
//
// This function should not care about Namespace entries, used
// by XML tags: they are handled automatically.
func (ann announce) MarkUsedNamespace(ns xmldoc.Namespace) {
	ann.Types.MarkUsedNamespace(ns)
}
