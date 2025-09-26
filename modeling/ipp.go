// MFP - Miulti-Function Printers and scanners toolkit
// Printer and scanner modeling.
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// IPP part of the Model

package modeling

import "github.com/OpenPrinting/go-mfp/proto/ipp"

// SetIPPPrinterAttrs sets the [escl.ScannerCapabilities].
func (model *Model) SetIPPPrinterAttrs(attrs *ipp.PrinterAttributes) {
	model.ippPrinterAttrs = attrs
}

// GetIPPPrinterAttrs returns the [escl.ScannerCapabilities].
func (model *Model) GetIPPPrinterAttrs() *ipp.PrinterAttributes {
	return model.ippPrinterAttrs
}
