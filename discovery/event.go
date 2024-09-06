// MFP - Miulti-Function Printers and scanners toolkit
// Device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Discovery events

package discovery

// Event is the common interface for all events
type Event interface {
	GetID() UnitID
}

var (
	_ = Event(&EventAddUnit{})
	_ = Event(&EventDelUnit{})
	_ = Event(&EventPrinterParameters{})
	_ = Event(&EventScannerParameters{})
	_ = Event(&EventAddEndpoints{})
	_ = Event(&EventDelEndpoints{})
)

// EventAddUnit generated when new print or scan unit is discovered.
type EventAddUnit struct {
	ID UnitID // Unit identity
}

// GetID returns the UnitID this event related to.
func (evnt *EventAddUnit) GetID() UnitID {
	return evnt.ID
}

// EventDelUnit generated when previously announced unit is not
// longer available.
type EventDelUnit struct {
	ID UnitID // Unit identity
}

// GetID returns the UnitID this event related to.
func (evnt *EventDelUnit) GetID() UnitID {
	return evnt.ID
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

// GetID returns the UnitID this event related to.
func (evnt *EventPrinterParameters) GetID() UnitID {
	return evnt.ID
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

// GetID returns the UnitID this event related to.
func (evnt *EventScannerParameters) GetID() UnitID {
	return evnt.ID
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

// GetID returns the UnitID this event related to.
func (evnt *EventAddEndpoints) GetID() UnitID {
	return evnt.ID
}

// EventDelEndpoints is generated, when one ore more printer's or scanner's
// endpoints are not longer available.
//
// Backend responsibilities:
//   - Unit MUST exist
//   - The removed endpoints MUST exist.
type EventDelEndpoints struct {
	ID       UnitID   // Unit identity
	Endpoint []string // URLs of removed endpoints
}

// GetID returns the UnitID this event related to.
func (evnt *EventDelEndpoints) GetID() UnitID {
	return evnt.ID
}
