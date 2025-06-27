// MFP - Miulti-Function Printers and scanners toolkit
// WSD core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// ProbeMatches test

package wsd

import (
	"reflect"
	"strings"
	"testing"

	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// TestProbeMatches tests ProbeMatches encoding and decoding
func TestProbeMatches(t *testing.T) {
	type testData struct {
		pm     ProbeMatches
		xml    xmldoc.Element
		nsused string
	}

	tests := []testData{
		{
			pm: ProbeMatches{
				ProbeMatch: []ProbeMatch{
					{
						EndpointReference: EndpointReference{
							Address: "urn:uuid:1fccdddc-380e-41df-8d38-b5df20bc47ef",
						},
						Types: []Type{Device,
							PrinterServiceType, ScannerServiceType},
						XAddrs: XAddrs{
							"http://127.0.0.1/",
							"https://[::1]/",
						},
						MetadataVersion: 1,
					},
				},
			},

			xml: xmldoc.Element{
				Name: NsDiscovery + ":ProbeMatches",
				Children: []xmldoc.Element{
					{
						Name: NsDiscovery + ":ProbeMatch",
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
		xml := test.pm.ToXML()
		if !reflect.DeepEqual(xml, test.xml) {
			t.Errorf("ToXML:\nexpected: %s\npresent:  %s\n",
				test.xml.EncodeString(NsMap),
				xml.EncodeString(NsMap))
		}

		pm, err := DecodeProbeMatches(xml)
		if err != nil {
			t.Errorf("DecodeProbeMatches: %s", err)
			continue
		}

		if !reflect.DeepEqual(pm, test.pm) {
			t.Errorf("DecodeProbeMatches:\n"+
				"expected: %#v\npresent:  %#v\n",
				test.pm, pm)
		}

		ns := NsMap.Clone()
		pm.MarkUsedNamespace(ns)

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

// TestProbeMatchesDecodeErrors additionally tests ProbeMatches decode errors
func TestProbeMatchesDecodeErrors(t *testing.T) {
	type testData struct {
		xml  xmldoc.Element
		estr string
	}

	tests := []testData{
		{
			// Empty ProbeMatches is OK
			xml: xmldoc.Element{
				Name: NsDiscovery + ":ProbeMatches",
			},
		},

		{
			// But empty ProbeMatch inside is not
			xml: xmldoc.Element{
				Name: NsDiscovery + ":ProbeMatches",
				Children: []xmldoc.Element{
					{
						Name:     NsDiscovery + ":ProbeMatch",
						Children: []xmldoc.Element{},
					},
				},
			},

			estr: "/d:ProbeMatches/d:ProbeMatch/a:EndpointReference: missed",
		},
	}

	for _, test := range tests {
		_, err := DecodeProbeMatches(test.xml)
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
