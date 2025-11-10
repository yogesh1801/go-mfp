// MFP - Multi-Function Printers and scanners toolkit
// PPD handling (libppd wrapper)
//
// Copyright (C) 2025 and up by GO-MFP authors.
// See LICENSE for license terms and conditions
//
// IPP->PPD conversion

package ppd

import (
	"errors"

	"github.com/OpenPrinting/goipp"
)

// FromIPP generates the PPD file from the IPP printer attributes.
func FromIPP(attrs goipp.Attributes) ([]byte, error) {
	return nil, errors.New("not implemented")
}
