// MFP - Miulti-Function Printers and scanners toolkit
// WSD core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Transport addresses (XAddrs)

package wsd

import (
	"net/url"
	"strings"

	"github.com/alexpevzner/mfp/xmldoc"
)

// XAddrs represents a collection of transport addresses (URLs)
type XAddrs []string

// DecodeXAddrs decodes [XAddrs] from the XML tree
func DecodeXAddrs(root xmldoc.Element) (xaddrs XAddrs, err error) {
	ss := strings.Fields(root.Text)
	xaddrs = make(XAddrs, 0, len(ss))

	for _, s := range ss {
		u, err := url.ParseRequestURI(s)
		if err != nil {
			// Silently skip invalid URLs
			continue
		}

		if u.Scheme != "http" && u.Scheme != "https" {
			// Silently skip non-HTTP URLs
			continue
		}

		xaddrs = append(xaddrs, s)
	}

	return
}

// ToXML generates XML tree for XAddrs
func (xaddrs XAddrs) ToXML() xmldoc.Element {
	elm := xmldoc.Element{
		Name: NsDiscovery + ":XAddrs",
		Text: strings.Join(xaddrs, " "),
	}

	return elm
}
