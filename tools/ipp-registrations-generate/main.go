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

	"github.com/OpenPrinting/go-mfp/argv"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// Command describes command options
var Command = argv.Command{
	Name: "ipp-registrations-generate",
	Help: "Processing tool for IANA IPP registrations database",
	Options: []argv.Option{
		argv.Option{
			Name:      "-i",
			Aliases:   []string{"--input"},
			Help:      "input file",
			HelpArg:   "file",
			Required:  true,
			Singleton: true,
			Validate:  argv.ValidateAny,
			Complete:  argv.CompleteOSPath,
		},
		argv.Option{
			Name:      "-e",
			Aliases:   []string{"--errata"},
			Help:      "errata file (takes precedence over input)",
			HelpArg:   "file",
			Singleton: true,
			Validate:  argv.ValidateAny,
			Complete:  argv.CompleteOSPath,
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

	var input, errata xmldoc.Element
	var err error

	// Load errata file, if any
	if file, ok := inv.Get("-e"); ok {
		errata, err = XMLLoad(file)
		if err != nil {
			return err
		}
	}

	// Load input file
	file, _ := inv.Get("-i")
	input, err = XMLLoad(file)
	if err != nil {
		return err
	}

	// Process loaded files
	err = db.Load(errata, true)
	if err == nil {
		err = db.Load(input, false)
	}

	if err != nil {
		return err
	}

	// Check for errors
	if len(db.Errors) != 0 {
		for _, err := range db.Errors {
			fmt.Println(err)
		}

		err := fmt.Errorf("%d errors encountered", len(db.Errors))
		fmt.Println(err)

		if !inv.Flag("-w") {
			return err
		}
	}

	// Open output file
	file, _ = inv.Get("-o")
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
	Command.Main(nil)
}
