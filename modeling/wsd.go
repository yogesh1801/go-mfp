// MFP - Miulti-Function Printers and scanners toolkit
// Printer and scanner modeling.
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// WSD (WS-Scan) part of Model

package modeling

import "github.com/OpenPrinting/go-mfp/proto/wsscan"

// SetWSDScanCaps sets the [escl.ScannerCapabilities].
func (model *Model) SetWSDScanCaps(caps *wsscan.GetScannerElementsResponse) {
	model.wsdScanCaps = caps
}

// GetWSDScanCaps returns the [escl.ScannerCapabilities].
func (model *Model) GetWSDScanCaps() *wsscan.GetScannerElementsResponse {
	return model.wsdScanCaps
}
