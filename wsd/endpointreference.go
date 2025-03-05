// MFP - Miulti-Function Printers and scanners toolkit
// WSD core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// EndpointReference

package wsd

import (
	"github.com/alexpevzner/mfp/util/xmldoc"
)

// EndpointReference represents a WSA endpoint address.
type EndpointReference struct {
	Address AnyURI // Endpoint address
}

// DecodeEndpointReference decodes EndpointReference from the XML tree
func DecodeEndpointReference(root xmldoc.Element) (
	ref EndpointReference, err error) {

	defer func() { err = xmldoc.XMLErrWrap(root, err) }()

	address := xmldoc.Lookup{Name: NsAddressing + ":Address", Required: true}
	missed := root.Lookup(&address)
	if missed != nil {
		err = xmldoc.XMLErrMissed(missed.Name)
		return
	}

	ref.Address, err = DecodeAnyURI(address.Elem)

	return
}

// ToXML generates XML tree for the EndpointReference
func (ref EndpointReference) ToXML(name string) xmldoc.Element {
	elm := xmldoc.Element{
		Name: name,
		Children: []xmldoc.Element{
			{
				Name: NsAddressing + ":Address",
				Text: string(ref.Address),
			},
		},
	}

	return elm
}
