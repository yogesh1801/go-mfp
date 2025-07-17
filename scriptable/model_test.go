// MFP - Miulti-Function Printers and scanners toolkit
// Print and scam servers with added scriptability.
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Package documentation

package scriptable

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/OpenPrinting/go-mfp/internal/assert"
	"github.com/OpenPrinting/go-mfp/internal/testutils"
	"github.com/OpenPrinting/go-mfp/proto/escl"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

func TestKyoceraESCLScannerCapabilities(t *testing.T) {
	rd := bytes.NewReader(testutils.Kyocera.ECOSYS.M2040dn.ESCL.ScannerCapabilities)
	xml, err := xmldoc.Decode(escl.NsMap, rd)
	assert.NoError(err)

	scancaps, err := escl.DecodeScannerCapabilities(xml)
	assert.NoError(err)

	model, err := NewModel()
	assert.NoError(err)

	obj, err := model.pyExportStruct(scancaps)
	assert.NoError(err)

	err = model.py.Exec("from pprint import pprint", "")
	assert.NoError(err)

	//err = model.pyFormat(obj, os.Stdout)
	//assert.NoError(err)

	var scancaps2 escl.ScannerCapabilities
	err = model.pyImportStruct(&scancaps2, obj)
	assert.NoError(err)

	if !reflect.DeepEqual(*scancaps, scancaps2) {
		t.Errorf("eSCL test failed")
	}
}
