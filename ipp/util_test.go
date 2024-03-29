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
	"testing"
)

func testFoo() {
	println("FOO")
}

// testDiffStruct dumps difference between two structures of
// the same type. It is intended for testing
//
// s1 and s2 must be pointers to the structures of the
// same type
func testDiffStruct(t *testing.T, s1, s2 interface{}) string {
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
