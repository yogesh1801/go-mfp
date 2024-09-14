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
