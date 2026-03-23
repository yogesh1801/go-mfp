// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// AnyURI type for WS-Scan protocol

package wsscan

import (
	"errors"

	"github.com/OpenPrinting/go-mfp/util/uuid"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// AnyURI represents the anyURI type, per XML Schema Part 2: Datatypes,
// 3.2.17.
type AnyURI string

// DecodeAnyURI decodes an [AnyURI] from the XML tree.
func DecodeAnyURI(root xmldoc.Element) (v AnyURI, err error) {
	if root.Text != "" {
		return AnyURI(root.Text), nil
	}
	return "", xmldoc.XMLErrNew(root, "invalid URI")
}

// DecodeAnyURIAttr decodes an [AnyURI] from an XML attribute.
func DecodeAnyURIAttr(attr xmldoc.Attr) (v AnyURI, err error) {
	if attr.Value != "" {
		return AnyURI(attr.Value), nil
	}
	return "", xmldoc.XMLErrWrapAttr(attr, errors.New("invalid URI"))
}

// UUID converts AnyURI into a [uuid.UUID].
//
// If AnyURI is a syntactically correct UUID (for example, in
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
