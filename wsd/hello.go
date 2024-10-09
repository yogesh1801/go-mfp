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
	"strconv"
	"strings"

	"github.com/alexpevzner/mfp/xmldoc"
)

// Hello represents body of the protocol Hello message.
// Each device must multicast this message when it enters the network.
type Hello struct {
	EndpointReference EndpointReference // Stable identifier of the device
	Types             []string          // Service types
	XAddrs            []string          // Transport addresses (URLs)
	MetadataVersion   uint64            // Incremented when metadata changes
}

// DecodeHello decodes [Hello] from the XML tree
func DecodeHello(root xmldoc.Element) (hello Hello, err error) {
	defer func() { err = xmlErrWrap(root, err) }()
	err = errors.New("not implemented")
	return
}

// ToXML generates XML tree for the message body
func (hello Hello) ToXML() xmldoc.Element {
	elm := xmldoc.Element{
		Name: NsSOAP + ":" + "Hello",
		Children: []xmldoc.Element{
			hello.EndpointReference.ToXML(
				NsAddressing + ":EndpointReference"),
			{
				Name: NsDiscovery + ":" + "MetadataVersion",
				Text: strconv.FormatUint(hello.MetadataVersion, 10),
			},
		},
	}

	if len(hello.Types) != 0 {
		chld := xmldoc.Element{
			Name: NsDiscovery + ":" + "Types",
			Text: strings.Join(hello.Types, " "),
		}

		elm.Children = append(elm.Children, chld)
	}

	if len(hello.XAddrs) != 0 {
		chld := xmldoc.Element{
			Name: NsDiscovery + ":" + "XAddrs",
			Text: strings.Join(hello.XAddrs, " "),
		}

		elm.Children = append(elm.Children, chld)
	}

	return elm
}
