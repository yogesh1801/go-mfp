// MFP - Miulti-Function Printers and scanners toolkit
// WSD core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// ResolveMatches message body

package wsd

import (
	"github.com/alexpevzner/mfp/xmldoc"
)

// ResolveMatches represents a protocol ResolveMatches message.
//
// The matching devices respond with the ResolveMatches to the received
// [Resolve] solicitation.
//
// Please notice, if device matches multiple [Types] it may respond with
// either separate [ResolveMatch] for each type, of with combined ResolveMatch
// for all (most, but not all, device prefer the second option).
type ResolveMatches struct {
	ResolveMatch []ResolveMatch
}

// ResolveMatch represents a single Resolve match.
type ResolveMatch struct {
	EndpointReference EndpointReference // Stable identifier of the device
	Types             Types             // Device types
	XAddrs            XAddrs            // Transport addresses (URLs)
	MetadataVersion   uint64            // Incremented when metadata changes
}

// DecodeResolveMatches decodes [ResolveMatches] from the XML tree
func DecodeResolveMatches(root xmldoc.Element) (rm ResolveMatches, err error) {
	defer func() { err = xmlErrWrap(root, err) }()

	const name = NsDiscovery + ":ResolveMatch"
	for _, chld := range root.Children {
		if chld.Name == name {
			var ann announce
			ann, err = decodeAnnounce(chld)
			if err != nil {
				return
			}

			rm.ResolveMatch = append(rm.ResolveMatch, ResolveMatch(ann))
		}
	}

	return
}

// Action returns [Action] to be used with the [ResolveMatches] message
func (ResolveMatches) Action() Action {
	return ActResolveMatches
}

// ToXML generates XML tree for the message body
func (rm ResolveMatches) ToXML() xmldoc.Element {
	elm := xmldoc.Element{
		Name: NsDiscovery + ":ResolveMatches",
	}

	for _, match := range rm.ResolveMatch {
		chld := announce(match).ToXML(NsDiscovery + ":ResolveMatch")
		elm.Children = append(elm.Children, chld)
	}

	return elm
}

// MarkUsedNamespace marks [xmldoc.Namespace] entries used by
// data elements within the message body, if any.
//
// This function should not care about Namespace entries, used
// by XML tags: they are handled automatically.
func (rm ResolveMatches) MarkUsedNamespace(ns xmldoc.Namespace) {
	for _, match := range rm.ResolveMatch {
		match.Types.MarkUsedNamespace(ns)
	}
}
