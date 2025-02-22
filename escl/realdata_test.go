// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Real-data tests

package escl

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/alexpevzner/mfp/testutils"
	"github.com/alexpevzner/mfp/xmldoc"
)

// TestKyoceraECOSYSM2040dnScannerCapabilities ScannerCapabilities decoding
// for Kyocera ECOSYS M2040dn MFP
func TestKyoceraECOSYSM2040dnScannerCapabilities(t *testing.T) {
	// Parse XML
	data := bytes.NewReader(testutils.
		Kyocera.ECOSYS.M2040dn.ESCL.ScannerCapabilities)
	xml, err := xmldoc.Decode(NsMap, data)
	if err != nil {
		panic(err)
	}

	// Decode ScannerCapabilities
	scancaps, err := DecodeScannerCapabilities(xml)
	if err != nil {
		t.Errorf("%s", err)
		return
	}

	// Verify ScannerCapabilities
	// TODO
	return

	fmt.Printf("%#v", scancaps)
}
