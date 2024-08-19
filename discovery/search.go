// MFP - Miulti-Function Printers and scanners toolkit
// Device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Search parameters

package discovery

// SearchRealm identifies a search realm (search domain).
type SearchRealm int

// SearchRealm values:
const (
	SearchRealmDNSSD SearchRealm = iota // DNS-SD search
	SearchRealmWSD                      // Microsoft WS-Discovery
	SearchRealmSNMP                     // SNMP search
	SearchRealmUSB                      // USB
)

// SearchDeviceKind identifies a kind of device
type SearchDeviceKind int

// SearchDeviceKind values:
const (
	SearchDeviceInvalid SearchDeviceKind = iota

	// Printers
	SearchDeviceIPPPrinter       // IPP/IPPS printer
	SearchDeviceLPDPrinter       // LPD protocol printer
	SearchDeviceAppSocketPrinter // AppSocket (AKA JetDirect) Printer
	SearchDeviceWSDPrinter       // WSD printer
	SearchDeviceCUPSPrinter      // CUPS-shred printer
	SearchDeviceSMBPrinter       // SMB-shred printer
	SearchDeviceUSBPrinter       // USB printer
	SearchDeviceUnknownPrinter   // Unknown printer

	// Scanners
	SearchDeviceIPPScanner     // IPP/IPPS scanner
	SearchDeviceESCLScanner    // ESCL scanner
	SearchDeviceWSDScanner     // WSD scanner
	SearchDeviceUnknownScanner // Unknown scanner
)
