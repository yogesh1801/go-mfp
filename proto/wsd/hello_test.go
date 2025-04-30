// MFP - Miulti-Function Printers and scanners toolkit
// WSD core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Hello methods test

package wsd

import (
	"reflect"
	"strings"
	"testing"

	"github.com/alexpevzner/mfp/util/xmldoc"
)

// TestXHello tests Hello encoding and decoding
func TestHello(t *testing.T) {
	type testData struct {
		hello  Hello
		xml    xmldoc.Element
		nsused string
	}

	tests := []testData{
		{
			hello: Hello{
				EndpointReference: EndpointReference{
					Address: "urn:uuid:1fccdddc-380e-41df-8d38-b5df20bc47ef",
				},
				MetadataVersion: 1,
			},

			xml: xmldoc.WithChildren(NsDiscovery+":Hello",
				xmldoc.WithChildren(NsAddressing+":EndpointReference",
					xmldoc.WithText(NsAddressing+":Address",
						"urn:uuid:1fccdddc-380e-41df-8d38-b5df20bc47ef",
					),
				),
				xmldoc.WithText(NsDiscovery+":MetadataVersion", "1"),
			),
		},

		{
			hello: Hello{
				EndpointReference: EndpointReference{
					Address: "urn:uuid:1fccdddc-380e-41df-8d38-b5df20bc47ef",
				},
				Types: TypeDevice | TypePrinter | TypeScanner,
				XAddrs: XAddrs{
					"http://127.0.0.1/",
					"https://[::1]/",
				},
				MetadataVersion: 1,
			},

			xml: xmldoc.WithChildren(NsDiscovery+":Hello",
				xmldoc.WithChildren(NsAddressing+":EndpointReference",
					xmldoc.WithText(NsAddressing+":Address",
						"urn:uuid:1fccdddc-380e-41df-8d38-b5df20bc47ef",
					),
				),
				xmldoc.WithText(NsDiscovery+":MetadataVersion", "1"),
				xmldoc.WithText(NsDiscovery+":Types",
					"devprof:Device print:PrintDeviceType scan:ScanDeviceType",
				),
				xmldoc.WithText(NsDiscovery+":XAddrs",
					"http://127.0.0.1/ https://[::1]/",
				),
			),

			nsused: "devprof,scan,print",
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

		ns := NsMap.Clone()
		hello.MarkUsedNamespace(ns)

		nsused := []string{}
		for _, n := range ns {
			if n.Used {
				nsused = append(nsused, n.Prefix)
			}
		}

		nsusedPresent := strings.Join(nsused, ",")

		if test.nsused != nsusedPresent {
			t.Errorf("Hello.MarkUsedNamespace:\n"+
				"expected: %s\n"+
				"present:  %s\n",
				test.nsused, nsusedPresent)
		}
	}
}

// TestHelloDecodeErrors additionally tests Hello decode errors
func TestHelloDecodeErrors(t *testing.T) {
	type testData struct {
		xml  xmldoc.Element
		estr string
	}

	tests := []testData{
		{
			xml: xmldoc.WithChildren(NsDiscovery+":Hello",
				xmldoc.WithChildren(NsAddressing+":EndpointReference",
					xmldoc.WithText(NsAddressing+":Address",
						"urn:uuid:1fccdddc-380e-41df-8d38-b5df20bc47ef",
					),
				),
				xmldoc.WithText(NsDiscovery+":MetadataVersion", "1"),
			),
		},

		{
			xml: xmldoc.WithChildren(NsDiscovery+":Hello",
				xmldoc.WithChildren(NsAddressing+":EndpointReference",
					xmldoc.WithText(NsAddressing+":Address",
						"urn:uuid:1fccdddc-380e-41df-8d38-b5df20bc47ef",
					),
				),
			),

			estr: "/d:Hello/d:MetadataVersion: missed",
		},

		{
			xml: xmldoc.WithChildren(NsDiscovery+":Hello",
				xmldoc.WithText(NsDiscovery+":MetadataVersion", "1"),
			),

			estr: "/d:Hello/a:EndpointReference: missed",
		},

		{
			xml: xmldoc.WithChildren(NsDiscovery+":Hello",
				xmldoc.WithChildren(NsAddressing+":EndpointReference"),
				xmldoc.WithText(NsDiscovery+":MetadataVersion", "1"),
			),

			estr: "/d:Hello/a:EndpointReference/a:Address: missed",
		},
	}

	for _, test := range tests {
		_, err := DecodeHello(test.xml)
		estr := ""
		if err != nil {
			estr = err.Error()
		}

		if estr != test.estr {
			t.Errorf("%s\nexpected: %q\npresent:  %q",
				test.xml.EncodeString(NsMap),
				test.estr, estr)
		}
	}
}
