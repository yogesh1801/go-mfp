// MFP - Miulti-Function Printers and scanners toolkit
// Printer and scanner modeling.
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// WSD (WS-Scan) part of Model

package modeling

import (
	"fmt"

	"github.com/OpenPrinting/go-mfp/proto/wsscan"
)

// SetWSDScanCaps sets the [escl.ScannerCapabilities].
func (model *Model) SetWSDScanCaps(caps *wsscan.GetScannerElementsResponse) {
	model.wsdScanCaps = caps
}

// GetWSDScanCaps returns the [escl.ScannerCapabilities].
func (model *Model) GetWSDScanCaps() *wsscan.GetScannerElementsResponse {
	return model.wsdScanCaps
}

// wsdLoad decodes WS-Scan part of model. The model file assumed to
// be preloaded into the Model's Python interpreter (model.py).
func (model *Model) wsdLoad() error {
	// Load wsscan.GetScannerElementsResponse
	obj := model.py.Eval("wsd.caps")
	if err := obj.Err(); err != nil {
		err = fmt.Errorf("wsd.caps: %w", err)
		return err
	}

	if obj.IsNone() {
		return nil
	}

	// Decode wsscan.GetScannerElementsResponse
	var caps *wsscan.GetScannerElementsResponse
	err := model.pyImportStruct(keywordMapWSD, &caps, obj)
	if err != nil {
		err = fmt.Errorf("wsd.caps: %w", err)
		return err
	}

	model.wsdScanCaps = caps
	return nil
}
