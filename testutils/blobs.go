// MFP - Miulti-Function Printers and scanners toolkit
// Utility functions and data BLOBs for testing
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Test BLOBs

package testutils

import (
	// Import "embed" for its side effects
	_ "embed"

	"github.com/OpenPrinting/goipp"
)

// Parsed BLOBs
var (
	ParsedKyoceraM2040dnPrinterAttributes *goipp.Message = ippMustParse(
		KyoceraM2040dnPrinterAttributes)
)

// KyoceraM2040dnPrinterAttributes contains Kyocera-ECOSYS-M2040dn
// response to the Get-Printer-Attributes request
//
//go:embed "Kyocera-ECOSYS-M2040dn-Printer-Attributes.ipp"
var KyoceraM2040dnPrinterAttributes []byte
