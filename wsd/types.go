// MFP - Miulti-Function Printers and scanners toolkit
// WSD core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Device types, for discovery

package wsd

import (
	"strings"

	"github.com/alexpevzner/mfp/xmldoc"
)

// Types represents set of device types, for discovery
type Types int

// Known types
const (
	TypeDevice Types = 1 << iota
	TypePrinter
	TypeScanner
)

// DecodeTypes decodes [Types] from the XML tree
func DecodeTypes(root xmldoc.Element) (types Types, err error) {
	names := strings.Fields(root.Text)

	for _, n := range names {
		switch n {
		case "devprof:Device":
			types |= TypeDevice
		case "print:PrintDeviceType":
			types |= TypePrinter
		case "scan:ScanDeviceType":
			types |= TypeScanner
		}
	}

	return
}

// String returns text representation for [Types].
// The representation is compatible for XML.
func (types Types) String() string {
	names := make([]string, 0, 3)

	if types&TypeDevice != 0 {
		names = append(names, "devprof:Device")
	}

	if types&TypePrinter != 0 {
		names = append(names, "print:PrintDeviceType")
	}

	if types&TypeScanner != 0 {
		names = append(names, "scan:ScanDeviceType")
	}

	return strings.Join(names, " ")
}

// ToXML generates XML tree for the Types
func (types Types) ToXML() xmldoc.Element {
	elm := xmldoc.Element{
		Name: NsDiscovery + ":Types",
		Text: types.String(),
	}

	return elm
}

// MarkUsedNamespace marks [xmldoc.Namespace] entries used by
// data elements within the message body, if any.
func (types Types) MarkUsedNamespace(ns xmldoc.Namespace) {
	for i := range ns {
		ent := &ns[i]
		var used bool

		switch ent.Prefix {
		case "devprof":
			used = types&TypeDevice != 0
		case "print":
			used = types&TypePrinter != 0
		case "scan":
			used = types&TypeScanner != 0
		}

		ent.Used = ent.Used || used
	}
}
