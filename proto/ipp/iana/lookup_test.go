// MFP - Miulti-Function Printers and scanners toolkit
// IANA registrations for IPP
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// LookupAttribute tests

package iana

import (
	"sort"
	"testing"
)

// TestLookupAttributeNotFound tests for attribute names that must not be found
func TestLookupAttributeNotFound(t *testing.T) {
	// Prepare test cases
	paths := []string{
		"",                                      // Empty string
		"Not Exist/attribute",                   // Not existing collection
		"Job Template/not-exist",                // Attribute doesn't exist
		"Job Template/print-accuracy/not-exist", // Member doesn't exist
		"Job Template/overrides/overrides",      // Attribute includes self as a member
	}

	for col := range Collections {
		paths = append(paths, col)
	}

	exceptions.ForEach(func(path string) {
		paths = append(paths, path)
	})

	sort.Strings(paths)

	// Run tests
	for _, path := range paths {
		def := LookupAttribute(path)
		if def != nil {
			t.Errorf("%q: LookupAttribute must return nil", path)
		}
	}
}
