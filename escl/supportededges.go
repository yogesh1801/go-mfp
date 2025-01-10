// MFP - Miulti-Function Printers and scanners toolkit
// eSCL core protocol
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Set of CCD image edges

package escl

import "github.com/alexpevzner/mfp/internal/generic"

// SupportedEdges contains a set of [SupportedEdge]s.
type SupportedEdges struct {
	generic.Bitset[SupportedEdge] // Set of edges
}

// MakeSupportedEdges makes [SupportedEdges] from the list of [SupportedEdge]s.
func MakeSupportedEdges(list ...SupportedEdge) SupportedEdges {
	return SupportedEdges{
		generic.MakeBitset(list...),
	}
}
