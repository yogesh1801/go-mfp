// MFP - Miulti-Function Printers and scanners toolkit
// WSD core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Transport addresses (XAddrs) test

package wsd

import (
	"reflect"
	"testing"

	"github.com/alexpevzner/mfp/xmldoc"
)

// TestXAddrs tests XAddrs encoding and decoding
func TestXAddrs(t *testing.T) {
	type testData struct {
		xaddrs XAddrs
		xml    xmldoc.Element
	}

	tests := []testData{
		{
			xaddrs: XAddrs{
				"http://127.0.0.1/",
				"http://[::1]/",
			},

			xml: xmldoc.Element{
				Name: NsDiscovery + ":XAddrs",
				Text: "http://127.0.0.1/ http://[::1]/",
			},
		},

		{
			xaddrs: XAddrs{},

			xml: xmldoc.Element{
				Name: NsDiscovery + ":XAddrs",
			},
		},
	}

	for _, test := range tests {
		xml := test.xaddrs.ToXML()
		if !reflect.DeepEqual(xml, test.xml) {
			t.Errorf("ToXML:\nexpected: %s\npresent: %s\n",
				test.xml.EncodeString(NsMap),
				xml.EncodeString(NsMap))
		}
	}
}
