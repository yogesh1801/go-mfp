// MFP - Miulti-Function Printers and scanners toolkit
// WSD core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// ResolveMatches test

package wsd

import (
	"reflect"
	"strings"
	"testing"

	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// TestResolveMatches tests ResolveMatches encoding and decoding
func TestResolveMatches(t *testing.T) {
	type testData struct {
		rm     ResolveMatches
		xml    xmldoc.Element
		nsused string
	}

	tests := []testData{
		{
			rm: ResolveMatches{
				ResolveMatch: []ResolveMatch{
					{
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
				},
			},

			xml: xmldoc.Element{
				Name: NsDiscovery + ":ResolveMatches",
				Children: []xmldoc.Element{
					{
						Name: NsDiscovery + ":ResolveMatch",
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
				},
			},

			nsused: "devprof,scan,print",
		},
	}

	for _, test := range tests {
		xml := test.rm.ToXML()
		if !reflect.DeepEqual(xml, test.xml) {
			t.Errorf("ToXML:\nexpected: %s\npresent:  %s\n",
				test.xml.EncodeString(NsMap),
				xml.EncodeString(NsMap))
		}

		rm, err := DecodeResolveMatches(xml)
		if err != nil {
			t.Errorf("DecodeResolveMatches: %s", err)
			continue
		}

		if !reflect.DeepEqual(rm, test.rm) {
			t.Errorf("DecodeResolveMatches:\n"+
				"expected: %#v\npresent:  %#v\n",
				test.rm, rm)
		}

		ns := NsMap.Clone()
		rm.MarkUsedNamespace(ns)

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

// TestResolveMatchesDecodeErrors additionally tests ResolveMatches decode errors
func TestResolveMatchesDecodeErrors(t *testing.T) {
	type testData struct {
		xml  xmldoc.Element
		estr string
	}

	tests := []testData{
		{
			// Empty ResolveMatches is OK
			xml: xmldoc.Element{
				Name: NsDiscovery + ":ResolveMatches",
			},
		},

		{
			// But empty ResolveMatch inside is not
			xml: xmldoc.Element{
				Name: NsDiscovery + ":ResolveMatches",
				Children: []xmldoc.Element{
					{
						Name:     NsDiscovery + ":ResolveMatch",
						Children: []xmldoc.Element{},
					},
				},
			},

			estr: "/d:ResolveMatches/d:ResolveMatch/a:EndpointReference: missed",
		},
	}

	for _, test := range tests {
		_, err := DecodeResolveMatches(test.xml)
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
