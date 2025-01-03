// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Set of scan color modes

package escl

import "strings"

// ADFOptions contains a set (bitmask) of [ADFOption]s.
type ADFOptions int

// MakeADFOptions makes [ADFOptions] from the list of [ADFOption]s.
func MakeADFOptions(list ...ADFOption) ADFOptions {
	var opts ADFOptions

	for _, opt := range list {
		opts.Add(opt)
	}

	return opts
}

// String returns a string representation of the [ADFOptions],
// for debugging.
func (opts ADFOptions) String() string {
	modes := [...]ADFOption{DetectPaperLoaded, SelectSinglePage, Duplex}
	s := make([]string, 0, len(modes))

	for _, opt := range modes {
		if opts.Contains(opt) {
			s = append(s, opt.String())
		}
	}

	return strings.Join(s, ",")
}

// Add adds [ADFOption] to the set.
func (opts *ADFOptions) Add(opt ADFOption) *ADFOptions {
	*opts |= 1 << opt
	return opts
}

// Del deletes [ADFOption] from the set.
func (opts *ADFOptions) Del(opt ADFOption) *ADFOptions {
	*opts &^= 1 << opt
	return opts
}

// Contains reports if [ADFOption] exists in the set.
func (opts ADFOptions) Contains(opt ADFOption) bool {
	return (opts & (1 << opt)) != 0
}
