// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// Test for Format

package wsscan

import (
	"reflect"
	"testing"

	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

func TestFormat_RoundTrip_AllAttributes(t *testing.T) {
	orig := Format(
		ValWithOptions[FormatValue]{
			Text:        JFIF,
			Override:    optional.New(BooleanElement("true")),
			UsedDefault: optional.New(BooleanElement("0")),
		},
	)

	elm := orig.toXML("wscn:Format")
	if elm.Name != "wscn:Format" {
		t.Errorf("expected element name 'wscn:Format', got '%s'",
			elm.Name)
	}
	if elm.Text != "jfif" {
		t.Errorf("expected text 'jfif', got '%s'", elm.Text)
	}
	if len(elm.Attrs) != 2 {
		t.Errorf("expected 2 attributes, got %d: %+v", len(elm.Attrs),
			elm.Attrs)
	}

	decoded, err := decodeFormat(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}
}

func TestFormat_RoundTrip_NoAttributes(t *testing.T) {
	orig := Format(
		ValWithOptions[FormatValue]{
			Text: PNG,
		},
	)

	elm := orig.toXML("wscn:Format")
	if len(elm.Attrs) != 0 {
		t.Errorf("expected no attributes, got %+v", elm.Attrs)
	}

	decoded, err := decodeFormat(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}
	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}
}

func TestFormat_StandardValues(t *testing.T) {
	cases := []struct {
		name     string
		value    FormatValue
		expected string
	}{
		{"dib", DIB, "dib"},
		{"exif", EXIF, "exif"},
		{"jfif", JFIF, "jfif"},
		{"png", PNG, "png"},
		{"pdf-a", PDFA, "pdf-a"},
		{"tiff-single-g4", TIFFSingleG4, "tiff-single-g4"},
		{"tiff-multi-uncompressed", TIFFMultiUncompressed,
			"tiff-multi-uncompressed"},
		{"xps", XPS, "xps"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			f := Format(
				ValWithOptions[FormatValue]{
					Text: c.value,
				},
			)
			elm := f.toXML("wscn:Format")

			if elm.Text != c.expected {
				t.Errorf("expected text '%s', got '%s'", c.expected, elm.Text)
			}

			decoded, err := decodeFormat(elm)
			if err != nil {
				t.Fatalf("decode returned error: %v", err)
			}
			if decoded.Text != c.value {
				t.Errorf("expected value %v, got %v", c.value, decoded.Text)
			}
		})
	}
}

func TestFormat_NoMustHonor(t *testing.T) {
	// Verify that MustHonor is not used (only Override and UsedDefault)
	f := Format(
		ValWithOptions[FormatValue]{
			Text:     JPEG2K,
			Override: optional.New(BooleanElement("false")),
		},
	)

	elm := f.toXML("wscn:Format")

	// Should only have Override attribute, not MustHonor
	if len(elm.Attrs) != 1 {
		t.Errorf("expected 1 attribute, got %d: %+v", len(elm.Attrs), elm.Attrs)
	}

	for _, attr := range elm.Attrs {
		if attr.Name == NsWSCN+":MustHonor" {
			t.Error("MustHonor attribute should not be present in Format")
		}
	}
}

func TestFormat_WithOverride(t *testing.T) {
	orig := Format(
		ValWithOptions[FormatValue]{
			Text:     TIFFSingleJPEGTN2,
			Override: optional.New(BooleanElement("1")),
		},
	)

	elm := orig.toXML("wscn:Format")
	decoded, err := decodeFormat(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}

	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}

	// Verify Override attribute is present
	if decoded.Override == nil {
		t.Error("expected Override attribute to be present")
	}
	if !optional.Get(decoded.Override).Bool() {
		t.Error("expected Override to be true")
	}
}

func TestFormat_WithUsedDefault(t *testing.T) {
	orig := Format(
		ValWithOptions[FormatValue]{
			Text:        TIFFMultiG3MH,
			UsedDefault: optional.New(BooleanElement("true")),
		},
	)

	elm := orig.toXML("wscn:Format")
	decoded, err := decodeFormat(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}

	if !reflect.DeepEqual(orig, decoded) {
		t.Errorf("expected %+v, got %+v", orig, decoded)
	}

	// Verify UsedDefault attribute is present
	if decoded.UsedDefault == nil {
		t.Error("expected UsedDefault attribute to be present")
	}
	if !optional.Get(decoded.UsedDefault).Bool() {
		t.Error("expected UsedDefault to be true")
	}
}

func TestFormat_UnknownValue(t *testing.T) {
	elm := xmldoc.Element{
		Name: "wscn:Format",
		Text: "unknown-format",
	}

	decoded, err := decodeFormat(elm)
	if err != nil {
		t.Fatalf("decode returned error: %v", err)
	}
	if decoded.Text != UnknownFormatValue {
		t.Errorf("expected UnknownFormatValue, got %v", decoded.Text)
	}
}

func TestFormat_InvalidBooleanAttribute(t *testing.T) {
	elm := xmldoc.Element{
		Name: "wscn:Format",
		Text: "pdf-a",
		Attrs: []xmldoc.Attr{
			{Name: NsWSCN + ":Override", Value: "invalid"},
		},
	}

	_, err := decodeFormat(elm)
	if err == nil {
		t.Error("expected error for invalid boolean attribute, got nil")
	}
}
