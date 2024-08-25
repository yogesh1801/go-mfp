// MFP - Miulti-Function Printers and scanners toolkit
// Device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Discovery events

package discovery

// EventAddPrintUnit generated when new printer is discovered.
type EventAddPrintUnit struct {
	ID      UnitID            // Unit identity
	Printer PrinterParameters // Unit parameters
}

// EventDelPrintUnit generated when printer is not longer available.
type EventDelPrintUnit struct {
	ID UnitID // Unit identity
}

// EventAddScanUnit generated when new scanner is discovered.
type EventAddScanUnit struct {
	ID      UnitID            // Unit identity
	Scanner ScannerParameters // Unit parameters
}

// EventDelScanUnit generated when scanner is not longer available.
type EventDelScanUnit struct {
	ID UnitID // Unit identity
}

// EventAddEndpoints is generated, when one or more new printer's
// or scanner's endpoints are discovered.
type EventAddEndpoints struct {
	ID       UnitID   // Unit identity
	Endpoint []string // URLs of added endpoints
}

// EventDelEndpoint is generated, when one ore more printer's or scanner's
// endpoints are not longer available.
type EventDelEndpoint struct {
	ID       UnitID   // Unit identity
	Endpoint []string // URLs of removed endpoints
}
