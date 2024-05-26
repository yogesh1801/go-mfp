// MFP - Miulti-Function Printers and scanners toolkit
// IPP - Internet Printing Protocol implementation
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Tests for PrinterAttributes

package ipp

import (
	_ "embed"
	"testing"

	"github.com/OpenPrinting/goipp"
)

func TestKyoceraM2040dnPrinterAttributes(t *testing.T) {
	msg := testIPPMessage(testdataKyoceraM2040dnPrinterAttributes)

	// Decode printer attributes from real printer
	var pa PrinterAttributes
	err := ippDecodeAttrs(&pa, msg.Printer)

	if err != nil {
		t.Errorf("%s", err)
		return
	}

	// Perform few tests
	if !pa.IsCharsetSupported("utf-8") {
		t.Errorf("TestKyoceraM2040dnPrinterAttributes:" +
			"IsCharsetSupported must be true for utf-8")
	}

	if pa.IsCharsetSupported("xxx-yyy") {
		t.Errorf("TestKyoceraM2040dnPrinterAttributes:" +
			"IsCharsetSupported must be false for xxx-yyy")
	}

	if !pa.IsOperationSupported(goipp.OpGetPrinterAttributes) {
		t.Errorf("TestKyoceraM2040dnPrinterAttributes:"+
			"IsOperationSupported must be true for %s",
			goipp.OpGetPrinterAttributes)
	}

	if pa.IsOperationSupported(goipp.OpCupsGetPrinters) {
		t.Errorf("TestKyoceraM2040dnPrinterAttributes:"+
			"IsOperationSupported must be false for %s",
			goipp.OpCupsGetPrinters)
	}

	// Now encode and compare
	attrs := ippEncodeAttrs(&pa)
	diff := testDiffAttrs(msg.Printer, attrs)
	_ = diff
	//println(diff)
}

//go:embed "testdata/Kyocera-ECOSYS-M2040dn/printer-attributes.ipp"
var testdataKyoceraM2040dnPrinterAttributes []byte
