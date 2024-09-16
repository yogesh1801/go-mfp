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
	"strings"

	"github.com/alexpevzner/mfp/uuid"
)

// PrintUnit represents a print unit
type PrintUnit struct {
	ID        UnitID            // Unit identity
	Meta      Metadata          // Unit metadata
	Params    PrinterParameters // Printer parameters
	Endpoints []string          // URLs of printer endpoints
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
//	UnitName   - specifies a logical unit within a device (for example,
//	             queue name for LPD printer which may have multiple
//	             distinct queues). Optional
//	Realm      - search realm. Different realms are treated as
//	             independent namespaces.
//	SubRealm   - allows backend to further divide its namespace
//	             (for example, to split it between IP4/IP6)
//	Kind       - specifies device kind (e.g., "IPP printer")
//	Serial     - device serial number, if appropriate (i.e., for USB)
type UnitID struct {
	DeviceName string       // Realm-unique device name
	UUID       uuid.UUID    // uuid.NilUUID if not available
	UnitName   string       // Logical unit within a device
	Realm      SearchRealm  // Search realm
	SubRealm   string       // Backend-specific subrealm
	SvcType    ServiceType  // Service type
	SvcProto   ServiceProto // Service protocol
	Serial     string       // "" if not avaliable
}

// MarshalText dumps [UnitID] as text, for [log.Object].
// It implements [encoding.TextMarshaler] interface.
func (id UnitID) MarshalText() ([]byte, error) {
	lines := make([]string, 0, 6)

	if id.DeviceName != "" {
		lines = append(lines, fmt.Sprintf("Name:     %q", id.DeviceName))
	}
	if id.UUID != uuid.NilUUID {
		lines = append(lines, fmt.Sprintf("UUID:     %s", id.UUID))
	}
	if id.UnitName != "" {
		lines = append(lines, fmt.Sprintf("UnitName: %q", id.UnitName))
	}

	realm := id.Realm.String()
	if id.SubRealm != "" {
		realm += "-" + id.SubRealm
	}
	lines = append(lines, fmt.Sprintf("Realm:    %s", realm))

	lines = append(lines, fmt.Sprintf("Service:  %s %s",
		id.SvcProto, id.SvcType))

	if id.Serial != "" {
		lines = append(lines, fmt.Sprintf("Serial:   %s", id.Serial))
	}

	return []byte(strings.Join(lines, "\n")), nil
}
