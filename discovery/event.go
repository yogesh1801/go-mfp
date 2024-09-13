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
	// Name returns the Event name.
	Name() string

	// GetID returns UnitID this Event related to.
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
//
// Note, if printer has multiple print queues, each queue must be
// announced as a separate unit with the separate ID. Backend may use
// UnitID.SubRealm to make IDs of these unit distinguishable.
type EventAddUnit struct {
	ID UnitID // Unit identity
}

// Name returns the Event name.
func (*EventAddUnit) Name() string {
	return "add-unit"
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

// Name returns the Event name.
func (*EventDelUnit) Name() string {
	return "del-unit"
}

// GetID returns the UnitID this event related to.
func (evnt *EventDelUnit) GetID() UnitID {
	return evnt.ID
}

// EventPrinterParameters generated when printer parameters
// become available or updated.
//
// Backend responsibilities:
//   - Unit MUST exist (i.e., previously announced with the
//     EventAddUnit event and not revoked with the EventDelUnit event)
type EventPrinterParameters struct {
	ID      UnitID            // Unit identity
	Printer PrinterParameters // Printer parameters
}

// Name returns the Event name.
func (*EventPrinterParameters) Name() string {
	return "printer-parameters"
}

// GetID returns the UnitID this event related to.
func (evnt *EventPrinterParameters) GetID() UnitID {
	return evnt.ID
}

// EventScannerParameters generated when scanner parameters
// become available or updated.
//
// Backend responsibilities:
//   - Unit MUST exist
type EventScannerParameters struct {
	ID      UnitID            // Unit identity
	Scanner ScannerParameters // Scanner parameters
}

// Name returns the Event name.
func (*EventScannerParameters) Name() string {
	return "scanner-parameters"
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
	ID        UnitID   // Unit identity
	Endpoints []string // URLs of added endpoints
}

// Name returns the Event name.
func (*EventAddEndpoints) Name() string {
	return "add-endpoints"
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
	ID        UnitID   // Unit identity
	Endpoints []string // URLs of removed endpoints
}

// Name returns the Event name.
func (*EventDelEndpoints) Name() string {
	return "del-endpoints"
}

// GetID returns the UnitID this event related to.
func (evnt *EventDelEndpoints) GetID() UnitID {
	return evnt.ID
}
