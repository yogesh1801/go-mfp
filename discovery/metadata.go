// MFP - Miulti-Function Printers and scanners toolkit
// Device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Device metadata

package discovery

// Metadata contains device (or unit) metadata.
//
// Please note, historically Manufacturer, Product and MakeModel come
// from different sources, so it is not guaranteed that MakeModel is
// just a concatenation of above two strings.
//
// [Backend] MUST either provide Manufacturer and Model or MakeModel
// and MAY provide all of these parameters.
type Metadata struct {
	Manufacturer string // I.e., "Hewlett Packard" or "Canon"
	Model        string // Model name
	MakeModel    string // Manufacturer + Model
}
