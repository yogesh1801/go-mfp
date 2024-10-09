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

// TestXHello tests Hello encoding and decoding
func TestHello(t *testing.T) {
	type testData struct {
		hello Hello
		xml   xmldoc.Element
	}

	tests := []testData{
		{
			hello: Hello{
				EndpointReference: EndpointReference{
					Address: "urn:uuid:1fccdddc-380e-41df-8d38-b5df20bc47ef",
				},
				MetadataVersion: 1,
			},

			xml: xmldoc.Element{
				Name: NsDiscovery + ":Hello",
				Children: []xmldoc.Element{
					{
						Name: NsAddressing + ":EndpointReference",
						Children: []xmldoc.Element{
							{
								Name: NsAddressing + ":Address",
								Text: "urn:uuid:1fccdddc-380e-41df-8d38-b5df20bc47ef",
							},
						},
					},
					{
						Name: NsDiscovery + ":MetadataVersion",
						Text: "1",
					},
				},
			},
		},

		{
			hello: Hello{
				EndpointReference: EndpointReference{
					Address: "urn:uuid:1fccdddc-380e-41df-8d38-b5df20bc47ef",
				},
				Types: []string{
					"devprof:Device",
					"scan:ScanDeviceType",
					"print:PrintDeviceType",
				},
				XAddrs: XAddrs{
					"http://127.0.0.1/",
					"https://[::1]/",
				},
				MetadataVersion: 1,
			},

			xml: xmldoc.Element{
				Name: NsDiscovery + ":Hello",
				Children: []xmldoc.Element{
					{
						Name: NsAddressing + ":EndpointReference",
						Children: []xmldoc.Element{
							{
								Name: NsAddressing + ":Address",
								Text: "urn:uuid:1fccdddc-380e-41df-8d38-b5df20bc47ef",
							},
						},
					},
					{
						Name: NsDiscovery + ":MetadataVersion",
						Text: "1",
					},
					{
						Name: NsDiscovery + ":Types",
						Text: "" +
							"devprof:Device " +
							"scan:ScanDeviceType " +
							"print:PrintDeviceType",
					},
					{
						Name: NsDiscovery + ":XAddrs",
						Text: "" +
							"http://127.0.0.1/ " +
							"https://[::1]/",
					},
				},
			},
		},
	}

	for _, test := range tests {
		xml := test.hello.ToXML()
		if !reflect.DeepEqual(xml, test.xml) {
			t.Errorf("ToXML:\nexpected: %s\npresent:  %s\n",
				test.xml.EncodeString(NsMap),
				xml.EncodeString(NsMap))
		}

		hello, err := DecodeHello(xml)
		if err != nil {
			t.Errorf("DecodeHello: %s", err)
			continue
		}

		if !reflect.DeepEqual(hello, test.hello) {
			t.Errorf("DecodeHello:\n"+
				"expected: %#v\npresent:  %#v\n",
				test.hello, hello)
		}

	}
}

// TestXAddrsDecode additionally tests corner cases of
// the XAddrs decoding
func TestXAddrsDecode(t *testing.T) {
	type testData struct {
		xml    xmldoc.Element
		xaddrs XAddrs
	}

	tests := []testData{
		{
			xml: xmldoc.Element{
				Name: NsDiscovery + ":XAddrs",
				Text: "http://127.0.0.1/ " +
					"invalud-url " +
					"http://[::1]/",
			},

			xaddrs: XAddrs{
				"http://127.0.0.1/",
				"http://[::1]/",
			},
		},

		{
			xml: xmldoc.Element{
				Name: NsDiscovery + ":XAddrs",
				Text: "http://127.0.0.1/ " +
					"ftp://localhost/ " +
					"http://[::1]/",
			},

			xaddrs: XAddrs{
				"http://127.0.0.1/",
				"http://[::1]/",
			},
		},
	}

	for _, test := range tests {
		xaddrs, err := DecodeXAddrs(test.xml)
		if err != nil {
			t.Errorf("%s", err)
			continue
		}

		if !reflect.DeepEqual(xaddrs, test.xaddrs) {
			t.Errorf("expected: %s\npresent:  %s\n",
				test.xaddrs, xaddrs)
		}
	}
}
