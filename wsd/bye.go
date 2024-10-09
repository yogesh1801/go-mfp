// MFP - Miulti-Function Printers and scanners toolkit
// WSD core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Package documentation

package wsd

import (
	"errors"

	"github.com/alexpevzner/mfp/xmldoc"
)

// Bye represents a protocol Bye message.
// Each device must multicast this message before it enters the network.
type Bye struct {
	EndpointReference EndpointReference // Stable identifier of the device
}

// DecodeBye decodes [Bye from the XML tree
func DecodeBye(root xmldoc.Element) (bye Bye, err error) {
	defer func() { err = xmlErrWrap(root, err) }()
	err = errors.New("not implemented")
	return
}

// ToXML generates XML tree for the message body
func (bye Bye) ToXML() xmldoc.Element {
	elm := xmldoc.Element{
		Name: NsSOAP + ":" + "Bye",
		Children: []xmldoc.Element{
			bye.EndpointReference.ToXML(
				NsAddressing + ":EndpointReference"),
		},
	}

	return elm
}
