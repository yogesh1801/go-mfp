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

	"github.com/alexpevzner/mfp/xmldoc"
)

// AnyURI represents anyURI type, per XMS Schema Part 2: Datatypes, 3.2.17
type AnyURI string

// DecodeAnyURI decodes anyURI from the XML tree
func DecodeAnyURI(root xmldoc.Element) (v AnyURI, err error) {
	if root.Text != "" {
		return AnyURI(root.Text), nil
	}
	return "", xmlErrNew(root, "invalid URi")
}

// DecodeAnyURIAttr decodes anyURI from the XML attribute
func DecodeAnyURIAttr(attr xmldoc.Attr) (v AnyURI, err error) {
	if attr.Value != "" {
		return AnyURI(attr.Value), nil
	}
	return "", xmlErrWrapAttr(attr, errors.New("invalid URi"))
}
