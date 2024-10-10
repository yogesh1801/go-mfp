// MFP - Miulti-Function Printers and scanners toolkit
// WSD core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Probe test

package wsd

import (
	"reflect"
	"testing"

	"github.com/alexpevzner/mfp/xmldoc"
)

// TestProbe tests Probe encoding and decoding
func TestProbe(t *testing.T) {
	type testData struct {
		probe Probe
		xml   xmldoc.Element
	}

	tests := []testData{
		{
			probe: Probe{
				Types: TypeDevice,
			},

			xml: xmldoc.Element{
				Name: NsDiscovery + ":Probe",
				Children: []xmldoc.Element{
					{
						Name: NsDiscovery + ":Types",
						Text: "devprof:Device",
					},
				},
			},
		},
	}

	for _, test := range tests {
		xml := test.probe.ToXML()
		if !reflect.DeepEqual(xml, test.xml) {
			t.Errorf("ToXML:\nexpected: %s\npresent:  %s\n",
				test.xml.EncodeString(NsMap),
				xml.EncodeString(NsMap))
		}

		probe, err := DecodeProbe(xml)
		if err != nil {
			t.Errorf("DecodeProbe: %s", err)
			continue
		}

		if !reflect.DeepEqual(probe, test.probe) {
			t.Errorf("DecodeProbe:\n"+
				"expected: %#v\npresent:  %#v\n",
				test.probe, probe)
		}
	}
}

// TestProbeDecodeErrors additionally tests Probe decoding errors
func TestProbeDecodeErrors(t *testing.T) {
	type testData struct {
		xml  xmldoc.Element
		estr string
	}

	tests := []testData{
		{
			xml: xmldoc.Element{
				Name: NsDiscovery + ":Probe",
				Children: []xmldoc.Element{
					{
						Name: NsDiscovery + ":Types",
						Text: "devprof:Device",
					},
				},
			},
		},

		{
			xml: xmldoc.Element{
				Name: NsDiscovery + ":Probe",
			},

			estr: "/d:Probe/d:Types: missed",
		},
	}

	for _, test := range tests {
		_, err := DecodeProbe(test.xml)
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
