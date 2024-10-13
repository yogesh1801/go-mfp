// MFP - Miulti-Function Printers and scanners toolkit
// WSD core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// announce test

package wsd

import (
	"reflect"
	"strings"
	"testing"

	"github.com/alexpevzner/mfp/xmldoc"
)

// TestAnnounce tests announce encoding and decoding
func TestAnnounce(t *testing.T) {
	type testData struct {
		ann    announce
		xml    xmldoc.Element
		nsused string
	}

	tests := []testData{
		{
			ann: announce{
				EndpointReference: EndpointReference{
					Address: "urn:uuid:1fccdddc-380e-41df-8d38-b5df20bc47ef",
				},
				MetadataVersion: 1,
			},

			xml: xmldoc.Element{
				Name: NsDiscovery + ":Test",
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
			ann: announce{
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

			xml: xmldoc.Element{
				Name: NsDiscovery + ":Test",
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
							"print:PrintDeviceType " +
							"scan:ScanDeviceType",
					},
					{
						Name: NsDiscovery + ":XAddrs",
						Text: "" +
							"http://127.0.0.1/ " +
							"https://[::1]/",
					},
				},
			},

			nsused: "devprof,scan,print",
		},
	}

	for _, test := range tests {
		xml := test.ann.ToXML(NsDiscovery + ":Test")
		if !reflect.DeepEqual(xml, test.xml) {
			t.Errorf("ToXML:\nexpected: %s\npresent:  %s\n",
				test.xml.EncodeString(NsMap),
				xml.EncodeString(NsMap))
		}

		ann, err := decodeAnnounce(xml)
		if err != nil {
			t.Errorf("decodeAnnounce: %s", err)
			continue
		}

		if !reflect.DeepEqual(ann, test.ann) {
			t.Errorf("decodeAnnounce:\n"+
				"expected: %#v\npresent:  %#v\n",
				test.ann, ann)
		}

		ns := NsMap.Clone()
		ann.MarkUsedNamespace(ns)

		nsused := []string{}
		for _, n := range ns {
			if n.Used {
				nsused = append(nsused, n.Prefix)
			}
		}

		nsusedPresent := strings.Join(nsused, ",")

		if test.nsused != nsusedPresent {
			t.Errorf("announce.MarkUsedNamespace:\n"+
				"expected: %s\n"+
				"present:  %s\n",
				test.nsused, nsusedPresent)
		}
	}
}

// TestAnnounceDecodeErrors additionally tests announce decode errors
func TestAnnounceDecodeErrors(t *testing.T) {
	type testData struct {
		xml  xmldoc.Element
		estr string
	}

	tests := []testData{
		{
			xml: xmldoc.Element{
				Name: NsDiscovery + ":Test",
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
			xml: xmldoc.Element{
				Name: NsDiscovery + ":Test",
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
				},
			},

			estr: "/d:Test/d:MetadataVersion: missed",
		},

		{
			xml: xmldoc.Element{
				Name: NsDiscovery + ":Test",
				Children: []xmldoc.Element{
					{
						Name: NsDiscovery + ":MetadataVersion",
						Text: "1",
					},
				},
			},

			estr: "/d:Test/a:EndpointReference: missed",
		},

		{
			xml: xmldoc.Element{
				Name: NsDiscovery + ":Test",
				Children: []xmldoc.Element{
					{
						Name: NsAddressing + ":EndpointReference",
					},
					{
						Name: NsDiscovery + ":MetadataVersion",
						Text: "1",
					},
				},
			},

			estr: "/d:Test/a:EndpointReference/a:Address: missed",
		},
	}

	for _, test := range tests {
		_, err := decodeAnnounce(test.xml)
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
