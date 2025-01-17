// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Platen capabilities

package escl

import "github.com/alexpevzner/mfp/optional"

// Platen contains scanner capabilities for the Platen source.
type Platen struct {
	PlatenInputCaps optional.Val[InputSourceCaps] // Platen capabilities
}
