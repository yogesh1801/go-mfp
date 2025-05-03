// MFP - Miulti-Function Printers and scanners toolkit
// WSD core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// ProbeMatches message body

package wsd

import (
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// ProbeMatches represents a protocol ProbeMatches message.
//
// The matching devices respond with the ProbeMatches to the received
// [Probe] solicitation.
//
// Please notice, if device matches multiple [Types] it may respond with
// either separate [ProbeMatch] for each type, of with combined ProbeMatch
// for all (most, but not all, device prefer the second option).
type ProbeMatches struct {
	ProbeMatch []ProbeMatch
}

// ProbeMatch represents a single Probe match.
type ProbeMatch struct {
	EndpointReference EndpointReference // Stable identifier of the device
	Types             Types             // Device types
	XAddrs            XAddrs            // Transport addresses (URLs)
	MetadataVersion   uint64            // Incremented when metadata changes
}

// DecodeProbeMatches decodes [ProbeMatches] from the XML tree
func DecodeProbeMatches(root xmldoc.Element) (pm ProbeMatches, err error) {
	defer func() { err = xmldoc.XMLErrWrap(root, err) }()

	const name = NsDiscovery + ":ProbeMatch"
	for _, chld := range root.Children {
		if chld.Name == name {
			var ann Announce
			ann, err = decodeAnnounce(chld)
			if err != nil {
				return
			}

			pm.ProbeMatch = append(pm.ProbeMatch, ProbeMatch(ann))
		}
	}

	return
}

// Action returns [Action] to be used with the [ProbeMatches] message
func (ProbeMatches) Action() Action {
	return ActProbeMatches
}

// Announces returns payload of the ProbeMatches message as a slice
// of the [Announce] structures.
func (pm ProbeMatches) Announces() []Announce {
	if pm.ProbeMatch == nil {
		return nil
	}

	anns := make([]Announce, len(pm.ProbeMatch))
	for i := range anns {
		anns[i] = Announce(pm.ProbeMatch[i])
	}

	return anns
}

// ToXML generates XML tree for the message body
func (pm ProbeMatches) ToXML() xmldoc.Element {
	elm := xmldoc.Element{
		Name: NsDiscovery + ":ProbeMatches",
	}

	for _, match := range pm.ProbeMatch {
		chld := Announce(match).ToXML(NsDiscovery + ":ProbeMatch")
		elm.Children = append(elm.Children, chld)
	}

	return elm
}

// MarkUsedNamespace marks [xmldoc.Namespace] entries used by
// data elements within the message body, if any.
//
// This function should not care about Namespace entries, used
// by XML tags: they are handled automatically.
func (pm ProbeMatches) MarkUsedNamespace(ns xmldoc.Namespace) {
	for _, match := range pm.ProbeMatch {
		match.Types.MarkUsedNamespace(ns)
	}
}
