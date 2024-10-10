// MFP - Miulti-Function Printers and scanners toolkit
// WSD core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Probe message body

package wsd

import (
	"github.com/alexpevzner/mfp/xmldoc"
)

// Probe represents a protocol Probe message.
// Each device must multicast this message before it enters the network.
type Probe struct {
	Types Types // Device types sender searched for
}

// DecodeProbe decodes [Probe] from the XML tree
func DecodeProbe(root xmldoc.Element) (probe Probe, err error) {
	defer func() { err = xmlErrWrap(root, err) }()

	// Lookup message elements
	Types := xmldoc.Lookup{Name: NsDiscovery + ":Types", Required: true}

	missed := root.Lookup(&Types)
	if missed != nil {
		err = xmlErrMissed(missed.Name)
		return
	}

	// Decode elements
	probe.Types, err = DecodeTypes(Types.Elem)

	return
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
	// Nothing to mark for Probe
}
