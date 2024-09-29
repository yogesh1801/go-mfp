// MFP - Miulti-Function Printers and scanners toolkit
// Device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Device units

package discovery

import (
	"fmt"
	"net/netip"
	"strings"

	"github.com/alexpevzner/mfp/uuid"
)

// PrintUnit represents a print unit.
type PrintUnit struct {
	Proto     ServiceProto      // Printing protocol
	Params    PrinterParameters // Printer parameters
	Endpoints []string          // URLs of printer endpoints
}

// ScanUnit represents a scan unit.
type ScanUnit struct {
	Proto     ServiceProto      // Scanning protocol
	Params    ScannerParameters // Scanner parameters
	Endpoints []string          // URLs of printer endpoints
}

// FaxoutUnit represents a fax unit.
type FaxoutUnit struct {
	Proto     ServiceProto      // Faxing protocol
	Params    PrinterParameters // Printer parameters
	Endpoints []string          // URLs of printer endpoints
}

// unit is the internal representation of the PrintUnit, ScanUnit
// or FaxoutUnit
type unit struct {
	id              UnitID       // Unit identity
	MakeModel       string       // Manufacturer + Model
	USBManufacturer string       // I.e., "Hewlett Packard" or "Canon"
	USBModel        string       // Model name
	params          any          // PrinterParameters or ScannerParameters
	endpoints       []string     // Unit endpoints
	addrs           []netip.Addr // Addresses that unit use
}

// Merge merges two units
func (un *unit) Merge(un2 unit) {
	un.endpoints = endpointsMerge(un.endpoints, un2.endpoints)
	un.addrs = addrsMerge(un.addrs, un2.addrs)
}

// Export exports unit ad PrintUnit, ScanUnit or FaxoutUnit
func (un unit) Export() any {
	switch params := un.params.(type) {
	case PrinterParameters:
		// PrinterParameters can be used either with PrintUnit
		// or FaxoutUnit
		switch un.id.SvcType {
		case ServicePrinter:
			return PrintUnit{
				Proto:     un.id.SvcProto,
				Params:    params,
				Endpoints: un.endpoints,
			}
		case ServiceFaxout:
			return FaxoutUnit{
				Proto:     un.id.SvcProto,
				Params:    params,
				Endpoints: un.endpoints,
			}
		}

	case ScannerParameters:
		return ScanUnit{
			Proto:     un.id.SvcProto,
			Params:    params,
			Endpoints: un.endpoints,
		}
	}

	return nil
}

// UnitID contains combination of parameters that identifies a device.
//
// Please note, depending on a discovery protocol being used, not
// all the fields of the following structure will have any sense.
//
// Note also, that device UUID is not necessary the same between
// protocols. Some Canon devices known to use different UUID for
// DNS-SD and WS-Discovery.
//
// The intended fields usage is the following:
//
//	DeviceName - realm-unique device name, in the DNS-SD sense.
//	             E.g., "Kyocera ECOSYS M2040dn",
//	UUID       - device UUID
//	Queue      - Job queue name for units with logical sub-units,
//	             like LPD server with multiple queues
//	Realm      - search realm. Different realms are treated as
//	             independent namespaces.
//	Zone       - allows backend to further divide its namespace
//	             (for example, to split it between network interfaces)
//	Variant    - used to distinguish between logically equivalent
//	             variants of discovered units, that backend sees as
//	             independent instances (for example IP4/IP6, HTTP/HTTPS)
//	SvcType    - service type, printer/scanner/faxout
//	SvcProto   - service protocol, i.e., IPP, LPD, eSCL etc
//	Serial     - device serial number, if appropriate (i.e., for USB)
type UnitID struct {
	DNSSDName string       // DNS-SD name, "" if not available
	UUID      uuid.UUID    // uuid.NilUUID if not available
	Queue     string       // Logical unit within a device
	Realm     SearchRealm  // Search realm
	Zone      string       // Namespace zone within the Realm
	Variant   string       // Finding variant of the same unit
	SvcType   ServiceType  // Service type
	SvcProto  ServiceProto // Service protocol
	USBSerial string       // "" if not avaliable
}

// SameDevice reports if two [UnitID]s belong to the same device.
func (id UnitID) SameDevice(id2 UnitID) bool {
	if id.UUID == id2.UUID {
		return true
	}

	if id.DNSSDName == id2.DNSSDName &&
		id.Realm == id2.Realm &&
		id.Zone == id2.Zone {
		return true
	}

	return false
}

// SameService reports if two [UnitID]s belong to the same service of
// the same device.
func (id UnitID) SameService(id2 UnitID) bool {
	return id.SvcType == id2.SvcType && id.SameDevice(id2)
}

// SameUnit reports if two [UnitID]s belong to the same unit of
// the same device.
func (id UnitID) SameUnit(id2 UnitID) bool {
	return id.Queue == id2.Queue && id.SameService(id2)
}

// MarshalLog dumps [UnitID] as text, for [log.Object].
// It implements [log.Marshaler].
func (id UnitID) MarshalLog() []byte {
	var line string
	lines := make([]string, 0, 6)

	if id.DNSSDName != "" {
		line = fmt.Sprintf("DNSSDName: %q", id.DNSSDName)
		lines = append(lines, line)
	}
	if id.UUID != uuid.NilUUID {
		line = fmt.Sprintf("UUID:      %s", id.UUID)
		lines = append(lines, line)
	}
	if id.Queue != "" {
		line = fmt.Sprintf("Queue:     %q", id.Queue)
		lines = append(lines, line)
	}

	line = fmt.Sprintf("Realm:     %s", id.Realm)
	lines = append(lines, line)

	if id.Zone != "" {
		line = fmt.Sprintf("Zone:      %s", id.Zone)
		lines = append(lines, line)
	}

	if id.Variant != "" {
		line = fmt.Sprintf("Variant:   %s", id.Variant)
		lines = append(lines, line)
	}

	line = fmt.Sprintf("Service:   %s %s", id.SvcProto, id.SvcType)
	lines = append(lines, line)

	if id.USBSerial != "" {
		line = fmt.Sprintf("Serial:    %s", id.USBSerial)
		lines = append(lines, line)
	}

	return []byte(strings.Join(lines, "\n"))
}
