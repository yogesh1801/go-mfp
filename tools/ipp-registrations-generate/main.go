// MFP - Miulti-Function Printers and scanners toolkit
// IPP registrations to Go converter.
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// The main function

package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/OpenPrinting/go-mfp/argv"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// Command describes command options
var Command = argv.Command{
	Name: "ipp-registrations-generate",
	Help: "Processing tool for IANA IPP registrations database",
	Options: []argv.Option{
		argv.Option{
			Name:     "-i",
			Aliases:  []string{"--input"},
			Help:     "input file",
			HelpArg:  "file",
			Required: true,
			Validate: argv.ValidateAny,
			Complete: argv.CompleteOSPath,
		},
		argv.Option{
			Name:     "-e",
			Aliases:  []string{"--errata"},
			Help:     "errata file (takes precedence over input)",
			HelpArg:  "file",
			Validate: argv.ValidateAny,
			Complete: argv.CompleteOSPath,
		},
		argv.Option{
			Name:      "-o",
			Aliases:   []string{"--output"},
			Help:      "output file",
			HelpArg:   "data.go",
			Required:  true,
			Singleton: true,
			Validate:  argv.ValidateAny,
			Complete:  argv.CompleteOSPath,
		},
		argv.Option{
			Name:      "-w",
			Aliases:   []string{"--allow-warnings"},
			Help:      "generate output regardless of warnings",
			Singleton: true,
		},
		argv.HelpOption,
	},
	Handler: commandHandler,
}

// commandHandler executes the command
func commandHandler(ctx context.Context, inv *argv.Invocation) error {
	// Create the database
	db := NewRegDB()

	var input, errata []xmldoc.Element
	var err error

	// Load errata files
	for _, file := range inv.Values("-e") {
		xml, err := XMLLoad(file)
		if err != nil {
			return err
		}

		errata = append(errata, xml)
	}

	// Load input file
	for _, file := range inv.Values("-i") {
		xml, err := XMLLoad(file)
		if err != nil {
			return err
		}

		input = append(input, xml)
	}

	// Process loaded files
	for _, xml := range errata {
		err := db.Load(xml, true)
		if err != nil {
			return err
		}
	}

	for _, xml := range input {
		err := db.Load(xml, false)
		if err != nil {
			return err
		}
	}

	if err := db.Finalize(); err != nil {
		return err
	}

	// Check for errors
	if len(db.Errors) != 0 {
		for _, err := range db.Errors {
			s := strings.TrimRight(err.Error(), "\n")
			fmt.Println(s)
		}

		err := fmt.Errorf("%d errors encountered", len(db.Errors))
		fmt.Println(err)

		if !inv.Flag("-w") {
			return err
		}
	}

	// Open output file
	file, _ := inv.Get("-o")
	output, err := os.OpenFile(file, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	// Generate output
	err = Output(output, db)
	output.Close()

	if err != nil {
		os.Remove(file)
		return err
	}

	return nil
}

// The main function
func main() {
	Command.Main(context.Background())
}
