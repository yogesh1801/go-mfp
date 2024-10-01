// MFP - Miulti-Function Printers and scanners toolkit
// XML mini library
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Package documentation

package xml

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

// TestDecode tests Decode function
func TestDecode(t *testing.T) {
	ns := map[string]string{
		`http://example.com/a`: `a`,
		`http://example.com/b`: `b`,
		`http://example.com/c`: `c`,
	}

	in := `` +
		`<?xml version="1.0" ?>` +
		`<env xmlns:ns-a="http://example.com/a" xmlns:ns-b="http://example.com/b" xmlns:ns-c="http://example.com/c">` +
		`<ns-a:elem-a>body a</ns-a:elem-a>` +
		`<ns-b:elem-b>body b` +
		`<ns-b:nested-1>nested body 1</ns-b:nested-1>` +
		`<ns-b:nested-2>nested body 2` +
		`<ns-b:nested-2-1>nested body 2-1</ns-b:nested-2-1>` +
		`</ns-b:nested-2>` +
		`</ns-b:elem-b>` +
		`<ns-c:elem-c>body c</ns-c:elem-c>` +
		`<ns-d:elem-d>body d</ns-d:elem-d>` +
		`</env>` +
		``

	expect := `` +
		`/env: ""` + "\n" +
		`  /env/a:elem-a: "body a"` + "\n" +
		`  /env/b:elem-b: "body b"` + "\n" +
		`    /env/b:elem-b/b:nested-1: "nested body 1"` + "\n" +
		`    /env/b:elem-b/b:nested-2: "nested body 2"` + "\n" +
		`      /env/b:elem-b/b:nested-2/b:nested-2-1: "nested body 2-1"` + "\n" +
		`    /env/b:elem-b/b:nested-2/b:nested-2-1: "nested body 2-1"` + "\n" +
		`  /env/b:elem-b/b:nested-1: "nested body 1"` + "\n" +
		`  /env/b:elem-b/b:nested-2: "nested body 2"` + "\n" +
		`    /env/b:elem-b/b:nested-2/b:nested-2-1: "nested body 2-1"` + "\n" +
		`  /env/b:elem-b/b:nested-2/b:nested-2-1: "nested body 2-1"` + "\n" +
		`  /env/c:elem-c: "body c"` + "\n" +
		`  /env/-:elem-d: "body d"` + "\n" +
		`/env/a:elem-a: "body a"` + "\n" +
		`/env/b:elem-b: "body b"` + "\n" +
		`  /env/b:elem-b/b:nested-1: "nested body 1"` + "\n" +
		`  /env/b:elem-b/b:nested-2: "nested body 2"` + "\n" +
		`    /env/b:elem-b/b:nested-2/b:nested-2-1: "nested body 2-1"` + "\n" +
		`  /env/b:elem-b/b:nested-2/b:nested-2-1: "nested body 2-1"` + "\n" +
		`/env/b:elem-b/b:nested-1: "nested body 1"` + "\n" +
		`/env/b:elem-b/b:nested-2: "nested body 2"` + "\n" +
		`  /env/b:elem-b/b:nested-2/b:nested-2-1: "nested body 2-1"` + "\n" +
		`/env/b:elem-b/b:nested-2/b:nested-2-1: "nested body 2-1"` + "\n" +
		`/env/c:elem-c: "body c"` + "\n" +
		`/env/-:elem-d: "body d"` + "\n" +
		``

	out, err := Decode(ns, bytes.NewReader([]byte(in)))
	if err != nil {
		t.Errorf("%s", err)
		return
	}

	buf := &bytes.Buffer{}
	dump(out, buf, "")

	if buf.String() != expect {
		t.Errorf("decode mismatch\nexpected: %s:\npresent: %s",
			expect, buf.String())
	}
}

// dump dumps decoded elements into io.Writer
func dump(elements []*Element, out io.Writer, indent string) {
	for _, elm := range elements {
		fmt.Fprintf(out, "%s%s: %q\n", indent, elm.Path, elm.Text)
		dump(elm.Children, out, indent+"  ")
	}
}
