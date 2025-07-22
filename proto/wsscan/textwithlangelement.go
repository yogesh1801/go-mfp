// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// TextWithLangElement: reusable type for elements with text and optional xml:lang

package wsscan

import (
	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// TextWithLangElement holds a text value and an optional xml:lang attribute.
type TextWithLangElement struct {
	Text string
	Lang optional.Val[string]
}

// Decode fills the struct from an XML element.
func (t *TextWithLangElement) Decode(root xmldoc.Element) {
	t.Text = root.Text
	if attr, found := root.AttrByName("xml:lang"); found {
		t.Lang = optional.New(attr.Value)
	}
}

// ToXML creates an XML element from the struct.
func (t TextWithLangElement) ToXML(name string) xmldoc.Element {
	elm := xmldoc.Element{Name: name, Text: t.Text}
	lang := optional.Get(t.Lang)
	if lang != "" {
		elm.Attrs = []xmldoc.Attr{{
			Name:  "xml:lang",
			Value: lang,
		}}
	}
	return elm
}
