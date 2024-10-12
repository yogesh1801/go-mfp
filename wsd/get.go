// MFP - Miulti-Function Printers and scanners toolkit
// WSD core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Get message body

package wsd

import (
	"github.com/alexpevzner/mfp/xmldoc"
)

// Get represents a protocol Get message.
//
// This message is send using HTTP POST via some of XAddrs URLs
// to obtain the device metadata.
//
// This message is trivial and contains no children elements.
type Get struct {
}

// DecodeGet decodes [Get] from the XML tree
func DecodeGet(root xmldoc.Element) (get Get, err error) {
	// Nothing to do
	return
}

// ToXML generates XML tree for the message body
func (get Get) ToXML() xmldoc.Element {
	return xmldoc.Element{}
}

// MarkUsedNamespace marks [xmldoc.Namespace] entries used by
// data elements within the message body, if any.
//
// This function should not care about Namespace entries, used
// by XML tags: they are handled automatically.
func (get Get) MarkUsedNamespace(ns xmldoc.Namespace) {
	// Nothing to mark for Get
}
