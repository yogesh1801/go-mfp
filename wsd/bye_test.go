// MFP - Miulti-Function Printers and scanners toolkit
// WSD core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Bye test

package wsd

import (
	"reflect"
	"testing"

	"github.com/alexpevzner/mfp/xmldoc"
)

// TestBye tests Bye encoding and decoding
func TestBye(t *testing.T) {
	type testData struct {
		bye Bye
		xml xmldoc.Element
	}

	tests := []testData{
		{
			bye: Bye{
				EndpointReference: EndpointReference{
					Address: "urn:uuid:1fccdddc-380e-41df-8d38-b5df20bc47ef",
				},
			},

			xml: xmldoc.WithChildren(NsDiscovery+":Bye",
				xmldoc.WithChildren(NsAddressing+":EndpointReference",
					xmldoc.WithText(NsAddressing+":Address",
						"urn:uuid:1fccdddc-380e-41df-8d38-b5df20bc47ef",
					),
				),
			),
		},
	}

	for _, test := range tests {
		xml := test.bye.ToXML()
		if !reflect.DeepEqual(xml, test.xml) {
			t.Errorf("ToXML:\nexpected: %s\npresent:  %s\n",
				test.xml.EncodeString(NsMap),
				xml.EncodeString(NsMap))
		}

		bye, err := DecodeBye(xml)
		if err != nil {
			t.Errorf("DecodeBye: %s", err)
			continue
		}

		if !reflect.DeepEqual(bye, test.bye) {
			t.Errorf("DecodeBye:\n"+
				"expected: %#v\npresent:  %#v\n",
				test.bye, bye)
		}
	}
}

// TestByeDecodeErrors additionally tests Bye decoding errors
func TestByeDecodeErrors(t *testing.T) {
	type testData struct {
		xml  xmldoc.Element
		estr string
	}

	tests := []testData{
		{
			xml: xmldoc.WithChildren(NsDiscovery+":Bye",
				xmldoc.WithChildren(NsAddressing+":EndpointReference",
					xmldoc.WithText(NsAddressing+":Address",
						"urn:uuid:1fccdddc-380e-41df-8d38-b5df20bc47ef",
					),
				),
			),
		},

		{
			xml: xmldoc.WithChildren(NsDiscovery + ":Bye"),

			estr: "/d:Bye/a:EndpointReference: missed",
		},
	}

	for _, test := range tests {
		_, err := DecodeBye(test.xml)
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
