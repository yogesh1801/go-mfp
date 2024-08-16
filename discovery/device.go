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
//	HostName   - host name, in the DNS-SD sense. E.g., "KM7B6A91.local".
//	UUID       - device UUID
type DeviceID struct {
	DeviceName string
	HostName   string
	UUID       uuid.UUID
	Scope      SearchScope
}
