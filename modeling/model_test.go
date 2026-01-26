// MFP - Miulti-Function Printers and scanners toolkit
// Printer and scanner modeling.
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Device model test

package modeling

import (
	"bytes"
	"testing"

	"github.com/OpenPrinting/go-mfp/internal/assert"
	"github.com/OpenPrinting/go-mfp/internal/testutils"
	"github.com/OpenPrinting/go-mfp/proto/escl"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

func TestKyoceraESCLScannerCapabilities(t *testing.T) {
	// Decode Kyocera ScannerCapabilities
	rd := bytes.NewReader(testutils.Kyocera.
		ECOSYS.M2040dn.ESCL.ScannerCapabilities)
	xml, err := xmldoc.Decode(escl.NsMap, rd)
	assert.NoError(err)

	scancaps, err := escl.DecodeScannerCapabilities(xml)
	assert.NoError(err)

	// Create a new, empty Model
	model, err := NewModel()
	assert.NoError(err)

	defer model.Close()

	// Roll over Model.pyExportStruct/Model.pyImportStruct
	obj := model.pyExportStruct(scancaps)
	if err := obj.Err(); err != nil {
		t.Errorf("Model.pyExportStruct: %s", err)
		return
	}

	var scancaps2 *escl.ScannerCapabilities
	err = model.pyImportStruct(&scancaps2, obj)
	if err != nil {
		t.Errorf("Model.pyImportStruct: %s", err)
		return
	}

	diff := testutils.Diff(scancaps, scancaps2)
	if diff != "" {
		t.Errorf("Model.pyExportStruct/Model.pyImportStruct:\n%s", diff)
	}

	// Roll over Model.Write/Model.Read
	buf := &bytes.Buffer{}

	model.SetESCLScanCaps(scancaps)
	err = model.Write(buf)
	if err != nil {
		t.Errorf("Model.Write: %s", err)
	}

	model2, err := NewModel()
	assert.NoError(err)

	defer model2.Close()

	err = model2.Read("test", buf)
	if err != nil {
		t.Errorf("Model.Read: %s", err)
	}

	diff = testutils.Diff(scancaps, scancaps2)
	if diff != "" {
		t.Errorf("Model.Write/Model.Read:\n%s", diff)
	}
}
