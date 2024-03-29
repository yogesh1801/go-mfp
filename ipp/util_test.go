// IPPX - High-level implementation of IPP printing protocol on Go
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Helper functions for tests

package ipp

import (
	"bytes"
	"fmt"
	"reflect"
	"sort"

	"github.com/OpenPrinting/goipp"
)

// testDiffStruct dumps difference between two structures of
// the same type. It is intended for testing
//
// s1 and s2 must be pointers to the structures of the
// same type
func testDiffStruct(s1, s2 interface{}) string {
	// Validate arguments
	t1, t2 := reflect.TypeOf(s1), reflect.TypeOf(s2)
	if t1 != t2 {
		err := fmt.Errorf("testDiffStruct: %s != %s", t1, t2)
		panic(err)
	}

	if t1.Kind() != reflect.Pointer || t1.Elem().Kind() != reflect.Struct {
		err := fmt.Errorf(
			"testDiffStruct: %s must be pointer to struct", t1)
		panic(err)
	}

	// Compare field by field
	buf := &bytes.Buffer{}

	stype := t1.Elem()
	struct1 := reflect.ValueOf(s1).Elem()
	struct2 := reflect.ValueOf(s2).Elem()

	for i := 0; i < stype.NumField(); i++ {
		fld := stype.Field(i)
		v1 := struct1.Field(i).Interface()
		v2 := struct2.Field(i).Interface()

		if !reflect.DeepEqual(v1, v2) {
			fmt.Fprintf(buf, "%s:\n  <<< %#v\n  >>> %#v",
				fld.Name, v1, v2)
		}
	}

	return buf.String()
}

// testDiffAttrs dumps difference between two sets of attributes
func testDiffAttrs(attrs1, attrs2 goipp.Attributes) string {
	// Make maps to access attributes by name
	m1 := make(map[string]goipp.Attribute, len(attrs1))
	m2 := make(map[string]goipp.Attribute, len(attrs2))

	for _, attr := range attrs1 {
		if _, found := m1[attr.Name]; !found {
			m1[attr.Name] = attr
		}
	}

	for _, attr := range attrs2 {
		if _, found := m2[attr.Name]; !found {
			m2[attr.Name] = attr
		}
	}

	// Compare, attribute by attribute. Build slice
	// of different attributes
	type diffItem struct {
		name   string
		v1, v2 goipp.Values
	}

	diffList := []diffItem{}

	for name, attr1 := range m1 {
		attr2 := m2[name]

		if !attr1.Equal(attr2) {
			diffList = append(diffList,
				diffItem{
					name: name,
					v1:   attr1.Values,
					v2:   attr2.Values,
				})
		}
	}

	for name, attr2 := range m2 {
		_, found := m1[name]
		if !found {
			diffList = append(diffList,
				diffItem{
					name: name,
					v1:   nil,
					v2:   attr2.Values,
				})
		}
	}

	sort.Slice(diffList, func(i, j int) bool {
		return diffList[i].name < diffList[2].name
	})

	// Generate output
	buf := &bytes.Buffer{}

	for _, diff := range diffList {
		fmt.Fprintf(buf, "%s:\n", diff.name)
		if diff.v1 != nil {
			fmt.Fprintf(buf, "  <<< %s\n", diff.v1)
		}
		if diff.v2 != nil {
			fmt.Fprintf(buf, "  >>> %s\n", diff.v2)
		}
	}

	return buf.String()
}
