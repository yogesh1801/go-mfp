// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scan "content type"

package escl

import "github.com/OpenPrinting/go-mfp/util/xmldoc"

// ContentType is similar to [Intent] but is limited to image processing.
//
// According to the eSCL protocol specification, it does not imply default
// assumptions about other parameters (resolutions, data format, etc…).
// Especially, it SHOULD be used when the Intent isn’t specified. It can be
// omitted when ‘Intent’ is present.
type ContentType int

// Known intents
const (
	UnknownContentType ContentType = iota // Unknown ContentType
	ContentTypePhoto
	ContentTypeText
	ContentTypeTextAndPhoto
	ContentTypeLineArt
	ContentTypeMagazine
	ContentTypeHalftone
	ContentTypeAuto
)

// decodeContentType decodes [ContentType] from the XML tree.
func decodeContentType(root xmldoc.Element) (ct ContentType, err error) {
	return decodeEnum(root, DecodeContentType)
}

// toXML generates XML tree for the [ContentType].
func (ct ContentType) toXML(name string) xmldoc.Element {
	return xmldoc.Element{
		Name: name,
		Text: ct.String(),
	}
}

// String returns a string representation of the [ContentType]
func (ct ContentType) String() string {
	switch ct {
	case ContentTypePhoto:
		return "Photo"
	case ContentTypeText:
		return "Text"
	case ContentTypeTextAndPhoto:
		return "TextAndPhoto"
	case ContentTypeLineArt:
		return "LineArt"
	case ContentTypeMagazine:
		return "Magazine"
	case ContentTypeHalftone:
		return "Halftone"
	case ContentTypeAuto:
		return "Auto"
	}

	return "Unknown"
}

// DecodeContentType decodes [ContentType] out of its XML string representation.
func DecodeContentType(s string) ContentType {
	switch s {
	case "Photo":
		return ContentTypePhoto
	case "Text":
		return ContentTypeText
	case "TextAndPhoto":
		return ContentTypeTextAndPhoto
	case "LineArt":
		return ContentTypeLineArt
	case "Magazine":
		return ContentTypeMagazine
	case "Halftone":
		return ContentTypeHalftone
	case "Auto":
		return ContentTypeAuto
	}

	return UnknownContentType
}
