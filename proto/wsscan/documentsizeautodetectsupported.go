// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// document size auto detect supported

package wsscan

import (
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// DocumentSizeAutoDetectSupported is a BooleanElement representing whether
// the scan device can detect the size of the original media.
type DocumentSizeAutoDetectSupported BooleanElement

// IsValid returns true if the value is a valid
// DocumentSizeAutoDetectSupported value.
func (dsads DocumentSizeAutoDetectSupported) IsValid() bool {
	return BooleanElement(dsads).IsValid()
}

// Bool returns the boolean value of DocumentSizeAutoDetectSupported.
func (dsads DocumentSizeAutoDetectSupported) Bool() (bool, error) {
	return BooleanElement(dsads).Bool()
}

// toXML converts a DocumentSizeAutoDetectSupported to an XML element.
func (dsads DocumentSizeAutoDetectSupported) toXML(name string) xmldoc.Element {
	return BooleanElement(dsads).toXML(name)
}

// decodeDocumentSizeAutoDetectSupported decodes a
// DocumentSizeAutoDetectSupported from an XML element.
func decodeDocumentSizeAutoDetectSupported(
	root xmldoc.Element,
) (DocumentSizeAutoDetectSupported, error) {
	val, err := decodeBooleanElement(root)
	return DocumentSizeAutoDetectSupported(val), err
}
