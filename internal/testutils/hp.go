// MFP - Miulti-Function Printers and scanners toolkit
// Utility functions and data BLOBs for testing
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Test data examples for HP printers

package testutils

import (
	// Import "embed" for its side effects
	_ "embed"
)

// HP contains data samples taken from the Kyocera printers
var HP struct {
	// LaserJet series
	LaserJet struct {
		// M426fdn model
		M426fdn struct {
			// IPP protocol samples
			// ESCL protocol samples
			ESCL struct {
				ScannerCapabilities []byte
				ScannerStatus       []byte
			}
		}
	}
}

func init() {
	HP.LaserJet.M426fdn.ESCL.ScannerCapabilities =
		hpLaserJetM426fdnScannerCapabilities
	HP.LaserJet.M426fdn.ESCL.ScannerStatus =
		hpLaserJetM426fdnScannerStatus
}

//go:embed "data/HP-LaserJet-MFP-M426fdn-ScannerCapabilities.xml"
var hpLaserJetM426fdnScannerCapabilities []byte

//go:embed "data/HP-LaserJet-MFP-M426fdn-ScannerStatus.xml"
var hpLaserJetM426fdnScannerStatus []byte
