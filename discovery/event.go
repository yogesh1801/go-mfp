// MFP - Miulti-Function Printers and scanners toolkit
// Device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Discovery events

package discovery

// EventAddUnit generated when new print or scan unit is discovered.
type EventAddUnit struct {
	ID UnitID // Unit identity
}

// EventDelUnit generated when previously announced unit is not
// longer available.
type EventDelUnit struct {
	ID UnitID // Unit identity
}

// EventPrinterParameters generated when printer parameters
// become available.
//
// Printer may have a multiple print queues with different parameters.
// All these queues belong to the same Unit and share the same set of
// endpoints.
//
// To distinguish between queues, the 'Queue' parameter is used. Backend
// may leave this parameter empty, if device doesn't support multiple
// queues.
//
// Backend responsibilities:
//   - Unit MUST exist
type EventPrinterParameters struct {
	ID      UnitID            // Unit identity
	Queue   string            // The queue name (optional)
	Printer PrinterParameters // Printer parameters
}

// EventScannerParameters generated when printer parameters
// become available.
//
// Backend responsibilities:
//   - Unit MUST exist
type EventScannerParameters struct {
	ID      UnitID            // Unit identity
	Scanner ScannerParameters // Scanner parameters
}

// EventAddEndpoints is generated, when one or more new printer's
// or scanner's endpoints are discovered.
//
// Backend responsibilities:
//   - Unit MUST exist
//   - The same endpoint MUST NOT be added multiple times.
type EventAddEndpoints struct {
	ID       UnitID   // Unit identity
	Endpoint []string // URLs of added endpoints
}

// EventDelEndpoint is generated, when one ore more printer's or scanner's
// endpoints are not longer available.
//
// Backend responsibilities:
//   - Unit MUST exist
//   - The removed endpoints MUST exist.
type EventDelEndpoint struct {
	ID       UnitID   // Unit identity
	Endpoint []string // URLs of removed endpoints
}
