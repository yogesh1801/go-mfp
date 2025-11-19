// MFP - Miulti-Function Printers and scanners toolkit
// IPP registrations to Go converter.
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Output generation

package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sort"
	"strings"
)

// Output writes RegDB to io.Writer.
func Output(w io.Writer, db *RegDB) error {
	// Write title
	buf := &bytes.Buffer{}
	buf.Write([]byte(outputTitle))

	// Output collections
	outputCollections(buf, db)

	// Save generated code to the temporary file, for formatting
	temp, err := os.CreateTemp("", "iana-ipp*.go")
	if err != nil {
		return err
	}

	name := temp.Name()
	defer temp.Close()
	defer os.Remove(name)

	_, err = io.Copy(temp, buf)
	if err != nil {
		return err
	}

	_, err = temp.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}

	// Use gofmt to format the output
	buf.Reset()
	gofmt := exec.Command("gofmt")
	gofmt.Args = append(gofmt.Args, temp.Name())
	gofmt.Stdin = temp
	gofmt.Stdout = buf

	err = gofmt.Run()
	if err != nil {
		return err
	}

	_, err = io.Copy(w, buf)
	return err
}

func outputCollections(buf *bytes.Buffer, db *RegDB) {
	for _, name := range db.CollectionNames() {
		ident := strings.Join(strings.Fields(name), "")
		fmt.Fprintf(buf, "\n")
		fmt.Fprintf(buf, "// %s is the %s attributes\n", ident, name)
		fmt.Fprintf(buf, "var %s = map[string]Attribute{\n", ident)
		outputAttributes(buf, db.Collections[name])
		fmt.Fprintf(buf, "}")
	}
}

func outputAttributes(buf *bytes.Buffer, attrs map[string]*RegDBAttr) {
	// Sort attributes by name
	names := make([]string, 0, len(attrs))
	for name := range attrs {
		names = append(names, name)
	}
	sort.Strings(names)

	// Write attributes, one by one
	for _, name := range names {
		attr := attrs[name]
		fmt.Fprintf(buf, "%q: Attribute{\n", attr.PureName())
		fmt.Fprintf(buf, "SetOf: %v,\n", attr.Syntax.SetOf)
		fmt.Fprintf(buf, "Min: %s,\n", attr.Syntax.FormatMin())
		fmt.Fprintf(buf, "Max: %s,\n", attr.Syntax.FormatMax())
		fmt.Fprintf(buf, "Tags: []goipp.Tag{")

		for _, tag := range attr.Syntax.Tags {
			fmt.Fprintf(buf, "%#v,", tag)
		}
		fmt.Fprintf(buf, "},\n")

		if len(attr.Members) > 0 {
			fmt.Fprintf(buf, "Members: map[string]Attribute{\n")
			outputAttributes(buf, attr.Members)
			fmt.Fprintf(buf, "},\n")

		}
		fmt.Fprintf(buf, "},\n")
	}
}

const outputTitle = `// MFP - Miulti-Function Printers and scanners toolkit
// IANA registrations for IPP
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// IPP registrations database
//
// THIS IS GENERATED FILE. DON'T EDIT!

package iana

import(
	"github.com/OpenPrinting/goipp"
)
`
