// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// ContentType: specifies the main characteristics of the original document

package wsscan

import (
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// ContentType specifies the main characteristics of the original document.
// The text value is one of: Auto, Text, Photo, Halftone, Mixed.
type ContentType ValWithOptions[ContentTypeValue]

// decodeContentType decodes a ContentType from an XML element.
func decodeContentType(root xmldoc.Element) (ContentType, error) {
	var base ValWithOptions[ContentTypeValue]
	decoded, err := base.decodeValWithOptions(
		root,
		contentTypeDecoder,
	)
	if err != nil {
		return ContentType{}, err
	}
	return ContentType(decoded), nil
}

// toXML converts a ContentType to an XML element.
func (ct ContentType) toXML(name string) xmldoc.Element {
	return ValWithOptions[ContentTypeValue](ct).toXML(name, contentTypeEncoder)
}

// contentTypeDecoder converts a string to a ContentTypeValue.
func contentTypeDecoder(s string) (ContentTypeValue, error) {
	return DecodeContentTypeValue(s), nil
}

// contentTypeEncoder converts a ContentTypeValue to a string.
func contentTypeEncoder(ctv ContentTypeValue) string {
	return ctv.String()
}
