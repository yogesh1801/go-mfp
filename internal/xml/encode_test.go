// MFP - Miulti-Function Printers and scanners toolkit
// XML mini library
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// XML encoder test

package xml

import (
	"fmt"
	"os"
	"testing"
)

// TestEncoder tests XML encoder
func TestEncoder(t *testing.T) {
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
						Name: "ns.el-1-1",
						Text: "element 1-1",
					},
					{
						Name: "ns.el-1-2",
						Text: "element 1-2",
					},
				},
			},
			{
				Name: "ns:el-2",
				Text: "element 2",
				Children: []*Element{
					{
						Name: "ns.el-2-1",
						Text: "element 2-1",
					},
					{
						Name: "ns.el-2-2",
						Text: "element 2-2",
					},
				},
			},
		},
	}

	err := root.EncodeIndent(os.Stdout, " ")
	if err != nil {
		panic(err)
	}

	iter := root.Iterate()
	for !iter.Done() {
		cur := iter.Elem()
		fmt.Printf("%s: %s\n", iter.Path(), cur.Text)
		iter.Next()
	}
}
