// MFP - Miulti-Function Printers and scanners toolkit
// WSD core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Resolve test

package wsd

import (
	"reflect"
	"strings"
	"testing"

	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// TestResolve tests Resolve encoding and decoding
func TestResolve(t *testing.T) {
	type testData struct {
		resolve Resolve
		xml     xmldoc.Element
		nsused  string
	}

	tests := []testData{
		{
			resolve: Resolve{
				EndpointReference: EndpointReference{
					Address: "urn:uuid:1fccdddc-380e-41df-8d38-b5df20bc47ef",
				},
			},

			xml: xmldoc.Element{
				Name: NsDiscovery + ":Resolve",
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

			nsused: "",
		},
	}

	for _, test := range tests {
		xml := test.resolve.ToXML()
		if !reflect.DeepEqual(xml, test.xml) {
			t.Errorf("ToXML:\nexpected: %s\npresent:  %s\n",
				test.xml.EncodeString(NsMap),
				xml.EncodeString(NsMap))
		}

		resolve, err := DecodeResolve(xml)
		if err != nil {
			t.Errorf("DecodeResolve: %s", err)
			continue
		}

		if !reflect.DeepEqual(resolve, test.resolve) {
			t.Errorf("DecodeResolve:\n"+
				"expected: %#v\npresent:  %#v\n",
				test.resolve, resolve)
		}

		ns := NsMap.Clone()
		resolve.MarkUsedNamespace(ns)

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

// TestResolveDecodeErrors additionally tests Resolve decoding errors
func TestResolveDecodeErrors(t *testing.T) {
	type testData struct {
		xml  xmldoc.Element
		estr string
	}

	tests := []testData{
		{
			xml: xmldoc.Element{
				Name: NsDiscovery + ":Resolve",
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
		},

		{
			xml: xmldoc.Element{
				Name: NsDiscovery + ":Resolve",
			},

			estr: "/d:Resolve/a:EndpointReference: missed",
		},
	}

	for _, test := range tests {
		_, err := DecodeResolve(test.xml)
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
