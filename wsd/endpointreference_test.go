// MFP - Miulti-Function Printers and scanners toolkit
// WSD core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// EndpointReference test

package wsd

import (
	"reflect"
	"testing"

	"github.com/alexpevzner/mfp/xmldoc"
)

// TestEndpointReference tests EndpointReference
func TestEndpointReference(t *testing.T) {
	type testData struct {
		ref EndpointReference
		xml xmldoc.Element
	}

	tests := []testData{
		{
			ref: EndpointReference{
				Address: "http://[fe80::217:c8ff:fe7b:6a91]:5358/WSDScanner",
			},
			xml: xmldoc.Element{
				Name: NsAddressing + ":EndpointReference",
				Children: []xmldoc.Element{
					{
						Name: NsAddressing + ":Address",
						Text: "http://[fe80::217:c8ff:fe7b:6a91]:5358/WSDScanner",
					},
				},
			},
		},
	}

	for _, test := range tests {
		xml := test.ref.ToXML(NsAddressing + ":EndpointReference")
		if !reflect.DeepEqual(xml, test.xml) {
			t.Errorf("ToXML:\nexpected: %s\npresent: %s\n",
				test.xml.EncodeString(NsMap),
				xml.EncodeString(NsMap))
		}

		ref, err := DecodeEndpointReference(xml)
		if err != nil {
			t.Errorf("DecodeEndpointReference: %s", err)
			continue
		}

		if !reflect.DeepEqual(ref, test.ref) {
			t.Errorf("DecodeAppSequence:\n"+
				"expected: %#v\npresent:  %#v\n",
				test.ref, ref)
		}
	}
}

// TestEndpointReferenceErrors tests EndpointReference decode errors
func TestEndpointReferenceErrors(t *testing.T) {
	type testData struct {
		xml  xmldoc.Element
		estr string
	}

	tests := []testData{
		{
			xml: xmldoc.Element{
				Name: NsAddressing + ":EndpointReference",
				Children: []xmldoc.Element{
					{
						Name: NsAddressing + ":Address",
						Text: "http://[fe80::217:c8ff:fe7b:6a91]:5358/WSDScanner",
					},
				},
			},
			estr: "",
		},

		{
			xml: xmldoc.Element{
				Name: NsAddressing + ":EndpointReference",
				Children: []xmldoc.Element{
					{
						Name: NsAddressing + ":Address",
					},
				},
			},
			estr: "a:EndpointReference/a:Address: invalid URI",
		},

		{
			xml: xmldoc.Element{
				Name: NsAddressing + ":EndpointReference",
			},
			estr: "a:EndpointReference/a:Address: missed",
		},
	}

	for _, test := range tests {
		_, err := DecodeEndpointReference(test.xml)
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
