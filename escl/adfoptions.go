// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Set of scan color modes

package escl

import "github.com/alexpevzner/mfp/internal/generic"

// ADFOptions contains a set of [ADFOption]s.
type ADFOptions struct {
	generic.Bitset[ADFOption]
}

// MakeADFOptions makes [ADFOptions] from the list of [ADFOption]s.
func MakeADFOptions(list ...ADFOption) ADFOptions {
	return ADFOptions{
		generic.MakeBitset(list...),
	}
}
