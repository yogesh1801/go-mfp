// MFP - Miulti-Function Printers and scanners toolkit
// XML mini library
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// XML encoder test

package xml

import (
	"testing"
)

// TestEncoder tests XML encoder
func TestEncoder(t *testing.T) {
	ns := Namespace{
		{URL: "http://example.com/ns", Prefix: "ns"},
		{URL: "https://example.com/ns", Prefix: "ns"},
		{URL: "http://example.com/ns1", Prefix: "ns1"},
		{URL: "http://example.com/ns2", Prefix: "ns2"},
	}

	root := Element{
		Name: "env",
		Attrs: []Attr{
			{"a1", "attr 1"},
			{"a2", "attr 2"},
			{"a3", "attr 3"},
		},
		Children: []Element{
			{
				Name: "ns:el-1",
				Text: "element 1",
				Children: []Element{
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
				Children: []Element{
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

	compact := `<?xml version="1.0"?><env xmlns:ns="http://example.com/ns" a1="attr 1" a2="attr 2" a3="attr 3"><ns:el-1>element 1<ns:el-1-1>element 1-1</ns:el-1-1><ns:el-1-2>element 1-2</ns:el-1-2></ns:el-1><ns:el-2>element 2<ns:el-2-1>element 2-1</ns:el-2-1><ns:el-2-2>element 2-2</ns:el-2-2></ns:el-2></env>`
	indent :=
		`<?xml version="1.0"?>
<env xmlns:ns="http://example.com/ns" a1="attr 1" a2="attr 2" a3="attr 3">
  <ns:el-1>element 1
    <ns:el-1-1>element 1-1</ns:el-1-1>
    <ns:el-1-2>element 1-2</ns:el-1-2>
  </ns:el-1>
  <ns:el-2>element 2
    <ns:el-2-1>element 2-1</ns:el-2-1>
    <ns:el-2-2>element 2-2</ns:el-2-2>
  </ns:el-2>
</env>
`

	out := root.EncodeString(ns)
	if out != compact {
		t.Errorf("EncodeString:\n"+
			"expected:\n%s\n"+
			"present:\n%s\n",
			compact, out)
	}

	out = root.EncodeIndentString(ns, "  ")
	if out != indent {
		t.Errorf("EncodeIndentString failed")
	}
}
