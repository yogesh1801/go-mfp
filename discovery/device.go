// MFP - Miulti-Function Printers and scanners toolkit
// Device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Common device information

package discovery

// Device consist of the multiple functional units. There are
// two types of units:
//   - [PrintUnit]
//   - [ScanUnit]
//
// Each unit has its unique [UnitID], the combination of parameters,
// that uniquely identifies the unit.
type Device struct {
	PrintUnits  []PrintUnit
	ScanUnits   []ScanUnit
	FaxoutUnits []FaxoutUnit
}

// device is the internal representation of the Device
type device struct {
	units []unit
}

// Export exports device as Device
func (dev device) Export() Device {
	out := Device{}

	for _, un := range dev.units {
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
