// MFP - Miulti-Function Printers and scanners toolkit
// WSD core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Message Body interfaces

package wsd

import "github.com/alexpevzner/mfp/xmldoc"

// Body represents a message body.
//
// Body can be one of the following types:
//   - [Bye]
//   - [Get]
//   - [GetResponse]
//   - [Hello]
//   - [Probe]
//   - [ProbeMatches]
//   - [Resolve]
//   - [ResolveMatches]
type Body interface {
	// Action returns [Action] to be used when sending message
	// with this Body.
	Action() Action

	// ToXML encodes Body into the XML tree.
	ToXML() xmldoc.Element

	// MarkUsedNamespace marks [xmldoc.Namespace] used by
	// encoding of this body.
	MarkUsedNamespace(xmldoc.Namespace)
}

// AnnouncesBody represents a message [Body] that contains device
// announces.
//
// AnnouncesBody can be one of the following types:
//   - [Hello]
//   - [ProbeMatches]
//   - [ResolveMatches]
type AnnouncesBody interface {
	Body
	Announces() []Announce
}
