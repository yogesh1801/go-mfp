// MFP - Miulti-Function Printers and scanners toolkit
// XML mini library
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// XML encoder

package xml

import (
	"encoding/xml"
	"io"
	"strings"
)

// Encode writes XML into [io.Writer] in the compact form.
func (root *Element) Encode(w io.Writer) error {
	return root.encode(w, true, "")
}

// EncodeIndent writes XML into [io.Writer] in the indented form.
func (root *Element) EncodeIndent(w io.Writer, indent string) error {
	return root.encode(w, false, indent)
}

// encode is the internal function that implements XML encoder.
func (root *Element) encode(w io.Writer, compact bool, indent string) error {
	// Create xml.Encoder
	encoder := xml.NewEncoder(w)
	if !compact {
		encoder.Indent("", indent)
	}

	// Write XML version
	tok := xml.ProcInst{Target: "xml", Inst: []byte(`version="1.0"`)}
	encoder.EncodeToken(tok)

	// Write NL after version if pretty-printing.
	// We have to do it manually with Go stdlib.
	if !compact {
		encoder.EncodeToken(xml.CharData("\n"))
	}

	// Recursively encode all elements
	root.encodeRecursive(encoder)

	// Write NL after XML body
	if !compact {
		encoder.EncodeToken(xml.CharData("\n"))
	}

	// And finally, we are done!
	return encoder.Flush()
}

// encodeRecursive recursively encodes XMS element and its children.
func (root *Element) encodeRecursive(encoder *xml.Encoder) error {
	var tok xml.Token
	var err error

	// Write xml.StartElement
	name := xml.Name{Space: "", Local: root.Name}
	attrs := []xml.Attr{}

	for _, attr := range root.Attrs {
		name = xml.Name{Space: "", Local: attr.Name}
		attrs = append(attrs,
			xml.Attr{Name: name, Value: attr.Value})
	}

	tok = xml.StartElement{Name: name, Attr: attrs}

	err = encoder.EncodeToken(tok)
	if err != nil {
		return err
	}

	// Write body
	text := strings.TrimSpace(root.Text)
	if text != "" {
		tok = xml.CharData(text)
		err = encoder.EncodeToken(tok)
		if err != nil {
			return err
		}
	}

	// Write children
	for _, elm := range root.Children {
		err = elm.encodeRecursive(encoder)
		if err != nil {
			return err
		}
	}

	// Write xml.EndElement
	tok = xml.EndElement{Name: name}
	err = encoder.EncodeToken(tok)
	if err != nil {
		return err
	}

	return nil
}
