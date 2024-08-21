// MFP - Miulti-Function Printers and scanners toolkit
// Device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Common device information

package discovery

import "github.com/alexpevzner/mfp/uuid"

// DeviceID contains information that identifies a device.
//
// Please note, depending on a discovery protocol being used, not
// all the fields of the following structure will have any sense.
// And although this package attempts to be protocol-neutral, this
// fact is hard to ignore here.
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
type DeviceID struct {
	DeviceName string      // Realm-unique device name
	Realm      DeviceRealm // Search realm
	IfIdx      int         // For multicast-based network discovery
	Kind       DeviceKind  // Kind of device
	UUID       uuid.UUID   // uuid.NilUUID if not available
	Serial     string      // "" if not avaliable
	MakeModel  string      // Just for user information
}

// DeviceRealm identifies a search realm (search domain) where
// device is found.
type DeviceRealm int

// DeviceRealm values:
const (
	RealmInvalid DeviceRealm = iota

	RealmDNSSD // DNS-SD search
	RealmWSD   // Microsoft WS-Discovery
	RealmSNMP  // SNMP search
	RealmUSB   // USB
)

// DeviceKind identifies a kind of device.
type DeviceKind int

// SearchDeviceKind values:
const (
	KindInvalid DeviceKind = iota

	// Printers
	KindIPPPrinter       // IPP/IPPS printer
	KindLPDPrinter       // LPD protocol printer
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
