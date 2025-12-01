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

	visited := generic.NewSet[*DefAttr]()
	for _, col := range collections {
		testDataIntegrityRecursive(t, col, Collections[col], visited)
	}
}

// testDataIntegrityRecursive is the internal function that does the
// work of TestDataIntegrity, recursively over all attributes in
// the set.
func testDataIntegrityRecursive(t *testing.T,
	path string, attrs map[string]*DefAttr,
	visited generic.Set[*DefAttr]) {

	// Process all attributes in the predictable way
	names := make([]string, 0, len(attrs))
	for name := range attrs {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		def := attrs[name]
		if !visited.TestAndAdd(def) {
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
		def2 := LookupAttribute(attrpath)
		if exceptions.Contains(attrpath) {
			if def2 != nil {
				t.Errorf("%q: must not resolve", attrpath)
			}
		} else if def2 != def && !reflect.DeepEqual(def, def2) {
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
		case def.IsCollection() && len(def.Members) == 0:
			t.Errorf("%q: empty collection", attrpath)

		case !def.IsCollection() && len(def.Members) != 0:
			t.Errorf("%q: non-collection with members", attrpath)
		}

		// Recursively visit all collection members
		if def.IsCollection() {
			for _, mbr := range def.Members {
				testDataIntegrityRecursive(t, attrpath, mbr, visited)
			}
		}
	}
}
