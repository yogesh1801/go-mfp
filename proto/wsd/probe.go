// MFP - Miulti-Function Printers and scanners toolkit
// WSD core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Probe message body

package wsd

import (
	"github.com/alexpevzner/mfp/util/xmldoc"
)

// Probe represents a protocol Probe message.
//
// This message usually sent as UDP multicast to the
// 239.255.255.250:3702 or [ff02::c]:3702 in order to solicit
// devices, that match the [Probe.Types] to respond with the
// [ProbeMatches] message.
type Probe struct {
	Types Types // Device types sender searched for
}

// DecodeProbe decodes [Probe] from the XML tree
func DecodeProbe(root xmldoc.Element) (probe Probe, err error) {
	defer func() { err = xmldoc.XMLErrWrap(root, err) }()

	// Lookup message elements
	types := xmldoc.Lookup{Name: NsDiscovery + ":Types", Required: true}

	missed := root.Lookup(&types)
	if missed != nil {
		err = xmldoc.XMLErrMissed(missed.Name)
		return
	}

	// Decode elements
	probe.Types, err = DecodeTypes(types.Elem)

	return
}

// Action returns [Action] to be used with the [Probe] message
func (Probe) Action() Action {
	return ActProbe
}

// ToXML generates XML tree for the message body
func (probe Probe) ToXML() xmldoc.Element {
	elm := xmldoc.Element{
		Name:     NsDiscovery + ":" + "Probe",
		Children: []xmldoc.Element{probe.Types.ToXML()},
	}

	return elm
}

// MarkUsedNamespace marks [xmldoc.Namespace] entries used by
// data elements within the message body, if any.
//
// This function should not care about Namespace entries, used
// by XML tags: they are handled automatically.
func (probe Probe) MarkUsedNamespace(ns xmldoc.Namespace) {
	probe.Types.MarkUsedNamespace(ns)
}
