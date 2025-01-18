// MFP - Miulti-Function Printers and scanners toolkit
// WSD core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// XMS Schema Part 2: Datatypes

package wsd

import (
	"errors"

	"github.com/alexpevzner/mfp/uuid"
	"github.com/alexpevzner/mfp/xmldoc"
)

// AnyURI represents anyURI type, per XMS Schema Part 2: Datatypes, 3.2.17
type AnyURI string

// DecodeAnyURI decodes anyURI from the XML tree
func DecodeAnyURI(root xmldoc.Element) (v AnyURI, err error) {
	if root.Text != "" {
		return AnyURI(root.Text), nil
	}
	return "", xmldoc.XMLErrNew(root, "invalid URI")
}

// DecodeAnyURIAttr decodes anyURI from the XML attribute
func DecodeAnyURIAttr(attr xmldoc.Attr) (v AnyURI, err error) {
	if attr.Value != "" {
		return AnyURI(attr.Value), nil
	}
	return "", xmldoc.XMLErrWrapAttr(attr, errors.New("invalid URi"))
}

// UUID converts AnyURI into the [uuid.UUID].
//
// If AnyURI is the syntactically correct UUID (for example, in
// the urn:uuid:xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx form), it is
// parsed and returned.
//
// Otherwise, it returns uuid.SHA1(uuid.NameSpaceURL, string(s)).
func (s AnyURI) UUID() uuid.UUID {
	u, err := uuid.Parse(string(s))
	if err == nil {
		return u
	}

	return uuid.SHA1(uuid.NameSpaceURL, string(s))
}
