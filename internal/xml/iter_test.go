// MFP - Miulti-Function Printers and scanners toolkit
// XML mini library
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// XML iterator test

package xml

import (
	"reflect"
	"strings"
	"testing"
)

// TestIterate tests XML iterator
func TestIterate(t *testing.T) {
	root := Element{
		Name: "env",
		Attrs: []Attr{
			{"a1", "attr 1"},
			{"a2", "attr 2"},
			{"a3", "attr 3"},
		},
		Children: []*Element{
			{
				Name: "ns:el-1",
				Text: "element 1",
				Children: []*Element{
					{
						Name: "ns:el-1-1",
						Text: "element 1-1",
					},
					{
						Name: "ns:el-1-2",
						Text: "element 1-2",
					},
				},
			},
			{
				Name: "ns:el-2",
				Text: "element 2",
				Children: []*Element{
					{
						Name: "ns:el-2-1",
						Text: "element 2-1",
					},
					{
						Name: "ns:el-2-2",
						Text: "element 2-2",
					},
				},
			},
		},
	}

	expected := []string{
		"/env",
		"/env/ns:el-1",
		"/env/ns:el-1/ns:el-1-1",
		"/env/ns:el-1/ns:el-1-2",
		"/env/ns:el-2-1",
		"/env/ns:el-2-2",
	}

	present := []string{}

	iter := root.Iterate()
	for iter.Next() {
		present = append(present, iter.Path())
	}

	if !reflect.DeepEqual(expected, present) {
		t.Errorf("\nexpected:\n%s\npresent:\n%s\n",
			strings.Join(expected, "\n"),
			strings.Join(present, "\n"))
	}
}
