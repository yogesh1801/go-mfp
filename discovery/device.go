// MFP - Miulti-Function Printers and scanners toolkit
// Device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Common device information

package discovery

import "net/netip"

// Device consist of the multiple functional units. There are
// two types of units:
//   - [PrintUnit]
//   - [ScanUnit]
//
// Each unit has its unique [UnitID], the combination of parameters,
// that uniquely identifies the unit.
type Device struct {
	MakeModel   string       // Device make and model
	Addrs       []netip.Addr // Device's IP addresses
	PrintUnits  []PrintUnit  // Print units
	ScanUnits   []ScanUnit   // Scan units
	FaxoutUnits []FaxoutUnit // Faxout units
}

// device is the internal representation of the Device
type device struct {
	units []unit       // Device's units
	addrs []netip.Addr // Device's IP addresses
}

// Export exports device as Device
func (dev device) Export() Device {
	out := Device{Addrs: dev.addrs}

	makeModelFrom := RealmInvalid

	for _, un := range dev.units {
		// Save device MakeModel
		setMakeModel := false

		switch un.id.Realm {
		case RealmUSB:
			setMakeModel = true
		case RealmDNSSD:
			setMakeModel = makeModelFrom != RealmUSB
		default:
			setMakeModel = out.MakeModel == ""
		}

		if setMakeModel {
			out.MakeModel = un.meta.MakeModel
			makeModelFrom = un.id.Realm
		}

		// Save unit
		exp := un.Export()
		switch exp := exp.(type) {
		case PrintUnit:
			out.PrintUnits = append(out.PrintUnits, exp)
		case ScanUnit:
			out.ScanUnits = append(out.ScanUnits, exp)
		case FaxoutUnit:
			out.FaxoutUnits = append(out.FaxoutUnits, exp)
		}
	}

	return out
}
