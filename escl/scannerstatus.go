// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scanner status

package escl

// ScannerStatus represents the scanner status.
//
// eSCL Technical Specification, 9.
type ScannerStatus struct {
	Version  Version      // eSCL protocol version
	State    ScannerState // Overall scanner state
	ADFState ADFState     // ADF state
	Jobs     []JobInfo    // State of particular jobs
}
