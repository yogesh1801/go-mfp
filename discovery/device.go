// MFP - Miulti-Function Printers and scanners toolkit
// Device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Common device information

package discovery

import (
	"fmt"
	"strings"

	"github.com/alexpevzner/mfp/uuid"
)

// Device consist of the multiple functional units. There are
// two types of units:
//   - [PrintUnit]
//   - [ScanUnit]
//
// Each unit has its unique [UnitID], the combination of parameters,
// that uniquely identifies the unit.
//
// If, due to the peculiarities of the search protocol, the same device
// can appear as several different ones, this is at the search [Backend]
// discretion, either to merge these multiple instances by itself or to
// leave this work up to the discovery system.
//
// If Backend decides to merge by itself, the resulting unit should appear
// as a single unit with merged endpoints. Otherwise, each appearance should
// appear as distinct unit (units with distinct UnitID), and discovery
// subsystem will merge them, if UnitUDs is "similar enough".
type Device struct {
	PrintUnits []PrintUnit
	ScanUnits  []ScanUnit
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
//	DeviceName - unique device name, in the DNS-SD sence.
//	             E.g., "Kyocera ECOSYS M2040dn",
//	UUID       - device UUID
//	Realm      - search realm. Different realms are treated as
//	             independent namespaces.
//	SubRealm   - allows backend to further divide its namespace
//	             (for example, to split it between IP4/IP6)
//	Kind       - specifies device kind (e.g., "IPP printer")
//	Serial     - device serial number, if appropriate (i.e., for USB)
//	MakeModel  - device make and model (e.g., "HP DeskJet 2540")
type UnitID struct {
	DeviceName string      // Realm-unique device name
	UUID       uuid.UUID   // uuid.NilUUID if not available
	Realm      SearchRealm // Search realm
	SubRealm   string      // Backend-specific subrealm
	Kind       UnitKind    // Kind of the unit
	Serial     string      // "" if not avaliable
}

// MarshalText dumps [UnitID] as text, for [log.Object].
// It implements [encoding.TextMarshaler] interface.
func (id UnitID) MarshalText() ([]byte, error) {
	lines := make([]string, 0, 6)

	if id.DeviceName != "" {
		lines = append(lines, fmt.Sprintf("Name:   %q", id.DeviceName))
	}
	if id.UUID != uuid.NilUUID {
		lines = append(lines, fmt.Sprintf("UUID:   %s", id.UUID))
	}

	realm := id.Realm.String()
	if id.SubRealm != "" {
		realm += "-" + id.SubRealm
	}
	lines = append(lines, fmt.Sprintf("Realm:  %s", realm))

	lines = append(lines, fmt.Sprintf("Kind:   %s", id.Kind))

	if id.Serial != "" {
		lines = append(lines, fmt.Sprintf("Serial: %s", id.Serial))
	}

	return []byte(strings.Join(lines, "\n")), nil
}

// SearchRealm identifies a search realm (search domain) where
// device is found.
type SearchRealm int

// SearchRealm values:
const (
	RealmInvalid SearchRealm = iota

	RealmDNSSD // DNS-SD search
	RealmWSD   // Microsoft WS-Discovery
	RealmSNMP  // SNMP search
	RealmUSB   // USB
)

// String returns SearchRealm name.
func (realm SearchRealm) String() string {
	return realmNames[realm]
}

// realmNames contains SearchRealm names
var realmNames = map[SearchRealm]string{
	RealmInvalid: "invalid",
	RealmDNSSD:   "dnssd",
	RealmWSD:     "wsd",
	RealmSNMP:    "snmp",
	RealmUSB:     "usb",
}

// UnitKind identifies a kind of device.
type UnitKind int

// UnitKind values:
const (
	KindInvalid UnitKind = iota

	// Printers
	KindIPPPrinter       // IPP/IPPS printer
	KindLPRPrinter       // LPR protocol printer
	KindAppSocketPrinter // AppSocket (AKA JetDirect) Printer
	KindWSDPrinter       // WSD printer
	KindCUPSPrinter      // CUPS-shred printer
	KindSMBPrinter       // SMB-shred printer
	KindUSBPrinter       // USB printer
	KindUnknownPrinter   // Unknown printer

	// Scanners
	KindIPPScanner     // IPP/IPPS scanner
	KindESCLScanner    // ESCL scanner
	KindWSDScanner     // WSD scanner
	KindUnknownScanner // Unknown scanner
)

// String returns UnitKind name.
func (realm UnitKind) String() string {
	return kindNames[realm]
}

// realmNames contains SearchRealm names
var kindNames = map[UnitKind]string{
	KindInvalid:          "invalid",
	KindIPPPrinter:       "IPP printer",
	KindLPRPrinter:       "LPR printer",
	KindAppSocketPrinter: "AppSocket printer",
	KindWSDPrinter:       "WSD printer",
	KindCUPSPrinter:      "CUPS printer",
	KindSMBPrinter:       "SMB printer",
	KindUSBPrinter:       "USB printer",
	KindUnknownPrinter:   "Unknown printer",
	KindIPPScanner:       "IPP scanner",
	KindESCLScanner:      "ESCL scanner",
	KindWSDScanner:       "WSD scanner",
	KindUnknownScanner:   "Unknown scanner",
}
