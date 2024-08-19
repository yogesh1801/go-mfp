// MFP - Miulti-Function Printers and scanners toolkit
// Device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Discovery events

package discovery

// EventAddPrinter generated when new printer is discovered.
type EventAddPrinter struct {
	Printer PrinterInfo
}

// EventDelPrinter generated when printer is not longer available.
type EventDelPrinter struct {
	ID DeviceID
}

// EventAddScanner generated when new scanner is discovered.
type EventAddScanner struct {
	Scanner ScannerInfo
}

// EventDelScanner generated when scanner is not longer available.
type EventDelScanner struct {
	ID DeviceID
}

// EventAddEndpoint is generated, when new printer's or scanner's
// endpoint is descovered.
type EventAddEndpoint struct {
	ID       DeviceID
	Endpoint string // In URI syntax
}

// EventDelEndpoint is generated, when printer's or scanner's
// endpoint is not longer available.
type EventDelEndpoint struct {
	ID       DeviceID
	Endpoint string // In URI syntax
}
