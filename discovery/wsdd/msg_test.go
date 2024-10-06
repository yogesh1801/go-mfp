// MFP - Miulti-Function Printers and scanners toolkit
// WSD device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Protocol messages test

package wsdd

import (
	"os"
	"testing"
)

func TestMsg(t *testing.T) {
	hdr := msgHdr{
		Action:    actHello,
		MessageID: "urn:uuid:34ee82d9-12b0-92d5-1d3e-011f30c09984",
		To:        msgToDiscovery,
	}

	hello := msgHello{
		Address: "uuid:4509a320-00a0-008f-00b6-002507510eca",
		Types: []string{
			"scan:ScannerServiceType",
			"print:PrinterServiceType",
			"wsdp:Device",
		},
		XAddrs: []string{
			"http://192.168.1.102:5358/DeviceService/",
			"http://[fe80::217:c8ff:fe7b:6a91]:5358/DeviceService/",
		},
		MetadataVersion: 1,
	}

	bye := msgBye{
		Address: "uuid:4509a320-00a0-008f-00b6-002507510eca",
	}

	println()
	msg{hdr, hello}.ToXML().EncodeIndent(os.Stdout, msgNsMap, "  ")

	println()
	hdr.Action = actBye
	msg{hdr, bye}.ToXML().EncodeIndent(os.Stdout, msgNsMap, "  ")
}
