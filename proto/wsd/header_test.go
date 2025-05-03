// MFP - Miulti-Function Printers and scanners toolkit
// WSD core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// WSD Message Header test

package wsd

import (
	"reflect"
	"testing"

	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// TestHeader tests Header encoding and decoding
func TestHeader(t *testing.T) {
	type testData struct {
		hdr Header
		xml xmldoc.Element
	}

	tests := []testData{
		{
			hdr: Header{
				Action:    ActHello,
				MessageID: "urn:uuid:1cf1d308-cb65-494c-9d60-2232c57462e1",
				To:        "urn:uuid:b8310cdf-157f-4e5b-a042-4588f7149ec0",
			},
			xml: xmldoc.WithChildren(NsSOAP+":Header",
				xmldoc.WithText(NsAddressing+":Action", ActHello.Encode()),
				xmldoc.WithText(NsAddressing+":MessageID",
					"urn:uuid:1cf1d308-cb65-494c-9d60-2232c57462e1",
				),
				xmldoc.WithText(NsAddressing+":To",
					"urn:uuid:b8310cdf-157f-4e5b-a042-4588f7149ec0",
				),
			),
		},

		{
			hdr: Header{
				Action:    ActHello,
				MessageID: "urn:uuid:1cf1d308-cb65-494c-9d60-2232c57462e1",
			},
			xml: xmldoc.WithChildren(NsSOAP+":Header",
				xmldoc.WithText(NsAddressing+":Action", ActHello.Encode()),
				xmldoc.WithText(NsAddressing+":MessageID",
					"urn:uuid:1cf1d308-cb65-494c-9d60-2232c57462e1",
				),
			),
		},

		{
			hdr: Header{
				Action:    ActHello,
				MessageID: "urn:uuid:1cf1d308-cb65-494c-9d60-2232c57462e1",
				To:        "urn:uuid:b8310cdf-157f-4e5b-a042-4588f7149ec0",
				ReplyTo: EndpointReference{
					Address: "urn:uuid:02b3be49-ccd5-4074-93ac-313c05050a1f",
				},
				RelatesTo: "urn:uuid:9a6942f8-f5dd-47fc-a4c4-9af559a2bc1a",
				AppSequence: &AppSequence{
					InstanceID:    123456789,
					MessageNumber: 12345,
					SequenceID:    "urn:uuid:b8f0cf58-d83d-4d78-944f-de3e648aaaf0",
				},
			},
			xml: xmldoc.WithChildren(NsSOAP+":Header",
				xmldoc.WithText(NsAddressing+":Action", ActHello.Encode()),
				xmldoc.WithText(NsAddressing+":MessageID",
					"urn:uuid:1cf1d308-cb65-494c-9d60-2232c57462e1",
				),
				xmldoc.WithText(NsAddressing+":To",
					"urn:uuid:b8310cdf-157f-4e5b-a042-4588f7149ec0",
				),
				xmldoc.WithChildren(NsAddressing+":ReplyTo",
					xmldoc.WithText(NsAddressing+":Address",
						"urn:uuid:02b3be49-ccd5-4074-93ac-313c05050a1f",
					),
				),
				xmldoc.WithText(NsAddressing+":RelatesTo",
					"urn:uuid:9a6942f8-f5dd-47fc-a4c4-9af559a2bc1a",
				),
				xmldoc.WithAttrs(NsDiscovery+":AppSequence",
					xmldoc.Attr{Name: "InstanceId", Value: "123456789"},
					xmldoc.Attr{Name: "MessageNumber", Value: "12345"},
					xmldoc.Attr{
						Name:  "SequenceId",
						Value: "urn:uuid:b8f0cf58-d83d-4d78-944f-de3e648aaaf0",
					},
				),
			),
		},
	}

	for _, test := range tests {
		xml := test.hdr.ToXML()
		if !reflect.DeepEqual(xml, test.xml) {
			t.Errorf("ToXML:\nexpected: %s\npresent:  %s\n",
				test.xml.EncodeString(NsMap),
				xml.EncodeString(NsMap))
		}

		hdr, err := DecodeHeader(xml)
		if err != nil {
			t.Errorf("DecodeHeader: %s", err)
			continue
		}

		if !reflect.DeepEqual(hdr, test.hdr) {
			t.Errorf("DecodeHeader:\n"+
				"expected: %#v\npresent:  %#v\n",
				test.hdr, hdr)
		}
	}
}

// TestHeaderErrors tests Header decode errors
func TestHeaderErrors(t *testing.T) {
	type testData struct {
		xml  xmldoc.Element
		estr string
	}

	tests := []testData{
		{
			xml: xmldoc.WithChildren(NsSOAP+":Header",
				xmldoc.WithText(NsAddressing+":Action", ActHello.Encode()),
				xmldoc.WithText(NsAddressing+":MessageID",
					"urn:uuid:1cf1d308-cb65-494c-9d60-2232c57462e1",
				),
				xmldoc.WithText(NsAddressing+":To",
					"urn:uuid:b8310cdf-157f-4e5b-a042-4588f7149ec0",
				),
			),
			estr: "",
		},

		{
			xml: xmldoc.WithChildren(NsSOAP+":Header",
				xmldoc.WithText(NsAddressing+":Action", ActHello.Encode()),
				xmldoc.WithText(NsAddressing+":To",
					"urn:uuid:b8310cdf-157f-4e5b-a042-4588f7149ec0",
				),
			),
			estr: "/s:Header/a:MessageID: missed",
		},

		{
			xml: xmldoc.WithChildren(NsSOAP+":Header",
				xmldoc.WithText(NsAddressing+":Action", ActHello.Encode()),
				xmldoc.WithText(NsAddressing+":MessageID",
					"urn:uuid:1cf1d308-cb65-494c-9d60-2232c57462e1",
				),
				xmldoc.WithText(NsAddressing+":To",
					"urn:uuid:b8310cdf-157f-4e5b-a042-4588f7149ec0",
				),
				xmldoc.WithChildren(NsDiscovery+":AppSequence"),
			),
			estr: "/s:Header/d:AppSequence/d:AppSequence/@InstanceId: missed attribyte",
		},
	}

	for _, test := range tests {
		_, err := DecodeHeader(test.xml)
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
