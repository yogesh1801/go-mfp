// MFP - Multi-Function Printers and scanners toolkit
// WS-Scan core protocol
//
// Copyright (C) 2024 and up by Yogesh Singla (yogeshsingla481@gmail.com)
// See LICENSE for license terms and conditions
//
// ValWithOptions: reusable type for elements with
// a value and optional boolean attributes (MustHonor, Override, UsedDefault)

package wsscan

import (
	"github.com/OpenPrinting/go-mfp/util/optional"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// ValWithOptions holds a value and optional boolean attributes.
// This is a generic element for patterns like:
// <wscn:Element
//
//	wscn:MustHonor="true"
//	wscn:Override="false"
//	wscn:UsedDefault="true">
//	    value
//
// </wscn:Element>
// The type parameter T allows the value to be any type (string, int, etc.)
type ValWithOptions[T any] struct {
	Text        T
	MustHonor   optional.Val[BooleanElement]
	Override    optional.Val[BooleanElement]
	UsedDefault optional.Val[BooleanElement]
}

// decodeValWithOptions fills the struct from an XML element.
// The decoder function converts the XML text to the desired type T.
func (t *ValWithOptions[T]) decodeValWithOptions(
	root xmldoc.Element,
	decoder func(string) (T, error),
) (ValWithOptions[T], error) {
	// Decode the text value using the provided decoder
	val, err := decoder(root.Text)
	if err != nil {
		return *t, err
	}
	t.Text = val

	// Decode MustHonor attribute
	if attr, found := root.AttrByName(NsWSCN + ":MustHonor"); found {
		boolVal := BooleanElement(attr.Value)
		if err := boolVal.Validate(); err != nil {
			return *t, err
		}
		t.MustHonor = optional.New(boolVal)
	}

	// Decode Override attribute
	if attr, found := root.AttrByName(NsWSCN + ":Override"); found {
		boolVal := BooleanElement(attr.Value)
		if err := boolVal.Validate(); err != nil {
			return *t, err
		}
		t.Override = optional.New(boolVal)
	}

	// Decode UsedDefault attribute
	if attr, found := root.AttrByName(NsWSCN + ":UsedDefault"); found {
		boolVal := BooleanElement(attr.Value)
		if err := boolVal.Validate(); err != nil {
			return *t, err
		}
		t.UsedDefault = optional.New(boolVal)
	}

	return *t, nil
}

// toXML creates an XML element from the struct.
// The encoder function converts the value of type T to a string.
func (t ValWithOptions[T]) toXML(
	name string,
	encoder func(T) string,
) xmldoc.Element {
	elm := xmldoc.Element{Name: name, Text: encoder(t.Text)}
	var attrs []xmldoc.Attr

	// Add MustHonor attribute if present
	if t.MustHonor != nil {
		attrs = append(attrs, xmldoc.Attr{
			Name:  NsWSCN + ":MustHonor",
			Value: string(optional.Get(t.MustHonor)),
		})
	}

	// Add Override attribute if present
	if t.Override != nil {
		attrs = append(attrs, xmldoc.Attr{
			Name:  NsWSCN + ":Override",
			Value: string(optional.Get(t.Override)),
		})
	}

	// Add UsedDefault attribute if present
	if t.UsedDefault != nil {
		attrs = append(attrs, xmldoc.Attr{
			Name:  NsWSCN + ":UsedDefault",
			Value: string(optional.Get(t.UsedDefault)),
		})
	}

	if len(attrs) > 0 {
		elm.Attrs = attrs
	}

	return elm
}
