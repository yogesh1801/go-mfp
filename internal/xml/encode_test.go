// MFP - Miulti-Function Printers and scanners toolkit
// XML mini library
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// XML encoder test

package xml

import (
	"os"
	"testing"
)

// TestEncoder tests XML encoder
func TestEncoder(t *testing.T) {
	elements := []*Element{
		{
			Name: "env",
			Children: []*Element{
				{
					Name: "ns:el-1",
					Text: "element 1",
				},
			},
		},
	}

	err := EncodeIndent(os.Stdout, elements, " ")
	if err != nil {
		panic(err)
	}
}
