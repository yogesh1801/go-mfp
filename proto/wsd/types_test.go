// MFP - Miulti-Function Printers and scanners toolkit
// WSD core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Types test

package wsd

import (
	"reflect"
	"strings"
	"testing"

	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// TestTypes tests Types encoding and decoding
func TestTypes(t *testing.T) {
	type testData struct {
		types  Types
		xml    xmldoc.Element
		nsused string
	}

	tests := []testData{
		{
			types: TypeDevice,
			xml: xmldoc.Element{
				Name: NsDiscovery + ":Types",
				Text: "devprof:Device",
			},
			nsused: "devprof",
		},

		{
			types: TypePrinter,
			xml: xmldoc.Element{
				Name: NsDiscovery + ":Types",
				Text: "print:PrintDeviceType",
			},
			nsused: "print",
		},

		{
			types: TypeScanner,
			xml: xmldoc.Element{
				Name: NsDiscovery + ":Types",
				Text: "scan:ScanDeviceType",
			},
			nsused: "scan",
		},

		{
			types: TypeDevice | TypePrinter | TypeScanner,
			xml: xmldoc.Element{
				Name: NsDiscovery + ":Types",
				Text: "devprof:Device print:PrintDeviceType scan:ScanDeviceType",
			},
			nsused: "devprof,scan,print",
		},
	}

	for _, test := range tests {
		xml := test.types.ToXML()
		if !reflect.DeepEqual(xml, test.xml) {
			t.Errorf("ToXML:\nexpected: %s\npresent:  %s\n",
				test.xml.EncodeString(NsMap),
				xml.EncodeString(NsMap))
		}

		types, err := DecodeTypes(xml)
		if err != nil {
			t.Errorf("DecodeBye: %s", err)
			continue
		}

		if !reflect.DeepEqual(types, test.types) {
			t.Errorf("DecodeBye:\n"+
				"expected: %q\npresent:  %q\n",
				test.types, types)
		}

		ns := NsMap.Clone()
		types.MarkUsedNamespace(ns)

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
