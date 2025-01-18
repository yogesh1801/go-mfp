// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// ADF capabilities

package escl

import "github.com/alexpevzner/mfp/optional"

// ADF contains scanner capabilities for the automated document feeded.
type ADF struct {
	ADFSimplexInputCaps optional.Val[InputSourceCaps] // ADF simplex caps
	ADFDuplexInputCaps  optional.Val[InputSourceCaps] // ADF duples caps
	FeederCapacity      optional.Val[int]             // Feeder capacity
	ADFOptions          ADFOptions                    // ADF options
	Justification       optional.Val[Justification]   // Image justification
}
