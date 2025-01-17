// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Set of CCD feed directions

package escl

import "github.com/alexpevzner/mfp/internal/generic"

// FeedDirections contains a set of [FeedDirection]s.
type FeedDirections struct {
	generic.Bitset[FeedDirection] // Set of channels
}

// MakeFeedDirections makes [FeedDirections] from the list of [FeedDirection]s.
func MakeFeedDirections(list ...FeedDirection) FeedDirections {
	return FeedDirections{
		generic.MakeBitset(list...),
	}
}
