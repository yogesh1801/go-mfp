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

// outputCollections writes top-level connections.
func outputCollections(buf *bytes.Buffer, db *RegDB) {
	for _, name := range db.CollectionNames() {
		ident := strings.Join(strings.Fields(name), "")
		fmt.Fprintf(buf, "\n")
		fmt.Fprintf(buf, "// %s is the %s attributes\n", ident, name)
		fmt.Fprintf(buf, "var %s = map[string]Attribute{\n", ident)
		outputAttributes(buf, db.Collections[name], name)
		fmt.Fprintf(buf, "}")
	}
}

// outputAttributes writes set of attributes, recursively.
func outputAttributes(buf *bytes.Buffer, attrs map[string]*RegDBAttr, path string) {
	// Sort attributes by name
	names := make([]string, 0, len(attrs))
	for name := range attrs {
		names = append(names, name)
	}
	sort.Strings(names)

	// Write attributes, one by one
	for _, name := range names {
		attr := attrs[name]
		purename := attr.PureName()
		attrpath := path + "/" + purename

		fmt.Fprintf(buf, "// %s\n", attrpath)
		fmt.Fprintf(buf, "%q: Attribute{\n", purename)
		fmt.Fprintf(buf, "SetOf: %v,\n", attr.Syntax.SetOf)
		fmt.Fprintf(buf, "Min: %s,\n", attr.Syntax.FormatMin())
		fmt.Fprintf(buf, "Max: %s,\n", attr.Syntax.FormatMax())
		fmt.Fprintf(buf, "Tags: []goipp.Tag{")

		for _, tag := range attr.Syntax.Tags {
			fmt.Fprintf(buf, "%#v,", tag)
		}
		fmt.Fprintf(buf, "},\n")

		cntMembers := len(attr.Members) + len(attr.Borrowed)
		if cntMembers > 0 {
			members := make(map[string]*RegDBAttr, cntMembers)

			for mbrame, member := range attr.Members {
				members[mbrame] = member
			}

			for mbrame, member := range attr.Borrowed {
				// Borrowed member may refer to the
				// top-level collection that includes
				// the current attribute.
				//
				// To prevent looping we skip members
				// that refer to the current attribute.
				if members[mbrame] == nil && member != attr {
					members[mbrame] = member
				}
			}

			fmt.Fprintf(buf, "Members: map[string]Attribute{\n")
			outputAttributes(buf, members, attrpath)
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
