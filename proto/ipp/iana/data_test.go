// MFP - Miulti-Function Printers and scanners toolkit
// IANA registrations for IPP
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// IPP registrations database integrity tests

package iana

import (
	"reflect"
	"sort"
	"testing"

	"github.com/OpenPrinting/go-mfp/util/generic"
)

// TestDataIntegrity performs various tests of the dataset integrity.
func TestDataIntegrity(t *testing.T) {
	// Roll over all top-level collections in the predictable way
	collections := make([]string, 0, len(Collections))
	for col := range Collections {
		collections = append(collections, col)
	}
	sort.Strings(collections)

	visited := generic.NewSet[*Attribute]()
	for _, col := range collections {
		testDataIntegrityRecursive(t, col, Collections[col], visited)
	}
}

// testDataIntegrityRecursive is the internal function that does the
// work of TestDataIntegrity, recursively over all attributes in
// the set.
func testDataIntegrityRecursive(t *testing.T,
	path string, attrs map[string]*Attribute,
	visited generic.Set[*Attribute]) {

	// Process all attributes in the predictable way
	names := make([]string, 0, len(attrs))
	for name := range attrs {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		attr := attrs[name]
		if !visited.TestAndAdd(attr) {
			// Skip already visited attributes to prevent
			// endless recursion
			continue
		}

		attrpath := path + "/" + name

		// Check that attrpath properly resolves
		//
		// Note, few attributes use members of multiple collections,
		// and collections may shadow each other, sometimes yelding
		// lookup mismatch.
		//
		// So if lookup result is reflect.DeepEqual to our
		// expectations, we still consider the lookup successful
		// and even after that there are still few exceptions.
		//
		// FIXME: this place requires more attention.
		attr2 := LookupAttribute(attrpath)
		if attr2 != attr && !reflect.DeepEqual(attr, attr2) {
			switch attrpath {
			case "Job Template/destination-uris/destination-attributes/finishings-col":
			case "Job Template/destination-uris/destination-attributes/media-col":
				// Ignore lookup mismatch for these paths
			default:
				t.Errorf("%q: doesn't resolve", attrpath)
			}
		}

		// Check that collections do really have members, while
		// non-collections doesn't
		switch {
		case attr.IsCollection() && len(attr.Members) == 0:
			t.Errorf("%q: empty collection", attrpath)

		case !attr.IsCollection() && len(attr.Members) != 0:
			t.Errorf("%q: non-collection with members", attrpath)
		}

		// Recursively visit all collection members
		if attr.IsCollection() {
			for _, mbr := range attr.Members {
				testDataIntegrityRecursive(t, attrpath, mbr, visited)
			}
		}
	}
}
