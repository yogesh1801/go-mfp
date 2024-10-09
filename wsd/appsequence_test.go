// MFP - Miulti-Function Printers and scanners toolkit
// WSD core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// AppSequence test

package wsd

import (
	"reflect"
	"testing"

	"github.com/alexpevzner/mfp/xmldoc"
)

// TestAppSequence tests AppSequence
func TestAppSequence(t *testing.T) {
	type testData struct {
		seq AppSequence
		xml xmldoc.Element
	}

	tests := []testData{
		{
			seq: AppSequence{
				InstanceID:    123456789,
				MessageNumber: 123,
			},
			xml: xmldoc.Element{
				Name: NsDiscovery + ":" + "AppSequence",
				Attrs: []xmldoc.Attr{
					{
						Name:  "InstanceId",
						Value: "123456789",
					},
					{
						Name:  "MessageNumber",
						Value: "123",
					},
				},
			},
		},

		{
			seq: AppSequence{
				InstanceID:    987654321,
				MessageNumber: 321,
				SequenceID:    "urn:uuid:2a443ed7-5ee5-498d-a302-73ff91ea9ea0",
			},
			xml: xmldoc.Element{
				Name: NsDiscovery + ":" + "AppSequence",
				Attrs: []xmldoc.Attr{
					{
						Name:  "InstanceId",
						Value: "987654321",
					},
					{
						Name:  "MessageNumber",
						Value: "321",
					},
					{
						Name:  "SequenceId",
						Value: "urn:uuid:2a443ed7-5ee5-498d-a302-73ff91ea9ea0",
					},
				},
			},
		},
	}

	for _, test := range tests {
		xml := test.seq.ToXML()
		if !reflect.DeepEqual(xml, test.xml) {
			t.Errorf("ToXML:\nexpected: %s\npresent: %s\n",
				test.xml.EncodeString(NsMap),
				xml.EncodeString(NsMap))
		}

		seq, err := DecodeAppSequence(xml)
		if err != nil {
			t.Errorf("DecodeAppSequence: %s", err)
			continue
		}

		if !reflect.DeepEqual(seq, test.seq) {
			t.Errorf("DecodeAppSequence:\n"+
				"expected: %#v\npresent:  %#v\n",
				test.seq, seq)
		}
	}
}

// TestAppSequenceDecodeErrors tests AppSequence decode errors
func TestAppSequenceDecodeErrors(t *testing.T) {
	type testData struct {
		xml  xmldoc.Element
		estr string
	}

	tests := []testData{
		{
			xml: xmldoc.Element{
				Name: NsDiscovery + ":" + "AppSequence",
				Attrs: []xmldoc.Attr{
					{
						Name:  "InstanceId",
						Value: "123456789",
					},
					{
						Name:  "MessageNumber",
						Value: "123",
					},
				},
			},
			estr: "",
		},

		{
			xml: xmldoc.Element{
				Name: NsDiscovery + ":" + "AppSequence",
				Attrs: []xmldoc.Attr{
					{
						Name:  "InstanceId",
						Value: "123456789",
					},
				},
			},
			estr: "d:AppSequence/d:AppSequence/MessageNumber: missed attribyte",
		},

		{
			xml: xmldoc.Element{
				Name: NsDiscovery + ":" + "AppSequence",
				Attrs: []xmldoc.Attr{
					{
						Name:  "MessageNumber",
						Value: "123",
					},
				},
			},
			estr: "d:AppSequence/d:AppSequence/InstanceId: missed attribyte",
		},
	}

	for _, test := range tests {
		_, err := DecodeAppSequence(test.xml)
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
