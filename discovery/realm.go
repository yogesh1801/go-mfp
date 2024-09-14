// MFP - Miulti-Function Printers and scanners toolkit
// Device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Device units

package discovery

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
