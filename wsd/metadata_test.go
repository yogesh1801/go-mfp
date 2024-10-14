// MFP - Miulti-Function Printers and scanners toolkit
// WSD core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Metadata test

package wsd

import (
	"reflect"
	"testing"
)

func TestMetagata(t *testing.T) {
	meta := Metadata{
		ThisDevice: ThisDeviceMetadata{
			FriendlyName: LocalizedStringList{
				{String: "I.Fyodorov FP-0001"},
				{String: "И.Фёдоров ПФ-0001", Lang: "ru-RU"},
			},
			FirmwareVersion: "0.0.1",
			SerialNumber:    "FP-8322017",
		},
		ThisModel: ThisModelMetadata{
			Manufacturer: LocalizedStringList{
				{String: "I.Fyodorov"},
				{String: "И.Фёдоров", Lang: "ru-RU"},
			},
			ManufacturerURL: "http://example.com",
			ModelName: LocalizedStringList{
				{String: "FP-0001"},
				{String: "ПФ-0001", Lang: "ru-RU"},
			},
			ModelNumber:     "FP-0001",
			ModelURL:        "http://example.com/FP-0001",
			PresentationURL: "http://example.com/FP-0001/pres",
		},
		Relationship: Relationship{
			Host: &ServiceMetadata{
				EndpointReference: []EndpointReference{
					{"http://127.0.0.1/"},
				},
			},
			Hosted: []ServiceMetadata{
				{
					EndpointReference: []EndpointReference{
						{"http://127.0.0.1/print"},
					},
					Types:     TypePrinter,
					ServiceID: "uri:b827bd97-925c-4502-a7db-4918a0abfc11",
				},
				{
					EndpointReference: []EndpointReference{
						{"http://127.0.0.1/scan"},
					},
					Types:     TypeScanner,
					ServiceID: "uri:6499d366-62a5-4da9-8c18-5af6eea01f22",
				},
			},
		},
	}

	meta2, err := DecodeMetadata(meta.ToXML())
	if err != nil {
		t.Errorf("DecodeMetadata: %s", err)
		return
	}

	if !reflect.DeepEqual(meta, meta2) {
		t.Errorf("encode/decode mismatch\n"+
			"expected: %s\n"+
			"present:  %s\n",
			meta.ToXML().EncodeString(NsMap),
			meta2.ToXML().EncodeString(NsMap),
		)
	}
}
