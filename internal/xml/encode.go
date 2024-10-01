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
func Encode(w io.Writer, elements []*Element) error {
	return encode(w, elements, true, "")
}

// EncodeIndent writes XML into [io.Writer] in the indented form.
func EncodeIndent(w io.Writer, elements []*Element, indent string) error {
	return encode(w, elements, false, indent)
}

// encode is the internal function that implements XML encoder.
func encode(w io.Writer, elements []*Element,
	compact bool, indent string) error {

	// Create xml.Encoder
	encoder := xml.NewEncoder(w)
	if !compact {
		encoder.Indent("", indent)
	}

	// Write XML version
	tok := xml.ProcInst{Target: "xml", Inst: []byte(`version="1.0"`)}
	encoder.EncodeToken(tok)

	// Write NL after version, if pretty-printing.
	// We have to do it manually with Go stdlib
	if !compact {
		encoder.EncodeToken(xml.CharData("\n"))
	}

	// Recursively encode all elements
	encodeRecursive(encoder, elements)

	// Write NL after XML body
	if !compact {
		encoder.EncodeToken(xml.CharData("\n"))
	}

	// And finally, we are done!
	return encoder.Flush()
}

// encodeRecursive recursively encodes XMS elements
func encodeRecursive(encoder *xml.Encoder, elements []*Element) error {
	for _, elm := range elements {
		var tok xml.Token
		var err error

		name := xml.Name{Space: "", Local: elm.Name}

		tok = xml.StartElement{Name: name}
		err = encoder.EncodeToken(tok)
		if err != nil {
			return err
		}

		text := strings.TrimSpace(elm.Text)
		if text != "" {
			tok = xml.CharData(text)
			err = encoder.EncodeToken(tok)
			if err != nil {
				return err
			}
		}

		if len(elm.Children) != 0 {
			err = encodeRecursive(encoder, elm.Children)
			if err != nil {
				return err
			}
		}

		tok = xml.EndElement{Name: name}
		err = encoder.EncodeToken(tok)
		if err != nil {
			return err
		}
	}

	return nil
}
