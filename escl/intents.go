// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Set of CCD intents

package escl

import "github.com/alexpevzner/mfp/internal/generic"

// Intents contains a set of [Intent]s.
type Intents struct {
	generic.Bitset[Intent]        // Set of channels
	Default                Intent // Default channel
}

// MakeIntents makes [Intents] from the list of [Intent]s.
func MakeIntents(list ...Intent) Intents {
	return Intents{
		generic.MakeBitset(list...),
		UnknownIntent,
	}
}
