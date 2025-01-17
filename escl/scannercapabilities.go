// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Scanner capabilities

package escl

import (
	"github.com/alexpevzner/mfp/optional"
	"github.com/alexpevzner/mfp/uuid"
)

// ScannerCapabilities defines the scanner capabilities.
//
// eSCL Technical Specification, 8.1.4.
type ScannerCapabilities struct {
	Version         Version                 // eSCL protocol version
	MakeAndModel    optional.Val[string]    // Device make and model
	SerialNumber    optional.Val[string]    // Device-unique serial number
	Manufacturer    optional.Val[string]    // Device manufacturer
	UUID            optional.Val[uuid.UUID] // Device UUID
	AdminURI        optional.Val[string]    // Configuration mage URL
	IconURI         optional.Val[string]    // Device icon URL
	SettingProfiles []SettingProfile        // Common settings profs
	Platen          optional.Val[Platen]    // Platen capabilities
	Camera          optional.Val[Camera]    // Camera capabilities
	ADF             optional.Val[ADF]       // ADF capabilities
}
